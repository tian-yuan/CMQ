<template>
<section :class="$style.banner">
    <div :class="$style.bannerContent">
        <h1 :class="$style.heading">Hello World！</h1>
        <h2 :class="$style.subtitle">Welcome to Vusion App</h2>
    </div>
</section>
</template>

<style module>
.banner {
    height: 480px;
    background: #f0f6fa;
    color: #676975;
}

.bannerContent {
    width: 1180px;
    margin: 0 auto;
    padding: 80px;
}

.heading {
    font-size: 52px;
    font-weight: normal;
    margin-bottom: 20px;
}

.subtitle {
    font-size: 28px;
    font-weight: normal;
    margin: 0;
}
</style>
