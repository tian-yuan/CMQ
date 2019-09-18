<template>
    <div block-name="$head">
        <u-notice-user v-if="showUserNotice"></u-notice-user>
        <u-notice v-if="tip" :icon="tip.status" :color="tip.status">
            {{ tip.text }}
        </u-notice>
        <u-service-notice v-if="notice" :service-name="notice.serviceName"></u-service-notice>
        <u-page-sum v-if="suggest && ownDocs">
            <p>
                {{ suggest.text }}
                <u-link-support :hash="link.hash" v-for="(link, index) in suggest.links" :key="index" @click="link.click">{{ link.text }}</u-link-support>
            </p>
        </u-page-sum>
        <u-page-sum v-if="noCustomeSuggest && ownDocs">
            <p>
                {{ noCustomeSuggest.text }}
                <u-link-support :hash="link.hash" v-for="(link, index) in noCustomeSuggest.links" :key="index">{{ link.text }}</u-link-support>
            </p>
        </u-page-sum>
        <u-tabs-data :data="tabs" v-if="tabs"></u-tabs-data>
        <div class="tit">
            <u-linear-layout style="display: inline-block;">
                <u-button v-if="$options.auth.isSub && !subAccess" icon="create" color="primary" @click="showNoAccessTip">{{ create.text }}</u-button>
                <u-button v-else color="primary" :to="create.url && {path: create.url, query: create.query}" icon="create" :disabled="!hasQuotab || create.disabled" @click="create.click" v-track="create.track">{{ create.text }}</u-button>
                <template v-if="actions">
                    <u-button v-bind="action.options" @click="action.handle" v-for="action in actions" :key="action">{{ action.text }}</u-button>
                </template>
                <u-button square icon="refresh" @click="refresh"></u-button>
                <span :class="$head.quotaHint" v-if="!hasQuotab">{{ outQuota }}</span>
            </u-linear-layout>
            <u-linear-layout style="float:right;margin-bottom:10px;" v-if="handlers && handlers.length" class="righthead">
                <component v-if="action.exist !== false" v-bind="action.options" v-on="action.listeners" @click="action.handle ? action.handle($event) : ()=>{}" :is="action.type" v-for="action in handlers" :key="action" v-track="action.track">{{ action.text }}</component>
            </u-linear-layout>
        </div>
        <div class="tit" v-if="line2Handlers && line2Handlers.length" style="text-align: right;margin-top: -10px;padding-bottom: 10px;">
            <u-linear-layout class="righthead">
                <component v-if="action.exist !== false" v-bind="action.options" v-on="action.listeners" @click="action.handle" :is="action.type" v-for="action in line2Handlers" :key="action">{{ action.text }}</component>
            </u-linear-layout>
        </div>
    </div>

</template>

<style module="$head">
.quotaHint {
    font-size: 12px;
    color: #666;
}
</style>

<script>
import Auth from '../utils/mixins/auth/index';
import Page from '../utils/mixins/page';
import ListBase from './list.base.vue';

export default {
    extends: ListBase,
    data() {
        return {
            // tip: {
            //     status: 'warning',
            //     text: '提示',
            // }, // 提示,黄色的
            // create: {
            //     text: '创建 VPC',
            //     url: '/vpc/create',
            //     disabled: false,
            // }, // 创建信息
            // outQuota: `VPC 配额已经用完，${this.countEnoughLabel()}`,
            // suggest: {
            //     text: '这里汇聚了你在网易云基础服务上的所有 VPC。',
            //     links: [
            //         {
            //             text: '帮助文档',
            //             hash: '83641380367421440',
            //         },
            //     ],
            // }, // 建议
            // tabs: {
            //     children: [
            //         {
            //             text: 'a',
            //             router: '/a',
            //         },
            //         {
            //             text: 'b',
            //             router: {
            //                 path: '/b',
            //                 query: {

            //                 },
            //             },
            //         }
            //     ],
            // }, // 路由
            // subAccess: false, // 子帐号访问
            showUserNotice: false,
            tabs: undefined,
            suggest: undefined,
            create: undefined,
            tip: undefined,
            notice: undefined, // 产品模块公告
            ownDocs: this.$options.auth.ownDocs,
            IS_CUSTOM: this.$options.auth.custom,
            IS_PRTZ: this.$options.auth.isPrtz,
        };
    },
};
</script>
