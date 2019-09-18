const topics = {};
export default {
    publish(topic, data) {
        if (!topics[topic]) {
            topics[topic] = {
                queue: [],
            };
        } else {
            topics[topic].queue.forEach((func) => {
                func(data);
            });
        }
        topics[topic].last = data;
    },
    unpublish(topic) {
        if (topics[topic])
            delete topics[topic].last;
    },
    subscribe(topic, func) {
        if (!topics[topic]) {
            topics[topic] = {
                queue: [],
            };
        } else if ('last' in topics[topic])
            func(topics[topic].last);

        topics[topic].queue.push(func);
    },
    unsubscribe(topic, func) {
        if (topics[topic])
            topics[topic].queue.splice(topics[topic].queue.indexOf(func), 1);
    },
};
