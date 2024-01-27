package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _MacActorMgr struct {
	*_BaseMgr
}

// MacActorMgr open func
func MacActorMgr(db *gorm.DB) *_MacActorMgr {
	if db == nil {
		panic(fmt.Errorf("MacActorMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MacActorMgr{_BaseMgr: &_BaseMgr{DB: db.Table("mac_actor"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_MacActorMgr) Debug() *_MacActorMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MacActorMgr) GetTableName() string {
	return "mac_actor"
}

// Reset 重置gorm会话
func (obj *_MacActorMgr) Reset() *_MacActorMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MacActorMgr) Get() (result MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MacActorMgr) Gets() (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MacActorMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MacActor{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithActorID actor_id获取 演员ID
func (obj *_MacActorMgr) WithActorID(actorID int) Option {
	return optionFunc(func(o *options) { o.query["actor_id"] = actorID })
}

// WithTypeID type_id获取 类型ID
func (obj *_MacActorMgr) WithTypeID(typeID uint16) Option {
	return optionFunc(func(o *options) { o.query["type_id"] = typeID })
}

// WithTypeID1 type_id_1获取 另一类型ID
func (obj *_MacActorMgr) WithTypeID1(typeID1 uint16) Option {
	return optionFunc(func(o *options) { o.query["type_id_1"] = typeID1 })
}

// WithActorName actor_name获取 演员姓名
func (obj *_MacActorMgr) WithActorName(actorName string) Option {
	return optionFunc(func(o *options) { o.query["actor_name"] = actorName })
}

// WithActorEn actor_en获取 演员英文名
func (obj *_MacActorMgr) WithActorEn(actorEn string) Option {
	return optionFunc(func(o *options) { o.query["actor_en"] = actorEn })
}

// WithActorAlias actor_alias获取 演员别名
func (obj *_MacActorMgr) WithActorAlias(actorAlias string) Option {
	return optionFunc(func(o *options) { o.query["actor_alias"] = actorAlias })
}

// WithActorStatus actor_status获取 演员状态
func (obj *_MacActorMgr) WithActorStatus(actorStatus uint8) Option {
	return optionFunc(func(o *options) { o.query["actor_status"] = actorStatus })
}

// WithActorLock actor_lock获取 演员锁定状态
func (obj *_MacActorMgr) WithActorLock(actorLock uint8) Option {
	return optionFunc(func(o *options) { o.query["actor_lock"] = actorLock })
}

// WithActorLetter actor_letter获取 演员姓氏首字母
func (obj *_MacActorMgr) WithActorLetter(actorLetter string) Option {
	return optionFunc(func(o *options) { o.query["actor_letter"] = actorLetter })
}

// WithActorSex actor_sex获取 演员性别
func (obj *_MacActorMgr) WithActorSex(actorSex string) Option {
	return optionFunc(func(o *options) { o.query["actor_sex"] = actorSex })
}

// WithActorColor actor_color获取 演员颜色
func (obj *_MacActorMgr) WithActorColor(actorColor string) Option {
	return optionFunc(func(o *options) { o.query["actor_color"] = actorColor })
}

// WithActorPic actor_pic获取 演员照片链接
func (obj *_MacActorMgr) WithActorPic(actorPic string) Option {
	return optionFunc(func(o *options) { o.query["actor_pic"] = actorPic })
}

// WithActorBlurb actor_blurb获取 演员简介
func (obj *_MacActorMgr) WithActorBlurb(actorBlurb string) Option {
	return optionFunc(func(o *options) { o.query["actor_blurb"] = actorBlurb })
}

// WithActorRemarks actor_remarks获取 演员备注
func (obj *_MacActorMgr) WithActorRemarks(actorRemarks string) Option {
	return optionFunc(func(o *options) { o.query["actor_remarks"] = actorRemarks })
}

// WithActorArea actor_area获取 演员地区
func (obj *_MacActorMgr) WithActorArea(actorArea string) Option {
	return optionFunc(func(o *options) { o.query["actor_area"] = actorArea })
}

// WithActorHeight actor_height获取 演员身高
func (obj *_MacActorMgr) WithActorHeight(actorHeight string) Option {
	return optionFunc(func(o *options) { o.query["actor_height"] = actorHeight })
}

// WithActorWeight actor_weight获取 演员体重
func (obj *_MacActorMgr) WithActorWeight(actorWeight string) Option {
	return optionFunc(func(o *options) { o.query["actor_weight"] = actorWeight })
}

// WithActorBirthday actor_birthday获取 演员生日
func (obj *_MacActorMgr) WithActorBirthday(actorBirthday string) Option {
	return optionFunc(func(o *options) { o.query["actor_birthday"] = actorBirthday })
}

// WithActorBirtharea actor_birtharea获取 演员出生地区
func (obj *_MacActorMgr) WithActorBirtharea(actorBirtharea string) Option {
	return optionFunc(func(o *options) { o.query["actor_birtharea"] = actorBirtharea })
}

// WithActorBlood actor_blood获取 演员血型
func (obj *_MacActorMgr) WithActorBlood(actorBlood string) Option {
	return optionFunc(func(o *options) { o.query["actor_blood"] = actorBlood })
}

// WithActorStarsign actor_starsign获取 演员星座
func (obj *_MacActorMgr) WithActorStarsign(actorStarsign string) Option {
	return optionFunc(func(o *options) { o.query["actor_starsign"] = actorStarsign })
}

// WithActorSchool actor_school获取 演员毕业学校
func (obj *_MacActorMgr) WithActorSchool(actorSchool string) Option {
	return optionFunc(func(o *options) { o.query["actor_school"] = actorSchool })
}

// WithActorWorks actor_works获取 演员作品
func (obj *_MacActorMgr) WithActorWorks(actorWorks string) Option {
	return optionFunc(func(o *options) { o.query["actor_works"] = actorWorks })
}

// WithActorTag actor_tag获取 演员标签
func (obj *_MacActorMgr) WithActorTag(actorTag string) Option {
	return optionFunc(func(o *options) { o.query["actor_tag"] = actorTag })
}

// WithActorClass actor_class获取 演员类别
func (obj *_MacActorMgr) WithActorClass(actorClass string) Option {
	return optionFunc(func(o *options) { o.query["actor_class"] = actorClass })
}

// WithActorLevel actor_level获取 演员级别
func (obj *_MacActorMgr) WithActorLevel(actorLevel uint8) Option {
	return optionFunc(func(o *options) { o.query["actor_level"] = actorLevel })
}

// WithActorTime actor_time获取 演员时间
func (obj *_MacActorMgr) WithActorTime(actorTime int) Option {
	return optionFunc(func(o *options) { o.query["actor_time"] = actorTime })
}

// WithActorTimeAdd actor_time_add获取 演员添加时间
func (obj *_MacActorMgr) WithActorTimeAdd(actorTimeAdd int) Option {
	return optionFunc(func(o *options) { o.query["actor_time_add"] = actorTimeAdd })
}

// WithActorTimeHits actor_time_hits获取 演员点击时间
func (obj *_MacActorMgr) WithActorTimeHits(actorTimeHits int) Option {
	return optionFunc(func(o *options) { o.query["actor_time_hits"] = actorTimeHits })
}

// WithActorTimeMake actor_time_make获取 演员制作时间
func (obj *_MacActorMgr) WithActorTimeMake(actorTimeMake int) Option {
	return optionFunc(func(o *options) { o.query["actor_time_make"] = actorTimeMake })
}

// WithActorHits actor_hits获取 演员点击次数
func (obj *_MacActorMgr) WithActorHits(actorHits int16) Option {
	return optionFunc(func(o *options) { o.query["actor_hits"] = actorHits })
}

// WithActorHitsDay actor_hits_day获取 演员每日点击次数
func (obj *_MacActorMgr) WithActorHitsDay(actorHitsDay int16) Option {
	return optionFunc(func(o *options) { o.query["actor_hits_day"] = actorHitsDay })
}

// WithActorHitsWeek actor_hits_week获取 演员每周点击次数
func (obj *_MacActorMgr) WithActorHitsWeek(actorHitsWeek int16) Option {
	return optionFunc(func(o *options) { o.query["actor_hits_week"] = actorHitsWeek })
}

// WithActorHitsMonth actor_hits_month获取 演员每月点击次数
func (obj *_MacActorMgr) WithActorHitsMonth(actorHitsMonth int16) Option {
	return optionFunc(func(o *options) { o.query["actor_hits_month"] = actorHitsMonth })
}

// WithActorScore actor_score获取 演员评分
func (obj *_MacActorMgr) WithActorScore(actorScore float64) Option {
	return optionFunc(func(o *options) { o.query["actor_score"] = actorScore })
}

// WithActorScoreAll actor_score_all获取 演员总评分
func (obj *_MacActorMgr) WithActorScoreAll(actorScoreAll int16) Option {
	return optionFunc(func(o *options) { o.query["actor_score_all"] = actorScoreAll })
}

// WithActorScoreNum actor_score_num获取 演员评分次数
func (obj *_MacActorMgr) WithActorScoreNum(actorScoreNum int16) Option {
	return optionFunc(func(o *options) { o.query["actor_score_num"] = actorScoreNum })
}

// WithActorUp actor_up获取 演员点赞次数
func (obj *_MacActorMgr) WithActorUp(actorUp int16) Option {
	return optionFunc(func(o *options) { o.query["actor_up"] = actorUp })
}

// WithActorDown actor_down获取 演员踩次数
func (obj *_MacActorMgr) WithActorDown(actorDown int16) Option {
	return optionFunc(func(o *options) { o.query["actor_down"] = actorDown })
}

// WithActorTpl actor_tpl获取 演员模板
func (obj *_MacActorMgr) WithActorTpl(actorTpl string) Option {
	return optionFunc(func(o *options) { o.query["actor_tpl"] = actorTpl })
}

// WithActorJumpurl actor_jumpurl获取 演员跳转链接
func (obj *_MacActorMgr) WithActorJumpurl(actorJumpurl string) Option {
	return optionFunc(func(o *options) { o.query["actor_jumpurl"] = actorJumpurl })
}

// WithActorContent actor_content获取 演员内容
func (obj *_MacActorMgr) WithActorContent(actorContent string) Option {
	return optionFunc(func(o *options) { o.query["actor_content"] = actorContent })
}

// GetByOption 功能选项模式获取
func (obj *_MacActorMgr) GetByOption(opts ...Option) (result MacActor, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MacActorMgr) GetByOptions(opts ...Option) (results []*MacActor, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MacActorMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MacActor, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where(options.query)
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

// GetFromActorID 通过actor_id获取内容 演员ID
func (obj *_MacActorMgr) GetFromActorID(actorID int) (result MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_id` = ?", actorID).First(&result).Error

	return
}

// GetBatchFromActorID 批量查找 演员ID
func (obj *_MacActorMgr) GetBatchFromActorID(actorIDs []int) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_id` IN (?)", actorIDs).Find(&results).Error

	return
}

// GetFromTypeID 通过type_id获取内容 类型ID
func (obj *_MacActorMgr) GetFromTypeID(typeID uint16) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`type_id` = ?", typeID).Find(&results).Error

	return
}

// GetBatchFromTypeID 批量查找 类型ID
func (obj *_MacActorMgr) GetBatchFromTypeID(typeIDs []uint16) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`type_id` IN (?)", typeIDs).Find(&results).Error

	return
}

// GetFromTypeID1 通过type_id_1获取内容 另一类型ID
func (obj *_MacActorMgr) GetFromTypeID1(typeID1 uint16) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`type_id_1` = ?", typeID1).Find(&results).Error

	return
}

// GetBatchFromTypeID1 批量查找 另一类型ID
func (obj *_MacActorMgr) GetBatchFromTypeID1(typeID1s []uint16) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`type_id_1` IN (?)", typeID1s).Find(&results).Error

	return
}

// GetFromActorName 通过actor_name获取内容 演员姓名
func (obj *_MacActorMgr) GetFromActorName(actorName string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_name` = ?", actorName).Find(&results).Error

	return
}

// GetBatchFromActorName 批量查找 演员姓名
func (obj *_MacActorMgr) GetBatchFromActorName(actorNames []string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_name` IN (?)", actorNames).Find(&results).Error

	return
}

// GetFromActorEn 通过actor_en获取内容 演员英文名
func (obj *_MacActorMgr) GetFromActorEn(actorEn string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_en` = ?", actorEn).Find(&results).Error

	return
}

// GetBatchFromActorEn 批量查找 演员英文名
func (obj *_MacActorMgr) GetBatchFromActorEn(actorEns []string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_en` IN (?)", actorEns).Find(&results).Error

	return
}

// GetFromActorAlias 通过actor_alias获取内容 演员别名
func (obj *_MacActorMgr) GetFromActorAlias(actorAlias string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_alias` = ?", actorAlias).Find(&results).Error

	return
}

// GetBatchFromActorAlias 批量查找 演员别名
func (obj *_MacActorMgr) GetBatchFromActorAlias(actorAliass []string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_alias` IN (?)", actorAliass).Find(&results).Error

	return
}

// GetFromActorStatus 通过actor_status获取内容 演员状态
func (obj *_MacActorMgr) GetFromActorStatus(actorStatus uint8) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_status` = ?", actorStatus).Find(&results).Error

	return
}

// GetBatchFromActorStatus 批量查找 演员状态
func (obj *_MacActorMgr) GetBatchFromActorStatus(actorStatuss []uint8) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_status` IN (?)", actorStatuss).Find(&results).Error

	return
}

