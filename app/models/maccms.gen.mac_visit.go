package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _MacVisitMgr struct {
	*_BaseMgr
}

// MacVisitMgr open func
func MacVisitMgr(db *gorm.DB) *_MacVisitMgr {
	if db == nil {
		panic(fmt.Errorf("MacVisitMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MacVisitMgr{_BaseMgr: &_BaseMgr{DB: db.Table("mac_visit"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_MacVisitMgr) Debug() *_MacVisitMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MacVisitMgr) GetTableName() string {
	return "mac_visit"
}

// Reset 重置gorm会话
func (obj *_MacVisitMgr) Reset() *_MacVisitMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MacVisitMgr) Get() (result MacVisit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVisit{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MacVisitMgr) Gets() (results []*MacVisit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVisit{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MacVisitMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MacVisit{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithVisitID visit_id获取 访问记录ID
func (obj *_MacVisitMgr) WithVisitID(visitID int) Option {
	return optionFunc(func(o *options) { o.query["visit_id"] = visitID })
}

// WithUserID user_id获取 用户ID
func (obj *_MacVisitMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithVisitIP visit_ip获取 访问IP
func (obj *_MacVisitMgr) WithVisitIP(visitIP int) Option {
	return optionFunc(func(o *options) { o.query["visit_ip"] = visitIP })
}

// WithVisitLy visit_ly获取 访问来源
func (obj *_MacVisitMgr) WithVisitLy(visitLy string) Option {
	return optionFunc(func(o *options) { o.query["visit_ly"] = visitLy })
}

// WithVisitTime visit_time获取 访问时间
func (obj *_MacVisitMgr) WithVisitTime(visitTime int) Option {
	return optionFunc(func(o *options) { o.query["visit_time"] = visitTime })
}

// GetByOption 功能选项模式获取
func (obj *_MacVisitMgr) GetByOption(opts ...Option) (result MacVisit, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacVisit{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MacVisitMgr) GetByOptions(opts ...Option) (results []*MacVisit, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacVisit{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MacVisitMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MacVisit, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MacVisit{}).Where(options.query)
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

// GetFromVisitID 通过visit_id获取内容 访问记录ID
func (obj *_MacVisitMgr) GetFromVisitID(visitID int) (result MacVisit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVisit{}).Where("`visit_id` = ?", visitID).First(&result).Error

	return
}

// GetBatchFromVisitID 批量查找 访问记录ID
func (obj *_MacVisitMgr) GetBatchFromVisitID(visitIDs []int) (results []*MacVisit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVisit{}).Where("`visit_id` IN (?)", visitIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户ID
func (obj *_MacVisitMgr) GetFromUserID(userID int) (results []*MacVisit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVisit{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找 用户ID
func (obj *_MacVisitMgr) GetBatchFromUserID(userIDs []int) (results []*MacVisit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVisit{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromVisitIP 通过visit_ip获取内容 访问IP
func (obj *_MacVisitMgr) GetFromVisitIP(visitIP int) (results []*MacVisit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVisit{}).Where("`visit_ip` = ?", visitIP).Find(&results).Error

	return
}

// GetBatchFromVisitIP 批量查找 访问IP
func (obj *_MacVisitMgr) GetBatchFromVisitIP(visitIPs []int) (results []*MacVisit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVisit{}).Where("`visit_ip` IN (?)", visitIPs).Find(&results).Error

	return
}

// GetFromVisitLy 通过visit_ly获取内容 访问来源
func (obj *_MacVisitMgr) GetFromVisitLy(visitLy string) (results []*MacVisit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVisit{}).Where("`visit_ly` = ?", visitLy).Find(&results).Error

	return
}

// GetBatchFromVisitLy 批量查找 访问来源
func (obj *_MacVisitMgr) GetBatchFromVisitLy(visitLys []string) (results []*MacVisit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVisit{}).Where("`visit_ly` IN (?)", visitLys).Find(&results).Error

	return
}

// GetFromVisitTime 通过visit_time获取内容 访问时间
func (obj *_MacVisitMgr) GetFromVisitTime(visitTime int) (results []*MacVisit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVisit{}).Where("`visit_time` = ?", visitTime).Find(&results).Error

	return
}

// GetBatchFromVisitTime 批量查找 访问时间
func (obj *_MacVisitMgr) GetBatchFromVisitTime(visitTimes []int) (results []*MacVisit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVisit{}).Where("`visit_time` IN (?)", visitTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MacVisitMgr) FetchByPrimaryKey(visitID int) (result MacVisit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVisit{}).Where("`visit_id` = ?", visitID).First(&result).Error

	return
}

// FetchIndexByUserID  获取多个内容
func (obj *_MacVisitMgr) FetchIndexByUserID(userID int) (results []*MacVisit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVisit{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// FetchIndexByVisitTime  获取多个内容
func (obj *_MacVisitMgr) FetchIndexByVisitTime(visitTime int) (results []*MacVisit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVisit{}).Where("`visit_time` = ?", visitTime).Find(&results).Error

	return
}
