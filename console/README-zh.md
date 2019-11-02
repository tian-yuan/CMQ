## Build Setup

```bash
# Clone project
first clone project from git

# Install dependencies
npm install

# 建议不要用cnpm  安装有各种诡异的bug 可以通过如下操作解决npm速度慢的问题
npm install --registry=https://registry.npm.taobao.org

# Serve with hot reload at localhost:9528
npm run dev

# Build for production with minification
npm run build

> npm run build, 会将生成文件放到与 app 同级目录 dis 下面

# Build for production and view the bundle analyzer report
npm run build --report
```

## run

after build for production, run the http server

```
node app.js
```

