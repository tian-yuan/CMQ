import emitter from './emitter';

export default {
    mixins: [emitter],
    name: 'routeComponents',
    data() {
        return {
            routerChildList: [],
            parentRouteComponent: undefined,
        };
    },
    created() {
        this.$on('addNewRouteComponents', (event) => {
            event.component.parentRouteComponent = this;
            this.routerChildList.push(event.component);
            event.component.$emit('addParent');
        });
        this.$on('removeNewRouteComponents', (event) => {
            const childIndex = this.routerChildList.indexOf(event.component);
            this.routerChildList.splice(childIndex, 1);
        });
        this.$on('onMessageFromParent', (event) => {
            const childIndex = this.routerChildList.indexOf(event.component);
            this.routerChildList.splice(childIndex, 1);
        });
        this.dispatch('routeComponents', 'addNewRouteComponents', {
            component: this,
        });
    },
    methods: {
        sendMessage(options, type) {
            if (!type)
                type = 'all';
            switch (type) {
                case 'child':
                    this.__sendMessageToChild(options);
                    break;
                case 'parent':
                    this.__sendMessageToParent(options);
                    break;
                default:
                    this.__sendMessageToParent(options);
                    this.__sendMessageToChild(options);
            }
        },
        getParent() {
            return this.parentRouteComponent;
        },
        __sendMessageToParent(options) {
            this.dispatch('routeComponents', 'onMessage', JSON.parse(JSON.stringify(options)));
        },
        __sendMessageToChild(options) {
            const children = this.routerChildList;
            for (const child of children)
                child.$emit('onMessage', JSON.parse(JSON.stringify(options)));
        },
    },
    destroyed() {
        this.dispatch('routeComponents', 'removeNewRouteComponents', {
            component: this,
        });
    },
};
