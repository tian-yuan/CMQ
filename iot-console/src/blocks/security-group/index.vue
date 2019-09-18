<template>
    <div class="m-box">
        <div :class="$style.wrap">
            <u-linear-layout>
                <u-button color="primary" icon="create" :disabled="!hasQuota || disabled || !networkId" @click="add">{{ $t('global.vpc.sgAdd') }}</u-button>
                <span v-if="!hasQuota && !disabledTip" class="u-enoughLabel">{{ $t('global.sg.quotaTip',[serviceTitle]) }}{{ countEnoughLabel() }}</span>
                <span v-if="disabledTip" class="u-enoughLabel">{{ $t('global.blocks.workTip.security',[serviceTitle]) }}</span>
            </u-linear-layout>
        </div>
        <u-table-view :data="list" :loading="loading" layout="fixed">
            <u-table-view-column :title="$t('global.security.name')" width="20%">
                <div slot-scope="props" class="f-toe">
                    <i v-if="props.row.IsDefault" v-tooltip.top="$t('global.sg.defaultTip')" class="u-securityGroup-tip"></i>
                    <span :title="props.row.Name || '-'">
                        {{ props.row.Name || '-' }}
                    </span>
                </div>
            </u-table-view-column>
            <u-table-view-column :title="$t('global.tab.security.id')" width="30%">
                <div slot-scope="props" class="f-toe  f-toe-rtl">
                    {{ props.row.Id }}
                </div>
            </u-table-view-column>
            <u-table-view-column :title="$t('global.description')" width="30%">
                <div slot-scope="props" class="f-toe">
                    {{ props.row.Description || '-' }}
                </div>
            </u-table-view-column>
            <u-table-view-column :title="$t('global.label.op')" width="20%">
                <div slot-scope="props">
                    <u-linear-layout>
                        <u-link v-if="hasSecurityRuleLink" :to="{ path:'/vpc/securitygroup/rule', query: {vpcId: props.row.VpcId, sgId: props.row.Id} }">{{ $t('global.vpc.sgManage') }}</u-link>
                        <u-link @click="deleteItem(props.row)" v-tooltip="total === 1 ? $t('global.vpc.sg.deleteTip') : ''" :disabled="total === 1 || props.row.deleting || removeDisabled">{{ $t('global.vpc.sg.deleteTitle') }}</u-link>
                    </u-linear-layout>
                </div>
            </u-table-view-column>
            <div slot="no-data-text">
                <div>{{ $t('global.prompt.info.empty') }}</div>
            </div>
        </u-table-view>
        <div class="pager" v-if="total >= limitList[0] && totalPage >= 1">
            <u-linear-layout direction="vertical">
                <u-pagination-pro :limitList="limitList" :limit="form.limit" :total="totalPage" :page="page" @change="changePage" @changeLimit="changeLimit"></u-pagination-pro>
            </u-linear-layout>
        </div>
        <security-group-add @done="refresh" modal-name="security.group.add" v-bind="modalData"></security-group-add>
    </div>
</template>

<style module>
.wrap {
    margin-bottom: 20px;
}
</style>

<script>
import Bridge from '@/utils/Bridge';
import Pager from 'mixins/page';
import Auth from 'mixins/auth';
import service from '@/views/dashboard/vpc/services/index';
import securityGroupAdd from './security-group-add.vue';
export default {
    components: {
        'security-group-add': securityGroupAdd,
    },
    mixins: [Pager, Auth],
    data() {
        return {
            needRestorePage: true,
        };
    },
    computed: {
        modalData() {
            return {
                formData: {
                    InstanceId: this.instanceId,
                    InstanceType: this.serviceName,
                },
                selected: this.list.map((item) => item.Id),
                networkId: this.networkId,
                serviceTitle: this.serviceTitle,
                titleTip: this.titleTip,
                addQuota: this.hasQuota,
            };
        },
    },
    watch: {
        instanceId(value) {
            this.loadList();
        },
    },
    methods: {
        loadQuota() {
            service.GetQuota().then(({ Quota }) => {
                const { InstanceInSecurityGroupQuota } = Quota;
                this.quota = InstanceInSecurityGroupQuota;
            });
        },
        loadList() {
            if (!this.instanceId) {
                // mutil line
                return true;
            }
            return service.ListInstanceSecurityGroupsAll({
                query: this.getFormForOAI({
                    InstanceId: this.instanceId,
                    InstanceType: this.serviceName,
                }),
            })
                .then(({ SecurityGroups, Count }) => {
                    this.total = Count;
                    this.list = SecurityGroups;
                });
        },
        deleteItem(item) {
            return new Promise((res, rej) => {
                this.dispatchParent('confirm', {
                    title: `${this.$t('global.vpc.sg.deleteTitle')}`,
                    content: `${this.$t('global.vpc.sg.deleteContent', [this.serviceTitle])}`,
                }, (err, status) => {
                    if (status) {
                        service.LeaveSecurityGroup({
                            body: {
                                InstanceId: this.instanceId,
                                InstanceType: this.serviceName,
                                SecurityGroupIds: [item.Id],

                            },
                        }).then(res, rej);
                    } else {
                        rej();
                    }
                });
            });
        },
        add() {
            this.$modal.show('security.group.add');
        },
    },
};
</script>

