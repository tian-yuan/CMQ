export default {
    update(el, binding, vnode) {
        if (binding.value) {
            el.scrollIntoView(false);
        }
    },
};
