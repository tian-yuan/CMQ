<template>
    <u-form-item label="操作权限">
        <u-capsules v-model="crtValue">
            <u-capsule
                v-for="item in privilegeList"
                :key="item.key"
                :value="item.key">{{ item.label }}</u-capsule>
        </u-capsules>
    </u-form-item>
</template>

<script>
import { config } from 'module/iothub/utils/config';
import vmodel from 'mixins/vmodel';

export default {
    mixins: [vmodel],
    data() {
        return {
            privilegeList: config.privilegeList,
        };
    },
};
</script>

