package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _MacVodSearchMgr struct {
	*_BaseMgr
}

// MacVodSearchMgr open func
func MacVodSearchMgr(db *gorm.DB) *_MacVodSearchMgr {
	if db == nil {
		panic(fmt.Errorf("MacVodSearchMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MacVodSearchMgr{_BaseMgr: &_BaseMgr{DB: db.Table("mac_vod_search"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_MacVodSearchMgr) Debug() *_MacVodSearchMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MacVodSearchMgr) GetTableName() string {
	return "mac_vod_search"
}

// Reset 重置gorm会话
func (obj *_MacVodSearchMgr) Reset() *_MacVodSearchMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MacVodSearchMgr) Get() (result MacVodSearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVodSearch{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MacVodSearchMgr) Gets() (results []*MacVodSearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVodSearch{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MacVodSearchMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MacVodSearch{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSearchKey search_key获取 搜索键（关键词md5）
func (obj *_MacVodSearchMgr) WithSearchKey(searchKey string) Option {
	return optionFunc(func(o *options) { o.query["search_key"] = searchKey })
}

// WithSearchWord search_word获取 搜索关键词
func (obj *_MacVodSearchMgr) WithSearchWord(searchWord string) Option {
	return optionFunc(func(o *options) { o.query["search_word"] = searchWord })
}

// WithSearchField search_field获取 搜索字段名（可有多个，用|分隔）
func (obj *_MacVodSearchMgr) WithSearchField(searchField string) Option {
	return optionFunc(func(o *options) { o.query["search_field"] = searchField })
}

// WithSearchHitCount search_hit_count获取 搜索命中次数
func (obj *_MacVodSearchMgr) WithSearchHitCount(searchHitCount uint64) Option {
	return optionFunc(func(o *options) { o.query["search_hit_count"] = searchHitCount })
}

// WithSearchLastHitTime search_last_hit_time获取 最近命中时间
func (obj *_MacVodSearchMgr) WithSearchLastHitTime(searchLastHitTime int) Option {
	return optionFunc(func(o *options) { o.query["search_last_hit_time"] = searchLastHitTime })
}

// WithSearchUpdateTime search_update_time获取 添加时间
func (obj *_MacVodSearchMgr) WithSearchUpdateTime(searchUpdateTime int) Option {
	return optionFunc(func(o *options) { o.query["search_update_time"] = searchUpdateTime })
}

// WithSearchResultCount search_result_count获取 结果Id数量
func (obj *_MacVodSearchMgr) WithSearchResultCount(searchResultCount int) Option {
	return optionFunc(func(o *options) { o.query["search_result_count"] = searchResultCount })
}

// WithSearchResultIDs search_result_ids获取 搜索结果Id列表，英文半角逗号分隔
func (obj *_MacVodSearchMgr) WithSearchResultIDs(searchResultIDs string) Option {
	return optionFunc(func(o *options) { o.query["search_result_ids"] = searchResultIDs })
}

// GetByOption 功能选项模式获取
func (obj *_MacVodSearchMgr) GetByOption(opts ...Option) (result MacVodSearch, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacVodSearch{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MacVodSearchMgr) GetByOptions(opts ...Option) (results []*MacVodSearch, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacVodSearch{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MacVodSearchMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MacVodSearch, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MacVodSearch{}).Where(options.query)
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

// GetFromSearchKey 通过search_key获取内容 搜索键（关键词md5）
func (obj *_MacVodSearchMgr) GetFromSearchKey(searchKey string) (result MacVodSearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVodSearch{}).Where("`search_key` = ?", searchKey).First(&result).Error

	return
}

// GetBatchFromSearchKey 批量查找 搜索键（关键词md5）
func (obj *_MacVodSearchMgr) GetBatchFromSearchKey(searchKeys []string) (results []*MacVodSearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVodSearch{}).Where("`search_key` IN (?)", searchKeys).Find(&results).Error

	return
}

// GetFromSearchWord 通过search_word获取内容 搜索关键词
func (obj *_MacVodSearchMgr) GetFromSearchWord(searchWord string) (results []*MacVodSearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVodSearch{}).Where("`search_word` = ?", searchWord).Find(&results).Error

	return
}

// GetBatchFromSearchWord 批量查找 搜索关键词
func (obj *_MacVodSearchMgr) GetBatchFromSearchWord(searchWords []string) (results []*MacVodSearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVodSearch{}).Where("`search_word` IN (?)", searchWords).Find(&results).Error

	return
}

// GetFromSearchField 通过search_field获取内容 搜索字段名（可有多个，用|分隔）
func (obj *_MacVodSearchMgr) GetFromSearchField(searchField string) (results []*MacVodSearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVodSearch{}).Where("`search_field` = ?", searchField).Find(&results).Error

	return
}

// GetBatchFromSearchField 批量查找 搜索字段名（可有多个，用|分隔）
func (obj *_MacVodSearchMgr) GetBatchFromSearchField(searchFields []string) (results []*MacVodSearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVodSearch{}).Where("`search_field` IN (?)", searchFields).Find(&results).Error

	return
}

// GetFromSearchHitCount 通过search_hit_count获取内容 搜索命中次数
func (obj *_MacVodSearchMgr) GetFromSearchHitCount(searchHitCount uint64) (results []*MacVodSearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVodSearch{}).Where("`search_hit_count` = ?", searchHitCount).Find(&results).Error

	return
}

// GetBatchFromSearchHitCount 批量查找 搜索命中次数
func (obj *_MacVodSearchMgr) GetBatchFromSearchHitCount(searchHitCounts []uint64) (results []*MacVodSearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVodSearch{}).Where("`search_hit_count` IN (?)", searchHitCounts).Find(&results).Error

	return
}

