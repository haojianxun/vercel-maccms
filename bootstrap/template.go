package bootstrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goapi/app/models"
	"goapi/pkg/maccms"
	"goapi/templates"
	"html/template"
	"io/fs"
	"net/http"
	"strings"
)

func SetupTemplate(router *gin.Engine) {
	// 将静态文件夹目录绑定到 statics 路由下访问
	defaultFads, _ := fs.Sub(templates.Statics, "statics")
	router.StaticFS("/statics", http.FS(defaultFads))
	templateNew := template.New("")
	// 注入函数
	templateNewFunc := templateNew.Funcs(SetFuncMap())
	// 设置模板
	_template, err := templateNewFunc.ParseFS(
		templates.Default,
		"default/*.html",
		"default/**/*",
	)
	if err != nil {
		panic(err)
	}
	// 注册模板函数 使用自定义模板函数
	router.SetHTMLTemplate(_template)
}

// 注册自定义模板函数

func SetFuncMap() template.FuncMap {
	// 注册自定义模板函数
	funcMap := template.FuncMap{
		"split":         split,
		"add":           add,
		"sub":           sub,
		"len":           lenFunction,
		"GetVodCount":   GetVodCount,
		"GetSplitFirst": GetSplitFirst,
		"ListMacVod":    ListMacVod,
		"EncryptID":     maccms.EncryptID,
		"DecryptID":     maccms.DecryptID,
		"TypeName":      TypeName,
	}
	return funcMap
}

func split(str, sep string) []string {
	return strings.Split(str, sep)
}

func lenFunction(v map[string]interface{}) int {
	return len(v)
}

func GetSplitFirst(str, sep string) string {
	res := strings.Split(str, sep)
	if len(res) > 0 {
		return res[0]
	}
	return ""
}

func GetVodCount(str, sep string, index int) int {
	res := strings.Split(str, sep)
	return len(strings.Split(res[index], "#"))
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func ListMacVod(PageData map[string]interface{}, name string) []models.MacVod {
	list := PageData[name].([]models.MacVod)
	return list
}

func TypeName(listMacType, TypeID interface{}) string {
	name := ""
	list := listMacType.([]models.MacType)
	for _, data := range list {
		if fmt.Sprintf("%v", TypeID) == fmt.Sprintf("%v", data.TypeID) {
			name = data.TypeName
			break
		}
	}
	return name
}
