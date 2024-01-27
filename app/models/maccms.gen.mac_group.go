package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _MacGroupMgr struct {
	*_BaseMgr
}

// MacGroupMgr open func
func MacGroupMgr(db *gorm.DB) *_MacGroupMgr {
	if db == nil {
		panic(fmt.Errorf("MacGroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MacGroupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("mac_group"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_MacGroupMgr) Debug() *_MacGroupMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MacGroupMgr) GetTableName() string {
	return "mac_group"
}

// Reset 重置gorm会话
func (obj *_MacGroupMgr) Reset() *_MacGroupMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MacGroupMgr) Get() (result MacGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGroup{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MacGroupMgr) Gets() (results []*MacGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGroup{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MacGroupMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MacGroup{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGroupID group_id获取 用户组ID
func (obj *_MacGroupMgr) WithGroupID(groupID int16) Option {
	return optionFunc(func(o *options) { o.query["group_id"] = groupID })
}

// WithGroupName group_name获取 用户组名称
func (obj *_MacGroupMgr) WithGroupName(groupName string) Option {
	return optionFunc(func(o *options) { o.query["group_name"] = groupName })
}

// WithGroupStatus group_status获取 用户组状态
func (obj *_MacGroupMgr) WithGroupStatus(groupStatus uint8) Option {
	return optionFunc(func(o *options) { o.query["group_status"] = groupStatus })
}

// WithGroupType group_type获取 用户组类型
func (obj *_MacGroupMgr) WithGroupType(groupType string) Option {
	return optionFunc(func(o *options) { o.query["group_type"] = groupType })
}

// WithGroupPopedom group_popedom获取 用户组权限
func (obj *_MacGroupMgr) WithGroupPopedom(groupPopedom string) Option {
	return optionFunc(func(o *options) { o.query["group_popedom"] = groupPopedom })
}

// WithGroupPointsDay group_points_day获取 用户组每日积分
func (obj *_MacGroupMgr) WithGroupPointsDay(groupPointsDay uint16) Option {
	return optionFunc(func(o *options) { o.query["group_points_day"] = groupPointsDay })
}

// WithGroupPointsWeek group_points_week获取 用户组每周积分
func (obj *_MacGroupMgr) WithGroupPointsWeek(groupPointsWeek int16) Option {
	return optionFunc(func(o *options) { o.query["group_points_week"] = groupPointsWeek })
}

// WithGroupPointsMonth group_points_month获取 用户组每月积分
func (obj *_MacGroupMgr) WithGroupPointsMonth(groupPointsMonth uint16) Option {
	return optionFunc(func(o *options) { o.query["group_points_month"] = groupPointsMonth })
}

// WithGroupPointsYear group_points_year获取 用户组每年积分
func (obj *_MacGroupMgr) WithGroupPointsYear(groupPointsYear uint16) Option {
	return optionFunc(func(o *options) { o.query["group_points_year"] = groupPointsYear })
}

// WithGroupPointsFree group_points_free获取 用户组免费积分
func (obj *_MacGroupMgr) WithGroupPointsFree(groupPointsFree uint8) Option {
	return optionFunc(func(o *options) { o.query["group_points_free"] = groupPointsFree })
}

// GetByOption 功能选项模式获取
func (obj *_MacGroupMgr) GetByOption(opts ...Option) (result MacGroup, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacGroup{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MacGroupMgr) GetByOptions(opts ...Option) (results []*MacGroup, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacGroup{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MacGroupMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MacGroup, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MacGroup{}).Where(options.query)
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

// GetFromGroupID 通过group_id获取内容 用户组ID
func (obj *_MacGroupMgr) GetFromGroupID(groupID int16) (result MacGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGroup{}).Where("`group_id` = ?", groupID).First(&result).Error

	return
}

// GetBatchFromGroupID 批量查找 用户组ID
func (obj *_MacGroupMgr) GetBatchFromGroupID(groupIDs []int16) (results []*MacGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGroup{}).Where("`group_id` IN (?)", groupIDs).Find(&results).Error

	return
}

// GetFromGroupName 通过group_name获取内容 用户组名称
func (obj *_MacGroupMgr) GetFromGroupName(groupName string) (results []*MacGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGroup{}).Where("`group_name` = ?", groupName).Find(&results).Error

	return
}

// GetBatchFromGroupName 批量查找 用户组名称
func (obj *_MacGroupMgr) GetBatchFromGroupName(groupNames []string) (results []*MacGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGroup{}).Where("`group_name` IN (?)", groupNames).Find(&results).Error

	return
}

// GetFromGroupStatus 通过group_status获取内容 用户组状态
func (obj *_MacGroupMgr) GetFromGroupStatus(groupStatus uint8) (results []*MacGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGroup{}).Where("`group_status` = ?", groupStatus).Find(&results).Error

	return
}

// GetBatchFromGroupStatus 批量查找 用户组状态
func (obj *_MacGroupMgr) GetBatchFromGroupStatus(groupStatuss []uint8) (results []*MacGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGroup{}).Where("`group_status` IN (?)", groupStatuss).Find(&results).Error

	return
}

// GetFromGroupType 通过group_type获取内容 用户组类型
func (obj *_MacGroupMgr) GetFromGroupType(groupType string) (results []*MacGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGroup{}).Where("`group_type` = ?", groupType).Find(&results).Error

	return
}

// GetBatchFromGroupType 批量查找 用户组类型
func (obj *_MacGroupMgr) GetBatchFromGroupType(groupTypes []string) (results []*MacGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGroup{}).Where("`group_type` IN (?)", groupTypes).Find(&results).Error

	return
}

