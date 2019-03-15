<template>
<u-linear-layout :class="$style.root" direction="vertical">
    <u-linear-layout>
        <u-link href="#">链接</u-link>
        <u-link href="#">查看详情</u-link>
        <u-label>Label</u-label>
        <u-label color="primary">Label</u-label>
    </u-linear-layout>
    <u-linear-layout>
        <u-button color="primary">创建</u-button>
        <u-button>取消</u-button>
        <span>消息 <u-badge :value="3"></u-badge></span>
        <u-badge :value="3" corner>
            <u-button>评论</u-button>
        </u-badge>
    </u-linear-layout>
    <u-linear-layout>
        <u-capsules>
            <u-capsule>Apple</u-capsule>
            <u-capsule>Banana</u-capsule>
            <u-capsule>Cake</u-capsule>
        </u-capsules>
        <u-pills>
            <u-pill>Apple</u-pill>
            <u-pill>Banana</u-pill>
            <u-pill>Cake</u-pill>
        </u-pills>
        <u-tablets>
            <u-tablet>Apple</u-tablet>
            <u-tablet>Banana</u-tablet>
            <u-tablet>Cake</u-tablet>
        </u-tablets>
    </u-linear-layout>
    <u-tabs>
        <u-tab title="详细信息"><u-loading></u-loading></u-tab>
        <u-tab title="性能监控"><u-loading></u-loading></u-tab>
        <u-tab title="操作日志"><u-loading></u-loading></u-tab>
    </u-tabs>
</u-linear-layout>
</template>

<style module>
.root {
    width: 1180px;
    margin: 0 auto;
    padding: 40px;
}
</style>
