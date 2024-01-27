package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _MacTopicMgr struct {
	*_BaseMgr
}

// MacTopicMgr open func
func MacTopicMgr(db *gorm.DB) *_MacTopicMgr {
	if db == nil {
		panic(fmt.Errorf("MacTopicMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MacTopicMgr{_BaseMgr: &_BaseMgr{DB: db.Table("mac_topic"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_MacTopicMgr) Debug() *_MacTopicMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MacTopicMgr) GetTableName() string {
	return "mac_topic"
}

// Reset 重置gorm会话
func (obj *_MacTopicMgr) Reset() *_MacTopicMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MacTopicMgr) Get() (result MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MacTopicMgr) Gets() (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MacTopicMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTopicID topic_id获取 专题ID
func (obj *_MacTopicMgr) WithTopicID(topicID uint16) Option {
	return optionFunc(func(o *options) { o.query["topic_id"] = topicID })
}

// WithTopicName topic_name获取 专题名称
func (obj *_MacTopicMgr) WithTopicName(topicName string) Option {
	return optionFunc(func(o *options) { o.query["topic_name"] = topicName })
}

// WithTopicEn topic_en获取 专题英文名
func (obj *_MacTopicMgr) WithTopicEn(topicEn string) Option {
	return optionFunc(func(o *options) { o.query["topic_en"] = topicEn })
}

// WithTopicSub topic_sub获取 专题副标题
func (obj *_MacTopicMgr) WithTopicSub(topicSub string) Option {
	return optionFunc(func(o *options) { o.query["topic_sub"] = topicSub })
}

// WithTopicStatus topic_status获取 专题状态
func (obj *_MacTopicMgr) WithTopicStatus(topicStatus uint8) Option {
	return optionFunc(func(o *options) { o.query["topic_status"] = topicStatus })
}

// WithTopicSort topic_sort获取 专题排序
func (obj *_MacTopicMgr) WithTopicSort(topicSort uint16) Option {
	return optionFunc(func(o *options) { o.query["topic_sort"] = topicSort })
}

// WithTopicLetter topic_letter获取 专题首字母
func (obj *_MacTopicMgr) WithTopicLetter(topicLetter string) Option {
	return optionFunc(func(o *options) { o.query["topic_letter"] = topicLetter })
}

// WithTopicColor topic_color获取 专题颜色
func (obj *_MacTopicMgr) WithTopicColor(topicColor string) Option {
	return optionFunc(func(o *options) { o.query["topic_color"] = topicColor })
}

// WithTopicTpl topic_tpl获取 专题模板
func (obj *_MacTopicMgr) WithTopicTpl(topicTpl string) Option {
	return optionFunc(func(o *options) { o.query["topic_tpl"] = topicTpl })
}

// WithTopicType topic_type获取 专题类型
func (obj *_MacTopicMgr) WithTopicType(topicType string) Option {
	return optionFunc(func(o *options) { o.query["topic_type"] = topicType })
}

// WithTopicPic topic_pic获取 专题图片
func (obj *_MacTopicMgr) WithTopicPic(topicPic string) Option {
	return optionFunc(func(o *options) { o.query["topic_pic"] = topicPic })
}

// WithTopicPicThumb topic_pic_thumb获取 专题缩略图
func (obj *_MacTopicMgr) WithTopicPicThumb(topicPicThumb string) Option {
	return optionFunc(func(o *options) { o.query["topic_pic_thumb"] = topicPicThumb })
}

// WithTopicPicSlide topic_pic_slide获取 专题幻灯片
func (obj *_MacTopicMgr) WithTopicPicSlide(topicPicSlide string) Option {
	return optionFunc(func(o *options) { o.query["topic_pic_slide"] = topicPicSlide })
}

// WithTopicKey topic_key获取 专题关键字
func (obj *_MacTopicMgr) WithTopicKey(topicKey string) Option {
	return optionFunc(func(o *options) { o.query["topic_key"] = topicKey })
}

// WithTopicDes topic_des获取 专题描述
func (obj *_MacTopicMgr) WithTopicDes(topicDes string) Option {
	return optionFunc(func(o *options) { o.query["topic_des"] = topicDes })
}

// WithTopicTitle topic_title获取 专题标题
func (obj *_MacTopicMgr) WithTopicTitle(topicTitle string) Option {
	return optionFunc(func(o *options) { o.query["topic_title"] = topicTitle })
}

// WithTopicBlurb topic_blurb获取 专题简介
func (obj *_MacTopicMgr) WithTopicBlurb(topicBlurb string) Option {
	return optionFunc(func(o *options) { o.query["topic_blurb"] = topicBlurb })
}

// WithTopicRemarks topic_remarks获取 专题备注
func (obj *_MacTopicMgr) WithTopicRemarks(topicRemarks string) Option {
	return optionFunc(func(o *options) { o.query["topic_remarks"] = topicRemarks })
}

// WithTopicLevel topic_level获取 专题级别
func (obj *_MacTopicMgr) WithTopicLevel(topicLevel uint8) Option {
	return optionFunc(func(o *options) { o.query["topic_level"] = topicLevel })
}

// WithTopicUp topic_up获取 专题点赞数
func (obj *_MacTopicMgr) WithTopicUp(topicUp int16) Option {
	return optionFunc(func(o *options) { o.query["topic_up"] = topicUp })
}

// WithTopicDown topic_down获取 专题踩数
func (obj *_MacTopicMgr) WithTopicDown(topicDown int16) Option {
	return optionFunc(func(o *options) { o.query["topic_down"] = topicDown })
}

// WithTopicScore topic_score获取 专题评分
func (obj *_MacTopicMgr) WithTopicScore(topicScore float64) Option {
	return optionFunc(func(o *options) { o.query["topic_score"] = topicScore })
}

// WithTopicScoreAll topic_score_all获取 专题总评分
func (obj *_MacTopicMgr) WithTopicScoreAll(topicScoreAll int16) Option {
	return optionFunc(func(o *options) { o.query["topic_score_all"] = topicScoreAll })
}

// WithTopicScoreNum topic_score_num获取 专题评分次数
func (obj *_MacTopicMgr) WithTopicScoreNum(topicScoreNum int16) Option {
	return optionFunc(func(o *options) { o.query["topic_score_num"] = topicScoreNum })
}

// WithTopicHits topic_hits获取 专题点击数
func (obj *_MacTopicMgr) WithTopicHits(topicHits int16) Option {
	return optionFunc(func(o *options) { o.query["topic_hits"] = topicHits })
}

// WithTopicHitsDay topic_hits_day获取 专题日点击数
func (obj *_MacTopicMgr) WithTopicHitsDay(topicHitsDay int16) Option {
	return optionFunc(func(o *options) { o.query["topic_hits_day"] = topicHitsDay })
}

// WithTopicHitsWeek topic_hits_week获取 专题周点击数
func (obj *_MacTopicMgr) WithTopicHitsWeek(topicHitsWeek int16) Option {
	return optionFunc(func(o *options) { o.query["topic_hits_week"] = topicHitsWeek })
}

// WithTopicHitsMonth topic_hits_month获取 专题月点击数
func (obj *_MacTopicMgr) WithTopicHitsMonth(topicHitsMonth int16) Option {
	return optionFunc(func(o *options) { o.query["topic_hits_month"] = topicHitsMonth })
}

// WithTopicTime topic_time获取 专题时间
func (obj *_MacTopicMgr) WithTopicTime(topicTime int) Option {
	return optionFunc(func(o *options) { o.query["topic_time"] = topicTime })
}

// WithTopicTimeAdd topic_time_add获取 专题添加时间
func (obj *_MacTopicMgr) WithTopicTimeAdd(topicTimeAdd int) Option {
	return optionFunc(func(o *options) { o.query["topic_time_add"] = topicTimeAdd })
}

// WithTopicTimeHits topic_time_hits获取 专题点击时间
func (obj *_MacTopicMgr) WithTopicTimeHits(topicTimeHits int) Option {
	return optionFunc(func(o *options) { o.query["topic_time_hits"] = topicTimeHits })
}

// WithTopicTimeMake topic_time_make获取 专题制作时间
func (obj *_MacTopicMgr) WithTopicTimeMake(topicTimeMake int) Option {
	return optionFunc(func(o *options) { o.query["topic_time_make"] = topicTimeMake })
}

// WithTopicTag topic_tag获取 专题标签
func (obj *_MacTopicMgr) WithTopicTag(topicTag string) Option {
	return optionFunc(func(o *options) { o.query["topic_tag"] = topicTag })
}

// WithTopicRelVod topic_rel_vod获取 专题关联视频
func (obj *_MacTopicMgr) WithTopicRelVod(topicRelVod string) Option {
	return optionFunc(func(o *options) { o.query["topic_rel_vod"] = topicRelVod })
}

// WithTopicRelArt topic_rel_art获取 专题关联文章
func (obj *_MacTopicMgr) WithTopicRelArt(topicRelArt string) Option {
	return optionFunc(func(o *options) { o.query["topic_rel_art"] = topicRelArt })
}

// WithTopicContent topic_content获取 专题内容
func (obj *_MacTopicMgr) WithTopicContent(topicContent string) Option {
	return optionFunc(func(o *options) { o.query["topic_content"] = topicContent })
}

// WithTopicExtend topic_extend获取 专题扩展信息
func (obj *_MacTopicMgr) WithTopicExtend(topicExtend string) Option {
	return optionFunc(func(o *options) { o.query["topic_extend"] = topicExtend })
}

// GetByOption 功能选项模式获取
func (obj *_MacTopicMgr) GetByOption(opts ...Option) (result MacTopic, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MacTopicMgr) GetByOptions(opts ...Option) (results []*MacTopic, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MacTopicMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MacTopic, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where(options.query)
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

// GetFromTopicID 通过topic_id获取内容 专题ID
func (obj *_MacTopicMgr) GetFromTopicID(topicID uint16) (result MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_id` = ?", topicID).First(&result).Error

	return
}

// GetBatchFromTopicID 批量查找 专题ID
func (obj *_MacTopicMgr) GetBatchFromTopicID(topicIDs []uint16) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_id` IN (?)", topicIDs).Find(&results).Error

	return
}

// GetFromTopicName 通过topic_name获取内容 专题名称
func (obj *_MacTopicMgr) GetFromTopicName(topicName string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_name` = ?", topicName).Find(&results).Error

	return
}

// GetBatchFromTopicName 批量查找 专题名称
func (obj *_MacTopicMgr) GetBatchFromTopicName(topicNames []string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_name` IN (?)", topicNames).Find(&results).Error

	return
}

// GetFromTopicEn 通过topic_en获取内容 专题英文名
func (obj *_MacTopicMgr) GetFromTopicEn(topicEn string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_en` = ?", topicEn).Find(&results).Error

	return
}

// GetBatchFromTopicEn 批量查找 专题英文名
func (obj *_MacTopicMgr) GetBatchFromTopicEn(topicEns []string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_en` IN (?)", topicEns).Find(&results).Error

	return
}

// GetFromTopicSub 通过topic_sub获取内容 专题副标题
func (obj *_MacTopicMgr) GetFromTopicSub(topicSub string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_sub` = ?", topicSub).Find(&results).Error

	return
}

// GetBatchFromTopicSub 批量查找 专题副标题
func (obj *_MacTopicMgr) GetBatchFromTopicSub(topicSubs []string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_sub` IN (?)", topicSubs).Find(&results).Error

	return
}

// GetFromTopicStatus 通过topic_status获取内容 专题状态
func (obj *_MacTopicMgr) GetFromTopicStatus(topicStatus uint8) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_status` = ?", topicStatus).Find(&results).Error

	return
}

// GetBatchFromTopicStatus 批量查找 专题状态
func (obj *_MacTopicMgr) GetBatchFromTopicStatus(topicStatuss []uint8) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_status` IN (?)", topicStatuss).Find(&results).Error

	return
}

// GetFromTopicSort 通过topic_sort获取内容 专题排序
func (obj *_MacTopicMgr) GetFromTopicSort(topicSort uint16) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_sort` = ?", topicSort).Find(&results).Error

	return
}

// GetBatchFromTopicSort 批量查找 专题排序
func (obj *_MacTopicMgr) GetBatchFromTopicSort(topicSorts []uint16) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_sort` IN (?)", topicSorts).Find(&results).Error

	return
}

// GetFromTopicLetter 通过topic_letter获取内容 专题首字母
func (obj *_MacTopicMgr) GetFromTopicLetter(topicLetter string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_letter` = ?", topicLetter).Find(&results).Error

	return
}

// GetBatchFromTopicLetter 批量查找 专题首字母
func (obj *_MacTopicMgr) GetBatchFromTopicLetter(topicLetters []string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_letter` IN (?)", topicLetters).Find(&results).Error

	return
}

// GetFromTopicColor 通过topic_color获取内容 专题颜色
func (obj *_MacTopicMgr) GetFromTopicColor(topicColor string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_color` = ?", topicColor).Find(&results).Error

	return
}

// GetBatchFromTopicColor 批量查找 专题颜色
func (obj *_MacTopicMgr) GetBatchFromTopicColor(topicColors []string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_color` IN (?)", topicColors).Find(&results).Error

	return
}

// GetFromTopicTpl 通过topic_tpl获取内容 专题模板
func (obj *_MacTopicMgr) GetFromTopicTpl(topicTpl string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_tpl` = ?", topicTpl).Find(&results).Error

	return
}

// GetBatchFromTopicTpl 批量查找 专题模板
func (obj *_MacTopicMgr) GetBatchFromTopicTpl(topicTpls []string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_tpl` IN (?)", topicTpls).Find(&results).Error

	return
}

// GetFromTopicType 通过topic_type获取内容 专题类型
func (obj *_MacTopicMgr) GetFromTopicType(topicType string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_type` = ?", topicType).Find(&results).Error

	return
}

// GetBatchFromTopicType 批量查找 专题类型
func (obj *_MacTopicMgr) GetBatchFromTopicType(topicTypes []string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_type` IN (?)", topicTypes).Find(&results).Error

	return
}

// GetFromTopicPic 通过topic_pic获取内容 专题图片
func (obj *_MacTopicMgr) GetFromTopicPic(topicPic string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_pic` = ?", topicPic).Find(&results).Error

	return
}

// GetBatchFromTopicPic 批量查找 专题图片
func (obj *_MacTopicMgr) GetBatchFromTopicPic(topicPics []string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_pic` IN (?)", topicPics).Find(&results).Error

	return
}

// GetFromTopicPicThumb 通过topic_pic_thumb获取内容 专题缩略图
func (obj *_MacTopicMgr) GetFromTopicPicThumb(topicPicThumb string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_pic_thumb` = ?", topicPicThumb).Find(&results).Error

	return
}

// GetBatchFromTopicPicThumb 批量查找 专题缩略图
func (obj *_MacTopicMgr) GetBatchFromTopicPicThumb(topicPicThumbs []string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_pic_thumb` IN (?)", topicPicThumbs).Find(&results).Error

	return
}

// GetFromTopicPicSlide 通过topic_pic_slide获取内容 专题幻灯片
func (obj *_MacTopicMgr) GetFromTopicPicSlide(topicPicSlide string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_pic_slide` = ?", topicPicSlide).Find(&results).Error

	return
}

// GetBatchFromTopicPicSlide 批量查找 专题幻灯片
func (obj *_MacTopicMgr) GetBatchFromTopicPicSlide(topicPicSlides []string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_pic_slide` IN (?)", topicPicSlides).Find(&results).Error

	return
}

// GetFromTopicKey 通过topic_key获取内容 专题关键字
func (obj *_MacTopicMgr) GetFromTopicKey(topicKey string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_key` = ?", topicKey).Find(&results).Error

	return
}

// GetBatchFromTopicKey 批量查找 专题关键字
func (obj *_MacTopicMgr) GetBatchFromTopicKey(topicKeys []string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_key` IN (?)", topicKeys).Find(&results).Error

	return
}

// GetFromTopicDes 通过topic_des获取内容 专题描述
func (obj *_MacTopicMgr) GetFromTopicDes(topicDes string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_des` = ?", topicDes).Find(&results).Error

	return
}

// GetBatchFromTopicDes 批量查找 专题描述
func (obj *_MacTopicMgr) GetBatchFromTopicDes(topicDess []string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_des` IN (?)", topicDess).Find(&results).Error

	return
}

// GetFromTopicTitle 通过topic_title获取内容 专题标题
func (obj *_MacTopicMgr) GetFromTopicTitle(topicTitle string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_title` = ?", topicTitle).Find(&results).Error

	return
}

// GetBatchFromTopicTitle 批量查找 专题标题
func (obj *_MacTopicMgr) GetBatchFromTopicTitle(topicTitles []string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_title` IN (?)", topicTitles).Find(&results).Error

	return
}

// GetFromTopicBlurb 通过topic_blurb获取内容 专题简介
func (obj *_MacTopicMgr) GetFromTopicBlurb(topicBlurb string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_blurb` = ?", topicBlurb).Find(&results).Error

	return
}

// GetBatchFromTopicBlurb 批量查找 专题简介
func (obj *_MacTopicMgr) GetBatchFromTopicBlurb(topicBlurbs []string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_blurb` IN (?)", topicBlurbs).Find(&results).Error

	return
}

// GetFromTopicRemarks 通过topic_remarks获取内容 专题备注
func (obj *_MacTopicMgr) GetFromTopicRemarks(topicRemarks string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_remarks` = ?", topicRemarks).Find(&results).Error

	return
}

// GetBatchFromTopicRemarks 批量查找 专题备注
func (obj *_MacTopicMgr) GetBatchFromTopicRemarks(topicRemarkss []string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_remarks` IN (?)", topicRemarkss).Find(&results).Error

	return
}

// GetFromTopicLevel 通过topic_level获取内容 专题级别
func (obj *_MacTopicMgr) GetFromTopicLevel(topicLevel uint8) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_level` = ?", topicLevel).Find(&results).Error

	return
}

// GetBatchFromTopicLevel 批量查找 专题级别
func (obj *_MacTopicMgr) GetBatchFromTopicLevel(topicLevels []uint8) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_level` IN (?)", topicLevels).Find(&results).Error

	return
}

// GetFromTopicUp 通过topic_up获取内容 专题点赞数
func (obj *_MacTopicMgr) GetFromTopicUp(topicUp int16) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_up` = ?", topicUp).Find(&results).Error

	return
}

// GetBatchFromTopicUp 批量查找 专题点赞数
func (obj *_MacTopicMgr) GetBatchFromTopicUp(topicUps []int16) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_up` IN (?)", topicUps).Find(&results).Error

	return
}

