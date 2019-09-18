import i18n from '@/utils/i18n';
import { Service, createAPI } from '@/lib/sdk-core/dist/neyun.js';
import axios, { response } from './axios';
import convert from 'xml-js';
import transform, { queryFirst } from './transform';
import map from './map';
import Bridge from '@/utils/Bridge';
import { detect } from 'detect-browser';
import { stringify, parse } from 'qs';

const browser = detect();

const ERROR_CODE = {
    REQUEST_ERROR: 1,
    JSON_ERROR: 10,
};
const isPromise = function (func) {
    return func && typeof func.then === 'function';
};
const crossDomain = function (baseURL, url, method, body, headers, config, withCredentials) {
    return new Promise((res, rej) => {
        Bridge.send('parent', 'fetch', {
            url: baseURL + url,
            params: {
                method,
                data: body,
                headers,
                noAlert: config.noAlert,
                credentials: withCredentials ? 'include' : 'same-origin',
            },
        }, (err, msg) => {
            if (err)
                rej(msg);
            else
                res(msg);
        });
    });
};
let request = null;
const errHandles = {
    406({ config }) {
        // cookie失效，跳到登录页
        // noRedirectToIndex主要是一些页面是直接弹窗登录而非跳转到首页
        if (!config.noRedirectToIndex)
            location.href = '/';
    },
    444() {
        // API限流
        Bridge.send('parent', 'error', i18n.t('global.req.mutilError'));
    },
    422({ config, baseURL, url, method, body, headers }, err) {
        // 安全验证弹窗
        return new Promise((res, rej) => {
            Bridge.send('parent', 'security', {
                tel: err.tel,
                isStrong: err.isStrong,
                email: err.email,
            }, (err, msg) => {
                if (err) {
                    rej('cancelSecurity');// 取消安全验证，返回reject
                } else
                    res(msg);
            });
        }).then(() => {
            const wrap = request(baseURL).bind(self);
            return wrap({ url, method, body, headers, config });
        }).catch((res) => Promise.reject(res));
    },
    432({ config }, err) {
        // 用戶被禁用
        location.href = location.origin + '/distribute/user/disabled';
    },
    403({ config }, err) {
        if (err.Code === 'UnauthorizedOperation') {
            window.samErrorMsg && window.samErrorMsg.add(err.msg);
            err.msg = '';
        } else if (!config.noAlert) {
            Bridge.send('parent', 'alert', {
                content: err.msg || i18n.t('global.errorTip.operation'),
                title: i18n.t('global.abnormalTip'),
            });
        }
    },
    defaults({ config }, err) {
        // 出错信息统一弹框处理
        // 不需要弹框的在options里传入noAlert字段
        if ((/^4/.test(err.code) || /^6/.test(err.code)) && !!err.msg) {
            // 如果多个接口同时出错，会有多个弹框，先去掉弹框
            // _cu._$hideModal();
            Bridge.send('parent', 'hideModal');
            Bridge.send('parent', 'alert', {
                content: err.msg,
                title: i18n.t('global.abnormalTip'),
            });
        } else
            Bridge.send('parent', 'error', i18n.t('global.systemBusyTip'));
    },
    localError() {
        Bridge.send('parent', 'error', i18n.t('global.req.client.error'));
    },
};
const formatContentType = function (contentType, data) {
    const map = {
        'application/x-www-form-urlencoded'(data) {
            return stringify(data);
        },
    };
    return map[contentType] && map[contentType](data) || data;
};
export const downloadFile = function (url, data = {}, method = 'post') {
    const formDom = document.createElement('form');
    let inputDom;
    const iframe = document.getElementById('download_iframe') || document.createElement('iframe');
    if (!iframe.getAttribute('id'))
        iframe.setAttribute('id', 'download_iframe');

    iframe.setAttribute('name', 'download_iframe');
    iframe.style.visibility = 'hidden';
    formDom.setAttribute('method', method);
    formDom.setAttribute('action', url);
    formDom.setAttribute('target', 'download_iframe');
    for (const key of Object.keys(data)) {
        inputDom = document.createElement('input');
        inputDom.setAttribute('type', 'hidden');
        inputDom.setAttribute('name', key);
        inputDom.setAttribute('value', data[key]);
        formDom.appendChild(inputDom);
    }
    // formDom.setAttribute("enctype","multipart/form-data")
    document.body.appendChild(iframe);
    document.body.appendChild(formDom);
    formDom.submit();
    document.body.removeChild(formDom);
};
const formatURL = function (str, indexHttp, saveEnds) {
    if (!str.startsWith('/') && !indexHttp) {
        str = '/' + str;
    }
    if (str.endsWith('/') && !saveEnds) {
        str = str.replace(/\/$/, '');
    }
    return str;
};
/**
 * 发送接口
 * @param {Object} data 请求参数
 * {
 *     url,
 *     body,
 *     headers,
 *     method,
 * }
 */
