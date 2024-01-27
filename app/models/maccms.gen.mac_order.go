package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _MacOrderMgr struct {
	*_BaseMgr
}

// MacOrderMgr open func
func MacOrderMgr(db *gorm.DB) *_MacOrderMgr {
	if db == nil {
		panic(fmt.Errorf("MacOrderMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MacOrderMgr{_BaseMgr: &_BaseMgr{DB: db.Table("mac_order"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_MacOrderMgr) Debug() *_MacOrderMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MacOrderMgr) GetTableName() string {
	return "mac_order"
}

// Reset 重置gorm会话
func (obj *_MacOrderMgr) Reset() *_MacOrderMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MacOrderMgr) Get() (result MacOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacOrder{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MacOrderMgr) Gets() (results []*MacOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacOrder{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MacOrderMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MacOrder{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithOrderID order_id获取 订单ID
func (obj *_MacOrderMgr) WithOrderID(orderID int) Option {
	return optionFunc(func(o *options) { o.query["order_id"] = orderID })
}

// WithUserID user_id获取 用户ID
func (obj *_MacOrderMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithOrderStatus order_status获取 订单状态
func (obj *_MacOrderMgr) WithOrderStatus(orderStatus uint8) Option {
	return optionFunc(func(o *options) { o.query["order_status"] = orderStatus })
}

// WithOrderCode order_code获取 订单代码
func (obj *_MacOrderMgr) WithOrderCode(orderCode string) Option {
	return optionFunc(func(o *options) { o.query["order_code"] = orderCode })
}

// WithOrderPrice order_price获取 订单价格
func (obj *_MacOrderMgr) WithOrderPrice(orderPrice float64) Option {
	return optionFunc(func(o *options) { o.query["order_price"] = orderPrice })
}

// WithOrderTime order_time获取 订单时间
func (obj *_MacOrderMgr) WithOrderTime(orderTime int) Option {
	return optionFunc(func(o *options) { o.query["order_time"] = orderTime })
}

// WithOrderPoints order_points获取 订单积分
func (obj *_MacOrderMgr) WithOrderPoints(orderPoints int16) Option {
	return optionFunc(func(o *options) { o.query["order_points"] = orderPoints })
}

// WithOrderPayType order_pay_type获取 支付方式
func (obj *_MacOrderMgr) WithOrderPayType(orderPayType string) Option {
	return optionFunc(func(o *options) { o.query["order_pay_type"] = orderPayType })
}

// WithOrderPayTime order_pay_time获取 支付时间
func (obj *_MacOrderMgr) WithOrderPayTime(orderPayTime int) Option {
	return optionFunc(func(o *options) { o.query["order_pay_time"] = orderPayTime })
}

// WithOrderRemarks order_remarks获取 订单备注
func (obj *_MacOrderMgr) WithOrderRemarks(orderRemarks string) Option {
	return optionFunc(func(o *options) { o.query["order_remarks"] = orderRemarks })
}

// GetByOption 功能选项模式获取
func (obj *_MacOrderMgr) GetByOption(opts ...Option) (result MacOrder, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacOrder{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MacOrderMgr) GetByOptions(opts ...Option) (results []*MacOrder, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacOrder{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MacOrderMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MacOrder, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MacOrder{}).Where(options.query)
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

// GetFromOrderID 通过order_id获取内容 订单ID
func (obj *_MacOrderMgr) GetFromOrderID(orderID int) (result MacOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacOrder{}).Where("`order_id` = ?", orderID).First(&result).Error

	return
}

// GetBatchFromOrderID 批量查找 订单ID
func (obj *_MacOrderMgr) GetBatchFromOrderID(orderIDs []int) (results []*MacOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacOrder{}).Where("`order_id` IN (?)", orderIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户ID
func (obj *_MacOrderMgr) GetFromUserID(userID int) (results []*MacOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacOrder{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找 用户ID
func (obj *_MacOrderMgr) GetBatchFromUserID(userIDs []int) (results []*MacOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacOrder{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromOrderStatus 通过order_status获取内容 订单状态
func (obj *_MacOrderMgr) GetFromOrderStatus(orderStatus uint8) (results []*MacOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacOrder{}).Where("`order_status` = ?", orderStatus).Find(&results).Error

	return
}

// GetBatchFromOrderStatus 批量查找 订单状态
func (obj *_MacOrderMgr) GetBatchFromOrderStatus(orderStatuss []uint8) (results []*MacOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacOrder{}).Where("`order_status` IN (?)", orderStatuss).Find(&results).Error

	return
}

// GetFromOrderCode 通过order_code获取内容 订单代码
func (obj *_MacOrderMgr) GetFromOrderCode(orderCode string) (results []*MacOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacOrder{}).Where("`order_code` = ?", orderCode).Find(&results).Error

	return
}

// GetBatchFromOrderCode 批量查找 订单代码
func (obj *_MacOrderMgr) GetBatchFromOrderCode(orderCodes []string) (results []*MacOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacOrder{}).Where("`order_code` IN (?)", orderCodes).Find(&results).Error

	return
}

// GetFromOrderPrice 通过order_price获取内容 订单价格
func (obj *_MacOrderMgr) GetFromOrderPrice(orderPrice float64) (results []*MacOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacOrder{}).Where("`order_price` = ?", orderPrice).Find(&results).Error

	return
}

// GetBatchFromOrderPrice 批量查找 订单价格
func (obj *_MacOrderMgr) GetBatchFromOrderPrice(orderPrices []float64) (results []*MacOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacOrder{}).Where("`order_price` IN (?)", orderPrices).Find(&results).Error

	return
}

// GetFromOrderTime 通过order_time获取内容 订单时间
func (obj *_MacOrderMgr) GetFromOrderTime(orderTime int) (results []*MacOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacOrder{}).Where("`order_time` = ?", orderTime).Find(&results).Error

	return
}

// GetBatchFromOrderTime 批量查找 订单时间
func (obj *_MacOrderMgr) GetBatchFromOrderTime(orderTimes []int) (results []*MacOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacOrder{}).Where("`order_time` IN (?)", orderTimes).Find(&results).Error

	return
}

// GetFromOrderPoints 通过order_points获取内容 订单积分
func (obj *_MacOrderMgr) GetFromOrderPoints(orderPoints int16) (results []*MacOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacOrder{}).Where("`order_points` = ?", orderPoints).Find(&results).Error

	return
}

// GetBatchFromOrderPoints 批量查找 订单积分
func (obj *_MacOrderMgr) GetBatchFromOrderPoints(orderPointss []int16) (results []*MacOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacOrder{}).Where("`order_points` IN (?)", orderPointss).Find(&results).Error

	return
}

// GetFromOrderPayType 通过order_pay_type获取内容 支付方式
func (obj *_MacOrderMgr) GetFromOrderPayType(orderPayType string) (results []*MacOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacOrder{}).Where("`order_pay_type` = ?", orderPayType).Find(&results).Error

	return
}

// GetBatchFromOrderPayType 批量查找 支付方式
func (obj *_MacOrderMgr) GetBatchFromOrderPayType(orderPayTypes []string) (results []*MacOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacOrder{}).Where("`order_pay_type` IN (?)", orderPayTypes).Find(&results).Error

	return
}

// GetFromOrderPayTime 通过order_pay_time获取内容 支付时间
func (obj *_MacOrderMgr) GetFromOrderPayTime(orderPayTime int) (results []*MacOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacOrder{}).Where("`order_pay_time` = ?", orderPayTime).Find(&results).Error

	return
}

