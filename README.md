# vercel-maccms

- 基于 [vercel-golang](https://github.com/iszmxw/vercel-golang) 开发的一键免费部署到 `vercel` 上运行的 `maccms` 影视资源浏览站

## 示例网址

- [https://maccms.vercel.app](https://maccms.vercel.app)


## 俺也要一个

- 点击 `Deploy` 快速克隆一个站点

[![Deploy with Vercel](https://vercel.com/button?utm_source=busiyi&utm_campaign=oss)](https://vercel.com/new/clone?utm_source=busiyi&utm_campaign=oss&repository-url=https://github.com/dyxs-site/vercel-maccms&env=CONFIG)

- 点击 `部署静态资源` 快速部署 statics 资源

[![部署静态资源](https://img.shields.io/badge/%E9%83%A8%E7%BD%B2-%E9%9D%99%E6%80%81%E8%B5%84%E6%BA%90-brightgreen)](https://vercel.com/new/clone?utm_source=busiyi&utm_campaign=oss&repository-url=https://github.com/dyxs-site/vercel-maccms/tree/main/templates/statics)


## 配置文件

- 配置文件文 `json` 格式，`vercel` 后台设置环境变量 `CONFIG` 字段，字段内容如下
- 请自行替换成自己的配置后再进行使用

```json

{
    "APP_NAME": "GoApi",
    "APP_ENV": "dev",
    "APP_URL": "http://127.0.0.1",
    "APP_PORT": 80,
    "APP_PPROF": true,
    "HTTPS": 0,
    "ADDRESS_LIMIT": true,
    "DB_CONNECTION": "mysql",
    "DB_HOST": "db4free.net",
    "DB_PORT": 3306,
    "DB_DATABASE": "vercel_golang",
    "DB_USERNAME": "vercel_golang",
    "DB_PASSWORD": "123456",
    "DB_PREFIX": "t_",
    "REDIS_HOST": "ethical-walrus-49379.kv.vercel-storage.com",
    "REDIS_PASSWORD": "9f9659as78sdf90as78qw7fsd45cfa3b918",
    "REDIS_PORT": 6379,
    "REDIS_DB": 0,
    "PW_SALT": "xw2023-06-08",
    "LOG_MYSQL_DEBUG": true,
    "LOG_MYSQL_ERROR": true,
    "LOG_MYSQL_WARN": true,
    "NOTIFY_EMAIL": "mail@54zm.com",
    "bark": {
        "url": "https://api.day.app/",
        "key": {
            "mac": "YaemvezPAqvc2ASewcP2KH",
            "iphone": "nPuDoXfXsMpVrDaqU9to88"
        },
        "logo": "https://avatars.githubusercontent.com/u/135967790?s=400&u=0efb3cc947e9f0c2165c11f65f374524cb48915d&v=4"
    },
    "wechat": {
        "botUrl": "http://106.52.198.173:8000",
        "guid": "68d4ebef-f854-3387-ab40-7832d51dab25",
        "BelongWx": "xiaomg_zs"
    }
}

```

## 现阶段待完善功能

- [ ] 1、继续套套模板，将静态的html全部套上
- [ ] 2、编写爬虫，收集影视资源渠道，欢迎大家来投稿
- [ ] 3、选择后台模板框架，搭建基本后台功能
- [ ] 4、后台接口编写提供管理功能
- [ ] 5、完善说明文档
