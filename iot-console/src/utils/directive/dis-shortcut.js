import event from '../event';
export default {
    bind(el, binding, vnode) {
        const events = 'keydown';
        const cb = function (e) {
            let cur = e.target;
            const map = [37, 40, 38, 39, 8];
            const tagNames = ['INPUT', 'TEXTAREA'];
            if (map.includes(e.which) && !tagNames.includes(cur.tagName)) {
                while (cur && cur !== el) {
                    cur = cur.parentElement;
                }
                if (cur) {
                    e.preventDefault && e.preventDefault();
                    e.stopPropagation && e.stopPropagation();
                    return false;
                }
            }
        };
        vnode.context.$customEvents = {
            events,
            cb,
        };
        event.on(document, events, cb);
    },
    unbind(el, binding, vnode) {
        const { events, cb } = vnode.context.$customEvents || {};
        event.off(document, events, cb);
    },
};