// GetFromActorLock 通过actor_lock获取内容 演员锁定状态
func (obj *_MacActorMgr) GetFromActorLock(actorLock uint8) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_lock` = ?", actorLock).Find(&results).Error

	return
}

// GetBatchFromActorLock 批量查找 演员锁定状态
func (obj *_MacActorMgr) GetBatchFromActorLock(actorLocks []uint8) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_lock` IN (?)", actorLocks).Find(&results).Error

	return
}

// GetFromActorLetter 通过actor_letter获取内容 演员姓氏首字母
func (obj *_MacActorMgr) GetFromActorLetter(actorLetter string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_letter` = ?", actorLetter).Find(&results).Error

	return
}

// GetBatchFromActorLetter 批量查找 演员姓氏首字母
func (obj *_MacActorMgr) GetBatchFromActorLetter(actorLetters []string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_letter` IN (?)", actorLetters).Find(&results).Error

	return
}

// GetFromActorSex 通过actor_sex获取内容 演员性别
func (obj *_MacActorMgr) GetFromActorSex(actorSex string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_sex` = ?", actorSex).Find(&results).Error

	return
}

// GetBatchFromActorSex 批量查找 演员性别
func (obj *_MacActorMgr) GetBatchFromActorSex(actorSexs []string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_sex` IN (?)", actorSexs).Find(&results).Error

	return
}

