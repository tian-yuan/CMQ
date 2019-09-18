import i18n from '@/utils/i18n';
import { cache } from '../cache';
export const cacheInit = function (router) {
    cache.add('tagCache', {});
    router.beforeEach((to, from, next) => {
        if (to.path !== from.path) {
            const tagCache = cache.get('tagCache');
            const [, fromKey] = /\/([^/]+)/.exec(from.path) || [];
            const [, toKey] = /\/([^/]+)/.exec(to.path) || [];
            if (toKey !== fromKey) {
                Object.keys(tagCache).forEach((item) => {
                    tagCache[item].data = [];
                    tagCache[item].newData = [];
                    tagCache[item].LineIndex = undefined;
                });
            } else if (tagCache[from.name]) {
                if (!tagCache[from.name].pages.some((item) => to.path.indexOf(item) === 0)) {
                    tagCache[from.name].data = [];
                    tagCache[from.name].newData = [];
                    tagCache[from.name].LineIndex = undefined;
                }
            } else if (!tagCache[from.name]) {
                Object.keys(tagCache).forEach((item) => {
                    if (!tagCache[item].pages.some((item) => to.path.indexOf(item) === 0)) {
                        tagCache[item].data = [];
                        tagCache[item].newData = [];
                        tagCache[item].LineIndex = undefined;
                    }
                });
            }
        }
        next();
    });
};
export default {
    data() {
        const name = this.$route.name;
        if (DEV) {
            if (!name) {
                throw new Error(i18n.t('global.rule.router.required'));
            }
        }
        const tagCache = cache.get('tagCache');
        const selfPage = this.$route.path;
        if (!tagCache[name]) {
            tagCache[name] = {
                data: [],
                pages: this.$options.tagCachePage ? [...this.$options.tagCachePage, selfPage] : [selfPage],
            };
        }
        const curPage = tagCache[name];
        curPage.selected = undefined;
        return {
            searchTags: curPage,
        };
    },
};
