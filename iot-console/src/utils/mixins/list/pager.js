/**
 * 列表分页的功能MinXin，列表页引入这个minxin，提供分页功能数据
 */

const getPage = function (curLen, limit, curPage, total) {
    if (limit * curPage + curLen < total)
        return (curPage + 1) * limit;

    return curPage * limit + curLen;
};

export default {
    data() {
        return {
            form: {
                offset: 0,
                limit: 20,
            },
            total: 0,
            page: 1,
        };
    },
    computed: {
        totalPage() {
            return Math.ceil(this.total / this.form.limit) || 1;
        },
    },
    methods: {
        refresh() {
            this.form.offset = 0;
            this.loadList();
        },
        changePage($event) {
            this.form.offset = ($event.page - 1) * this.form.limit;
            this.page = $event.page;
            this.loadList();
        },
        getForm() {
            return Object.assign({}, this.form);
        },
        setTotal(total) {
            this.total = total;
        },
        updateOffset(listLength) {
            const { offset, limit } = this.form;
            const total = this.total;
            this.form.offset = getPage(listLength, limit, Math.ceil(offset / limit), total);
            this.page = Math.ceil(this.form.offset / limit);
        },
    },
};
