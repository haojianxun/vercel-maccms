package bing

import (
	"fmt"
	"goapi/pkg/config"
	"io/ioutil"
	"net/http"
	"strings"
)

// SubmitURLBatch 提交 URL 批处理请求到 Bing Webmaster API
func SubmitURLBatch(siteURL string, urlList []string) (error, string) {
	// 构建 API URL
	apiURL := "https://ssl.bing.com/webmaster/api.svc/json/SubmitUrlbatch?apikey=" + config.GetString("seo.bing.apikey")

	// 构建 JSON 请求体
	payload := fmt.Sprintf(`{
        "siteUrl": "%s",
        "urlList": %s
    }`, siteURL, toJSONString(urlList))

	// 创建 HTTP 请求
	req, err := http.NewRequest(http.MethodPost, apiURL, strings.NewReader(payload))
	if err != nil {
		return fmt.Errorf("创建请求失败: %v", err), ""
	}
	// 设置请求头
	req.Header.Set("Host", "ssl.bing.com")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	// 发送请求
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("发送请求失败: %v", err), ""
	}
	defer res.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("读取响应失败: %v", err), ""
	}
	return nil, string(body)
}

// toJSONString 将字符串切片转换为 JSON 字符串
func toJSONString(strs []string) string {
	jsonStrings := make([]string, len(strs))
	for i, str := range strs {
		jsonStrings[i] = fmt.Sprintf(`"%s"`, str)
	}
	return fmt.Sprintf("[%s]", strings.Join(jsonStrings, ","))
}
