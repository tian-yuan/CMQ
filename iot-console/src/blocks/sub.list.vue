<template>
    <div block-name="$head" v-show="!hideHead">
        <div class="tit">
            <u-linear-layout style="display: inline-block;">
                <template v-if="create">
                    <u-button color="primary" :to="create.url" @click="(create.click || createItem || noop)()" icon="create" :disabled="!hasQuota || create.disabled">{{ create.text }}</u-button>
                </template>
                <u-button square icon="refresh" @click="refresh" v-if="needRefresh"></u-button>
                <u-button v-if="action" v-bind="action.options" @click="action.handle" v-for="action in actions" :key="action">{{ action.text }}</u-button>
                <span :class="$head.quotaHint" v-if="create && !hasQuota">{{ outQuota }}</span>
            </u-linear-layout>
            <u-linear-layout style="float:right" v-if="handlers && handlers.length" class="righthead">
                <template v-for="action in handlers">
                    <component v-if="action.exist !== false && action.text" v-bind="action.options" v-on="action.listeners" :is="action.type" :key="action">{{ action.text }}</component>
                    <component v-else-if="action.exist !== false" v-bind="action.options" v-on="action.listeners" :is="action.type" :key="action"></component>
                </template>

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
import ListBase from './list.base.vue';

export default {
    extends: ListBase,
    data() {
        return {
            create: undefined,
            hideHead: false,
        };
    },
    methods: {
        noop() {
            // noop
        },
    },
};
</script>
