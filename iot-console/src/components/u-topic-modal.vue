<template>
    <u-modal :title="title" :visible.sync="show" size="huge" okButton="x">
        <u-form ref="form" gap="large">
            <u-privilege v-model="model.Operation"></u-privilege>

            <u-topic-name
                v-model="model.TopicName"
                :ProductKey="model.ProductKey"
                :Privilege="model.Operation"
                :oldTopicName="TopicName"
                :oldPrivilege="Operation"></u-topic-name>

            <u-description
                v-model="model.Description"
                placement="bottom"></u-description>
        </u-form>
        <template slot="foot">
            <u-submit-button
                auto-focus
                :formRef="() => $refs.form"
                :formChanged="formChanged"
                :click="submit.bind(this)">
                <template slot-scope="scope">
                    <u-linear-layout>
                        <u-button color="primary"
                            :icon="scope.submitting ? 'loading' : ''"
                            :disabled="scope.submitting"
                            @click="scope.submit">确定</u-button>
                        <u-button @click="close">取消</u-button>
                    </u-linear-layout>
                </template>
            </u-submit-button>
        </template>
    </u-modal>
</template>

<script>
import Modal from 'mixins/modal/base';
import iothubService from 'module/iothub/services/index';
import UTopicName from 'module/iothub/components/form-item/u-topic-name.vue';
import UDescription from '@/components/common/form/form-item/u-description.vue';
import UPrivilege from 'module/iothub/components/form-item/u-privilege.vue';

export default {
    components: {
        UTopicName, UDescription, UPrivilege,
    },
    mixins: [Modal],
    props: {
        ProductKey: { type: String, default: '' },
        TopicName: { type: String, default: '' },
        TopicId: { type: String, default: undefined }, // 创建时，没有TopicId
        Operation: { type: String, default: 'PUB' }, // 可选值[SUB, PUB, ALL]，SUB-订阅，PUB-发布，ALL-订阅和发布
        Description: { type: String, default: '' },
    },
    data() {
        const model = {
            Qos: 1, // 默认填写 1
            TopicType: 'CUSTOM', // 默认填写 CUSTOM
        };
        const propList = ['ProductKey', 'TopicName', 'TopicId', 'Operation', 'Description'];
        propList.forEach((key) => {
            model[key] = this[key];
        });
        return {
            model,
            propList,
        };
    },
    computed: {
        formChanged() {
            const {
                TopicName: oldTopicName,
                Description: oldDescription,
                Operation: oldOperation,
                isModify,
            } = this;
            const { TopicName, Description, Operation } = this.model;
            return !isModify
                || TopicName !== oldTopicName
                || Description !== oldDescription
                || Operation !== oldOperation;
        },
        isModify() {
            const { TopicId } = this;
            return !!TopicId;
        },
        title() {
            const { isModify } = this;
            return isModify ? '修改 Topic 类' : '定义 Topic 类';
        },
    },
    created() {
        this.propList.forEach((key) => {
            this.$watch(() => this[key], (val) => {
                this.model[key] = val;
            });
        });
    },
    methods: {
        submit() {
            const model = JSON.parse(JSON.stringify(this.model));

            const action = this.isModify ? 'UpdateTopicClass' : 'CreateTopicClass';
            return iothubService[action]({
                query: model,
                config: { noAlert: true },
            }).then(() => {
                this.$emit('ok');
                this.close();
            });
        },
    },
};
</script>