const requestCache = {};
request = function request(cacheBaseURL = '/gateway/v2') {
    return function wrap({ url, method, body, headers, query, config = {} }) {
        const mocker = this && this.mocker;
        const mock = function mock(data) {
            if (config.mock && mocker) {
                // mock 数据
                return Promise.resolve(mocker(config.mock));
            }
            return Promise.resolve(data);
        };
        let baseURL = cacheBaseURL;
        if (config.baseURL)
            baseURL = config.baseURL;

        const httpReg = /^(?:http|https):\/\//g;
        const indexHttp = httpReg.test(baseURL);
        baseURL = formatURL(baseURL, indexHttp);
        const saveEnds = url === 'domain/';
        url = formatURL(url, false, saveEnds);
        if (config.download) {
            if (method.toUpperCase() === 'POST') {
                return downloadFile(baseURL + url, body, method);
            } else {
                const urlArr = (baseURL + url).split('?');
                return downloadFile(urlArr[0], parse(urlArr[1]), method);
            }
        }

        if (url && url.indexOf('/nos') === 0) {
            // NOS的接口需要移除 If-Modified-Since 头，Cache-Control 头在设置文件元数据接口会有冲突
            if (!headers.hasOwnProperty('Cache-Control')) {
                Object.assign(headers, {
                    'Cache-Control': 'no-cache',
                    Pragma: 'no-cache',
                });
            }
        } else {
            Object.assign(headers, {
                'Cache-Control': 'no-cache',
                Pragma: 'no-cache',
                'If-Modified-Since': '0',
            });
        }

        headers = headers || {};
        headers['X-163-AcceptLanguage'] = window.backend.lang.split('-')[0];
        headers['X-Language-Mark'] = window.backend.lang.replace('-', '_');
        let withCredentials = false;
        if (query && query.Action && query.Version) {
            headers['Content-Type'] = headers['Content-Type'] || 'application/json';
        }
        if (indexHttp)
            withCredentials = true;
        if (withCredentials && browser.name === 'ie' && browser.version.split('.')[0] === '9')
            return crossDomain(baseURL, url, method, body, headers, config, withCredentials);
        else {
            let req = axios({
                baseURL,
                method,
                url,
                data: formatContentType(headers['Content-Type'], body),
                headers,
                withCredentials,
                result: config.result,
            });
            if (DEV) {
                // dev 模式下如果参数有 mock ，则会自动 mock 数据, mock 的数据只会是 http code 200 的情况
                req = req.then(mock, mock);
            }
            // 服务方后端给xml数据的时候用，<ant>wwee</ant>类似的接口会被转成{ant:{_text:'wwee'}}内容优化，业务调用中自己处理
            if (config.xml) {
                req = req.then((response) => {
                    const data = response.data;
                    if (data) {
                        // 部分返回字符串形式，需要转换成JSON
                        if (data.params && typeof data.params === 'string')
                            data.params = convert.xml2json(data.params, { compact: true, spaces: 4 });
                    }
                    return response;
                });
            }

            req = req.then(response);

            if (config.process) {
                if (Array.isArray(config.process))
                    config.process.forEach((item) => req = req.then(item));
                else
                    req = req.then(config.process);
            }
            if (config.order) {
                const temp = /Action=([a-zA-Z]*)(&|$)/.exec(url);
                if (temp && temp[1]) {
                    const orderAction = temp[1];
                    req.order = {};
                    req.order.action = orderAction;
                    requestCache[orderAction] = req.order.id = requestCache[orderAction] ? requestCache[orderAction] + 1 : +new Date();// 有的请求隔得很近
                    // console.log('aaa', orderAction, requestCache[orderAction]);
                }
            }
            return req.then((res) => {
                if (req.order) {
                    // console.log('aaaa', req.order);
                    if (req.order.id !== requestCache[req.order.action])
                        return Promise.reject('expired request'); // 过期请求
                }
                return res;
            }).catch((err) => {
                // 处理code
                let handleOut;
                if (err === 'expired request') {
                    throw err;
                }
                if (err.code) {
                    let handle = errHandles[err.code];
                    if (!handle && !config.noAlert)
                        handle = errHandles.defaults;

                    if (handle) {
                        handleOut = handle({
                            config, baseURL, url, method, body, headers,
                        }, err);
                    }
                } else if (err.code === undefined)
                    handleOut = errHandles.localError();

                if (isPromise(handleOut))
                    return handleOut;

                throw err;
            });
        }
    };
};
export default function createService(serverName, apiSchemaList, baseURL = '/gateway/v2') {
    const apiSchemaListArr = Object.keys(apiSchemaList);

    // 通过api单项是否有url属性判断是否需要transform
    if (!apiSchemaList[apiSchemaListArr[0]].url)
        apiSchemaList = transform(serverName, apiSchemaList);

    const service = new Service(serverName, request(baseURL), createAPI(apiSchemaList));
    Object.keys(service).forEach((key) => {
        service[`$${key}`] = service[key];
    });

    // 只有全模块的开关开启之后，单个的接口才能够去配置isQueryFirst信息，否则无效
    if (Object.keys(map).find((item) => map[item].isQueryFirst && item === serverName)) {
        const config = {};

        apiSchemaListArr.forEach((api) => {
            const { preProcess, isQueryFirst } = apiSchemaList[api].config;
            const tmp = Array.isArray(preProcess) ? preProcess : [preProcess];
            const preProcessArr = preProcess ? tmp : [];
            config[api] = isQueryFirst !== false ? [queryFirst].concat(preProcessArr) : preProcessArr;
        });

        setPreProcess(service, config, 100);
    }

    return service;
}

