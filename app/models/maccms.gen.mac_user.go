package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _MacUserMgr struct {
	*_BaseMgr
}

// MacUserMgr open func
func MacUserMgr(db *gorm.DB) *_MacUserMgr {
	if db == nil {
		panic(fmt.Errorf("MacUserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MacUserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("mac_user"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_MacUserMgr) Debug() *_MacUserMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MacUserMgr) GetTableName() string {
	return "mac_user"
}

// Reset 重置gorm会话
func (obj *_MacUserMgr) Reset() *_MacUserMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MacUserMgr) Get() (result MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MacUserMgr) Gets() (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MacUserMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MacUser{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithUserID user_id获取 用户ID
func (obj *_MacUserMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithGroupID group_id获取 用户组ID
func (obj *_MacUserMgr) WithGroupID(groupID uint16) Option {
	return optionFunc(func(o *options) { o.query["group_id"] = groupID })
}

// WithUserName user_name获取 用户名
func (obj *_MacUserMgr) WithUserName(userName string) Option {
	return optionFunc(func(o *options) { o.query["user_name"] = userName })
}

// WithUserPwd user_pwd获取 用户密码
func (obj *_MacUserMgr) WithUserPwd(userPwd string) Option {
	return optionFunc(func(o *options) { o.query["user_pwd"] = userPwd })
}

// WithUserNickName user_nick_name获取 用户昵称
func (obj *_MacUserMgr) WithUserNickName(userNickName string) Option {
	return optionFunc(func(o *options) { o.query["user_nick_name"] = userNickName })
}

// WithUserQq user_qq获取 用户QQ
func (obj *_MacUserMgr) WithUserQq(userQq string) Option {
	return optionFunc(func(o *options) { o.query["user_qq"] = userQq })
}

// WithUserEmail user_email获取 用户邮箱
func (obj *_MacUserMgr) WithUserEmail(userEmail string) Option {
	return optionFunc(func(o *options) { o.query["user_email"] = userEmail })
}

// WithUserPhone user_phone获取 用户手机号
func (obj *_MacUserMgr) WithUserPhone(userPhone string) Option {
	return optionFunc(func(o *options) { o.query["user_phone"] = userPhone })
}

// WithUserStatus user_status获取 用户状态
func (obj *_MacUserMgr) WithUserStatus(userStatus uint8) Option {
	return optionFunc(func(o *options) { o.query["user_status"] = userStatus })
}

// WithUserPortrait user_portrait获取 用户头像
func (obj *_MacUserMgr) WithUserPortrait(userPortrait string) Option {
	return optionFunc(func(o *options) { o.query["user_portrait"] = userPortrait })
}

// WithUserPortraitThumb user_portrait_thumb获取 用户头像缩略图
func (obj *_MacUserMgr) WithUserPortraitThumb(userPortraitThumb string) Option {
	return optionFunc(func(o *options) { o.query["user_portrait_thumb"] = userPortraitThumb })
}

// WithUserOpenidQq user_openid_qq获取 用户QQ登录OpenID
func (obj *_MacUserMgr) WithUserOpenidQq(userOpenidQq string) Option {
	return optionFunc(func(o *options) { o.query["user_openid_qq"] = userOpenidQq })
}

// WithUserOpenidWeixin user_openid_weixin获取 用户微信登录OpenID
func (obj *_MacUserMgr) WithUserOpenidWeixin(userOpenidWeixin string) Option {
	return optionFunc(func(o *options) { o.query["user_openid_weixin"] = userOpenidWeixin })
}

// WithUserQuestion user_question获取 找回密码问题
func (obj *_MacUserMgr) WithUserQuestion(userQuestion string) Option {
	return optionFunc(func(o *options) { o.query["user_question"] = userQuestion })
}

// WithUserAnswer user_answer获取 找回密码答案
func (obj *_MacUserMgr) WithUserAnswer(userAnswer string) Option {
	return optionFunc(func(o *options) { o.query["user_answer"] = userAnswer })
}

// WithUserPoints user_points获取 用户积分
func (obj *_MacUserMgr) WithUserPoints(userPoints int) Option {
	return optionFunc(func(o *options) { o.query["user_points"] = userPoints })
}

// WithUserPointsFroze user_points_froze获取 冻结积分
func (obj *_MacUserMgr) WithUserPointsFroze(userPointsFroze int) Option {
	return optionFunc(func(o *options) { o.query["user_points_froze"] = userPointsFroze })
}

// WithUserRegTime user_reg_time获取 注册时间
func (obj *_MacUserMgr) WithUserRegTime(userRegTime int) Option {
	return optionFunc(func(o *options) { o.query["user_reg_time"] = userRegTime })
}

// WithUserRegIP user_reg_ip获取 注册IP
func (obj *_MacUserMgr) WithUserRegIP(userRegIP int) Option {
	return optionFunc(func(o *options) { o.query["user_reg_ip"] = userRegIP })
}

// WithUserLoginTime user_login_time获取 登录时间
func (obj *_MacUserMgr) WithUserLoginTime(userLoginTime int) Option {
	return optionFunc(func(o *options) { o.query["user_login_time"] = userLoginTime })
}

// WithUserLoginIP user_login_ip获取 登录IP
func (obj *_MacUserMgr) WithUserLoginIP(userLoginIP int) Option {
	return optionFunc(func(o *options) { o.query["user_login_ip"] = userLoginIP })
}

// WithUserLastLoginTime user_last_login_time获取 上次登录时间
func (obj *_MacUserMgr) WithUserLastLoginTime(userLastLoginTime int) Option {
	return optionFunc(func(o *options) { o.query["user_last_login_time"] = userLastLoginTime })
}

// WithUserLastLoginIP user_last_login_ip获取 上次登录IP
func (obj *_MacUserMgr) WithUserLastLoginIP(userLastLoginIP int) Option {
	return optionFunc(func(o *options) { o.query["user_last_login_ip"] = userLastLoginIP })
}

// WithUserLoginNum user_login_num获取 登录次数
func (obj *_MacUserMgr) WithUserLoginNum(userLoginNum uint16) Option {
	return optionFunc(func(o *options) { o.query["user_login_num"] = userLoginNum })
}

// WithUserExtend user_extend获取 用户扩展字段
func (obj *_MacUserMgr) WithUserExtend(userExtend uint16) Option {
	return optionFunc(func(o *options) { o.query["user_extend"] = userExtend })
}

// WithUserRandom user_random获取 用户随机码
func (obj *_MacUserMgr) WithUserRandom(userRandom string) Option {
	return optionFunc(func(o *options) { o.query["user_random"] = userRandom })
}

// WithUserEndTime user_end_time获取 VIP到期时间
func (obj *_MacUserMgr) WithUserEndTime(userEndTime int) Option {
	return optionFunc(func(o *options) { o.query["user_end_time"] = userEndTime })
}

// WithUserPid user_pid获取 上级推广员ID
func (obj *_MacUserMgr) WithUserPid(userPid int) Option {
	return optionFunc(func(o *options) { o.query["user_pid"] = userPid })
}

// WithUserPid2 user_pid_2获取 上级推广员ID_2
func (obj *_MacUserMgr) WithUserPid2(userPid2 int) Option {
	return optionFunc(func(o *options) { o.query["user_pid_2"] = userPid2 })
}

// WithUserPid3 user_pid_3获取 上级推广员ID_3
func (obj *_MacUserMgr) WithUserPid3(userPid3 int) Option {
	return optionFunc(func(o *options) { o.query["user_pid_3"] = userPid3 })
}

// GetByOption 功能选项模式获取
func (obj *_MacUserMgr) GetByOption(opts ...Option) (result MacUser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MacUserMgr) GetByOptions(opts ...Option) (results []*MacUser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MacUserMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MacUser, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where(options.query)
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

// GetFromUserID 通过user_id获取内容 用户ID
func (obj *_MacUserMgr) GetFromUserID(userID int) (result MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_id` = ?", userID).First(&result).Error

	return
}

// GetBatchFromUserID 批量查找 用户ID
func (obj *_MacUserMgr) GetBatchFromUserID(userIDs []int) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromGroupID 通过group_id获取内容 用户组ID
func (obj *_MacUserMgr) GetFromGroupID(groupID uint16) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`group_id` = ?", groupID).Find(&results).Error

	return
}

// GetBatchFromGroupID 批量查找 用户组ID
func (obj *_MacUserMgr) GetBatchFromGroupID(groupIDs []uint16) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`group_id` IN (?)", groupIDs).Find(&results).Error

	return
}

// GetFromUserName 通过user_name获取内容 用户名
func (obj *_MacUserMgr) GetFromUserName(userName string) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_name` = ?", userName).Find(&results).Error

	return
}

// GetBatchFromUserName 批量查找 用户名
func (obj *_MacUserMgr) GetBatchFromUserName(userNames []string) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_name` IN (?)", userNames).Find(&results).Error

	return
}

// GetFromUserPwd 通过user_pwd获取内容 用户密码
func (obj *_MacUserMgr) GetFromUserPwd(userPwd string) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_pwd` = ?", userPwd).Find(&results).Error

	return
}

// GetBatchFromUserPwd 批量查找 用户密码
func (obj *_MacUserMgr) GetBatchFromUserPwd(userPwds []string) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_pwd` IN (?)", userPwds).Find(&results).Error

	return
}

// GetFromUserNickName 通过user_nick_name获取内容 用户昵称
func (obj *_MacUserMgr) GetFromUserNickName(userNickName string) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_nick_name` = ?", userNickName).Find(&results).Error

	return
}

// GetBatchFromUserNickName 批量查找 用户昵称
func (obj *_MacUserMgr) GetBatchFromUserNickName(userNickNames []string) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_nick_name` IN (?)", userNickNames).Find(&results).Error

	return
}

// GetFromUserQq 通过user_qq获取内容 用户QQ
func (obj *_MacUserMgr) GetFromUserQq(userQq string) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_qq` = ?", userQq).Find(&results).Error

	return
}

// GetBatchFromUserQq 批量查找 用户QQ
func (obj *_MacUserMgr) GetBatchFromUserQq(userQqs []string) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_qq` IN (?)", userQqs).Find(&results).Error

	return
}

// GetFromUserEmail 通过user_email获取内容 用户邮箱
func (obj *_MacUserMgr) GetFromUserEmail(userEmail string) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_email` = ?", userEmail).Find(&results).Error

	return
}

// GetBatchFromUserEmail 批量查找 用户邮箱
func (obj *_MacUserMgr) GetBatchFromUserEmail(userEmails []string) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_email` IN (?)", userEmails).Find(&results).Error

	return
}

// GetFromUserPhone 通过user_phone获取内容 用户手机号
func (obj *_MacUserMgr) GetFromUserPhone(userPhone string) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_phone` = ?", userPhone).Find(&results).Error

	return
}

// GetBatchFromUserPhone 批量查找 用户手机号
func (obj *_MacUserMgr) GetBatchFromUserPhone(userPhones []string) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_phone` IN (?)", userPhones).Find(&results).Error

	return
}

// GetFromUserStatus 通过user_status获取内容 用户状态
func (obj *_MacUserMgr) GetFromUserStatus(userStatus uint8) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_status` = ?", userStatus).Find(&results).Error

	return
}

