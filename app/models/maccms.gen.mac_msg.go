package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _MacMsgMgr struct {
	*_BaseMgr
}

// MacMsgMgr open func
func MacMsgMgr(db *gorm.DB) *_MacMsgMgr {
	if db == nil {
		panic(fmt.Errorf("MacMsgMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MacMsgMgr{_BaseMgr: &_BaseMgr{DB: db.Table("mac_msg"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_MacMsgMgr) Debug() *_MacMsgMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MacMsgMgr) GetTableName() string {
	return "mac_msg"
}

// Reset 重置gorm会话
func (obj *_MacMsgMgr) Reset() *_MacMsgMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MacMsgMgr) Get() (result MacMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacMsg{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MacMsgMgr) Gets() (results []*MacMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacMsg{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MacMsgMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MacMsg{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithMsgID msg_id获取 消息ID
func (obj *_MacMsgMgr) WithMsgID(msgID int) Option {
	return optionFunc(func(o *options) { o.query["msg_id"] = msgID })
}

// WithUserID user_id获取 用户ID
func (obj *_MacMsgMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithMsgType msg_type获取 消息类型
func (obj *_MacMsgMgr) WithMsgType(msgType uint8) Option {
	return optionFunc(func(o *options) { o.query["msg_type"] = msgType })
}

// WithMsgStatus msg_status获取 消息状态
func (obj *_MacMsgMgr) WithMsgStatus(msgStatus uint8) Option {
	return optionFunc(func(o *options) { o.query["msg_status"] = msgStatus })
}

// WithMsgTo msg_to获取 消息接收者
func (obj *_MacMsgMgr) WithMsgTo(msgTo string) Option {
	return optionFunc(func(o *options) { o.query["msg_to"] = msgTo })
}

// WithMsgCode msg_code获取 消息代码
func (obj *_MacMsgMgr) WithMsgCode(msgCode string) Option {
	return optionFunc(func(o *options) { o.query["msg_code"] = msgCode })
}

// WithMsgContent msg_content获取 消息内容
func (obj *_MacMsgMgr) WithMsgContent(msgContent string) Option {
	return optionFunc(func(o *options) { o.query["msg_content"] = msgContent })
}

// WithMsgTime msg_time获取 消息时间
func (obj *_MacMsgMgr) WithMsgTime(msgTime int) Option {
	return optionFunc(func(o *options) { o.query["msg_time"] = msgTime })
}

// GetByOption 功能选项模式获取
func (obj *_MacMsgMgr) GetByOption(opts ...Option) (result MacMsg, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacMsg{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MacMsgMgr) GetByOptions(opts ...Option) (results []*MacMsg, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacMsg{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MacMsgMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MacMsg, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MacMsg{}).Where(options.query)
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

// GetFromMsgID 通过msg_id获取内容 消息ID
func (obj *_MacMsgMgr) GetFromMsgID(msgID int) (result MacMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacMsg{}).Where("`msg_id` = ?", msgID).First(&result).Error

	return
}

// GetBatchFromMsgID 批量查找 消息ID
func (obj *_MacMsgMgr) GetBatchFromMsgID(msgIDs []int) (results []*MacMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacMsg{}).Where("`msg_id` IN (?)", msgIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户ID
func (obj *_MacMsgMgr) GetFromUserID(userID int) (results []*MacMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacMsg{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找 用户ID
func (obj *_MacMsgMgr) GetBatchFromUserID(userIDs []int) (results []*MacMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacMsg{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromMsgType 通过msg_type获取内容 消息类型
func (obj *_MacMsgMgr) GetFromMsgType(msgType uint8) (results []*MacMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacMsg{}).Where("`msg_type` = ?", msgType).Find(&results).Error

	return
}

// GetBatchFromMsgType 批量查找 消息类型
func (obj *_MacMsgMgr) GetBatchFromMsgType(msgTypes []uint8) (results []*MacMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacMsg{}).Where("`msg_type` IN (?)", msgTypes).Find(&results).Error

	return
}

// GetFromMsgStatus 通过msg_status获取内容 消息状态
func (obj *_MacMsgMgr) GetFromMsgStatus(msgStatus uint8) (results []*MacMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacMsg{}).Where("`msg_status` = ?", msgStatus).Find(&results).Error

	return
}

// GetBatchFromMsgStatus 批量查找 消息状态
func (obj *_MacMsgMgr) GetBatchFromMsgStatus(msgStatuss []uint8) (results []*MacMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacMsg{}).Where("`msg_status` IN (?)", msgStatuss).Find(&results).Error

	return
}

// GetFromMsgTo 通过msg_to获取内容 消息接收者
func (obj *_MacMsgMgr) GetFromMsgTo(msgTo string) (results []*MacMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacMsg{}).Where("`msg_to` = ?", msgTo).Find(&results).Error

	return
}

// GetBatchFromMsgTo 批量查找 消息接收者
func (obj *_MacMsgMgr) GetBatchFromMsgTo(msgTos []string) (results []*MacMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacMsg{}).Where("`msg_to` IN (?)", msgTos).Find(&results).Error

	return
}

// GetFromMsgCode 通过msg_code获取内容 消息代码
func (obj *_MacMsgMgr) GetFromMsgCode(msgCode string) (results []*MacMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacMsg{}).Where("`msg_code` = ?", msgCode).Find(&results).Error

	return
}

// GetBatchFromMsgCode 批量查找 消息代码
func (obj *_MacMsgMgr) GetBatchFromMsgCode(msgCodes []string) (results []*MacMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacMsg{}).Where("`msg_code` IN (?)", msgCodes).Find(&results).Error

	return
}

// GetFromMsgContent 通过msg_content获取内容 消息内容
func (obj *_MacMsgMgr) GetFromMsgContent(msgContent string) (results []*MacMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacMsg{}).Where("`msg_content` = ?", msgContent).Find(&results).Error

	return
}

// GetBatchFromMsgContent 批量查找 消息内容
func (obj *_MacMsgMgr) GetBatchFromMsgContent(msgContents []string) (results []*MacMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacMsg{}).Where("`msg_content` IN (?)", msgContents).Find(&results).Error

	return
}

// GetFromMsgTime 通过msg_time获取内容 消息时间
func (obj *_MacMsgMgr) GetFromMsgTime(msgTime int) (results []*MacMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacMsg{}).Where("`msg_time` = ?", msgTime).Find(&results).Error

	return
}

// GetBatchFromMsgTime 批量查找 消息时间
func (obj *_MacMsgMgr) GetBatchFromMsgTime(msgTimes []int) (results []*MacMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacMsg{}).Where("`msg_time` IN (?)", msgTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MacMsgMgr) FetchByPrimaryKey(msgID int) (result MacMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacMsg{}).Where("`msg_id` = ?", msgID).First(&result).Error

	return
}

// FetchIndexByUserID  获取多个内容
func (obj *_MacMsgMgr) FetchIndexByUserID(userID int) (results []*MacMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacMsg{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// FetchIndexByMsgCode  获取多个内容
func (obj *_MacMsgMgr) FetchIndexByMsgCode(msgCode string) (results []*MacMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacMsg{}).Where("`msg_code` = ?", msgCode).Find(&results).Error

	return
}

// FetchIndexByMsgTime  获取多个内容
func (obj *_MacMsgMgr) FetchIndexByMsgTime(msgTime int) (results []*MacMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacMsg{}).Where("`msg_time` = ?", msgTime).Find(&results).Error

	return
}