/**
 * 给生成的service对象生成对应的接口方法，生成对应的前置数据处理方法
 * @param {Object} service - 通过createService生成的Service对象
 * @param {object|array} config - 对应函数的前置数据处理方法.
 * @param {number} [index=10] - 对应的处理函数的权重，默认为1，则按照设置的先后顺序，权重越高，越后处理（方法没有index属性的也当做1）手动可设置范围【1-999】
 * @param {function|array} unifiedFunc - 给apis添加统一的处理一个或多个方法.如果同时声明了unifiedFunc和config中的单独的函数处理方法
 * @desc  有两种方式添加preProcess：在api层添加preProcess字段；在 createService之后使用setPreProcess添加。
 *        建议使用setPreProcess添加。因为处理方法涉及较多的业务代码，而且较多可以复用代码。放在api层不方便。
 *        api层默认不经常使用，即定义完之后基本不变动。但是涉及到业务代码，就会造成经常性的变动和查看。
 * @example
 *
 *      import createService, { setPreProcess } from '@/services';
 *      import esAPIs from './apis/elasticSearch';
 *      const preProcess1 = (params) => {
 *           params.query ? (params.query.Limit = 20) : '';
 *           return params;
 *      }
 *      const preProcess2 = (params) => {
 *           params.query ? (params.query.Limit2 = 20) : '';
 *           return params;
 *      }
 *      const esService = createService('nes', esAPIs);
 *      setPreProcess(esService, {
 *          loads: [preProcess1, preProcess2]
 *          // ...otherServiceName
 *      })
 * @see
 *      services/monitor/index          webapi的实现
 *      views/dashboard/nes/services    openapi的实现
 */