// GetFromActorColor 通过actor_color获取内容 演员颜色
func (obj *_MacActorMgr) GetFromActorColor(actorColor string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_color` = ?", actorColor).Find(&results).Error

	return
}

// GetBatchFromActorColor 批量查找 演员颜色
func (obj *_MacActorMgr) GetBatchFromActorColor(actorColors []string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_color` IN (?)", actorColors).Find(&results).Error

	return
}

// GetFromActorPic 通过actor_pic获取内容 演员照片链接
func (obj *_MacActorMgr) GetFromActorPic(actorPic string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_pic` = ?", actorPic).Find(&results).Error

	return
}

// GetBatchFromActorPic 批量查找 演员照片链接
func (obj *_MacActorMgr) GetBatchFromActorPic(actorPics []string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_pic` IN (?)", actorPics).Find(&results).Error

	return
}

// GetFromActorBlurb 通过actor_blurb获取内容 演员简介
func (obj *_MacActorMgr) GetFromActorBlurb(actorBlurb string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_blurb` = ?", actorBlurb).Find(&results).Error

	return
}

// GetBatchFromActorBlurb 批量查找 演员简介
func (obj *_MacActorMgr) GetBatchFromActorBlurb(actorBlurbs []string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_blurb` IN (?)", actorBlurbs).Find(&results).Error

	return
}

