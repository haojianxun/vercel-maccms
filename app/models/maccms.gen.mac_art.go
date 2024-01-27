package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _MacArtMgr struct {
	*_BaseMgr
}

// MacArtMgr open func
func MacArtMgr(db *gorm.DB) *_MacArtMgr {
	if db == nil {
		panic(fmt.Errorf("MacArtMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MacArtMgr{_BaseMgr: &_BaseMgr{DB: db.Table("mac_art"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_MacArtMgr) Debug() *_MacArtMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MacArtMgr) GetTableName() string {
	return "mac_art"
}

// Reset 重置gorm会话
func (obj *_MacArtMgr) Reset() *_MacArtMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MacArtMgr) Get() (result MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MacArtMgr) Gets() (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MacArtMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MacArt{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithArtID art_id获取 文章ID
func (obj *_MacArtMgr) WithArtID(artID int) Option {
	return optionFunc(func(o *options) { o.query["art_id"] = artID })
}

// WithTypeID type_id获取 类型ID
func (obj *_MacArtMgr) WithTypeID(typeID uint16) Option {
	return optionFunc(func(o *options) { o.query["type_id"] = typeID })
}

// WithTypeID1 type_id_1获取 另一类型ID
func (obj *_MacArtMgr) WithTypeID1(typeID1 uint16) Option {
	return optionFunc(func(o *options) { o.query["type_id_1"] = typeID1 })
}

// WithGroupID group_id获取 分组ID
func (obj *_MacArtMgr) WithGroupID(groupID uint16) Option {
	return optionFunc(func(o *options) { o.query["group_id"] = groupID })
}

// WithArtName art_name获取 文章名称
func (obj *_MacArtMgr) WithArtName(artName string) Option {
	return optionFunc(func(o *options) { o.query["art_name"] = artName })
}

// WithArtSub art_sub获取 文章副标题
func (obj *_MacArtMgr) WithArtSub(artSub string) Option {
	return optionFunc(func(o *options) { o.query["art_sub"] = artSub })
}

// WithArtEn art_en获取 文章英文名
func (obj *_MacArtMgr) WithArtEn(artEn string) Option {
	return optionFunc(func(o *options) { o.query["art_en"] = artEn })
}

// WithArtStatus art_status获取 文章状态
func (obj *_MacArtMgr) WithArtStatus(artStatus uint8) Option {
	return optionFunc(func(o *options) { o.query["art_status"] = artStatus })
}

// WithArtLetter art_letter获取 文章姓氏首字母
func (obj *_MacArtMgr) WithArtLetter(artLetter string) Option {
	return optionFunc(func(o *options) { o.query["art_letter"] = artLetter })
}

// WithArtColor art_color获取 文章颜色
func (obj *_MacArtMgr) WithArtColor(artColor string) Option {
	return optionFunc(func(o *options) { o.query["art_color"] = artColor })
}

// WithArtFrom art_from获取 文章来源
func (obj *_MacArtMgr) WithArtFrom(artFrom string) Option {
	return optionFunc(func(o *options) { o.query["art_from"] = artFrom })
}

// WithArtAuthor art_author获取 文章作者
func (obj *_MacArtMgr) WithArtAuthor(artAuthor string) Option {
	return optionFunc(func(o *options) { o.query["art_author"] = artAuthor })
}

// WithArtTag art_tag获取 文章标签
func (obj *_MacArtMgr) WithArtTag(artTag string) Option {
	return optionFunc(func(o *options) { o.query["art_tag"] = artTag })
}

// WithArtClass art_class获取 文章类别
func (obj *_MacArtMgr) WithArtClass(artClass string) Option {
	return optionFunc(func(o *options) { o.query["art_class"] = artClass })
}

// WithArtPic art_pic获取 文章图片链接
func (obj *_MacArtMgr) WithArtPic(artPic string) Option {
	return optionFunc(func(o *options) { o.query["art_pic"] = artPic })
}

// WithArtPicThumb art_pic_thumb获取 文章缩略图链接
func (obj *_MacArtMgr) WithArtPicThumb(artPicThumb string) Option {
	return optionFunc(func(o *options) { o.query["art_pic_thumb"] = artPicThumb })
}

// WithArtPicSlide art_pic_slide获取 文章幻灯片图片链接
func (obj *_MacArtMgr) WithArtPicSlide(artPicSlide string) Option {
	return optionFunc(func(o *options) { o.query["art_pic_slide"] = artPicSlide })
}

// WithArtPicScreenshot art_pic_screenshot获取 文章截图图片链接
func (obj *_MacArtMgr) WithArtPicScreenshot(artPicScreenshot string) Option {
	return optionFunc(func(o *options) { o.query["art_pic_screenshot"] = artPicScreenshot })
}

// WithArtBlurb art_blurb获取 文章简介
func (obj *_MacArtMgr) WithArtBlurb(artBlurb string) Option {
	return optionFunc(func(o *options) { o.query["art_blurb"] = artBlurb })
}

// WithArtRemarks art_remarks获取 文章备注
func (obj *_MacArtMgr) WithArtRemarks(artRemarks string) Option {
	return optionFunc(func(o *options) { o.query["art_remarks"] = artRemarks })
}

// WithArtJumpurl art_jumpurl获取 文章跳转链接
func (obj *_MacArtMgr) WithArtJumpurl(artJumpurl string) Option {
	return optionFunc(func(o *options) { o.query["art_jumpurl"] = artJumpurl })
}

// WithArtTpl art_tpl获取 文章模板
func (obj *_MacArtMgr) WithArtTpl(artTpl string) Option {
	return optionFunc(func(o *options) { o.query["art_tpl"] = artTpl })
}

// WithArtLevel art_level获取 文章级别
func (obj *_MacArtMgr) WithArtLevel(artLevel uint8) Option {
	return optionFunc(func(o *options) { o.query["art_level"] = artLevel })
}

// WithArtLock art_lock获取 文章锁定状态
func (obj *_MacArtMgr) WithArtLock(artLock uint8) Option {
	return optionFunc(func(o *options) { o.query["art_lock"] = artLock })
}

// WithArtPoints art_points获取 文章积分
func (obj *_MacArtMgr) WithArtPoints(artPoints uint16) Option {
	return optionFunc(func(o *options) { o.query["art_points"] = artPoints })
}

// WithArtPointsDetail art_points_detail获取 文章详细积分
func (obj *_MacArtMgr) WithArtPointsDetail(artPointsDetail uint16) Option {
	return optionFunc(func(o *options) { o.query["art_points_detail"] = artPointsDetail })
}

// WithArtUp art_up获取 文章点赞次数
func (obj *_MacArtMgr) WithArtUp(artUp int16) Option {
	return optionFunc(func(o *options) { o.query["art_up"] = artUp })
}

// WithArtDown art_down获取 文章踩次数
func (obj *_MacArtMgr) WithArtDown(artDown int16) Option {
	return optionFunc(func(o *options) { o.query["art_down"] = artDown })
}

// WithArtHits art_hits获取 文章点击次数
func (obj *_MacArtMgr) WithArtHits(artHits int16) Option {
	return optionFunc(func(o *options) { o.query["art_hits"] = artHits })
}

// WithArtHitsDay art_hits_day获取 文章每日点击次数
func (obj *_MacArtMgr) WithArtHitsDay(artHitsDay int16) Option {
	return optionFunc(func(o *options) { o.query["art_hits_day"] = artHitsDay })
}

// WithArtHitsWeek art_hits_week获取 文章每周点击次数
func (obj *_MacArtMgr) WithArtHitsWeek(artHitsWeek int16) Option {
	return optionFunc(func(o *options) { o.query["art_hits_week"] = artHitsWeek })
}

// WithArtHitsMonth art_hits_month获取 文章每月点击次数
func (obj *_MacArtMgr) WithArtHitsMonth(artHitsMonth int16) Option {
	return optionFunc(func(o *options) { o.query["art_hits_month"] = artHitsMonth })
}

// WithArtTime art_time获取 文章时间
func (obj *_MacArtMgr) WithArtTime(artTime int) Option {
	return optionFunc(func(o *options) { o.query["art_time"] = artTime })
}

// WithArtTimeAdd art_time_add获取 文章添加时间
func (obj *_MacArtMgr) WithArtTimeAdd(artTimeAdd int) Option {
	return optionFunc(func(o *options) { o.query["art_time_add"] = artTimeAdd })
}

// WithArtTimeHits art_time_hits获取 文章点击时间
func (obj *_MacArtMgr) WithArtTimeHits(artTimeHits int) Option {
	return optionFunc(func(o *options) { o.query["art_time_hits"] = artTimeHits })
}

// WithArtTimeMake art_time_make获取 文章制作时间
func (obj *_MacArtMgr) WithArtTimeMake(artTimeMake int) Option {
	return optionFunc(func(o *options) { o.query["art_time_make"] = artTimeMake })
}

// WithArtScore art_score获取 文章评分
func (obj *_MacArtMgr) WithArtScore(artScore float64) Option {
	return optionFunc(func(o *options) { o.query["art_score"] = artScore })
}

// WithArtScoreAll art_score_all获取 文章总评分
func (obj *_MacArtMgr) WithArtScoreAll(artScoreAll int16) Option {
	return optionFunc(func(o *options) { o.query["art_score_all"] = artScoreAll })
}

// WithArtScoreNum art_score_num获取 文章评分次数
func (obj *_MacArtMgr) WithArtScoreNum(artScoreNum int16) Option {
	return optionFunc(func(o *options) { o.query["art_score_num"] = artScoreNum })
}

// WithArtRelArt art_rel_art获取 相关文章链接
func (obj *_MacArtMgr) WithArtRelArt(artRelArt string) Option {
	return optionFunc(func(o *options) { o.query["art_rel_art"] = artRelArt })
}

// WithArtRelVod art_rel_vod获取 相关视频链接
func (obj *_MacArtMgr) WithArtRelVod(artRelVod string) Option {
	return optionFunc(func(o *options) { o.query["art_rel_vod"] = artRelVod })
}

// WithArtPwd art_pwd获取 文章密码
func (obj *_MacArtMgr) WithArtPwd(artPwd string) Option {
	return optionFunc(func(o *options) { o.query["art_pwd"] = artPwd })
}

// WithArtPwdURL art_pwd_url获取 文章密码链接
func (obj *_MacArtMgr) WithArtPwdURL(artPwdURL string) Option {
	return optionFunc(func(o *options) { o.query["art_pwd_url"] = artPwdURL })
}

// WithArtTitle art_title获取 文章标题
func (obj *_MacArtMgr) WithArtTitle(artTitle string) Option {
	return optionFunc(func(o *options) { o.query["art_title"] = artTitle })
}

// WithArtNote art_note获取 文章注释
func (obj *_MacArtMgr) WithArtNote(artNote string) Option {
	return optionFunc(func(o *options) { o.query["art_note"] = artNote })
}

// WithArtContent art_content获取 文章内容
func (obj *_MacArtMgr) WithArtContent(artContent string) Option {
	return optionFunc(func(o *options) { o.query["art_content"] = artContent })
}

// GetByOption 功能选项模式获取
func (obj *_MacArtMgr) GetByOption(opts ...Option) (result MacArt, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MacArtMgr) GetByOptions(opts ...Option) (results []*MacArt, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MacArtMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MacArt, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where(options.query)
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

// GetFromArtID 通过art_id获取内容 文章ID
func (obj *_MacArtMgr) GetFromArtID(artID int) (result MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_id` = ?", artID).First(&result).Error

	return
}

// GetBatchFromArtID 批量查找 文章ID
func (obj *_MacArtMgr) GetBatchFromArtID(artIDs []int) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_id` IN (?)", artIDs).Find(&results).Error

	return
}

// GetFromTypeID 通过type_id获取内容 类型ID
func (obj *_MacArtMgr) GetFromTypeID(typeID uint16) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`type_id` = ?", typeID).Find(&results).Error

	return
}

// GetBatchFromTypeID 批量查找 类型ID
func (obj *_MacArtMgr) GetBatchFromTypeID(typeIDs []uint16) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`type_id` IN (?)", typeIDs).Find(&results).Error

	return
}

// GetFromTypeID1 通过type_id_1获取内容 另一类型ID
func (obj *_MacArtMgr) GetFromTypeID1(typeID1 uint16) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`type_id_1` = ?", typeID1).Find(&results).Error

	return
}

// GetBatchFromTypeID1 批量查找 另一类型ID
func (obj *_MacArtMgr) GetBatchFromTypeID1(typeID1s []uint16) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`type_id_1` IN (?)", typeID1s).Find(&results).Error

	return
}

// GetFromGroupID 通过group_id获取内容 分组ID
func (obj *_MacArtMgr) GetFromGroupID(groupID uint16) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`group_id` = ?", groupID).Find(&results).Error

	return
}