// GetBatchFromOrderPayTime 批量查找 支付时间
func (obj *_MacOrderMgr) GetBatchFromOrderPayTime(orderPayTimes []int) (results []*MacOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacOrder{}).Where("`order_pay_time` IN (?)", orderPayTimes).Find(&results).Error

	return
}

// GetFromOrderRemarks 通过order_remarks获取内容 订单备注
func (obj *_MacOrderMgr) GetFromOrderRemarks(orderRemarks string) (results []*MacOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacOrder{}).Where("`order_remarks` = ?", orderRemarks).Find(&results).Error

	return
}

// GetBatchFromOrderRemarks 批量查找 订单备注
func (obj *_MacOrderMgr) GetBatchFromOrderRemarks(orderRemarkss []string) (results []*MacOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacOrder{}).Where("`order_remarks` IN (?)", orderRemarkss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MacOrderMgr) FetchByPrimaryKey(orderID int) (result MacOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacOrder{}).Where("`order_id` = ?", orderID).First(&result).Error

	return
}

// FetchIndexByUserID  获取多个内容
func (obj *_MacOrderMgr) FetchIndexByUserID(userID int) (results []*MacOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacOrder{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// FetchIndexByOrderCode  获取多个内容
func (obj *_MacOrderMgr) FetchIndexByOrderCode(orderCode string) (results []*MacOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacOrder{}).Where("`order_code` = ?", orderCode).Find(&results).Error

	return
}

// FetchIndexByOrderTime  获取多个内容
func (obj *_MacOrderMgr) FetchIndexByOrderTime(orderTime int) (results []*MacOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacOrder{}).Where("`order_time` = ?", orderTime).Find(&results).Error

	return
}
