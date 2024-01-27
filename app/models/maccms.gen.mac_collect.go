package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _MacCollectMgr struct {
	*_BaseMgr
}

// MacCollectMgr open func
func MacCollectMgr(db *gorm.DB) *_MacCollectMgr {
	if db == nil {
		panic(fmt.Errorf("MacCollectMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MacCollectMgr{_BaseMgr: &_BaseMgr{DB: db.Table("mac_collect"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_MacCollectMgr) Debug() *_MacCollectMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MacCollectMgr) GetTableName() string {
	return "mac_collect"
}

// Reset 重置gorm会话
func (obj *_MacCollectMgr) Reset() *_MacCollectMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MacCollectMgr) Get() (result MacCollect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCollect{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MacCollectMgr) Gets() (results []*MacCollect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCollect{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MacCollectMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MacCollect{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithCollectID collect_id获取 采集ID
func (obj *_MacCollectMgr) WithCollectID(collectID int) Option {
	return optionFunc(func(o *options) { o.query["collect_id"] = collectID })
}

// WithCollectName collect_name获取 采集名称
func (obj *_MacCollectMgr) WithCollectName(collectName string) Option {
	return optionFunc(func(o *options) { o.query["collect_name"] = collectName })
}

// WithCollectURL collect_url获取 采集URL
func (obj *_MacCollectMgr) WithCollectURL(collectURL string) Option {
	return optionFunc(func(o *options) { o.query["collect_url"] = collectURL })
}

// WithCollectType collect_type获取 采集类型
func (obj *_MacCollectMgr) WithCollectType(collectType uint8) Option {
	return optionFunc(func(o *options) { o.query["collect_type"] = collectType })
}

// WithCollectMid collect_mid获取 模型ID
func (obj *_MacCollectMgr) WithCollectMid(collectMid uint8) Option {
	return optionFunc(func(o *options) { o.query["collect_mid"] = collectMid })
}

// WithCollectAppid collect_appid获取 采集AppID
func (obj *_MacCollectMgr) WithCollectAppid(collectAppid string) Option {
	return optionFunc(func(o *options) { o.query["collect_appid"] = collectAppid })
}

// WithCollectAppkey collect_appkey获取 采集AppKey
func (obj *_MacCollectMgr) WithCollectAppkey(collectAppkey string) Option {
	return optionFunc(func(o *options) { o.query["collect_appkey"] = collectAppkey })
}

// WithCollectParam collect_param获取 采集参数
func (obj *_MacCollectMgr) WithCollectParam(collectParam string) Option {
	return optionFunc(func(o *options) { o.query["collect_param"] = collectParam })
}

// WithCollectFilter collect_filter获取 采集过滤
func (obj *_MacCollectMgr) WithCollectFilter(collectFilter uint8) Option {
	return optionFunc(func(o *options) { o.query["collect_filter"] = collectFilter })
}

// WithCollectFilterFrom collect_filter_from获取 采集过滤来源
func (obj *_MacCollectMgr) WithCollectFilterFrom(collectFilterFrom string) Option {
	return optionFunc(func(o *options) { o.query["collect_filter_from"] = collectFilterFrom })
}

// WithCollectFilterYear collect_filter_year获取 采集时，过滤年份
func (obj *_MacCollectMgr) WithCollectFilterYear(collectFilterYear string) Option {
	return optionFunc(func(o *options) { o.query["collect_filter_year"] = collectFilterYear })
}

// WithCollectOpt collect_opt获取 采集选项
func (obj *_MacCollectMgr) WithCollectOpt(collectOpt uint8) Option {
	return optionFunc(func(o *options) { o.query["collect_opt"] = collectOpt })
}

// WithCollectSyncPicOpt collect_sync_pic_opt获取 同步图片选项，0-跟随全局，1-开启，2-关闭
func (obj *_MacCollectMgr) WithCollectSyncPicOpt(collectSyncPicOpt uint8) Option {
	return optionFunc(func(o *options) { o.query["collect_sync_pic_opt"] = collectSyncPicOpt })
}

// GetByOption 功能选项模式获取
func (obj *_MacCollectMgr) GetByOption(opts ...Option) (result MacCollect, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacCollect{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MacCollectMgr) GetByOptions(opts ...Option) (results []*MacCollect, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacCollect{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MacCollectMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MacCollect, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MacCollect{}).Where(options.query)
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

// GetFromCollectID 通过collect_id获取内容 采集ID
func (obj *_MacCollectMgr) GetFromCollectID(collectID int) (result MacCollect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCollect{}).Where("`collect_id` = ?", collectID).First(&result).Error

	return
}

// GetBatchFromCollectID 批量查找 采集ID
func (obj *_MacCollectMgr) GetBatchFromCollectID(collectIDs []int) (results []*MacCollect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCollect{}).Where("`collect_id` IN (?)", collectIDs).Find(&results).Error

	return
}

// GetFromCollectName 通过collect_name获取内容 采集名称
func (obj *_MacCollectMgr) GetFromCollectName(collectName string) (results []*MacCollect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCollect{}).Where("`collect_name` = ?", collectName).Find(&results).Error

	return
}

// GetBatchFromCollectName 批量查找 采集名称
func (obj *_MacCollectMgr) GetBatchFromCollectName(collectNames []string) (results []*MacCollect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCollect{}).Where("`collect_name` IN (?)", collectNames).Find(&results).Error

	return
}

// GetFromCollectURL 通过collect_url获取内容 采集URL
func (obj *_MacCollectMgr) GetFromCollectURL(collectURL string) (results []*MacCollect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCollect{}).Where("`collect_url` = ?", collectURL).Find(&results).Error

	return
}

// GetBatchFromCollectURL 批量查找 采集URL
func (obj *_MacCollectMgr) GetBatchFromCollectURL(collectURLs []string) (results []*MacCollect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCollect{}).Where("`collect_url` IN (?)", collectURLs).Find(&results).Error

	return
}

// GetFromCollectType 通过collect_type获取内容 采集类型
func (obj *_MacCollectMgr) GetFromCollectType(collectType uint8) (results []*MacCollect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCollect{}).Where("`collect_type` = ?", collectType).Find(&results).Error

	return
}

// GetBatchFromCollectType 批量查找 采集类型
func (obj *_MacCollectMgr) GetBatchFromCollectType(collectTypes []uint8) (results []*MacCollect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCollect{}).Where("`collect_type` IN (?)", collectTypes).Find(&results).Error

	return
}