// GetFromTopicDown 通过topic_down获取内容 专题踩数
func (obj *_MacTopicMgr) GetFromTopicDown(topicDown int16) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_down` = ?", topicDown).Find(&results).Error

	return
}

// GetBatchFromTopicDown 批量查找 专题踩数
func (obj *_MacTopicMgr) GetBatchFromTopicDown(topicDowns []int16) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_down` IN (?)", topicDowns).Find(&results).Error

	return
}

// GetFromTopicScore 通过topic_score获取内容 专题评分
func (obj *_MacTopicMgr) GetFromTopicScore(topicScore float64) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_score` = ?", topicScore).Find(&results).Error

	return
}

// GetBatchFromTopicScore 批量查找 专题评分
func (obj *_MacTopicMgr) GetBatchFromTopicScore(topicScores []float64) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_score` IN (?)", topicScores).Find(&results).Error

	return
}

// GetFromTopicScoreAll 通过topic_score_all获取内容 专题总评分
func (obj *_MacTopicMgr) GetFromTopicScoreAll(topicScoreAll int16) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_score_all` = ?", topicScoreAll).Find(&results).Error

	return
}

// GetBatchFromTopicScoreAll 批量查找 专题总评分
func (obj *_MacTopicMgr) GetBatchFromTopicScoreAll(topicScoreAlls []int16) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_score_all` IN (?)", topicScoreAlls).Find(&results).Error

	return
}

