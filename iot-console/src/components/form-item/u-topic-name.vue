<template>
    <u-form-item ref="formItem" label="Topic 类" required layout="block"
        :rules="rules"
        placement="bottom">
        <template slot="description">
            {{ prefix }}
            <u-addon>
                Topic 格式必须以“/”进行分层，区分每个类目。前三个类目已规定好，第一个代表产品标识 ProductKey，第二个${deviceName} 通配 DeviceName，第三个 user 用来标识自定义的 Topic 类。
            </u-addon>
        </template>
        <u-input size="huge"
            v-model="crtValue"
            :placeholder="placeholder"
            maxlength="63"
            maxlength-message="不得超过63个字符">
        </u-input>
    </u-form-item>
</template>

<script>
import iothubService from 'module/iothub/services/index';

export default {
    props: {
        value: { type: String, default: '' },
        ProductKey: String,
        Privilege: String,
        oldTopicName: String,
        oldPrivilege: String,
    },
    data() {
        const prefix = `${this.ProductKey}/\${deviceName}/user/`;
        return {
            prefix,
            crtValue: this.value.replace(prefix, ''),
            repeatRule: {
                trigger: 'blur',
                message: '该 Topic 类已存在',
                validator: (rull, value, callback) => {
                    const { ProductKey, oldTopicName, prefix } = this;
                    value = `${prefix}${value}`;
                    if (value === oldTopicName) {
                        callback();
                        return;
                    }

                    iothubService.CheckTopicName({
                        query: {
                            ProductKey,
                            TopicName: value,
                        },
                    }).then((res) => {
                        if (res.Status === 'EXISTED')
                            callback(new Error());
                        else
                            callback();
                    });
                },
            },
        };
    },
    computed: {
        placeholder() {
            const { Privilege } = this;
            return Privilege === 'SUB' ? '1-63位字母、数字、"_"、"+"、"#"或"/"组成，以"/"进行分层且不得以其结尾，"#"仅可为结尾，"+"需单独为一层' : '1-63位字母、数字、"_"或"/"组成，以"/"进行分层且其后不可为空';
        },
        rules() {
            const { Privilege, repeatRule } = this;

            if (Privilege === 'SUB')
                return this.$rules.type().parse([
                    'required ::Topic 类不得为空',
                    { pattern: /^[a-zA-Z0-9_+#/]*$/, trigger: 'input+blur', message: '字母、数字、"_"、"+"、"#" 或"/"组成' },
                    { pattern: /[^/]$/, trigger: 'blur', message: '不得以"/"结尾' },
                    { pattern: /^[^#]*#?$/, trigger: 'blur', message: '"#"仅可为结尾' },
                    { trigger: 'input+blur', message: '"+"需单独为一层', validator(rull, value, callback) {
                        const regs = [/[^/]\+/, /\+[^/]/];
                        const error = regs.some((reg) => reg.test(value)) ? new Error() : undefined;
                        callback(error);
                    } },
                    repeatRule,
                ]);
            else
                return this.$rules.type().parse([
                    'required ::Topic 类不得为空',
                    { pattern: /^[a-zA-Z0-9_/]*$/, trigger: 'input+blur', message: '字母、数字、"_" 或"/"组成' },
                    { pattern: /[^/]$/, trigger: 'blur', message: '不得以"/"结尾' },
                    repeatRule,
                ]);
        },
    },
    watch: {
        Privilege(val) {
            this.$nextTick(() => {
                const vm = this.$refs.formItem;
                vm && vm.validate();
            });
        },
        value(val) {
            const { prefix } = this;
            this.crtValue = val.replace(prefix, '');
        },
        crtValue(val) {
            const { prefix } = this;
            this.$emit('input', `${prefix}${val}`);
        },
    },
};
</script>

