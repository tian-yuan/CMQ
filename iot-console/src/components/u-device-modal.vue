<template>
    <u-modal :title="title" :visible.sync="show" :size="isBatch ? 'large' : 'huge'" okButton="X">
        <u-form ref="form" gap="large">
            <u-select-product
                v-model="model.ProductKey"
                @update:remain="remain = $event"></u-select-product>

            <u-device-name v-if="!isBatch"
                v-model="model.DeviceNames[0]"></u-device-name>

            <template v-else>
                <u-form-item label="添加方式">
                    <u-capsules value="auto">
                        <u-capsule value="auto">自动生成</u-capsule>
                    </u-capsules>
                </u-form-item>

                <u-integer
                    v-model="model.Count"
                    label="设备数量"
                    description="系统会自动成全局唯一的 DeviceName"
                    placement="bottom"
                    :min="1"
                    :max="remain"
                    unit="个"></u-integer>
            </template>
        </u-form>
        <template slot="foot">
            <u-submit-button
                :autoFocus="true"
                :formRef="() => $refs.form"
                :click="submit.bind(this)">
                <template slot-scope="scope">
                    <u-linear-layout>
                        <u-button color="primary"
                            :icon="scope.submitting ? 'loading' : ''"
                            :disabled="scope.submitting"
                            @click="scope.submit">
                            立即添加
                        </u-button>
                        <u-button @click="showModal=false">取消</u-button>
                    </u-linear-layout>
                </template>
            </u-submit-button>
        </template>
    </u-modal>
</template>

<script>
import Modal from 'mixins/modal/base';
import iothubService from 'module/iothub/services/index';
import UInteger from '@/components/common/form/form-item/u-integer.vue';
import USelectProduct from 'module/iothub/components/form-item/u-select-product.vue';
import UDeviceName from 'module/iothub/components/form-item/u-device-name.vue';

export default {
    components: {
        UInteger, USelectProduct, UDeviceName,
    },
    mixins: [Modal],
    props: {
        isBatch: { type: Boolean, default: false },
    },
    data() {
        return {
            model: {
                ProductKey: undefined,
                Count: 1,
                DeviceNames: [],
            },
            remain: 1,
        };
    },
    computed: {
        title() {
            const { isBatch } = this;
            return isBatch ? '批量添加设备' : '添加设备';
        },
    },
    methods: {
        submit() {
            const { ProductKey, Count, DeviceNames } = this.model;
            const { isBatch } = this;

            return iothubService.RegisterDevices({
                query: { ProductKey },
                body: isBatch ? { Count } : { Count, DeviceNames },
                config: { noAlert: true },
            }).then(() => {
                this.$emit('ok');
                this.close();
            });
        },
    },
};
</script>