// GetFromSearchLastHitTime 通过search_last_hit_time获取内容 最近命中时间
func (obj *_MacVodSearchMgr) GetFromSearchLastHitTime(searchLastHitTime int) (results []*MacVodSearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVodSearch{}).Where("`search_last_hit_time` = ?", searchLastHitTime).Find(&results).Error

	return
}

// GetBatchFromSearchLastHitTime 批量查找 最近命中时间
func (obj *_MacVodSearchMgr) GetBatchFromSearchLastHitTime(searchLastHitTimes []int) (results []*MacVodSearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVodSearch{}).Where("`search_last_hit_time` IN (?)", searchLastHitTimes).Find(&results).Error

	return
}

// GetFromSearchUpdateTime 通过search_update_time获取内容 添加时间
func (obj *_MacVodSearchMgr) GetFromSearchUpdateTime(searchUpdateTime int) (results []*MacVodSearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVodSearch{}).Where("`search_update_time` = ?", searchUpdateTime).Find(&results).Error

	return
}

// GetBatchFromSearchUpdateTime 批量查找 添加时间
func (obj *_MacVodSearchMgr) GetBatchFromSearchUpdateTime(searchUpdateTimes []int) (results []*MacVodSearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVodSearch{}).Where("`search_update_time` IN (?)", searchUpdateTimes).Find(&results).Error

	return
}

// GetFromSearchResultCount 通过search_result_count获取内容 结果Id数量
func (obj *_MacVodSearchMgr) GetFromSearchResultCount(searchResultCount int) (results []*MacVodSearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVodSearch{}).Where("`search_result_count` = ?", searchResultCount).Find(&results).Error

	return
}

// GetBatchFromSearchResultCount 批量查找 结果Id数量
func (obj *_MacVodSearchMgr) GetBatchFromSearchResultCount(searchResultCounts []int) (results []*MacVodSearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVodSearch{}).Where("`search_result_count` IN (?)", searchResultCounts).Find(&results).Error

	return
}

// GetFromSearchResultIDs 通过search_result_ids获取内容 搜索结果Id列表，英文半角逗号分隔
func (obj *_MacVodSearchMgr) GetFromSearchResultIDs(searchResultIDs string) (results []*MacVodSearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVodSearch{}).Where("`search_result_ids` = ?", searchResultIDs).Find(&results).Error

	return
}

// GetBatchFromSearchResultIDs 批量查找 搜索结果Id列表，英文半角逗号分隔
func (obj *_MacVodSearchMgr) GetBatchFromSearchResultIDs(searchResultIDss []string) (results []*MacVodSearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVodSearch{}).Where("`search_result_ids` IN (?)", searchResultIDss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MacVodSearchMgr) FetchByPrimaryKey(searchKey string) (result MacVodSearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVodSearch{}).Where("`search_key` = ?", searchKey).First(&result).Error

	return
}

// FetchIndexBySearchField  获取多个内容
func (obj *_MacVodSearchMgr) FetchIndexBySearchField(searchField string) (results []*MacVodSearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVodSearch{}).Where("`search_field` = ?", searchField).Find(&results).Error

	return
}

// FetchIndexBySearchHitCount  获取多个内容
func (obj *_MacVodSearchMgr) FetchIndexBySearchHitCount(searchHitCount uint64) (results []*MacVodSearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVodSearch{}).Where("`search_hit_count` = ?", searchHitCount).Find(&results).Error

	return
}

// FetchIndexBySearchLastHitTime  获取多个内容
func (obj *_MacVodSearchMgr) FetchIndexBySearchLastHitTime(searchLastHitTime int) (results []*MacVodSearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVodSearch{}).Where("`search_last_hit_time` = ?", searchLastHitTime).Find(&results).Error

	return
}

// FetchIndexBySearchUpdateTime  获取多个内容
func (obj *_MacVodSearchMgr) FetchIndexBySearchUpdateTime(searchUpdateTime int) (results []*MacVodSearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVodSearch{}).Where("`search_update_time` = ?", searchUpdateTime).Find(&results).Error

	return
}
