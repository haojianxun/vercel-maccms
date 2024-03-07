package maccms

import (
	"fmt"
	"goapi/app/models"
	"html/template"
	"strings"
)

func Split(str, sep string) []string {
	return strings.Split(str, sep)
}

func IsUnknown(str string) string {
	if len(str) == 0 {
		return "未知"
	} else {
		return str
	}
}

func Raw(text string) template.HTML {
	return template.HTML(text)
}

func LenFunction(v map[string]interface{}) int {
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

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
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

func TypeEn(listMacType, TypeID interface{}) string {
	name := ""
	list := listMacType.([]models.MacType)
	for _, data := range list {
		if fmt.Sprintf("%v", TypeID) == fmt.Sprintf("%v", data.TypeID) {
			name = data.TypeEn
			break
		}
	}
	return name
}
