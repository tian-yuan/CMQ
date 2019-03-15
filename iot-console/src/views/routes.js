import Index from './index.vue';
import Main from './main.vue';
import Overview from './overview.vue';

export default [
    { path: '/', component: Index, children: [
        { path: '', component: Main },
        { path: 'overview', component: Overview },
    ] },
];