// GetBatchFromGroupID 批量查找 分组ID
func (obj *_MacArtMgr) GetBatchFromGroupID(groupIDs []uint16) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`group_id` IN (?)", groupIDs).Find(&results).Error

	return
}

// GetFromArtName 通过art_name获取内容 文章名称
func (obj *_MacArtMgr) GetFromArtName(artName string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_name` = ?", artName).Find(&results).Error

	return
}

// GetBatchFromArtName 批量查找 文章名称
func (obj *_MacArtMgr) GetBatchFromArtName(artNames []string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_name` IN (?)", artNames).Find(&results).Error

	return
}

// GetFromArtSub 通过art_sub获取内容 文章副标题
func (obj *_MacArtMgr) GetFromArtSub(artSub string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_sub` = ?", artSub).Find(&results).Error

	return
}

// GetBatchFromArtSub 批量查找 文章副标题
func (obj *_MacArtMgr) GetBatchFromArtSub(artSubs []string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_sub` IN (?)", artSubs).Find(&results).Error

	return
}

// GetFromArtEn 通过art_en获取内容 文章英文名
func (obj *_MacArtMgr) GetFromArtEn(artEn string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_en` = ?", artEn).Find(&results).Error

	return
}

// GetBatchFromArtEn 批量查找 文章英文名
func (obj *_MacArtMgr) GetBatchFromArtEn(artEns []string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_en` IN (?)", artEns).Find(&results).Error

	return
}

