import Vue from 'vue';
Vue.use({
    install(Vue) {
        Vue.prototype.$modal = {
            show(name, data) {
                const modal = this.map[name];
                modal.open(data);
            },
            hide(name, data) {
                const modal = this.map[name];
                modal.close(data);
            },
            map: {},
            // 以下逻辑记录当前存在的Modal个数, 用以处理存在多个Modal时关闭某个Modal的情况
            // 在总体迁移到Vue后，可以去掉这些逻辑
            count: 0,
            addModal() {
                this.count++;
            },
            delModal() {
                this.count--;
            },
            hasModal() {
                return !!this.count;
            },
        };
    },
});
export default {
    beforeCreate() {
        const modalName = this.$attrs['modal-name'];
        if (modalName) {
            this.$modal.map[modalName] = this;
        }
    },
    destroyed() {
        const modalName = this.$attrs['modal-name'];
        if (modalName && this.$modal.map[modalName] === this) {
            delete this.$modal.map[modalName];
        }
    },
    props: {
        showModal: { type: Boolean, default: false },
    },
    data() {
        return {
            show: this.showModal || false,
        };
    },
    resetModal: true, // 设置是否需要reset
    watch: {
        showModal(value) {
            this.show = value;
        },
        show(value) {
            if (value)
                this.$modal.addModal();
            else
                this.$modal.delModal();

            if (this.$options.resetModal && !value)
                this.resetModal();
            this.$emit('update:showModal', value);
            this.dispatchParent('handleMask', {
                action: this.$modal.hasModal() ? 'open' : 'close',
            });
        },
    },
    methods: {
        open() {
            this.show = true;
        },
        close() {
            this.show = false;
        },
        resetModal() {
            const staticData = this.$options.staticData || [];
            const resetData = this.$options.data.apply(this);
            staticData.push('show');
            for (const name of staticData)
                delete resetData[name];
            Object.assign(this.$data, resetData);
        },
    },
};
