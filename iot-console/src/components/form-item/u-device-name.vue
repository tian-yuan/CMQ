<template>
    <u-form-item
        label="DeviceName" required
        placement="bottom"
        :rules="rules">
        <u-input size="huge"
            placeholder="1-63位字母、数字、&quot;-&quot;、&quot;_&quot;、&quot;@&quot;、&quot;.&quot;或&quot;:&quot;组成"
            maxlength="63"
            maxlength-message="不得超过63个字符"
            v-model="crtValue"></u-input>
    </u-form-item>
</template>

<script>
import vmodel from 'mixins/vmodel';
import iothubService from 'module/iothub/services/index';

export default {
    mixins: [vmodel],
    data() {
        return {
            rules: this.$rules.type().parse([
                'required ::DeviceName 不得为空',
                { pattern: /^[\w-@.:]+$/, trigger: 'input+blur', message: '字母、数字、"-"、"_"、"@"、"."或":"组成' },
                { trigger: 'blur', message: '该 DeiviceName 已存在', validator(rull, value, callback) {
                    iothubService.CheckDeviceName({
                        query: {
                            DeviceName: value,
                        },
                    }).then(({ Status }) => {
                        const error = Status === 'EXISTED' ? new Error() : undefined;
                        callback(error);
                    });
                } },
            ]),
        };
    },
};
</script>

