package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _MacVodMgr struct {
	*_BaseMgr
}

// MacVodMgr open func
func MacVodMgr(db *gorm.DB) *_MacVodMgr {
	if db == nil {
		panic(fmt.Errorf("MacVodMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MacVodMgr{_BaseMgr: &_BaseMgr{DB: db.Table("mac_vod"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_MacVodMgr) Debug() *_MacVodMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MacVodMgr) GetTableName() string {
	return "mac_vod"
}

// Reset 重置gorm会话
func (obj *_MacVodMgr) Reset() *_MacVodMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MacVodMgr) Get() (result MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MacVodMgr) Gets() (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MacVodMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MacVod{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithVodID vod_id获取 视频ID
func (obj *_MacVodMgr) WithVodID(vodID int) Option {
	return optionFunc(func(o *options) { o.query["vod_id"] = vodID })
}

// WithTypeID type_id获取 栏目ID
func (obj *_MacVodMgr) WithTypeID(typeID int16) Option {
	return optionFunc(func(o *options) { o.query["type_id"] = typeID })
}

// WithTypeID1 type_id_1获取 子栏目ID
func (obj *_MacVodMgr) WithTypeID1(typeID1 uint16) Option {
	return optionFunc(func(o *options) { o.query["type_id_1"] = typeID1 })
}

// WithGroupID group_id获取 分组ID
func (obj *_MacVodMgr) WithGroupID(groupID uint16) Option {
	return optionFunc(func(o *options) { o.query["group_id"] = groupID })
}

// WithVodName vod_name获取 视频标题
func (obj *_MacVodMgr) WithVodName(vodName string) Option {
	return optionFunc(func(o *options) { o.query["vod_name"] = vodName })
}

// WithVodSub vod_sub获取 视频子标题
func (obj *_MacVodMgr) WithVodSub(vodSub string) Option {
	return optionFunc(func(o *options) { o.query["vod_sub"] = vodSub })
}

// WithVodEn vod_en获取 视频英文或者拼音名称
func (obj *_MacVodMgr) WithVodEn(vodEn string) Option {
	return optionFunc(func(o *options) { o.query["vod_en"] = vodEn })
}

// WithVodStatus vod_status获取 视频状态
func (obj *_MacVodMgr) WithVodStatus(vodStatus uint8) Option {
	return optionFunc(func(o *options) { o.query["vod_status"] = vodStatus })
}

// WithVodLetter vod_letter获取 首字母索引排序
func (obj *_MacVodMgr) WithVodLetter(vodLetter string) Option {
	return optionFunc(func(o *options) { o.query["vod_letter"] = vodLetter })
}

// WithVodColor vod_color获取 颜色值
func (obj *_MacVodMgr) WithVodColor(vodColor string) Option {
	return optionFunc(func(o *options) { o.query["vod_color"] = vodColor })
}

// WithVodTag vod_tag获取 视频标签
func (obj *_MacVodMgr) WithVodTag(vodTag string) Option {
	return optionFunc(func(o *options) { o.query["vod_tag"] = vodTag })
}

// WithVodClass vod_class获取 视频分类
func (obj *_MacVodMgr) WithVodClass(vodClass string) Option {
	return optionFunc(func(o *options) { o.query["vod_class"] = vodClass })
}

// WithVodPic vod_pic获取 视频封面图
func (obj *_MacVodMgr) WithVodPic(vodPic string) Option {
	return optionFunc(func(o *options) { o.query["vod_pic"] = vodPic })
}

// WithVodPicThumb vod_pic_thumb获取 视频封面图的缩略图
func (obj *_MacVodMgr) WithVodPicThumb(vodPicThumb string) Option {
	return optionFunc(func(o *options) { o.query["vod_pic_thumb"] = vodPicThumb })
}

// WithVodPicSlide vod_pic_slide获取 视频封面图的幻灯片
func (obj *_MacVodMgr) WithVodPicSlide(vodPicSlide string) Option {
	return optionFunc(func(o *options) { o.query["vod_pic_slide"] = vodPicSlide })
}

// WithVodPicScreenshot vod_pic_screenshot获取 视频截图
func (obj *_MacVodMgr) WithVodPicScreenshot(vodPicScreenshot string) Option {
	return optionFunc(func(o *options) { o.query["vod_pic_screenshot"] = vodPicScreenshot })
}

// WithVodActor vod_actor获取 视频演员
func (obj *_MacVodMgr) WithVodActor(vodActor string) Option {
	return optionFunc(func(o *options) { o.query["vod_actor"] = vodActor })
}

// WithVodDirector vod_director获取 视频导演
func (obj *_MacVodMgr) WithVodDirector(vodDirector string) Option {
	return optionFunc(func(o *options) { o.query["vod_director"] = vodDirector })
}

// WithVodWriter vod_writer获取 视频编剧
func (obj *_MacVodMgr) WithVodWriter(vodWriter string) Option {
	return optionFunc(func(o *options) { o.query["vod_writer"] = vodWriter })
}

// WithVodBehind vod_behind获取 幕后制作人员
func (obj *_MacVodMgr) WithVodBehind(vodBehind string) Option {
	return optionFunc(func(o *options) { o.query["vod_behind"] = vodBehind })
}

// WithVodBlurb vod_blurb获取 视频简介
func (obj *_MacVodMgr) WithVodBlurb(vodBlurb string) Option {
	return optionFunc(func(o *options) { o.query["vod_blurb"] = vodBlurb })
}

// WithVodRemarks vod_remarks获取 备注
func (obj *_MacVodMgr) WithVodRemarks(vodRemarks string) Option {
	return optionFunc(func(o *options) { o.query["vod_remarks"] = vodRemarks })
}

// WithVodPubdate vod_pubdate获取 视频发布日期
func (obj *_MacVodMgr) WithVodPubdate(vodPubdate string) Option {
	return optionFunc(func(o *options) { o.query["vod_pubdate"] = vodPubdate })
}

// WithVodTotal vod_total获取 视频总数
func (obj *_MacVodMgr) WithVodTotal(vodTotal int16) Option {
	return optionFunc(func(o *options) { o.query["vod_total"] = vodTotal })
}

// WithVodSerial vod_serial获取 视频连载信息
func (obj *_MacVodMgr) WithVodSerial(vodSerial string) Option {
	return optionFunc(func(o *options) { o.query["vod_serial"] = vodSerial })
}

// WithVodTv vod_tv获取 视频播出的电视台
func (obj *_MacVodMgr) WithVodTv(vodTv string) Option {
	return optionFunc(func(o *options) { o.query["vod_tv"] = vodTv })
}

// WithVodWeekday vod_weekday获取 视频播出的星期几
func (obj *_MacVodMgr) WithVodWeekday(vodWeekday string) Option {
	return optionFunc(func(o *options) { o.query["vod_weekday"] = vodWeekday })
}

// WithVodArea vod_area获取 视频制作地区
func (obj *_MacVodMgr) WithVodArea(vodArea string) Option {
	return optionFunc(func(o *options) { o.query["vod_area"] = vodArea })
}

// WithVodLang vod_lang获取 视频语言
func (obj *_MacVodMgr) WithVodLang(vodLang string) Option {
	return optionFunc(func(o *options) { o.query["vod_lang"] = vodLang })
}

// WithVodYear vod_year获取 视频制作年份
func (obj *_MacVodMgr) WithVodYear(vodYear string) Option {
	return optionFunc(func(o *options) { o.query["vod_year"] = vodYear })
}

// WithVodVersion vod_version获取 视频版本
func (obj *_MacVodMgr) WithVodVersion(vodVersion string) Option {
	return optionFunc(func(o *options) { o.query["vod_version"] = vodVersion })
}

// WithVodState vod_state获取 视频状态
func (obj *_MacVodMgr) WithVodState(vodState string) Option {
	return optionFunc(func(o *options) { o.query["vod_state"] = vodState })
}

// WithVodAuthor vod_author获取 视频作者
func (obj *_MacVodMgr) WithVodAuthor(vodAuthor string) Option {
	return optionFunc(func(o *options) { o.query["vod_author"] = vodAuthor })
}

// WithVodJumpurl vod_jumpurl获取 跳转链接
func (obj *_MacVodMgr) WithVodJumpurl(vodJumpurl string) Option {
	return optionFunc(func(o *options) { o.query["vod_jumpurl"] = vodJumpurl })
}

// WithVodTpl vod_tpl获取 模板信息
func (obj *_MacVodMgr) WithVodTpl(vodTpl string) Option {
	return optionFunc(func(o *options) { o.query["vod_tpl"] = vodTpl })
}

// WithVodTplPlay vod_tpl_play获取 播放模板信息
func (obj *_MacVodMgr) WithVodTplPlay(vodTplPlay string) Option {
	return optionFunc(func(o *options) { o.query["vod_tpl_play"] = vodTplPlay })
}

// WithVodTplDown vod_tpl_down获取 下载模板信息
func (obj *_MacVodMgr) WithVodTplDown(vodTplDown string) Option {
	return optionFunc(func(o *options) { o.query["vod_tpl_down"] = vodTplDown })
}

// WithVodIsend vod_isend获取 视频是否完结
func (obj *_MacVodMgr) WithVodIsend(vodIsend uint8) Option {
	return optionFunc(func(o *options) { o.query["vod_isend"] = vodIsend })
}

// WithVodLock vod_lock获取 视频是否锁定
func (obj *_MacVodMgr) WithVodLock(vodLock uint8) Option {
	return optionFunc(func(o *options) { o.query["vod_lock"] = vodLock })
}

// WithVodLevel vod_level获取 视频级别
func (obj *_MacVodMgr) WithVodLevel(vodLevel uint8) Option {
	return optionFunc(func(o *options) { o.query["vod_level"] = vodLevel })
}

// WithVodCopyright vod_copyright获取 视频版权信息
func (obj *_MacVodMgr) WithVodCopyright(vodCopyright uint8) Option {
	return optionFunc(func(o *options) { o.query["vod_copyright"] = vodCopyright })
}

// WithVodPoints vod_points获取 视频积分
func (obj *_MacVodMgr) WithVodPoints(vodPoints uint16) Option {
	return optionFunc(func(o *options) { o.query["vod_points"] = vodPoints })
}

// WithVodPointsPlay vod_points_play获取 播放积分
func (obj *_MacVodMgr) WithVodPointsPlay(vodPointsPlay uint16) Option {
	return optionFunc(func(o *options) { o.query["vod_points_play"] = vodPointsPlay })
}

// WithVodPointsDown vod_points_down获取 下载积分
func (obj *_MacVodMgr) WithVodPointsDown(vodPointsDown uint16) Option {
	return optionFunc(func(o *options) { o.query["vod_points_down"] = vodPointsDown })
}

// WithVodHits vod_hits获取 视频点击数
func (obj *_MacVodMgr) WithVodHits(vodHits int16) Option {
	return optionFunc(func(o *options) { o.query["vod_hits"] = vodHits })
}

// WithVodHitsDay vod_hits_day获取 视频每日点击数
func (obj *_MacVodMgr) WithVodHitsDay(vodHitsDay int16) Option {
	return optionFunc(func(o *options) { o.query["vod_hits_day"] = vodHitsDay })
}

// WithVodHitsWeek vod_hits_week获取 视频每周点击数
func (obj *_MacVodMgr) WithVodHitsWeek(vodHitsWeek int16) Option {
	return optionFunc(func(o *options) { o.query["vod_hits_week"] = vodHitsWeek })
}

// WithVodHitsMonth vod_hits_month获取 视频每月点击数
func (obj *_MacVodMgr) WithVodHitsMonth(vodHitsMonth int16) Option {
	return optionFunc(func(o *options) { o.query["vod_hits_month"] = vodHitsMonth })
}

// WithVodDuration vod_duration获取 视频时长
func (obj *_MacVodMgr) WithVodDuration(vodDuration string) Option {
	return optionFunc(func(o *options) { o.query["vod_duration"] = vodDuration })
}

// WithVodUp vod_up获取 视频点赞数
func (obj *_MacVodMgr) WithVodUp(vodUp int16) Option {
	return optionFunc(func(o *options) { o.query["vod_up"] = vodUp })
}

// WithVodDown vod_down获取 视频踩数
func (obj *_MacVodMgr) WithVodDown(vodDown int16) Option {
	return optionFunc(func(o *options) { o.query["vod_down"] = vodDown })
}

// WithVodScore vod_score获取 视频评分
func (obj *_MacVodMgr) WithVodScore(vodScore float64) Option {
	return optionFunc(func(o *options) { o.query["vod_score"] = vodScore })
}

// WithVodScoreAll vod_score_all获取 视频总评分数
func (obj *_MacVodMgr) WithVodScoreAll(vodScoreAll int16) Option {
	return optionFunc(func(o *options) { o.query["vod_score_all"] = vodScoreAll })
}

// WithVodScoreNum vod_score_num获取 视频评分人数
func (obj *_MacVodMgr) WithVodScoreNum(vodScoreNum int16) Option {
	return optionFunc(func(o *options) { o.query["vod_score_num"] = vodScoreNum })
}

// WithVodTime vod_time获取 视频总时长
func (obj *_MacVodMgr) WithVodTime(vodTime int) Option {
	return optionFunc(func(o *options) { o.query["vod_time"] = vodTime })
}

// WithVodTimeAdd vod_time_add获取 视频添加时长
func (obj *_MacVodMgr) WithVodTimeAdd(vodTimeAdd int) Option {
	return optionFunc(func(o *options) { o.query["vod_time_add"] = vodTimeAdd })
}

// WithVodTimeHits vod_time_hits获取 视频点击时长
func (obj *_MacVodMgr) WithVodTimeHits(vodTimeHits int) Option {
	return optionFunc(func(o *options) { o.query["vod_time_hits"] = vodTimeHits })
}

// WithVodTimeMake vod_time_make获取 视频制作时长
func (obj *_MacVodMgr) WithVodTimeMake(vodTimeMake int) Option {
	return optionFunc(func(o *options) { o.query["vod_time_make"] = vodTimeMake })
}

// WithVodTrysee vod_trysee获取 试看时长
func (obj *_MacVodMgr) WithVodTrysee(vodTrysee uint16) Option {
	return optionFunc(func(o *options) { o.query["vod_trysee"] = vodTrysee })
}

// WithVodDoubanID vod_douban_id获取 豆瓣ID
func (obj *_MacVodMgr) WithVodDoubanID(vodDoubanID int) Option {
	return optionFunc(func(o *options) { o.query["vod_douban_id"] = vodDoubanID })
}

// WithVodDoubanScore vod_douban_score获取 豆瓣评分
func (obj *_MacVodMgr) WithVodDoubanScore(vodDoubanScore float64) Option {
	return optionFunc(func(o *options) { o.query["vod_douban_score"] = vodDoubanScore })
}

// WithVodReurl vod_reurl获取 重定向链接
func (obj *_MacVodMgr) WithVodReurl(vodReurl string) Option {
	return optionFunc(func(o *options) { o.query["vod_reurl"] = vodReurl })
}

// WithVodRelVod vod_rel_vod获取 相关视频链接
func (obj *_MacVodMgr) WithVodRelVod(vodRelVod string) Option {
	return optionFunc(func(o *options) { o.query["vod_rel_vod"] = vodRelVod })
}

// WithVodRelArt vod_rel_art获取 相关文章链接
func (obj *_MacVodMgr) WithVodRelArt(vodRelArt string) Option {
	return optionFunc(func(o *options) { o.query["vod_rel_art"] = vodRelArt })
}

// WithVodPwd vod_pwd获取 视频密码
func (obj *_MacVodMgr) WithVodPwd(vodPwd string) Option {
	return optionFunc(func(o *options) { o.query["vod_pwd"] = vodPwd })
}

// WithVodPwdURL vod_pwd_url获取 密码链接
func (obj *_MacVodMgr) WithVodPwdURL(vodPwdURL string) Option {
	return optionFunc(func(o *options) { o.query["vod_pwd_url"] = vodPwdURL })
}

// WithVodPwdPlay vod_pwd_play获取 播放密码
func (obj *_MacVodMgr) WithVodPwdPlay(vodPwdPlay string) Option {
	return optionFunc(func(o *options) { o.query["vod_pwd_play"] = vodPwdPlay })
}

// WithVodPwdPlayURL vod_pwd_play_url获取 播放密码链接
func (obj *_MacVodMgr) WithVodPwdPlayURL(vodPwdPlayURL string) Option {
	return optionFunc(func(o *options) { o.query["vod_pwd_play_url"] = vodPwdPlayURL })
}

// WithVodPwdDown vod_pwd_down获取 下载密码
func (obj *_MacVodMgr) WithVodPwdDown(vodPwdDown string) Option {
	return optionFunc(func(o *options) { o.query["vod_pwd_down"] = vodPwdDown })
}

// WithVodPwdDownURL vod_pwd_down_url获取 下载密码链接
func (obj *_MacVodMgr) WithVodPwdDownURL(vodPwdDownURL string) Option {
	return optionFunc(func(o *options) { o.query["vod_pwd_down_url"] = vodPwdDownURL })
}

// WithVodContent vod_content获取 视频内容
func (obj *_MacVodMgr) WithVodContent(vodContent string) Option {
	return optionFunc(func(o *options) { o.query["vod_content"] = vodContent })
}

// WithVodPlayFrom vod_play_from获取 播放来源
func (obj *_MacVodMgr) WithVodPlayFrom(vodPlayFrom string) Option {
	return optionFunc(func(o *options) { o.query["vod_play_from"] = vodPlayFrom })
}

// WithVodPlayServer vod_play_server获取 播放服务器
func (obj *_MacVodMgr) WithVodPlayServer(vodPlayServer string) Option {
	return optionFunc(func(o *options) { o.query["vod_play_server"] = vodPlayServer })
}

// WithVodPlayNote vod_play_note获取 播放注意事项
func (obj *_MacVodMgr) WithVodPlayNote(vodPlayNote string) Option {
	return optionFunc(func(o *options) { o.query["vod_play_note"] = vodPlayNote })
}

// WithVodPlayURL vod_play_url获取 播放链接
func (obj *_MacVodMgr) WithVodPlayURL(vodPlayURL string) Option {
	return optionFunc(func(o *options) { o.query["vod_play_url"] = vodPlayURL })
}

// WithVodDownFrom vod_down_from获取 下载来源
func (obj *_MacVodMgr) WithVodDownFrom(vodDownFrom string) Option {
	return optionFunc(func(o *options) { o.query["vod_down_from"] = vodDownFrom })
}

// WithVodDownServer vod_down_server获取 下载服务器
func (obj *_MacVodMgr) WithVodDownServer(vodDownServer string) Option {
	return optionFunc(func(o *options) { o.query["vod_down_server"] = vodDownServer })
}

// WithVodDownNote vod_down_note获取 下载注意事项
func (obj *_MacVodMgr) WithVodDownNote(vodDownNote string) Option {
	return optionFunc(func(o *options) { o.query["vod_down_note"] = vodDownNote })
}

// WithVodDownURL vod_down_url获取 下载链接
func (obj *_MacVodMgr) WithVodDownURL(vodDownURL string) Option {
	return optionFunc(func(o *options) { o.query["vod_down_url"] = vodDownURL })
}

// WithVodPlot vod_plot获取 剧情类型
func (obj *_MacVodMgr) WithVodPlot(vodPlot uint8) Option {
	return optionFunc(func(o *options) { o.query["vod_plot"] = vodPlot })
}

// WithVodPlotName vod_plot_name获取 剧情名称
func (obj *_MacVodMgr) WithVodPlotName(vodPlotName string) Option {
	return optionFunc(func(o *options) { o.query["vod_plot_name"] = vodPlotName })
}

// WithVodPlotDetail vod_plot_detail获取 剧情详情
func (obj *_MacVodMgr) WithVodPlotDetail(vodPlotDetail string) Option {
	return optionFunc(func(o *options) { o.query["vod_plot_detail"] = vodPlotDetail })
}

// GetByOption 功能选项模式获取
func (obj *_MacVodMgr) GetByOption(opts ...Option) (result MacVod, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MacVodMgr) GetByOptions(opts ...Option) (results []*MacVod, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MacVodMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MacVod, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where(options.query)
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

// GetFromVodID 通过vod_id获取内容 视频ID
func (obj *_MacVodMgr) GetFromVodID(vodID int) (result MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_id` = ?", vodID).First(&result).Error

	return
}

// GetBatchFromVodID 批量查找 视频ID
func (obj *_MacVodMgr) GetBatchFromVodID(vodIDs []int) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_id` IN (?)", vodIDs).Find(&results).Error

	return
}

// GetFromTypeID 通过type_id获取内容 栏目ID
func (obj *_MacVodMgr) GetFromTypeID(typeID int16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`type_id` = ?", typeID).Find(&results).Error

	return
}

// GetBatchFromTypeID 批量查找 栏目ID
func (obj *_MacVodMgr) GetBatchFromTypeID(typeIDs []int16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`type_id` IN (?)", typeIDs).Find(&results).Error

	return
}

// GetFromTypeID1 通过type_id_1获取内容 子栏目ID
func (obj *_MacVodMgr) GetFromTypeID1(typeID1 uint16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`type_id_1` = ?", typeID1).Find(&results).Error

	return
}

// GetBatchFromTypeID1 批量查找 子栏目ID
func (obj *_MacVodMgr) GetBatchFromTypeID1(typeID1s []uint16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`type_id_1` IN (?)", typeID1s).Find(&results).Error

	return
}

// GetFromGroupID 通过group_id获取内容 分组ID
func (obj *_MacVodMgr) GetFromGroupID(groupID uint16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`group_id` = ?", groupID).Find(&results).Error

	return
}

// GetBatchFromGroupID 批量查找 分组ID
func (obj *_MacVodMgr) GetBatchFromGroupID(groupIDs []uint16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`group_id` IN (?)", groupIDs).Find(&results).Error

	return
}

// GetFromVodName 通过vod_name获取内容 视频标题
func (obj *_MacVodMgr) GetFromVodName(vodName string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_name` = ?", vodName).Find(&results).Error

	return
}

// GetBatchFromVodName 批量查找 视频标题
func (obj *_MacVodMgr) GetBatchFromVodName(vodNames []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_name` IN (?)", vodNames).Find(&results).Error

	return
}

// GetFromVodSub 通过vod_sub获取内容 视频子标题
func (obj *_MacVodMgr) GetFromVodSub(vodSub string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_sub` = ?", vodSub).Find(&results).Error

	return
}

// GetBatchFromVodSub 批量查找 视频子标题
func (obj *_MacVodMgr) GetBatchFromVodSub(vodSubs []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_sub` IN (?)", vodSubs).Find(&results).Error

	return
}

// GetFromVodEn 通过vod_en获取内容 视频英文或者拼音名称
func (obj *_MacVodMgr) GetFromVodEn(vodEn string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_en` = ?", vodEn).Find(&results).Error

	return
}

// GetBatchFromVodEn 批量查找 视频英文或者拼音名称
func (obj *_MacVodMgr) GetBatchFromVodEn(vodEns []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_en` IN (?)", vodEns).Find(&results).Error

	return
}

// GetFromVodStatus 通过vod_status获取内容 视频状态
func (obj *_MacVodMgr) GetFromVodStatus(vodStatus uint8) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_status` = ?", vodStatus).Find(&results).Error

	return
}

// GetBatchFromVodStatus 批量查找 视频状态
func (obj *_MacVodMgr) GetBatchFromVodStatus(vodStatuss []uint8) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_status` IN (?)", vodStatuss).Find(&results).Error

	return
}

// GetFromVodLetter 通过vod_letter获取内容 首字母索引排序
func (obj *_MacVodMgr) GetFromVodLetter(vodLetter string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_letter` = ?", vodLetter).Find(&results).Error

	return
}

// GetBatchFromVodLetter 批量查找 首字母索引排序
func (obj *_MacVodMgr) GetBatchFromVodLetter(vodLetters []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_letter` IN (?)", vodLetters).Find(&results).Error

	return
}

// GetFromVodColor 通过vod_color获取内容 颜色值
func (obj *_MacVodMgr) GetFromVodColor(vodColor string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_color` = ?", vodColor).Find(&results).Error

	return
}

// GetBatchFromVodColor 批量查找 颜色值
func (obj *_MacVodMgr) GetBatchFromVodColor(vodColors []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_color` IN (?)", vodColors).Find(&results).Error

	return
}

// GetFromVodTag 通过vod_tag获取内容 视频标签
func (obj *_MacVodMgr) GetFromVodTag(vodTag string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_tag` = ?", vodTag).Find(&results).Error

	return
}

// GetBatchFromVodTag 批量查找 视频标签
func (obj *_MacVodMgr) GetBatchFromVodTag(vodTags []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_tag` IN (?)", vodTags).Find(&results).Error

	return
}

// GetFromVodClass 通过vod_class获取内容 视频分类
func (obj *_MacVodMgr) GetFromVodClass(vodClass string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_class` = ?", vodClass).Find(&results).Error

	return
}

// GetBatchFromVodClass 批量查找 视频分类
func (obj *_MacVodMgr) GetBatchFromVodClass(vodClasss []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_class` IN (?)", vodClasss).Find(&results).Error

	return
}

// GetFromVodPic 通过vod_pic获取内容 视频封面图
func (obj *_MacVodMgr) GetFromVodPic(vodPic string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_pic` = ?", vodPic).Find(&results).Error

	return
}

// GetBatchFromVodPic 批量查找 视频封面图
func (obj *_MacVodMgr) GetBatchFromVodPic(vodPics []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_pic` IN (?)", vodPics).Find(&results).Error

	return
}

// GetFromVodPicThumb 通过vod_pic_thumb获取内容 视频封面图的缩略图
func (obj *_MacVodMgr) GetFromVodPicThumb(vodPicThumb string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_pic_thumb` = ?", vodPicThumb).Find(&results).Error

	return
}

// GetBatchFromVodPicThumb 批量查找 视频封面图的缩略图
func (obj *_MacVodMgr) GetBatchFromVodPicThumb(vodPicThumbs []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_pic_thumb` IN (?)", vodPicThumbs).Find(&results).Error

	return
}

// GetFromVodPicSlide 通过vod_pic_slide获取内容 视频封面图的幻灯片
func (obj *_MacVodMgr) GetFromVodPicSlide(vodPicSlide string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_pic_slide` = ?", vodPicSlide).Find(&results).Error

	return
}

// GetBatchFromVodPicSlide 批量查找 视频封面图的幻灯片
func (obj *_MacVodMgr) GetBatchFromVodPicSlide(vodPicSlides []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_pic_slide` IN (?)", vodPicSlides).Find(&results).Error

	return
}

// GetFromVodPicScreenshot 通过vod_pic_screenshot获取内容 视频截图
func (obj *_MacVodMgr) GetFromVodPicScreenshot(vodPicScreenshot string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_pic_screenshot` = ?", vodPicScreenshot).Find(&results).Error

	return
}

// GetBatchFromVodPicScreenshot 批量查找 视频截图
func (obj *_MacVodMgr) GetBatchFromVodPicScreenshot(vodPicScreenshots []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_pic_screenshot` IN (?)", vodPicScreenshots).Find(&results).Error

	return
}

// GetFromVodActor 通过vod_actor获取内容 视频演员
func (obj *_MacVodMgr) GetFromVodActor(vodActor string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_actor` = ?", vodActor).Find(&results).Error

	return
}

// GetBatchFromVodActor 批量查找 视频演员
func (obj *_MacVodMgr) GetBatchFromVodActor(vodActors []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_actor` IN (?)", vodActors).Find(&results).Error

	return
}

// GetFromVodDirector 通过vod_director获取内容 视频导演
func (obj *_MacVodMgr) GetFromVodDirector(vodDirector string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_director` = ?", vodDirector).Find(&results).Error

	return
}

// GetBatchFromVodDirector 批量查找 视频导演
func (obj *_MacVodMgr) GetBatchFromVodDirector(vodDirectors []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_director` IN (?)", vodDirectors).Find(&results).Error

	return
}

// GetFromVodWriter 通过vod_writer获取内容 视频编剧
func (obj *_MacVodMgr) GetFromVodWriter(vodWriter string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_writer` = ?", vodWriter).Find(&results).Error

	return
}

// GetBatchFromVodWriter 批量查找 视频编剧
func (obj *_MacVodMgr) GetBatchFromVodWriter(vodWriters []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_writer` IN (?)", vodWriters).Find(&results).Error

	return
}

// GetFromVodBehind 通过vod_behind获取内容 幕后制作人员
func (obj *_MacVodMgr) GetFromVodBehind(vodBehind string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_behind` = ?", vodBehind).Find(&results).Error

	return
}

// GetBatchFromVodBehind 批量查找 幕后制作人员
func (obj *_MacVodMgr) GetBatchFromVodBehind(vodBehinds []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_behind` IN (?)", vodBehinds).Find(&results).Error

	return
}

// GetFromVodBlurb 通过vod_blurb获取内容 视频简介
func (obj *_MacVodMgr) GetFromVodBlurb(vodBlurb string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_blurb` = ?", vodBlurb).Find(&results).Error

	return
}

// GetBatchFromVodBlurb 批量查找 视频简介
func (obj *_MacVodMgr) GetBatchFromVodBlurb(vodBlurbs []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_blurb` IN (?)", vodBlurbs).Find(&results).Error

	return
}

// GetFromVodRemarks 通过vod_remarks获取内容 备注
func (obj *_MacVodMgr) GetFromVodRemarks(vodRemarks string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_remarks` = ?", vodRemarks).Find(&results).Error

	return
}

// GetBatchFromVodRemarks 批量查找 备注
func (obj *_MacVodMgr) GetBatchFromVodRemarks(vodRemarkss []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_remarks` IN (?)", vodRemarkss).Find(&results).Error

	return
}

// GetFromVodPubdate 通过vod_pubdate获取内容 视频发布日期
func (obj *_MacVodMgr) GetFromVodPubdate(vodPubdate string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_pubdate` = ?", vodPubdate).Find(&results).Error

	return
}

// GetBatchFromVodPubdate 批量查找 视频发布日期
func (obj *_MacVodMgr) GetBatchFromVodPubdate(vodPubdates []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_pubdate` IN (?)", vodPubdates).Find(&results).Error

	return
}

// GetFromVodTotal 通过vod_total获取内容 视频总数
func (obj *_MacVodMgr) GetFromVodTotal(vodTotal int16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_total` = ?", vodTotal).Find(&results).Error

	return
}

// GetBatchFromVodTotal 批量查找 视频总数
func (obj *_MacVodMgr) GetBatchFromVodTotal(vodTotals []int16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_total` IN (?)", vodTotals).Find(&results).Error

	return
}

// GetFromVodSerial 通过vod_serial获取内容 视频连载信息
func (obj *_MacVodMgr) GetFromVodSerial(vodSerial string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_serial` = ?", vodSerial).Find(&results).Error

	return
}

// GetBatchFromVodSerial 批量查找 视频连载信息
func (obj *_MacVodMgr) GetBatchFromVodSerial(vodSerials []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_serial` IN (?)", vodSerials).Find(&results).Error

	return
}

// GetFromVodTv 通过vod_tv获取内容 视频播出的电视台
func (obj *_MacVodMgr) GetFromVodTv(vodTv string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_tv` = ?", vodTv).Find(&results).Error

	return
}

// GetBatchFromVodTv 批量查找 视频播出的电视台
func (obj *_MacVodMgr) GetBatchFromVodTv(vodTvs []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_tv` IN (?)", vodTvs).Find(&results).Error

	return
}

// GetFromVodWeekday 通过vod_weekday获取内容 视频播出的星期几
func (obj *_MacVodMgr) GetFromVodWeekday(vodWeekday string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_weekday` = ?", vodWeekday).Find(&results).Error

	return
}

// GetBatchFromVodWeekday 批量查找 视频播出的星期几
func (obj *_MacVodMgr) GetBatchFromVodWeekday(vodWeekdays []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_weekday` IN (?)", vodWeekdays).Find(&results).Error

	return
}

// GetFromVodArea 通过vod_area获取内容 视频制作地区
func (obj *_MacVodMgr) GetFromVodArea(vodArea string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_area` = ?", vodArea).Find(&results).Error

	return
}

// GetBatchFromVodArea 批量查找 视频制作地区
func (obj *_MacVodMgr) GetBatchFromVodArea(vodAreas []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_area` IN (?)", vodAreas).Find(&results).Error

	return
}

// GetFromVodLang 通过vod_lang获取内容 视频语言
func (obj *_MacVodMgr) GetFromVodLang(vodLang string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_lang` = ?", vodLang).Find(&results).Error

	return
}

// GetBatchFromVodLang 批量查找 视频语言
func (obj *_MacVodMgr) GetBatchFromVodLang(vodLangs []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_lang` IN (?)", vodLangs).Find(&results).Error

	return
}

// GetFromVodYear 通过vod_year获取内容 视频制作年份
func (obj *_MacVodMgr) GetFromVodYear(vodYear string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_year` = ?", vodYear).Find(&results).Error

	return
}

// GetBatchFromVodYear 批量查找 视频制作年份
func (obj *_MacVodMgr) GetBatchFromVodYear(vodYears []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_year` IN (?)", vodYears).Find(&results).Error

	return
}

// GetFromVodVersion 通过vod_version获取内容 视频版本
func (obj *_MacVodMgr) GetFromVodVersion(vodVersion string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_version` = ?", vodVersion).Find(&results).Error

	return
}

// GetBatchFromVodVersion 批量查找 视频版本
func (obj *_MacVodMgr) GetBatchFromVodVersion(vodVersions []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_version` IN (?)", vodVersions).Find(&results).Error

	return
}

// GetFromVodState 通过vod_state获取内容 视频状态
func (obj *_MacVodMgr) GetFromVodState(vodState string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_state` = ?", vodState).Find(&results).Error

	return
}

// GetBatchFromVodState 批量查找 视频状态
func (obj *_MacVodMgr) GetBatchFromVodState(vodStates []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_state` IN (?)", vodStates).Find(&results).Error

	return
}

// GetFromVodAuthor 通过vod_author获取内容 视频作者
func (obj *_MacVodMgr) GetFromVodAuthor(vodAuthor string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_author` = ?", vodAuthor).Find(&results).Error

	return
}

// GetBatchFromVodAuthor 批量查找 视频作者
func (obj *_MacVodMgr) GetBatchFromVodAuthor(vodAuthors []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_author` IN (?)", vodAuthors).Find(&results).Error

	return
}

// GetFromVodJumpurl 通过vod_jumpurl获取内容 跳转链接
func (obj *_MacVodMgr) GetFromVodJumpurl(vodJumpurl string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_jumpurl` = ?", vodJumpurl).Find(&results).Error

	return
}

// GetBatchFromVodJumpurl 批量查找 跳转链接
func (obj *_MacVodMgr) GetBatchFromVodJumpurl(vodJumpurls []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_jumpurl` IN (?)", vodJumpurls).Find(&results).Error

	return
}

// GetFromVodTpl 通过vod_tpl获取内容 模板信息
func (obj *_MacVodMgr) GetFromVodTpl(vodTpl string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_tpl` = ?", vodTpl).Find(&results).Error

	return
}

// GetBatchFromVodTpl 批量查找 模板信息
func (obj *_MacVodMgr) GetBatchFromVodTpl(vodTpls []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_tpl` IN (?)", vodTpls).Find(&results).Error

	return
}

// GetFromVodTplPlay 通过vod_tpl_play获取内容 播放模板信息
func (obj *_MacVodMgr) GetFromVodTplPlay(vodTplPlay string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_tpl_play` = ?", vodTplPlay).Find(&results).Error

	return
}

// GetBatchFromVodTplPlay 批量查找 播放模板信息
func (obj *_MacVodMgr) GetBatchFromVodTplPlay(vodTplPlays []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_tpl_play` IN (?)", vodTplPlays).Find(&results).Error

	return
}

// GetFromVodTplDown 通过vod_tpl_down获取内容 下载模板信息
func (obj *_MacVodMgr) GetFromVodTplDown(vodTplDown string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_tpl_down` = ?", vodTplDown).Find(&results).Error

	return
}

// GetBatchFromVodTplDown 批量查找 下载模板信息
func (obj *_MacVodMgr) GetBatchFromVodTplDown(vodTplDowns []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_tpl_down` IN (?)", vodTplDowns).Find(&results).Error

	return
}

// GetFromVodIsend 通过vod_isend获取内容 视频是否完结
func (obj *_MacVodMgr) GetFromVodIsend(vodIsend uint8) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_isend` = ?", vodIsend).Find(&results).Error

	return
}

// GetBatchFromVodIsend 批量查找 视频是否完结
func (obj *_MacVodMgr) GetBatchFromVodIsend(vodIsends []uint8) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_isend` IN (?)", vodIsends).Find(&results).Error

	return
}

// GetFromVodLock 通过vod_lock获取内容 视频是否锁定
func (obj *_MacVodMgr) GetFromVodLock(vodLock uint8) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_lock` = ?", vodLock).Find(&results).Error

	return
}

// GetBatchFromVodLock 批量查找 视频是否锁定
func (obj *_MacVodMgr) GetBatchFromVodLock(vodLocks []uint8) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_lock` IN (?)", vodLocks).Find(&results).Error

	return
}

// GetFromVodLevel 通过vod_level获取内容 视频级别
func (obj *_MacVodMgr) GetFromVodLevel(vodLevel uint8) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_level` = ?", vodLevel).Find(&results).Error

	return
}

// GetBatchFromVodLevel 批量查找 视频级别
func (obj *_MacVodMgr) GetBatchFromVodLevel(vodLevels []uint8) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_level` IN (?)", vodLevels).Find(&results).Error

	return
}

// GetFromVodCopyright 通过vod_copyright获取内容 视频版权信息
func (obj *_MacVodMgr) GetFromVodCopyright(vodCopyright uint8) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_copyright` = ?", vodCopyright).Find(&results).Error

	return
}

// GetBatchFromVodCopyright 批量查找 视频版权信息
func (obj *_MacVodMgr) GetBatchFromVodCopyright(vodCopyrights []uint8) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_copyright` IN (?)", vodCopyrights).Find(&results).Error

	return
}

