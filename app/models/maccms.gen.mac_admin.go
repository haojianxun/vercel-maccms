package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _MacAdminMgr struct {
	*_BaseMgr
}

// MacAdminMgr open func
func MacAdminMgr(db *gorm.DB) *_MacAdminMgr {
	if db == nil {
		panic(fmt.Errorf("MacAdminMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MacAdminMgr{_BaseMgr: &_BaseMgr{DB: db.Table("mac_admin"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_MacAdminMgr) Debug() *_MacAdminMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MacAdminMgr) GetTableName() string {
	return "mac_admin"
}

// Reset 重置gorm会话
func (obj *_MacAdminMgr) Reset() *_MacAdminMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MacAdminMgr) Get() (result MacAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAdmin{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MacAdminMgr) Gets() (results []*MacAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAdmin{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MacAdminMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MacAdmin{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithAdminID admin_id获取 管理员ID
func (obj *_MacAdminMgr) WithAdminID(adminID uint16) Option {
	return optionFunc(func(o *options) { o.query["admin_id"] = adminID })
}

// WithAdminName admin_name获取 管理员姓名
func (obj *_MacAdminMgr) WithAdminName(adminName string) Option {
	return optionFunc(func(o *options) { o.query["admin_name"] = adminName })
}

// WithAdminPwd admin_pwd获取 管理员密码
func (obj *_MacAdminMgr) WithAdminPwd(adminPwd string) Option {
	return optionFunc(func(o *options) { o.query["admin_pwd"] = adminPwd })
}

// WithAdminRandom admin_random获取 管理员随机值
func (obj *_MacAdminMgr) WithAdminRandom(adminRandom string) Option {
	return optionFunc(func(o *options) { o.query["admin_random"] = adminRandom })
}

// WithAdminStatus admin_status获取 管理员状态
func (obj *_MacAdminMgr) WithAdminStatus(adminStatus uint8) Option {
	return optionFunc(func(o *options) { o.query["admin_status"] = adminStatus })
}

// WithAdminAuth admin_auth获取 管理员权限
func (obj *_MacAdminMgr) WithAdminAuth(adminAuth string) Option {
	return optionFunc(func(o *options) { o.query["admin_auth"] = adminAuth })
}

// WithAdminLoginTime admin_login_time获取 管理员登录时间
func (obj *_MacAdminMgr) WithAdminLoginTime(adminLoginTime int) Option {
	return optionFunc(func(o *options) { o.query["admin_login_time"] = adminLoginTime })
}

// WithAdminLoginIP admin_login_ip获取 管理员登录IP
func (obj *_MacAdminMgr) WithAdminLoginIP(adminLoginIP int) Option {
	return optionFunc(func(o *options) { o.query["admin_login_ip"] = adminLoginIP })
}

// WithAdminLoginNum admin_login_num获取 管理员登录次数
func (obj *_MacAdminMgr) WithAdminLoginNum(adminLoginNum int) Option {
	return optionFunc(func(o *options) { o.query["admin_login_num"] = adminLoginNum })
}

// WithAdminLastLoginTime admin_last_login_time获取 管理员上次登录时间
func (obj *_MacAdminMgr) WithAdminLastLoginTime(adminLastLoginTime int) Option {
	return optionFunc(func(o *options) { o.query["admin_last_login_time"] = adminLastLoginTime })
}

// WithAdminLastLoginIP admin_last_login_ip获取 管理员上次登录IP
func (obj *_MacAdminMgr) WithAdminLastLoginIP(adminLastLoginIP int) Option {
	return optionFunc(func(o *options) { o.query["admin_last_login_ip"] = adminLastLoginIP })
}

// GetByOption 功能选项模式获取
func (obj *_MacAdminMgr) GetByOption(opts ...Option) (result MacAdmin, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacAdmin{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MacAdminMgr) GetByOptions(opts ...Option) (results []*MacAdmin, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacAdmin{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MacAdminMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MacAdmin, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MacAdmin{}).Where(options.query)
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

// GetFromAdminID 通过admin_id获取内容 管理员ID
func (obj *_MacAdminMgr) GetFromAdminID(adminID uint16) (result MacAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAdmin{}).Where("`admin_id` = ?", adminID).First(&result).Error

	return
}

// GetBatchFromAdminID 批量查找 管理员ID
func (obj *_MacAdminMgr) GetBatchFromAdminID(adminIDs []uint16) (results []*MacAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAdmin{}).Where("`admin_id` IN (?)", adminIDs).Find(&results).Error

	return
}

// GetFromAdminName 通过admin_name获取内容 管理员姓名
func (obj *_MacAdminMgr) GetFromAdminName(adminName string) (results []*MacAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAdmin{}).Where("`admin_name` = ?", adminName).Find(&results).Error

	return
}

// GetBatchFromAdminName 批量查找 管理员姓名
func (obj *_MacAdminMgr) GetBatchFromAdminName(adminNames []string) (results []*MacAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAdmin{}).Where("`admin_name` IN (?)", adminNames).Find(&results).Error

	return
}

// GetFromAdminPwd 通过admin_pwd获取内容 管理员密码
func (obj *_MacAdminMgr) GetFromAdminPwd(adminPwd string) (results []*MacAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAdmin{}).Where("`admin_pwd` = ?", adminPwd).Find(&results).Error

	return
}

// GetBatchFromAdminPwd 批量查找 管理员密码
func (obj *_MacAdminMgr) GetBatchFromAdminPwd(adminPwds []string) (results []*MacAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAdmin{}).Where("`admin_pwd` IN (?)", adminPwds).Find(&results).Error

	return
}

// GetFromAdminRandom 通过admin_random获取内容 管理员随机值
func (obj *_MacAdminMgr) GetFromAdminRandom(adminRandom string) (results []*MacAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAdmin{}).Where("`admin_random` = ?", adminRandom).Find(&results).Error

	return
}

// GetBatchFromAdminRandom 批量查找 管理员随机值
func (obj *_MacAdminMgr) GetBatchFromAdminRandom(adminRandoms []string) (results []*MacAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAdmin{}).Where("`admin_random` IN (?)", adminRandoms).Find(&results).Error

	return
}

