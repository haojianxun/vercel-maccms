package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _MacCardMgr struct {
	*_BaseMgr
}

// MacCardMgr open func
func MacCardMgr(db *gorm.DB) *_MacCardMgr {
	if db == nil {
		panic(fmt.Errorf("MacCardMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MacCardMgr{_BaseMgr: &_BaseMgr{DB: db.Table("mac_card"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_MacCardMgr) Debug() *_MacCardMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MacCardMgr) GetTableName() string {
	return "mac_card"
}

// Reset 重置gorm会话
func (obj *_MacCardMgr) Reset() *_MacCardMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MacCardMgr) Get() (result MacCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCard{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MacCardMgr) Gets() (results []*MacCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCard{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MacCardMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MacCard{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithCardID card_id获取 卡片ID
func (obj *_MacCardMgr) WithCardID(cardID int) Option {
	return optionFunc(func(o *options) { o.query["card_id"] = cardID })
}

// WithCardNo card_no获取 卡号
func (obj *_MacCardMgr) WithCardNo(cardNo string) Option {
	return optionFunc(func(o *options) { o.query["card_no"] = cardNo })
}

// WithCardPwd card_pwd获取 卡密码
func (obj *_MacCardMgr) WithCardPwd(cardPwd string) Option {
	return optionFunc(func(o *options) { o.query["card_pwd"] = cardPwd })
}

// WithCardMoney card_money获取 卡上余额
func (obj *_MacCardMgr) WithCardMoney(cardMoney uint16) Option {
	return optionFunc(func(o *options) { o.query["card_money"] = cardMoney })
}

// WithCardPoints card_points获取 卡积分
func (obj *_MacCardMgr) WithCardPoints(cardPoints uint16) Option {
	return optionFunc(func(o *options) { o.query["card_points"] = cardPoints })
}

// WithCardUseStatus card_use_status获取 卡使用状态
func (obj *_MacCardMgr) WithCardUseStatus(cardUseStatus uint8) Option {
	return optionFunc(func(o *options) { o.query["card_use_status"] = cardUseStatus })
}

// WithCardSaleStatus card_sale_status获取 卡销售状态
func (obj *_MacCardMgr) WithCardSaleStatus(cardSaleStatus uint8) Option {
	return optionFunc(func(o *options) { o.query["card_sale_status"] = cardSaleStatus })
}

// WithUserID user_id获取 用户ID
func (obj *_MacCardMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithCardAddTime card_add_time获取 卡添加时间
func (obj *_MacCardMgr) WithCardAddTime(cardAddTime int) Option {
	return optionFunc(func(o *options) { o.query["card_add_time"] = cardAddTime })
}

// WithCardUseTime card_use_time获取 卡使用时间
func (obj *_MacCardMgr) WithCardUseTime(cardUseTime int) Option {
	return optionFunc(func(o *options) { o.query["card_use_time"] = cardUseTime })
}

// GetByOption 功能选项模式获取
func (obj *_MacCardMgr) GetByOption(opts ...Option) (result MacCard, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacCard{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MacCardMgr) GetByOptions(opts ...Option) (results []*MacCard, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacCard{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MacCardMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MacCard, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MacCard{}).Where(options.query)
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

// GetFromCardID 通过card_id获取内容 卡片ID
func (obj *_MacCardMgr) GetFromCardID(cardID int) (result MacCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCard{}).Where("`card_id` = ?", cardID).First(&result).Error

	return
}

// GetBatchFromCardID 批量查找 卡片ID
func (obj *_MacCardMgr) GetBatchFromCardID(cardIDs []int) (results []*MacCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCard{}).Where("`card_id` IN (?)", cardIDs).Find(&results).Error

	return
}

// GetFromCardNo 通过card_no获取内容 卡号
func (obj *_MacCardMgr) GetFromCardNo(cardNo string) (results []*MacCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCard{}).Where("`card_no` = ?", cardNo).Find(&results).Error

	return
}

// GetBatchFromCardNo 批量查找 卡号
func (obj *_MacCardMgr) GetBatchFromCardNo(cardNos []string) (results []*MacCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCard{}).Where("`card_no` IN (?)", cardNos).Find(&results).Error

	return
}

// GetFromCardPwd 通过card_pwd获取内容 卡密码
func (obj *_MacCardMgr) GetFromCardPwd(cardPwd string) (results []*MacCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCard{}).Where("`card_pwd` = ?", cardPwd).Find(&results).Error

	return
}

// GetBatchFromCardPwd 批量查找 卡密码
func (obj *_MacCardMgr) GetBatchFromCardPwd(cardPwds []string) (results []*MacCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCard{}).Where("`card_pwd` IN (?)", cardPwds).Find(&results).Error

	return
}

// GetFromCardMoney 通过card_money获取内容 卡上余额
func (obj *_MacCardMgr) GetFromCardMoney(cardMoney uint16) (results []*MacCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCard{}).Where("`card_money` = ?", cardMoney).Find(&results).Error

	return
}

// GetBatchFromCardMoney 批量查找 卡上余额
func (obj *_MacCardMgr) GetBatchFromCardMoney(cardMoneys []uint16) (results []*MacCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCard{}).Where("`card_money` IN (?)", cardMoneys).Find(&results).Error

	return
}

// GetFromCardPoints 通过card_points获取内容 卡积分
func (obj *_MacCardMgr) GetFromCardPoints(cardPoints uint16) (results []*MacCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCard{}).Where("`card_points` = ?", cardPoints).Find(&results).Error

	return
}