// GetFromVodPoints 通过vod_points获取内容 视频积分
func (obj *_MacVodMgr) GetFromVodPoints(vodPoints uint16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_points` = ?", vodPoints).Find(&results).Error

	return
}

// GetBatchFromVodPoints 批量查找 视频积分
func (obj *_MacVodMgr) GetBatchFromVodPoints(vodPointss []uint16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_points` IN (?)", vodPointss).Find(&results).Error

	return
}

// GetFromVodPointsPlay 通过vod_points_play获取内容 播放积分
func (obj *_MacVodMgr) GetFromVodPointsPlay(vodPointsPlay uint16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_points_play` = ?", vodPointsPlay).Find(&results).Error

	return
}

// GetBatchFromVodPointsPlay 批量查找 播放积分
func (obj *_MacVodMgr) GetBatchFromVodPointsPlay(vodPointsPlays []uint16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_points_play` IN (?)", vodPointsPlays).Find(&results).Error

	return
}

// GetFromVodPointsDown 通过vod_points_down获取内容 下载积分
func (obj *_MacVodMgr) GetFromVodPointsDown(vodPointsDown uint16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_points_down` = ?", vodPointsDown).Find(&results).Error

	return
}

// GetBatchFromVodPointsDown 批量查找 下载积分
func (obj *_MacVodMgr) GetBatchFromVodPointsDown(vodPointsDowns []uint16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_points_down` IN (?)", vodPointsDowns).Find(&results).Error

	return
}

