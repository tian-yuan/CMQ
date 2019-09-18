import { default as subscribe } from '../subscribe';
export default {
    created() {
        const subscribes = this.$options.subscribes;
        if (subscribes) {
            this.$options.subscribeQueues = {};
            Object.keys(subscribes).forEach((topic) => {
                this.$options.subscribeQueues[topic] = subscribes[topic].bind(this);
                subscribe.subscribe(topic, this.$options.subscribeQueues[topic]);
            });
        }
        const publishs = this.$options.publishs;
        if (publishs) {
            Object.keys(publishs).forEach((topic) => {
                this.$watch(publishs[topic], (newVal) => {
                    subscribe.publish(topic, newVal);
                });
            });
        }
    },
    beforeDestroy() {
        const subscribes = this.$options.subscribes;
        if (subscribes) {
            Object.keys(subscribes).forEach((topic) => {
                subscribe.unsubscribe(topic, this.$options.subscribeQueues[topic]);
            });
        }
        const publishs = this.$options.publishs;
        if (publishs) {
            Object.keys(publishs).forEach((topic) => {
                subscribe.unpublish(topic);
            });
        }
    },
    subscribe,
};