// GetFromArtStatus 通过art_status获取内容 文章状态
func (obj *_MacArtMgr) GetFromArtStatus(artStatus uint8) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_status` = ?", artStatus).Find(&results).Error

	return
}

// GetBatchFromArtStatus 批量查找 文章状态
func (obj *_MacArtMgr) GetBatchFromArtStatus(artStatuss []uint8) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_status` IN (?)", artStatuss).Find(&results).Error

	return
}

// GetFromArtLetter 通过art_letter获取内容 文章姓氏首字母
func (obj *_MacArtMgr) GetFromArtLetter(artLetter string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_letter` = ?", artLetter).Find(&results).Error

	return
}

// GetBatchFromArtLetter 批量查找 文章姓氏首字母
func (obj *_MacArtMgr) GetBatchFromArtLetter(artLetters []string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_letter` IN (?)", artLetters).Find(&results).Error

	return
}

// GetFromArtColor 通过art_color获取内容 文章颜色
func (obj *_MacArtMgr) GetFromArtColor(artColor string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_color` = ?", artColor).Find(&results).Error

	return
}

// GetBatchFromArtColor 批量查找 文章颜色
func (obj *_MacArtMgr) GetBatchFromArtColor(artColors []string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_color` IN (?)", artColors).Find(&results).Error

	return
}

