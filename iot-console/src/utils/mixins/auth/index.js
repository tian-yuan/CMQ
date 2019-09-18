import i18n from '@/utils/i18n';
import Vue from 'vue';
/**
 * needCertify 是否需要判断实例是否欠费
 */
export default {
    created() {
        this.initArrearMsg({ useCache: true, needCertify: this.$options.needCertify });
    },
    methods: {
        countEnoughLabel() {
            const { dec, isPrtz } = this.$options.auth;
            let xx = i18n.t('global.addQuota1');
            if (dec) {
                xx = i18n.t('global.addQuota2');
            }
            if (isPrtz) {
                xx = i18n.t('global.addQuota3');
            }
            if (window.backend.lang === 'en-US') {
                xx = i18n.t('global.addQuota2');
            }
            return xx;
        },
        showNoAccessTip() {
            this.dispatchParent('alert', {
                content: i18n.t('global.youdont.have.power'),
            });
        },
        listExpiredArrearMsg(list, needCertify) {
            list = Array.isArray(list) ? list : [list];
            list.forEach((item, i) => {
                if (item.arrearInfo && item.arrearInfo.status && item.arrearInfo.status.toLowerCase() === 'arrear') {
                    // item.arrearInfo.type的值为'expired'、为空、或者其他值
                    this.getArrearMsg({ useCache: true, status: item.arrearInfo.type, needCertify })
                        .then((msg) => {
                            Vue.set(item, 'arrearMsg', msg);
                        });
                }
            });
        },
    },
};
