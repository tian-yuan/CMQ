function noop() {
    return '';
}
const noSupport = {
    set: noop,
    get: noop,
    remove: noop,
    clear: noop,
};

const isLocalStorageAvailable = () => {
    try {
        const key = 'testlocastorage';
        window.localStorage.setItem(key, '1');
        window.localStorage.removeItem(key);

        return true;
    } catch (error) {
        return false;
    }
};

const storage = isLocalStorageAvailable() ? window.localStorage : null;

const storageObj = !storage ? noSupport : {
    set(key, value) {
        storage.setItem(key, value);
    },
    get(key) {
        return storage.getItem(key);
    },
    remove(key) {
        return storage.removeItem(key);
    },
    clear() {
        storage.clear();
    },
};

export default storageObj;

