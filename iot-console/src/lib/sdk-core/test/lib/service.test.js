import {
    assert,
} from 'chai';
import {
    default as Service,
} from '../../lib/service';
import {
    default as createAPI,
} from '../../lib/api';
import {
    apisJSON,
} from './apis';
const noop = function noop() {

};
const request = function request(data) {
    return data;
};
describe('lib/service', () => {
    describe('new Service()', () => {
        it('throw error', () => {
            assert.throws(() => {
                new Service();
            }, Error, /serverName is must/);
        });
    });
    const serverName = 'a';
    describe(`new Service(${serverName})`, () => {
        it('throw error', () => {
            assert.throws(() => {
                new Service(serverName);
            }, Error, /request is must/);
        });
    });
    describe(`new Service(${serverName}, request)`, () => {
        const s1 = new Service(serverName, request);
        it('has $env', () => {
            assert.isObject(s1.$env);
        });
        it('has correct serverName', () => {
            assert.equal(s1.serverName, serverName);
        });
        const config1 = {
            AK: 'aaaa',
            SK: 'bbbb',
            host: 'xx.com',
            region: 'cn-east-1',
        };
        it('$setENV() throw error', () => {
            assert.throws(() => {
                s1.$setENV();
            }, Error, 'region is required');
        });
        const config0 = {
            AK: 'b',
            SK: 'b',
            region: 'a',
        };
        it('$setENV(region, AK, SK) will pass', () => {
            const result = s1.$setENV(config0).$env;
            assert.equal(JSON.stringify(config0), JSON.stringify(result));
        });
        it('$setENV(config1) will get config1', () => {
            s1.$setENV(config1);
            assert.equal(JSON.stringify(s1.$env), JSON.stringify(Object.assign(config0, config1)));
        });
        const config2 = {
            AK: 'cccc',
            SK: 'bbbb',
        };
        it('$setENV(config2) will get mix config1 and config2', () => {
            s1.$setENV(config2);
            assert.equal(JSON.stringify(s1.$env), JSON.stringify(Object.assign(config0, config1, config2)));
        });
        const value = 'ddd';
        it(`$set("AK", ${value})`, () => {
            s1.$set('AK', value);
            assert.equal(s1.$env.AK, value);
        });
        it(`$get("SK")`, () => {
            assert.equal(s1.$env.SK, s1.$get('SK'));
        });
        it(`$config(object) throw error`, () => {
            assert.throws(() => {
                s1.$config({
                    a: 1,
                    b: 2,
                });
            }, Error, /apis is must/);
        });

        describe(`$config(object, apis)`, () => {
            const sa = s1.$config({
                get: true,
                set: true,
                a: true,
                c() {
                    return 'custom';
                },
            }, createAPI(apisJSON, serverName));
            it(`[server][action] is function`, () => {
                assert.isFunction(sa.a);
            });
            it(`[server][action]() will throw error if not api`, () => {
                assert.throws(() => {
                    sa.a();
                }, Error, /no such api, check api config/);
            });
            it(`[server][action]() api not sign`, () => {
                assert.equal(JSON.stringify({ url: '/get?id=2', body: '', headers: {}, method: 'GET', config:{},query:{id:2} }), JSON.stringify(sa.get({
                    query: {},
                })));
            });
            it(`[server][action]() api sign`, () => {
                const now = new Date().toISOString().split('.')[0] + 'Z';
                const yyyymmdd = now.split('T')[0].replace(/-/g, '');
                const result = sa.set({
                    path: {
                        id: 2,
                    },
                    headers: {
                        host: 'cc.com',
                    },
                });
                delete result.headers['X-163-SignatureNonce'];
                delete result.headers['X-163-Date'];
                result.headers.Authorization = result.headers.Authorization.split('Signature')[0];
                assert.equal(JSON.stringify({
                    url: '/set/2?name=%E6%88%91',
                    body: {},
                    headers: {
                        a: 1,
                        host: 'cc.com',
                        'X-163-SignatureVersion': '2.0',
                        Authorization: `HMAC-SHA256 Credential=ddd/${yyyymmdd}/cn-east-1/a/163_request, SignedHeaders=a;host;x-163-date;x-163-signaturenonce;x-163-signatureversion, `,
                    },
                    method: 'POST',
                    config:{},
                    query:{name:"我"},
                    path:{id:2}
                }), JSON.stringify(result));
            });
            it(`[server][action]() will throw error if query not pass valid`, () => {
                assert.throws(() => {
                    sa.get();
                }, Error);
            });
        });
    });
    describe(`new Service(${serverName}, request, apis)`, () => {
        const sa = new Service(serverName, request, createAPI(apisJSON, serverName));
        it('[server][action]() will return', () => {
            assert.equal(JSON.stringify({ url: '/get?id=2', body: '', headers: {}, method: 'GET', config: {}, query: {id: 2 } }), JSON.stringify(sa.get({
                query: {},
            })));
        });
    });
});
