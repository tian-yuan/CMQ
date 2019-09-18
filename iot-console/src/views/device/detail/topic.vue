<template>
    <u-table-view :data="list" :loading="loading" layout="fixed" block-name="table">
        <u-table-view-column title="设备 Topic" class="f-toe">
            <div slot-scope="{row}">
                <span :title="row.Topic" v-clamp="2">{{ row.Topic }}</span>
            </div>
        </u-table-view-column>
        <u-table-view-column title="设备具有的操作权限">
            <div slot-scope="{row}">
                <span>{{ row.Operation | privilege }}</span>
            </div>
        </u-table-view-column>
        <u-table-view-column title="发布消息数">
            <div slot-scope="{row}">
                <span>{{ row.isCustom ? row.MessageCount : '-' }}</span>
            </div>
        </u-table-view-column>
        <u-table-view-column title="操作">
            <div slot-scope="{row}">
                <u-link-list>
                    <u-link-list-item v-if="row.isCustom" @click="publish(row)">发布消息</u-link-list-item>
                    <u-link-list-item v-else v-tooltip="'系统 Topic，不可发布消息'" disabled>发布消息</u-link-list-item>
                </u-link-list>
            </div>
        </u-table-view-column>
        <div slot="no-data-text" v-if="loadError">
            获取数据失败，请<u-link @click="refresh">重试</u-link>
        </div>
        <u-publish-message
            :show-modal.sync="showModal"
            :ProductKey="ProductKey"
            :Topic="crtTopic.Topic"></u-publish-message>
    </u-table-view>
</template>

<script>
import iothubService from 'module/iothub/services/index';
import filters from 'module/iothub/utils/filters';
import List from '@/blocks/sub.list.vue';
import UPublishMessage from 'module/iothub/components/u-publish-message.vue';

export default {
    components: {
        UPublishMessage,
    },
    filters,
    extends: List,
    tplMixins: [List],
    data() {
        return {
            hideHead: true,
            needRestorePage: true,
            ProductKey: this.$route.query.ProductKey,
            DeviceName: this.$route.query.DeviceName,
            showModal: false,

            crtTopic: {},
        };
    },
    methods: {
        loadList() {
            const { ProductKey, DeviceName } = this;
            const query = this.getFormForOAI({ ProductKey, DeviceName });
            return iothubService.QueryTopicList({
                query,
            })
                .then((res) => {
                    this.total = res.TotalCount;
                    this.list = res.TopicInfoList;
                });
        },
        publish(item) {
            const { Topic } = item;
            this.crtTopic = { Topic };
            this.showModal = true;
        },
    },
};
</script>