export const setPreProcess = function (service, config = {}, index = 10, unifiedFunc) {
    if (!(config instanceof Object))
        return;

    if (index > 999)
        index = 999;
    if (index < 1)
        index = 1;
    index = parseInt(index);

    const apiNames = Object.keys(config);
    const apis = service.apis;

    // 复制 this.apis 下已有的 preProcess 处理方式放到 config 中
    apiNames.forEach((item) => {
        if (index !== 10)
            Array.isArray(config[item]) ? (config[item].forEach((item) => item.index = index)) : (config[item].index = index);

        const preProcess = apis[item].config.preProcess;
        const tmp = Array.isArray(preProcess) ? preProcess : [preProcess];
        const preProcessArr = preProcess ? tmp : [];
        config[item] = preProcessArr.concat(config[item]);
    });

    // 将 unifiedFunc 先绑定到对应的 config 属性内
    // unifiedFunc权重默认最高(设置一个极大值，手动设置范围之外)，即放到最后处理
    if (unifiedFunc) {
        const unifiedFuncArr = Array.isArray(unifiedFunc) ? unifiedFunc : [unifiedFunc];
        unifiedFuncArr.forEach((item) => item.index = 1000);
        Object.keys(apis).forEach((api) => {
            if (apiNames.includes(api)) {
                if (Array.isArray(config[api]))
                    config[api] = unifiedFuncArr.concat(config[api]);
                else
                    config[api] = unifiedFuncArr.concat([config[api]]);
            } else
                config[api] = apis[api];
        });
    }

    apiNames.forEach((item) => {
        const preFunc = config[item];
        // 权重排序
        preFunc.sort((a, b) => {
            a.index === undefined && (a.index = 1);
            b.index === undefined && (b.index = 1);

            return a.index > b.index;
        });
        apis[item].config.preProcess = preFunc;
        config[item] = function (params) {
            Array.isArray(preFunc) ? preFunc.forEach((item) => params = item(params)) : (params = preFunc(params));
            return this.request(this.apis[item].init(this)(params));
        };
    });
    service.$config(config);
};

/**
 * 给定义apis(即：接口们的元数据信息)添加额外的后置数据处理方法
 * @param {Object} apis - 接口们的元数据信息
 * @param {function|Array} config - 对应函数的后置数据处理方法
 * @desc  在使用的时候需要在调用 createService 之前调用
 * @example
 *
 *      import createService, { setProcess } from '@/services';
 *      import esAPIs from './apis/elasticSearch';
 *      const esService = createService('nes', esAPIs);
 *      setProcess(esService.apis, {
 *          loads: (result) => {
 *              const Instances = (result.Instances || []).map((item) => edit(item));
 *              return { Instances, Total: result.Total };
 *          },
 *          load: [(result) => {
 *              const Instance = edit(result.Instance);
 *              return Instance;
 *          }, (result) => {
 *              result.EngineVersion = '111';
 *              return result;
 *          }],
 *      });
 *      // todo
 */

export const setProcess = function (apis, config = {}) {
    if (!(config instanceof Object))
        return;

    Object.keys(config).forEach((key) => {
        const processFunc = config[key];
        // 兼容transform前后的两种方式的APIs
        if (apis[key].url) {
            if (!apis[key].config)
                apis[key].config = {};
            apis[key].config.process = processFunc;
        } else
            apis[key].process = processFunc;
    });
    return apis;
};

export const getAll = function (req, { get: format, set: readd, limit, offset, addBody }, data, ...args) {
    let all = null;
    let count = 0;
    const Limit = limit || 1000;
    const Offset = offset || 0;
    const getData = function (Offset) {
        const formatData = Object.assign({ }, data);
        const action = addBody ? 'body' : 'query';
        formatData[action] = Object.assign({}, formatData[action], {
            Limit,
            Offset,
        });
        return formatData;
    };
    return req(getData(Offset), ...args).then((json) => {
        const { list, Count } = format(json);
        if (Count <= Limit) {
            return json;
        } else {
            all = list;
            count = Count;
            const queue = [];
            for (let i = 1; i < Math.ceil(Count / Limit); i++) {
                queue.push(req(getData(i * Limit), ...args));
            }
            return Promise.all(queue).then((results) => {
                results.forEach((item) => {
                    const { list, Count } = format(item);
                    all = all.concat(list);
                    count = Count;
                });
                readd(json, all, count);
                return json;
            });
        }
    });
};

export const regionURL = function regionURL(data, regionId, prefix = '') {
    if (!regionId)
        return;
    const config = {
        baseURL: getDomain(regionId) + prefix,
    };
    data.config = Object.assign(data.config || {}, config);
};