// GetFromTopicScoreNum 通过topic_score_num获取内容 专题评分次数
func (obj *_MacTopicMgr) GetFromTopicScoreNum(topicScoreNum int16) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_score_num` = ?", topicScoreNum).Find(&results).Error

	return
}

// GetBatchFromTopicScoreNum 批量查找 专题评分次数
func (obj *_MacTopicMgr) GetBatchFromTopicScoreNum(topicScoreNums []int16) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_score_num` IN (?)", topicScoreNums).Find(&results).Error

	return
}

// GetFromTopicHits 通过topic_hits获取内容 专题点击数
func (obj *_MacTopicMgr) GetFromTopicHits(topicHits int16) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_hits` = ?", topicHits).Find(&results).Error

	return
}

// GetBatchFromTopicHits 批量查找 专题点击数
func (obj *_MacTopicMgr) GetBatchFromTopicHits(topicHitss []int16) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_hits` IN (?)", topicHitss).Find(&results).Error

	return
}

// GetFromTopicHitsDay 通过topic_hits_day获取内容 专题日点击数
func (obj *_MacTopicMgr) GetFromTopicHitsDay(topicHitsDay int16) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_hits_day` = ?", topicHitsDay).Find(&results).Error

	return
}

// GetBatchFromTopicHitsDay 批量查找 专题日点击数
func (obj *_MacTopicMgr) GetBatchFromTopicHitsDay(topicHitsDays []int16) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_hits_day` IN (?)", topicHitsDays).Find(&results).Error

	return
}

