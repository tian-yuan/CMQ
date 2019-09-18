function getArr(target) {
    const targetArr = target.split('.');
    targetArr.forEach((item, index) => {
        targetArr[index] = targetArr[index].split('[').map((item) => item.replace(']', ''));
    });
    return targetArr.reduce((a, b) => a.concat(b), []);
}
function getValue(source, keyArr) {
    const keyLastIndex = keyArr.length - 1;
    let cur = source;
    let value;
    keyArr.some((item, index) => {
        if (index !== keyLastIndex) {
            if (!cur[item])
                return true;

            cur = cur[item];
        } else
            value = cur[item];

        return false;
    });
    return value;
}
function trans(source, key, target) {
    if (Array.isArray(key)) {
        key.forEach((i) => {
            trans(source, i, target);
        });
        return;
    }
    if (Array.isArray(target)) {
        target.forEach((i) => {
            trans(source, key, i);
        });
        return;
    }
    const targetArr = getArr(target);
    const targetLastIndex = targetArr.length - 1;
    const keyArr = getArr(key);
    let cur = source;

    targetArr.forEach((item, index) => {
        if (index !== targetLastIndex) {
            cur[item] = cur[item] || {};
            cur = cur[item];
        } else
            cur[item] = getValue(source, keyArr);
    });
}
export default function transform(source, rules, revert, cb) {
    if (Array.isArray(source)) {
        source.forEach((item) => {
            transform(item, rules, revert, cb);
        });
    } else if (source instanceof Object) {
        Object.keys(rules).forEach((key) => {
            if (revert)
                trans(source, rules[key], key);
            else
                trans(source, key, rules[key]);
        });
        cb && cb(source);
    }
    return source;
}