// GetFromGroupPopedom 通过group_popedom获取内容 用户组权限
func (obj *_MacGroupMgr) GetFromGroupPopedom(groupPopedom string) (results []*MacGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGroup{}).Where("`group_popedom` = ?", groupPopedom).Find(&results).Error

	return
}

// GetBatchFromGroupPopedom 批量查找 用户组权限
func (obj *_MacGroupMgr) GetBatchFromGroupPopedom(groupPopedoms []string) (results []*MacGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGroup{}).Where("`group_popedom` IN (?)", groupPopedoms).Find(&results).Error

	return
}

// GetFromGroupPointsDay 通过group_points_day获取内容 用户组每日积分
func (obj *_MacGroupMgr) GetFromGroupPointsDay(groupPointsDay uint16) (results []*MacGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGroup{}).Where("`group_points_day` = ?", groupPointsDay).Find(&results).Error

	return
}

// GetBatchFromGroupPointsDay 批量查找 用户组每日积分
func (obj *_MacGroupMgr) GetBatchFromGroupPointsDay(groupPointsDays []uint16) (results []*MacGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGroup{}).Where("`group_points_day` IN (?)", groupPointsDays).Find(&results).Error

	return
}

// GetFromGroupPointsWeek 通过group_points_week获取内容 用户组每周积分
func (obj *_MacGroupMgr) GetFromGroupPointsWeek(groupPointsWeek int16) (results []*MacGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGroup{}).Where("`group_points_week` = ?", groupPointsWeek).Find(&results).Error

	return
}

// GetBatchFromGroupPointsWeek 批量查找 用户组每周积分
func (obj *_MacGroupMgr) GetBatchFromGroupPointsWeek(groupPointsWeeks []int16) (results []*MacGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGroup{}).Where("`group_points_week` IN (?)", groupPointsWeeks).Find(&results).Error

	return
}

// GetFromGroupPointsMonth 通过group_points_month获取内容 用户组每月积分
func (obj *_MacGroupMgr) GetFromGroupPointsMonth(groupPointsMonth uint16) (results []*MacGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGroup{}).Where("`group_points_month` = ?", groupPointsMonth).Find(&results).Error

	return
}

// GetBatchFromGroupPointsMonth 批量查找 用户组每月积分
func (obj *_MacGroupMgr) GetBatchFromGroupPointsMonth(groupPointsMonths []uint16) (results []*MacGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGroup{}).Where("`group_points_month` IN (?)", groupPointsMonths).Find(&results).Error

	return
}

// GetFromGroupPointsYear 通过group_points_year获取内容 用户组每年积分
func (obj *_MacGroupMgr) GetFromGroupPointsYear(groupPointsYear uint16) (results []*MacGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGroup{}).Where("`group_points_year` = ?", groupPointsYear).Find(&results).Error

	return
}

// GetBatchFromGroupPointsYear 批量查找 用户组每年积分
func (obj *_MacGroupMgr) GetBatchFromGroupPointsYear(groupPointsYears []uint16) (results []*MacGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGroup{}).Where("`group_points_year` IN (?)", groupPointsYears).Find(&results).Error

	return
}

// GetFromGroupPointsFree 通过group_points_free获取内容 用户组免费积分
func (obj *_MacGroupMgr) GetFromGroupPointsFree(groupPointsFree uint8) (results []*MacGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGroup{}).Where("`group_points_free` = ?", groupPointsFree).Find(&results).Error

	return
}

// GetBatchFromGroupPointsFree 批量查找 用户组免费积分
func (obj *_MacGroupMgr) GetBatchFromGroupPointsFree(groupPointsFrees []uint8) (results []*MacGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGroup{}).Where("`group_points_free` IN (?)", groupPointsFrees).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MacGroupMgr) FetchByPrimaryKey(groupID int16) (result MacGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGroup{}).Where("`group_id` = ?", groupID).First(&result).Error

	return
}

// FetchIndexByGroupStatus  获取多个内容
func (obj *_MacGroupMgr) FetchIndexByGroupStatus(groupStatus uint8) (results []*MacGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGroup{}).Where("`group_status` = ?", groupStatus).Find(&results).Error

	return
}
