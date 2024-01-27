package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _MacLinkMgr struct {
	*_BaseMgr
}

// MacLinkMgr open func
func MacLinkMgr(db *gorm.DB) *_MacLinkMgr {
	if db == nil {
		panic(fmt.Errorf("MacLinkMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MacLinkMgr{_BaseMgr: &_BaseMgr{DB: db.Table("mac_link"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_MacLinkMgr) Debug() *_MacLinkMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MacLinkMgr) GetTableName() string {
	return "mac_link"
}

// Reset 重置gorm会话
func (obj *_MacLinkMgr) Reset() *_MacLinkMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MacLinkMgr) Get() (result MacLink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacLink{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MacLinkMgr) Gets() (results []*MacLink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacLink{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MacLinkMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MacLink{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithLinkID link_id获取 友情链接ID
func (obj *_MacLinkMgr) WithLinkID(linkID uint16) Option {
	return optionFunc(func(o *options) { o.query["link_id"] = linkID })
}

// WithLinkType link_type获取 链接类型
func (obj *_MacLinkMgr) WithLinkType(linkType uint8) Option {
	return optionFunc(func(o *options) { o.query["link_type"] = linkType })
}

// WithLinkName link_name获取 链接名称
func (obj *_MacLinkMgr) WithLinkName(linkName string) Option {
	return optionFunc(func(o *options) { o.query["link_name"] = linkName })
}

// WithLinkSort link_sort获取 链接排序
func (obj *_MacLinkMgr) WithLinkSort(linkSort int16) Option {
	return optionFunc(func(o *options) { o.query["link_sort"] = linkSort })
}

// WithLinkAddTime link_add_time获取 链接添加时间
func (obj *_MacLinkMgr) WithLinkAddTime(linkAddTime int) Option {
	return optionFunc(func(o *options) { o.query["link_add_time"] = linkAddTime })
}

// WithLinkTime link_time获取 链接更新时间
func (obj *_MacLinkMgr) WithLinkTime(linkTime int) Option {
	return optionFunc(func(o *options) { o.query["link_time"] = linkTime })
}

// WithLinkURL link_url获取 链接URL
func (obj *_MacLinkMgr) WithLinkURL(linkURL string) Option {
	return optionFunc(func(o *options) { o.query["link_url"] = linkURL })
}

// WithLinkLogo link_logo获取 链接Logo
func (obj *_MacLinkMgr) WithLinkLogo(linkLogo string) Option {
	return optionFunc(func(o *options) { o.query["link_logo"] = linkLogo })
}

// GetByOption 功能选项模式获取
func (obj *_MacLinkMgr) GetByOption(opts ...Option) (result MacLink, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacLink{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MacLinkMgr) GetByOptions(opts ...Option) (results []*MacLink, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacLink{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MacLinkMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MacLink, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MacLink{}).Where(options.query)
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

// GetFromLinkID 通过link_id获取内容 友情链接ID
func (obj *_MacLinkMgr) GetFromLinkID(linkID uint16) (result MacLink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacLink{}).Where("`link_id` = ?", linkID).First(&result).Error

	return
}

// GetBatchFromLinkID 批量查找 友情链接ID
func (obj *_MacLinkMgr) GetBatchFromLinkID(linkIDs []uint16) (results []*MacLink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacLink{}).Where("`link_id` IN (?)", linkIDs).Find(&results).Error

	return
}

// GetFromLinkType 通过link_type获取内容 链接类型
func (obj *_MacLinkMgr) GetFromLinkType(linkType uint8) (results []*MacLink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacLink{}).Where("`link_type` = ?", linkType).Find(&results).Error

	return
}

// GetBatchFromLinkType 批量查找 链接类型
func (obj *_MacLinkMgr) GetBatchFromLinkType(linkTypes []uint8) (results []*MacLink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacLink{}).Where("`link_type` IN (?)", linkTypes).Find(&results).Error

	return
}

// GetFromLinkName 通过link_name获取内容 链接名称
func (obj *_MacLinkMgr) GetFromLinkName(linkName string) (results []*MacLink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacLink{}).Where("`link_name` = ?", linkName).Find(&results).Error

	return
}

// GetBatchFromLinkName 批量查找 链接名称
func (obj *_MacLinkMgr) GetBatchFromLinkName(linkNames []string) (results []*MacLink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacLink{}).Where("`link_name` IN (?)", linkNames).Find(&results).Error

	return
}

// GetFromLinkSort 通过link_sort获取内容 链接排序
func (obj *_MacLinkMgr) GetFromLinkSort(linkSort int16) (results []*MacLink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacLink{}).Where("`link_sort` = ?", linkSort).Find(&results).Error

	return
}

// GetBatchFromLinkSort 批量查找 链接排序
func (obj *_MacLinkMgr) GetBatchFromLinkSort(linkSorts []int16) (results []*MacLink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacLink{}).Where("`link_sort` IN (?)", linkSorts).Find(&results).Error

	return
}