// GetFromVodHits 通过vod_hits获取内容 视频点击数
func (obj *_MacVodMgr) GetFromVodHits(vodHits int16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_hits` = ?", vodHits).Find(&results).Error

	return
}

// GetBatchFromVodHits 批量查找 视频点击数
func (obj *_MacVodMgr) GetBatchFromVodHits(vodHitss []int16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_hits` IN (?)", vodHitss).Find(&results).Error

	return
}

// GetFromVodHitsDay 通过vod_hits_day获取内容 视频每日点击数
func (obj *_MacVodMgr) GetFromVodHitsDay(vodHitsDay int16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_hits_day` = ?", vodHitsDay).Find(&results).Error

	return
}

// GetBatchFromVodHitsDay 批量查找 视频每日点击数
func (obj *_MacVodMgr) GetBatchFromVodHitsDay(vodHitsDays []int16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_hits_day` IN (?)", vodHitsDays).Find(&results).Error

	return
}

// GetFromVodHitsWeek 通过vod_hits_week获取内容 视频每周点击数
func (obj *_MacVodMgr) GetFromVodHitsWeek(vodHitsWeek int16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_hits_week` = ?", vodHitsWeek).Find(&results).Error

	return
}

// GetBatchFromVodHitsWeek 批量查找 视频每周点击数
func (obj *_MacVodMgr) GetBatchFromVodHitsWeek(vodHitsWeeks []int16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_hits_week` IN (?)", vodHitsWeeks).Find(&results).Error

	return
}

