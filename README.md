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
  "app": {
    "app_name": "GoApi",
    "app_env": "dev",
    "app_url": "http://127.0.0.1",
    "app_port": 1234,
    "app_pprof": true
  },
  "database": {
    "connection": "mysql",
    "host": "127.0.0.1",
    "port": 3306,
    "name": "maccms",
    "username": "root",
    "password": "root",
    "prefix": "mac_"
  },
  "redis": {
    "host": "127.0.0.1",
    "password": null,
    "port": 6379,
    "db": 0
  },
  "log": {
    "mysql_debug": true,
    "mysql_error": true,
    "mysql_warn": true
  },
  "maccms": {
    "site": {
      "site_name": "电影先生 - dyxs.site 在线 播放 观看",
      "site_url": "dyxs.site",
      "site_wapurl": "dyxs.site",
      "site_keywords": "电影先生,在线播放,在线观看,百度网盘,免费视频,在线视频,预告片,dyxs.site",
      "site_description": "电影先生提供最新最快的视频分享数据-dyxs.site",
      "site_icp": "",
      "site_qq": "",
      "site_email": "mail@dyxs.site",
      "install_dir": "/",
      "site_logo": "",
      "site_waplogo": "",
      "template_dir": "DYXS2",
      "html_dir": "html",
      "mob_status": "0",
      "mob_template_dir": "DYXS2",
      "mob_html_dir": "html",
      "site_tj": "<script>\nvar _hmt = _hmt || [];\n(function() {\n  var hm = document.createElement(\"script\");\n  hm.src = \"https://hm.baidu.com/hm.js?f46a2caaf84b781dc103b13289d90383\";\n  var s = document.getElementsByTagName(\"script\")[0];\n  s.parentNode.insertBefore(hm, s);\n})();\n</script>\n<script>let u=\"https:\\/\\/api.cgyx.tv:66\",token=\"f7dd072abb9af17700b0b8a1bd8b8e1dc43e13838257436c31c30f4495aba5a79382ef3b36808ab90351f85de968be49838554c9b1e809e46631bdd330fc2189\";var cltj = document.createElement(\"script\");cltj.src = u + \"\\/tj\\/tongji.js?v=1.3\";var s = document.getElementsByTagName(\"script\")[0];s.parentNode.insertBefore(cltj,s);<\\/script>\n",
      "site_status": "1",
      "site_close_tip": "站点暂时关闭，请稍后访问",
      "ads_dir": "ads",
      "mob_ads_dir": "ads"
    }
  },
  "seo": {
    "bing": {
      "apikey": "12121212121212121212"
    },
    "baidu": {
      "apikey": null
    },
    "sogou": {
      "apikey": null
    }
  },
  "bark": {
    "url": "https://api.day.app/",
    "key": {
      "mac": "YaemvezPAqvc2ASewcP2KH",
      "iphone": "nPuDoXfXsMpVrDaqU9to88"
    },
    "logo": "https://avatars.githubusercontent.com/u/135967790?s=400&u=0efb3cc947e9f0c2165c11f65f374524cb48915d&v=4"
  },
  "wechat": {
    "botUrl": "http://localhost:8000",
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
