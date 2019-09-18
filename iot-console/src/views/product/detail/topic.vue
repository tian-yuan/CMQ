<template>
    <u-table-view :data="list" :loading="loading" layout="fixed" block-name="table">
        <u-table-view-column title="Topic 类" class="f-toe">
            <div slot-scope="{row}">
                <span :title="row.TopicName" v-clamp="2">{{ row.TopicName }}</span>
            </div>
        </u-table-view-column>
        <u-table-view-column title="操作权限">
            <div slot-scope="{row}">
                <span>{{ row.Operation | privilege }}</span>
            </div>
        </u-table-view-column>
        <u-table-view-column title="描述" class="f-toe">
            <div slot-scope="{row}">
                <span :title="row.Description" v-clamp="2">{{ row.Description | cross }}</span>
            </div>
        </u-table-view-column>
        <u-table-view-column title="操作">
            <div slot-scope="{row}">
                <u-link-list>
                    <u-link-list-item v-if="row.isCustom" @click="modify(row)">修改</u-link-list-item>
                    <u-link-list-item v-else v-tooltip="'系统 Topic 类，不可修改'" disabled>修改</u-link-list-item>

                    <u-link-list-item v-if="row.isCustom" @click="deleteItem(row)">删除</u-link-list-item>
                    <u-link-list-item v-else v-tooltip="'系统 Topic 类，不可删除'" disabled>删除</u-link-list-item>
                </u-link-list>
            </div>
        </u-table-view-column>
        <div slot="no-data-text" v-if="loadError">
            获取数据失败，请<u-link @click="refresh">重试</u-link>
        </div>
        <u-topic-modal
            :show-modal.sync="showModal"
            :ProductKey="ProductKey"
            :Operation="crtTopic.Operation"
            :TopicName="crtTopic.TopicName"
            :TopicId="crtTopic.TopicId"
            :Description="crtTopic.Description"
            @ok="refresh">
        </u-topic-modal>
    </u-table-view>
</template>

<script>
import iothubService from 'module/iothub/services/index';
import filters from 'module/iothub/utils/filters';
import List from '@/blocks/list.vue';
import UTopicModal from 'module/iothub/components/u-topic-modal.vue';

export default {
    filters,
    components: {
        UTopicModal,
    },
    extends: List,
    tplMixins: [List],
    data() {
        const self = this;
        return {
            create: {
                text: '定义 Topic 类',
                click() {
                    self.modify();
                },
            },
            outQuota: '该产品的自定义 Topic 类配额已经用完，' + this.countEnoughLabel(),
            needRestorePage: true,
            showModal: false,
            crtTopic: {},
            ProductKey: this.$route.query.ProductKey,
        };
    },
    methods: {
        modify(item = {}) {
            const { Operation, TopicName, TopicId, Description } = item;
            this.crtTopic = { Operation, TopicName, TopicId, Description };
            this.showModal = true;
        },
        loadList() {
            const { ProductKey } = this;
            const query = this.getFormForOAI({ ProductKey });
            return iothubService.QueryTopicClassList({
                query,
            })
                .then((res) => {
                    this.list = res.TopicClassInfoList;
                });
        },
        loadQuota() {
            const { ProductKey } = this;
            iothubService.QueryTopicClassQuota({
                query: { ProductKey },
            })
                .then(({ Quota, UsedQuota }) => {
                    this.quota = Quota;
                    this.total = UsedQuota;
                });
        },
        deleteItem(item) {
            const { ProductKey } = this;
            const { TopicId } = item;
            return this.$confirm({
                content: '确定删除该 Topic 类吗?',
                submit() {
                    return iothubService.DeleteTopicClass({
                        query: { ProductKey, TopicId },
                        config: { noAlert: true },
                    });
                },
            });
        },
    },
};
</script>
