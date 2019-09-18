import { path, Version } from '../base.js';
import { getConfig } from '@/utils';
const actions = [
    'CreateProduct',
    'QueryProduct',
    'QueryProductList',
    'UpdateProduct',
    'DeleteProduct',
    'QueryProductQuota',
    { action: 'RegisterDevices', method: 'post' },
    'QueryApplyState',
    'QueryDeviceListByApplyId',
    'QueryDeviceInfo',
    'QueryDeviceList',
    'DeleteDevice',
    'QueryDeviceQuota',
    'CreateTopicClass',
    'QueryTopicClass',
    'QueryTopicClassList',
    'UpdateTopicClass',
    'DeleteTopicClass',
    'QueryTopicClassQuota',
    { action: 'PublishMessage', method: 'post' },
    'QueryMessageLengthQuota',
    'CheckProductName',
    'CheckDeviceName',
    'CheckTopicName',
    'QueryTopicList',
];

const config = getConfig(actions, Version, path);

export default config;
