package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _MacWebsiteMgr struct {
	*_BaseMgr
}

// MacWebsiteMgr open func
func MacWebsiteMgr(db *gorm.DB) *_MacWebsiteMgr {
	if db == nil {
		panic(fmt.Errorf("MacWebsiteMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MacWebsiteMgr{_BaseMgr: &_BaseMgr{DB: db.Table("mac_website"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_MacWebsiteMgr) Debug() *_MacWebsiteMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MacWebsiteMgr) GetTableName() string {
	return "mac_website"
}

// Reset 重置gorm会话
func (obj *_MacWebsiteMgr) Reset() *_MacWebsiteMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MacWebsiteMgr) Get() (result MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MacWebsiteMgr) Gets() (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MacWebsiteMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithWebsiteID website_id获取 网站ID
func (obj *_MacWebsiteMgr) WithWebsiteID(websiteID int) Option {
	return optionFunc(func(o *options) { o.query["website_id"] = websiteID })
}

// WithTypeID type_id获取 类型ID
func (obj *_MacWebsiteMgr) WithTypeID(typeID uint16) Option {
	return optionFunc(func(o *options) { o.query["type_id"] = typeID })
}

// WithTypeID1 type_id_1获取 类型ID_1
func (obj *_MacWebsiteMgr) WithTypeID1(typeID1 uint16) Option {
	return optionFunc(func(o *options) { o.query["type_id_1"] = typeID1 })
}

// WithWebsiteName website_name获取 网站名称
func (obj *_MacWebsiteMgr) WithWebsiteName(websiteName string) Option {
	return optionFunc(func(o *options) { o.query["website_name"] = websiteName })
}

// WithWebsiteSub website_sub获取 网站子标题
func (obj *_MacWebsiteMgr) WithWebsiteSub(websiteSub string) Option {
	return optionFunc(func(o *options) { o.query["website_sub"] = websiteSub })
}

// WithWebsiteEn website_en获取 网站英文名称
func (obj *_MacWebsiteMgr) WithWebsiteEn(websiteEn string) Option {
	return optionFunc(func(o *options) { o.query["website_en"] = websiteEn })
}

// WithWebsiteStatus website_status获取 网站状态
func (obj *_MacWebsiteMgr) WithWebsiteStatus(websiteStatus uint8) Option {
	return optionFunc(func(o *options) { o.query["website_status"] = websiteStatus })
}

// WithWebsiteLetter website_letter获取 网站首字母
func (obj *_MacWebsiteMgr) WithWebsiteLetter(websiteLetter string) Option {
	return optionFunc(func(o *options) { o.query["website_letter"] = websiteLetter })
}

// WithWebsiteColor website_color获取 网站颜色
func (obj *_MacWebsiteMgr) WithWebsiteColor(websiteColor string) Option {
	return optionFunc(func(o *options) { o.query["website_color"] = websiteColor })
}

// WithWebsiteLock website_lock获取 网站锁定状态
func (obj *_MacWebsiteMgr) WithWebsiteLock(websiteLock uint8) Option {
	return optionFunc(func(o *options) { o.query["website_lock"] = websiteLock })
}

// WithWebsiteSort website_sort获取 网站排序
func (obj *_MacWebsiteMgr) WithWebsiteSort(websiteSort int) Option {
	return optionFunc(func(o *options) { o.query["website_sort"] = websiteSort })
}

// WithWebsiteJumpurl website_jumpurl获取 网站跳转URL
func (obj *_MacWebsiteMgr) WithWebsiteJumpurl(websiteJumpurl string) Option {
	return optionFunc(func(o *options) { o.query["website_jumpurl"] = websiteJumpurl })
}

// WithWebsitePic website_pic获取 网站图片
func (obj *_MacWebsiteMgr) WithWebsitePic(websitePic string) Option {
	return optionFunc(func(o *options) { o.query["website_pic"] = websitePic })
}

// WithWebsitePicScreenshot website_pic_screenshot获取 网站图片截图
func (obj *_MacWebsiteMgr) WithWebsitePicScreenshot(websitePicScreenshot string) Option {
	return optionFunc(func(o *options) { o.query["website_pic_screenshot"] = websitePicScreenshot })
}

// WithWebsiteLogo website_logo获取 网站Logo
func (obj *_MacWebsiteMgr) WithWebsiteLogo(websiteLogo string) Option {
	return optionFunc(func(o *options) { o.query["website_logo"] = websiteLogo })
}

// WithWebsiteArea website_area获取 网站地区
func (obj *_MacWebsiteMgr) WithWebsiteArea(websiteArea string) Option {
	return optionFunc(func(o *options) { o.query["website_area"] = websiteArea })
}

// WithWebsiteLang website_lang获取 网站语言
func (obj *_MacWebsiteMgr) WithWebsiteLang(websiteLang string) Option {
	return optionFunc(func(o *options) { o.query["website_lang"] = websiteLang })
}

// WithWebsiteLevel website_level获取 网站级别
func (obj *_MacWebsiteMgr) WithWebsiteLevel(websiteLevel uint8) Option {
	return optionFunc(func(o *options) { o.query["website_level"] = websiteLevel })
}

// WithWebsiteTime website_time获取 网站时间
func (obj *_MacWebsiteMgr) WithWebsiteTime(websiteTime int) Option {
	return optionFunc(func(o *options) { o.query["website_time"] = websiteTime })
}

// WithWebsiteTimeAdd website_time_add获取 网站添加时间
func (obj *_MacWebsiteMgr) WithWebsiteTimeAdd(websiteTimeAdd int) Option {
	return optionFunc(func(o *options) { o.query["website_time_add"] = websiteTimeAdd })
}

// WithWebsiteTimeHits website_time_hits获取 网站点击时间
func (obj *_MacWebsiteMgr) WithWebsiteTimeHits(websiteTimeHits int) Option {
	return optionFunc(func(o *options) { o.query["website_time_hits"] = websiteTimeHits })
}

// WithWebsiteTimeMake website_time_make获取 网站制作时间
func (obj *_MacWebsiteMgr) WithWebsiteTimeMake(websiteTimeMake int) Option {
	return optionFunc(func(o *options) { o.query["website_time_make"] = websiteTimeMake })
}

// WithWebsiteTimeReferer website_time_referer获取 网站引用时间
func (obj *_MacWebsiteMgr) WithWebsiteTimeReferer(websiteTimeReferer int) Option {
	return optionFunc(func(o *options) { o.query["website_time_referer"] = websiteTimeReferer })
}

// WithWebsiteHits website_hits获取 网站点击次数
func (obj *_MacWebsiteMgr) WithWebsiteHits(websiteHits int16) Option {
	return optionFunc(func(o *options) { o.query["website_hits"] = websiteHits })
}

// WithWebsiteHitsDay website_hits_day获取 网站日点击次数
func (obj *_MacWebsiteMgr) WithWebsiteHitsDay(websiteHitsDay int16) Option {
	return optionFunc(func(o *options) { o.query["website_hits_day"] = websiteHitsDay })
}

// WithWebsiteHitsWeek website_hits_week获取 网站周点击次数
func (obj *_MacWebsiteMgr) WithWebsiteHitsWeek(websiteHitsWeek int16) Option {
	return optionFunc(func(o *options) { o.query["website_hits_week"] = websiteHitsWeek })
}

// WithWebsiteHitsMonth website_hits_month获取 网站月点击次数
func (obj *_MacWebsiteMgr) WithWebsiteHitsMonth(websiteHitsMonth int16) Option {
	return optionFunc(func(o *options) { o.query["website_hits_month"] = websiteHitsMonth })
}

// WithWebsiteScore website_score获取 网站评分
func (obj *_MacWebsiteMgr) WithWebsiteScore(websiteScore float64) Option {
	return optionFunc(func(o *options) { o.query["website_score"] = websiteScore })
}

// WithWebsiteScoreAll website_score_all获取 网站总评分
func (obj *_MacWebsiteMgr) WithWebsiteScoreAll(websiteScoreAll int16) Option {
	return optionFunc(func(o *options) { o.query["website_score_all"] = websiteScoreAll })
}

// WithWebsiteScoreNum website_score_num获取 网站评分次数
func (obj *_MacWebsiteMgr) WithWebsiteScoreNum(websiteScoreNum int16) Option {
	return optionFunc(func(o *options) { o.query["website_score_num"] = websiteScoreNum })
}

// WithWebsiteUp website_up获取 网站点赞数
func (obj *_MacWebsiteMgr) WithWebsiteUp(websiteUp int16) Option {
	return optionFunc(func(o *options) { o.query["website_up"] = websiteUp })
}

// WithWebsiteDown website_down获取 网站踩数
func (obj *_MacWebsiteMgr) WithWebsiteDown(websiteDown int16) Option {
	return optionFunc(func(o *options) { o.query["website_down"] = websiteDown })
}

// WithWebsiteReferer website_referer获取 网站引用数
func (obj *_MacWebsiteMgr) WithWebsiteReferer(websiteReferer int16) Option {
	return optionFunc(func(o *options) { o.query["website_referer"] = websiteReferer })
}

// WithWebsiteRefererDay website_referer_day获取 网站日引用数
func (obj *_MacWebsiteMgr) WithWebsiteRefererDay(websiteRefererDay int16) Option {
	return optionFunc(func(o *options) { o.query["website_referer_day"] = websiteRefererDay })
}

// WithWebsiteRefererWeek website_referer_week获取 网站周引用数
func (obj *_MacWebsiteMgr) WithWebsiteRefererWeek(websiteRefererWeek int16) Option {
	return optionFunc(func(o *options) { o.query["website_referer_week"] = websiteRefererWeek })
}

// WithWebsiteRefererMonth website_referer_month获取 网站月引用数
func (obj *_MacWebsiteMgr) WithWebsiteRefererMonth(websiteRefererMonth int16) Option {
	return optionFunc(func(o *options) { o.query["website_referer_month"] = websiteRefererMonth })
}

// WithWebsiteTag website_tag获取 网站标签
func (obj *_MacWebsiteMgr) WithWebsiteTag(websiteTag string) Option {
	return optionFunc(func(o *options) { o.query["website_tag"] = websiteTag })
}

// WithWebsiteClass website_class获取 网站类别
func (obj *_MacWebsiteMgr) WithWebsiteClass(websiteClass string) Option {
	return optionFunc(func(o *options) { o.query["website_class"] = websiteClass })
}

// WithWebsiteRemarks website_remarks获取 网站备注
func (obj *_MacWebsiteMgr) WithWebsiteRemarks(websiteRemarks string) Option {
	return optionFunc(func(o *options) { o.query["website_remarks"] = websiteRemarks })
}

// WithWebsiteTpl website_tpl获取 网站模板
func (obj *_MacWebsiteMgr) WithWebsiteTpl(websiteTpl string) Option {
	return optionFunc(func(o *options) { o.query["website_tpl"] = websiteTpl })
}

// WithWebsiteBlurb website_blurb获取 网站简介
func (obj *_MacWebsiteMgr) WithWebsiteBlurb(websiteBlurb string) Option {
	return optionFunc(func(o *options) { o.query["website_blurb"] = websiteBlurb })
}

// WithWebsiteContent website_content获取 网站内容
func (obj *_MacWebsiteMgr) WithWebsiteContent(websiteContent string) Option {
	return optionFunc(func(o *options) { o.query["website_content"] = websiteContent })
}

// GetByOption 功能选项模式获取
func (obj *_MacWebsiteMgr) GetByOption(opts ...Option) (result MacWebsite, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MacWebsiteMgr) GetByOptions(opts ...Option) (results []*MacWebsite, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MacWebsiteMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MacWebsite, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where(options.query)
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

// GetFromWebsiteID 通过website_id获取内容 网站ID
func (obj *_MacWebsiteMgr) GetFromWebsiteID(websiteID int) (result MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_id` = ?", websiteID).First(&result).Error

	return
}

// GetBatchFromWebsiteID 批量查找 网站ID
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteID(websiteIDs []int) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_id` IN (?)", websiteIDs).Find(&results).Error

	return
}

// GetFromTypeID 通过type_id获取内容 类型ID
func (obj *_MacWebsiteMgr) GetFromTypeID(typeID uint16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`type_id` = ?", typeID).Find(&results).Error

	return
}

// GetBatchFromTypeID 批量查找 类型ID
func (obj *_MacWebsiteMgr) GetBatchFromTypeID(typeIDs []uint16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`type_id` IN (?)", typeIDs).Find(&results).Error

	return
}

// GetFromTypeID1 通过type_id_1获取内容 类型ID_1
func (obj *_MacWebsiteMgr) GetFromTypeID1(typeID1 uint16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`type_id_1` = ?", typeID1).Find(&results).Error

	return
}

// GetBatchFromTypeID1 批量查找 类型ID_1
func (obj *_MacWebsiteMgr) GetBatchFromTypeID1(typeID1s []uint16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`type_id_1` IN (?)", typeID1s).Find(&results).Error

	return
}

// GetFromWebsiteName 通过website_name获取内容 网站名称
func (obj *_MacWebsiteMgr) GetFromWebsiteName(websiteName string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_name` = ?", websiteName).Find(&results).Error

	return
}

// GetBatchFromWebsiteName 批量查找 网站名称
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteName(websiteNames []string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_name` IN (?)", websiteNames).Find(&results).Error

	return
}

// GetFromWebsiteSub 通过website_sub获取内容 网站子标题
func (obj *_MacWebsiteMgr) GetFromWebsiteSub(websiteSub string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_sub` = ?", websiteSub).Find(&results).Error

	return
}

// GetBatchFromWebsiteSub 批量查找 网站子标题
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteSub(websiteSubs []string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_sub` IN (?)", websiteSubs).Find(&results).Error

	return
}

// GetFromWebsiteEn 通过website_en获取内容 网站英文名称
func (obj *_MacWebsiteMgr) GetFromWebsiteEn(websiteEn string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_en` = ?", websiteEn).Find(&results).Error

	return
}

// GetBatchFromWebsiteEn 批量查找 网站英文名称
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteEn(websiteEns []string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_en` IN (?)", websiteEns).Find(&results).Error

	return
}

// GetFromWebsiteStatus 通过website_status获取内容 网站状态
func (obj *_MacWebsiteMgr) GetFromWebsiteStatus(websiteStatus uint8) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_status` = ?", websiteStatus).Find(&results).Error

	return
}

// GetBatchFromWebsiteStatus 批量查找 网站状态
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteStatus(websiteStatuss []uint8) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_status` IN (?)", websiteStatuss).Find(&results).Error

	return
}

// GetFromWebsiteLetter 通过website_letter获取内容 网站首字母
func (obj *_MacWebsiteMgr) GetFromWebsiteLetter(websiteLetter string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_letter` = ?", websiteLetter).Find(&results).Error

	return
}

// GetBatchFromWebsiteLetter 批量查找 网站首字母
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteLetter(websiteLetters []string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_letter` IN (?)", websiteLetters).Find(&results).Error

	return
}

// GetFromWebsiteColor 通过website_color获取内容 网站颜色
func (obj *_MacWebsiteMgr) GetFromWebsiteColor(websiteColor string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_color` = ?", websiteColor).Find(&results).Error

	return
}

// GetBatchFromWebsiteColor 批量查找 网站颜色
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteColor(websiteColors []string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_color` IN (?)", websiteColors).Find(&results).Error

	return
}

// GetFromWebsiteLock 通过website_lock获取内容 网站锁定状态
func (obj *_MacWebsiteMgr) GetFromWebsiteLock(websiteLock uint8) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_lock` = ?", websiteLock).Find(&results).Error

	return
}

// GetBatchFromWebsiteLock 批量查找 网站锁定状态
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteLock(websiteLocks []uint8) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_lock` IN (?)", websiteLocks).Find(&results).Error

	return
}

// GetFromWebsiteSort 通过website_sort获取内容 网站排序
func (obj *_MacWebsiteMgr) GetFromWebsiteSort(websiteSort int) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_sort` = ?", websiteSort).Find(&results).Error

	return
}

// GetBatchFromWebsiteSort 批量查找 网站排序
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteSort(websiteSorts []int) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_sort` IN (?)", websiteSorts).Find(&results).Error

	return
}

// GetFromWebsiteJumpurl 通过website_jumpurl获取内容 网站跳转URL
func (obj *_MacWebsiteMgr) GetFromWebsiteJumpurl(websiteJumpurl string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_jumpurl` = ?", websiteJumpurl).Find(&results).Error

	return
}

// GetBatchFromWebsiteJumpurl 批量查找 网站跳转URL
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteJumpurl(websiteJumpurls []string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_jumpurl` IN (?)", websiteJumpurls).Find(&results).Error

	return
}

// GetFromWebsitePic 通过website_pic获取内容 网站图片
func (obj *_MacWebsiteMgr) GetFromWebsitePic(websitePic string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_pic` = ?", websitePic).Find(&results).Error

	return
}

// GetBatchFromWebsitePic 批量查找 网站图片
func (obj *_MacWebsiteMgr) GetBatchFromWebsitePic(websitePics []string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_pic` IN (?)", websitePics).Find(&results).Error

	return
}

// GetFromWebsitePicScreenshot 通过website_pic_screenshot获取内容 网站图片截图
func (obj *_MacWebsiteMgr) GetFromWebsitePicScreenshot(websitePicScreenshot string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_pic_screenshot` = ?", websitePicScreenshot).Find(&results).Error

	return
}

// GetBatchFromWebsitePicScreenshot 批量查找 网站图片截图
func (obj *_MacWebsiteMgr) GetBatchFromWebsitePicScreenshot(websitePicScreenshots []string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_pic_screenshot` IN (?)", websitePicScreenshots).Find(&results).Error

	return
}

// GetFromWebsiteLogo 通过website_logo获取内容 网站Logo
func (obj *_MacWebsiteMgr) GetFromWebsiteLogo(websiteLogo string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_logo` = ?", websiteLogo).Find(&results).Error

	return
}

// GetBatchFromWebsiteLogo 批量查找 网站Logo
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteLogo(websiteLogos []string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_logo` IN (?)", websiteLogos).Find(&results).Error

	return
}

// GetFromWebsiteArea 通过website_area获取内容 网站地区
func (obj *_MacWebsiteMgr) GetFromWebsiteArea(websiteArea string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_area` = ?", websiteArea).Find(&results).Error

	return
}

// GetBatchFromWebsiteArea 批量查找 网站地区
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteArea(websiteAreas []string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_area` IN (?)", websiteAreas).Find(&results).Error

	return
}

// GetFromWebsiteLang 通过website_lang获取内容 网站语言
func (obj *_MacWebsiteMgr) GetFromWebsiteLang(websiteLang string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_lang` = ?", websiteLang).Find(&results).Error

	return
}

// GetBatchFromWebsiteLang 批量查找 网站语言
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteLang(websiteLangs []string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_lang` IN (?)", websiteLangs).Find(&results).Error

	return
}

// GetFromWebsiteLevel 通过website_level获取内容 网站级别
func (obj *_MacWebsiteMgr) GetFromWebsiteLevel(websiteLevel uint8) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_level` = ?", websiteLevel).Find(&results).Error

	return
}

// GetBatchFromWebsiteLevel 批量查找 网站级别
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteLevel(websiteLevels []uint8) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_level` IN (?)", websiteLevels).Find(&results).Error

	return
}

