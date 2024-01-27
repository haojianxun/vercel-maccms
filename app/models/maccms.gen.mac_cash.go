package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _MacCashMgr struct {
	*_BaseMgr
}

// MacCashMgr open func
func MacCashMgr(db *gorm.DB) *_MacCashMgr {
	if db == nil {
		panic(fmt.Errorf("MacCashMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MacCashMgr{_BaseMgr: &_BaseMgr{DB: db.Table("mac_cash"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_MacCashMgr) Debug() *_MacCashMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MacCashMgr) GetTableName() string {
	return "mac_cash"
}

// Reset 重置gorm会话
func (obj *_MacCashMgr) Reset() *_MacCashMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MacCashMgr) Get() (result MacCash, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCash{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MacCashMgr) Gets() (results []*MacCash, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCash{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MacCashMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MacCash{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithCashID cash_id获取 提现记录ID
func (obj *_MacCashMgr) WithCashID(cashID int) Option {
	return optionFunc(func(o *options) { o.query["cash_id"] = cashID })
}

// WithUserID user_id获取 用户ID
func (obj *_MacCashMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithCashStatus cash_status获取 提现状态
func (obj *_MacCashMgr) WithCashStatus(cashStatus uint8) Option {
	return optionFunc(func(o *options) { o.query["cash_status"] = cashStatus })
}

// WithCashPoints cash_points获取 提现积分
func (obj *_MacCashMgr) WithCashPoints(cashPoints uint16) Option {
	return optionFunc(func(o *options) { o.query["cash_points"] = cashPoints })
}

// WithCashMoney cash_money获取 提现金额
func (obj *_MacCashMgr) WithCashMoney(cashMoney float64) Option {
	return optionFunc(func(o *options) { o.query["cash_money"] = cashMoney })
}

// WithCashBankName cash_bank_name获取 提现银行名称
func (obj *_MacCashMgr) WithCashBankName(cashBankName string) Option {
	return optionFunc(func(o *options) { o.query["cash_bank_name"] = cashBankName })
}

// WithCashBankNo cash_bank_no获取 提现银行账号
func (obj *_MacCashMgr) WithCashBankNo(cashBankNo string) Option {
	return optionFunc(func(o *options) { o.query["cash_bank_no"] = cashBankNo })
}

// WithCashPayeeName cash_payee_name获取 提现收款人姓名
func (obj *_MacCashMgr) WithCashPayeeName(cashPayeeName string) Option {
	return optionFunc(func(o *options) { o.query["cash_payee_name"] = cashPayeeName })
}

// WithCashTime cash_time获取 提现申请时间
func (obj *_MacCashMgr) WithCashTime(cashTime int) Option {
	return optionFunc(func(o *options) { o.query["cash_time"] = cashTime })
}

// WithCashTimeAudit cash_time_audit获取 提现审核时间
func (obj *_MacCashMgr) WithCashTimeAudit(cashTimeAudit int) Option {
	return optionFunc(func(o *options) { o.query["cash_time_audit"] = cashTimeAudit })
}

// GetByOption 功能选项模式获取
func (obj *_MacCashMgr) GetByOption(opts ...Option) (result MacCash, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacCash{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MacCashMgr) GetByOptions(opts ...Option) (results []*MacCash, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacCash{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MacCashMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MacCash, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MacCash{}).Where(options.query)
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

// GetFromCashID 通过cash_id获取内容 提现记录ID
func (obj *_MacCashMgr) GetFromCashID(cashID int) (result MacCash, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCash{}).Where("`cash_id` = ?", cashID).First(&result).Error

	return
}

// GetBatchFromCashID 批量查找 提现记录ID
func (obj *_MacCashMgr) GetBatchFromCashID(cashIDs []int) (results []*MacCash, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCash{}).Where("`cash_id` IN (?)", cashIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户ID
func (obj *_MacCashMgr) GetFromUserID(userID int) (results []*MacCash, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCash{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找 用户ID
func (obj *_MacCashMgr) GetBatchFromUserID(userIDs []int) (results []*MacCash, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCash{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromCashStatus 通过cash_status获取内容 提现状态
func (obj *_MacCashMgr) GetFromCashStatus(cashStatus uint8) (results []*MacCash, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCash{}).Where("`cash_status` = ?", cashStatus).Find(&results).Error

	return
}

// GetBatchFromCashStatus 批量查找 提现状态
func (obj *_MacCashMgr) GetBatchFromCashStatus(cashStatuss []uint8) (results []*MacCash, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCash{}).Where("`cash_status` IN (?)", cashStatuss).Find(&results).Error

	return
}

// GetFromCashPoints 通过cash_points获取内容 提现积分
func (obj *_MacCashMgr) GetFromCashPoints(cashPoints uint16) (results []*MacCash, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCash{}).Where("`cash_points` = ?", cashPoints).Find(&results).Error

	return
}

// GetBatchFromCashPoints 批量查找 提现积分
func (obj *_MacCashMgr) GetBatchFromCashPoints(cashPointss []uint16) (results []*MacCash, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCash{}).Where("`cash_points` IN (?)", cashPointss).Find(&results).Error

	return
}

// GetFromCashMoney 通过cash_money获取内容 提现金额
func (obj *_MacCashMgr) GetFromCashMoney(cashMoney float64) (results []*MacCash, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCash{}).Where("`cash_money` = ?", cashMoney).Find(&results).Error

	return
}

// GetBatchFromCashMoney 批量查找 提现金额
func (obj *_MacCashMgr) GetBatchFromCashMoney(cashMoneys []float64) (results []*MacCash, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCash{}).Where("`cash_money` IN (?)", cashMoneys).Find(&results).Error

	return
}

// GetFromCashBankName 通过cash_bank_name获取内容 提现银行名称
func (obj *_MacCashMgr) GetFromCashBankName(cashBankName string) (results []*MacCash, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCash{}).Where("`cash_bank_name` = ?", cashBankName).Find(&results).Error

	return
}

// GetBatchFromCashBankName 批量查找 提现银行名称
func (obj *_MacCashMgr) GetBatchFromCashBankName(cashBankNames []string) (results []*MacCash, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCash{}).Where("`cash_bank_name` IN (?)", cashBankNames).Find(&results).Error

	return
}

// GetFromCashBankNo 通过cash_bank_no获取内容 提现银行账号
func (obj *_MacCashMgr) GetFromCashBankNo(cashBankNo string) (results []*MacCash, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCash{}).Where("`cash_bank_no` = ?", cashBankNo).Find(&results).Error

	return
}

// GetBatchFromCashBankNo 批量查找 提现银行账号
func (obj *_MacCashMgr) GetBatchFromCashBankNo(cashBankNos []string) (results []*MacCash, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCash{}).Where("`cash_bank_no` IN (?)", cashBankNos).Find(&results).Error

	return
}

// GetFromCashPayeeName 通过cash_payee_name获取内容 提现收款人姓名
func (obj *_MacCashMgr) GetFromCashPayeeName(cashPayeeName string) (results []*MacCash, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCash{}).Where("`cash_payee_name` = ?", cashPayeeName).Find(&results).Error

	return
}

// GetBatchFromCashPayeeName 批量查找 提现收款人姓名
func (obj *_MacCashMgr) GetBatchFromCashPayeeName(cashPayeeNames []string) (results []*MacCash, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCash{}).Where("`cash_payee_name` IN (?)", cashPayeeNames).Find(&results).Error

	return
}

// GetFromCashTime 通过cash_time获取内容 提现申请时间
func (obj *_MacCashMgr) GetFromCashTime(cashTime int) (results []*MacCash, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCash{}).Where("`cash_time` = ?", cashTime).Find(&results).Error

	return
}

