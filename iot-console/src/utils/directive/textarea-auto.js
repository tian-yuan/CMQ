import event from '../event';
export default {
    bind(el, binding, vnode) {
        const events = ['change', 'cut', 'paste', 'drop', 'input', 'focus', 'keydown', 'keyup'];
        let timer;
        const cb = function () {
            clearTimeout(timer);
            timer = setTimeout(() => {
                if (el) {
                    el.style.height = '1px';
                    if (!el.value) {
                        el.setAttribute('rows', '1');
                    }
                    el.style.height = el.scrollHeight + 'px';
                    el.scrollTop = el.scrollHeight;
                }
            }, 30);
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
