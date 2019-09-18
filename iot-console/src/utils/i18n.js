import VueI18n from 'vue-i18n';
import Vue from 'vue';
Vue.use(VueI18n);
window.langMessage = window.langMessage || {};
const locale = window.backend.lang;
const instance = new VueI18n({
    locale: locale || 'zh-CN',
    messages: {
        'zh-CN': window.langMessage['zh-CN'],
        'en-US': window.langMessage['en-US'],
    },
});
Vue.prototype.$tl = instance.tl = function (...args) {
    return instance.t(...args).replace(/^[A-Z]/, (a) => a.toLowerCase());
};
Vue.prototype.$tlb = instance.tlb = function (...args) {
    return (instance.locale === 'en-US' ? ' ' : '') + instance.tl(...args);
};
Vue.prototype.$tu = instance.tu = function (...args) {
    return instance.t(...args).replace(/^[a-z]/, (a) => a.toUpperCase());
};
Vue.prototype.$tub = instance.tub = function (...args) {
    return instance.tu(...args) + (instance.locale === 'en-US' ? ' ' : '');
};
Vue.i18n = instance;
{
    const document = window.document.documentElement;
    const classes = document.className;
    document.className = classes + ' lang-' + locale + ' ';
}
export default instance;
