<template>
    <u-table-view :data="list" :loading="loading" layout="fixed" block-name="table"
        @filter-change="onFilterChange">
        <u-table-view-column title="DeviceName">
            <div slot-scope="{row}" class="f-toe">
                <u-link :to="{path: '/iothub/device/detail/info', query: {ProductKey: row.ProductKey, DeviceName: row.DeviceName}}">
                    <span :title="row.DeviceName">{{ row.DeviceName }}</span>
                </u-link>
            </div>
        </u-table-view-column>
        <u-table-view-column
            title="所属产品"
            label="ProductName"
            ellipsis
            width="30%"
            filter
            :value="filterProductKey"
            :options="options"
            :filter-method="() => true"></u-table-view-column>
        <u-table-view-column title="类型" width="10%">设备</u-table-view-column>
        <u-table-view-column title="最近上线时间">
            <div slot-scope="{row}">
                <span>{{ row.LastActiveAt | dateFormat }}</span>
            </div>
        </u-table-view-column>
        <u-table-view-column title="操作">
            <div slot-scope="{row}">
                <u-link-list>
                    <u-link-list-item @click="deleteItem(row)">删除</u-link-list-item>
                </u-link-list>
            </div>
        </u-table-view-column>
        <div slot="no-data-text" v-if="!loadError">
            <template v-if="Keyword">没有搜索到相关内容</template>
            <template v-else-if="filterProductKey">{{ ProductMap[filterProductKey] }} 产品中没有设备</template>
            <template v-else>
                还没有任何设备，现在就<u-link @click="add">立即添加</u-link>一个吧
            </template>
        </div>
        <div slot="no-data-text" v-else>
            获取数据失败，请<u-link @click="refresh">重试</u-link>
        </div>
        <u-device-modal
            :show-modal.sync="showModal"
            :isBatch="isBatch"
            @ok="refresh">
        </u-device-modal>
    </u-table-view>
</template>

<script>
import iothubService from 'module/iothub/services/index';
import List from '@/blocks/list.vue';
import UDeviceModal from 'module/iothub/components/u-device-modal.vue';
import simpleSearch from 'mixins/list/simple-search';

export default {
    components: {
        UDeviceModal,
    },
    extends: List,
    tplMixins: [List],
    mixins: [simpleSearch({
        tagCachePage: ['/iothub/device/detail'],
        placeholder: '输入 DeviceName',
    })],
    data() {
        return {
            create: {
                text: '添加设备',
                click: this.add.bind(this, false),
            },
            actions: [{
                text: '批量添加设备',
                options: {
                    color: 'primary',
                    icon: 'create',
                },
                handle: this.add.bind(this, true),
            }],
            needRestorePage: true,
            showModal: false,
            filterProductKey: '',
            ProductMap: {}, // { ProductKey: ProductName}
            isBatch: false, // 是否批量添加
            options: [{ name: '全部', value: '' }],
        };
    },
    created() {
        this.loadProductList();
    },
    methods: {
        add(isBatch = false) {
            this.isBatch = isBatch;
            this.showModal = true;
        },
        onFilterChange(event) {
            this.filterProductKey = event.value;

            if (this.filterProductKey)
                this.updateKeyword();

            this.resetToPage1();
        },
        loadProductList() {
            return iothubService.QueryProductListAll()
                .then((res) => {
                    const list = res.ProductInfoList;
                    const options = index.map(({ ProductName, ProductKey }) => {
                        this.ProductMap[ProductKey] = ProductName;
                        return {
                            name: ProductName,
                            value: ProductKey,
                        };
                    });
                    this.options = [...this.options, ...options];
                });
        },
        loadList() {
            if (this.isBeforeKeywordLoaded && this.isBeforeKeywordLoaded())
                return true;

            const { Keyword, filterProductKey: ProductKey } = this;
            const query = this.getFormForOAI({ Keyword, ProductKey });
            return iothubService.QueryDeviceList({
                query,
            })
                .then((res) => {
                    this.total = res.TotalCount;
                    this.list = res.DeviceInfoList;
                });
        },
        deleteItem(item) {
            const { ProductKey, DeviceName } = item;
            return this.$confirm({
                content: '确定删除该设备吗?',
                submit: () => iothubService.DeleteDevice({
                    query: { ProductKey, DeviceName },
                    config: { noAlert: true },
                }),
            });
        },
    },
};
</script>
