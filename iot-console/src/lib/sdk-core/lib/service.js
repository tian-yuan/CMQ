export default class Service {
    constructor(serverName, request, apis) {
        if (DEV) {
            if (!serverName)
                throw new Error('serverName is must');

            if (!request)
                throw new Error('request is must');
        }
        this.serverName = serverName;
        this.request = request;
        this.$env = {};
        if (apis) {
            this.apis = apis;
            this.$config(Object.keys(apis), apis);
        }
    }
    $requestWrap(data, serviceItem) {
        let promise;
        if (serviceItem.check) {
            if (serviceItem.check(data) === false) {
                throw new Error('check data error', data);
            }
        }
        if (serviceItem.mock) {
            promise = serviceItem.mock(data);
        }
        if (promise && promise.then) {
            return promise;
        }
        return this.request(data);
    }
    $config(config, apis) {
        apis = apis || this.apis;
        const apiKeys = Array.isArray(config) ? config : Object.keys(config);
        if (DEV) {
            if (!apis)
                throw new Error('apis is required');
        }
        apiKeys.forEach((item) => {
            const self = this;
            if (Array.isArray(config) ? config.includes(item) : config[item] === true) {
                this[item] = function tmp(data) {
                    const api = apis[item];
                    if (!api)
                        throw new Error('no such api, check api config');
                    return self.$requestWrap(api.init(this)(data), tmp);
                };
            } else
                this[item] = config[item];
        });
        return this;
    }
    $setENV(env = {}) {
        if (DEV) {
            ['region', 'AK', 'SK'].forEach((key) => {
                if (!env[key] && !this.$env[key])
                    throw new Error(`${key} is required`);
            });
            if (!env.host && !this.$env.host)
                console.warn('host can set in here or set in data.headers');
        }
        Object.assign(this.$env, env);
        return this;
    }
    $set(key, value) {
        this.$env[key] = value;
        return this;
    }
    $get(key) {
        return this.$env[key];
    }
}
