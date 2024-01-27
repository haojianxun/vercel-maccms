package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _MacWechatMgr struct {
	*_BaseMgr
}

// MacWechatMgr open func
func MacWechatMgr(db *gorm.DB) *_MacWechatMgr {
	if db == nil {
		panic(fmt.Errorf("MacWechatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MacWechatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("mac_wechat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_MacWechatMgr) Debug() *_MacWechatMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MacWechatMgr) GetTableName() string {
	return "mac_wechat"
}

// Reset 重置gorm会话
func (obj *_MacWechatMgr) Reset() *_MacWechatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MacWechatMgr) Get() (result MacWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWechat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MacWechatMgr) Gets() (results []*MacWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWechat{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MacWechatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MacWechat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_MacWechatMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithBelongWx belong_wx获取 所属账号
func (obj *_MacWechatMgr) WithBelongWx(belongWx string) Option {
	return optionFunc(func(o *options) { o.query["belong_wx"] = belongWx })
}

// WithAccount account获取 微信账号
func (obj *_MacWechatMgr) WithAccount(account string) Option {
	return optionFunc(func(o *options) { o.query["account"] = account })
}

// WithAvatar avatar获取 微信头像
func (obj *_MacWechatMgr) WithAvatar(avatar string) Option {
	return optionFunc(func(o *options) { o.query["avatar"] = avatar })
}

// WithCity city获取 微信城市
func (obj *_MacWechatMgr) WithCity(city string) Option {
	return optionFunc(func(o *options) { o.query["city"] = city })
}

// WithCountry country获取 国家
func (obj *_MacWechatMgr) WithCountry(country string) Option {
	return optionFunc(func(o *options) { o.query["country"] = country })
}

// WithLabelidList labelid_list获取 标签列表
func (obj *_MacWechatMgr) WithLabelidList(labelidList string) Option {
	return optionFunc(func(o *options) { o.query["labelid_list"] = labelidList })
}

// WithNickname nickname获取 昵称
func (obj *_MacWechatMgr) WithNickname(nickname string) Option {
	return optionFunc(func(o *options) { o.query["nickname"] = nickname })
}

// WithProvince province获取 省份
func (obj *_MacWechatMgr) WithProvince(province string) Option {
	return optionFunc(func(o *options) { o.query["province"] = province })
}

// WithRemark remark获取 备注
func (obj *_MacWechatMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithSex sex获取 性别:1-男，2-女 0-未知
func (obj *_MacWechatMgr) WithSex(sex int) Option {
	return optionFunc(func(o *options) { o.query["sex"] = sex })
}

// WithWxid wxid获取 微信ID
func (obj *_MacWechatMgr) WithWxid(wxid string) Option {
	return optionFunc(func(o *options) { o.query["wxid"] = wxid })
}

// WithCreateTs create_ts获取
func (obj *_MacWechatMgr) WithCreateTs(createTs int64) Option {
	return optionFunc(func(o *options) { o.query["create_ts"] = createTs })
}

// WithUpdateTs update_ts获取
func (obj *_MacWechatMgr) WithUpdateTs(updateTs int64) Option {
	return optionFunc(func(o *options) { o.query["update_ts"] = updateTs })
}

// WithDeleteTs delete_ts获取
func (obj *_MacWechatMgr) WithDeleteTs(deleteTs int64) Option {
	return optionFunc(func(o *options) { o.query["delete_ts"] = deleteTs })
}

// GetByOption 功能选项模式获取
func (obj *_MacWechatMgr) GetByOption(opts ...Option) (result MacWechat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacWechat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MacWechatMgr) GetByOptions(opts ...Option) (results []*MacWechat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacWechat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MacWechatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MacWechat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MacWechat{}).Where(options.query)
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

// GetFromID 通过id获取内容
func (obj *_MacWechatMgr) GetFromID(id int64) (result MacWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWechat{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_MacWechatMgr) GetBatchFromID(ids []int64) (results []*MacWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWechat{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromBelongWx 通过belong_wx获取内容 所属账号
func (obj *_MacWechatMgr) GetFromBelongWx(belongWx string) (results []*MacWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWechat{}).Where("`belong_wx` = ?", belongWx).Find(&results).Error

	return
}

// GetBatchFromBelongWx 批量查找 所属账号
func (obj *_MacWechatMgr) GetBatchFromBelongWx(belongWxs []string) (results []*MacWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWechat{}).Where("`belong_wx` IN (?)", belongWxs).Find(&results).Error

	return
}

// GetFromAccount 通过account获取内容 微信账号
func (obj *_MacWechatMgr) GetFromAccount(account string) (results []*MacWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWechat{}).Where("`account` = ?", account).Find(&results).Error

	return
}

// GetBatchFromAccount 批量查找 微信账号
func (obj *_MacWechatMgr) GetBatchFromAccount(accounts []string) (results []*MacWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWechat{}).Where("`account` IN (?)", accounts).Find(&results).Error

	return
}

// GetFromAvatar 通过avatar获取内容 微信头像
func (obj *_MacWechatMgr) GetFromAvatar(avatar string) (results []*MacWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWechat{}).Where("`avatar` = ?", avatar).Find(&results).Error

	return
}

// GetBatchFromAvatar 批量查找 微信头像
func (obj *_MacWechatMgr) GetBatchFromAvatar(avatars []string) (results []*MacWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWechat{}).Where("`avatar` IN (?)", avatars).Find(&results).Error

	return
}

// GetFromCity 通过city获取内容 微信城市
func (obj *_MacWechatMgr) GetFromCity(city string) (results []*MacWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWechat{}).Where("`city` = ?", city).Find(&results).Error

	return
}

