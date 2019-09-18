<template>
    <div>
        <u-linear-layout>
            <u-select :value="limit" @select="selectLimit" :class="$style.select">
                <u-select-item v-for="item in limitList" :key="item" :value="item">
                    {{ item }}
                </u-select-item>
            </u-select>
            <span :class="$style.unit">{{ $t('global.bar.perPage') }}</span>
            <div :class="$style.pagerbox">
                <u-pagination :total="total" :page="page" @select="changePage"></u-pagination>
            </div>
        </u-linear-layout>
    </div>
</template>

<style module>
.select{
    text-align: left;
    width: 80px;
    height: 30px;
    line-height: 28px;
    vertical-align: bottom;
}
.select > ul{
    width: 80px;
}
.unit{
    margin-left: -10px;
}
.pagerbox{
    display: inline-block;
}
</style>

<script>
export default {
    name: 'u-pagination-pro',
    props: {
        total: { type: Number, default: 0 },
        page: { type: Number, default: 1 },
        limitList: Array,
        limit: Number,
    },
    methods: {
        selectLimit(event) {
            this.$emit('changeLimit', event);
        },
        changePage(event) {
            this.$emit('change', event);
        },
    },
};
</script>
