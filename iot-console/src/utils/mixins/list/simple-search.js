import tagCache from '@/utils/mixins/tagCache';
import USearchInput from './u-search-input.vue';

/**
 * 列表页的简单搜索框
 * 在使用该mixin的页面，可访问this.Keyword，或调用this.updateKeyword()
 * 列表页loadList方法中，需要用this.isBeforeKeywordLoaded()判断是否需要延迟加载
 * @param {Array} tagCachePage  需要缓存搜索关键字的页面
 * @param {String}  placeholder 搜索输入框的placeholder
 */
export default function simpleSearch({ tagCachePage, placeholder }) {
    return {
        components: {
            USearchInput,
        },
        mixins: [tagCache],
        tagCachePage,
        data() {
            const self = this;
            return {
                Keyword: '',
                handlers: [{
                    type: 'u-search-input',
                    options: {
                        placeholder,
                        value: '',
                    },
                    listeners: {
                        'enter'(event) {
                            if (event)
                                self.filterProductKey = '';

                            self.updateKeyword(event);
                            self.resetToPage1();
                        },
                    },
                }],
            };
        },
        created() {
            const Keyword = this.searchTags.data[0]; // 读取关键字
            if (Keyword) {
                this.updateKeyword(Keyword);
                this.loadList();
            }
        },
        methods: {
            updateKeyword(Keyword = '') {
                this.Keyword = Keyword;
                this.handlers[0].options.value = Keyword;
                this.searchTags.data[0] = Keyword; // 保存关键字
            },
            isBeforeKeywordLoaded() {
                return this.searchTags.data[0] && !this.Keyword;
            },
        },
    };
}
