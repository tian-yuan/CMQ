import { concatURL, resolvePath } from './util/help';
if (DEV) {
    const Ajv = require('ajv');
    var ajv = new Ajv();
    var keys = ['query', 'path', 'body', 'response'];
    if (SERVER)
        var jsf = require('json-schema-faker/lib/index.js');
    else
        var jsf = require('json-schema-faker/dist/json-schema-faker.js');
}
if (SERVER) {
    const request = require('request');
    var urlNode = require('url');
    var sign = require('./util/sign');
}
const bodyMethods = ['PATCH', 'POST', 'PUT'];

class API {
    constructor(model, serverName) {
        const url = model.url;
        if (SERVER)
            this.sign = model.sign; // 是否需要签名

        this.url = url.path;
        this.serverName = serverName;
        this.query = url.query;
        this.method = url.method.toUpperCase();
        this.headers = model.headers;
        this.config = model.config;
        this.body = model.body;
        if (DEV) {
            // 初始化验证
            keys.forEach((key) => {
                if (model[key]) {
                    if (key !== 'response')
                        this[`${key}Schema`] = ajv.compile(model[key]);
                    else {
                        Object.keys(model[key]).forEach((status) => {
                            this[`req${status}`] = ajv.compile(model[key][status]);
                        });
                    }
                }
            });
        }
    }
    init(service) {
        return (data) => {
            const { $env: env } = service;
            const method = this.method;
            data = data || {};
            if (DEV) {
                // 验证参数
                keys.forEach((key) => {
                    const keyData = data[key];
                    const keySchema = this[`${key}Schema`];
                    if (keySchema) {
                        if (!keySchema(keyData)) {
                            const msg = `[${method}] ${this.url} ${key} params is error.\n${JSON.stringify(keySchema.errors, null, '\t')}`;
                            throw new Error(msg);
                        }
                    }
                });
            }
            let url = this.url;
            const path = data.path;
            if (path) {
                // 在不需要签名的情况下，可以写 `/a/{id}` 的路径
                url = resolvePath(url, path);
            }
            // 在调用接口方法时传递的query参数
            const querys = data.query;
            if(!querys) {
                Object.keys(data).forEach((key) => {
                    const value = data[key];
                    if (!['headers', 'config', 'body', 'path'].includes(key))
                        querys[key] = value;
                })
            }

            const query = Object.assign({}, this.query, querys);
            const headers = Object.assign({}, this.headers, data.headers);
            const config = Object.assign({}, this.config, data.config);
            let body = '';
            if (bodyMethods.includes(method)) {
                if (Array.isArray(data.body)) {
                    body = data.body;
                } else {
                    body = Object.assign({}, this.body, data.body) || body;
                }
            }

            if (SERVER) {
                if (this.sign) {
                    if (!(headers.host) && !(env.host))
                        throw new Error('host is required if you want to sign');

                    headers.host = (headers.host) || urlNode.parse(env.host).host;
                    // query and headers will change.
                    sign.sign(env.region, env.AK, env.SK, method, this.url, this.serverName, query, headers, body);
                }
            }
            url = concatURL(url, query);
            return {
                url,
                body,
                headers,
                method,
                config,
                query,
                path,
            };
        };
    }
}
export default function createAPI(apis, serverName) {
    const modelAPI = {};
    Object.keys(apis).forEach((action) => {
        modelAPI[action] = new API(apis[action], serverName);
    });
    // 释放内存
    apis = {};
    return modelAPI;
}