// GetFromCollectMid 通过collect_mid获取内容 模型ID
func (obj *_MacCollectMgr) GetFromCollectMid(collectMid uint8) (results []*MacCollect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCollect{}).Where("`collect_mid` = ?", collectMid).Find(&results).Error

	return
}

// GetBatchFromCollectMid 批量查找 模型ID
func (obj *_MacCollectMgr) GetBatchFromCollectMid(collectMids []uint8) (results []*MacCollect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCollect{}).Where("`collect_mid` IN (?)", collectMids).Find(&results).Error

	return
}

// GetFromCollectAppid 通过collect_appid获取内容 采集AppID
func (obj *_MacCollectMgr) GetFromCollectAppid(collectAppid string) (results []*MacCollect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCollect{}).Where("`collect_appid` = ?", collectAppid).Find(&results).Error

	return
}

// GetBatchFromCollectAppid 批量查找 采集AppID
func (obj *_MacCollectMgr) GetBatchFromCollectAppid(collectAppids []string) (results []*MacCollect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCollect{}).Where("`collect_appid` IN (?)", collectAppids).Find(&results).Error

	return
}

// GetFromCollectAppkey 通过collect_appkey获取内容 采集AppKey
func (obj *_MacCollectMgr) GetFromCollectAppkey(collectAppkey string) (results []*MacCollect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCollect{}).Where("`collect_appkey` = ?", collectAppkey).Find(&results).Error

	return
}

// GetBatchFromCollectAppkey 批量查找 采集AppKey
func (obj *_MacCollectMgr) GetBatchFromCollectAppkey(collectAppkeys []string) (results []*MacCollect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCollect{}).Where("`collect_appkey` IN (?)", collectAppkeys).Find(&results).Error

	return
}