// GetFromTopicHitsWeek 通过topic_hits_week获取内容 专题周点击数
func (obj *_MacTopicMgr) GetFromTopicHitsWeek(topicHitsWeek int16) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_hits_week` = ?", topicHitsWeek).Find(&results).Error

	return
}

// GetBatchFromTopicHitsWeek 批量查找 专题周点击数
func (obj *_MacTopicMgr) GetBatchFromTopicHitsWeek(topicHitsWeeks []int16) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_hits_week` IN (?)", topicHitsWeeks).Find(&results).Error

	return
}

// GetFromTopicHitsMonth 通过topic_hits_month获取内容 专题月点击数
func (obj *_MacTopicMgr) GetFromTopicHitsMonth(topicHitsMonth int16) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_hits_month` = ?", topicHitsMonth).Find(&results).Error

	return
}

// GetBatchFromTopicHitsMonth 批量查找 专题月点击数
func (obj *_MacTopicMgr) GetBatchFromTopicHitsMonth(topicHitsMonths []int16) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_hits_month` IN (?)", topicHitsMonths).Find(&results).Error

	return
}

// GetFromTopicTime 通过topic_time获取内容 专题时间
func (obj *_MacTopicMgr) GetFromTopicTime(topicTime int) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_time` = ?", topicTime).Find(&results).Error

	return
}

// GetBatchFromTopicTime 批量查找 专题时间
func (obj *_MacTopicMgr) GetBatchFromTopicTime(topicTimes []int) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_time` IN (?)", topicTimes).Find(&results).Error

	return
}

