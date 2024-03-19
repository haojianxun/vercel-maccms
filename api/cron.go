package handler

import (
	"fmt"
	"goapi/bootstrap"
	"goapi/config"
	"goapi/pkg/logger"
	"io/ioutil"
	"net/http"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init()
	// 初始化 SQL
	logger.Info("初始化 SQL")
	bootstrap.SetupDB()
	// 初始化 Redis
	logger.Info("初始化 Redis")
	bootstrap.SetupRedis(0)
}

/**

vercel 入口

定时器执行函数
*/

func Timer(w http.ResponseWriter, r *http.Request) {
	// 启动路由
	logger.Info("开始执行定时器")
	logger.Info("百度提交")
	fetchGet("https://seo-baidu.dyxs.site/app/v1/seo/baidu")
	logger.Info("Bing提交")
	fetchGet("https://seo-baidu.dyxs.site/app/v1/seo/bing")
}

func fetchGet(url string) {
	// 发起 HTTP GET 请求
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("HTTP 请求失败:", err)
		return
	}
	defer response.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("读取响应内容失败:", err)
		return
	}

	// 打印响应内容
	fmt.Println("HTTP 请求成功:", string(body))
}