// GetFromArtFrom 通过art_from获取内容 文章来源
func (obj *_MacArtMgr) GetFromArtFrom(artFrom string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_from` = ?", artFrom).Find(&results).Error

	return
}

// GetBatchFromArtFrom 批量查找 文章来源
func (obj *_MacArtMgr) GetBatchFromArtFrom(artFroms []string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_from` IN (?)", artFroms).Find(&results).Error

	return
}

// GetFromArtAuthor 通过art_author获取内容 文章作者
func (obj *_MacArtMgr) GetFromArtAuthor(artAuthor string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_author` = ?", artAuthor).Find(&results).Error

	return
}

// GetBatchFromArtAuthor 批量查找 文章作者
func (obj *_MacArtMgr) GetBatchFromArtAuthor(artAuthors []string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_author` IN (?)", artAuthors).Find(&results).Error

	return
}

// GetFromArtTag 通过art_tag获取内容 文章标签
func (obj *_MacArtMgr) GetFromArtTag(artTag string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_tag` = ?", artTag).Find(&results).Error

	return
}

// GetBatchFromArtTag 批量查找 文章标签
func (obj *_MacArtMgr) GetBatchFromArtTag(artTags []string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_tag` IN (?)", artTags).Find(&results).Error

	return
}

// GetFromArtClass 通过art_class获取内容 文章类别
func (obj *_MacArtMgr) GetFromArtClass(artClass string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_class` = ?", artClass).Find(&results).Error

	return
}

// GetBatchFromArtClass 批量查找 文章类别
func (obj *_MacArtMgr) GetBatchFromArtClass(artClasss []string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_class` IN (?)", artClasss).Find(&results).Error

	return
}

// GetFromArtPic 通过art_pic获取内容 文章图片链接
func (obj *_MacArtMgr) GetFromArtPic(artPic string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_pic` = ?", artPic).Find(&results).Error

	return
}

// GetBatchFromArtPic 批量查找 文章图片链接
func (obj *_MacArtMgr) GetBatchFromArtPic(artPics []string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_pic` IN (?)", artPics).Find(&results).Error

	return
}

// GetFromArtPicThumb 通过art_pic_thumb获取内容 文章缩略图链接
func (obj *_MacArtMgr) GetFromArtPicThumb(artPicThumb string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_pic_thumb` = ?", artPicThumb).Find(&results).Error

	return
}

// GetBatchFromArtPicThumb 批量查找 文章缩略图链接
func (obj *_MacArtMgr) GetBatchFromArtPicThumb(artPicThumbs []string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_pic_thumb` IN (?)", artPicThumbs).Find(&results).Error

	return
}

