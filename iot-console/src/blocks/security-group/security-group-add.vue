<template>
    <u-modal-fix class="form-error-block" @close="close" ok-button="" cancel-button="" :visible.sync="show" size="modallist">
        <template slot="title">
            {{ $t('global.vpc.sgAdd') }}
        </template>
        <div v-if="titleTip" :class="$style.tip" style="margin-top: -15px;margin-bottom: 25px;">
            <u-status-icon name="warning" size="14"></u-status-icon>{{ titleTip }}
        </div>
        <div style="margin-top: -15px;margin-bottom: 15px;">
            <u-linear-layout style="display: inline-block;">
                <u-submit-button :click="addList.bind(this)" inline :hideMessage="true">
                    <template slot-scope="scope">
                        <u-button @click="scope.submit" color="primary" :icon="scope.submitting ? 'loading' : ''"
                            :disabled="scope.submitting || !securityGroups.length || addLeftQuota < 0">{{ $t('global.vpc.sg.mutilIn') }}</u-button>
                    </template>
                </u-submit-button>
                <span :class="$style.subtitle">{{ $t('global.moveIn.less.security',[addLeftQuota === true ? '-' : addLeftQuota]) }}</span>
            </u-linear-layout>
            <u-linear-layout style="float:right">
                <u-input :placeholder="$t('global.input.sgName')" size="medium medium" v-on="listeners"></u-input>
            </u-linear-layout>
        </div>
        <u-table-view :data="list" :loading="loading" layout="fixed" @selection-change="selectionChange($event)">
            <u-table-view-column type="selection"></u-table-view-column>
            <u-table-view-column :title="$t('global.security.name')" width="30%">
                <template slot-scope="scope">
                    <u-text-hide :text="scope.row.Name"></u-text-hide>
                </template>
            </u-table-view-column>
            <u-table-view-column :title="$t('global.tab.security.id')" width="30%">
                <template slot-scope="scope">
                    <u-text-hide rtl :text="scope.row.Id"></u-text-hide>
                </template>
            </u-table-view-column>
            <u-table-view-column :title="$t('global.description')" width="20%">
                <template slot-scope="scope">
                    <u-text-hide :text="scope.row.Description"></u-text-hide>
                </template>
            </u-table-view-column>
            <u-table-view-column :title="$t('global.label.op')" width="100px">
                <template slot-scope="scope2">
                    <u-submit-button inline :click="addItem.bind(this, scope2.row)" :hideMessage="true">
                        <template slot-scope="scope">
                            <u-link @click="scope.submit" :disabled="scope.submitting || scope2.row.disabled || addQuota <= 0">{{ $t('global.link.movein') }}</u-link>
                        </template>
                    </u-submit-button>
                </template>
            </u-table-view-column>
            <div slot="no-data-text" v-if="!loadError">
                <template v-if="search">
                    {{ $t('global.search.empty') }}
                </template>
                <template v-else>
                    {{ $t('global.empty2.now',[name]) }}
                </template>
            </div>
            <div slot="no-data-text" v-else>
                <i18n path="global.load.base2">
                    <u-link @click="refresh">{{ $t('global.load.error.retry') }}</u-link>
                    <span>{{ name }}</span>
                </i18n>
            </div>
        </u-table-view>
        <div class="leftfooterwrapper">
            <u-desc :class="$style.tip2">
                <u-text color="secondary">{{ $t('global.vpc.sg.newTip.base') }}<u-link :to="networkId?'/vpc/detail/securitygroup?vpcId='+networkId:'/vpc/list'" target="_blank" :class="$style.link">{{ $t('global.clickHere3') }}</u-link>
                    <span style="line-height: 24px;">
                        <u-refresh @click="refresh" :loading="loading"></u-refresh>
                    </span>
                </u-text>
            </u-desc>
        </div>
        <div class="pager" style="float:right;margin:0;" v-if="totalPage > 1">
            <u-linear-layout direction="vertical">
                <u-pagination :total="totalPage" :page="page" @select="changePage"></u-pagination>
            </u-linear-layout></div>
    </u-modal-fix>
</template>
<style module>
.tip2{
    margin: 0!important;
    line-height: 30px;
}
.tip {
    background: #fbf7cf;
    color: #dfb050;
    border-color: #f6e0c4;
    padding: 6px 14px;
    margin-bottom: 15px;
}
.subtitle {
    margin-left: -4px;
    font-size: 12px;
    color: #aaa;
}
</style>
<script>
import Modal from '@/utils/mixins/modal/base';
import Page from '@/utils/mixins/page';
import services from '@/views/dashboard/vpc/services/index';
import modalFix from '@/components/common/u-modal-fix.vue';
import Vue from 'vue';
import TableViewFix from '@/components/common/u-table-view-fix.vue';
import TableViewFixColumn from '@/components/common/u-table-view-fix-column.vue';
export default {
    components: {
        'u-modal-fix': modalFix,
        'u-table-view': TableViewFix,
        'u-table-view-column': TableViewFixColumn,
    },
    mixins: [Modal, Page],
    props: {
        formData: Object,
        networkId: String,
        titleTip: String,
        serviceTitle: String,
        selected: Array,
        addQuota: {
            type: [Boolean, Number],
            default: true,
        },
    },
    data() {
        return {
            name: `${this.$t('global.canMoveIn.security')}`,
            securityGroups: [],
            list: [],
            search: '',
            searchInner: '',
            listeners: {
                blur: ($event) => {
                    this.search = this.searchInner;
                },
                input: ($event) => {
                    this.searchInner = $event;
                },
                keyup: ($event) => {
                    if ($event.keyCode === 13) {
                        this.search = this.searchInner;
                    }
                },
            },
            ...this.initLoadStatus('load'),
        };
    },
    computed: {
        canSubmit() {
            const { addLeftQuota } = this;
            return !!this.securityGroups.length && addLeftQuota >= 0;
        },
        addLeftQuota() {
            const securityGroups = this.securityGroups;
            return this.addQuota === true ? true : this.addQuota - securityGroups.length;
        },
    },
    watch: {
        show() {
            if (this.show) {
                this.securityGroups = [];
                this.list = [];
                this.loadList();
            }
        },
        search() {
            this.reset();
            this.refresh();
        },
        selected(selected, oldSelected) {
            const ifChange = JSON.stringify(selected.sort()) !== JSON.stringify(oldSelected.sort());
            if (selected && selected.length && ifChange) {
                this.loadList();
            }
        },
    },
    methods: {
        addItem(item) {
            return this.addList([item]);
        },
        addList(list) {
            return services.JoinSecurityGroup({
                body: {
                    ...this.formData,
                    SecurityGroupIds: (list || this.securityGroups).map((item) => item.Id),
                },
            }).then(() => {
                this.$emit('done');
            });
        },
        selectionChange(securityGroups) {
            this.securityGroups = securityGroups;
        },
        loadList() {
            if (!this.show) {
                return true;
            }
            const Name = this.search;
            return this.addLoadStatus(services.ListSecurityGroup({
                query: this.getFormForOAI(Name ? {
                    VpcId: this.networkId,
                    Name,
                    FuzzySearch: true,
                } : {
                    VpcId: this.networkId,
                }),
            }).then(({ SecurityGroups, Count }) => {
                if (this.search !== Name) {
                    return;
                }
                SecurityGroups.forEach((item) => {
                    if (this.selected.includes(item.Id) || this.addQuota <= 0) {
                        item.disabled = true;
                    }
                });
                this.list = SecurityGroups;
                this.total = Count;
            }), 'load');
        },
    },
};
</script>