// GetFromWebsiteTime 通过website_time获取内容 网站时间
func (obj *_MacWebsiteMgr) GetFromWebsiteTime(websiteTime int) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_time` = ?", websiteTime).Find(&results).Error

	return
}

// GetBatchFromWebsiteTime 批量查找 网站时间
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteTime(websiteTimes []int) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_time` IN (?)", websiteTimes).Find(&results).Error

	return
}

// GetFromWebsiteTimeAdd 通过website_time_add获取内容 网站添加时间
func (obj *_MacWebsiteMgr) GetFromWebsiteTimeAdd(websiteTimeAdd int) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_time_add` = ?", websiteTimeAdd).Find(&results).Error

	return
}

// GetBatchFromWebsiteTimeAdd 批量查找 网站添加时间
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteTimeAdd(websiteTimeAdds []int) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_time_add` IN (?)", websiteTimeAdds).Find(&results).Error

	return
}

// GetFromWebsiteTimeHits 通过website_time_hits获取内容 网站点击时间
func (obj *_MacWebsiteMgr) GetFromWebsiteTimeHits(websiteTimeHits int) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_time_hits` = ?", websiteTimeHits).Find(&results).Error

	return
}

// GetBatchFromWebsiteTimeHits 批量查找 网站点击时间
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteTimeHits(websiteTimeHitss []int) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_time_hits` IN (?)", websiteTimeHitss).Find(&results).Error

	return
}

