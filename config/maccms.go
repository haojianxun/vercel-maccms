package config

import (
	"goapi/pkg/config"
)

// redis 参数
func init() {
	config.Add("maccms", config.StrMap{
		"site": config.StrMap{
			"site_name":        config.Env("site.site_name", "电影先生 - dianyingxs.cc 在线 播放 观看"),
			"site_url":         config.Env("site.site_url", "dianyingxs.cc"),
			"site_wapurl":      config.Env("site.site_wapurl", "dianyingxs.cc"),
			"site_keywords":    config.Env("site.site_keywords", "电影先生,在线播放,在线观看,百度网盘,免费视频,在线视频,预告片,dianyingxs.cc"),
			"site_description": config.Env("site.site_description", "电影先生提供最新最快的视频分享数据-dianyingxs.cc"),
			"site_icp":         config.Env("site.site_icp", ""),
			"site_qq":          config.Env("site.site_qq", ""),
			"site_email":       config.Env("site.site_email", "mail@dianyingxs.cc"),
			"install_dir":      config.Env("site.install_dir", "/"),
			"site_logo":        config.Env("site.site_logo", ""),
			"site_waplogo":     config.Env("site.site_waplogo", ""),
			"template_dir":     config.Env("site.template_dir", "DYXS2"),
			"html_dir":         config.Env("site.html_dir", "html"),
			"mob_status":       config.Env("site.mob_status", "0"),
			"mob_template_dir": config.Env("site.mob_template_dir", "DYXS2"),
			"mob_html_dir":     config.Env("site.mob_html_dir", "html"),
			"site_tj": config.Env("site.site_tj", `<script>
                                                         var _hmt = _hmt || [];
                                                         (function() {
                                                           var hm = document.createElement("script");
                                                           hm.src = "https://hm.baidu.com/hm.js?f46a2caaf84b781dc103b13289d90383";
                                                           var s = document.getElementsByTagName("script")[0];
                                                           s.parentNode.insertBefore(hm, s);
                                                         })();
                                                         </script>
                                                         <script>let u="https:\/\/api.cgyx.tv:66",token="f7dd072abb9af17700b0b8a1bd8b8e1dc43e13838257436c31c30f4495aba5a79382ef3b36808ab90351f85de968be49838554c9b1e809e46631bdd330fc2189";var cltj = document.createElement("script");cltj.src = u + "\/tj\/tongji.js?v=1.3";var s = document.getElementsByTagName("script")[0];s.parentNode.insertBefore(cltj,s);<\/script>`),
			"site_status":    config.Env("site.site_status", "1"),
			"site_close_tip": config.Env("site.site_close_tip", "站点暂时关闭，请稍后访问"),
			"ads_dir":        config.Env("site.ads_dir", "ads"),
			"mob_ads_dir":    config.Env("site.mob_ads_dir", "ads"),
		},
	})

}
