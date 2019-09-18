import moduleMap from './map';
/**
 * 转化请求源信息的组织方式
 * @param {string} serverName - 模块名，如:ncs、nvm、ncv、nes等。
 * @param {object} map - 请求源信息的一种组织方式
 *
 * @example
 *  set: {
        action: 'ConfigLogstash',
        method: 'POST',
        noAlert: true,
        preProcess...
        process...
        version...
        isQueryFirst(默认为false)...
    },

    to

    set: {
        url: {
            path,
            method: 'post',
            query: {
                Action: 'ConfigLogstash',
                Version,
            },
        },
        config: {
            noAlert: true,
            preProcess...
            process...
            isQueryFirst(默认为false)...
        }
    },
 */
export default function transform(serverName, map = {}) {
    const tmp = {};
    // webapi 没有对应的Version字段
    const Version = moduleMap[serverName] ? moduleMap[serverName].version : '';

    Object.keys(map).forEach((key) => {
        const obj = map[key];
        const query = {};
        // 这里用action参数来区分是否是OpenAPI
        if (obj.action) {
            Object.assign(query, {
                Version: obj.version || Version,
                Action: obj.action,
            });
        }

        tmp[key] = {
            url: {
                path: obj.path || serverName,
                method: (obj.method && obj.method.toLowerCase()) || 'get',
                query: Object.assign(query, obj.query),
            },
            config: {
                process: obj.process,
                preProcess: obj.preProcess,
                noAlert: !!obj.noAlert,
                // undefined有作用，作为第三种状态。表示默认态。根据全模块的设置走
                // 为true没有意义
                // 为false表示单独的不开启queryFirst开关
                isQueryFirst: obj.isQueryFirst,
            },
        };
    });

    return tmp;
}

const titleCase = (str) => str.replace(/(^|\s)\S/g, (L) => L.toUpperCase());

export const trans2apis = (map, instance = '') => {
    const apis = {};

    map.forEach((item) => {
        const methodName = item.methodName;
        apis[methodName] = {
            method: item.method || 'get',
            action: item.action || (titleCase(methodName) + titleCase(instance)),
            preProcess: item.preProcess,
            process: item.process,
            noAlert: !!item.noAlert,
        };
    });
    return apis;
};

export const queryFirst = (params = {}) => {
    let newParams = {};

    if (!params.query) {
        newParams.query = {};
        Object.keys(params).forEach((key) => {
            const value = params[key];
            if (!['headers', 'config', 'body', 'path'].includes(key))
                newParams.query[key] = value;
            else
                newParams[key] = value;
        });
    } else
        newParams = params;

    return newParams;
};