// GetFromWebsiteTimeMake 通过website_time_make获取内容 网站制作时间
func (obj *_MacWebsiteMgr) GetFromWebsiteTimeMake(websiteTimeMake int) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_time_make` = ?", websiteTimeMake).Find(&results).Error

	return
}

// GetBatchFromWebsiteTimeMake 批量查找 网站制作时间
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteTimeMake(websiteTimeMakes []int) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_time_make` IN (?)", websiteTimeMakes).Find(&results).Error

	return
}

// GetFromWebsiteTimeReferer 通过website_time_referer获取内容 网站引用时间
func (obj *_MacWebsiteMgr) GetFromWebsiteTimeReferer(websiteTimeReferer int) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_time_referer` = ?", websiteTimeReferer).Find(&results).Error

	return
}

// GetBatchFromWebsiteTimeReferer 批量查找 网站引用时间
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteTimeReferer(websiteTimeReferers []int) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_time_referer` IN (?)", websiteTimeReferers).Find(&results).Error

	return
}

// GetFromWebsiteHits 通过website_hits获取内容 网站点击次数
func (obj *_MacWebsiteMgr) GetFromWebsiteHits(websiteHits int16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_hits` = ?", websiteHits).Find(&results).Error

	return
}

// GetBatchFromWebsiteHits 批量查找 网站点击次数
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteHits(websiteHitss []int16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_hits` IN (?)", websiteHitss).Find(&results).Error

	return
}

