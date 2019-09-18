<template>
    <div>
        <u-info-list-group title="基本信息" column="3">
            <u-info-list-item label="产品名称">{{ instance.ProductName }}</u-info-list-item>
            <u-info-list-item label="ProductKey">
                {{ instance.ProductKey }}
                <u-copy :message="instance.ProductKey" placement="right" class="f-ml10"></u-copy>
            </u-info-list-item>
            <u-info-list-item label="设备数">{{ instance.DeviceCount }}个</u-info-list-item>
            <u-info-list-item label="创建时间">{{ instance.CreateAt | dateFormat }}</u-info-list-item>
            <u-info-list-item label="描述" column="1">{{ instance.Description | cross }}</u-info-list-item>
        </u-info-list-group>
    </div>
</template>

<style module>
</style>

<script>
import subscribe from '@/utils/mixins/subscribe';
import filters from 'module/tsdb/utils/filters';
import azfilter from '@/utils/filters/vpc';

export default {
    filters: Object.assign(filters, azfilter),
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
        };
    },
};
</script>
