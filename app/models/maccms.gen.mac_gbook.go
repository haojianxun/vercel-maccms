package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _MacGbookMgr struct {
	*_BaseMgr
}

// MacGbookMgr open func
func MacGbookMgr(db *gorm.DB) *_MacGbookMgr {
	if db == nil {
		panic(fmt.Errorf("MacGbookMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MacGbookMgr{_BaseMgr: &_BaseMgr{DB: db.Table("mac_gbook"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_MacGbookMgr) Debug() *_MacGbookMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MacGbookMgr) GetTableName() string {
	return "mac_gbook"
}

// Reset 重置gorm会话
func (obj *_MacGbookMgr) Reset() *_MacGbookMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MacGbookMgr) Get() (result MacGbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGbook{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MacGbookMgr) Gets() (results []*MacGbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGbook{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MacGbookMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MacGbook{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGbookID gbook_id获取 留言ID
func (obj *_MacGbookMgr) WithGbookID(gbookID int) Option {
	return optionFunc(func(o *options) { o.query["gbook_id"] = gbookID })
}

// WithGbookRid gbook_rid获取 回复留言ID
func (obj *_MacGbookMgr) WithGbookRid(gbookRid int) Option {
	return optionFunc(func(o *options) { o.query["gbook_rid"] = gbookRid })
}

// WithUserID user_id获取 用户ID
func (obj *_MacGbookMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithGbookStatus gbook_status获取 留言状态
func (obj *_MacGbookMgr) WithGbookStatus(gbookStatus uint8) Option {
	return optionFunc(func(o *options) { o.query["gbook_status"] = gbookStatus })
}

// WithGbookName gbook_name获取 留言者姓名
func (obj *_MacGbookMgr) WithGbookName(gbookName string) Option {
	return optionFunc(func(o *options) { o.query["gbook_name"] = gbookName })
}

// WithGbookIP gbook_ip获取 留言者IP
func (obj *_MacGbookMgr) WithGbookIP(gbookIP int) Option {
	return optionFunc(func(o *options) { o.query["gbook_ip"] = gbookIP })
}

// WithGbookTime gbook_time获取 留言时间
func (obj *_MacGbookMgr) WithGbookTime(gbookTime int) Option {
	return optionFunc(func(o *options) { o.query["gbook_time"] = gbookTime })
}

// WithGbookReplyTime gbook_reply_time获取 留言回复时间
func (obj *_MacGbookMgr) WithGbookReplyTime(gbookReplyTime int) Option {
	return optionFunc(func(o *options) { o.query["gbook_reply_time"] = gbookReplyTime })
}

// WithGbookContent gbook_content获取 留言内容
func (obj *_MacGbookMgr) WithGbookContent(gbookContent string) Option {
	return optionFunc(func(o *options) { o.query["gbook_content"] = gbookContent })
}

// WithGbookReply gbook_reply获取 留言回复
func (obj *_MacGbookMgr) WithGbookReply(gbookReply string) Option {
	return optionFunc(func(o *options) { o.query["gbook_reply"] = gbookReply })
}

// GetByOption 功能选项模式获取
func (obj *_MacGbookMgr) GetByOption(opts ...Option) (result MacGbook, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacGbook{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MacGbookMgr) GetByOptions(opts ...Option) (results []*MacGbook, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacGbook{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MacGbookMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MacGbook, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MacGbook{}).Where(options.query)
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

// GetFromGbookID 通过gbook_id获取内容 留言ID
func (obj *_MacGbookMgr) GetFromGbookID(gbookID int) (result MacGbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGbook{}).Where("`gbook_id` = ?", gbookID).First(&result).Error

	return
}

// GetBatchFromGbookID 批量查找 留言ID
func (obj *_MacGbookMgr) GetBatchFromGbookID(gbookIDs []int) (results []*MacGbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGbook{}).Where("`gbook_id` IN (?)", gbookIDs).Find(&results).Error

	return
}

// GetFromGbookRid 通过gbook_rid获取内容 回复留言ID
func (obj *_MacGbookMgr) GetFromGbookRid(gbookRid int) (results []*MacGbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGbook{}).Where("`gbook_rid` = ?", gbookRid).Find(&results).Error

	return
}

// GetBatchFromGbookRid 批量查找 回复留言ID
func (obj *_MacGbookMgr) GetBatchFromGbookRid(gbookRids []int) (results []*MacGbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGbook{}).Where("`gbook_rid` IN (?)", gbookRids).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户ID
func (obj *_MacGbookMgr) GetFromUserID(userID int) (results []*MacGbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGbook{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找 用户ID
func (obj *_MacGbookMgr) GetBatchFromUserID(userIDs []int) (results []*MacGbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGbook{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromGbookStatus 通过gbook_status获取内容 留言状态
func (obj *_MacGbookMgr) GetFromGbookStatus(gbookStatus uint8) (results []*MacGbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGbook{}).Where("`gbook_status` = ?", gbookStatus).Find(&results).Error

	return
}

// GetBatchFromGbookStatus 批量查找 留言状态
func (obj *_MacGbookMgr) GetBatchFromGbookStatus(gbookStatuss []uint8) (results []*MacGbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGbook{}).Where("`gbook_status` IN (?)", gbookStatuss).Find(&results).Error

	return
}

// GetFromGbookName 通过gbook_name获取内容 留言者姓名
func (obj *_MacGbookMgr) GetFromGbookName(gbookName string) (results []*MacGbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGbook{}).Where("`gbook_name` = ?", gbookName).Find(&results).Error

	return
}

// GetBatchFromGbookName 批量查找 留言者姓名
func (obj *_MacGbookMgr) GetBatchFromGbookName(gbookNames []string) (results []*MacGbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGbook{}).Where("`gbook_name` IN (?)", gbookNames).Find(&results).Error

	return
}

// GetFromGbookIP 通过gbook_ip获取内容 留言者IP
func (obj *_MacGbookMgr) GetFromGbookIP(gbookIP int) (results []*MacGbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGbook{}).Where("`gbook_ip` = ?", gbookIP).Find(&results).Error

	return
}

// GetBatchFromGbookIP 批量查找 留言者IP
func (obj *_MacGbookMgr) GetBatchFromGbookIP(gbookIPs []int) (results []*MacGbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGbook{}).Where("`gbook_ip` IN (?)", gbookIPs).Find(&results).Error

	return
}

// GetFromGbookTime 通过gbook_time获取内容 留言时间
func (obj *_MacGbookMgr) GetFromGbookTime(gbookTime int) (results []*MacGbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGbook{}).Where("`gbook_time` = ?", gbookTime).Find(&results).Error

	return
}

// GetBatchFromGbookTime 批量查找 留言时间
func (obj *_MacGbookMgr) GetBatchFromGbookTime(gbookTimes []int) (results []*MacGbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGbook{}).Where("`gbook_time` IN (?)", gbookTimes).Find(&results).Error

	return
}

// GetFromGbookReplyTime 通过gbook_reply_time获取内容 留言回复时间
func (obj *_MacGbookMgr) GetFromGbookReplyTime(gbookReplyTime int) (results []*MacGbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGbook{}).Where("`gbook_reply_time` = ?", gbookReplyTime).Find(&results).Error

	return
}

// GetBatchFromGbookReplyTime 批量查找 留言回复时间
func (obj *_MacGbookMgr) GetBatchFromGbookReplyTime(gbookReplyTimes []int) (results []*MacGbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGbook{}).Where("`gbook_reply_time` IN (?)", gbookReplyTimes).Find(&results).Error

	return
}

