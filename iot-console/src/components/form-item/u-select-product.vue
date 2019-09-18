<template>
    <u-form-item label="选择产品" required
        placement="bottom"
        layout="block"
        :rules="rules">
        <u-select size="medium"
            v-model="crtValue"
            :disabled="filterList.length === 0">
            <u-select-item>
                {{ filterList.length === 0 ? '暂无可选择的产品' : '选择产品' }}
            </u-select-item>
            <u-select-item
                v-for="item in filterList"
                :key="item"
                :value="item.ProductKey">
                {{ item.ProductName }}
            </u-select-item>
        </u-select>
        <u-desc v-if="crtValue">所选产品还可以添加 {{ remain | cross }} 个设备</u-desc>
    </u-form-item>
</template>

<script>
import vmodel from 'mixins/vmodel';
import iothubService from 'module/iothub/services/index';

export default {
    mixins: [vmodel],
    data() {
        return {
            list: [],
            rules: this.$rules.type().parse([
                'required ::产品为必填项',
            ]),
            quota: undefined, // 设配配额，对与所有产品都一样
            remain: undefined, // 所选产品还可以添加的设备数
        };
    },
    computed: {
        filterList() {
            const { list, quota } = this;
            if (quota)
                return list.filter((item) => item.DeviceCount < quota);
            else
                return [];
        },
    },
    watch: {
        crtValue(val) {
            if (val)
                this.getQuota(val);
        },
    },
    created() {
        this.getList();
    },
    methods: {
        getList() {
            iothubService.QueryProductListAll()
                .then((res) => {
                    this.list = res.ProductInfoList;

                    if (this.list[0]) {
                        const ProductKey = this.list[0].ProductKey;
                        this.getQuota(ProductKey);
                    }
                });
        },
        getQuota(ProductKey) {
            this.remain = undefined;
            iothubService.QueryDeviceQuota({
                query: {
                    ProductKey,
                },
            }).then(({ Quota, UsedQuota }) => {
                this.remain = Quota - UsedQuota;
                this.quota = Quota;
                this.$emit('update:remain', this.remain);
            });
        },
    },
};
</script>

