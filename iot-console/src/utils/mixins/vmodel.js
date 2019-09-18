export default {
    props: ['value'],
    data() {
        return {
            crtValue: this.value,
        };
    },
    watch: {
        value(val) {
            this.crtValue = val;
        },
        crtValue(val) {
            this.$emit('input', val);
        },
    },
};