// GetFromVodHitsMonth 通过vod_hits_month获取内容 视频每月点击数
func (obj *_MacVodMgr) GetFromVodHitsMonth(vodHitsMonth int16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_hits_month` = ?", vodHitsMonth).Find(&results).Error

	return
}

// GetBatchFromVodHitsMonth 批量查找 视频每月点击数
func (obj *_MacVodMgr) GetBatchFromVodHitsMonth(vodHitsMonths []int16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_hits_month` IN (?)", vodHitsMonths).Find(&results).Error

	return
}

// GetFromVodDuration 通过vod_duration获取内容 视频时长
func (obj *_MacVodMgr) GetFromVodDuration(vodDuration string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_duration` = ?", vodDuration).Find(&results).Error

	return
}

// GetBatchFromVodDuration 批量查找 视频时长
func (obj *_MacVodMgr) GetBatchFromVodDuration(vodDurations []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_duration` IN (?)", vodDurations).Find(&results).Error

	return
}

// GetFromVodUp 通过vod_up获取内容 视频点赞数
func (obj *_MacVodMgr) GetFromVodUp(vodUp int16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_up` = ?", vodUp).Find(&results).Error

	return
}

// GetBatchFromVodUp 批量查找 视频点赞数
func (obj *_MacVodMgr) GetBatchFromVodUp(vodUps []int16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_up` IN (?)", vodUps).Find(&results).Error

	return
}

