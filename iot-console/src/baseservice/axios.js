import i18n from '@/utils/i18n';
import axios from 'axios';

/**
 * @TODO: pcache
 * @TODO: nos head
 * @TODO: error message
 */

axios.defaults.xsrfCookieName = 'NICE_CSRF';
axios.defaults.xsrfHeaderName = 'X-CSRF-Token';

const SUCCESS_RE = /^2/;
const ERROR_CODE = {
    REQUEST_ERROR: 1,
    JSON_ERROR: 10,
};

const options = {};
export function response(response) {
    const data = response.data;

    // NOS
    // 增加method为head时的处理
    if (response.config.result === 'headers')
        return response.headers;

    if (data) {
        // 部分返回字符串形式，需要转换成JSON
        if (data.params && typeof data.params === 'string')
            data.params = JSON.parse(data.params);

        // 请求的操作成功、
        if (SUCCESS_RE.test(data.code))
            return data.params || data;

        // nqs的登录管理页接口。没有msg，params,code等字段
        if (data.name && data.tags)
            return data;
    } else
        return Promise.reject({});

    // 处理message
    let msg = data.msg;
    if (msg && typeof msg === 'string') {
        try {
            msg = JSON.parse(msg);
        } catch (e) { }
    }
    if (msg) {
        if (msg.Code)
            data.Code = msg.Code;

        // nlb
        if (msg.message)
            msg = msg.message;
        else if (msg.msg)
            msg = msg.msg;

        // rds
        if (msg.Error && msg.Error.Description) {
            Object.assign(data, msg.Error);
            msg = msg.Error.Description;
        }
        // NOS
        if (msg.Error && msg.Error.Message)
            msg = msg.Error.Message;

        // nqs
        if (msg.errorMessage)
            msg = msg.errorMessage;

        // kafka
        if (msg.Message)
            msg = msg.Message;
    } else if (data.params) {
        // 一些透传的接口，错误信息在params里
        if (data.params.message)
            msg = data.params.message;
        else if (data.params.msg)
            msg = data.params.msg;
        else if (data.params.Error && data.params.Error.Message)
            msg = data.params.Error.Message;
    }
    // 表示是RDS，code430特殊处理
    if (data.code === 430 && options.data && options.data.Action)
        msg = i18n.t('global.bill.moneyOut');

    // NOS接口，无权限访问时
    if (data.code === 403 && data.params
        && data.params.Error && data.params.Error.Code === 'AccessDenied')
        msg = i18n.t('global.noPermissionTip');

    data.msg = msg;

    const error = new Error();
    error.result = data;
    return Promise.reject(data);
}

// axios.interceptors.response.use((r) => r, (error) => {
//     const data = error.result || {};
//     console.log(data.code + ': ' + data.msg);
//     // new Error()得到的message是string类型
//     switch (error.message) {
//         case ERROR_CODE.REQUEST_ERROR + '':
//         case 'Failed to fetch':
//             window.vueBridge.error('网络或浏览器出现问题，请稍后再试');
//             break;
//         default:
//             // 处理code
//             if (data.code === 406) {
//                 // cookie失效，跳到登录页
//                 // noRedirectToIndex主要是一些页面是直接弹窗登录而非跳转到首页
//                 if (!options.noRedirectToIndex)
//                     location.href = '/';
//             } else if (data.code === 444) {
//                 // API限流
//                 window.vueBridge.error('接口访问频繁，请稍候再试');
//             } else if (data.code === 422) {
//                 // 安全验证弹窗
//                 // return SecurityModal.show({
//                 //     data: { telNumber: data.tel, isStrong: data.isStrong },
//                 // })
//                 //     .then(() => base.fetch(url, options))
//                 //     .catch((res) => {
//                 //         if (res && res.sender)// 取消安全验证，返回reject
//                 //             return Promise.reject('cancelSecurity');
//                 //         return Promise.reject(res);
//                 //     });
//             } else if (data.code === 432) {
//                 // 用户被禁用
//                 location.href = location.origin + '/distribute/user/disabled';
//             } else if (!options.noAlert) {
//                 // 出错信息统一弹框处理
//                 // 不需要弹框的在options里传入noAlert字段
//                 if ((/^4/.test(data.code) || /^6/.test(data.code)) && !!data.msg) {
//                     // 如果多个接口同时出错，会有多个弹框，先去掉弹框
//                     window.vueBridge.hideModal();
//                     window.vueBridge.alert({
//                         content: data.msg,
//                         title: '异常提示',
//                     });
//                 } else
//                     window.vueBridge.error('系统正忙，请稍后再试');
//             }
//             break;
//     }

//     return Promise.reject(error.data);
// });

export default axios;