// GetFromArtPicSlide 通过art_pic_slide获取内容 文章幻灯片图片链接
func (obj *_MacArtMgr) GetFromArtPicSlide(artPicSlide string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_pic_slide` = ?", artPicSlide).Find(&results).Error

	return
}

// GetBatchFromArtPicSlide 批量查找 文章幻灯片图片链接
func (obj *_MacArtMgr) GetBatchFromArtPicSlide(artPicSlides []string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_pic_slide` IN (?)", artPicSlides).Find(&results).Error

	return
}

// GetFromArtPicScreenshot 通过art_pic_screenshot获取内容 文章截图图片链接
func (obj *_MacArtMgr) GetFromArtPicScreenshot(artPicScreenshot string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_pic_screenshot` = ?", artPicScreenshot).Find(&results).Error

	return
}

// GetBatchFromArtPicScreenshot 批量查找 文章截图图片链接
func (obj *_MacArtMgr) GetBatchFromArtPicScreenshot(artPicScreenshots []string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_pic_screenshot` IN (?)", artPicScreenshots).Find(&results).Error

	return
}

// GetFromArtBlurb 通过art_blurb获取内容 文章简介
func (obj *_MacArtMgr) GetFromArtBlurb(artBlurb string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_blurb` = ?", artBlurb).Find(&results).Error

	return
}

// GetBatchFromArtBlurb 批量查找 文章简介
func (obj *_MacArtMgr) GetBatchFromArtBlurb(artBlurbs []string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_blurb` IN (?)", artBlurbs).Find(&results).Error

	return
}

// GetFromArtRemarks 通过art_remarks获取内容 文章备注
func (obj *_MacArtMgr) GetFromArtRemarks(artRemarks string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_remarks` = ?", artRemarks).Find(&results).Error

	return
}

// GetBatchFromArtRemarks 批量查找 文章备注
func (obj *_MacArtMgr) GetBatchFromArtRemarks(artRemarkss []string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_remarks` IN (?)", artRemarkss).Find(&results).Error

	return
}

// GetFromArtJumpurl 通过art_jumpurl获取内容 文章跳转链接
func (obj *_MacArtMgr) GetFromArtJumpurl(artJumpurl string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_jumpurl` = ?", artJumpurl).Find(&results).Error

	return
}

// GetBatchFromArtJumpurl 批量查找 文章跳转链接
func (obj *_MacArtMgr) GetBatchFromArtJumpurl(artJumpurls []string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_jumpurl` IN (?)", artJumpurls).Find(&results).Error

	return
}

// GetFromArtTpl 通过art_tpl获取内容 文章模板
func (obj *_MacArtMgr) GetFromArtTpl(artTpl string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_tpl` = ?", artTpl).Find(&results).Error

	return
}

// GetBatchFromArtTpl 批量查找 文章模板
func (obj *_MacArtMgr) GetBatchFromArtTpl(artTpls []string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_tpl` IN (?)", artTpls).Find(&results).Error

	return
}

// GetFromArtLevel 通过art_level获取内容 文章级别
func (obj *_MacArtMgr) GetFromArtLevel(artLevel uint8) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_level` = ?", artLevel).Find(&results).Error

	return
}

// GetBatchFromArtLevel 批量查找 文章级别
func (obj *_MacArtMgr) GetBatchFromArtLevel(artLevels []uint8) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_level` IN (?)", artLevels).Find(&results).Error

	return
}

// GetFromArtLock 通过art_lock获取内容 文章锁定状态
func (obj *_MacArtMgr) GetFromArtLock(artLock uint8) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_lock` = ?", artLock).Find(&results).Error

	return
}

// GetBatchFromArtLock 批量查找 文章锁定状态
func (obj *_MacArtMgr) GetBatchFromArtLock(artLocks []uint8) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_lock` IN (?)", artLocks).Find(&results).Error

	return
}

// GetFromArtPoints 通过art_points获取内容 文章积分
func (obj *_MacArtMgr) GetFromArtPoints(artPoints uint16) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_points` = ?", artPoints).Find(&results).Error

	return
}

// GetBatchFromArtPoints 批量查找 文章积分
func (obj *_MacArtMgr) GetBatchFromArtPoints(artPointss []uint16) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_points` IN (?)", artPointss).Find(&results).Error

	return
}

// GetFromArtPointsDetail 通过art_points_detail获取内容 文章详细积分
func (obj *_MacArtMgr) GetFromArtPointsDetail(artPointsDetail uint16) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_points_detail` = ?", artPointsDetail).Find(&results).Error

	return
}

// GetBatchFromArtPointsDetail 批量查找 文章详细积分
func (obj *_MacArtMgr) GetBatchFromArtPointsDetail(artPointsDetails []uint16) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_points_detail` IN (?)", artPointsDetails).Find(&results).Error

	return
}