// GetFromTopicTimeAdd 通过topic_time_add获取内容 专题添加时间
func (obj *_MacTopicMgr) GetFromTopicTimeAdd(topicTimeAdd int) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_time_add` = ?", topicTimeAdd).Find(&results).Error

	return
}

// GetBatchFromTopicTimeAdd 批量查找 专题添加时间
func (obj *_MacTopicMgr) GetBatchFromTopicTimeAdd(topicTimeAdds []int) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_time_add` IN (?)", topicTimeAdds).Find(&results).Error

	return
}

// GetFromTopicTimeHits 通过topic_time_hits获取内容 专题点击时间
func (obj *_MacTopicMgr) GetFromTopicTimeHits(topicTimeHits int) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_time_hits` = ?", topicTimeHits).Find(&results).Error

	return
}

// GetBatchFromTopicTimeHits 批量查找 专题点击时间
func (obj *_MacTopicMgr) GetBatchFromTopicTimeHits(topicTimeHitss []int) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_time_hits` IN (?)", topicTimeHitss).Find(&results).Error

	return
}

// GetFromTopicTimeMake 通过topic_time_make获取内容 专题制作时间
func (obj *_MacTopicMgr) GetFromTopicTimeMake(topicTimeMake int) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_time_make` = ?", topicTimeMake).Find(&results).Error

	return
}

// GetBatchFromTopicTimeMake 批量查找 专题制作时间
func (obj *_MacTopicMgr) GetBatchFromTopicTimeMake(topicTimeMakes []int) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_time_make` IN (?)", topicTimeMakes).Find(&results).Error

	return
}