// GetFromWebsiteHitsDay 通过website_hits_day获取内容 网站日点击次数
func (obj *_MacWebsiteMgr) GetFromWebsiteHitsDay(websiteHitsDay int16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_hits_day` = ?", websiteHitsDay).Find(&results).Error

	return
}

// GetBatchFromWebsiteHitsDay 批量查找 网站日点击次数
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteHitsDay(websiteHitsDays []int16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_hits_day` IN (?)", websiteHitsDays).Find(&results).Error

	return
}

// GetFromWebsiteHitsWeek 通过website_hits_week获取内容 网站周点击次数
func (obj *_MacWebsiteMgr) GetFromWebsiteHitsWeek(websiteHitsWeek int16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_hits_week` = ?", websiteHitsWeek).Find(&results).Error

	return
}

// GetBatchFromWebsiteHitsWeek 批量查找 网站周点击次数
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteHitsWeek(websiteHitsWeeks []int16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_hits_week` IN (?)", websiteHitsWeeks).Find(&results).Error

	return
}

// GetFromWebsiteHitsMonth 通过website_hits_month获取内容 网站月点击次数
func (obj *_MacWebsiteMgr) GetFromWebsiteHitsMonth(websiteHitsMonth int16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_hits_month` = ?", websiteHitsMonth).Find(&results).Error

	return
}

// GetBatchFromWebsiteHitsMonth 批量查找 网站月点击次数
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteHitsMonth(websiteHitsMonths []int16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_hits_month` IN (?)", websiteHitsMonths).Find(&results).Error

	return
}

