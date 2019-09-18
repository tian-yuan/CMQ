import Vue from 'vue';
import VueRouter from 'vue-router';
Vue.use(VueRouter);

window.backend = {};
window.backend.lang = 'zh';

import * as Components from 'library';
import { install } from 'vusion-utils';
install(Components, Vue);

import routes from './routes';

new Vue({
    el: '#app',
    router: new VueRouter({
        routes,
    }),
});
