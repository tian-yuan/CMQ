import { sign } from '../../../lib/util/sign';
import { assert } from 'chai';
describe('util/sign', () => {
    const now = new Date().toISOString().split('.')[0] + 'Z';
    const yyyymmdd = now.split('T')[0].replace(/-/g, '');
    const req1 = ['cn-east-1', 'aaaa', 'bbbb', 'GET', '/ncs', 'ncs', {
        b: 1,
        a: 2,
    }, {
        HOST: 'xxx.com',
    }, ''];
    const req1Str = JSON.stringify(req1);
    const query1Str = JSON.stringify(req1[6]);
    sign(...req1);
    describe(`sign(${req1Str})`, () => {
        it(`query(${query1Str}) is same`, () => {
            assert.equal(query1Str, JSON.stringify(req1[6]));
        });
        it(`!!header['X-163-SignatureNonce'] return true`, () => {
            assert.equal(true, !!req1[7]['X-163-SignatureNonce']);
        });
        it(`!!header['X-163-Date'] return true`, () => {
            assert.equal(true, !!req1[7]['X-163-Date']);
        });
        it(`header.HOST is same`, () => {
            assert.equal(req1[7].HOST, req1[7].HOST);
        });
        it(`header['X-163-SignatureVersion'] is 2.0`, () => {
            assert.equal('2.0', req1[7]['X-163-SignatureVersion']);
        });
        // if 23:59:59:xxx maybe error.
        it(`header.Authorization`, () => {
            const result = `HMAC-SHA256 Credential=aaaa/${yyyymmdd}/cn-east-1/ncs/163_request, SignedHeaders=host;x-163-date;x-163-signaturenonce;x-163-signatureversion, `;
            assert.equal(result, req1[7].Authorization.split('Signature')[0]);
        });
    });

    const req2 = ['cn-east-1', 'aaaa', 'bbbb', 'POST', '/ncs', 'ncs', {
        d: 1,
        a: 2,
    }, {
        HOST: 'xxx.com',
        'x-host': 'x2.com',
        'X-163-SignatureVersion': '3.0',
    }, {
        xxx: 'ddd',
    }];
    const req2Str = JSON.stringify(req2);
    const query2Str = JSON.stringify(req2[6]);
    sign(...req2);
    describe(`sign(${req2Str})`, () => {
        it(`query(${query2Str}) is same`, () => {
            assert.equal(query2Str, JSON.stringify(req2[6]));
        });
        it(`!!header['X-163-SignatureNonce'] return true`, () => {
            assert.equal(true, !!req2[7]['X-163-SignatureNonce']);
        });
        it(`!!header['X-163-Date'] return true`, () => {
            assert.equal(true, !!req2[7]['X-163-Date']);
        });
        it(`header.HOST is same`, () => {
            assert.equal(req2[7].HOST, req2[7].HOST);
        });
        it(`header['x-host] is same`, () => {
            assert.equal(req2[7]['x-host'], req2[7]['x-host']);
        });
        it(`header['X-163-SignatureVersion'] is 2.0`, () => {
            assert.equal('2.0', req2[7]['X-163-SignatureVersion']);
        });
        it(`header.Authorization`, () => {
            const result = `HMAC-SHA256 Credential=aaaa/${yyyymmdd}/cn-east-1/ncs/163_request, SignedHeaders=host;x-163-date;x-163-signaturenonce;x-163-signatureversion;x-host, `;
            assert.equal(result, req2[7].Authorization.split('Signature')[0]);
        });
    });
});