// GetFromWebsiteScore 通过website_score获取内容 网站评分
func (obj *_MacWebsiteMgr) GetFromWebsiteScore(websiteScore float64) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_score` = ?", websiteScore).Find(&results).Error

	return
}

// GetBatchFromWebsiteScore 批量查找 网站评分
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteScore(websiteScores []float64) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_score` IN (?)", websiteScores).Find(&results).Error

	return
}

// GetFromWebsiteScoreAll 通过website_score_all获取内容 网站总评分
func (obj *_MacWebsiteMgr) GetFromWebsiteScoreAll(websiteScoreAll int16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_score_all` = ?", websiteScoreAll).Find(&results).Error

	return
}

// GetBatchFromWebsiteScoreAll 批量查找 网站总评分
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteScoreAll(websiteScoreAlls []int16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_score_all` IN (?)", websiteScoreAlls).Find(&results).Error

	return
}

// GetFromWebsiteScoreNum 通过website_score_num获取内容 网站评分次数
func (obj *_MacWebsiteMgr) GetFromWebsiteScoreNum(websiteScoreNum int16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_score_num` = ?", websiteScoreNum).Find(&results).Error

	return
}

// GetBatchFromWebsiteScoreNum 批量查找 网站评分次数
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteScoreNum(websiteScoreNums []int16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_score_num` IN (?)", websiteScoreNums).Find(&results).Error

	return
}