// GetBatchFromCashTime 批量查找 提现申请时间
func (obj *_MacCashMgr) GetBatchFromCashTime(cashTimes []int) (results []*MacCash, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCash{}).Where("`cash_time` IN (?)", cashTimes).Find(&results).Error

	return
}

// GetFromCashTimeAudit 通过cash_time_audit获取内容 提现审核时间
func (obj *_MacCashMgr) GetFromCashTimeAudit(cashTimeAudit int) (results []*MacCash, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCash{}).Where("`cash_time_audit` = ?", cashTimeAudit).Find(&results).Error

	return
}

// GetBatchFromCashTimeAudit 批量查找 提现审核时间
func (obj *_MacCashMgr) GetBatchFromCashTimeAudit(cashTimeAudits []int) (results []*MacCash, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCash{}).Where("`cash_time_audit` IN (?)", cashTimeAudits).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MacCashMgr) FetchByPrimaryKey(cashID int) (result MacCash, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCash{}).Where("`cash_id` = ?", cashID).First(&result).Error

	return
}

// FetchIndexByUserID  获取多个内容
func (obj *_MacCashMgr) FetchIndexByUserID(userID int) (results []*MacCash, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCash{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// FetchIndexByCashStatus  获取多个内容
func (obj *_MacCashMgr) FetchIndexByCashStatus(cashStatus uint8) (results []*MacCash, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCash{}).Where("`cash_status` = ?", cashStatus).Find(&results).Error

	return
}