// GetBatchFromCardPoints 批量查找 卡积分
func (obj *_MacCardMgr) GetBatchFromCardPoints(cardPointss []uint16) (results []*MacCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCard{}).Where("`card_points` IN (?)", cardPointss).Find(&results).Error

	return
}

// GetFromCardUseStatus 通过card_use_status获取内容 卡使用状态
func (obj *_MacCardMgr) GetFromCardUseStatus(cardUseStatus uint8) (results []*MacCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCard{}).Where("`card_use_status` = ?", cardUseStatus).Find(&results).Error

	return
}

// GetBatchFromCardUseStatus 批量查找 卡使用状态
func (obj *_MacCardMgr) GetBatchFromCardUseStatus(cardUseStatuss []uint8) (results []*MacCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCard{}).Where("`card_use_status` IN (?)", cardUseStatuss).Find(&results).Error

	return
}

// GetFromCardSaleStatus 通过card_sale_status获取内容 卡销售状态
func (obj *_MacCardMgr) GetFromCardSaleStatus(cardSaleStatus uint8) (results []*MacCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCard{}).Where("`card_sale_status` = ?", cardSaleStatus).Find(&results).Error

	return
}

// GetBatchFromCardSaleStatus 批量查找 卡销售状态
func (obj *_MacCardMgr) GetBatchFromCardSaleStatus(cardSaleStatuss []uint8) (results []*MacCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCard{}).Where("`card_sale_status` IN (?)", cardSaleStatuss).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户ID
func (obj *_MacCardMgr) GetFromUserID(userID int) (results []*MacCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCard{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找 用户ID
func (obj *_MacCardMgr) GetBatchFromUserID(userIDs []int) (results []*MacCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCard{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromCardAddTime 通过card_add_time获取内容 卡添加时间
func (obj *_MacCardMgr) GetFromCardAddTime(cardAddTime int) (results []*MacCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCard{}).Where("`card_add_time` = ?", cardAddTime).Find(&results).Error

	return
}

// GetBatchFromCardAddTime 批量查找 卡添加时间
func (obj *_MacCardMgr) GetBatchFromCardAddTime(cardAddTimes []int) (results []*MacCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCard{}).Where("`card_add_time` IN (?)", cardAddTimes).Find(&results).Error

	return
}

// GetFromCardUseTime 通过card_use_time获取内容 卡使用时间
func (obj *_MacCardMgr) GetFromCardUseTime(cardUseTime int) (results []*MacCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCard{}).Where("`card_use_time` = ?", cardUseTime).Find(&results).Error

	return
}

// GetBatchFromCardUseTime 批量查找 卡使用时间
func (obj *_MacCardMgr) GetBatchFromCardUseTime(cardUseTimes []int) (results []*MacCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCard{}).Where("`card_use_time` IN (?)", cardUseTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MacCardMgr) FetchByPrimaryKey(cardID int) (result MacCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCard{}).Where("`card_id` = ?", cardID).First(&result).Error

	return
}

// FetchIndexByCardNo  获取多个内容
func (obj *_MacCardMgr) FetchIndexByCardNo(cardNo string) (results []*MacCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCard{}).Where("`card_no` = ?", cardNo).Find(&results).Error

	return
}

// FetchIndexByCardPwd  获取多个内容
func (obj *_MacCardMgr) FetchIndexByCardPwd(cardPwd string) (results []*MacCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCard{}).Where("`card_pwd` = ?", cardPwd).Find(&results).Error

	return
}

// FetchIndexByUserID  获取多个内容
func (obj *_MacCardMgr) FetchIndexByUserID(userID int) (results []*MacCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCard{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// FetchIndexByCardAddTime  获取多个内容
func (obj *_MacCardMgr) FetchIndexByCardAddTime(cardAddTime int) (results []*MacCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCard{}).Where("`card_add_time` = ?", cardAddTime).Find(&results).Error

	return
}

// FetchIndexByCardUseTime  获取多个内容
func (obj *_MacCardMgr) FetchIndexByCardUseTime(cardUseTime int) (results []*MacCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCard{}).Where("`card_use_time` = ?", cardUseTime).Find(&results).Error

	return
}
