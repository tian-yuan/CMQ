<template>
    <div>
        <u-info-list-group title="基本信息" column="3">
            <u-info-list-item label="产品名称">{{ instance.ProductName }}</u-info-list-item>
            <u-info-list-item label="ProductKey">
                {{ instance.ProductKey }}
                <u-copy :message="instance.ProductKey" placement="right" class="f-ml10"></u-copy>
            </u-info-list-item>
            <u-info-list-item label="DeviceName">
                {{ instance.DeviceName }}
                <u-copy :message="instance.DeviceName" placement="right" class="f-ml10"></u-copy>
            </u-info-list-item>
            <u-info-list-item label="DeviceSecret">
                {{ DeviceSecret }}
                <u-link @click="showSecret = !showSecret">{{ showSecret ? '隐藏' : '显示' }}</u-link>
                <u-copy
                    v-if="showSecret"
                    :message="instance.DeviceSecret"
                    placement="right"
                    class="f-ml10"></u-copy>
            </u-info-list-item>
            <u-info-list-item label="类型">设备</u-info-list-item>
            <u-info-list-item label="类型">{{ instance.State | deviceState }}</u-info-list-item>
            <u-info-list-item label="最近上线时间">{{ instance.LastActiveAt | dateFormat }}</u-info-list-item>
            <u-info-list-item label="添加时间">{{ instance.CreateAt | dateFormat }}</u-info-list-item>
        </u-info-list-group>
    </div>
</template>

<style module>
</style>

<script>
import subscribe from '@/utils/mixins/subscribe';
import filters from 'module/iothub/utils/filters';

export default {
    filters,
    mixins: [subscribe],
    subscribes: {
        loadInstance(data) {
            this.instance = data;
        },
    },
    data() {
        return {
            instance: {},
            ClusterId: this.$route.query.ClusterId,
            networkId: undefined,
            subnetId: undefined,
            showSecret: false,
        };
    },
    computed: {
        DeviceSecret() {
            let secret = this.instance.DeviceSecret || '';
            const { showSecret } = this;
            if (!showSecret) {
                secret = secret.replace(/./g, '*');
            }
            return secret;
        },
    },
};
</script>