// GetFromAdminStatus 通过admin_status获取内容 管理员状态
func (obj *_MacAdminMgr) GetFromAdminStatus(adminStatus uint8) (results []*MacAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAdmin{}).Where("`admin_status` = ?", adminStatus).Find(&results).Error

	return
}

// GetBatchFromAdminStatus 批量查找 管理员状态
func (obj *_MacAdminMgr) GetBatchFromAdminStatus(adminStatuss []uint8) (results []*MacAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAdmin{}).Where("`admin_status` IN (?)", adminStatuss).Find(&results).Error

	return
}

// GetFromAdminAuth 通过admin_auth获取内容 管理员权限
func (obj *_MacAdminMgr) GetFromAdminAuth(adminAuth string) (results []*MacAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAdmin{}).Where("`admin_auth` = ?", adminAuth).Find(&results).Error

	return
}

// GetBatchFromAdminAuth 批量查找 管理员权限
func (obj *_MacAdminMgr) GetBatchFromAdminAuth(adminAuths []string) (results []*MacAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAdmin{}).Where("`admin_auth` IN (?)", adminAuths).Find(&results).Error

	return
}

// GetFromAdminLoginTime 通过admin_login_time获取内容 管理员登录时间
func (obj *_MacAdminMgr) GetFromAdminLoginTime(adminLoginTime int) (results []*MacAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAdmin{}).Where("`admin_login_time` = ?", adminLoginTime).Find(&results).Error

	return
}

// GetBatchFromAdminLoginTime 批量查找 管理员登录时间
func (obj *_MacAdminMgr) GetBatchFromAdminLoginTime(adminLoginTimes []int) (results []*MacAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAdmin{}).Where("`admin_login_time` IN (?)", adminLoginTimes).Find(&results).Error

	return
}

// GetFromAdminLoginIP 通过admin_login_ip获取内容 管理员登录IP
func (obj *_MacAdminMgr) GetFromAdminLoginIP(adminLoginIP int) (results []*MacAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAdmin{}).Where("`admin_login_ip` = ?", adminLoginIP).Find(&results).Error

	return
}

// GetBatchFromAdminLoginIP 批量查找 管理员登录IP
func (obj *_MacAdminMgr) GetBatchFromAdminLoginIP(adminLoginIPs []int) (results []*MacAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAdmin{}).Where("`admin_login_ip` IN (?)", adminLoginIPs).Find(&results).Error

	return
}

// GetFromAdminLoginNum 通过admin_login_num获取内容 管理员登录次数
func (obj *_MacAdminMgr) GetFromAdminLoginNum(adminLoginNum int) (results []*MacAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAdmin{}).Where("`admin_login_num` = ?", adminLoginNum).Find(&results).Error

	return
}

// GetBatchFromAdminLoginNum 批量查找 管理员登录次数
func (obj *_MacAdminMgr) GetBatchFromAdminLoginNum(adminLoginNums []int) (results []*MacAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAdmin{}).Where("`admin_login_num` IN (?)", adminLoginNums).Find(&results).Error

	return
}

// GetFromAdminLastLoginTime 通过admin_last_login_time获取内容 管理员上次登录时间
func (obj *_MacAdminMgr) GetFromAdminLastLoginTime(adminLastLoginTime int) (results []*MacAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAdmin{}).Where("`admin_last_login_time` = ?", adminLastLoginTime).Find(&results).Error

	return
}

// GetBatchFromAdminLastLoginTime 批量查找 管理员上次登录时间
func (obj *_MacAdminMgr) GetBatchFromAdminLastLoginTime(adminLastLoginTimes []int) (results []*MacAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAdmin{}).Where("`admin_last_login_time` IN (?)", adminLastLoginTimes).Find(&results).Error

	return
}

// GetFromAdminLastLoginIP 通过admin_last_login_ip获取内容 管理员上次登录IP
func (obj *_MacAdminMgr) GetFromAdminLastLoginIP(adminLastLoginIP int) (results []*MacAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAdmin{}).Where("`admin_last_login_ip` = ?", adminLastLoginIP).Find(&results).Error

	return
}

// GetBatchFromAdminLastLoginIP 批量查找 管理员上次登录IP
func (obj *_MacAdminMgr) GetBatchFromAdminLastLoginIP(adminLastLoginIPs []int) (results []*MacAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAdmin{}).Where("`admin_last_login_ip` IN (?)", adminLastLoginIPs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MacAdminMgr) FetchByPrimaryKey(adminID uint16) (result MacAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAdmin{}).Where("`admin_id` = ?", adminID).First(&result).Error

	return
}

// FetchIndexByAdminName  获取多个内容
func (obj *_MacAdminMgr) FetchIndexByAdminName(adminName string) (results []*MacAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacAdmin{}).Where("`admin_name` = ?", adminName).Find(&results).Error

	return
}
