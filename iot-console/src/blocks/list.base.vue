<script>
import Page from '../utils/mixins/page';
import Tpl from '../utils/mixins/tpl';
import pagePro from '../components/common/u-pagination-pro.vue';

const getTpl = function (tpl) {
    return typeof tpl === 'function' ? tpl() : tpl;
};
const vIf = function (data, tplTrue, tplFalse) {
    return data ? getTpl(tplTrue) : getTpl(tplFalse);
};
export default {
    components: {
        'u-pagination-pro': pagePro,
    },
    mixins: [Page, Tpl],
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
            <div class="m-box">
                {map.$head}
                {map.table}
                <div class="leftfooterwrapper">{map.leftFooter}</div>
                {vIf(
                    !needRestorePage,
                    vIf(totalPage > 1,
                        () => <div class="pager"><u-linear-layout direction="vertical">
                            <u-pagination total={totalPage} page={page} onSelect={this.changePage} />
                        </u-linear-layout></div>),
                    vIf(total > limitList[0] && totalPage >= 1, () => <div class="pager"><u-linear-layout direction="vertical">
                        <u-pagination-pro limitList={limitList} limit={form.limit} total={totalPage} page={page} onChange={this.changePage} onChangeLimit={this.changeLimit}/>
                    </u-linear-layout></div>)
                )}
            </div>
        );
    },
};
</script>