// GetFromWebsiteUp 通过website_up获取内容 网站点赞数
func (obj *_MacWebsiteMgr) GetFromWebsiteUp(websiteUp int16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_up` = ?", websiteUp).Find(&results).Error

	return
}

// GetBatchFromWebsiteUp 批量查找 网站点赞数
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteUp(websiteUps []int16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_up` IN (?)", websiteUps).Find(&results).Error

	return
}

// GetFromWebsiteDown 通过website_down获取内容 网站踩数
func (obj *_MacWebsiteMgr) GetFromWebsiteDown(websiteDown int16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_down` = ?", websiteDown).Find(&results).Error

	return
}

// GetBatchFromWebsiteDown 批量查找 网站踩数
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteDown(websiteDowns []int16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_down` IN (?)", websiteDowns).Find(&results).Error

	return
}

// GetFromWebsiteReferer 通过website_referer获取内容 网站引用数
func (obj *_MacWebsiteMgr) GetFromWebsiteReferer(websiteReferer int16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_referer` = ?", websiteReferer).Find(&results).Error

	return
}

// GetBatchFromWebsiteReferer 批量查找 网站引用数
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteReferer(websiteReferers []int16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_referer` IN (?)", websiteReferers).Find(&results).Error

	return
}

// GetFromWebsiteRefererDay 通过website_referer_day获取内容 网站日引用数
func (obj *_MacWebsiteMgr) GetFromWebsiteRefererDay(websiteRefererDay int16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_referer_day` = ?", websiteRefererDay).Find(&results).Error

	return
}

// GetBatchFromWebsiteRefererDay 批量查找 网站日引用数
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteRefererDay(websiteRefererDays []int16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_referer_day` IN (?)", websiteRefererDays).Find(&results).Error

	return
}

// GetFromWebsiteRefererWeek 通过website_referer_week获取内容 网站周引用数
func (obj *_MacWebsiteMgr) GetFromWebsiteRefererWeek(websiteRefererWeek int16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_referer_week` = ?", websiteRefererWeek).Find(&results).Error

	return
}

// GetBatchFromWebsiteRefererWeek 批量查找 网站周引用数
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteRefererWeek(websiteRefererWeeks []int16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_referer_week` IN (?)", websiteRefererWeeks).Find(&results).Error

	return
}

// GetFromWebsiteRefererMonth 通过website_referer_month获取内容 网站月引用数
func (obj *_MacWebsiteMgr) GetFromWebsiteRefererMonth(websiteRefererMonth int16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_referer_month` = ?", websiteRefererMonth).Find(&results).Error

	return
}

// GetBatchFromWebsiteRefererMonth 批量查找 网站月引用数
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteRefererMonth(websiteRefererMonths []int16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_referer_month` IN (?)", websiteRefererMonths).Find(&results).Error

	return
}

