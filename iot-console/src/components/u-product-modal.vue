<template>
    <u-modal :title="isModify ? '设置产品' : '创建产品'" :visible.sync="show" size="huge"
        okButton="立即创建" cancelButton="取消">
        <u-form ref="form" gap="large">
            <u-form-item label="产品名称" required
                :rules="rules.ProductName"
                placement="bottom">
                <u-input size="huge"
                    v-model="model.ProductName"
                    placeholder="1-63位字母、数字或&quot;-&quot;组成，以字母开头，字母或数字结尾"
                    maxlength="63"
                    maxlength-message="不得超过63个字符">
                </u-input>
            </u-form-item>
            <u-form-item label="描述"
                :rules="rules.Description"
                placement="bottom">
                <u-input size="huge"
                    v-model="model.Description"
                    placeholder="100字符以内"
                    maxlength="100"
                    maxlength-message="100字符以内">
                </u-input>
            </u-form-item>
        </u-form>
        <template slot="foot">
            <u-submit-button
                :autoFocus="true"
                :formRef="() => $refs.form"
                :formChanged="formChanged"
                :click="submit.bind(this)">
                <template slot-scope="scope">
                    <u-linear-layout>
                        <u-button color="primary"
                            :icon="scope.submitting ? 'loading' : ''"
                            :disabled="scope.submitting"
                            @click="scope.submit">
                            {{ isModify ? '提交设置' : '立即创建' }}
                        </u-button>
                        <u-button @click="showModal=false">取消</u-button>
                    </u-linear-layout>
                </template>
            </u-submit-button>
        </template>
    </u-modal>
</template>

<script>
import Modal from '@/utils/mixins/modal/base';
import iothubService from '@/services/index';

export default {
    mixins: [Modal],
    props: {
        ProductName: { type: String, default: '' },
        Description: { type: String, default: '' },
        ProductKey: { type: String, default: undefined }, // 设置时，需要传ProductKey
    },
    data() {
        return {
            model: {
                ProductName: this.ProductName,
                Description: this.Description,
                ProductKey: this.ProductKey,
            },
            rules: {
                ProductName: this.$rules.type('string').parse([
                    'required ::产品名称不得为空',
                    'startazAZ',
                    'withazAZ09-',
                    'endazAZ09',
                    ['exist', null, () => this.nameList2],
                ]),
            },
            nameList: [],
        };
    },
    computed: {
        formChanged() {
            const { ProductName: oldProductName, Description: oldDescription, isModify } = this;
            const { ProductName, Description } = this.model;
            return !isModify
                || ProductName !== oldProductName
                || Description !== oldDescription;
        },
        isModify() {
            return !!this.ProductKey;
        },
        // 设置产品时，允许使用原产品名
        nameList2() {
            let { isModify, nameList, ProductName } = this;
            if (isModify) {
                nameList = nameList.filter((name) => name !== ProductName);
            }
            return nameList;
        },
    },
    watch: {
        ProductName(val) {
            this.model.ProductName = val;
            this.isModify = !!val;
        },
        Description(val) {
            this.model.Description = val;
        },
        ProductKey(val) {
            this.model.ProductKey = val;
        },
        show(val) {
            if (val)
                this.getNameList();
        },
    },
    methods: {
        getNameList() {
            iothubService.QueryProductListAll()
                .then((res) => {
                    this.nameList = res.ProductInfoList.map((item) => item.ProductName);
                });
        },
        submit() {
            const model = JSON.parse(JSON.stringify(this.model));

            const action = this.isModify ? 'UpdateProduct' : 'CreateProduct';
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
