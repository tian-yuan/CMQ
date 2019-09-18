import { cloneDeep } from 'lodash';
const wm = new Map();
export const cache = {
    add(key, value, notClone) {
        wm.set(key, notClone ? value : cloneDeep(value));
    },
    get(key) {
        return wm.get(key);
    },
    has(key) {
        return wm.has(key);
    },
    clear(key) {
        return wm.delete(key);
    },
    safeAdd(key, func, isPromise = true) {
        if (this.has(key)) {
            return isPromise ? Promise.resolve(cache.get(key)) : cache.get(key);
        } else {
            return func(key);
        }
    },
    safeAddSync(key, func) {
        if (this.has(key)) {
            return this.get(key);
        } else {
            const value = func(key);
            this.add(key, value);
            return value;
        }
    },
};
