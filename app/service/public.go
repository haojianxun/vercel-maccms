package service

import (
	"goapi/app/models"
	"goapi/pkg/mysql"
)

// 热播数据

func CurrentlyTrending(table string, typeId1, limit int, CurrentlyTrending *[]models.MacVod) {
	CacheCurrentlyTrending, _ := GetTable(table, "CurrentlyTrending", []models.MacVod{})
	if CacheCurrentlyTrending == nil {
		models.MacVodMgr(mysql.DB).Debug().Where("type_id_1", typeId1).Where("vod_status", 1).Order("vod_hits desc").Limit(limit).Find(&CurrentlyTrending)
		SaveTable(table, "CurrentlyTrending", CurrentlyTrending)
	} else {
		*CurrentlyTrending = *CacheCurrentlyTrending.(*[]models.MacVod) // 更新指针指向的值
	}
}

// 二级分类

func ListMacType(table string, typePid int, listMacType *[]models.MacType) {
	CacheMacType, _ := GetTable(table, "listMacType", []models.MacType{})
	if CacheMacType == nil {
		models.MacTypeMgr(mysql.DB).Debug().Where("type_pid", typePid).Where("type_status", 1).Order("type_sort asc").Find(&listMacType)
		SaveTable(table, "listMacType", listMacType)
	} else {
		*listMacType = *CacheMacType.(*[]models.MacType) // 更新指针指向的值
	}
}
