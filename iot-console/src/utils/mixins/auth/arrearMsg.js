import i18n from '@/utils/i18n';
import authService from '@/services/global/auth';
import { cache } from '@/utils/cache';
import configAuth from '@/utils/config/auth';

const MSG = {
    arrear: configAuth.isSub ? i18n.t('global.master.moneyOut') : i18n.t('global.moneyOut'),
    certify: i18n.t('global.auth.notPass'),
    expired: i18n.t('global.errorTip.prepay'),
};

export default {
    data() {
        return {
            arrearMsgGlobal: '',
        };
    },
    methods: {
        getCertInfo(options) {
            if (options.needCertify === false)
                return Promise.resolve(MSG.arrear);
            else {
                const useCache = options.useCache;
                const prePomise = cache.get('arrearMsgCertifyInfo');
                if (useCache && prePomise)
                    return prePomise;
                else {
                    const cachetemp = authService.getCertInfo().then((result) => {
                        const certify = result.realCertify;
                        return certify ? MSG.arrear : MSG.certify;
                    });
                    cache.add('arrearMsgCertifyInfo', cachetemp, true);
                    return cachetemp;
                }
            }
        },
        getArrearMsg(options = {}) {
            return this.getCertInfo(options).then((result) => {
                const msg = options && options.status === 'expired' ? MSG.expired : result;
                return i18n.t('global.auth.stopServer', [msg]);
            });
        },
        clearCache() {
            cache.clear('arrearMsgCertifyInfo');
        },
        initArrearMsg(options = {}) {
            this.clearCache();
            this.getArrearMsg(options).then((msg) => {
                this.arrearMsgGlobal = msg;
                return msg;
            });
        },
    },
};