// GetFromActorRemarks 通过actor_remarks获取内容 演员备注
func (obj *_MacActorMgr) GetFromActorRemarks(actorRemarks string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_remarks` = ?", actorRemarks).Find(&results).Error

	return
}

// GetBatchFromActorRemarks 批量查找 演员备注
func (obj *_MacActorMgr) GetBatchFromActorRemarks(actorRemarkss []string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_remarks` IN (?)", actorRemarkss).Find(&results).Error

	return
}

// GetFromActorArea 通过actor_area获取内容 演员地区
func (obj *_MacActorMgr) GetFromActorArea(actorArea string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_area` = ?", actorArea).Find(&results).Error

	return
}

// GetBatchFromActorArea 批量查找 演员地区
func (obj *_MacActorMgr) GetBatchFromActorArea(actorAreas []string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_area` IN (?)", actorAreas).Find(&results).Error

	return
}

// GetFromActorHeight 通过actor_height获取内容 演员身高
func (obj *_MacActorMgr) GetFromActorHeight(actorHeight string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_height` = ?", actorHeight).Find(&results).Error

	return
}

// GetBatchFromActorHeight 批量查找 演员身高
func (obj *_MacActorMgr) GetBatchFromActorHeight(actorHeights []string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_height` IN (?)", actorHeights).Find(&results).Error

	return
}

