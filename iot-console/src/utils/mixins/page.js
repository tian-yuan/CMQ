import i18n from '@/utils/i18n';
/**
 * mixins提供了：
 *  1. 对列表接口 loadList 方法进行了增强，可获取状态 loading 、loadError。
 *     需要定义 loadList 方法，返回值需要是 promise 或者 true（true 代表延迟调用 loadList，调用时机需用户自行把握），
 *     同时需要对 list total 进行赋值。
 *  2. 配额是否有剩余 hasQuota （默认是 true，在有数据的情况下会返回实际剩余配额）
 *     需要定义 loadQuota 方法（可选）主要是获取配额，同时需要对 quota 赋值配额。
 *  3. 分页控件事件响应 changePage，提供 totalPage page 取值
 *  4. 提供了获取调用远程接口参数的方法，支持传参，eg: getFormForOAI({other: 1})
 * getForm（用于接口发送 offset、limit），getFormForOAI（用于接口发送 Offset、Limit），getFormForPage（用于接口发生 pageNum、pageSize）
 *  5. 提供安全删除实例的方法，用于（最后一页只有一条的情况），同时会自动在实例上添加 deleting 属性
 *     需要定义 deleteItem 方法，返回值需要是 promise 或者 true（true 代表延迟调用 deleteItem）
 *  6. 提供前端分页模式 localPage，如需刷新，调用 refresh 即可
 *  7. 提供分页细粒度模式 needRestorePage，需要在路由配置相关信息（包括 name、meta.children、meta.noPageInfo），
 *
 */
import storage from '@/utils/storage/localStorage';