// GetFromLinkAddTime 通过link_add_time获取内容 链接添加时间
func (obj *_MacLinkMgr) GetFromLinkAddTime(linkAddTime int) (results []*MacLink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacLink{}).Where("`link_add_time` = ?", linkAddTime).Find(&results).Error

	return
}

// GetBatchFromLinkAddTime 批量查找 链接添加时间
func (obj *_MacLinkMgr) GetBatchFromLinkAddTime(linkAddTimes []int) (results []*MacLink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacLink{}).Where("`link_add_time` IN (?)", linkAddTimes).Find(&results).Error

	return
}

// GetFromLinkTime 通过link_time获取内容 链接更新时间
func (obj *_MacLinkMgr) GetFromLinkTime(linkTime int) (results []*MacLink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacLink{}).Where("`link_time` = ?", linkTime).Find(&results).Error

	return
}

// GetBatchFromLinkTime 批量查找 链接更新时间
func (obj *_MacLinkMgr) GetBatchFromLinkTime(linkTimes []int) (results []*MacLink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacLink{}).Where("`link_time` IN (?)", linkTimes).Find(&results).Error

	return
}

// GetFromLinkURL 通过link_url获取内容 链接URL
func (obj *_MacLinkMgr) GetFromLinkURL(linkURL string) (results []*MacLink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacLink{}).Where("`link_url` = ?", linkURL).Find(&results).Error

	return
}

// GetBatchFromLinkURL 批量查找 链接URL
func (obj *_MacLinkMgr) GetBatchFromLinkURL(linkURLs []string) (results []*MacLink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacLink{}).Where("`link_url` IN (?)", linkURLs).Find(&results).Error

	return
}

// GetFromLinkLogo 通过link_logo获取内容 链接Logo
func (obj *_MacLinkMgr) GetFromLinkLogo(linkLogo string) (results []*MacLink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacLink{}).Where("`link_logo` = ?", linkLogo).Find(&results).Error

	return
}

// GetBatchFromLinkLogo 批量查找 链接Logo
func (obj *_MacLinkMgr) GetBatchFromLinkLogo(linkLogos []string) (results []*MacLink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacLink{}).Where("`link_logo` IN (?)", linkLogos).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MacLinkMgr) FetchByPrimaryKey(linkID uint16) (result MacLink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacLink{}).Where("`link_id` = ?", linkID).First(&result).Error

	return
}

// FetchIndexByLinkType  获取多个内容
func (obj *_MacLinkMgr) FetchIndexByLinkType(linkType uint8) (results []*MacLink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacLink{}).Where("`link_type` = ?", linkType).Find(&results).Error

	return
}

// FetchIndexByLinkSort  获取多个内容
func (obj *_MacLinkMgr) FetchIndexByLinkSort(linkSort int16) (results []*MacLink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacLink{}).Where("`link_sort` = ?", linkSort).Find(&results).Error

	return
}

// FetchIndexByLinkAddTime  获取多个内容
func (obj *_MacLinkMgr) FetchIndexByLinkAddTime(linkAddTime int) (results []*MacLink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacLink{}).Where("`link_add_time` = ?", linkAddTime).Find(&results).Error

	return
}

// FetchIndexByLinkTime  获取多个内容
func (obj *_MacLinkMgr) FetchIndexByLinkTime(linkTime int) (results []*MacLink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacLink{}).Where("`link_time` = ?", linkTime).Find(&results).Error

	return
}
