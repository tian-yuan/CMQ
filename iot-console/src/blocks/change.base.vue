<script>
import Tpl from 'mixins/tpl';
import billNewService from '@/services/billnew/index.js';
import OrderService from '@/services/order/index.js';
import i18n from '@/utils/i18n';
import Bridge from '@/utils/Bridge';

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
            standard: undefined,
            instanceName: '',
            currentValue: 1,
            terminateMethod: '',
            timeLable: `${this.$t('global.pay.period')}`,
            seletcValue: [
                { text: `${this.$t('global.terminateMethod.unselect')}`, value: '' },
                { text: `${this.$t('global.terminateMethod.delete')}`, value: 'DELETE' },
                { text: `${this.$t('global.terminateMethod.persist')}`, value: 'PERSIST' },
            ],
            rules: {
                TerminateMethod: [
                    { type: 'string', message: i18n.t('nvm.endData.required'), validator: (rule, value, callback) => {
                        !this.terminateMethod ? callback(new Error()) : callback();
                    } },
                ],
            },
            needTerminateMethod: false,
            listUrl: '/m/', // 无需支付直接跳到列表页url
        };
    },
    created() {
        this.instanceName = this.$route.query.name;
        this.service = this.$route.path.split('/')[1];
        this.getItem();
    },
    methods: {
        getItem() {
            billNewService
                .getInstanceChargeInfo({
                    query: {
                        Service: this.product,
                        InstanceId: this.$route.query.id,
                    },
                })
                .then((result) => {
                    this.getStandard(result);
                });
        },
        getStandard(result) {
            this.standard = JSON.parse(result.Standard);
        },
        changeValue(value) {
            this.currentValue = value.value;
        },
        selectValue(value) {
            this.terminateMethod = value.value;
        },
        getBody(standard) {
            const Body = [];
            return Body;
        },
        showOrder(result) {
            if (result.isDirectPay) { // 没计费的环境直接跳转
                Bridge.send('parent', 'urlchange', this.listUrl);
            } else {
                Bridge.send('parent', 'urlchange', '/order/payment?orderId=' + result.OrderId);
            }
        },
        submit() {
            const terminateMethod = this.terminateMethod === '' ? 'PERSIST' : this.terminateMethod;
            const { product, standard, currentValue, instanceName } = this;
            const Body = this.getBody(standard);
            const body = {
                Body,
                Service: product,
                InstanceId: this.$route.query.id,
                Period: currentValue,
                InstanceName: instanceName,
                RegionId: this.$route.query.regionId,
            };
            if (this.needTerminateMethod && terminateMethod) {
                Object.assign(body, {
                    Extra: { TerminateMethod: terminateMethod },
                });
            }
            if (window.DATracker) {
                DATracker.track('upgrade', {
                    product: this.service,
                    stage: 'submit',
                    type: 'convert',
                });
            }
            return OrderService.ModifyChargeType({ body }).then((result) => {
                this.showOrder(result);
            });
        },
    },
    getRender(map) {
        return (
            <u-form gap="large" ref="form">
                <u-form-item label={`${this.$t('global.instance.name.label')}`} layout="block" class={this.$style.config}>
                    {this.instanceName}
                </u-form-item>
                <u-form-item label={`${this.$t('global.instance.configuration.label')}`} layout="block">
                    {map.configuration}
                </u-form-item>
                <u-form-item label={this.timeLable} required>
                    {vIf(this.standard, <u-purchase-period-capsules value={1} product={this.product} standard={this.standard} onChange={this.changeValue}></u-purchase-period-capsules>)}
                </u-form-item>
                {vIf(this.needTerminateMethod, <u-form-item label={`${this.$t('global.terminateMethod.label')}`} rules={this.rules.TerminateMethod} required>
                    <u-select onSelect={this.selectValue} value={this.terminateMethod}>
                        {this.seletcValue.map((item, index) => (<u-select-item value={item.value}>{item.text}</u-select-item>))}
                    </u-select>
                </u-form-item>)}
                <u-form-item>
                    {map.purchaseCard}
                </u-form-item>
                <u-form-item>
                    <u-submit-button click={this.submit.bind(this)} autoFocus={true} { ...{
                        scopedSlots: {
                            default: (scope) => (
                                <u-button color="primary" icon={scope.submitting ? 'loading' : ''} disabled={scope.submitting} onClick={scope.submit}>
                                    {this.$t('global.buy.immediately')}
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
