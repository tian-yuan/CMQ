<template>
    <div>
        <u-head-card :title="instance.DeviceName">
            <div slot="logo">
                <i :class="$style.logo"></i>
            </div>
            <div slot="info">
                <p><label>所属产品：</label>
                    <u-text-hide :text="instance.ProductName" inline style="max-width:400px;"></u-text-hide>
                </p>
                <p><label>最近上线时间：</label>{{ instance.LastActiveAt | dateFormat }}</p>
            </div>
        </u-head-card>
        <div>
            <u-tabs router>
                <u-tab title="详细信息" :to="{path:'/iothub/device/detail/info', query: {ProductKey, DeviceName}}"></u-tab>
                <u-tab title="设备 Topic 信息" :to="{path:'/iothub/device/detail/topic', query: {ProductKey, DeviceName}}"></u-tab>
            </u-tabs>
            <div>
                <router-view></router-view>
            </div>
        </div>
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

export default {
    mixins: [subscribe, Auth],
    data() {
        return {
            instance: {},
            ProductKey: this.$route.query.ProductKey,
            DeviceName: this.$route.query.DeviceName,
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
            iothubService.QueryDeviceInfo({
                query: {
                    ProductKey: this.ProductKey,
                    DeviceName: this.DeviceName,
                },
            }).then((res) => {
                res.loadDone = true;
                this.instance = res.DeviceInfo;
            });
        },
    },
};
</script>
