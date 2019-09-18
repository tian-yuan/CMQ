<template>
    <div>
        <u-head-card :title="detail.Name || name" :infos="infoMessage"></u-head-card>
        <div>
            <u-tabs-data :data="tabs"></u-tabs-data>
            <div>
                <router-view></router-view>
            </div>
        </div>
    </div>
</template>

<style module>

</style>

<script>
import subscribe from '@/utils/mixins/subscribe';
export default {
    mixins: [subscribe],
    data() {
        return {
            detail: {},
        };
    },
};
</script>
