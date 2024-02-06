package service

import (
	"goapi/app/models"
	"goapi/pkg/mysql"
)

// 获取某个栏目下的视频数据

func ListMacVod(table, name string, typeId1, limit int, listMacVod *[]models.MacVod) {
	CacheListMacVod, _ := GetTable(table, name, []models.MacVod{})
	if CacheListMacVod == nil {
		models.MacVodMgr(mysql.DB).Debug().Where("type_id_1", typeId1).Where("vod_status", 1).Order("vod_hits desc").Limit(limit).Find(&listMacVod)
		SaveTable(table, name, listMacVod)
	} else {
		*listMacVod = *CacheListMacVod.(*[]models.MacVod) // 更新指针指向的值
	}
}

// 获取某个栏目下的视频月榜数据

func MonthListMacVod(table, name string, typeId1, limit int, MonthList *[]models.MacVod) {
	CacheMonthList, _ := GetTable(table, name, []models.MacVod{})
	if CacheMonthList == nil {
		models.MacVodMgr(mysql.DB).Debug().Where("type_id_1", typeId1).Order("vod_hits_month desc").Limit(limit).Find(&MonthList)
		SaveTable(table, name, MonthList)
	} else {
		*MonthList = *CacheMonthList.(*[]models.MacVod) // 更新指针指向的值
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
