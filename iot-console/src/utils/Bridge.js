import { stringify } from 'qs';
const contacts = {}; // 联系人
const handles = {
    samErrorMsg(msg) {
        window.samErrorMsg && window.samErrorMsg.add(msg);
    },
}; // 信息处理
const queue = {}; // 信息队列，当未送达的时候缓存在这里
const receipt = function receipt(handleId, cb) {
    return function (data) {
        delete handles[handleId];
        cb(data.err, data.msg);
    };
};
let id = 0;
const guid = function () {
    return ++id;
};
const Bridge = {
    init(router) {
        // 根据window有没有全局变量来判断是否已初始化
        if (window.VueBridge)
            return window.VueBridge;

        // 一定要先设置handles，再注册
        handles.urlchange = function urlchange(path) {
            router.replace(path);
        };

        // 不触发 afterEach
        let fromParent = false;
        handles.justchange = function justchange(path) {
            fromParent = true;
            router.replace(path);
            // 不会触发 afterEach 的情况
            console.log(router.currentRoute.fullPath);
            console.log(path.split('?')[0]);
            if (router.currentRoute.fullPath.split('?')[0] === path.split('?')[0])
                fromParent = false;
        };

        router.afterEach((to, from) => {
            if (!fromParent && to.path !== '/') {
                const query = stringify(to.query);
                Bridge.send('parent', 'urlchange', to.path + (query ? '?' : '') + query);
            }
            fromParent = false;
        });

        Bridge.reg('self', Bridge);
        const parent = window.parent.window.VueBridge;
        if (window.parent !== window && parent) {
            Bridge.reg('parent', parent);
            parent.reg('sub', Bridge);
        }

        return window.VueBridge = Bridge;
    },
    /**
     * 发送信息
     * @param {*} receiver 收件人
     * @param {*} action 描述
     * @param {*} msg 内容
     * @param {*} cb 回执函数
     * @param {*} isTemp 如果收件人不存在，是否丢弃，默认不丢弃，当存在收件人的时候，会依次读取 queue 里面的信息
     */
    send(receiver, action, msg, cb, isTemp) {
        const receiptAction = cb ? this.receipt(cb) : null;
        msg = JSON.stringify(msg || '');
        if (contacts[receiver])
            contacts[receiver].receive(action, msg, receiptAction);
        else if (!isTemp) {
            queue[receiver] = queue[receiver] || [];
            queue[receiver].push([action, msg, receiptAction]);
        }
        return receiptAction ? handles[receiptAction] : null;
    },
    receipt(cb) {
        const receiptAction = 'receipt' + guid();
        handles[receiptAction] = receipt(receiptAction, cb);
        return receiptAction;
    },
    /**
     * 接收信息
     * @param {*} action 描述
     * @param {*} msg 内容
     * @param {*} receiptAction 回执消息
     */
    receive(action, msg, receiptAction) {
        msg = JSON.parse(msg || '');
        if (handles[action])
            handles[action].call(this, msg, receiptAction ? (receipter, receiptMsg) => this.send(receipter, receiptAction, receiptMsg) : null);

        // self.receive(action, msg);
    },
    /**
     * 用户注册
     * @param {*} name 用户名
     * @param {*} person 用户联系方式
     */
    reg(name, person) {
        contacts[name] = person;
        if (queue[name]) {
            queue[name].forEach((send) => {
                person.receive(...send);
            });
            queue[name].length = 0;
        }
    },
};

export default Bridge;