export default {
    loadListWhenCreated: true, // 弹窗里的列表不用created时loadList
    created() {
        this.loadListSource = this.loadList; // cache this
        this.deleteItemSource = this.deleteItem; // cache this
        this.deleteItem = this.deleteItemWrap;
        this.updateItemSource = this.updateItem; // cache this
        this.updateItem = this.updateItemWrap;
        this.loadList = this.loadListWrap;
        // 还原分页粒度和页码
        if (this.needRestorePage) {
            const currentPageName = this.$route.name;
            const name = currentPageName.substring(currentPageName.indexOf('.') + 1);
            this.urlPageName = `${name}.page`;
            this.limitStorageId = `page-limit-${currentPageName}`;
            let limit = storage.get(this.limitStorageId);
            const urlPageData = this.$route.query[this.urlPageName] || '';
            // localStorage里有limit，或者route里带有page信息，说明设置过，需要设置url
            if (limit || urlPageData) {
                limit = this.checkLimit(+limit);
                const pageData = urlPageData.split('_');
                const urlPage = +pageData[1] || this.page;
                let urlLimit = pageData[0] ? this.checkLimit(+pageData[0]) : '';
                urlLimit = urlLimit || limit; // url上的limit优先
                this.updateUrl({ limit: urlLimit, page: urlPage }, this.$options.loadListWhenCreated);
                storage.set(this.limitStorageId, urlLimit);
            } else {
                this.$options.loadListWhenCreated && this.loadList();
                this.$options.loadListWhenCreated && this.loadQuota && this.loadQuota();
            }
        } else {
            this.$options.loadListWhenCreated && this.loadList();
            this.$options.loadListWhenCreated && this.loadQuota && this.loadQuota();
        }
    },
    beforeRouteLeave(to, from, next) {
        if (this.needRestorePage) {
            // 非跳到列表页的子页面，删除当页的page参数
            const children = from.meta.children || [];
            const toSubPage = children.some((item) => to.path.startsWith(item));
            // 创建页是子页面，创建后需要跳转到第一页，不需要保留分页信息
            const noPageInfo = from.meta.noPageInfo || [];
            const toNoPageInfo = noPageInfo.some((item) => to.path.startsWith(item));
            if (!toSubPage || toNoPageInfo) {
                this.resetPageUrl();
            }
        }
        next();
    },
    beforeRouteUpdate(to, from, next) {
        // 处理浏览器回退历史记录时列表刷新问题
        if (this.needRestorePage) {
            const query = to.query;
            const pageData = query[this.urlPageName] ? query[this.urlPageName].split('_') : [];
            const limit = this.checkLimit(+pageData[0]);
            const page = +pageData[1] || 1; // 如果page没有设置，回退到第一页
            if (limit !== this.form.limit || page !== this.page) {
                const dataTemp = {
                    limit,
                    page,
                };
                this.updatePageData(dataTemp);
                this.refresh();
            }
        }
        next();
    },
    data() {
        return {
            ...this.reset(true),
            ...this.initLoadStatus('load'),
            list: [],
            originList: null,
            quota: undefined,
            totalAll: 0,
            smoothLoad: false, // 是否刷新不loading: 实时刷新新数据，但不显示loading状态。在外层控制loading状态
        };
    },
    computed: {
        hasQuota() {
            const list = this.list;
            const quota = this.quota;
            const total = this.total;
            const totalAll = this.totalAll;
            // 改这里之前，请先看一下最上面的描述
            if ((total || list || totalAll) && quota !== undefined)
                return quota - (totalAll || total || list.length);
            return true;
        },
        hasQuotab() { // 配额为负数的情况
            return this.hasQuota > 0;
        },
    },
    watch: {
        total(total) {
            const { limit } = this.form;
            this.totalPage = Math.ceil(total / limit) || 1;
        },
        'form.limit'(limit) {
            this.totalPage = Math.ceil(this.total / limit) || 1;
        },
    },
    methods: {
        /**
         * 刷新时，是否显示加载中间态
         * @param {*} smoothLoad    若传该参数，会优先使用该参数，不使用this.smoothLoad
         */
        refresh(params = {}) {
            if (this.localPage) {
                this.originList = null;
            }
            this.loadList(params);
            this.loadQuota && this.loadQuota();
        },
        reset(isInit) {
            const limitList = [20, 50, 100];
            const data = {
                form: {
                    offset: 0,
                    limit: limitList[0],
                },
                page: 1,
                totalPage: 1,
                total: 0,
                limitList,
            };
            if (isInit)
                return data;
            else {
                Object.assign(this.form, data.form);
                this.page = data.page;
                this.totalPage = data.totalPage;
                this.total = data.total;
            }
        },
        resetAll() {
            this.total = 0;
            this.reset();
        },
        previousPage() {
            this.page -= 1;
            this.changePage({ page: this.page });
        },
        changePage($event) {
            if (this.needRestorePage) {
                this.updateUrl({ page: $event.page });
            } else {
                this.page = $event.page;
                this.form.offset = ($event.page - 1) * this.form.limit;
                this.loadList();
            }
        },
        $getLocalList() {
            if (this.originList) {
                const { pageNum, pageSize } = this.getFormForPage();
                this.list = this.originList.slice((pageNum - 1) * pageSize, pageNum * pageSize);
            }
        },
        loadListWrap(params = {}) {
            if (this.localPage && this.originList) {
                this.$getLocalList();
                return;
            }
            const req = this.wrap('loadListSource', params);
            if (req) {
                const smoothLoad = params.smoothLoad !== undefined ? params.smoothLoad : this.smoothLoad;
                this.addLoadStatus(() => {
                    if (!smoothLoad) {
                        this.list = [];
                    } else {
                        this.loading = false;
                    }
                    return req;
                }, 'load').then(() => {
                    if (this.localPage) {
                        this.$getLocalList();
                    }
                    if (this.list && !this.list.length && this.form.offset !== 0) {
                        if (this.needRestorePage) {
                            this.updateUrl({ page: 1 });
                        } else {
                            this.form.offset = 0;
                            this.page = 1;
                            this.loadList();
                        }
                    } else if (this.list) {
                        const { offset, limit } = this.form;
                        this.page = Math.ceil((this.list.length + offset) / limit);
                    }
                });
            }
        },
        wrap(source, params) {
            const req = this[source](params);
            if (req === true) {
                // 不执行
                return false;
            }
            if (DEV) {
                if (!req || !req.then) {
                    console.error('必须返回 promise 或者 true (不执行)');
                    return false;
                }
            }
            return req;
        },
        deleteItemWrap(item) {
            if (item.deleting)
                return;
            const req = this.wrap('deleteItemSource', item);
            if (req) {
                this.$set(item, 'deleting', true);
                req.then(() => {
                    item.deleting = false;
                    if (this.totalAll) {
                        this.totalAll -= 1;
                    }
                    if (this.total) {
                        this.total -= 1;
                    }
                    // 删除最后一页，自动回到前一页
                    if (this.list.length === 1) {
                        if (this.needRestorePage) {
                            const page = this.page - 1;
                            this.updateUrl({ page: page <= 0 ? 1 : page });
                            return;
                        } else {
                            const prev = this.page - 2;
                            this.form.offset = (prev < 0 ? 0 : prev) * this.form.limit;
                        }
                    }
                    this.refresh();
                }, () => {
                    item.deleting = false;
                });
            }
        },
        updateItemWrap(item) {
            if (item.updating)
                return;
            const req = this.wrap('updateItemSource', item);
            if (req) {
                this.$set(item, 'updating', true);
                req.then(() => {
                    item.updating = false;
                    this.refresh();
                }, () => {
                    item.updating = false;
                });
            }
        },
        getForm(ops) {
            return Object.assign({}, this.form, ops);
        },
        getFormForOAI(ops) {
            const { offset, limit } = this.form;
            return Object.assign({
                Offset: offset,
                Limit: limit,
            }, ops);
        },
        getFormForPage(ops) {
            const { offset, limit } = this.form;
            return Object.assign({
                pageNum: (Math.ceil(offset / limit) + 1) || 1,
                pageSize: limit,
            }, ops);
        },
        /**
         * 修改分页粒度
         */
        changeLimit($event) {
            // limit变化后，分页总数也会跟着变，所以需要重置total，使得分页控件不显示
            this.resetToPage1({ limit: $event.value });
            storage.set(this.limitStorageId, $event.value);
        },
        /**
         * 更新列表中的limit和page
         * 统一从url中拿取
         */
        updatePageData(data) {
            this.form.limit = +data.limit;
            this.page = +data.page;
            this.form.offset = (this.page - 1) * this.form.limit;
        },
        /**
         * limit or page 更改的时候，更新url
         */
        updateUrl(data = {}, loadListWhenCreated = true) {
            const dataTemp = Object.assign({
                limit: this.form.limit,
                page: this.page,
            }, data);
            const query = Object.assign({}, this.$route.query);
            const urlPageName = this.urlPageName;
            query[urlPageName] = [dataTemp.limit, dataTemp.page].join('_');
            this.changeQueryUrl(query);
            this.updatePageData(dataTemp);
            loadListWhenCreated && this.refresh();
        },
        /**
         * 子列表退出时，需要将page的最后一位清除
         */
        resetPageUrl() {
            const query = this.$route.query;
            if (query[this.urlPageName]) {
                delete query[this.urlPageName];
            }
        },
        checkLimit(limit) {
            if (this.limitList.includes(limit))
                return limit;
            return this.limitList[0];
        },
        resetToPage1(data = {}) {
            // this.totalPage = 0;
            this.total = 0;
            this.updateUrl(Object.assign(data, { page: 1 }));
        },
    },
};
