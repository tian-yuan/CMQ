import crypto from 'crypto';
const defaultConfig = {
    signatureMethod: 'HMAC-SHA256',
    hash: 'sha256',
};
const defaultEncoding = 'utf8';
const excludeHeaders = new Set(['X-163-SignedHeaders', 'X-163-Signature', 'Authorization']);
const protectHeaders = new Set([...excludeHeaders, 'X-163-Date', 'X-163-Credential', 'X-163-SignatureVersion', 'X-163-SignatureMethod', 'X-163-SignatureNonce', 'X-163-DryRun']);

const orderParams = function orderParams(query) {
    return Object.keys(query).sort().map((key) => `${encodeURIComponent(key)}=${encodeURIComponent(query[key])}`).join('&');
};
/* code from https://gist.github.com/jed/982883 */
const uuid4 = function uuid4(a, b) {
    for (b = a = ''; a++ < 36; b += a * 51 & 52 ? (a ^ 15 ? 8 ^ Math.random() * (a ^ 20 ? 16 : 4) : 4).toString(16) : '-')
        ;
    return b;
};
const getDefaultSignHeader = function getDefaultSignHeader(now) {
    const headers = {
        'X-163-Date': now,
        'X-163-SignatureVersion': '2.0',
        'X-163-SignatureNonce': uuid4().replace(/-/, ''),
    };
    return headers;
};
const fixHeaders = function fixHeaders(headers) {
    Object.keys(headers).forEach((header) => {
        if (protectHeaders.has(header))
            delete headers[header];
    });
};
const formatHeaders = function formatHeaders(headers) {
    return Object.keys(headers).sort().map((key) => `${key.toLowerCase()}:${(headers[key] || '')}`).join('\n') + '\n';
};
const hashHex = function hashHex(str) {
    const hash = crypto.createHash(defaultConfig.hash);
    hash.update(str);
    return hash.digest('hex');
};
const getSignedHeaders = function getSignedHeaders(headers) {
    const result = [];
    Object.keys(headers).forEach((header) => {
        if (!excludeHeaders.has(header))
            result.push(header.toLowerCase());
    });
    return result.sort().join(';');
};
const hashPlayLoad = function hashPlayLoad(requestPayload) {
    return hashHex(requestPayload).toLowerCase();
};
const HMAC = function HMAC(secret, str) {
    return crypto.createHmac('sha256', secret).update(str, defaultEncoding).digest();
};
const HMACToString = function HMAC(secret, str) {
    return crypto.createHmac('sha256', secret).update(str, defaultEncoding).digest('hex');
};
const getSignKey = function getSignKey(SK, regionId, yyyymmdd, serverName) {
    return HMAC(HMAC(HMAC(HMAC(Buffer.from('163' + SK), yyyymmdd), regionId), serverName), '163_request');
};
const getSign = function getSign(toSign, signKey) {
    return HMACToString(signKey, toSign);
};
export function sign(regionId, AK, SK, httpRequestMethod, canonicalURI, serverName, query, headers, requestPayload) {
    const result = [];
    result.push(httpRequestMethod);
    result.push(canonicalURI);

    // 清洗整理参数
    fixHeaders(headers);

    result.push(orderParams(query));

    const now = new Date().toISOString().split('.')[0] + 'Z';
    const yyyymmdd = now.split('T')[0].replace(/-/g, '');

    const credentialScope = `${yyyymmdd}/${regionId}/${serverName}/163_request`;

    const defaultHeaders = getDefaultSignHeader(now, true);
    Object.assign(headers, defaultHeaders);

    const signedHeaders = getSignedHeaders(headers);

    result.push(formatHeaders(headers));
    result.push(signedHeaders);
    result.push(hashPlayLoad(JSON.stringify(requestPayload)));

    const toSign = [defaultConfig.signatureMethod, now, credentialScope, hashHex(result.join('\n')).toLowerCase()].join('\n');
    const signKey = getSignKey(SK, regionId, yyyymmdd, serverName);

    headers.Authorization = `${defaultConfig.signatureMethod} Credential=${AK}/${credentialScope}, SignedHeaders=${signedHeaders}, Signature=${getSign(toSign, signKey)}`;
}
