package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _MacTmpvodMgr struct {
	*_BaseMgr
}

// MacTmpvodMgr open func
func MacTmpvodMgr(db *gorm.DB) *_MacTmpvodMgr {
	if db == nil {
		panic(fmt.Errorf("MacTmpvodMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MacTmpvodMgr{_BaseMgr: &_BaseMgr{DB: db.Table("mac_tmpvod"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_MacTmpvodMgr) Debug() *_MacTmpvodMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MacTmpvodMgr) GetTableName() string {
	return "mac_tmpvod"
}

// Reset 重置gorm会话
func (obj *_MacTmpvodMgr) Reset() *_MacTmpvodMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MacTmpvodMgr) Get() (result MacTmpvod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTmpvod{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MacTmpvodMgr) Gets() (results []*MacTmpvod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTmpvod{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MacTmpvodMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MacTmpvod{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID1 id1获取 临时视频ID
func (obj *_MacTmpvodMgr) WithID1(id1 int) Option {
	return optionFunc(func(o *options) { o.query["id1"] = id1 })
}

// WithName1 name1获取 视频名称
func (obj *_MacTmpvodMgr) WithName1(name1 string) Option {
	return optionFunc(func(o *options) { o.query["name1"] = name1 })
}

// GetByOption 功能选项模式获取
func (obj *_MacTmpvodMgr) GetByOption(opts ...Option) (result MacTmpvod, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacTmpvod{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MacTmpvodMgr) GetByOptions(opts ...Option) (results []*MacTmpvod, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacTmpvod{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MacTmpvodMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MacTmpvod, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MacTmpvod{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID1 通过id1获取内容 临时视频ID
func (obj *_MacTmpvodMgr) GetFromID1(id1 int) (results []*MacTmpvod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTmpvod{}).Where("`id1` = ?", id1).Find(&results).Error

	return
}

// GetBatchFromID1 批量查找 临时视频ID
func (obj *_MacTmpvodMgr) GetBatchFromID1(id1s []int) (results []*MacTmpvod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTmpvod{}).Where("`id1` IN (?)", id1s).Find(&results).Error

	return
}

// GetFromName1 通过name1获取内容 视频名称
func (obj *_MacTmpvodMgr) GetFromName1(name1 string) (results []*MacTmpvod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTmpvod{}).Where("`name1` = ?", name1).Find(&results).Error

	return
}

// GetBatchFromName1 批量查找 视频名称
func (obj *_MacTmpvodMgr) GetBatchFromName1(name1s []string) (results []*MacTmpvod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTmpvod{}).Where("`name1` IN (?)", name1s).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
