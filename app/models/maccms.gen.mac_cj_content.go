package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _MacCjContentMgr struct {
	*_BaseMgr
}

// MacCjContentMgr open func
func MacCjContentMgr(db *gorm.DB) *_MacCjContentMgr {
	if db == nil {
		panic(fmt.Errorf("MacCjContentMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MacCjContentMgr{_BaseMgr: &_BaseMgr{DB: db.Table("mac_cj_content"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_MacCjContentMgr) Debug() *_MacCjContentMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MacCjContentMgr) GetTableName() string {
	return "mac_cj_content"
}

// Reset 重置gorm会话
func (obj *_MacCjContentMgr) Reset() *_MacCjContentMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MacCjContentMgr) Get() (result MacCjContent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjContent{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MacCjContentMgr) Gets() (results []*MacCjContent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjContent{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MacCjContentMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MacCjContent{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 内容ID
func (obj *_MacCjContentMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithNodeid nodeid获取 节点ID
func (obj *_MacCjContentMgr) WithNodeid(nodeid int) Option {
	return optionFunc(func(o *options) { o.query["nodeid"] = nodeid })
}

// WithStatus status获取 状态
func (obj *_MacCjContentMgr) WithStatus(status uint8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithURL url获取 URL
func (obj *_MacCjContentMgr) WithURL(url string) Option {
	return optionFunc(func(o *options) { o.query["url"] = url })
}

// WithTitle title获取 标题
func (obj *_MacCjContentMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithData data获取 数据
func (obj *_MacCjContentMgr) WithData(data string) Option {
	return optionFunc(func(o *options) { o.query["data"] = data })
}

// GetByOption 功能选项模式获取
func (obj *_MacCjContentMgr) GetByOption(opts ...Option) (result MacCjContent, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacCjContent{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MacCjContentMgr) GetByOptions(opts ...Option) (results []*MacCjContent, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacCjContent{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MacCjContentMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MacCjContent, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MacCjContent{}).Where(options.query)
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

// GetFromID 通过id获取内容 内容ID
func (obj *_MacCjContentMgr) GetFromID(id int) (result MacCjContent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjContent{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 内容ID
func (obj *_MacCjContentMgr) GetBatchFromID(ids []int) (results []*MacCjContent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjContent{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromNodeid 通过nodeid获取内容 节点ID
func (obj *_MacCjContentMgr) GetFromNodeid(nodeid int) (results []*MacCjContent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjContent{}).Where("`nodeid` = ?", nodeid).Find(&results).Error

	return
}

// GetBatchFromNodeid 批量查找 节点ID
func (obj *_MacCjContentMgr) GetBatchFromNodeid(nodeids []int) (results []*MacCjContent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjContent{}).Where("`nodeid` IN (?)", nodeids).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态
func (obj *_MacCjContentMgr) GetFromStatus(status uint8) (results []*MacCjContent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjContent{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态
func (obj *_MacCjContentMgr) GetBatchFromStatus(statuss []uint8) (results []*MacCjContent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjContent{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromURL 通过url获取内容 URL
func (obj *_MacCjContentMgr) GetFromURL(url string) (results []*MacCjContent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjContent{}).Where("`url` = ?", url).Find(&results).Error

	return
}

// GetBatchFromURL 批量查找 URL
func (obj *_MacCjContentMgr) GetBatchFromURL(urls []string) (results []*MacCjContent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjContent{}).Where("`url` IN (?)", urls).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 标题
func (obj *_MacCjContentMgr) GetFromTitle(title string) (results []*MacCjContent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjContent{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找 标题
func (obj *_MacCjContentMgr) GetBatchFromTitle(titles []string) (results []*MacCjContent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjContent{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromData 通过data获取内容 数据
func (obj *_MacCjContentMgr) GetFromData(data string) (results []*MacCjContent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjContent{}).Where("`data` = ?", data).Find(&results).Error

	return
}

// GetBatchFromData 批量查找 数据
func (obj *_MacCjContentMgr) GetBatchFromData(datas []string) (results []*MacCjContent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjContent{}).Where("`data` IN (?)", datas).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MacCjContentMgr) FetchByPrimaryKey(id int) (result MacCjContent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjContent{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByNodeid  获取多个内容
func (obj *_MacCjContentMgr) FetchIndexByNodeid(nodeid int) (results []*MacCjContent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjContent{}).Where("`nodeid` = ?", nodeid).Find(&results).Error

	return
}

// FetchIndexByStatus  获取多个内容
func (obj *_MacCjContentMgr) FetchIndexByStatus(status uint8) (results []*MacCjContent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjContent{}).Where("`status` = ?", status).Find(&results).Error

	return
}
