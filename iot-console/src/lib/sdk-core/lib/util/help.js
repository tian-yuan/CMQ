export function isString(str) {
    return typeof str === 'string';
}
export function isObject(obj) {
    return Object.prototype.toString.call(obj).indexOf('Object') !== -1;
}
const concatParams = function concatParams(query) {
    return Object.keys(query).map((key) => `${encodeURIComponent(key)}=${encodeURIComponent(query[key])}`).join('&');
};
export function concatURL(url, query) {
    let queryStr = '';
    if (isString(query))
        queryStr = query;
    else if (isObject(query))
        queryStr = concatParams(query);
    else
        queryStr = (query || '').toString();

    if (queryStr)
        url += ((url.indexOf('?') === -1 ? '?' : '&') + queryStr);

    return url;
}
export function resolvePath(url, path) {
    if (url && isObject(path))
        return url.replace(/\{(.*?)\}/g, (a, b) => path[b] || a);

    return url;
}
