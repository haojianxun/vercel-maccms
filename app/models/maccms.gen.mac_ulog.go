package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _MacUlogMgr struct {
	*_BaseMgr
}

// MacUlogMgr open func
func MacUlogMgr(db *gorm.DB) *_MacUlogMgr {
	if db == nil {
		panic(fmt.Errorf("MacUlogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MacUlogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("mac_ulog"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_MacUlogMgr) Debug() *_MacUlogMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MacUlogMgr) GetTableName() string {
	return "mac_ulog"
}

// Reset 重置gorm会话
func (obj *_MacUlogMgr) Reset() *_MacUlogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MacUlogMgr) Get() (result MacUlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUlog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MacUlogMgr) Gets() (results []*MacUlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUlog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MacUlogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MacUlog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithUlogID ulog_id获取 播放记录ID
func (obj *_MacUlogMgr) WithUlogID(ulogID int) Option {
	return optionFunc(func(o *options) { o.query["ulog_id"] = ulogID })
}

// WithUserID user_id获取 用户ID
func (obj *_MacUlogMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithUlogMid ulog_mid获取 模型ID
func (obj *_MacUlogMgr) WithUlogMid(ulogMid uint8) Option {
	return optionFunc(func(o *options) { o.query["ulog_mid"] = ulogMid })
}

// WithUlogType ulog_type获取 记录类型
func (obj *_MacUlogMgr) WithUlogType(ulogType uint8) Option {
	return optionFunc(func(o *options) { o.query["ulog_type"] = ulogType })
}

// WithUlogRid ulog_rid获取 资源ID
func (obj *_MacUlogMgr) WithUlogRid(ulogRid int) Option {
	return optionFunc(func(o *options) { o.query["ulog_rid"] = ulogRid })
}

// WithUlogSid ulog_sid获取 影片集数ID
func (obj *_MacUlogMgr) WithUlogSid(ulogSid uint8) Option {
	return optionFunc(func(o *options) { o.query["ulog_sid"] = ulogSid })
}

// WithUlogNid ulog_nid获取 影片ID
func (obj *_MacUlogMgr) WithUlogNid(ulogNid uint16) Option {
	return optionFunc(func(o *options) { o.query["ulog_nid"] = ulogNid })
}

// WithUlogPoints ulog_points获取 观看进度
func (obj *_MacUlogMgr) WithUlogPoints(ulogPoints uint16) Option {
	return optionFunc(func(o *options) { o.query["ulog_points"] = ulogPoints })
}

// WithUlogTime ulog_time获取 记录时间
func (obj *_MacUlogMgr) WithUlogTime(ulogTime int) Option {
	return optionFunc(func(o *options) { o.query["ulog_time"] = ulogTime })
}

// GetByOption 功能选项模式获取
func (obj *_MacUlogMgr) GetByOption(opts ...Option) (result MacUlog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacUlog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MacUlogMgr) GetByOptions(opts ...Option) (results []*MacUlog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacUlog{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MacUlogMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MacUlog, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MacUlog{}).Where(options.query)
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

// GetFromUlogID 通过ulog_id获取内容 播放记录ID
func (obj *_MacUlogMgr) GetFromUlogID(ulogID int) (result MacUlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUlog{}).Where("`ulog_id` = ?", ulogID).First(&result).Error

	return
}

// GetBatchFromUlogID 批量查找 播放记录ID
func (obj *_MacUlogMgr) GetBatchFromUlogID(ulogIDs []int) (results []*MacUlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUlog{}).Where("`ulog_id` IN (?)", ulogIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户ID
func (obj *_MacUlogMgr) GetFromUserID(userID int) (results []*MacUlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUlog{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找 用户ID
func (obj *_MacUlogMgr) GetBatchFromUserID(userIDs []int) (results []*MacUlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUlog{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromUlogMid 通过ulog_mid获取内容 模型ID
func (obj *_MacUlogMgr) GetFromUlogMid(ulogMid uint8) (results []*MacUlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUlog{}).Where("`ulog_mid` = ?", ulogMid).Find(&results).Error

	return
}

// GetBatchFromUlogMid 批量查找 模型ID
func (obj *_MacUlogMgr) GetBatchFromUlogMid(ulogMids []uint8) (results []*MacUlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUlog{}).Where("`ulog_mid` IN (?)", ulogMids).Find(&results).Error

	return
}

// GetFromUlogType 通过ulog_type获取内容 记录类型
func (obj *_MacUlogMgr) GetFromUlogType(ulogType uint8) (results []*MacUlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUlog{}).Where("`ulog_type` = ?", ulogType).Find(&results).Error

	return
}

// GetBatchFromUlogType 批量查找 记录类型
func (obj *_MacUlogMgr) GetBatchFromUlogType(ulogTypes []uint8) (results []*MacUlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUlog{}).Where("`ulog_type` IN (?)", ulogTypes).Find(&results).Error

	return
}

// GetFromUlogRid 通过ulog_rid获取内容 资源ID
func (obj *_MacUlogMgr) GetFromUlogRid(ulogRid int) (results []*MacUlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUlog{}).Where("`ulog_rid` = ?", ulogRid).Find(&results).Error

	return
}

// GetBatchFromUlogRid 批量查找 资源ID
func (obj *_MacUlogMgr) GetBatchFromUlogRid(ulogRids []int) (results []*MacUlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUlog{}).Where("`ulog_rid` IN (?)", ulogRids).Find(&results).Error

	return
}

// GetFromUlogSid 通过ulog_sid获取内容 影片集数ID
func (obj *_MacUlogMgr) GetFromUlogSid(ulogSid uint8) (results []*MacUlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUlog{}).Where("`ulog_sid` = ?", ulogSid).Find(&results).Error

	return
}

// GetBatchFromUlogSid 批量查找 影片集数ID
func (obj *_MacUlogMgr) GetBatchFromUlogSid(ulogSids []uint8) (results []*MacUlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUlog{}).Where("`ulog_sid` IN (?)", ulogSids).Find(&results).Error

	return
}

// GetFromUlogNid 通过ulog_nid获取内容 影片ID
func (obj *_MacUlogMgr) GetFromUlogNid(ulogNid uint16) (results []*MacUlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUlog{}).Where("`ulog_nid` = ?", ulogNid).Find(&results).Error

	return
}

