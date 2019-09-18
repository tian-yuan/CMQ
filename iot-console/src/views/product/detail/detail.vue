<template>
    <div>
        <u-head-card :title="instance.ProductName">
            <div slot="logo">
                <i :class="$style.logo"></i>
            </div>
            <div slot="info">
                <p><label>设备数：</label>{{ instance.DeviceCount }}个</p>
                <p><label>创建时间：</label>{{ instance.CreateAt | dateFormat }}</p>
            </div>
            <div slot="act">
                <u-button
                    @click="showModal=true">设置
                </u-button>
            </div>
        </u-head-card>
        <div>
            <u-tabs router>
                <u-tab title="详细信息" :to="{path:'/iothub/product/detail/info', query: {ProductKey}}"></u-tab>
                <u-tab title="Topic 信息" :to="{path:'/iothub/product/detail/topic', query: {ProductKey}}"></u-tab>
            </u-tabs>
            <div>
                <router-view></router-view>
            </div>
        </div>
        <u-product-modal
            :show-modal.sync="showModal"
            :ProductName="instance.ProductName"
            :Description="instance.Description"
            :ProductKey="ProductKey"
            @ok="load">
        </u-product-modal>
    </div>
</template>

<style module>
.logo::after{
    icon-font: url('@/assets/icons/svg/font/home-iothub-thin.svg');
}
</style>

<script>
import subscribe from '@/utils/mixins/subscribe';
import iothubService from 'module/iothub/services/index';
import Auth from 'mixins/auth';
import UProductModal from 'module/iothub/components/u-product-modal.vue';

export default {
    components: {
        UProductModal,
    },
    mixins: [subscribe, Auth],
    data() {
        return {
            instance: {},
            ProductKey: this.$route.query.ProductKey,
            showModal: false,
        };
    },
    publishs: {
        loadInstance: 'instance',
    },
    created() {
        this.load();
    },
    methods: {
        load() {
            iothubService.QueryProduct({
                query: {
                    ProductKey: this.ProductKey,
                },
            }).then((res) => {
                res.loadDone = true;
                this.instance = res.ProductInfo;
            });
        },
    },
};
</script>