// GetFromActorWeight 通过actor_weight获取内容 演员体重
func (obj *_MacActorMgr) GetFromActorWeight(actorWeight string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_weight` = ?", actorWeight).Find(&results).Error

	return
}

// GetBatchFromActorWeight 批量查找 演员体重
func (obj *_MacActorMgr) GetBatchFromActorWeight(actorWeights []string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_weight` IN (?)", actorWeights).Find(&results).Error

	return
}

// GetFromActorBirthday 通过actor_birthday获取内容 演员生日
func (obj *_MacActorMgr) GetFromActorBirthday(actorBirthday string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_birthday` = ?", actorBirthday).Find(&results).Error

	return
}

// GetBatchFromActorBirthday 批量查找 演员生日
func (obj *_MacActorMgr) GetBatchFromActorBirthday(actorBirthdays []string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_birthday` IN (?)", actorBirthdays).Find(&results).Error

	return
}

// GetFromActorBirtharea 通过actor_birtharea获取内容 演员出生地区
func (obj *_MacActorMgr) GetFromActorBirtharea(actorBirtharea string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_birtharea` = ?", actorBirtharea).Find(&results).Error

	return
}

// GetBatchFromActorBirtharea 批量查找 演员出生地区
func (obj *_MacActorMgr) GetBatchFromActorBirtharea(actorBirthareas []string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_birtharea` IN (?)", actorBirthareas).Find(&results).Error

	return
}

// GetFromActorBlood 通过actor_blood获取内容 演员血型
func (obj *_MacActorMgr) GetFromActorBlood(actorBlood string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_blood` = ?", actorBlood).Find(&results).Error

	return
}

// GetBatchFromActorBlood 批量查找 演员血型
func (obj *_MacActorMgr) GetBatchFromActorBlood(actorBloods []string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_blood` IN (?)", actorBloods).Find(&results).Error

	return
}

// GetFromActorStarsign 通过actor_starsign获取内容 演员星座
func (obj *_MacActorMgr) GetFromActorStarsign(actorStarsign string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_starsign` = ?", actorStarsign).Find(&results).Error

	return
}

// GetBatchFromActorStarsign 批量查找 演员星座
func (obj *_MacActorMgr) GetBatchFromActorStarsign(actorStarsigns []string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_starsign` IN (?)", actorStarsigns).Find(&results).Error

	return
}

// GetFromActorSchool 通过actor_school获取内容 演员毕业学校
func (obj *_MacActorMgr) GetFromActorSchool(actorSchool string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_school` = ?", actorSchool).Find(&results).Error

	return
}

// GetBatchFromActorSchool 批量查找 演员毕业学校
func (obj *_MacActorMgr) GetBatchFromActorSchool(actorSchools []string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_school` IN (?)", actorSchools).Find(&results).Error

	return
}

// GetFromActorWorks 通过actor_works获取内容 演员作品
func (obj *_MacActorMgr) GetFromActorWorks(actorWorks string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_works` = ?", actorWorks).Find(&results).Error

	return
}

// GetBatchFromActorWorks 批量查找 演员作品
func (obj *_MacActorMgr) GetBatchFromActorWorks(actorWorkss []string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_works` IN (?)", actorWorkss).Find(&results).Error

	return
}

