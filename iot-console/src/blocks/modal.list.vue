<script>
/**
 * hasInput: 使用简单input框
 * startSearch: input搜索框变化后回调
 * inputValue: input值
 * errorMsg: 请求失败，设置错误信息
 * */
import Auth from 'mixins/auth';
import Page from 'mixins/page';
import Tpl from 'mixins/tpl';
import Modal from 'mixins/modal/base';
import searchTag from '@/components/common/search/index.vue';
import modalFix from '@/components/common/u-modal-fix.vue';

const getTpl = function (tpl) {
    return typeof tpl === 'function' ? tpl() : tpl;
};
const vIf = function (data, tplTrue, tplFalse) {
    return data ? getTpl(tplTrue) : getTpl(tplFalse);
};
export default {
    components: {
        'search-tag': searchTag,
        'u-modal-fix': modalFix,
    },
    mixins: [Auth, Page, Tpl, Modal],
    loadListWhenCreated: false, // 在loadList里return true来控制更好
    data() {
        return {
            errorMsg: null,
            hasSearchTag: false,
            hasInput: false,
            inputValue: '',
            oldInputValue: '', // 当前结果对应的搜索值
        };
    },
    watch: {
        errorMsg(value) {
            if (value)
                this.dispatchParent('error', value);
        },
        show(value) {
            if (value)
                this.loadList && this.loadList();
        },
    },
    methods: {
        wrapStartSearch() {
            if (this.inputValue === this.oldInputValue)
                return;
            this.reset(); // 回到第一页
            this.oldInputValue = this.inputValue;
            if (this.inputValue === '' || this.inputValue === undefined)
                this.loadList();
            else {
                if (this.startSearch)
                    this.startSearch();
                else if (this.loadList) {
                    this.loadList();
                }
            }
        },
    },
    getRender(map) {
        const {
            totalPage,
            page,
            total,
            form,
            limitList,
            needRestorePage,
        } = this;
        return (
            <u-modal-fix onClose={this.close} title={this.title} ok-button={''} cancel-button={''} visible={this.show} {...{ on: { 'update:visible': (val) => { this.show = val; } } }} size={'modallist'}>
                {map.tabs}
                {map.notice}
                {map.input}
                {vIf(
                    this.hasInput,
                    () => <div style="text-align: center;margin-bottom: 20px;"><u-icon-input style="width: 900px; height: 38px;" value={this.inputValue} placeholder={this.placeholder} {...{ on: { keypress: (event) => { event.keyCode === 13 && this.wrapStartSearch(); }, change: (val) => { this.inputValue = val.value; } } }}></u-icon-input></div>
                )}
                {vIf(
                    this.hasSearchTag,
                    () => <div style="text-align: center;margin-bottom: 20px;"><search-tag tagTypes={this.tagTypes} placeholder={this.placeholder} onSearch={this.onSearch}></search-tag></div>
                )}
                {map.table}
                <div class="modalleftfooterwrapper">{map.leftFooter}</div>
                {vIf(
                    totalPage > 1,
                    () => <div class="pager" style="margin-top: 10px;"><u-linear-layout direction="vertical">
                        <u-pagination total={totalPage} page={page} onSelect={this.changePage} />
                    </u-linear-layout></div>
                )}
            </u-modal-fix>
        );
    },
};
</script>