// GetBatchFromCity 批量查找 微信城市
func (obj *_MacWechatMgr) GetBatchFromCity(citys []string) (results []*MacWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWechat{}).Where("`city` IN (?)", citys).Find(&results).Error

	return
}

// GetFromCountry 通过country获取内容 国家
func (obj *_MacWechatMgr) GetFromCountry(country string) (results []*MacWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWechat{}).Where("`country` = ?", country).Find(&results).Error

	return
}

// GetBatchFromCountry 批量查找 国家
func (obj *_MacWechatMgr) GetBatchFromCountry(countrys []string) (results []*MacWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWechat{}).Where("`country` IN (?)", countrys).Find(&results).Error

	return
}

// GetFromLabelidList 通过labelid_list获取内容 标签列表
func (obj *_MacWechatMgr) GetFromLabelidList(labelidList string) (results []*MacWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWechat{}).Where("`labelid_list` = ?", labelidList).Find(&results).Error

	return
}

// GetBatchFromLabelidList 批量查找 标签列表
func (obj *_MacWechatMgr) GetBatchFromLabelidList(labelidLists []string) (results []*MacWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWechat{}).Where("`labelid_list` IN (?)", labelidLists).Find(&results).Error

	return
}

// GetFromNickname 通过nickname获取内容 昵称
func (obj *_MacWechatMgr) GetFromNickname(nickname string) (results []*MacWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWechat{}).Where("`nickname` = ?", nickname).Find(&results).Error

	return
}

// GetBatchFromNickname 批量查找 昵称
func (obj *_MacWechatMgr) GetBatchFromNickname(nicknames []string) (results []*MacWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWechat{}).Where("`nickname` IN (?)", nicknames).Find(&results).Error

	return
}

// GetFromProvince 通过province获取内容 省份
func (obj *_MacWechatMgr) GetFromProvince(province string) (results []*MacWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWechat{}).Where("`province` = ?", province).Find(&results).Error

	return
}

// GetBatchFromProvince 批量查找 省份
func (obj *_MacWechatMgr) GetBatchFromProvince(provinces []string) (results []*MacWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWechat{}).Where("`province` IN (?)", provinces).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_MacWechatMgr) GetFromRemark(remark string) (results []*MacWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWechat{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_MacWechatMgr) GetBatchFromRemark(remarks []string) (results []*MacWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWechat{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromSex 通过sex获取内容 性别:1-男，2-女 0-未知
func (obj *_MacWechatMgr) GetFromSex(sex int) (results []*MacWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWechat{}).Where("`sex` = ?", sex).Find(&results).Error

	return
}

// GetBatchFromSex 批量查找 性别:1-男，2-女 0-未知
func (obj *_MacWechatMgr) GetBatchFromSex(sexs []int) (results []*MacWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWechat{}).Where("`sex` IN (?)", sexs).Find(&results).Error

	return
}

// GetFromWxid 通过wxid获取内容 微信ID
func (obj *_MacWechatMgr) GetFromWxid(wxid string) (results []*MacWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWechat{}).Where("`wxid` = ?", wxid).Find(&results).Error

	return
}

// GetBatchFromWxid 批量查找 微信ID
func (obj *_MacWechatMgr) GetBatchFromWxid(wxids []string) (results []*MacWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWechat{}).Where("`wxid` IN (?)", wxids).Find(&results).Error

	return
}

// GetFromCreateTs 通过create_ts获取内容
func (obj *_MacWechatMgr) GetFromCreateTs(createTs int64) (results []*MacWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWechat{}).Where("`create_ts` = ?", createTs).Find(&results).Error

	return
}

// GetBatchFromCreateTs 批量查找
func (obj *_MacWechatMgr) GetBatchFromCreateTs(createTss []int64) (results []*MacWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWechat{}).Where("`create_ts` IN (?)", createTss).Find(&results).Error

	return
}

// GetFromUpdateTs 通过update_ts获取内容
func (obj *_MacWechatMgr) GetFromUpdateTs(updateTs int64) (results []*MacWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWechat{}).Where("`update_ts` = ?", updateTs).Find(&results).Error

	return
}

// GetBatchFromUpdateTs 批量查找
func (obj *_MacWechatMgr) GetBatchFromUpdateTs(updateTss []int64) (results []*MacWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWechat{}).Where("`update_ts` IN (?)", updateTss).Find(&results).Error

	return
}

// GetFromDeleteTs 通过delete_ts获取内容
func (obj *_MacWechatMgr) GetFromDeleteTs(deleteTs int64) (results []*MacWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWechat{}).Where("`delete_ts` = ?", deleteTs).Find(&results).Error

	return
}

// GetBatchFromDeleteTs 批量查找
func (obj *_MacWechatMgr) GetBatchFromDeleteTs(deleteTss []int64) (results []*MacWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWechat{}).Where("`delete_ts` IN (?)", deleteTss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MacWechatMgr) FetchByPrimaryKey(id int64) (result MacWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWechat{}).Where("`id` = ?", id).First(&result).Error

	return
}