// GetBatchFromUserStatus 批量查找 用户状态
func (obj *_MacUserMgr) GetBatchFromUserStatus(userStatuss []uint8) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_status` IN (?)", userStatuss).Find(&results).Error

	return
}

// GetFromUserPortrait 通过user_portrait获取内容 用户头像
func (obj *_MacUserMgr) GetFromUserPortrait(userPortrait string) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_portrait` = ?", userPortrait).Find(&results).Error

	return
}

// GetBatchFromUserPortrait 批量查找 用户头像
func (obj *_MacUserMgr) GetBatchFromUserPortrait(userPortraits []string) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_portrait` IN (?)", userPortraits).Find(&results).Error

	return
}

// GetFromUserPortraitThumb 通过user_portrait_thumb获取内容 用户头像缩略图
func (obj *_MacUserMgr) GetFromUserPortraitThumb(userPortraitThumb string) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_portrait_thumb` = ?", userPortraitThumb).Find(&results).Error

	return
}

// GetBatchFromUserPortraitThumb 批量查找 用户头像缩略图
func (obj *_MacUserMgr) GetBatchFromUserPortraitThumb(userPortraitThumbs []string) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_portrait_thumb` IN (?)", userPortraitThumbs).Find(&results).Error

	return
}

// GetFromUserOpenidQq 通过user_openid_qq获取内容 用户QQ登录OpenID
func (obj *_MacUserMgr) GetFromUserOpenidQq(userOpenidQq string) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_openid_qq` = ?", userOpenidQq).Find(&results).Error

	return
}

