package main_test

import (
	"fmt"
	"goapi/pkg/maccms"
	"goapi/pkg/seoTools/bing"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"testing"
)

func TestBase64(t *testing.T) {
	ASAS := "%68%74%74%70%73%3A%2F%2F%76%2E%63%64%6E%6C%7A%32%2E%63%6F%6D%2F%32%30%32%34%30%32%30%35%2F%32%38%37%30%38%5F%64%35%34%66%37%66%61%64%2F%69%6E%64%65%78%2E%6D%33%75%38"
	encoded := maccms.Base64encode(ASAS)
	fmt.Println("encoded", encoded) // Output: SGVsbG8sIOS4lueVjA==

	decoded := maccms.Base64decode(encoded)
	fmt.Println("decoded", decoded) // Output: Hello, 世界!
}

func TestUrl(t *testing.T) {
	Urls()
}

func Urls() {
	encodedURL := "%68%74%74%70%73%3A%2F%2F%76%2E%63%64%6E%6C%7A%32%2E%63%6F%6D%2F%32%30%32%34%30%32%30%35%2F%32%38%37%30%38%5F%64%35%34%66%37%66%61%64%2F%69%6E%64%65%78%2E%6D%33%75%38"

	decodedURL, err := url.QueryUnescape(encodedURL)
	if err != nil {
		fmt.Println("Error decoding URL:", err)
		return
	}

	fmt.Println("Decoded URL:", decodedURL)
	encodedURL2 := maccms.EncodeURL(decodedURL)

	fmt.Println("EncodedURL URL:", encodedURL2)
}

func TestReg(t *testing.T) {
	// 调用 submitURLBatch 方法提交 URL 批处理请求
	var urlList []string
	for i := 1; i <= 88; i++ {
		urlList = append(urlList, fmt.Sprintf("https://dianyingxs.cc/vod/detail/id/%v.html", i))
	}
	err, _ := bing.SubmitURLBatch("https://dianyingxs.cc", urlList)
	if err != nil {
		fmt.Println("提交请求失败:", err)
		return
	}
}

func TestJump(t *testing.T) {
	url := "https://www.baidu.com/link?url=noGxhGv9xu_hMEFjzNHkjYShs7RKLLeuGpwTo4lKYfN6w65OxqigdLePVUsCnTlo&wd=&eqid=abfdebd9000cee030000000665f93e1e"
	refreshURL := getRedirectURL(url)
	if refreshURL != "" {
		fmt.Println("提取到的重定向URL:", refreshURL)
	} else {
		fmt.Println("未找到重定向URL")
	}
}

func getRedirectURL(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("请求失败:", err)
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return ""
	}

	// 提取JavaScript中的重定向URL
	htmlContent := string(body)

	// 使用正则表达式提取重定向URL
	re := regexp.MustCompile(`window\.location\.replace\("([^"]+)"\)`)
	matches := re.FindStringSubmatch(htmlContent)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}