// GetFromArtUp 通过art_up获取内容 文章点赞次数
func (obj *_MacArtMgr) GetFromArtUp(artUp int16) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_up` = ?", artUp).Find(&results).Error

	return
}

// GetBatchFromArtUp 批量查找 文章点赞次数
func (obj *_MacArtMgr) GetBatchFromArtUp(artUps []int16) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_up` IN (?)", artUps).Find(&results).Error

	return
}

// GetFromArtDown 通过art_down获取内容 文章踩次数
func (obj *_MacArtMgr) GetFromArtDown(artDown int16) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_down` = ?", artDown).Find(&results).Error

	return
}

// GetBatchFromArtDown 批量查找 文章踩次数
func (obj *_MacArtMgr) GetBatchFromArtDown(artDowns []int16) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_down` IN (?)", artDowns).Find(&results).Error

	return
}

// GetFromArtHits 通过art_hits获取内容 文章点击次数
func (obj *_MacArtMgr) GetFromArtHits(artHits int16) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_hits` = ?", artHits).Find(&results).Error

	return
}

// GetBatchFromArtHits 批量查找 文章点击次数
func (obj *_MacArtMgr) GetBatchFromArtHits(artHitss []int16) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_hits` IN (?)", artHitss).Find(&results).Error

	return
}

// GetFromArtHitsDay 通过art_hits_day获取内容 文章每日点击次数
func (obj *_MacArtMgr) GetFromArtHitsDay(artHitsDay int16) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_hits_day` = ?", artHitsDay).Find(&results).Error

	return
}

// GetBatchFromArtHitsDay 批量查找 文章每日点击次数
func (obj *_MacArtMgr) GetBatchFromArtHitsDay(artHitsDays []int16) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_hits_day` IN (?)", artHitsDays).Find(&results).Error

	return
}

// GetFromArtHitsWeek 通过art_hits_week获取内容 文章每周点击次数
func (obj *_MacArtMgr) GetFromArtHitsWeek(artHitsWeek int16) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_hits_week` = ?", artHitsWeek).Find(&results).Error

	return
}

// GetBatchFromArtHitsWeek 批量查找 文章每周点击次数
func (obj *_MacArtMgr) GetBatchFromArtHitsWeek(artHitsWeeks []int16) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_hits_week` IN (?)", artHitsWeeks).Find(&results).Error

	return
}

// GetFromArtHitsMonth 通过art_hits_month获取内容 文章每月点击次数
func (obj *_MacArtMgr) GetFromArtHitsMonth(artHitsMonth int16) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_hits_month` = ?", artHitsMonth).Find(&results).Error

	return
}

// GetBatchFromArtHitsMonth 批量查找 文章每月点击次数
func (obj *_MacArtMgr) GetBatchFromArtHitsMonth(artHitsMonths []int16) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_hits_month` IN (?)", artHitsMonths).Find(&results).Error

	return
}

// GetFromArtTime 通过art_time获取内容 文章时间
func (obj *_MacArtMgr) GetFromArtTime(artTime int) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_time` = ?", artTime).Find(&results).Error

	return
}

// GetBatchFromArtTime 批量查找 文章时间
func (obj *_MacArtMgr) GetBatchFromArtTime(artTimes []int) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_time` IN (?)", artTimes).Find(&results).Error

	return
}

// GetFromArtTimeAdd 通过art_time_add获取内容 文章添加时间
func (obj *_MacArtMgr) GetFromArtTimeAdd(artTimeAdd int) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_time_add` = ?", artTimeAdd).Find(&results).Error

	return
}

// GetBatchFromArtTimeAdd 批量查找 文章添加时间
func (obj *_MacArtMgr) GetBatchFromArtTimeAdd(artTimeAdds []int) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_time_add` IN (?)", artTimeAdds).Find(&results).Error

	return
}

// GetFromArtTimeHits 通过art_time_hits获取内容 文章点击时间
func (obj *_MacArtMgr) GetFromArtTimeHits(artTimeHits int) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_time_hits` = ?", artTimeHits).Find(&results).Error

	return
}

// GetBatchFromArtTimeHits 批量查找 文章点击时间
func (obj *_MacArtMgr) GetBatchFromArtTimeHits(artTimeHitss []int) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_time_hits` IN (?)", artTimeHitss).Find(&results).Error

	return
}