// GetFromVodDown 通过vod_down获取内容 视频踩数
func (obj *_MacVodMgr) GetFromVodDown(vodDown int16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_down` = ?", vodDown).Find(&results).Error

	return
}

// GetBatchFromVodDown 批量查找 视频踩数
func (obj *_MacVodMgr) GetBatchFromVodDown(vodDowns []int16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_down` IN (?)", vodDowns).Find(&results).Error

	return
}

// GetFromVodScore 通过vod_score获取内容 视频评分
func (obj *_MacVodMgr) GetFromVodScore(vodScore float64) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_score` = ?", vodScore).Find(&results).Error

	return
}

// GetBatchFromVodScore 批量查找 视频评分
func (obj *_MacVodMgr) GetBatchFromVodScore(vodScores []float64) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_score` IN (?)", vodScores).Find(&results).Error

	return
}

// GetFromVodScoreAll 通过vod_score_all获取内容 视频总评分数
func (obj *_MacVodMgr) GetFromVodScoreAll(vodScoreAll int16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_score_all` = ?", vodScoreAll).Find(&results).Error

	return
}

// GetBatchFromVodScoreAll 批量查找 视频总评分数
func (obj *_MacVodMgr) GetBatchFromVodScoreAll(vodScoreAlls []int16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_score_all` IN (?)", vodScoreAlls).Find(&results).Error

	return
}

