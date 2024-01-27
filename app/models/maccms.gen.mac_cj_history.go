package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _MacCjHistoryMgr struct {
	*_BaseMgr
}

// MacCjHistoryMgr open func
func MacCjHistoryMgr(db *gorm.DB) *_MacCjHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("MacCjHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MacCjHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("mac_cj_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_MacCjHistoryMgr) Debug() *_MacCjHistoryMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MacCjHistoryMgr) GetTableName() string {
	return "mac_cj_history"
}

// Reset 重置gorm会话
func (obj *_MacCjHistoryMgr) Reset() *_MacCjHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MacCjHistoryMgr) Get() (result MacCjHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MacCjHistoryMgr) Gets() (results []*MacCjHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjHistory{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MacCjHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MacCjHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithMd5 md5获取 MD5值
func (obj *_MacCjHistoryMgr) WithMd5(md5 string) Option {
	return optionFunc(func(o *options) { o.query["md5"] = md5 })
}

// GetByOption 功能选项模式获取
func (obj *_MacCjHistoryMgr) GetByOption(opts ...Option) (result MacCjHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacCjHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MacCjHistoryMgr) GetByOptions(opts ...Option) (results []*MacCjHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacCjHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MacCjHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MacCjHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MacCjHistory{}).Where(options.query)
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

// GetFromMd5 通过md5获取内容 MD5值
func (obj *_MacCjHistoryMgr) GetFromMd5(md5 string) (result MacCjHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjHistory{}).Where("`md5` = ?", md5).First(&result).Error

	return
}

// GetBatchFromMd5 批量查找 MD5值
func (obj *_MacCjHistoryMgr) GetBatchFromMd5(md5s []string) (results []*MacCjHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjHistory{}).Where("`md5` IN (?)", md5s).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MacCjHistoryMgr) FetchByPrimaryKey(md5 string) (result MacCjHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjHistory{}).Where("`md5` = ?", md5).First(&result).Error

	return
}

// FetchIndexByMd5  获取多个内容
func (obj *_MacCjHistoryMgr) FetchIndexByMd5(md5 string) (results []*MacCjHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjHistory{}).Where("`md5` = ?", md5).Find(&results).Error

	return
}