// GetBatchFromUserOpenidQq 批量查找 用户QQ登录OpenID
func (obj *_MacUserMgr) GetBatchFromUserOpenidQq(userOpenidQqs []string) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_openid_qq` IN (?)", userOpenidQqs).Find(&results).Error

	return
}

// GetFromUserOpenidWeixin 通过user_openid_weixin获取内容 用户微信登录OpenID
func (obj *_MacUserMgr) GetFromUserOpenidWeixin(userOpenidWeixin string) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_openid_weixin` = ?", userOpenidWeixin).Find(&results).Error

	return
}

// GetBatchFromUserOpenidWeixin 批量查找 用户微信登录OpenID
func (obj *_MacUserMgr) GetBatchFromUserOpenidWeixin(userOpenidWeixins []string) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_openid_weixin` IN (?)", userOpenidWeixins).Find(&results).Error

	return
}

// GetFromUserQuestion 通过user_question获取内容 找回密码问题
func (obj *_MacUserMgr) GetFromUserQuestion(userQuestion string) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_question` = ?", userQuestion).Find(&results).Error

	return
}

// GetBatchFromUserQuestion 批量查找 找回密码问题
func (obj *_MacUserMgr) GetBatchFromUserQuestion(userQuestions []string) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_question` IN (?)", userQuestions).Find(&results).Error

	return
}

// GetFromUserAnswer 通过user_answer获取内容 找回密码答案
func (obj *_MacUserMgr) GetFromUserAnswer(userAnswer string) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_answer` = ?", userAnswer).Find(&results).Error

	return
}