// GetFromVodScoreNum 通过vod_score_num获取内容 视频评分人数
func (obj *_MacVodMgr) GetFromVodScoreNum(vodScoreNum int16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_score_num` = ?", vodScoreNum).Find(&results).Error

	return
}

// GetBatchFromVodScoreNum 批量查找 视频评分人数
func (obj *_MacVodMgr) GetBatchFromVodScoreNum(vodScoreNums []int16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_score_num` IN (?)", vodScoreNums).Find(&results).Error

	return
}

// GetFromVodTime 通过vod_time获取内容 视频总时长
func (obj *_MacVodMgr) GetFromVodTime(vodTime int) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_time` = ?", vodTime).Find(&results).Error

	return
}

// GetBatchFromVodTime 批量查找 视频总时长
func (obj *_MacVodMgr) GetBatchFromVodTime(vodTimes []int) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_time` IN (?)", vodTimes).Find(&results).Error

	return
}

// GetFromVodTimeAdd 通过vod_time_add获取内容 视频添加时长
func (obj *_MacVodMgr) GetFromVodTimeAdd(vodTimeAdd int) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_time_add` = ?", vodTimeAdd).Find(&results).Error

	return
}

// GetBatchFromVodTimeAdd 批量查找 视频添加时长
func (obj *_MacVodMgr) GetBatchFromVodTimeAdd(vodTimeAdds []int) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_time_add` IN (?)", vodTimeAdds).Find(&results).Error

	return
}

// GetFromVodTimeHits 通过vod_time_hits获取内容 视频点击时长
func (obj *_MacVodMgr) GetFromVodTimeHits(vodTimeHits int) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_time_hits` = ?", vodTimeHits).Find(&results).Error

	return
}

// GetBatchFromVodTimeHits 批量查找 视频点击时长
func (obj *_MacVodMgr) GetBatchFromVodTimeHits(vodTimeHitss []int) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_time_hits` IN (?)", vodTimeHitss).Find(&results).Error

	return
}

// GetFromVodTimeMake 通过vod_time_make获取内容 视频制作时长
func (obj *_MacVodMgr) GetFromVodTimeMake(vodTimeMake int) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_time_make` = ?", vodTimeMake).Find(&results).Error

	return
}

// GetBatchFromVodTimeMake 批量查找 视频制作时长
func (obj *_MacVodMgr) GetBatchFromVodTimeMake(vodTimeMakes []int) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_time_make` IN (?)", vodTimeMakes).Find(&results).Error

	return
}

// GetFromVodTrysee 通过vod_trysee获取内容 试看时长
func (obj *_MacVodMgr) GetFromVodTrysee(vodTrysee uint16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_trysee` = ?", vodTrysee).Find(&results).Error

	return
}

// GetBatchFromVodTrysee 批量查找 试看时长
func (obj *_MacVodMgr) GetBatchFromVodTrysee(vodTrysees []uint16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_trysee` IN (?)", vodTrysees).Find(&results).Error

	return
}

// GetFromVodDoubanID 通过vod_douban_id获取内容 豆瓣ID
func (obj *_MacVodMgr) GetFromVodDoubanID(vodDoubanID int) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_douban_id` = ?", vodDoubanID).Find(&results).Error

	return
}

// GetBatchFromVodDoubanID 批量查找 豆瓣ID
func (obj *_MacVodMgr) GetBatchFromVodDoubanID(vodDoubanIDs []int) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_douban_id` IN (?)", vodDoubanIDs).Find(&results).Error

	return
}

// GetFromVodDoubanScore 通过vod_douban_score获取内容 豆瓣评分
func (obj *_MacVodMgr) GetFromVodDoubanScore(vodDoubanScore float64) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_douban_score` = ?", vodDoubanScore).Find(&results).Error

	return
}

// GetBatchFromVodDoubanScore 批量查找 豆瓣评分
func (obj *_MacVodMgr) GetBatchFromVodDoubanScore(vodDoubanScores []float64) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_douban_score` IN (?)", vodDoubanScores).Find(&results).Error

	return
}

// GetFromVodReurl 通过vod_reurl获取内容 重定向链接
func (obj *_MacVodMgr) GetFromVodReurl(vodReurl string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_reurl` = ?", vodReurl).Find(&results).Error

	return
}

// GetBatchFromVodReurl 批量查找 重定向链接
func (obj *_MacVodMgr) GetBatchFromVodReurl(vodReurls []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_reurl` IN (?)", vodReurls).Find(&results).Error

	return
}

// GetFromVodRelVod 通过vod_rel_vod获取内容 相关视频链接
func (obj *_MacVodMgr) GetFromVodRelVod(vodRelVod string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_rel_vod` = ?", vodRelVod).Find(&results).Error

	return
}

// GetBatchFromVodRelVod 批量查找 相关视频链接
func (obj *_MacVodMgr) GetBatchFromVodRelVod(vodRelVods []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_rel_vod` IN (?)", vodRelVods).Find(&results).Error

	return
}

