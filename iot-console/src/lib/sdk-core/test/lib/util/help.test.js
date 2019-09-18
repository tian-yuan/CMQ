import { isString, isObject, concatURL, resolvePath } from '../../../lib/util/help';
import { assert } from 'chai';

const types = ['a', {}, null, true, undefined, 2, []];
describe('util/help', () => {
    describe('isString', () => {
        it(`isString(${JSON.stringify(types[0])}) return true`, () => {
            assert.equal(true, isString(types[0]));
        });
        types.slice(1).forEach((item) => {
            it(`isString(${JSON.stringify(item)}) return false`, () => {
                assert.equal(false, isString(item));
            });
        });
    });
    describe('isObject', () => {
        it(`isObject(${JSON.stringify(types[1])}) return true`, () => {
            assert.equal(true, isObject(types[1]));
        });
        [types[0], ...types.slice(2)].forEach((item) => {
            it(`isObject(${JSON.stringify(item)}) return false`, () => {
                assert.equal(false, isObject(item));
            });
        });
    });
    describe('concatURL', () => {
        const url = 'a/b/';
        const query = {
            id: 2,
            n: '我',
        };
        const url2 = 'a/b/?c=2';
        const url3 = 'a/b/?123';
        const query2 = {
            id: 2,
            n: '我',
        };
        const query3 = 'c=1&d=2';
        it(`concatURL(${url}, ${JSON.stringify(query)}) return a/b/?id=2&n=%E6%88%91`, () => {
            assert.equal('a/b/?id=2&n=%E6%88%91', concatURL(url, query));
        });
        it(`concatURL(${url}, ${JSON.stringify(query3)}) return a/b/?c=1&d=2`, () => {
            assert.equal('a/b/?c=1&d=2', concatURL(url, query3));
        });
        it(`concatURL(${url2}, ${JSON.stringify(query3)}) return a/b/?c=2&c=1&d=2`, () => {
            assert.equal('a/b/?c=2&c=1&d=2', concatURL(url2, query3));
        });
        it(`concatURL(${url}) return a/b/`, () => {
            assert.equal('a/b/', concatURL(url));
        });
        it(`concatURL(${url2}, ${JSON.stringify(query2)}) return a/b/?c=2&id=2&n=%E6%88%91`, () => {
            assert.equal('a/b/?c=2&id=2&n=%E6%88%91', concatURL(url2, query2));
        });
        it(`concatURL(${url3}, ${JSON.stringify(query2)}) return a/b/?123&id=2&n=%E6%88%91`, () => {
            assert.equal('a/b/?123&id=2&n=%E6%88%91', concatURL(url3, query2));
        });
    });
    describe('resolvePath', () => {
        const url = '/a/{id}';
        const url1 = '/a/{id}/detail';
        const path = {
            id: 2,
        };
        const path1 = undefined;
        it(`resolvePath(${url}, ${JSON.stringify(path)}) return /a/2`, () => {
            assert.equal('/a/2', resolvePath(url, path));
        });
        it(`resolvePath(${url1}, ${JSON.stringify(path)}) return /a/2/detail`, () => {
            assert.equal('/a/2/detail', resolvePath(url1, path));
        });
        it(`resolvePath(${url}, ${JSON.stringify(path1)}) return /a/{id}`, () => {
            assert.equal('/a/{id}', resolvePath(url, path1));
        });
    });
});