// GetBatchFromUserAnswer 批量查找 找回密码答案
func (obj *_MacUserMgr) GetBatchFromUserAnswer(userAnswers []string) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_answer` IN (?)", userAnswers).Find(&results).Error

	return
}

// GetFromUserPoints 通过user_points获取内容 用户积分
func (obj *_MacUserMgr) GetFromUserPoints(userPoints int) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_points` = ?", userPoints).Find(&results).Error

	return
}

// GetBatchFromUserPoints 批量查找 用户积分
func (obj *_MacUserMgr) GetBatchFromUserPoints(userPointss []int) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_points` IN (?)", userPointss).Find(&results).Error

	return
}

// GetFromUserPointsFroze 通过user_points_froze获取内容 冻结积分
func (obj *_MacUserMgr) GetFromUserPointsFroze(userPointsFroze int) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_points_froze` = ?", userPointsFroze).Find(&results).Error

	return
}

// GetBatchFromUserPointsFroze 批量查找 冻结积分
func (obj *_MacUserMgr) GetBatchFromUserPointsFroze(userPointsFrozes []int) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_points_froze` IN (?)", userPointsFrozes).Find(&results).Error

	return
}

// GetFromUserRegTime 通过user_reg_time获取内容 注册时间
func (obj *_MacUserMgr) GetFromUserRegTime(userRegTime int) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_reg_time` = ?", userRegTime).Find(&results).Error

	return
}

// GetBatchFromUserRegTime 批量查找 注册时间
func (obj *_MacUserMgr) GetBatchFromUserRegTime(userRegTimes []int) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_reg_time` IN (?)", userRegTimes).Find(&results).Error

	return
}

// GetFromUserRegIP 通过user_reg_ip获取内容 注册IP
func (obj *_MacUserMgr) GetFromUserRegIP(userRegIP int) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_reg_ip` = ?", userRegIP).Find(&results).Error

	return
}

// GetBatchFromUserRegIP 批量查找 注册IP
func (obj *_MacUserMgr) GetBatchFromUserRegIP(userRegIPs []int) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_reg_ip` IN (?)", userRegIPs).Find(&results).Error

	return
}

// GetFromUserLoginTime 通过user_login_time获取内容 登录时间
func (obj *_MacUserMgr) GetFromUserLoginTime(userLoginTime int) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_login_time` = ?", userLoginTime).Find(&results).Error

	return
}

// GetBatchFromUserLoginTime 批量查找 登录时间
func (obj *_MacUserMgr) GetBatchFromUserLoginTime(userLoginTimes []int) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_login_time` IN (?)", userLoginTimes).Find(&results).Error

	return
}

// GetFromUserLoginIP 通过user_login_ip获取内容 登录IP
func (obj *_MacUserMgr) GetFromUserLoginIP(userLoginIP int) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_login_ip` = ?", userLoginIP).Find(&results).Error

	return
}

// GetBatchFromUserLoginIP 批量查找 登录IP
func (obj *_MacUserMgr) GetBatchFromUserLoginIP(userLoginIPs []int) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_login_ip` IN (?)", userLoginIPs).Find(&results).Error

	return
}

// GetFromUserLastLoginTime 通过user_last_login_time获取内容 上次登录时间
func (obj *_MacUserMgr) GetFromUserLastLoginTime(userLastLoginTime int) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_last_login_time` = ?", userLastLoginTime).Find(&results).Error

	return
}

// GetBatchFromUserLastLoginTime 批量查找 上次登录时间
func (obj *_MacUserMgr) GetBatchFromUserLastLoginTime(userLastLoginTimes []int) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_last_login_time` IN (?)", userLastLoginTimes).Find(&results).Error

	return
}

// GetFromUserLastLoginIP 通过user_last_login_ip获取内容 上次登录IP
func (obj *_MacUserMgr) GetFromUserLastLoginIP(userLastLoginIP int) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_last_login_ip` = ?", userLastLoginIP).Find(&results).Error

	return
}

