/**
 * 列表分页的功能MinXin，列表页引入这个minxin，提供分页功能数据
 * @attention
 * 1.需要外界提供`loads`方法
 * 2.提供给外界refresh，可覆盖
 *
 * @todo
 * pager内的参数是否需要再包裹一层？
 *
 * @see
 * nes/index.vue
 * nes/detail/oplogs.vue
 */

export default {
    data() {
        return {
            limit: 20,
            page: 1, // 当前页数
        };
    },
    computed: {
        offset() {
            return (this.page - 1) * this.limit;
        },
        totalPage() {
            return Math.ceil((this.total || 0) / this.limit) || 1;
        },
    },
    methods: {
        refresh() {
            this.loads();
        },
        changePage($event) {
            this.page = $event.page;
            this.loads();
        },
        getFormForOAI(options) {
            const { offset, limit } = this;
            return Object.assign({
                Offset: offset,
                Limit: limit,
            }, options);
        },
        // getTotal(totalNum = 0) {
        //     totalNum = isNaN(+totalNum) ? 0 : totalNum;
        //     this.total =
        // },
    },
};
