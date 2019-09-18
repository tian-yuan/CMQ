export const mapComponents = (components) => {
    const result = {};
    components.forEach((component) => result[component.options ? component.options.name : component.name] = component);
    return result;
};

/**
 * 生成接口配置
 * @param {Array} actions       [
 *                                  { action: 'CreateDatabase', method: 'post' },
                                    'DeleteDatabase'
                                ]
 * @param {*} Version
 * @param {*} path
 */
export const getConfig = (actions, Version, path) => {
    const config = {};
    actions.forEach((action) => {
        const method = action.method || 'get';
        const Action = action.action || action;
        config[Action] = {
            url: {
                path,
                method,
                query: {
                    Action,
                    Version,
                },
            },
        };
    });
    return config;
};

export const parseBackDirect = (query) => {
    let backDirect = query.backDirect;
    try {
        backDirect = JSON.parse(decodeURIComponent(backDirect));
    } catch (error) {
        backDirect = {};
    }
    return backDirect;
};
