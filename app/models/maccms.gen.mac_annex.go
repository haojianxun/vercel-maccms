package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _MacAnnexMgr struct {
	*_BaseMgr
}

// MacAnnexMgr open func
func MacAnnexMgr(db *gorm.DB) *_MacAnnexMgr {
	if db == nil {
		panic(fmt.Errorf("MacAnnexMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MacAnnexMgr{_BaseMgr: &_BaseMgr{DB: db.Table("mac_annex"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_MacAnnexMgr) Debug() *_MacAnnexMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MacAnnexMgr) GetTableName() string {
	return "mac_annex"
}

// Reset 重置gorm会话
func (obj *_MacAnnexMgr) Reset() *_MacAnnexMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MacAnnexMgr) Get() (result MacAnnex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAnnex{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MacAnnexMgr) Gets() (results []*MacAnnex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAnnex{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MacAnnexMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MacAnnex{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithAnnexID annex_id获取 附件ID
func (obj *_MacAnnexMgr) WithAnnexID(annexID int) Option {
	return optionFunc(func(o *options) { o.query["annex_id"] = annexID })
}

// WithAnnexTime annex_time获取 附件上传时间
func (obj *_MacAnnexMgr) WithAnnexTime(annexTime int) Option {
	return optionFunc(func(o *options) { o.query["annex_time"] = annexTime })
}

// WithAnnexFile annex_file获取 附件文件名
func (obj *_MacAnnexMgr) WithAnnexFile(annexFile string) Option {
	return optionFunc(func(o *options) { o.query["annex_file"] = annexFile })
}

// WithAnnexSize annex_size获取 附件大小
func (obj *_MacAnnexMgr) WithAnnexSize(annexSize int) Option {
	return optionFunc(func(o *options) { o.query["annex_size"] = annexSize })
}

// WithAnnexType annex_type获取 附件类型
func (obj *_MacAnnexMgr) WithAnnexType(annexType string) Option {
	return optionFunc(func(o *options) { o.query["annex_type"] = annexType })
}

// GetByOption 功能选项模式获取
func (obj *_MacAnnexMgr) GetByOption(opts ...Option) (result MacAnnex, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacAnnex{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MacAnnexMgr) GetByOptions(opts ...Option) (results []*MacAnnex, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacAnnex{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MacAnnexMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MacAnnex, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MacAnnex{}).Where(options.query)
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

// GetFromAnnexID 通过annex_id获取内容 附件ID
func (obj *_MacAnnexMgr) GetFromAnnexID(annexID int) (result MacAnnex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAnnex{}).Where("`annex_id` = ?", annexID).First(&result).Error

	return
}

// GetBatchFromAnnexID 批量查找 附件ID
func (obj *_MacAnnexMgr) GetBatchFromAnnexID(annexIDs []int) (results []*MacAnnex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAnnex{}).Where("`annex_id` IN (?)", annexIDs).Find(&results).Error

	return
}

// GetFromAnnexTime 通过annex_time获取内容 附件上传时间
func (obj *_MacAnnexMgr) GetFromAnnexTime(annexTime int) (results []*MacAnnex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAnnex{}).Where("`annex_time` = ?", annexTime).Find(&results).Error

	return
}

// GetBatchFromAnnexTime 批量查找 附件上传时间
func (obj *_MacAnnexMgr) GetBatchFromAnnexTime(annexTimes []int) (results []*MacAnnex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAnnex{}).Where("`annex_time` IN (?)", annexTimes).Find(&results).Error

	return
}

// GetFromAnnexFile 通过annex_file获取内容 附件文件名
func (obj *_MacAnnexMgr) GetFromAnnexFile(annexFile string) (results []*MacAnnex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAnnex{}).Where("`annex_file` = ?", annexFile).Find(&results).Error

	return
}

// GetBatchFromAnnexFile 批量查找 附件文件名
func (obj *_MacAnnexMgr) GetBatchFromAnnexFile(annexFiles []string) (results []*MacAnnex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAnnex{}).Where("`annex_file` IN (?)", annexFiles).Find(&results).Error

	return
}

// GetFromAnnexSize 通过annex_size获取内容 附件大小
func (obj *_MacAnnexMgr) GetFromAnnexSize(annexSize int) (results []*MacAnnex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAnnex{}).Where("`annex_size` = ?", annexSize).Find(&results).Error

	return
}

// GetBatchFromAnnexSize 批量查找 附件大小
func (obj *_MacAnnexMgr) GetBatchFromAnnexSize(annexSizes []int) (results []*MacAnnex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAnnex{}).Where("`annex_size` IN (?)", annexSizes).Find(&results).Error

	return
}

// GetFromAnnexType 通过annex_type获取内容 附件类型
func (obj *_MacAnnexMgr) GetFromAnnexType(annexType string) (results []*MacAnnex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAnnex{}).Where("`annex_type` = ?", annexType).Find(&results).Error

	return
}

// GetBatchFromAnnexType 批量查找 附件类型
func (obj *_MacAnnexMgr) GetBatchFromAnnexType(annexTypes []string) (results []*MacAnnex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAnnex{}).Where("`annex_type` IN (?)", annexTypes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MacAnnexMgr) FetchByPrimaryKey(annexID int) (result MacAnnex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAnnex{}).Where("`annex_id` = ?", annexID).First(&result).Error

	return
}

// FetchIndexByAnnexTime  获取多个内容
func (obj *_MacAnnexMgr) FetchIndexByAnnexTime(annexTime int) (results []*MacAnnex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAnnex{}).Where("`annex_time` = ?", annexTime).Find(&results).Error

	return
}

// FetchIndexByAnnexFile  获取多个内容
func (obj *_MacAnnexMgr) FetchIndexByAnnexFile(annexFile string) (results []*MacAnnex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAnnex{}).Where("`annex_file` = ?", annexFile).Find(&results).Error

	return
}

// FetchIndexByAnnexType  获取多个内容
func (obj *_MacAnnexMgr) FetchIndexByAnnexType(annexType string) (results []*MacAnnex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAnnex{}).Where("`annex_type` = ?", annexType).Find(&results).Error

	return
}
