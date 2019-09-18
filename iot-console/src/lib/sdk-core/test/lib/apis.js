export const apisJSON = {
    get: {
        url: {
            path: '/get',
            query: {
                id: 2,
            },
            method: 'get',
        },
        query: {
            type: 'object',
        },
        path: {

        },
        body: {

        },
        response: {
            200: {
                type: 'object',
            },
            400: {

            },
        },
    },
    set: {
        url: {
            path: '/set/{id}',
            query: {
                name: '我',
            },
            method: 'post',
        },
        sign: true,
        headers: {
            a: 1,
        },
    },
    list: {
        url: {
            path: '/list',
            method: 'post',
        },
        sign: true,
        query: {
            type: 'object',
        },
    },
};
