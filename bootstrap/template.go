package bootstrap

import (
	"github.com/gin-gonic/gin"
	"goapi/pkg/maccms"
	"goapi/templates"
	"html/template"
	"io/fs"
	"net/http"
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
		"split":         maccms.Split,
		"add":           maccms.Add,
		"sub":           maccms.Sub,
		"raw":           maccms.Raw,
		"len":           maccms.LenFunction,
		"GetVodCount":   maccms.GetVodCount,
		"GetSplitFirst": maccms.GetSplitFirst,
		"ListMacVod":    maccms.ListMacVod,
		"EncryptID":     maccms.EncryptID,
		"DecryptID":     maccms.DecryptID,
		"TypeName":      maccms.TypeName,
		"TypeEn":        maccms.TypeEn,
		"IsUnknown":     maccms.IsUnknown,
	}
	return funcMap
}