// GetFromWebsiteTag 通过website_tag获取内容 网站标签
func (obj *_MacWebsiteMgr) GetFromWebsiteTag(websiteTag string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_tag` = ?", websiteTag).Find(&results).Error

	return
}

// GetBatchFromWebsiteTag 批量查找 网站标签
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteTag(websiteTags []string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_tag` IN (?)", websiteTags).Find(&results).Error

	return
}

// GetFromWebsiteClass 通过website_class获取内容 网站类别
func (obj *_MacWebsiteMgr) GetFromWebsiteClass(websiteClass string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_class` = ?", websiteClass).Find(&results).Error

	return
}

// GetBatchFromWebsiteClass 批量查找 网站类别
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteClass(websiteClasss []string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_class` IN (?)", websiteClasss).Find(&results).Error

	return
}

// GetFromWebsiteRemarks 通过website_remarks获取内容 网站备注
func (obj *_MacWebsiteMgr) GetFromWebsiteRemarks(websiteRemarks string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_remarks` = ?", websiteRemarks).Find(&results).Error

	return
}

// GetBatchFromWebsiteRemarks 批量查找 网站备注
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteRemarks(websiteRemarkss []string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_remarks` IN (?)", websiteRemarkss).Find(&results).Error

	return
}

// GetFromWebsiteTpl 通过website_tpl获取内容 网站模板
func (obj *_MacWebsiteMgr) GetFromWebsiteTpl(websiteTpl string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_tpl` = ?", websiteTpl).Find(&results).Error

	return
}

// GetBatchFromWebsiteTpl 批量查找 网站模板
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteTpl(websiteTpls []string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_tpl` IN (?)", websiteTpls).Find(&results).Error

	return
}

// GetFromWebsiteBlurb 通过website_blurb获取内容 网站简介
func (obj *_MacWebsiteMgr) GetFromWebsiteBlurb(websiteBlurb string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_blurb` = ?", websiteBlurb).Find(&results).Error

	return
}

// GetBatchFromWebsiteBlurb 批量查找 网站简介
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteBlurb(websiteBlurbs []string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_blurb` IN (?)", websiteBlurbs).Find(&results).Error

	return
}

// GetFromWebsiteContent 通过website_content获取内容 网站内容
func (obj *_MacWebsiteMgr) GetFromWebsiteContent(websiteContent string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_content` = ?", websiteContent).Find(&results).Error

	return
}

// GetBatchFromWebsiteContent 批量查找 网站内容
func (obj *_MacWebsiteMgr) GetBatchFromWebsiteContent(websiteContents []string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_content` IN (?)", websiteContents).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MacWebsiteMgr) FetchByPrimaryKey(websiteID int) (result MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_id` = ?", websiteID).First(&result).Error

	return
}

// FetchIndexByTypeID  获取多个内容
func (obj *_MacWebsiteMgr) FetchIndexByTypeID(typeID uint16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`type_id` = ?", typeID).Find(&results).Error

	return
}

// FetchIndexByTypeID1  获取多个内容
func (obj *_MacWebsiteMgr) FetchIndexByTypeID1(typeID1 uint16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`type_id_1` = ?", typeID1).Find(&results).Error

	return
}

// FetchIndexByWebsiteName  获取多个内容
func (obj *_MacWebsiteMgr) FetchIndexByWebsiteName(websiteName string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_name` = ?", websiteName).Find(&results).Error

	return
}

// FetchIndexByWebsiteEn  获取多个内容
func (obj *_MacWebsiteMgr) FetchIndexByWebsiteEn(websiteEn string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_en` = ?", websiteEn).Find(&results).Error

	return
}

// FetchIndexByWebsiteLetter  获取多个内容
func (obj *_MacWebsiteMgr) FetchIndexByWebsiteLetter(websiteLetter string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_letter` = ?", websiteLetter).Find(&results).Error

	return
}

// FetchIndexByWebsiteLock  获取多个内容
func (obj *_MacWebsiteMgr) FetchIndexByWebsiteLock(websiteLock uint8) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_lock` = ?", websiteLock).Find(&results).Error

	return
}

// FetchIndexByWebsiteSort  获取多个内容
func (obj *_MacWebsiteMgr) FetchIndexByWebsiteSort(websiteSort int) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_sort` = ?", websiteSort).Find(&results).Error

	return
}

// FetchIndexByWebsiteLevel  获取多个内容
func (obj *_MacWebsiteMgr) FetchIndexByWebsiteLevel(websiteLevel uint8) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_level` = ?", websiteLevel).Find(&results).Error

	return
}