// GetFromArtTimeMake 通过art_time_make获取内容 文章制作时间
func (obj *_MacArtMgr) GetFromArtTimeMake(artTimeMake int) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_time_make` = ?", artTimeMake).Find(&results).Error

	return
}

// GetBatchFromArtTimeMake 批量查找 文章制作时间
func (obj *_MacArtMgr) GetBatchFromArtTimeMake(artTimeMakes []int) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_time_make` IN (?)", artTimeMakes).Find(&results).Error

	return
}

// GetFromArtScore 通过art_score获取内容 文章评分
func (obj *_MacArtMgr) GetFromArtScore(artScore float64) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_score` = ?", artScore).Find(&results).Error

	return
}

// GetBatchFromArtScore 批量查找 文章评分
func (obj *_MacArtMgr) GetBatchFromArtScore(artScores []float64) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_score` IN (?)", artScores).Find(&results).Error

	return
}

// GetFromArtScoreAll 通过art_score_all获取内容 文章总评分
func (obj *_MacArtMgr) GetFromArtScoreAll(artScoreAll int16) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_score_all` = ?", artScoreAll).Find(&results).Error

	return
}

// GetBatchFromArtScoreAll 批量查找 文章总评分
func (obj *_MacArtMgr) GetBatchFromArtScoreAll(artScoreAlls []int16) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_score_all` IN (?)", artScoreAlls).Find(&results).Error

	return
}

// GetFromArtScoreNum 通过art_score_num获取内容 文章评分次数
func (obj *_MacArtMgr) GetFromArtScoreNum(artScoreNum int16) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_score_num` = ?", artScoreNum).Find(&results).Error

	return
}

// GetBatchFromArtScoreNum 批量查找 文章评分次数
func (obj *_MacArtMgr) GetBatchFromArtScoreNum(artScoreNums []int16) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_score_num` IN (?)", artScoreNums).Find(&results).Error

	return
}

// GetFromArtRelArt 通过art_rel_art获取内容 相关文章链接
func (obj *_MacArtMgr) GetFromArtRelArt(artRelArt string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_rel_art` = ?", artRelArt).Find(&results).Error

	return
}

// GetBatchFromArtRelArt 批量查找 相关文章链接
func (obj *_MacArtMgr) GetBatchFromArtRelArt(artRelArts []string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_rel_art` IN (?)", artRelArts).Find(&results).Error

	return
}

// GetFromArtRelVod 通过art_rel_vod获取内容 相关视频链接
func (obj *_MacArtMgr) GetFromArtRelVod(artRelVod string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_rel_vod` = ?", artRelVod).Find(&results).Error

	return
}

// GetBatchFromArtRelVod 批量查找 相关视频链接
func (obj *_MacArtMgr) GetBatchFromArtRelVod(artRelVods []string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_rel_vod` IN (?)", artRelVods).Find(&results).Error

	return
}

// GetFromArtPwd 通过art_pwd获取内容 文章密码
func (obj *_MacArtMgr) GetFromArtPwd(artPwd string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_pwd` = ?", artPwd).Find(&results).Error

	return
}

// GetBatchFromArtPwd 批量查找 文章密码
func (obj *_MacArtMgr) GetBatchFromArtPwd(artPwds []string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_pwd` IN (?)", artPwds).Find(&results).Error

	return
}

// GetFromArtPwdURL 通过art_pwd_url获取内容 文章密码链接
func (obj *_MacArtMgr) GetFromArtPwdURL(artPwdURL string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_pwd_url` = ?", artPwdURL).Find(&results).Error

	return
}

// GetBatchFromArtPwdURL 批量查找 文章密码链接
func (obj *_MacArtMgr) GetBatchFromArtPwdURL(artPwdURLs []string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_pwd_url` IN (?)", artPwdURLs).Find(&results).Error

	return
}

// GetFromArtTitle 通过art_title获取内容 文章标题
func (obj *_MacArtMgr) GetFromArtTitle(artTitle string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_title` = ?", artTitle).Find(&results).Error

	return
}

// GetBatchFromArtTitle 批量查找 文章标题
func (obj *_MacArtMgr) GetBatchFromArtTitle(artTitles []string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_title` IN (?)", artTitles).Find(&results).Error

	return
}

// GetFromArtNote 通过art_note获取内容 文章注释
func (obj *_MacArtMgr) GetFromArtNote(artNote string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_note` = ?", artNote).Find(&results).Error

	return
}

// GetBatchFromArtNote 批量查找 文章注释
func (obj *_MacArtMgr) GetBatchFromArtNote(artNotes []string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_note` IN (?)", artNotes).Find(&results).Error

	return
}