// GetFromActorTag 通过actor_tag获取内容 演员标签
func (obj *_MacActorMgr) GetFromActorTag(actorTag string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_tag` = ?", actorTag).Find(&results).Error

	return
}

// GetBatchFromActorTag 批量查找 演员标签
func (obj *_MacActorMgr) GetBatchFromActorTag(actorTags []string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_tag` IN (?)", actorTags).Find(&results).Error

	return
}

// GetFromActorClass 通过actor_class获取内容 演员类别
func (obj *_MacActorMgr) GetFromActorClass(actorClass string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_class` = ?", actorClass).Find(&results).Error

	return
}

// GetBatchFromActorClass 批量查找 演员类别
func (obj *_MacActorMgr) GetBatchFromActorClass(actorClasss []string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_class` IN (?)", actorClasss).Find(&results).Error

	return
}

// GetFromActorLevel 通过actor_level获取内容 演员级别
func (obj *_MacActorMgr) GetFromActorLevel(actorLevel uint8) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_level` = ?", actorLevel).Find(&results).Error

	return
}

// GetBatchFromActorLevel 批量查找 演员级别
func (obj *_MacActorMgr) GetBatchFromActorLevel(actorLevels []uint8) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_level` IN (?)", actorLevels).Find(&results).Error

	return
}

// GetFromActorTime 通过actor_time获取内容 演员时间
func (obj *_MacActorMgr) GetFromActorTime(actorTime int) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_time` = ?", actorTime).Find(&results).Error

	return
}

// GetBatchFromActorTime 批量查找 演员时间
func (obj *_MacActorMgr) GetBatchFromActorTime(actorTimes []int) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_time` IN (?)", actorTimes).Find(&results).Error

	return
}

// GetFromActorTimeAdd 通过actor_time_add获取内容 演员添加时间
func (obj *_MacActorMgr) GetFromActorTimeAdd(actorTimeAdd int) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_time_add` = ?", actorTimeAdd).Find(&results).Error

	return
}

// GetBatchFromActorTimeAdd 批量查找 演员添加时间
func (obj *_MacActorMgr) GetBatchFromActorTimeAdd(actorTimeAdds []int) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_time_add` IN (?)", actorTimeAdds).Find(&results).Error

	return
}

// GetFromActorTimeHits 通过actor_time_hits获取内容 演员点击时间
func (obj *_MacActorMgr) GetFromActorTimeHits(actorTimeHits int) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_time_hits` = ?", actorTimeHits).Find(&results).Error

	return
}

// GetBatchFromActorTimeHits 批量查找 演员点击时间
func (obj *_MacActorMgr) GetBatchFromActorTimeHits(actorTimeHitss []int) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_time_hits` IN (?)", actorTimeHitss).Find(&results).Error

	return
}

// GetFromActorTimeMake 通过actor_time_make获取内容 演员制作时间
func (obj *_MacActorMgr) GetFromActorTimeMake(actorTimeMake int) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_time_make` = ?", actorTimeMake).Find(&results).Error

	return
}

// GetBatchFromActorTimeMake 批量查找 演员制作时间
func (obj *_MacActorMgr) GetBatchFromActorTimeMake(actorTimeMakes []int) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_time_make` IN (?)", actorTimeMakes).Find(&results).Error

	return
}

// GetFromActorHits 通过actor_hits获取内容 演员点击次数
func (obj *_MacActorMgr) GetFromActorHits(actorHits int16) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_hits` = ?", actorHits).Find(&results).Error

	return
}

// GetBatchFromActorHits 批量查找 演员点击次数
func (obj *_MacActorMgr) GetBatchFromActorHits(actorHitss []int16) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_hits` IN (?)", actorHitss).Find(&results).Error

	return
}

// GetFromActorHitsDay 通过actor_hits_day获取内容 演员每日点击次数
func (obj *_MacActorMgr) GetFromActorHitsDay(actorHitsDay int16) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_hits_day` = ?", actorHitsDay).Find(&results).Error

	return
}

// GetBatchFromActorHitsDay 批量查找 演员每日点击次数
func (obj *_MacActorMgr) GetBatchFromActorHitsDay(actorHitsDays []int16) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_hits_day` IN (?)", actorHitsDays).Find(&results).Error

	return
}