// GetFromGbookContent 通过gbook_content获取内容 留言内容
func (obj *_MacGbookMgr) GetFromGbookContent(gbookContent string) (results []*MacGbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGbook{}).Where("`gbook_content` = ?", gbookContent).Find(&results).Error

	return
}

// GetBatchFromGbookContent 批量查找 留言内容
func (obj *_MacGbookMgr) GetBatchFromGbookContent(gbookContents []string) (results []*MacGbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGbook{}).Where("`gbook_content` IN (?)", gbookContents).Find(&results).Error

	return
}

// GetFromGbookReply 通过gbook_reply获取内容 留言回复
func (obj *_MacGbookMgr) GetFromGbookReply(gbookReply string) (results []*MacGbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGbook{}).Where("`gbook_reply` = ?", gbookReply).Find(&results).Error

	return
}

// GetBatchFromGbookReply 批量查找 留言回复
func (obj *_MacGbookMgr) GetBatchFromGbookReply(gbookReplys []string) (results []*MacGbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGbook{}).Where("`gbook_reply` IN (?)", gbookReplys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MacGbookMgr) FetchByPrimaryKey(gbookID int) (result MacGbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGbook{}).Where("`gbook_id` = ?", gbookID).First(&result).Error

	return
}

// FetchIndexByGbookRid  获取多个内容
func (obj *_MacGbookMgr) FetchIndexByGbookRid(gbookRid int) (results []*MacGbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGbook{}).Where("`gbook_rid` = ?", gbookRid).Find(&results).Error

	return
}

// FetchIndexByUserID  获取多个内容
func (obj *_MacGbookMgr) FetchIndexByUserID(userID int) (results []*MacGbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGbook{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// FetchIndexByGbookTime  获取多个内容
func (obj *_MacGbookMgr) FetchIndexByGbookTime(gbookTime int) (results []*MacGbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGbook{}).Where("`gbook_time` = ?", gbookTime).Find(&results).Error

	return
}

// FetchIndexByGbookReplyTime  获取多个内容
func (obj *_MacGbookMgr) FetchIndexByGbookReplyTime(gbookReplyTime int) (results []*MacGbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGbook{}).Where("`gbook_reply_time` = ?", gbookReplyTime).Find(&results).Error

	return
}

// FetchIndexByGbookReply  获取多个内容
func (obj *_MacGbookMgr) FetchIndexByGbookReply(gbookReply string) (results []*MacGbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacGbook{}).Where("`gbook_reply` = ?", gbookReply).Find(&results).Error

	return
}
