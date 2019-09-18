import apis from './apis/index';
import { prefix } from './base.js';
import createService, { setProcess, getAll } from '@/baseservice';

const transform = {
    QueryProductList(res) {
        res.ProductInfoList.forEach((item, index) => {
            item.canDelete = item.DeviceCount <= 0;
        });
        return res;
    },
    QueryTopicClassList(res) {
        res.TopicClassInfoList.forEach((item) => {
            item.isCustom = ['CUSTOM'].includes(item.TopicType);
        });
        return res;
    },
    QueryTopicList(res) {
        res.TopicInfoList.forEach((item) => {
            item.isCustom = ['CUSTOM'].includes(item.TopicType);
        });
        return res;
    },
};
setProcess(apis, transform);
const service = createService('iothub', apis, prefix);

service.QueryProductListAll = () => getAll(service.QueryProductList.bind(service), {
    get(json) {
        return {
            list: json.ProductInfoList,
            Count: json.TotalCount,
        };
    },
    set(json, all, count) {
        json.ProductInfoList = all;
        json.TotalCount = count;
    },
});

service.QueryDeviceListAll = (params) => getAll(service.QueryDeviceList.bind(service), {
    get(json) {
        return {
            list: json.DeviceInfoList,
            Count: json.TotalCount,
        };
    },
    set(json, all, count) {
        json.DeviceInfoList = all;
        json.TotalCount = count;
    },
}, {
    query: params.query,
});

service.QueryTopicClassListAll = (params) => getAll(service.QueryTopicClassList.bind(service), {
    get(json) {
        return {
            list: json.TopicClassInfoList,
            Count: json.TotalCount,
        };
    },
    set(json, all, count) {
        json.TopicClassInfoList = all;
        json.TotalCount = count;
    },
}, {
    query: params.query,
});

export default service;
