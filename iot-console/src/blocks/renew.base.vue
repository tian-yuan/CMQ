<script>
import Tpl from 'mixins/tpl';
import billNewService from '@/services/billnew/index.js';
import OrderService from '@/services/order/index.js';
import Bridge from '@/utils/Bridge';
import Vue from 'vue';

const getTpl = function (tpl) {
    return typeof tpl === 'function' ? tpl() : tpl;
};
const vIf = function (data, tplTrue, tplFalse) {
    return data ? getTpl(tplTrue) : getTpl(tplFalse);
};
export default {
    mixins: [Tpl],
    data() {
        return {
            product: '',
            detail: {},
            standard: undefined,
            instanceName: '',
            expirationTime: undefined,
            proExpirationTime: undefined,
            currentValue: 1,
            timeLable: `${this.$t('global.instance.renewPeriod.label')}`,
        };
    },
    created() {
        this.getItem();
    },
    methods: {
        getItem() {
            billNewService
                .getBillInfo({
                    query: {
                        Service: this.product,
                        InstanceId: this.$route.query.id,
                    },
                })
                .then((result) => {
                    this.expirationTime = result.EndTime;
                    this.getStandard(result);
                });
        },
        getStandard(result) {
            this.standard = JSON.parse(result.Standard);
            this.instanceName = result.InstanceName;
        },
        renewTime(value) {
            this.currentValue = value.value;
            if (!this.expirationTime) {
                this.proExpirationTime = undefined;
            } else {
                const expirationTime = new Date(this.expirationTime);
                let year = expirationTime.getFullYear();
                let month = expirationTime.getMonth();
                if (this.currentValue < 10) {
                    month = month + this.currentValue;
                    expirationTime.setMonth(month);
                } else {
                    year = year + this.currentValue / 12;
                    expirationTime.setFullYear(year);
                }
                this.proExpirationTime = expirationTime;
            }
        },
        getBody() {
            const Body = [];
            return Body;
        },
        renewBridge(result) {
            if (result.isDirectPay) { // 没计费的环境跳转到列表页
                Bridge.send('parent', 'urlchange', this.listUrl);
            } else {
                Bridge.send('parent', 'urlchange', '/order/payment?orderId=' + result.OrderId);
            }
        },
        submit() {
            const { product, standard, currentValue } = this;
            const Body = this.getBody(standard);
            const body = {
                Body,
                Service: product,
                InstanceId: this.instanceId || this.$route.query.id,
                Period: currentValue,
                RegionId: this.$route.query.regionId,
            };
            return OrderService.CreateRenewOrder({ body }).then((result) => {
                this.renewBridge(result);
            });
        },
    },
    getRender(map) {
        return (
            <u-form gap="large" ref="form">
                {vIf(!this.hideInstanceName, () => <u-form-item label={`${this.$t('global.instance.name.label')}`} layout="block" class={this.$style.config}>
                    {this.instanceName}
                </u-form-item>)}
                <u-form-item label={`${this.$t('global.instance.configuration.label')}`} layout="block">
                    {map.configuration}
                </u-form-item>
                <u-form-item label={`${this.$t('global.instance.expriedTimeBeforRenew.label')}`} layout="block">
                    {vIf(this.expirationTime, <div class="f-mb10" v-if="expirationTime"><span class={this.$style.config}>{Vue.filter('dateFormat')(this.expirationTime)}</span></div>)}
                </u-form-item>
                <u-form-item label={`${this.$t('global.instance.renewPeriod.label')}`} layout="block">
                    {vIf(this.expirationTime && this.standard, <u-purchase-period-capsules value={1} product={this.product} standard={this.standard} onChange={this.renewTime}></u-purchase-period-capsules>)}
                    {vIf(this.proExpirationTime, <div class="f-mb10 f-mt10"><span class={this.$style.config}>{this.$t('global.newRenewTime')} {Vue.filter('dateFormat')(this.proExpirationTime)}</span></div>)}
                </u-form-item>
                <u-form-item>
                    {map.purchaseCard}
                </u-form-item>
                <u-form-item>
                    <u-submit-button click={this.submit.bind(this)} { ...{
                        scopedSlots: {
                            default: (scope) => (
                                <u-button color="primary" icon={scope.submitting ? 'loading' : ''} onClick={scope.submit}>
                                    {this.$t('global.instance.renew.immediately')}
                                </u-button>
                            ),
                        },
                    } }>
                    </u-submit-button>
                </u-form-item>
            </u-form>
        );
    },
};
</script>
<style module="$style">
.config {
  color: #999;
  margin-right: 5px;
}
</style>