// GetFromActorHitsWeek 通过actor_hits_week获取内容 演员每周点击次数
func (obj *_MacActorMgr) GetFromActorHitsWeek(actorHitsWeek int16) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_hits_week` = ?", actorHitsWeek).Find(&results).Error

	return
}

// GetBatchFromActorHitsWeek 批量查找 演员每周点击次数
func (obj *_MacActorMgr) GetBatchFromActorHitsWeek(actorHitsWeeks []int16) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_hits_week` IN (?)", actorHitsWeeks).Find(&results).Error

	return
}

// GetFromActorHitsMonth 通过actor_hits_month获取内容 演员每月点击次数
func (obj *_MacActorMgr) GetFromActorHitsMonth(actorHitsMonth int16) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_hits_month` = ?", actorHitsMonth).Find(&results).Error

	return
}

// GetBatchFromActorHitsMonth 批量查找 演员每月点击次数
func (obj *_MacActorMgr) GetBatchFromActorHitsMonth(actorHitsMonths []int16) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_hits_month` IN (?)", actorHitsMonths).Find(&results).Error

	return
}

// GetFromActorScore 通过actor_score获取内容 演员评分
func (obj *_MacActorMgr) GetFromActorScore(actorScore float64) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_score` = ?", actorScore).Find(&results).Error

	return
}

// GetBatchFromActorScore 批量查找 演员评分
func (obj *_MacActorMgr) GetBatchFromActorScore(actorScores []float64) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_score` IN (?)", actorScores).Find(&results).Error

	return
}

// GetFromActorScoreAll 通过actor_score_all获取内容 演员总评分
func (obj *_MacActorMgr) GetFromActorScoreAll(actorScoreAll int16) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_score_all` = ?", actorScoreAll).Find(&results).Error

	return
}

// GetBatchFromActorScoreAll 批量查找 演员总评分
func (obj *_MacActorMgr) GetBatchFromActorScoreAll(actorScoreAlls []int16) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_score_all` IN (?)", actorScoreAlls).Find(&results).Error

	return
}

// GetFromActorScoreNum 通过actor_score_num获取内容 演员评分次数
func (obj *_MacActorMgr) GetFromActorScoreNum(actorScoreNum int16) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_score_num` = ?", actorScoreNum).Find(&results).Error

	return
}

// GetBatchFromActorScoreNum 批量查找 演员评分次数
func (obj *_MacActorMgr) GetBatchFromActorScoreNum(actorScoreNums []int16) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_score_num` IN (?)", actorScoreNums).Find(&results).Error

	return
}

// GetFromActorUp 通过actor_up获取内容 演员点赞次数
func (obj *_MacActorMgr) GetFromActorUp(actorUp int16) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_up` = ?", actorUp).Find(&results).Error

	return
}

// GetBatchFromActorUp 批量查找 演员点赞次数
func (obj *_MacActorMgr) GetBatchFromActorUp(actorUps []int16) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_up` IN (?)", actorUps).Find(&results).Error

	return
}

// GetFromActorDown 通过actor_down获取内容 演员踩次数
func (obj *_MacActorMgr) GetFromActorDown(actorDown int16) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_down` = ?", actorDown).Find(&results).Error

	return
}

// GetBatchFromActorDown 批量查找 演员踩次数
func (obj *_MacActorMgr) GetBatchFromActorDown(actorDowns []int16) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_down` IN (?)", actorDowns).Find(&results).Error

	return
}

// GetFromActorTpl 通过actor_tpl获取内容 演员模板
func (obj *_MacActorMgr) GetFromActorTpl(actorTpl string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_tpl` = ?", actorTpl).Find(&results).Error

	return
}

// GetBatchFromActorTpl 批量查找 演员模板
func (obj *_MacActorMgr) GetBatchFromActorTpl(actorTpls []string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_tpl` IN (?)", actorTpls).Find(&results).Error

	return
}

// GetFromActorJumpurl 通过actor_jumpurl获取内容 演员跳转链接
func (obj *_MacActorMgr) GetFromActorJumpurl(actorJumpurl string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_jumpurl` = ?", actorJumpurl).Find(&results).Error

	return
}

// GetBatchFromActorJumpurl 批量查找 演员跳转链接
func (obj *_MacActorMgr) GetBatchFromActorJumpurl(actorJumpurls []string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_jumpurl` IN (?)", actorJumpurls).Find(&results).Error

	return
}

