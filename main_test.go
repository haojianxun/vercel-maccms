package main_test

import (
	"fmt"
	"goapi/pkg/maccms"
	"net/url"
	"strings"
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

type MovieInfo struct {
	MovieName    string
	ActorName    string
	AreaName     string
	DirectorName string
	Page         string
	Year         string
}

func extractParameters(url string) (page, movieName, actorName, areaName, directorName, year string) {
	// 使用字符串分割来提取参数
	parts := strings.Split(url, "---")
	if len(parts) < 4 {
		return "", "", "", "", "", ""
	}

	// 根据前缀判断参数类型
	switch {
	case strings.HasPrefix(parts[0], "-"):
		movieName = strings.ReplaceAll(parts[0], "-", "")
	case strings.HasPrefix(parts[0], "--"):
		actorName = strings.ReplaceAll(parts[0], "-", "")
	}
	areaName = strings.ReplaceAll(parts[1], "-", "")
	directorName = strings.ReplaceAll(parts[2], "-", "")
	page = strings.ReplaceAll(parts[3], "-", "")
	year = strings.ReplaceAll(parts[4], "-", "")
	return page, movieName, actorName, areaName, directorName, year
}

func TestReg(t *testing.T) {
	// 测试示例URL
	urls := []string{
		"-电影名称----------第几页---",
		"--演员名称---------第几页---",
		"---大陆--------第几页---",
		"------导演-----第几页---",
		"-----------第几页---年份",
	}

	// 提取参数并打印
	for _, url := range urls {
		page, name, actorName, areaName, directorName, year := extractParameters(url)
		fmt.Println("page", page, "name", name, "actorName", actorName, "areaName", areaName, "directorName", directorName, "year", year)
	}
}