// GetFromArtContent 通过art_content获取内容 文章内容
func (obj *_MacArtMgr) GetFromArtContent(artContent string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_content` = ?", artContent).Find(&results).Error

	return
}

// GetBatchFromArtContent 批量查找 文章内容
func (obj *_MacArtMgr) GetBatchFromArtContent(artContents []string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_content` IN (?)", artContents).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MacArtMgr) FetchByPrimaryKey(artID int) (result MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_id` = ?", artID).First(&result).Error

	return
}

// FetchIndexByTypeID  获取多个内容
func (obj *_MacArtMgr) FetchIndexByTypeID(typeID uint16) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`type_id` = ?", typeID).Find(&results).Error

	return
}

// FetchIndexByTypeID1  获取多个内容
func (obj *_MacArtMgr) FetchIndexByTypeID1(typeID1 uint16) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`type_id_1` = ?", typeID1).Find(&results).Error

	return
}

// FetchIndexByArtName  获取多个内容
func (obj *_MacArtMgr) FetchIndexByArtName(artName string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_name` = ?", artName).Find(&results).Error

	return
}

// FetchIndexByArtEnn  获取多个内容
func (obj *_MacArtMgr) FetchIndexByArtEnn(artEn string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_en` = ?", artEn).Find(&results).Error

	return
}

// FetchIndexByArtLetter  获取多个内容
func (obj *_MacArtMgr) FetchIndexByArtLetter(artLetter string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_letter` = ?", artLetter).Find(&results).Error

	return
}

// FetchIndexByArtTag  获取多个内容
func (obj *_MacArtMgr) FetchIndexByArtTag(artTag string) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_tag` = ?", artTag).Find(&results).Error

	return
}

// FetchIndexByArtLevel  获取多个内容
func (obj *_MacArtMgr) FetchIndexByArtLevel(artLevel uint8) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_level` = ?", artLevel).Find(&results).Error

	return
}

// FetchIndexByArtLock  获取多个内容
func (obj *_MacArtMgr) FetchIndexByArtLock(artLock uint8) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_lock` = ?", artLock).Find(&results).Error

	return
}

// FetchIndexByArtUp  获取多个内容
func (obj *_MacArtMgr) FetchIndexByArtUp(artUp int16) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_up` = ?", artUp).Find(&results).Error

	return
}

// FetchIndexByArtDown  获取多个内容
func (obj *_MacArtMgr) FetchIndexByArtDown(artDown int16) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_down` = ?", artDown).Find(&results).Error

	return
}

// FetchIndexByArtHits  获取多个内容
func (obj *_MacArtMgr) FetchIndexByArtHits(artHits int16) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_hits` = ?", artHits).Find(&results).Error

	return
}

// FetchIndexByArtHitsDay  获取多个内容
func (obj *_MacArtMgr) FetchIndexByArtHitsDay(artHitsDay int16) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_hits_day` = ?", artHitsDay).Find(&results).Error

	return
}

// FetchIndexByArtHitsWeek  获取多个内容
func (obj *_MacArtMgr) FetchIndexByArtHitsWeek(artHitsWeek int16) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_hits_week` = ?", artHitsWeek).Find(&results).Error

	return
}

// FetchIndexByArtHitsMonth  获取多个内容
func (obj *_MacArtMgr) FetchIndexByArtHitsMonth(artHitsMonth int16) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_hits_month` = ?", artHitsMonth).Find(&results).Error

	return
}

// FetchIndexByArtTime  获取多个内容
func (obj *_MacArtMgr) FetchIndexByArtTime(artTime int) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_time` = ?", artTime).Find(&results).Error

	return
}

// FetchIndexByArtTimeAdd  获取多个内容
func (obj *_MacArtMgr) FetchIndexByArtTimeAdd(artTimeAdd int) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_time_add` = ?", artTimeAdd).Find(&results).Error

	return
}

// FetchIndexByArtTimeMake  获取多个内容
func (obj *_MacArtMgr) FetchIndexByArtTimeMake(artTimeMake int) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_time_make` = ?", artTimeMake).Find(&results).Error

	return
}

// FetchIndexByArtScore  获取多个内容
func (obj *_MacArtMgr) FetchIndexByArtScore(artScore float64) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_score` = ?", artScore).Find(&results).Error

	return
}

// FetchIndexByArtScoreAll  获取多个内容
func (obj *_MacArtMgr) FetchIndexByArtScoreAll(artScoreAll int16) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_score_all` = ?", artScoreAll).Find(&results).Error

	return
}

// FetchIndexByArtScoreNum  获取多个内容
func (obj *_MacArtMgr) FetchIndexByArtScoreNum(artScoreNum int16) (results []*MacArt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacArt{}).Where("`art_score_num` = ?", artScoreNum).Find(&results).Error

	return
}
