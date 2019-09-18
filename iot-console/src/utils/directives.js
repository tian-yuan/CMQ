import event from './event';
import Bridge from '@/utils/Bridge';
import clampLib from '@/lib/clamp.js';

export const repeatClick = {
    bind(el, binding, vnode) {
        const wait = +binding.arg || 400;
        const interval = 100;
        const handler = vnode.context[binding.expression];
        let pressing = false;
        let timeout;

        const fn = () => {
            if (!pressing)
                return;

            handler();
            timeout = setTimeout(fn, interval);
        };

        event.on(el, 'mousedown', (e) => {
            if (e.button !== 0)
                return;

            event.once(document, 'mouseup', () => pressing = false);
            clearTimeout(timeout);
            pressing = true;
            handler();
            timeout = setTimeout(fn, wait);
        });
    },
};

export const track = {
    bind(el, binding, vnode) {
        if (!binding.value)
            return;
        const daTrack = (trackData) => {
            Bridge.send('parent', 'DATracker1', {
                action: 'track',
                data: trackData,
            });
        };
        const daTrackTime = (trackData) => {
            Bridge.send('parent', 'DATracker1', {
                action: 'time_event',
                data: trackData,
            });
        };
        let trackData = binding.value;
        const track = function () {
            trackData = trackData || {};
            daTrack(trackData);
            if (trackData.time_event)
                daTrack(trackData.time_event);
        };
        if (trackData instanceof Object && trackData.type === 'expression') {
            // this.$watch(value, (newValue, oldValue) => {
            //     trackData = newValue;
            //     if (trackData.time_event && window.DATracker)
            //         DATracker.time_event(trackData.time_event);
            //     dom.on(elem, 'click', track);
            // }, { deep: 1 });
        } else {
            // 有time_event参数，表示需要统计事件耗时， 目前代码里并没有。
            if (trackData.time_event && window.DATracker)
                daTrackTime(trackData.time_event);

            vnode.context.$customEvents = {
                eventName: 'click',
                cb: track,
            };
            event.on(el, vnode.context.$customEvents.eventName, vnode.context.$customEvents.cb);
        }
    },
    unbind(el, binding, vnode) {
        vnode.context.$customEvents && event.off(el, vnode.context.$customEvents.eventName, vnode.context.$customEvents.cb);
    },
};

export const gaq = {
    bind(el, binding, vnode) {
        if (!binding.value)
            return;
        const cb = () => {
            Bridge.send('parent', 'gaq', binding.value);
        };
        vnode.context.$customEvents = {
            eventName: 'click',
            cb,
        };
        event.on(el, vnode.context.$customEvents.eventName, vnode.context.$customEvents.cb);
    },
    unbind(el, binding, vnode) {
        vnode.context.$customEvents && event.off(el, vnode.context.$customEvents.eventName, vnode.context.$customEvents.cb);
    },
};

export const clamp = {
    bind(el, binding, vnode) {
        clampLib(el, { clamp: +binding.value });
    },
};
