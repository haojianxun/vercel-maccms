package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _MacPlogMgr struct {
	*_BaseMgr
}

// MacPlogMgr open func
func MacPlogMgr(db *gorm.DB) *_MacPlogMgr {
	if db == nil {
		panic(fmt.Errorf("MacPlogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MacPlogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("mac_plog"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_MacPlogMgr) Debug() *_MacPlogMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MacPlogMgr) GetTableName() string {
	return "mac_plog"
}

// Reset 重置gorm会话
func (obj *_MacPlogMgr) Reset() *_MacPlogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MacPlogMgr) Get() (result MacPlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacPlog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MacPlogMgr) Gets() (results []*MacPlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacPlog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MacPlogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MacPlog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithPlogID plog_id获取 积分日志ID
func (obj *_MacPlogMgr) WithPlogID(plogID int) Option {
	return optionFunc(func(o *options) { o.query["plog_id"] = plogID })
}

// WithUserID user_id获取 用户ID
func (obj *_MacPlogMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithUserID1 user_id_1获取 关联用户ID
func (obj *_MacPlogMgr) WithUserID1(userID1 int) Option {
	return optionFunc(func(o *options) { o.query["user_id_1"] = userID1 })
}

// WithPlogType plog_type获取 日志类型
func (obj *_MacPlogMgr) WithPlogType(plogType uint8) Option {
	return optionFunc(func(o *options) { o.query["plog_type"] = plogType })
}

// WithPlogPoints plog_points获取 积分数量
func (obj *_MacPlogMgr) WithPlogPoints(plogPoints uint16) Option {
	return optionFunc(func(o *options) { o.query["plog_points"] = plogPoints })
}

// WithPlogTime plog_time获取 日志时间
func (obj *_MacPlogMgr) WithPlogTime(plogTime int) Option {
	return optionFunc(func(o *options) { o.query["plog_time"] = plogTime })
}

// WithPlogRemarks plog_remarks获取 日志备注
func (obj *_MacPlogMgr) WithPlogRemarks(plogRemarks string) Option {
	return optionFunc(func(o *options) { o.query["plog_remarks"] = plogRemarks })
}

// GetByOption 功能选项模式获取
func (obj *_MacPlogMgr) GetByOption(opts ...Option) (result MacPlog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacPlog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MacPlogMgr) GetByOptions(opts ...Option) (results []*MacPlog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacPlog{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MacPlogMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MacPlog, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MacPlog{}).Where(options.query)
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

// GetFromPlogID 通过plog_id获取内容 积分日志ID
func (obj *_MacPlogMgr) GetFromPlogID(plogID int) (result MacPlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacPlog{}).Where("`plog_id` = ?", plogID).First(&result).Error

	return
}

// GetBatchFromPlogID 批量查找 积分日志ID
func (obj *_MacPlogMgr) GetBatchFromPlogID(plogIDs []int) (results []*MacPlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacPlog{}).Where("`plog_id` IN (?)", plogIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户ID
func (obj *_MacPlogMgr) GetFromUserID(userID int) (results []*MacPlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacPlog{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找 用户ID
func (obj *_MacPlogMgr) GetBatchFromUserID(userIDs []int) (results []*MacPlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacPlog{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromUserID1 通过user_id_1获取内容 关联用户ID
func (obj *_MacPlogMgr) GetFromUserID1(userID1 int) (results []*MacPlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacPlog{}).Where("`user_id_1` = ?", userID1).Find(&results).Error

	return
}

// GetBatchFromUserID1 批量查找 关联用户ID
func (obj *_MacPlogMgr) GetBatchFromUserID1(userID1s []int) (results []*MacPlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacPlog{}).Where("`user_id_1` IN (?)", userID1s).Find(&results).Error

	return
}

// GetFromPlogType 通过plog_type获取内容 日志类型
func (obj *_MacPlogMgr) GetFromPlogType(plogType uint8) (results []*MacPlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacPlog{}).Where("`plog_type` = ?", plogType).Find(&results).Error

	return
}

// GetBatchFromPlogType 批量查找 日志类型
func (obj *_MacPlogMgr) GetBatchFromPlogType(plogTypes []uint8) (results []*MacPlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacPlog{}).Where("`plog_type` IN (?)", plogTypes).Find(&results).Error

	return
}

// GetFromPlogPoints 通过plog_points获取内容 积分数量
func (obj *_MacPlogMgr) GetFromPlogPoints(plogPoints uint16) (results []*MacPlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacPlog{}).Where("`plog_points` = ?", plogPoints).Find(&results).Error

	return
}

// GetBatchFromPlogPoints 批量查找 积分数量
func (obj *_MacPlogMgr) GetBatchFromPlogPoints(plogPointss []uint16) (results []*MacPlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacPlog{}).Where("`plog_points` IN (?)", plogPointss).Find(&results).Error

	return
}

// GetFromPlogTime 通过plog_time获取内容 日志时间
func (obj *_MacPlogMgr) GetFromPlogTime(plogTime int) (results []*MacPlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacPlog{}).Where("`plog_time` = ?", plogTime).Find(&results).Error

	return
}

// GetBatchFromPlogTime 批量查找 日志时间
func (obj *_MacPlogMgr) GetBatchFromPlogTime(plogTimes []int) (results []*MacPlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacPlog{}).Where("`plog_time` IN (?)", plogTimes).Find(&results).Error

	return
}

// GetFromPlogRemarks 通过plog_remarks获取内容 日志备注
func (obj *_MacPlogMgr) GetFromPlogRemarks(plogRemarks string) (results []*MacPlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacPlog{}).Where("`plog_remarks` = ?", plogRemarks).Find(&results).Error

	return
}

// GetBatchFromPlogRemarks 批量查找 日志备注
func (obj *_MacPlogMgr) GetBatchFromPlogRemarks(plogRemarkss []string) (results []*MacPlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacPlog{}).Where("`plog_remarks` IN (?)", plogRemarkss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MacPlogMgr) FetchByPrimaryKey(plogID int) (result MacPlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacPlog{}).Where("`plog_id` = ?", plogID).First(&result).Error

	return
}

// FetchIndexByUserID  获取多个内容
func (obj *_MacPlogMgr) FetchIndexByUserID(userID int) (results []*MacPlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacPlog{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// FetchIndexByPlogType  获取多个内容
func (obj *_MacPlogMgr) FetchIndexByPlogType(plogType uint8) (results []*MacPlog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacPlog{}).Where("`plog_type` = ?", plogType).Find(&results).Error

	return
}