// GetFromTopicTag 通过topic_tag获取内容 专题标签
func (obj *_MacTopicMgr) GetFromTopicTag(topicTag string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_tag` = ?", topicTag).Find(&results).Error

	return
}

// GetBatchFromTopicTag 批量查找 专题标签
func (obj *_MacTopicMgr) GetBatchFromTopicTag(topicTags []string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_tag` IN (?)", topicTags).Find(&results).Error

	return
}

// GetFromTopicRelVod 通过topic_rel_vod获取内容 专题关联视频
func (obj *_MacTopicMgr) GetFromTopicRelVod(topicRelVod string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_rel_vod` = ?", topicRelVod).Find(&results).Error

	return
}

// GetBatchFromTopicRelVod 批量查找 专题关联视频
func (obj *_MacTopicMgr) GetBatchFromTopicRelVod(topicRelVods []string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_rel_vod` IN (?)", topicRelVods).Find(&results).Error

	return
}

// GetFromTopicRelArt 通过topic_rel_art获取内容 专题关联文章
func (obj *_MacTopicMgr) GetFromTopicRelArt(topicRelArt string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_rel_art` = ?", topicRelArt).Find(&results).Error

	return
}

// GetBatchFromTopicRelArt 批量查找 专题关联文章
func (obj *_MacTopicMgr) GetBatchFromTopicRelArt(topicRelArts []string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_rel_art` IN (?)", topicRelArts).Find(&results).Error

	return
}