// GetBatchFromUserLastLoginIP 批量查找 上次登录IP
func (obj *_MacUserMgr) GetBatchFromUserLastLoginIP(userLastLoginIPs []int) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_last_login_ip` IN (?)", userLastLoginIPs).Find(&results).Error

	return
}

// GetFromUserLoginNum 通过user_login_num获取内容 登录次数
func (obj *_MacUserMgr) GetFromUserLoginNum(userLoginNum uint16) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_login_num` = ?", userLoginNum).Find(&results).Error

	return
}

// GetBatchFromUserLoginNum 批量查找 登录次数
func (obj *_MacUserMgr) GetBatchFromUserLoginNum(userLoginNums []uint16) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_login_num` IN (?)", userLoginNums).Find(&results).Error

	return
}

// GetFromUserExtend 通过user_extend获取内容 用户扩展字段
func (obj *_MacUserMgr) GetFromUserExtend(userExtend uint16) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_extend` = ?", userExtend).Find(&results).Error

	return
}

// GetBatchFromUserExtend 批量查找 用户扩展字段
func (obj *_MacUserMgr) GetBatchFromUserExtend(userExtends []uint16) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_extend` IN (?)", userExtends).Find(&results).Error

	return
}

// GetFromUserRandom 通过user_random获取内容 用户随机码
func (obj *_MacUserMgr) GetFromUserRandom(userRandom string) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_random` = ?", userRandom).Find(&results).Error

	return
}

// GetBatchFromUserRandom 批量查找 用户随机码
func (obj *_MacUserMgr) GetBatchFromUserRandom(userRandoms []string) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_random` IN (?)", userRandoms).Find(&results).Error

	return
}

// GetFromUserEndTime 通过user_end_time获取内容 VIP到期时间
func (obj *_MacUserMgr) GetFromUserEndTime(userEndTime int) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_end_time` = ?", userEndTime).Find(&results).Error

	return
}

// GetBatchFromUserEndTime 批量查找 VIP到期时间
func (obj *_MacUserMgr) GetBatchFromUserEndTime(userEndTimes []int) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_end_time` IN (?)", userEndTimes).Find(&results).Error

	return
}

// GetFromUserPid 通过user_pid获取内容 上级推广员ID
func (obj *_MacUserMgr) GetFromUserPid(userPid int) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_pid` = ?", userPid).Find(&results).Error

	return
}

// GetBatchFromUserPid 批量查找 上级推广员ID
func (obj *_MacUserMgr) GetBatchFromUserPid(userPids []int) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_pid` IN (?)", userPids).Find(&results).Error

	return
}

// GetFromUserPid2 通过user_pid_2获取内容 上级推广员ID_2
func (obj *_MacUserMgr) GetFromUserPid2(userPid2 int) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_pid_2` = ?", userPid2).Find(&results).Error

	return
}

// GetBatchFromUserPid2 批量查找 上级推广员ID_2
func (obj *_MacUserMgr) GetBatchFromUserPid2(userPid2s []int) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_pid_2` IN (?)", userPid2s).Find(&results).Error

	return
}

// GetFromUserPid3 通过user_pid_3获取内容 上级推广员ID_3
func (obj *_MacUserMgr) GetFromUserPid3(userPid3 int) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_pid_3` = ?", userPid3).Find(&results).Error

	return
}

// GetBatchFromUserPid3 批量查找 上级推广员ID_3
func (obj *_MacUserMgr) GetBatchFromUserPid3(userPid3s []int) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_pid_3` IN (?)", userPid3s).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MacUserMgr) FetchByPrimaryKey(userID int) (result MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_id` = ?", userID).First(&result).Error

	return
}

// FetchIndexByTypeID  获取多个内容
func (obj *_MacUserMgr) FetchIndexByTypeID(groupID uint16) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`group_id` = ?", groupID).Find(&results).Error

	return
}

// FetchIndexByUserName  获取多个内容
func (obj *_MacUserMgr) FetchIndexByUserName(userName string) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_name` = ?", userName).Find(&results).Error

	return
}

// FetchIndexByUserRegTime  获取多个内容
func (obj *_MacUserMgr) FetchIndexByUserRegTime(userRegTime int) (results []*MacUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacUser{}).Where("`user_reg_time` = ?", userRegTime).Find(&results).Error

	return
}