// GetFromVodRelArt 通过vod_rel_art获取内容 相关文章链接
func (obj *_MacVodMgr) GetFromVodRelArt(vodRelArt string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_rel_art` = ?", vodRelArt).Find(&results).Error

	return
}

// GetBatchFromVodRelArt 批量查找 相关文章链接
func (obj *_MacVodMgr) GetBatchFromVodRelArt(vodRelArts []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_rel_art` IN (?)", vodRelArts).Find(&results).Error

	return
}

// GetFromVodPwd 通过vod_pwd获取内容 视频密码
func (obj *_MacVodMgr) GetFromVodPwd(vodPwd string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_pwd` = ?", vodPwd).Find(&results).Error

	return
}

// GetBatchFromVodPwd 批量查找 视频密码
func (obj *_MacVodMgr) GetBatchFromVodPwd(vodPwds []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_pwd` IN (?)", vodPwds).Find(&results).Error

	return
}

// GetFromVodPwdURL 通过vod_pwd_url获取内容 密码链接
func (obj *_MacVodMgr) GetFromVodPwdURL(vodPwdURL string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_pwd_url` = ?", vodPwdURL).Find(&results).Error

	return
}

// GetBatchFromVodPwdURL 批量查找 密码链接
func (obj *_MacVodMgr) GetBatchFromVodPwdURL(vodPwdURLs []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_pwd_url` IN (?)", vodPwdURLs).Find(&results).Error

	return
}

// GetFromVodPwdPlay 通过vod_pwd_play获取内容 播放密码
func (obj *_MacVodMgr) GetFromVodPwdPlay(vodPwdPlay string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_pwd_play` = ?", vodPwdPlay).Find(&results).Error

	return
}

// GetBatchFromVodPwdPlay 批量查找 播放密码
func (obj *_MacVodMgr) GetBatchFromVodPwdPlay(vodPwdPlays []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_pwd_play` IN (?)", vodPwdPlays).Find(&results).Error

	return
}

// GetFromVodPwdPlayURL 通过vod_pwd_play_url获取内容 播放密码链接
func (obj *_MacVodMgr) GetFromVodPwdPlayURL(vodPwdPlayURL string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_pwd_play_url` = ?", vodPwdPlayURL).Find(&results).Error

	return
}

// GetBatchFromVodPwdPlayURL 批量查找 播放密码链接
func (obj *_MacVodMgr) GetBatchFromVodPwdPlayURL(vodPwdPlayURLs []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_pwd_play_url` IN (?)", vodPwdPlayURLs).Find(&results).Error

	return
}

// GetFromVodPwdDown 通过vod_pwd_down获取内容 下载密码
func (obj *_MacVodMgr) GetFromVodPwdDown(vodPwdDown string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_pwd_down` = ?", vodPwdDown).Find(&results).Error

	return
}

// GetBatchFromVodPwdDown 批量查找 下载密码
func (obj *_MacVodMgr) GetBatchFromVodPwdDown(vodPwdDowns []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_pwd_down` IN (?)", vodPwdDowns).Find(&results).Error

	return
}

// GetFromVodPwdDownURL 通过vod_pwd_down_url获取内容 下载密码链接
func (obj *_MacVodMgr) GetFromVodPwdDownURL(vodPwdDownURL string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_pwd_down_url` = ?", vodPwdDownURL).Find(&results).Error

	return
}

// GetBatchFromVodPwdDownURL 批量查找 下载密码链接
func (obj *_MacVodMgr) GetBatchFromVodPwdDownURL(vodPwdDownURLs []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_pwd_down_url` IN (?)", vodPwdDownURLs).Find(&results).Error

	return
}

// GetFromVodContent 通过vod_content获取内容 视频内容
func (obj *_MacVodMgr) GetFromVodContent(vodContent string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_content` = ?", vodContent).Find(&results).Error

	return
}

// GetBatchFromVodContent 批量查找 视频内容
func (obj *_MacVodMgr) GetBatchFromVodContent(vodContents []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_content` IN (?)", vodContents).Find(&results).Error

	return
}

// GetFromVodPlayFrom 通过vod_play_from获取内容 播放来源
func (obj *_MacVodMgr) GetFromVodPlayFrom(vodPlayFrom string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_play_from` = ?", vodPlayFrom).Find(&results).Error

	return
}

// GetBatchFromVodPlayFrom 批量查找 播放来源
func (obj *_MacVodMgr) GetBatchFromVodPlayFrom(vodPlayFroms []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_play_from` IN (?)", vodPlayFroms).Find(&results).Error

	return
}

// GetFromVodPlayServer 通过vod_play_server获取内容 播放服务器
func (obj *_MacVodMgr) GetFromVodPlayServer(vodPlayServer string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_play_server` = ?", vodPlayServer).Find(&results).Error

	return
}

// GetBatchFromVodPlayServer 批量查找 播放服务器
func (obj *_MacVodMgr) GetBatchFromVodPlayServer(vodPlayServers []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_play_server` IN (?)", vodPlayServers).Find(&results).Error

	return
}

// GetFromVodPlayNote 通过vod_play_note获取内容 播放注意事项
func (obj *_MacVodMgr) GetFromVodPlayNote(vodPlayNote string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_play_note` = ?", vodPlayNote).Find(&results).Error

	return
}

// GetBatchFromVodPlayNote 批量查找 播放注意事项
func (obj *_MacVodMgr) GetBatchFromVodPlayNote(vodPlayNotes []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_play_note` IN (?)", vodPlayNotes).Find(&results).Error

	return
}

// GetFromVodPlayURL 通过vod_play_url获取内容 播放链接
func (obj *_MacVodMgr) GetFromVodPlayURL(vodPlayURL string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_play_url` = ?", vodPlayURL).Find(&results).Error

	return
}

// GetBatchFromVodPlayURL 批量查找 播放链接
func (obj *_MacVodMgr) GetBatchFromVodPlayURL(vodPlayURLs []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_play_url` IN (?)", vodPlayURLs).Find(&results).Error

	return
}

// GetFromVodDownFrom 通过vod_down_from获取内容 下载来源
func (obj *_MacVodMgr) GetFromVodDownFrom(vodDownFrom string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_down_from` = ?", vodDownFrom).Find(&results).Error

	return
}

// GetBatchFromVodDownFrom 批量查找 下载来源
func (obj *_MacVodMgr) GetBatchFromVodDownFrom(vodDownFroms []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_down_from` IN (?)", vodDownFroms).Find(&results).Error

	return
}

// GetFromVodDownServer 通过vod_down_server获取内容 下载服务器
func (obj *_MacVodMgr) GetFromVodDownServer(vodDownServer string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_down_server` = ?", vodDownServer).Find(&results).Error

	return
}

// GetBatchFromVodDownServer 批量查找 下载服务器
func (obj *_MacVodMgr) GetBatchFromVodDownServer(vodDownServers []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_down_server` IN (?)", vodDownServers).Find(&results).Error

	return
}

// GetFromVodDownNote 通过vod_down_note获取内容 下载注意事项
func (obj *_MacVodMgr) GetFromVodDownNote(vodDownNote string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_down_note` = ?", vodDownNote).Find(&results).Error

	return
}

// GetBatchFromVodDownNote 批量查找 下载注意事项
func (obj *_MacVodMgr) GetBatchFromVodDownNote(vodDownNotes []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_down_note` IN (?)", vodDownNotes).Find(&results).Error

	return
}

// GetFromVodDownURL 通过vod_down_url获取内容 下载链接
func (obj *_MacVodMgr) GetFromVodDownURL(vodDownURL string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_down_url` = ?", vodDownURL).Find(&results).Error

	return
}

// GetBatchFromVodDownURL 批量查找 下载链接
func (obj *_MacVodMgr) GetBatchFromVodDownURL(vodDownURLs []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_down_url` IN (?)", vodDownURLs).Find(&results).Error

	return
}

// GetFromVodPlot 通过vod_plot获取内容 剧情类型
func (obj *_MacVodMgr) GetFromVodPlot(vodPlot uint8) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_plot` = ?", vodPlot).Find(&results).Error

	return
}

// GetBatchFromVodPlot 批量查找 剧情类型
func (obj *_MacVodMgr) GetBatchFromVodPlot(vodPlots []uint8) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_plot` IN (?)", vodPlots).Find(&results).Error

	return
}

