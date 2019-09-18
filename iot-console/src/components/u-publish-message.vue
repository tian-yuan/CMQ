<template>
    <u-modal title="发布消息" :visible.sync="show" size="huge" okButton="X">
        <u-notice icon="warning">如果该 Topic 正在被使用，请谨慎操作，以防出现异常。这里发布的消息不会被服务端订阅。</u-notice>
        <u-form ref="form" gap="large">
            <u-form-item label="Topic" layout="block" style="word-break:break-all;">
                {{ Topic }}
            </u-form-item>

            <u-form-item label="消息内容" required placement="bottom"
                ref="content"
                :rules="rules.Content">
                <u-textarea size="huge"
                    resize="none"
                    v-model="model.Content"
                    placeholder="1000字符以内"
                    maxlength="1000"
                    @keypress="onKeypress(model.Content)"></u-textarea>
            </u-form-item>

            <u-form-item label="Qos">
                <u-radios v-model="model.Qos">
                    <u-radio :label="0">0</u-radio>
                    <u-radio :label="1">1</u-radio>
                </u-radios>
            </u-form-item>

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
                            立即发布
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
import { Base64 } from 'js-base64';

export default {
    mixins: [Modal],
    props: {
        ProductKey: String,
        Topic: String,
    },
    data() {
        return {
            model: {
                ProductKey: this.ProductKey,
                Topic: this.Topic,
                Qos: 0, // Qos 默认选中 0
                Content: '',
            },
            rules: {
                Content: this.$rules.type().parse([
                    'required ::消息内容不得为空',
                ]),
            },
        };
    },
    created() {
        ['ProductKey', 'Topic', 'Qos'].forEach((key) => {
            this.$watch(() => this[key], (val) => {
                this.model[key] = val;
            });
        });
    },
    methods: {
        // @tmp 临时处理，CloudUI支持后，移除该逻辑
        onKeypress(value) {
            const formItemVm = this.$refs.content;
            if (formItemVm && value.length === 1000) {
                formItemVm.color = 'error';
                formItemVm.currentMessage = '1000字符以内';
            }
        },
        submit() {
            const model = JSON.parse(JSON.stringify(this.model));
            model.Content = Base64.encode(model.Content);
            return iothubService.PublishMessage({
                body: model,
                config: { noAlert: true },
            }).then(() => {
                this.$emit('ok');
                this.close();
            });
        },
    },
};
</script>