// GetFromTopicContent 通过topic_content获取内容 专题内容
func (obj *_MacTopicMgr) GetFromTopicContent(topicContent string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_content` = ?", topicContent).Find(&results).Error

	return
}

// GetBatchFromTopicContent 批量查找 专题内容
func (obj *_MacTopicMgr) GetBatchFromTopicContent(topicContents []string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_content` IN (?)", topicContents).Find(&results).Error

	return
}

// GetFromTopicExtend 通过topic_extend获取内容 专题扩展信息
func (obj *_MacTopicMgr) GetFromTopicExtend(topicExtend string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_extend` = ?", topicExtend).Find(&results).Error

	return
}

// GetBatchFromTopicExtend 批量查找 专题扩展信息
func (obj *_MacTopicMgr) GetBatchFromTopicExtend(topicExtends []string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_extend` IN (?)", topicExtends).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MacTopicMgr) FetchByPrimaryKey(topicID uint16) (result MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_id` = ?", topicID).First(&result).Error

	return
}

// FetchIndexByTopicName  获取多个内容
func (obj *_MacTopicMgr) FetchIndexByTopicName(topicName string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_name` = ?", topicName).Find(&results).Error

	return
}

// FetchIndexByTopicEn  获取多个内容
func (obj *_MacTopicMgr) FetchIndexByTopicEn(topicEn string) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_en` = ?", topicEn).Find(&results).Error

	return
}