// GetBatchFromUlogNid 批量查找 影片ID
func (obj *_MacUlogMgr) GetBatchFromUlogNid(ulogNids []uint16) (results []*MacUlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUlog{}).Where("`ulog_nid` IN (?)", ulogNids).Find(&results).Error

	return
}

// GetFromUlogPoints 通过ulog_points获取内容 观看进度
func (obj *_MacUlogMgr) GetFromUlogPoints(ulogPoints uint16) (results []*MacUlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUlog{}).Where("`ulog_points` = ?", ulogPoints).Find(&results).Error

	return
}

// GetBatchFromUlogPoints 批量查找 观看进度
func (obj *_MacUlogMgr) GetBatchFromUlogPoints(ulogPointss []uint16) (results []*MacUlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUlog{}).Where("`ulog_points` IN (?)", ulogPointss).Find(&results).Error

	return
}

// GetFromUlogTime 通过ulog_time获取内容 记录时间
func (obj *_MacUlogMgr) GetFromUlogTime(ulogTime int) (results []*MacUlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUlog{}).Where("`ulog_time` = ?", ulogTime).Find(&results).Error

	return
}

// GetBatchFromUlogTime 批量查找 记录时间
func (obj *_MacUlogMgr) GetBatchFromUlogTime(ulogTimes []int) (results []*MacUlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUlog{}).Where("`ulog_time` IN (?)", ulogTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MacUlogMgr) FetchByPrimaryKey(ulogID int) (result MacUlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUlog{}).Where("`ulog_id` = ?", ulogID).First(&result).Error

	return
}

// FetchIndexByUserID  获取多个内容
func (obj *_MacUlogMgr) FetchIndexByUserID(userID int) (results []*MacUlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUlog{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// FetchIndexByUlogMid  获取多个内容
func (obj *_MacUlogMgr) FetchIndexByUlogMid(ulogMid uint8) (results []*MacUlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUlog{}).Where("`ulog_mid` = ?", ulogMid).Find(&results).Error

	return
}

// FetchIndexByUlogType  获取多个内容
func (obj *_MacUlogMgr) FetchIndexByUlogType(ulogType uint8) (results []*MacUlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUlog{}).Where("`ulog_type` = ?", ulogType).Find(&results).Error

	return
}

// FetchIndexByUlogRid  获取多个内容
func (obj *_MacUlogMgr) FetchIndexByUlogRid(ulogRid int) (results []*MacUlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUlog{}).Where("`ulog_rid` = ?", ulogRid).Find(&results).Error

	return
}
