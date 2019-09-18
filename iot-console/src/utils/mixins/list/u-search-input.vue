<template>
    <u-input
        :placeholder="placeholder"
        search="left"
        size="medium"
        v-model.trim="crtValue"
        @keyup.enter="$emit('enter', crtValue)"></u-input>
</template>

<script>
// simple-search.js依赖该组件
import vmodel from '@/utils/mixins/vmodel';

export default {
    mixins: [vmodel],
    props: {
        placeholder: { type: String, default: '' },
    },
};
</script>