// GetFromActorContent 通过actor_content获取内容 演员内容
func (obj *_MacActorMgr) GetFromActorContent(actorContent string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_content` = ?", actorContent).Find(&results).Error

	return
}

// GetBatchFromActorContent 批量查找 演员内容
func (obj *_MacActorMgr) GetBatchFromActorContent(actorContents []string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_content` IN (?)", actorContents).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MacActorMgr) FetchByPrimaryKey(actorID int) (result MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_id` = ?", actorID).First(&result).Error

	return
}

// FetchIndexByTypeID  获取多个内容
func (obj *_MacActorMgr) FetchIndexByTypeID(typeID uint16) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`type_id` = ?", typeID).Find(&results).Error

	return
}

// FetchIndexByTypeID1  获取多个内容
func (obj *_MacActorMgr) FetchIndexByTypeID1(typeID1 uint16) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`type_id_1` = ?", typeID1).Find(&results).Error

	return
}

// FetchIndexByActorName  获取多个内容
func (obj *_MacActorMgr) FetchIndexByActorName(actorName string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_name` = ?", actorName).Find(&results).Error

	return
}

// FetchIndexByActorEn  获取多个内容
func (obj *_MacActorMgr) FetchIndexByActorEn(actorEn string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_en` = ?", actorEn).Find(&results).Error

	return
}

// FetchIndexByActorLetter  获取多个内容
func (obj *_MacActorMgr) FetchIndexByActorLetter(actorLetter string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_letter` = ?", actorLetter).Find(&results).Error

	return
}

// FetchIndexByActorSex  获取多个内容
func (obj *_MacActorMgr) FetchIndexByActorSex(actorSex string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_sex` = ?", actorSex).Find(&results).Error

	return
}

// FetchIndexByActorArea  获取多个内容
func (obj *_MacActorMgr) FetchIndexByActorArea(actorArea string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_area` = ?", actorArea).Find(&results).Error

	return
}

// FetchIndexByActorTag  获取多个内容
func (obj *_MacActorMgr) FetchIndexByActorTag(actorTag string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_tag` = ?", actorTag).Find(&results).Error

	return
}

// FetchIndexByActorClass  获取多个内容
func (obj *_MacActorMgr) FetchIndexByActorClass(actorClass string) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_class` = ?", actorClass).Find(&results).Error

	return
}

// FetchIndexByActorLevel  获取多个内容
func (obj *_MacActorMgr) FetchIndexByActorLevel(actorLevel uint8) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_level` = ?", actorLevel).Find(&results).Error

	return
}

// FetchIndexByActorTime  获取多个内容
func (obj *_MacActorMgr) FetchIndexByActorTime(actorTime int) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_time` = ?", actorTime).Find(&results).Error

	return
}

// FetchIndexByActorTimeAdd  获取多个内容
func (obj *_MacActorMgr) FetchIndexByActorTimeAdd(actorTimeAdd int) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_time_add` = ?", actorTimeAdd).Find(&results).Error

	return
}

// FetchIndexByActorScore  获取多个内容
func (obj *_MacActorMgr) FetchIndexByActorScore(actorScore float64) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_score` = ?", actorScore).Find(&results).Error

	return
}

// FetchIndexByActorScoreAll  获取多个内容
func (obj *_MacActorMgr) FetchIndexByActorScoreAll(actorScoreAll int16) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_score_all` = ?", actorScoreAll).Find(&results).Error

	return
}

// FetchIndexByActorScoreNum  获取多个内容
func (obj *_MacActorMgr) FetchIndexByActorScoreNum(actorScoreNum int16) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_score_num` = ?", actorScoreNum).Find(&results).Error

	return
}

// FetchIndexByActorUp  获取多个内容
func (obj *_MacActorMgr) FetchIndexByActorUp(actorUp int16) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_up` = ?", actorUp).Find(&results).Error

	return
}

// FetchIndexByActorDown  获取多个内容
func (obj *_MacActorMgr) FetchIndexByActorDown(actorDown int16) (results []*MacActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacActor{}).Where("`actor_down` = ?", actorDown).Find(&results).Error

	return
}
