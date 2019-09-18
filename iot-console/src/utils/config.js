import i18n from '@/utils/i18n';

const privilegeList = [
    { label: '发布', key: 'PUB' },
    { label: '订阅', key: 'SUB' },
    { label: '发布和订阅', key: 'ALL' },
];
const privilegeMap = {};
privilegeList.forEach(({ key, label }) => {
    privilegeMap[key] = label;
});

const deviceStateMap = {
    Online: '在线',
    Offline: '离线',
};

export const config = {
    privilegeList,
    privilegeMap,
    deviceStateMap,
};
