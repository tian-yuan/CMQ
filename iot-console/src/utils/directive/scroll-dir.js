import event from '../event';
export default {
    bind(el, binding, vnode) {
        const events = ['scroll'];
        let scrollTop = 0;
        let timer;
        const resetScrollTop = function (scrollTop) {
            el.scrollTop = el.scrollTop + scrollTop;
        };
        const cb = function (event) {
            if (timer)
                clearTimeout(timer);
            timer = setTimeout(() => {
                const stance = scrollTop - this.scrollTop;
                const height = (window.getComputedStyle(this).height.replace('px', '') - 0) || 0;
                const safeHeight = 100;
                if (stance > 0) {
                    if (this.scrollTop <= (binding.value.up || safeHeight)) {
                        if (parseInt(this.scrollTop) === 0)
                            this.scrollTop = 1;
                        binding.value.scroll('up', (p) => {
                            resetScrollTop(p);
                        });
                    }
                } else if (stance < 0) {
                    if (this.scrollHeight - height - this.scrollTop <= (binding.value.down || safeHeight)) {
                        if (parseInt(this.scrollTop) === parseInt(this.scrollHeight - height))
                            this.scrollTop = this.scrollHeight - height - 1;
                        binding.value.scroll('down', (p) => {
                            resetScrollTop(p);
                        });
                    }
                }
                scrollTop = this.scrollTop;
            }, 17);
        };
        events.forEach((eventItem) => {
            event.on(el, eventItem, cb);
        });
        vnode.context.$customEvents = {
            events,
            cb,
        };
    },
    unbind(el, binding, vnode) {
        const { events, cb } = vnode.context.$customEvents || {};
        events.forEach((eventItem) => {
            event.off(el, eventItem, cb);
        });
    },
};