// FetchIndexByTopicSort  获取多个内容
func (obj *_MacTopicMgr) FetchIndexByTopicSort(topicSort uint16) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_sort` = ?", topicSort).Find(&results).Error

	return
}

// FetchIndexByTopicLevel  获取多个内容
func (obj *_MacTopicMgr) FetchIndexByTopicLevel(topicLevel uint8) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_level` = ?", topicLevel).Find(&results).Error

	return
}

// FetchIndexByTopicUp  获取多个内容
func (obj *_MacTopicMgr) FetchIndexByTopicUp(topicUp int16) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_up` = ?", topicUp).Find(&results).Error

	return
}

// FetchIndexByTopicDown  获取多个内容
func (obj *_MacTopicMgr) FetchIndexByTopicDown(topicDown int16) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_down` = ?", topicDown).Find(&results).Error

	return
}

// FetchIndexByTopicScore  获取多个内容
func (obj *_MacTopicMgr) FetchIndexByTopicScore(topicScore float64) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_score` = ?", topicScore).Find(&results).Error

	return
}

// FetchIndexByTopicScoreAll  获取多个内容
func (obj *_MacTopicMgr) FetchIndexByTopicScoreAll(topicScoreAll int16) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_score_all` = ?", topicScoreAll).Find(&results).Error

	return
}

// FetchIndexByTopicScoreNum  获取多个内容
func (obj *_MacTopicMgr) FetchIndexByTopicScoreNum(topicScoreNum int16) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_score_num` = ?", topicScoreNum).Find(&results).Error

	return
}

// FetchIndexByTopicHits  获取多个内容
func (obj *_MacTopicMgr) FetchIndexByTopicHits(topicHits int16) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_hits` = ?", topicHits).Find(&results).Error

	return
}

// FetchIndexByTopicHitsDay  获取多个内容
func (obj *_MacTopicMgr) FetchIndexByTopicHitsDay(topicHitsDay int16) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_hits_day` = ?", topicHitsDay).Find(&results).Error

	return
}

// FetchIndexByTopicHitsWeek  获取多个内容
func (obj *_MacTopicMgr) FetchIndexByTopicHitsWeek(topicHitsWeek int16) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_hits_week` = ?", topicHitsWeek).Find(&results).Error

	return
}

// FetchIndexByTopicHitsMonth  获取多个内容
func (obj *_MacTopicMgr) FetchIndexByTopicHitsMonth(topicHitsMonth int16) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_hits_month` = ?", topicHitsMonth).Find(&results).Error

	return
}

// FetchIndexByTopicTime  获取多个内容
func (obj *_MacTopicMgr) FetchIndexByTopicTime(topicTime int) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_time` = ?", topicTime).Find(&results).Error

	return
}

// FetchIndexByTopicTimeAdd  获取多个内容
func (obj *_MacTopicMgr) FetchIndexByTopicTimeAdd(topicTimeAdd int) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_time_add` = ?", topicTimeAdd).Find(&results).Error

	return
}

// FetchIndexByTopicTimeHits  获取多个内容
func (obj *_MacTopicMgr) FetchIndexByTopicTimeHits(topicTimeHits int) (results []*MacTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacTopic{}).Where("`topic_time_hits` = ?", topicTimeHits).Find(&results).Error

	return
}
