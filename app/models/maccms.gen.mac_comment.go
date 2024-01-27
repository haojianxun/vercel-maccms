package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _MacCommentMgr struct {
	*_BaseMgr
}

// MacCommentMgr open func
func MacCommentMgr(db *gorm.DB) *_MacCommentMgr {
	if db == nil {
		panic(fmt.Errorf("MacCommentMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MacCommentMgr{_BaseMgr: &_BaseMgr{DB: db.Table("mac_comment"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_MacCommentMgr) Debug() *_MacCommentMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MacCommentMgr) GetTableName() string {
	return "mac_comment"
}

// Reset 重置gorm会话
func (obj *_MacCommentMgr) Reset() *_MacCommentMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MacCommentMgr) Get() (result MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MacCommentMgr) Gets() (results []*MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MacCommentMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MacComment{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithCommentID comment_id获取 评论ID
func (obj *_MacCommentMgr) WithCommentID(commentID int) Option {
	return optionFunc(func(o *options) { o.query["comment_id"] = commentID })
}

// WithCommentMid comment_mid获取 模型ID
func (obj *_MacCommentMgr) WithCommentMid(commentMid uint8) Option {
	return optionFunc(func(o *options) { o.query["comment_mid"] = commentMid })
}

// WithCommentRid comment_rid获取 回复评论ID
func (obj *_MacCommentMgr) WithCommentRid(commentRid int) Option {
	return optionFunc(func(o *options) { o.query["comment_rid"] = commentRid })
}

// WithCommentPid comment_pid获取 父评论ID
func (obj *_MacCommentMgr) WithCommentPid(commentPid int) Option {
	return optionFunc(func(o *options) { o.query["comment_pid"] = commentPid })
}

// WithUserID user_id获取 用户ID
func (obj *_MacCommentMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithCommentStatus comment_status获取 评论状态
func (obj *_MacCommentMgr) WithCommentStatus(commentStatus uint8) Option {
	return optionFunc(func(o *options) { o.query["comment_status"] = commentStatus })
}

// WithCommentName comment_name获取 评论者姓名
func (obj *_MacCommentMgr) WithCommentName(commentName string) Option {
	return optionFunc(func(o *options) { o.query["comment_name"] = commentName })
}

// WithCommentIP comment_ip获取 评论者IP
func (obj *_MacCommentMgr) WithCommentIP(commentIP int) Option {
	return optionFunc(func(o *options) { o.query["comment_ip"] = commentIP })
}

// WithCommentTime comment_time获取 评论时间
func (obj *_MacCommentMgr) WithCommentTime(commentTime int) Option {
	return optionFunc(func(o *options) { o.query["comment_time"] = commentTime })
}

// WithCommentContent comment_content获取 评论内容
func (obj *_MacCommentMgr) WithCommentContent(commentContent string) Option {
	return optionFunc(func(o *options) { o.query["comment_content"] = commentContent })
}

// WithCommentUp comment_up获取 评论点赞数
func (obj *_MacCommentMgr) WithCommentUp(commentUp int16) Option {
	return optionFunc(func(o *options) { o.query["comment_up"] = commentUp })
}

// WithCommentDown comment_down获取 评论踩数
func (obj *_MacCommentMgr) WithCommentDown(commentDown int16) Option {
	return optionFunc(func(o *options) { o.query["comment_down"] = commentDown })
}

// WithCommentReply comment_reply获取 评论回复数
func (obj *_MacCommentMgr) WithCommentReply(commentReply int16) Option {
	return optionFunc(func(o *options) { o.query["comment_reply"] = commentReply })
}

// WithCommentReport comment_report获取 评论举报数
func (obj *_MacCommentMgr) WithCommentReport(commentReport int16) Option {
	return optionFunc(func(o *options) { o.query["comment_report"] = commentReport })
}

// GetByOption 功能选项模式获取
func (obj *_MacCommentMgr) GetByOption(opts ...Option) (result MacComment, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MacCommentMgr) GetByOptions(opts ...Option) (results []*MacComment, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MacCommentMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MacComment, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where(options.query)
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

// GetFromCommentID 通过comment_id获取内容 评论ID
func (obj *_MacCommentMgr) GetFromCommentID(commentID int) (result MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where("`comment_id` = ?", commentID).First(&result).Error

	return
}

// GetBatchFromCommentID 批量查找 评论ID
func (obj *_MacCommentMgr) GetBatchFromCommentID(commentIDs []int) (results []*MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where("`comment_id` IN (?)", commentIDs).Find(&results).Error

	return
}

// GetFromCommentMid 通过comment_mid获取内容 模型ID
func (obj *_MacCommentMgr) GetFromCommentMid(commentMid uint8) (results []*MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where("`comment_mid` = ?", commentMid).Find(&results).Error

	return
}

// GetBatchFromCommentMid 批量查找 模型ID
func (obj *_MacCommentMgr) GetBatchFromCommentMid(commentMids []uint8) (results []*MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where("`comment_mid` IN (?)", commentMids).Find(&results).Error

	return
}

// GetFromCommentRid 通过comment_rid获取内容 回复评论ID
func (obj *_MacCommentMgr) GetFromCommentRid(commentRid int) (results []*MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where("`comment_rid` = ?", commentRid).Find(&results).Error

	return
}

// GetBatchFromCommentRid 批量查找 回复评论ID
func (obj *_MacCommentMgr) GetBatchFromCommentRid(commentRids []int) (results []*MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where("`comment_rid` IN (?)", commentRids).Find(&results).Error

	return
}

// GetFromCommentPid 通过comment_pid获取内容 父评论ID
func (obj *_MacCommentMgr) GetFromCommentPid(commentPid int) (results []*MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where("`comment_pid` = ?", commentPid).Find(&results).Error

	return
}

// GetBatchFromCommentPid 批量查找 父评论ID
func (obj *_MacCommentMgr) GetBatchFromCommentPid(commentPids []int) (results []*MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where("`comment_pid` IN (?)", commentPids).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户ID
func (obj *_MacCommentMgr) GetFromUserID(userID int) (results []*MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找 用户ID
func (obj *_MacCommentMgr) GetBatchFromUserID(userIDs []int) (results []*MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromCommentStatus 通过comment_status获取内容 评论状态
func (obj *_MacCommentMgr) GetFromCommentStatus(commentStatus uint8) (results []*MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where("`comment_status` = ?", commentStatus).Find(&results).Error

	return
}

// GetBatchFromCommentStatus 批量查找 评论状态
func (obj *_MacCommentMgr) GetBatchFromCommentStatus(commentStatuss []uint8) (results []*MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where("`comment_status` IN (?)", commentStatuss).Find(&results).Error

	return
}

// GetFromCommentName 通过comment_name获取内容 评论者姓名
func (obj *_MacCommentMgr) GetFromCommentName(commentName string) (results []*MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where("`comment_name` = ?", commentName).Find(&results).Error

	return
}

// GetBatchFromCommentName 批量查找 评论者姓名
func (obj *_MacCommentMgr) GetBatchFromCommentName(commentNames []string) (results []*MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where("`comment_name` IN (?)", commentNames).Find(&results).Error

	return
}

// GetFromCommentIP 通过comment_ip获取内容 评论者IP
func (obj *_MacCommentMgr) GetFromCommentIP(commentIP int) (results []*MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where("`comment_ip` = ?", commentIP).Find(&results).Error

	return
}

// GetBatchFromCommentIP 批量查找 评论者IP
func (obj *_MacCommentMgr) GetBatchFromCommentIP(commentIPs []int) (results []*MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where("`comment_ip` IN (?)", commentIPs).Find(&results).Error

	return
}

// GetFromCommentTime 通过comment_time获取内容 评论时间
func (obj *_MacCommentMgr) GetFromCommentTime(commentTime int) (results []*MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where("`comment_time` = ?", commentTime).Find(&results).Error

	return
}

// GetBatchFromCommentTime 批量查找 评论时间
func (obj *_MacCommentMgr) GetBatchFromCommentTime(commentTimes []int) (results []*MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where("`comment_time` IN (?)", commentTimes).Find(&results).Error

	return
}

// GetFromCommentContent 通过comment_content获取内容 评论内容
func (obj *_MacCommentMgr) GetFromCommentContent(commentContent string) (results []*MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where("`comment_content` = ?", commentContent).Find(&results).Error

	return
}

// GetBatchFromCommentContent 批量查找 评论内容
func (obj *_MacCommentMgr) GetBatchFromCommentContent(commentContents []string) (results []*MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where("`comment_content` IN (?)", commentContents).Find(&results).Error

	return
}

// GetFromCommentUp 通过comment_up获取内容 评论点赞数
func (obj *_MacCommentMgr) GetFromCommentUp(commentUp int16) (results []*MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where("`comment_up` = ?", commentUp).Find(&results).Error

	return
}

// GetBatchFromCommentUp 批量查找 评论点赞数
func (obj *_MacCommentMgr) GetBatchFromCommentUp(commentUps []int16) (results []*MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where("`comment_up` IN (?)", commentUps).Find(&results).Error

	return
}

// GetFromCommentDown 通过comment_down获取内容 评论踩数
func (obj *_MacCommentMgr) GetFromCommentDown(commentDown int16) (results []*MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where("`comment_down` = ?", commentDown).Find(&results).Error

	return
}

// GetBatchFromCommentDown 批量查找 评论踩数
func (obj *_MacCommentMgr) GetBatchFromCommentDown(commentDowns []int16) (results []*MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where("`comment_down` IN (?)", commentDowns).Find(&results).Error

	return
}

// GetFromCommentReply 通过comment_reply获取内容 评论回复数
func (obj *_MacCommentMgr) GetFromCommentReply(commentReply int16) (results []*MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where("`comment_reply` = ?", commentReply).Find(&results).Error

	return
}

// GetBatchFromCommentReply 批量查找 评论回复数
func (obj *_MacCommentMgr) GetBatchFromCommentReply(commentReplys []int16) (results []*MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where("`comment_reply` IN (?)", commentReplys).Find(&results).Error

	return
}

// GetFromCommentReport 通过comment_report获取内容 评论举报数
func (obj *_MacCommentMgr) GetFromCommentReport(commentReport int16) (results []*MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where("`comment_report` = ?", commentReport).Find(&results).Error

	return
}

// GetBatchFromCommentReport 批量查找 评论举报数
func (obj *_MacCommentMgr) GetBatchFromCommentReport(commentReports []int16) (results []*MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where("`comment_report` IN (?)", commentReports).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MacCommentMgr) FetchByPrimaryKey(commentID int) (result MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where("`comment_id` = ?", commentID).First(&result).Error

	return
}

// FetchIndexByCommentMid  获取多个内容
func (obj *_MacCommentMgr) FetchIndexByCommentMid(commentMid uint8) (results []*MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where("`comment_mid` = ?", commentMid).Find(&results).Error

	return
}

// FetchIndexByCommentRid  获取多个内容
func (obj *_MacCommentMgr) FetchIndexByCommentRid(commentRid int) (results []*MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where("`comment_rid` = ?", commentRid).Find(&results).Error

	return
}

// FetchIndexByCommentPid  获取多个内容
func (obj *_MacCommentMgr) FetchIndexByCommentPid(commentPid int) (results []*MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where("`comment_pid` = ?", commentPid).Find(&results).Error

	return
}

// FetchIndexByUserID  获取多个内容
func (obj *_MacCommentMgr) FetchIndexByUserID(userID int) (results []*MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// FetchIndexByCommentTime  获取多个内容
func (obj *_MacCommentMgr) FetchIndexByCommentTime(commentTime int) (results []*MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where("`comment_time` = ?", commentTime).Find(&results).Error

	return
}

// FetchIndexByCommentReply  获取多个内容
func (obj *_MacCommentMgr) FetchIndexByCommentReply(commentReply int16) (results []*MacComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacComment{}).Where("`comment_reply` = ?", commentReply).Find(&results).Error

	return
}