// GetFromVodPlotName 通过vod_plot_name获取内容 剧情名称
func (obj *_MacVodMgr) GetFromVodPlotName(vodPlotName string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_plot_name` = ?", vodPlotName).Find(&results).Error

	return
}

// GetBatchFromVodPlotName 批量查找 剧情名称
func (obj *_MacVodMgr) GetBatchFromVodPlotName(vodPlotNames []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_plot_name` IN (?)", vodPlotNames).Find(&results).Error

	return
}

// GetFromVodPlotDetail 通过vod_plot_detail获取内容 剧情详情
func (obj *_MacVodMgr) GetFromVodPlotDetail(vodPlotDetail string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_plot_detail` = ?", vodPlotDetail).Find(&results).Error

	return
}

// GetBatchFromVodPlotDetail 批量查找 剧情详情
func (obj *_MacVodMgr) GetBatchFromVodPlotDetail(vodPlotDetails []string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_plot_detail` IN (?)", vodPlotDetails).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MacVodMgr) FetchByPrimaryKey(vodID int) (result MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_id` = ?", vodID).First(&result).Error

	return
}

// FetchIndexByTypeID  获取多个内容
func (obj *_MacVodMgr) FetchIndexByTypeID(typeID int16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`type_id` = ?", typeID).Find(&results).Error

	return
}

// FetchIndexByTypeID1  获取多个内容
func (obj *_MacVodMgr) FetchIndexByTypeID1(typeID1 uint16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`type_id_1` = ?", typeID1).Find(&results).Error

	return
}

// FetchIndexByGroupID  获取多个内容
func (obj *_MacVodMgr) FetchIndexByGroupID(groupID uint16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`group_id` = ?", groupID).Find(&results).Error

	return
}

// FetchIndexByVodName  获取多个内容
func (obj *_MacVodMgr) FetchIndexByVodName(vodName string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_name` = ?", vodName).Find(&results).Error

	return
}

// FetchIndexByVodEn  获取多个内容
func (obj *_MacVodMgr) FetchIndexByVodEn(vodEn string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_en` = ?", vodEn).Find(&results).Error

	return
}

// FetchIndexByVodLetter  获取多个内容
func (obj *_MacVodMgr) FetchIndexByVodLetter(vodLetter string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_letter` = ?", vodLetter).Find(&results).Error

	return
}

// FetchIndexByVodTag  获取多个内容
func (obj *_MacVodMgr) FetchIndexByVodTag(vodTag string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_tag` = ?", vodTag).Find(&results).Error

	return
}

// FetchIndexByVodClass  获取多个内容
func (obj *_MacVodMgr) FetchIndexByVodClass(vodClass string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_class` = ?", vodClass).Find(&results).Error

	return
}

// FetchIndexByVodActor  获取多个内容
func (obj *_MacVodMgr) FetchIndexByVodActor(vodActor string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_actor` = ?", vodActor).Find(&results).Error

	return
}

// FetchIndexByVodDirector  获取多个内容
func (obj *_MacVodMgr) FetchIndexByVodDirector(vodDirector string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_director` = ?", vodDirector).Find(&results).Error

	return
}

// FetchIndexByVodTotal  获取多个内容
func (obj *_MacVodMgr) FetchIndexByVodTotal(vodTotal int16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_total` = ?", vodTotal).Find(&results).Error

	return
}

// FetchIndexByVodArea  获取多个内容
func (obj *_MacVodMgr) FetchIndexByVodArea(vodArea string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_area` = ?", vodArea).Find(&results).Error

	return
}

// FetchIndexByVodLang  获取多个内容
func (obj *_MacVodMgr) FetchIndexByVodLang(vodLang string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_lang` = ?", vodLang).Find(&results).Error

	return
}

// FetchIndexByVodYear  获取多个内容
func (obj *_MacVodMgr) FetchIndexByVodYear(vodYear string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_year` = ?", vodYear).Find(&results).Error

	return
}

// FetchIndexByVodVersion  获取多个内容
func (obj *_MacVodMgr) FetchIndexByVodVersion(vodVersion string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_version` = ?", vodVersion).Find(&results).Error

	return
}

// FetchIndexByVodState  获取多个内容
func (obj *_MacVodMgr) FetchIndexByVodState(vodState string) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_state` = ?", vodState).Find(&results).Error

	return
}

// FetchIndexByVodIsend  获取多个内容
func (obj *_MacVodMgr) FetchIndexByVodIsend(vodIsend uint8) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_isend` = ?", vodIsend).Find(&results).Error

	return
}

// FetchIndexByVodLock  获取多个内容
func (obj *_MacVodMgr) FetchIndexByVodLock(vodLock uint8) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_lock` = ?", vodLock).Find(&results).Error

	return
}

// FetchIndexByVodLevel  获取多个内容
func (obj *_MacVodMgr) FetchIndexByVodLevel(vodLevel uint8) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_level` = ?", vodLevel).Find(&results).Error

	return
}

// FetchIndexByVodPointsPlay  获取多个内容
func (obj *_MacVodMgr) FetchIndexByVodPointsPlay(vodPointsPlay uint16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_points_play` = ?", vodPointsPlay).Find(&results).Error

	return
}

// FetchIndexByVodPointsDown  获取多个内容
func (obj *_MacVodMgr) FetchIndexByVodPointsDown(vodPointsDown uint16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_points_down` = ?", vodPointsDown).Find(&results).Error

	return
}

// FetchIndexByVodHits  获取多个内容
func (obj *_MacVodMgr) FetchIndexByVodHits(vodHits int16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_hits` = ?", vodHits).Find(&results).Error

	return
}

// FetchIndexByVodHitsDay  获取多个内容
func (obj *_MacVodMgr) FetchIndexByVodHitsDay(vodHitsDay int16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_hits_day` = ?", vodHitsDay).Find(&results).Error

	return
}

// FetchIndexByVodHitsWeek  获取多个内容
func (obj *_MacVodMgr) FetchIndexByVodHitsWeek(vodHitsWeek int16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_hits_week` = ?", vodHitsWeek).Find(&results).Error

	return
}

// FetchIndexByVodHitsMonth  获取多个内容
func (obj *_MacVodMgr) FetchIndexByVodHitsMonth(vodHitsMonth int16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_hits_month` = ?", vodHitsMonth).Find(&results).Error

	return
}

// FetchIndexByVodUp  获取多个内容
func (obj *_MacVodMgr) FetchIndexByVodUp(vodUp int16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_up` = ?", vodUp).Find(&results).Error

	return
}

// FetchIndexByVodDown  获取多个内容
func (obj *_MacVodMgr) FetchIndexByVodDown(vodDown int16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_down` = ?", vodDown).Find(&results).Error

	return
}

// FetchIndexByVodScore  获取多个内容
func (obj *_MacVodMgr) FetchIndexByVodScore(vodScore float64) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_score` = ?", vodScore).Find(&results).Error

	return
}

// FetchIndexByVodScoreAll  获取多个内容
func (obj *_MacVodMgr) FetchIndexByVodScoreAll(vodScoreAll int16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_score_all` = ?", vodScoreAll).Find(&results).Error

	return
}

// FetchIndexByVodScoreNum  获取多个内容
func (obj *_MacVodMgr) FetchIndexByVodScoreNum(vodScoreNum int16) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_score_num` = ?", vodScoreNum).Find(&results).Error

	return
}

// FetchIndexByVodTime  获取多个内容
func (obj *_MacVodMgr) FetchIndexByVodTime(vodTime int) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_time` = ?", vodTime).Find(&results).Error

	return
}

// FetchIndexByVodTimeAdd  获取多个内容
func (obj *_MacVodMgr) FetchIndexByVodTimeAdd(vodTimeAdd int) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_time_add` = ?", vodTimeAdd).Find(&results).Error

	return
}

// FetchIndexByVodTimeMake  获取多个内容
func (obj *_MacVodMgr) FetchIndexByVodTimeMake(vodTimeMake int) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_time_make` = ?", vodTimeMake).Find(&results).Error

	return
}

// FetchIndexByVodPlot  获取多个内容
func (obj *_MacVodMgr) FetchIndexByVodPlot(vodPlot uint8) (results []*MacVod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacVod{}).Where("`vod_plot` = ?", vodPlot).Find(&results).Error

	return
}
