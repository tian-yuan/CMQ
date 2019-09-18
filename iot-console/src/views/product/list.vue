<template>
    <u-table-view :data="list" :loading="loading" layout="fixed" block-name="table">
        <u-table-view-column title="产品名称">
            <div slot-scope="{row}" class="f-toe">
                <u-link :to="{path: '/iothub/product/detail/info', query: {ProductKey: row.ProductKey}}">
                    <span :title="row.ProductName">{{ row.ProductName }}</span>
                </u-link>
            </div>
        </u-table-view-column>
        <u-table-view-column title="ProductKey" label="ProductKey"></u-table-view-column>
        <u-table-view-column title="设备数" label="DeviceCount"></u-table-view-column>
        <u-table-view-column title="创建时间">
            <div slot-scope="{row}">
                <span>{{ row.CreateAt | timeFormat }}</span>
            </div>
        </u-table-view-column>
        <u-table-view-column title="操作">
            <div slot-scope="{row}">
                <u-link-list>
                    <u-link-list-item @click="modify(row)">设置</u-link-list-item>

                    <u-link-list-item v-if="row.canDelete" @click="deleteItem(row)">删除</u-link-list-item>
                    <u-link-list-item v-else disabled v-tooltip="'本产品中有接入的设备，不可删除'">删除</u-link-list-item>
                </u-link-list>
            </div>
        </u-table-view-column>
        <div slot="no-data-text" v-if="!loadError">
            <template v-if="Keyword">没有搜索到相关内容</template>
            <template v-else>
                还没有任何产品，现在就<u-link @click="modify">立即创建</u-link>一个吧
            </template>
        </div>
        <div slot="no-data-text" v-else>
            获取数据失败，请<u-link @click="refresh">重试</u-link>
        </div>
        <u-product-modal
            :show-modal.sync="showModal"
            :ProductName="crtProduct.ProductName"
            :Description="crtProduct.Description"
            :ProductKey="crtProduct.ProductKey"
            @ok="refresh">
        </u-product-modal>
    </u-table-view>
</template>

<script>
import iothubService from '@/services/index';
import UProductModal from '@/components/u-product-modal.vue';
import simpleSearch from '@/utils/mixins/list/simple-search';
import List from '@/blocks/list.vue';

export default {
    components: {
        UProductModal,
    },
    extends: List,
    tplMixins: [List],
    mixins: [simpleSearch({
        tagCachePage: ['/iothub/product/detail'],
        placeholder: '输入产品名称',
    })],
    data() {
        const self = this;
        return {
            create: {
                text: '创建产品',
                click() {
                    self.modify();
                },
            },
            outQuota: '产品配额已经用完，' + this.countEnoughLabel(),
            needRestorePage: true,
            showModal: false,
            crtProduct: {
                ProductName: '',
                Description: '',
                ProductKey: undefined,
            },
        };
    },
    methods: {
        modify(item = { ProductName: '', Description: '' }) {
            const { ProductName, Description, ProductKey } = item;
            this.crtProduct = { ProductName, Description, ProductKey };
            this.showModal = true;
        },
        loadList() {
            if (this.isBeforeKeywordLoaded && this.isBeforeKeywordLoaded())
                return true;

            const { Keyword } = this;
            const query = this.getFormForOAI({ Keyword });
            return iothubService.QueryProductList({
                query,
            })
                .then((res) => {
                    this.total = res.TotalCount;
                    this.list = res.ProductInfoList;
                });
        },
        loadQuota() {
            iothubService.QueryProductQuota()
                .then(({ Quota }) => {
                    this.quota = Quota;
                });
        },
        deleteItem(item) {
            const { ProductKey } = item;
            return this.$confirm({
                content: '确定删除该产品吗?',
                submit() {
                    return iothubService.DeleteProduct({
                        query: { ProductKey },
                        config: { noAlert: true },
                    });
                },
            });
        },
    },
};
</script>
