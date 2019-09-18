export default {
    install(Vue) {
        const bus = new Vue();
        bus.on = bus.$on;
        bus.once = bus.$once;
        bus.off = bus.$off;
        bus.emit = bus.$emit;
        Object.defineProperty(Vue.prototype, '$bus', {
            get() {
                return bus;
            },
        });
    },
};
