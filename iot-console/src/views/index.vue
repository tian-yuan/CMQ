<template>
<div :class="$style.root">
    <header :class="$style.head">
        <div :class="$style.wrap">
            <u-navbar>
                <u-navbar-item to="/overview">Overview</u-navbar-item>
                <!--<u-navbar-item to="/product">产品接入</u-navbar-item>-->
                <!--<u-navbar-item to="/device">设备接入</u-navbar-item>-->
                <div :class="$style.right">
                    <u-navbar-item href="https://vusion.github.io" target="_blank">文档</u-navbar-item>
                    <u-navbar-item href="https://github.com/tian-yuan" target="_blank">GitHub</u-navbar-item>
                </div>
            </u-navbar>
        </div>
    </header>
    <main :class="$style.body">
        <router-view></router-view>
    </main>
</div>
</template>

<script>
export default {};
</script>

<style module>
.head {
    position: fixed;
    top: 0;
    right: 0;
    left: 0;
    height: 68px;
    background: $brand-primary;
    color: white;
    box-shadow: 0 3px 4px rgba(0,0,0,0.1);
    padding: 0 10px;
    z-index: 1;
}

.wrap {
    width: 1180px;
    margin: 0 auto;
}

.left {
    float: left;
}

.right {
    float: right;
}

.body {
    margin-top: 68px;
}
</style>
