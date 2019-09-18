/**
 * 存在各个模块OPENAPI接口的Version或者请求前缀等其他信息
 * 因为请求前缀基本是以gateway/v2，可默认不填写
 * @param {string} version - 模块的版本号
 * @param {boolean} isQueryFirst - 在调用请求的时候，是否采用query优先的准则，实例如下。
 *
 * @notice
 * 新增接口、接口兼容性升级，都不需要修改Version；不兼容性升级的话，需要统一升级一下所有接口的Version。
 * 但是其他接口使用原有Version也需要支持，所有接口在帮助文档上统一一个最新的Version。
 * 模块的Version保证上述的准则，如果模块后端没有保证，则需要提醒，拉网关的负责-陈重阳一起讨论。
 *
 * @notice 请求参数下的其他属性正常书写，如: body, path, header, config等。
 * isQueryFirst === false
 *      service.load({
 *          query: {
 *              limit: 20,
 *              offset: 20,
 *          },
 *          // header: ....
 *      })
 *      ....
 * isQueryFirst === true
 *      service.load({
 *          limit: 20,
 *          offset: 20,
 *          // headers: ...
 *      })
 *      ....
 */

// export default {
const map = {
    ncs: {
        version: '2017-11-16',
        isQueryFirst: true,
    },
    pubproxy: {
        version: '2018-03-22',
    },
    nes: {
        version: '2018-02-08',
        isQueryFirst: true,
    },
    dns: {
        version: '2017-12-12',
    },
    nqs: {
        version: '2017-12-01',
    },
    order: {
        version: '2017-12-28',
    },
    nvm: {
        version: '2017-12-14',
    },
    ncv: {
        version: '2017-12-28',
    },
    mirror: {
        isQueryFirst: true,
    },
    ing: {
        version: '2017-12-05',
    },
    ccr: {
        version: '2018-03-08',
        isQueryFirst: true,
    },
};

export default map;