// FetchIndexByWebsiteTime  获取多个内容
func (obj *_MacWebsiteMgr) FetchIndexByWebsiteTime(websiteTime int) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_time` = ?", websiteTime).Find(&results).Error

	return
}

// FetchIndexByWebsiteTimeAdd  获取多个内容
func (obj *_MacWebsiteMgr) FetchIndexByWebsiteTimeAdd(websiteTimeAdd int) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_time_add` = ?", websiteTimeAdd).Find(&results).Error

	return
}

// FetchIndexByWebsiteTimeMake  获取多个内容
func (obj *_MacWebsiteMgr) FetchIndexByWebsiteTimeMake(websiteTimeMake int) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_time_make` = ?", websiteTimeMake).Find(&results).Error

	return
}

// FetchIndexByWebsiteTimeReferer  获取多个内容
func (obj *_MacWebsiteMgr) FetchIndexByWebsiteTimeReferer(websiteTimeReferer int) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_time_referer` = ?", websiteTimeReferer).Find(&results).Error

	return
}

// FetchIndexByWebsiteHits  获取多个内容
func (obj *_MacWebsiteMgr) FetchIndexByWebsiteHits(websiteHits int16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_hits` = ?", websiteHits).Find(&results).Error

	return
}

// FetchIndexByWebsiteHitsDay  获取多个内容
func (obj *_MacWebsiteMgr) FetchIndexByWebsiteHitsDay(websiteHitsDay int16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_hits_day` = ?", websiteHitsDay).Find(&results).Error

	return
}

// FetchIndexByWebsiteHitsWeek  获取多个内容
func (obj *_MacWebsiteMgr) FetchIndexByWebsiteHitsWeek(websiteHitsWeek int16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_hits_week` = ?", websiteHitsWeek).Find(&results).Error

	return
}

// FetchIndexByWebsiteHitsMonth  获取多个内容
func (obj *_MacWebsiteMgr) FetchIndexByWebsiteHitsMonth(websiteHitsMonth int16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_hits_month` = ?", websiteHitsMonth).Find(&results).Error

	return
}

// FetchIndexByWebsiteScore  获取多个内容
func (obj *_MacWebsiteMgr) FetchIndexByWebsiteScore(websiteScore float64) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_score` = ?", websiteScore).Find(&results).Error

	return
}

// FetchIndexByWebsiteScoreAll  获取多个内容
func (obj *_MacWebsiteMgr) FetchIndexByWebsiteScoreAll(websiteScoreAll int16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_score_all` = ?", websiteScoreAll).Find(&results).Error

	return
}

// FetchIndexByWebsiteScoreNum  获取多个内容
func (obj *_MacWebsiteMgr) FetchIndexByWebsiteScoreNum(websiteScoreNum int16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_score_num` = ?", websiteScoreNum).Find(&results).Error

	return
}

// FetchIndexByWebsiteUp  获取多个内容
func (obj *_MacWebsiteMgr) FetchIndexByWebsiteUp(websiteUp int16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_up` = ?", websiteUp).Find(&results).Error

	return
}

// FetchIndexByWebsiteDown  获取多个内容
func (obj *_MacWebsiteMgr) FetchIndexByWebsiteDown(websiteDown int16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_down` = ?", websiteDown).Find(&results).Error

	return
}

// FetchIndexByWebsiteReferer  获取多个内容
func (obj *_MacWebsiteMgr) FetchIndexByWebsiteReferer(websiteReferer int16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_referer` = ?", websiteReferer).Find(&results).Error

	return
}

// FetchIndexByWebsiteRefererDay  获取多个内容
func (obj *_MacWebsiteMgr) FetchIndexByWebsiteRefererDay(websiteRefererDay int16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_referer_day` = ?", websiteRefererDay).Find(&results).Error

	return
}

// FetchIndexByWebsiteRefererWeek  获取多个内容
func (obj *_MacWebsiteMgr) FetchIndexByWebsiteRefererWeek(websiteRefererWeek int16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_referer_week` = ?", websiteRefererWeek).Find(&results).Error

	return
}

// FetchIndexByWebsiteRefererMonth  获取多个内容
func (obj *_MacWebsiteMgr) FetchIndexByWebsiteRefererMonth(websiteRefererMonth int16) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_referer_month` = ?", websiteRefererMonth).Find(&results).Error

	return
}

// FetchIndexByWebsiteTag  获取多个内容
func (obj *_MacWebsiteMgr) FetchIndexByWebsiteTag(websiteTag string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_tag` = ?", websiteTag).Find(&results).Error

	return
}

// FetchIndexByWebsiteClass  获取多个内容
func (obj *_MacWebsiteMgr) FetchIndexByWebsiteClass(websiteClass string) (results []*MacWebsite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacWebsite{}).Where("`website_class` = ?", websiteClass).Find(&results).Error

	return
}
