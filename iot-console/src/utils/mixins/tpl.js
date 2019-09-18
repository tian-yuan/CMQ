export default {
    beforeCreate() {
        const getRender = this.$options.getRender;
        let tplMixins = this.$options.tplMixins || [];
        const map = {};
        const collect = function (component) {
            if (!component) {
                return;
            }
            if (Array.isArray(component)) {
                component.forEach(collect);
                return;
            }
            const blockName = ((component.data || {}).attrs || {})['block-name'];
            if (blockName) {
                map[blockName] = component;
                return true;
            }
            return false;
        };
        // 把当前的 render 也参与计算
        tplMixins = tplMixins.concat({
            render: this.$options.render,
        });
        this.$options.render = function (...args) {
            const solts = tplMixins.map((tplMixin) => tplMixin.render && tplMixin.render.apply(this, args));
            solts.forEach((solt) => {
                if (solt && !collect(solt) && solt.children)
                    collect(solt.children);
            });
            return getRender.apply(this, [map, ...args]);
        };
    },
};
