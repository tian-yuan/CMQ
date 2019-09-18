import {
    assert,
} from 'chai';
import {
    apisJSON,
} from './apis';
import {
    default as createAPI,
} from '../../lib/api';
describe('lib/api', () => {
    const serverName = 'people';
    const apis = createAPI(apisJSON, serverName);
    describe('createAPI', () => {
        it('all apis', () => {
            assert.equal(Object.keys(apisJSON).join(''), Object.keys(apis).join(''));
        });
        it('[api].sign', () => {
            assert.equal(!!apisJSON.get.sign, !!apis.get.sign);
        });
        it('[api].url', () => {
            assert.equal(apisJSON.get.url.path, apis.get.url);
        });
        it('[api].serverName', () => {
            assert.equal(serverName, apis.get.serverName);
        });
        it('[api].query', () => {
            assert.equal(JSON.stringify(apisJSON.get.url.query), JSON.stringify(apis.get.query));
        });
        it('[api].method', () => {
            assert.equal(apisJSON.get.url.method.toUpperCase(), apis.get.method.toUpperCase());
        });
        it('[api].headers', () => {
            assert.equal(JSON.stringify(apisJSON.set.headers), JSON.stringify(apis.set.headers));
        });
        const keys = ['query', 'path', 'body', 'response'];
        keys.forEach((key) => {
            if (apisJSON.get[key]) {
                if (key !== 'response') {
                    it(`[api].${key}Schema`, () => {
                        assert.equal(!!apisJSON.get[`${key}`], !!apis.get[`${key}Schema`]);
                    });
                } else {
                    Object.keys(apisJSON.get[key]).forEach((status) => {
                        it(`[api].response ${status}`, () => {
                            assert.equal(!!apisJSON.get.response[status], !!apis.get[`req${status}`]);
                        });
                    });
                }
            }
        });
        const service = {
            $env: {
                region: 'cn-east-1',
                AK: 'aaaa',
                SK: 'bbbb',
                host: 'xxx.com',
            },
        };
        const now = new Date().toISOString().split('.')[0] + 'Z';
        const yyyymmdd = now.split('T')[0].replace(/-/g, '');

        const init = apis.get.init(service);
        it('[api].init() return function', () => {
            assert.isFunction(init);
        });
        it('[api].init()() throws error', () => {
            assert.throws(init, Error, /\[GET\] \/get query params is error/);
        });
        const init2 = apis.set.init(service);
        const data = {
            path: {
                a: 1, // not work
            },
            query: {
                c: 1,
            },
            headers: {
                d: 2,
            },
        };
        it('[api].init()(data)', () => {
            const result = init2(data);
            delete result.headers['X-163-SignatureNonce'];
            delete result.headers['X-163-Date'];
            result.headers.Authorization = result.headers.Authorization.split('Signature')[0];
            assert.equal(JSON.stringify({
                url: '/set/{id}?name=%E6%88%91&c=1',
                body: {},
                headers: {
                    a: 1,
                    d: 2,
                    host: null,
                    'X-163-SignatureVersion': '2.0',
                    Authorization: `HMAC-SHA256 Credential=aaaa/${yyyymmdd}/cn-east-1/people/163_request, SignedHeaders=a;d;host;x-163-date;x-163-signaturenonce;x-163-signatureversion, `,
                },
                method: 'POST',
                config:{},
                query:{name:"我", c:1},
                path:{a:1}
            }), JSON.stringify(result));
        });
        const init3 = apis.list.init({
            $env: {
                region: 'cn-east-1',
                AK: 'aaaa',
                SK: 'bbbb',
            },
        });
        it('[api].init()(data) not host will throw error', () => {
            assert.throws(() => {
                init3(data);
            }, Error, /host is required if you want to sign/);
        });
    });
});
