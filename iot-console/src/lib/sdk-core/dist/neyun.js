(function webpackUniversalModuleDefinition(root, factory) {
	if(typeof exports === 'object' && typeof module === 'object')
		module.exports = factory();
	else if(typeof define === 'function' && define.amd)
		define([], factory);
	else if(typeof exports === 'object')
		exports["NEYUN"] = factory();
	else
		root["NEYUN"] = factory();
})(typeof self !== 'undefined' ? self : this, function() {
return /******/ (function(modules) { // webpackBootstrap
/******/ 	// The module cache
/******/ 	var installedModules = {};
/******/
/******/ 	// The require function
/******/ 	function __webpack_require__(moduleId) {
/******/
/******/ 		// Check if module is in cache
/******/ 		if(installedModules[moduleId]) {
/******/ 			return installedModules[moduleId].exports;
/******/ 		}
/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = installedModules[moduleId] = {
/******/ 			i: moduleId,
/******/ 			l: false,
/******/ 			exports: {}
/******/ 		};
/******/
/******/ 		// Execute the module function
/******/ 		modules[moduleId].call(module.exports, module, module.exports, __webpack_require__);
/******/
/******/ 		// Flag the module as loaded
/******/ 		module.l = true;
/******/
/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}
/******/
/******/
/******/ 	// expose the modules object (__webpack_modules__)
/******/ 	__webpack_require__.m = modules;
/******/
/******/ 	// expose the module cache
/******/ 	__webpack_require__.c = installedModules;
/******/
/******/ 	// define getter function for harmony exports
/******/ 	__webpack_require__.d = function(exports, name, getter) {
/******/ 		if(!__webpack_require__.o(exports, name)) {
/******/ 			Object.defineProperty(exports, name, {
/******/ 				configurable: false,
/******/ 				enumerable: true,
/******/ 				get: getter
/******/ 			});
/******/ 		}
/******/ 	};
/******/
/******/ 	// getDefaultExport function for compatibility with non-harmony modules
/******/ 	__webpack_require__.n = function(module) {
/******/ 		var getter = module && module.__esModule ?
/******/ 			function getDefault() { return module['default']; } :
/******/ 			function getModuleExports() { return module; };
/******/ 		__webpack_require__.d(getter, 'a', getter);
/******/ 		return getter;
/******/ 	};
/******/
/******/ 	// Object.prototype.hasOwnProperty.call
/******/ 	__webpack_require__.o = function(object, property) { return Object.prototype.hasOwnProperty.call(object, property); };
/******/
/******/ 	// __webpack_public_path__
/******/ 	__webpack_require__.p = "";
/******/
/******/ 	// Load entry module and return exports
/******/ 	return __webpack_require__(__webpack_require__.s = 0);
/******/ })
/************************************************************************/
/******/ ([
/* 0 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";


Object.defineProperty(exports, "__esModule", {
  value: true
});

var _service = __webpack_require__(1);

Object.defineProperty(exports, 'Service', {
  enumerable: true,
  get: function get() {
    return _interopRequireDefault(_service).default;
  }
});

var _api = __webpack_require__(2);

Object.defineProperty(exports, 'createAPI', {
  enumerable: true,
  get: function get() {
    return _interopRequireDefault(_api).default;
  }
});

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

/***/ }),
/* 1 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";


Object.defineProperty(exports, "__esModule", {
    value: true
});

var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }();

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

var Service = function () {
    function Service(serverName, request, apis) {
        _classCallCheck(this, Service);

        if (false) {
            if (!serverName) throw new Error('serverName is must');

            if (!request) throw new Error('request is must');
        }
        this.serverName = serverName;
        this.request = request;
        this.$env = {};
        if (apis) {
            this.apis = apis;
            this.$config(Object.keys(apis), apis);
        }
    }

    _createClass(Service, [{
        key: '$requestWrap',
        value: function $requestWrap(data, serviceItem) {
            var promise = void 0;
            if (serviceItem.check) {
                if (serviceItem.check(data) === false) {
                    throw new Error('check data error', data);
                }
            }
            if (serviceItem.mock) {
                promise = serviceItem.mock(data);
            }
            if (promise && promise.then) {
                return promise;
            }
            return this.request(data);
        }
    }, {
        key: '$config',
        value: function $config(config, apis) {
            var _this = this;

            apis = apis || this.apis;
            var apiKeys = Array.isArray(config) ? config : Object.keys(config);
            if (false) {
                if (!apis) throw new Error('apis is required');
            }
            apiKeys.forEach(function (item) {
                var self = _this;
                if (Array.isArray(config) ? config.includes(item) : config[item] === true) {
                    _this[item] = function tmp(data) {
                        var api = apis[item];
                        if (!api) throw new Error('no such api, check api config');
                        return self.$requestWrap(api.init(this)(data), tmp);
                    };
                } else _this[item] = config[item];
            });
            return this;
        }
    }, {
        key: '$setENV',
        value: function $setENV() {
            var _this2 = this;

            var env = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : {};

            if (false) {
                ['region', 'AK', 'SK'].forEach(function (key) {
                    if (!env[key] && !_this2.$env[key]) throw new Error(key + ' is required');
                });
                if (!env.host && !this.$env.host) console.warn('host can set in here or set in data.headers');
            }
            Object.assign(this.$env, env);
            return this;
        }
    }, {
        key: '$set',
        value: function $set(key, value) {
            this.$env[key] = value;
            return this;
        }
    }, {
        key: '$get',
        value: function $get(key) {
            return this.$env[key];
        }
    }]);

    return Service;
}();

exports.default = Service;

/***/ }),
/* 2 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";


Object.defineProperty(exports, "__esModule", {
    value: true
});

var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }();

exports.default = createAPI;

var _help = __webpack_require__(3);

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

if (false) {
    var Ajv = require('ajv');
    var ajv = new Ajv();
    var keys = ['query', 'path', 'body', 'response'];
    if (SERVER) var jsf = require('json-schema-faker/lib/index.js');else var jsf = require('json-schema-faker/dist/json-schema-faker.js');
}
if (false) {
    var request = require('request');
    var urlNode = require('url');
    var sign = require('./util/sign');
}
var bodyMethods = ['PATCH', 'POST', 'PUT'];

var API = function () {
    function API(model, serverName) {
        var _this = this;

        _classCallCheck(this, API);

        var url = model.url;
        if (false) this.sign = model.sign; // 是否需要签名

        this.url = url.path;
        this.serverName = serverName;
        this.query = url.query;
        this.method = url.method.toUpperCase();
        this.headers = model.headers;
        this.config = model.config;
        this.body = model.body;
        if (false) {
            // 初始化验证
            keys.forEach(function (key) {
                if (model[key]) {
                    if (key !== 'response') _this[key + 'Schema'] = ajv.compile(model[key]);else {
                        Object.keys(model[key]).forEach(function (status) {
                            _this['req' + status] = ajv.compile(model[key][status]);
                        });
                    }
                }
            });
        }
    }

    _createClass(API, [{
        key: 'init',
        value: function init(service) {
            var _this2 = this;

            return function (data) {
                var env = service.$env;

                var method = _this2.method;
                data = data || {};
                if (false) {
                    // 验证参数
                    keys.forEach(function (key) {
                        var keyData = data[key];
                        var keySchema = _this2[key + 'Schema'];
                        if (keySchema) {
                            if (!keySchema(keyData)) {
                                var msg = '[' + method + '] ' + _this2.url + ' ' + key + ' params is error.\n' + JSON.stringify(keySchema.errors, null, '\t');
                                throw new Error(msg);
                            }
                        }
                    });
                }
                var url = _this2.url;
                var path = data.path;
                if (path) {
                    // 在不需要签名的情况下，可以写 `/a/{id}` 的路径
                    url = (0, _help.resolvePath)(url, path);
                }
                // 在调用接口方法时传递的query参数
                var querys = data.query;
                if (!querys) {
                    Object.keys(data).forEach(function (key) {
                        var value = data[key];
                        if (!['headers', 'config', 'body', 'path'].includes(key)) querys[key] = value;
                    });
                }

                var query = Object.assign({}, _this2.query, querys);
                var headers = Object.assign({}, _this2.headers, data.headers);
                var config = Object.assign({}, _this2.config, data.config);
                var body = '';
                if (bodyMethods.includes(method)) {
                    if (Array.isArray(data.body)) {
                        body = data.body;
                    } else {
                        body = Object.assign({}, _this2.body, data.body) || body;
                    }
                }

                if (false) {
                    if (_this2.sign) {
                        if (!headers.host && !env.host) throw new Error('host is required if you want to sign');

                        headers.host = headers.host || urlNode.parse(env.host).host;
                        // query and headers will change.
                        sign.sign(env.region, env.AK, env.SK, method, _this2.url, _this2.serverName, query, headers, body);
                    }
                }
                url = (0, _help.concatURL)(url, query);
                return {
                    url: url,
                    body: body,
                    headers: headers,
                    method: method,
                    config: config,
                    query: query,
                    path: path
                };
            };
        }
    }]);

    return API;
}();

function createAPI(apis, serverName) {
    var modelAPI = {};
    Object.keys(apis).forEach(function (action) {
        modelAPI[action] = new API(apis[action], serverName);
    });
    // 释放内存
    apis = {};
    return modelAPI;
}

/***/ }),
/* 3 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";


Object.defineProperty(exports, "__esModule", {
    value: true
});
exports.isString = isString;
exports.isObject = isObject;
exports.concatURL = concatURL;
exports.resolvePath = resolvePath;
function isString(str) {
    return typeof str === 'string';
}
function isObject(obj) {
    return Object.prototype.toString.call(obj).indexOf('Object') !== -1;
}
var concatParams = function concatParams(query) {
    return Object.keys(query).map(function (key) {
        return encodeURIComponent(key) + '=' + encodeURIComponent(query[key]);
    }).join('&');
};
function concatURL(url, query) {
    var queryStr = '';
    if (isString(query)) queryStr = query;else if (isObject(query)) queryStr = concatParams(query);else queryStr = (query || '').toString();

    if (queryStr) url += (url.indexOf('?') === -1 ? '?' : '&') + queryStr;

    return url;
}
function resolvePath(url, path) {
    if (url && isObject(path)) return url.replace(/\{(.*?)\}/g, function (a, b) {
        return path[b] || a;
    });

    return url;
}

/***/ })
/******/ ]);
});