// GetFromCollectParam 通过collect_param获取内容 采集参数
func (obj *_MacCollectMgr) GetFromCollectParam(collectParam string) (results []*MacCollect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCollect{}).Where("`collect_param` = ?", collectParam).Find(&results).Error

	return
}

// GetBatchFromCollectParam 批量查找 采集参数
func (obj *_MacCollectMgr) GetBatchFromCollectParam(collectParams []string) (results []*MacCollect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCollect{}).Where("`collect_param` IN (?)", collectParams).Find(&results).Error

	return
}

// GetFromCollectFilter 通过collect_filter获取内容 采集过滤
func (obj *_MacCollectMgr) GetFromCollectFilter(collectFilter uint8) (results []*MacCollect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCollect{}).Where("`collect_filter` = ?", collectFilter).Find(&results).Error

	return
}

// GetBatchFromCollectFilter 批量查找 采集过滤
func (obj *_MacCollectMgr) GetBatchFromCollectFilter(collectFilters []uint8) (results []*MacCollect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCollect{}).Where("`collect_filter` IN (?)", collectFilters).Find(&results).Error

	return
}

// GetFromCollectFilterFrom 通过collect_filter_from获取内容 采集过滤来源
func (obj *_MacCollectMgr) GetFromCollectFilterFrom(collectFilterFrom string) (results []*MacCollect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCollect{}).Where("`collect_filter_from` = ?", collectFilterFrom).Find(&results).Error

	return
}

// GetBatchFromCollectFilterFrom 批量查找 采集过滤来源
func (obj *_MacCollectMgr) GetBatchFromCollectFilterFrom(collectFilterFroms []string) (results []*MacCollect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCollect{}).Where("`collect_filter_from` IN (?)", collectFilterFroms).Find(&results).Error

	return
}

// GetFromCollectFilterYear 通过collect_filter_year获取内容 采集时，过滤年份
func (obj *_MacCollectMgr) GetFromCollectFilterYear(collectFilterYear string) (results []*MacCollect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCollect{}).Where("`collect_filter_year` = ?", collectFilterYear).Find(&results).Error

	return
}

// GetBatchFromCollectFilterYear 批量查找 采集时，过滤年份
func (obj *_MacCollectMgr) GetBatchFromCollectFilterYear(collectFilterYears []string) (results []*MacCollect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCollect{}).Where("`collect_filter_year` IN (?)", collectFilterYears).Find(&results).Error

	return
}

// GetFromCollectOpt 通过collect_opt获取内容 采集选项
func (obj *_MacCollectMgr) GetFromCollectOpt(collectOpt uint8) (results []*MacCollect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCollect{}).Where("`collect_opt` = ?", collectOpt).Find(&results).Error

	return
}

// GetBatchFromCollectOpt 批量查找 采集选项
func (obj *_MacCollectMgr) GetBatchFromCollectOpt(collectOpts []uint8) (results []*MacCollect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCollect{}).Where("`collect_opt` IN (?)", collectOpts).Find(&results).Error

	return
}

// GetFromCollectSyncPicOpt 通过collect_sync_pic_opt获取内容 同步图片选项，0-跟随全局，1-开启，2-关闭
func (obj *_MacCollectMgr) GetFromCollectSyncPicOpt(collectSyncPicOpt uint8) (results []*MacCollect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCollect{}).Where("`collect_sync_pic_opt` = ?", collectSyncPicOpt).Find(&results).Error

	return
}

// GetBatchFromCollectSyncPicOpt 批量查找 同步图片选项，0-跟随全局，1-开启，2-关闭
func (obj *_MacCollectMgr) GetBatchFromCollectSyncPicOpt(collectSyncPicOpts []uint8) (results []*MacCollect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCollect{}).Where("`collect_sync_pic_opt` IN (?)", collectSyncPicOpts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MacCollectMgr) FetchByPrimaryKey(collectID int) (result MacCollect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCollect{}).Where("`collect_id` = ?", collectID).First(&result).Error

	return
}
