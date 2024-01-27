package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _MacCjNodeMgr struct {
	*_BaseMgr
}

// MacCjNodeMgr open func
func MacCjNodeMgr(db *gorm.DB) *_MacCjNodeMgr {
	if db == nil {
		panic(fmt.Errorf("MacCjNodeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MacCjNodeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("mac_cj_node"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_MacCjNodeMgr) Debug() *_MacCjNodeMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MacCjNodeMgr) GetTableName() string {
	return "mac_cj_node"
}

// Reset 重置gorm会话
func (obj *_MacCjNodeMgr) Reset() *_MacCjNodeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MacCjNodeMgr) Get() (result MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MacCjNodeMgr) Gets() (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MacCjNodeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithNodeid nodeid获取 节点ID
func (obj *_MacCjNodeMgr) WithNodeid(nodeid uint16) Option {
	return optionFunc(func(o *options) { o.query["nodeid"] = nodeid })
}

// WithName name获取 节点名称
func (obj *_MacCjNodeMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithLastdate lastdate获取 最后更新日期
func (obj *_MacCjNodeMgr) WithLastdate(lastdate int) Option {
	return optionFunc(func(o *options) { o.query["lastdate"] = lastdate })
}

// WithSourcecharset sourcecharset获取 源字符集
func (obj *_MacCjNodeMgr) WithSourcecharset(sourcecharset string) Option {
	return optionFunc(func(o *options) { o.query["sourcecharset"] = sourcecharset })
}

// WithSourcetype sourcetype获取 源类型
func (obj *_MacCjNodeMgr) WithSourcetype(sourcetype uint8) Option {
	return optionFunc(func(o *options) { o.query["sourcetype"] = sourcetype })
}

// WithURLpage urlpage获取 URL页面
func (obj *_MacCjNodeMgr) WithURLpage(urlpage string) Option {
	return optionFunc(func(o *options) { o.query["urlpage"] = urlpage })
}

// WithPagesizeStart pagesize_start获取 页面大小起始值
func (obj *_MacCjNodeMgr) WithPagesizeStart(pagesizeStart uint8) Option {
	return optionFunc(func(o *options) { o.query["pagesize_start"] = pagesizeStart })
}

// WithPagesizeEnd pagesize_end获取 页面大小结束值
func (obj *_MacCjNodeMgr) WithPagesizeEnd(pagesizeEnd int16) Option {
	return optionFunc(func(o *options) { o.query["pagesize_end"] = pagesizeEnd })
}

// WithPageBase page_base获取 页面基础URL
func (obj *_MacCjNodeMgr) WithPageBase(pageBase string) Option {
	return optionFunc(func(o *options) { o.query["page_base"] = pageBase })
}

// WithParNum par_num获取 参数数量
func (obj *_MacCjNodeMgr) WithParNum(parNum uint8) Option {
	return optionFunc(func(o *options) { o.query["par_num"] = parNum })
}

// WithURLContain url_contain获取 URL包含规则
func (obj *_MacCjNodeMgr) WithURLContain(urlContain string) Option {
	return optionFunc(func(o *options) { o.query["url_contain"] = urlContain })
}

// WithURLExcept url_except获取 URL排除规则
func (obj *_MacCjNodeMgr) WithURLExcept(urlExcept string) Option {
	return optionFunc(func(o *options) { o.query["url_except"] = urlExcept })
}

// WithURLStart url_start获取 URL起始规则
func (obj *_MacCjNodeMgr) WithURLStart(urlStart string) Option {
	return optionFunc(func(o *options) { o.query["url_start"] = urlStart })
}

// WithURLEnd url_end获取 URL结束规则
func (obj *_MacCjNodeMgr) WithURLEnd(urlEnd string) Option {
	return optionFunc(func(o *options) { o.query["url_end"] = urlEnd })
}

// WithTitleRule title_rule获取 标题规则
func (obj *_MacCjNodeMgr) WithTitleRule(titleRule string) Option {
	return optionFunc(func(o *options) { o.query["title_rule"] = titleRule })
}

// WithTitleHTMLRule title_html_rule获取 标题HTML规则
func (obj *_MacCjNodeMgr) WithTitleHTMLRule(titleHTMLRule string) Option {
	return optionFunc(func(o *options) { o.query["title_html_rule"] = titleHTMLRule })
}

// WithTypeRule type_rule获取 类型规则
func (obj *_MacCjNodeMgr) WithTypeRule(typeRule string) Option {
	return optionFunc(func(o *options) { o.query["type_rule"] = typeRule })
}

// WithTypeHTMLRule type_html_rule获取 类型HTML规则
func (obj *_MacCjNodeMgr) WithTypeHTMLRule(typeHTMLRule string) Option {
	return optionFunc(func(o *options) { o.query["type_html_rule"] = typeHTMLRule })
}

// WithContentRule content_rule获取 内容规则
func (obj *_MacCjNodeMgr) WithContentRule(contentRule string) Option {
	return optionFunc(func(o *options) { o.query["content_rule"] = contentRule })
}

// WithContentHTMLRule content_html_rule获取 内容HTML规则
func (obj *_MacCjNodeMgr) WithContentHTMLRule(contentHTMLRule string) Option {
	return optionFunc(func(o *options) { o.query["content_html_rule"] = contentHTMLRule })
}

// WithContentPageStart content_page_start获取 内容页起始规则
func (obj *_MacCjNodeMgr) WithContentPageStart(contentPageStart string) Option {
	return optionFunc(func(o *options) { o.query["content_page_start"] = contentPageStart })
}

// WithContentPageEnd content_page_end获取 内容页结束规则
func (obj *_MacCjNodeMgr) WithContentPageEnd(contentPageEnd string) Option {
	return optionFunc(func(o *options) { o.query["content_page_end"] = contentPageEnd })
}

// WithContentPageRule content_page_rule获取 内容页规则
func (obj *_MacCjNodeMgr) WithContentPageRule(contentPageRule uint8) Option {
	return optionFunc(func(o *options) { o.query["content_page_rule"] = contentPageRule })
}

// WithContentPage content_page获取 内容页
func (obj *_MacCjNodeMgr) WithContentPage(contentPage uint8) Option {
	return optionFunc(func(o *options) { o.query["content_page"] = contentPage })
}

// WithContentNextpage content_nextpage获取 下一页规则
func (obj *_MacCjNodeMgr) WithContentNextpage(contentNextpage string) Option {
	return optionFunc(func(o *options) { o.query["content_nextpage"] = contentNextpage })
}

// WithDownAttachment down_attachment获取 下载附件
func (obj *_MacCjNodeMgr) WithDownAttachment(downAttachment uint8) Option {
	return optionFunc(func(o *options) { o.query["down_attachment"] = downAttachment })
}

// WithWatermark watermark获取 水印
func (obj *_MacCjNodeMgr) WithWatermark(watermark uint8) Option {
	return optionFunc(func(o *options) { o.query["watermark"] = watermark })
}

// WithCollOrder coll_order获取 采集顺序
func (obj *_MacCjNodeMgr) WithCollOrder(collOrder uint8) Option {
	return optionFunc(func(o *options) { o.query["coll_order"] = collOrder })
}

// WithCustomizeConfig customize_config获取 自定义配置
func (obj *_MacCjNodeMgr) WithCustomizeConfig(customizeConfig string) Option {
	return optionFunc(func(o *options) { o.query["customize_config"] = customizeConfig })
}

// WithProgramConfig program_config获取 程序配置
func (obj *_MacCjNodeMgr) WithProgramConfig(programConfig string) Option {
	return optionFunc(func(o *options) { o.query["program_config"] = programConfig })
}

// WithMid mid获取 模型ID
func (obj *_MacCjNodeMgr) WithMid(mid uint8) Option {
	return optionFunc(func(o *options) { o.query["mid"] = mid })
}

// GetByOption 功能选项模式获取
func (obj *_MacCjNodeMgr) GetByOption(opts ...Option) (result MacCjNode, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MacCjNodeMgr) GetByOptions(opts ...Option) (results []*MacCjNode, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MacCjNodeMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MacCjNode, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where(options.query)
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

// GetFromNodeid 通过nodeid获取内容 节点ID
func (obj *_MacCjNodeMgr) GetFromNodeid(nodeid uint16) (result MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`nodeid` = ?", nodeid).First(&result).Error

	return
}

// GetBatchFromNodeid 批量查找 节点ID
func (obj *_MacCjNodeMgr) GetBatchFromNodeid(nodeids []uint16) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`nodeid` IN (?)", nodeids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 节点名称
func (obj *_MacCjNodeMgr) GetFromName(name string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 节点名称
func (obj *_MacCjNodeMgr) GetBatchFromName(names []string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromLastdate 通过lastdate获取内容 最后更新日期
func (obj *_MacCjNodeMgr) GetFromLastdate(lastdate int) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`lastdate` = ?", lastdate).Find(&results).Error

	return
}

// GetBatchFromLastdate 批量查找 最后更新日期
func (obj *_MacCjNodeMgr) GetBatchFromLastdate(lastdates []int) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`lastdate` IN (?)", lastdates).Find(&results).Error

	return
}

// GetFromSourcecharset 通过sourcecharset获取内容 源字符集
func (obj *_MacCjNodeMgr) GetFromSourcecharset(sourcecharset string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`sourcecharset` = ?", sourcecharset).Find(&results).Error

	return
}

// GetBatchFromSourcecharset 批量查找 源字符集
func (obj *_MacCjNodeMgr) GetBatchFromSourcecharset(sourcecharsets []string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`sourcecharset` IN (?)", sourcecharsets).Find(&results).Error

	return
}

// GetFromSourcetype 通过sourcetype获取内容 源类型
func (obj *_MacCjNodeMgr) GetFromSourcetype(sourcetype uint8) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`sourcetype` = ?", sourcetype).Find(&results).Error

	return
}

// GetBatchFromSourcetype 批量查找 源类型
func (obj *_MacCjNodeMgr) GetBatchFromSourcetype(sourcetypes []uint8) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`sourcetype` IN (?)", sourcetypes).Find(&results).Error

	return
}

// GetFromURLpage 通过urlpage获取内容 URL页面
func (obj *_MacCjNodeMgr) GetFromURLpage(urlpage string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`urlpage` = ?", urlpage).Find(&results).Error

	return
}

// GetBatchFromURLpage 批量查找 URL页面
func (obj *_MacCjNodeMgr) GetBatchFromURLpage(urlpages []string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`urlpage` IN (?)", urlpages).Find(&results).Error

	return
}

// GetFromPagesizeStart 通过pagesize_start获取内容 页面大小起始值
func (obj *_MacCjNodeMgr) GetFromPagesizeStart(pagesizeStart uint8) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`pagesize_start` = ?", pagesizeStart).Find(&results).Error

	return
}

// GetBatchFromPagesizeStart 批量查找 页面大小起始值
func (obj *_MacCjNodeMgr) GetBatchFromPagesizeStart(pagesizeStarts []uint8) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`pagesize_start` IN (?)", pagesizeStarts).Find(&results).Error

	return
}

// GetFromPagesizeEnd 通过pagesize_end获取内容 页面大小结束值
func (obj *_MacCjNodeMgr) GetFromPagesizeEnd(pagesizeEnd int16) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`pagesize_end` = ?", pagesizeEnd).Find(&results).Error

	return
}

// GetBatchFromPagesizeEnd 批量查找 页面大小结束值
func (obj *_MacCjNodeMgr) GetBatchFromPagesizeEnd(pagesizeEnds []int16) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`pagesize_end` IN (?)", pagesizeEnds).Find(&results).Error

	return
}

// GetFromPageBase 通过page_base获取内容 页面基础URL
func (obj *_MacCjNodeMgr) GetFromPageBase(pageBase string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`page_base` = ?", pageBase).Find(&results).Error

	return
}

// GetBatchFromPageBase 批量查找 页面基础URL
func (obj *_MacCjNodeMgr) GetBatchFromPageBase(pageBases []string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`page_base` IN (?)", pageBases).Find(&results).Error

	return
}

// GetFromParNum 通过par_num获取内容 参数数量
func (obj *_MacCjNodeMgr) GetFromParNum(parNum uint8) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`par_num` = ?", parNum).Find(&results).Error

	return
}

// GetBatchFromParNum 批量查找 参数数量
func (obj *_MacCjNodeMgr) GetBatchFromParNum(parNums []uint8) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`par_num` IN (?)", parNums).Find(&results).Error

	return
}

// GetFromURLContain 通过url_contain获取内容 URL包含规则
func (obj *_MacCjNodeMgr) GetFromURLContain(urlContain string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`url_contain` = ?", urlContain).Find(&results).Error

	return
}

// GetBatchFromURLContain 批量查找 URL包含规则
func (obj *_MacCjNodeMgr) GetBatchFromURLContain(urlContains []string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`url_contain` IN (?)", urlContains).Find(&results).Error

	return
}

// GetFromURLExcept 通过url_except获取内容 URL排除规则
func (obj *_MacCjNodeMgr) GetFromURLExcept(urlExcept string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`url_except` = ?", urlExcept).Find(&results).Error

	return
}

// GetBatchFromURLExcept 批量查找 URL排除规则
func (obj *_MacCjNodeMgr) GetBatchFromURLExcept(urlExcepts []string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`url_except` IN (?)", urlExcepts).Find(&results).Error

	return
}

// GetFromURLStart 通过url_start获取内容 URL起始规则
func (obj *_MacCjNodeMgr) GetFromURLStart(urlStart string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`url_start` = ?", urlStart).Find(&results).Error

	return
}

// GetBatchFromURLStart 批量查找 URL起始规则
func (obj *_MacCjNodeMgr) GetBatchFromURLStart(urlStarts []string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`url_start` IN (?)", urlStarts).Find(&results).Error

	return
}

// GetFromURLEnd 通过url_end获取内容 URL结束规则
func (obj *_MacCjNodeMgr) GetFromURLEnd(urlEnd string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`url_end` = ?", urlEnd).Find(&results).Error

	return
}

// GetBatchFromURLEnd 批量查找 URL结束规则
func (obj *_MacCjNodeMgr) GetBatchFromURLEnd(urlEnds []string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`url_end` IN (?)", urlEnds).Find(&results).Error

	return
}

// GetFromTitleRule 通过title_rule获取内容 标题规则
func (obj *_MacCjNodeMgr) GetFromTitleRule(titleRule string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`title_rule` = ?", titleRule).Find(&results).Error

	return
}

// GetBatchFromTitleRule 批量查找 标题规则
func (obj *_MacCjNodeMgr) GetBatchFromTitleRule(titleRules []string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`title_rule` IN (?)", titleRules).Find(&results).Error

	return
}

// GetFromTitleHTMLRule 通过title_html_rule获取内容 标题HTML规则
func (obj *_MacCjNodeMgr) GetFromTitleHTMLRule(titleHTMLRule string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`title_html_rule` = ?", titleHTMLRule).Find(&results).Error

	return
}

// GetBatchFromTitleHTMLRule 批量查找 标题HTML规则
func (obj *_MacCjNodeMgr) GetBatchFromTitleHTMLRule(titleHTMLRules []string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`title_html_rule` IN (?)", titleHTMLRules).Find(&results).Error

	return
}

// GetFromTypeRule 通过type_rule获取内容 类型规则
func (obj *_MacCjNodeMgr) GetFromTypeRule(typeRule string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`type_rule` = ?", typeRule).Find(&results).Error

	return
}

// GetBatchFromTypeRule 批量查找 类型规则
func (obj *_MacCjNodeMgr) GetBatchFromTypeRule(typeRules []string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`type_rule` IN (?)", typeRules).Find(&results).Error

	return
}

// GetFromTypeHTMLRule 通过type_html_rule获取内容 类型HTML规则
func (obj *_MacCjNodeMgr) GetFromTypeHTMLRule(typeHTMLRule string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`type_html_rule` = ?", typeHTMLRule).Find(&results).Error

	return
}

// GetBatchFromTypeHTMLRule 批量查找 类型HTML规则
func (obj *_MacCjNodeMgr) GetBatchFromTypeHTMLRule(typeHTMLRules []string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`type_html_rule` IN (?)", typeHTMLRules).Find(&results).Error

	return
}

// GetFromContentRule 通过content_rule获取内容 内容规则
func (obj *_MacCjNodeMgr) GetFromContentRule(contentRule string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`content_rule` = ?", contentRule).Find(&results).Error

	return
}

// GetBatchFromContentRule 批量查找 内容规则
func (obj *_MacCjNodeMgr) GetBatchFromContentRule(contentRules []string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`content_rule` IN (?)", contentRules).Find(&results).Error

	return
}

// GetFromContentHTMLRule 通过content_html_rule获取内容 内容HTML规则
func (obj *_MacCjNodeMgr) GetFromContentHTMLRule(contentHTMLRule string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`content_html_rule` = ?", contentHTMLRule).Find(&results).Error

	return
}

// GetBatchFromContentHTMLRule 批量查找 内容HTML规则
func (obj *_MacCjNodeMgr) GetBatchFromContentHTMLRule(contentHTMLRules []string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`content_html_rule` IN (?)", contentHTMLRules).Find(&results).Error

	return
}

// GetFromContentPageStart 通过content_page_start获取内容 内容页起始规则
func (obj *_MacCjNodeMgr) GetFromContentPageStart(contentPageStart string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`content_page_start` = ?", contentPageStart).Find(&results).Error

	return
}

// GetBatchFromContentPageStart 批量查找 内容页起始规则
func (obj *_MacCjNodeMgr) GetBatchFromContentPageStart(contentPageStarts []string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`content_page_start` IN (?)", contentPageStarts).Find(&results).Error

	return
}

// GetFromContentPageEnd 通过content_page_end获取内容 内容页结束规则
func (obj *_MacCjNodeMgr) GetFromContentPageEnd(contentPageEnd string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`content_page_end` = ?", contentPageEnd).Find(&results).Error

	return
}

// GetBatchFromContentPageEnd 批量查找 内容页结束规则
func (obj *_MacCjNodeMgr) GetBatchFromContentPageEnd(contentPageEnds []string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`content_page_end` IN (?)", contentPageEnds).Find(&results).Error

	return
}

// GetFromContentPageRule 通过content_page_rule获取内容 内容页规则
func (obj *_MacCjNodeMgr) GetFromContentPageRule(contentPageRule uint8) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`content_page_rule` = ?", contentPageRule).Find(&results).Error

	return
}

// GetBatchFromContentPageRule 批量查找 内容页规则
func (obj *_MacCjNodeMgr) GetBatchFromContentPageRule(contentPageRules []uint8) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`content_page_rule` IN (?)", contentPageRules).Find(&results).Error

	return
}

// GetFromContentPage 通过content_page获取内容 内容页
func (obj *_MacCjNodeMgr) GetFromContentPage(contentPage uint8) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`content_page` = ?", contentPage).Find(&results).Error

	return
}

// GetBatchFromContentPage 批量查找 内容页
func (obj *_MacCjNodeMgr) GetBatchFromContentPage(contentPages []uint8) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`content_page` IN (?)", contentPages).Find(&results).Error

	return
}

// GetFromContentNextpage 通过content_nextpage获取内容 下一页规则
func (obj *_MacCjNodeMgr) GetFromContentNextpage(contentNextpage string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`content_nextpage` = ?", contentNextpage).Find(&results).Error

	return
}

// GetBatchFromContentNextpage 批量查找 下一页规则
func (obj *_MacCjNodeMgr) GetBatchFromContentNextpage(contentNextpages []string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`content_nextpage` IN (?)", contentNextpages).Find(&results).Error

	return
}

// GetFromDownAttachment 通过down_attachment获取内容 下载附件
func (obj *_MacCjNodeMgr) GetFromDownAttachment(downAttachment uint8) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`down_attachment` = ?", downAttachment).Find(&results).Error

	return
}

// GetBatchFromDownAttachment 批量查找 下载附件
func (obj *_MacCjNodeMgr) GetBatchFromDownAttachment(downAttachments []uint8) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`down_attachment` IN (?)", downAttachments).Find(&results).Error

	return
}

// GetFromWatermark 通过watermark获取内容 水印
func (obj *_MacCjNodeMgr) GetFromWatermark(watermark uint8) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`watermark` = ?", watermark).Find(&results).Error

	return
}

// GetBatchFromWatermark 批量查找 水印
func (obj *_MacCjNodeMgr) GetBatchFromWatermark(watermarks []uint8) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`watermark` IN (?)", watermarks).Find(&results).Error

	return
}

// GetFromCollOrder 通过coll_order获取内容 采集顺序
func (obj *_MacCjNodeMgr) GetFromCollOrder(collOrder uint8) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`coll_order` = ?", collOrder).Find(&results).Error

	return
}

// GetBatchFromCollOrder 批量查找 采集顺序
func (obj *_MacCjNodeMgr) GetBatchFromCollOrder(collOrders []uint8) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`coll_order` IN (?)", collOrders).Find(&results).Error

	return
}

// GetFromCustomizeConfig 通过customize_config获取内容 自定义配置
func (obj *_MacCjNodeMgr) GetFromCustomizeConfig(customizeConfig string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`customize_config` = ?", customizeConfig).Find(&results).Error

	return
}

// GetBatchFromCustomizeConfig 批量查找 自定义配置
func (obj *_MacCjNodeMgr) GetBatchFromCustomizeConfig(customizeConfigs []string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`customize_config` IN (?)", customizeConfigs).Find(&results).Error

	return
}

// GetFromProgramConfig 通过program_config获取内容 程序配置
func (obj *_MacCjNodeMgr) GetFromProgramConfig(programConfig string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`program_config` = ?", programConfig).Find(&results).Error

	return
}

// GetBatchFromProgramConfig 批量查找 程序配置
func (obj *_MacCjNodeMgr) GetBatchFromProgramConfig(programConfigs []string) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`program_config` IN (?)", programConfigs).Find(&results).Error

	return
}

// GetFromMid 通过mid获取内容 模型ID
func (obj *_MacCjNodeMgr) GetFromMid(mid uint8) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`mid` = ?", mid).Find(&results).Error

	return
}

// GetBatchFromMid 批量查找 模型ID
func (obj *_MacCjNodeMgr) GetBatchFromMid(mids []uint8) (results []*MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`mid` IN (?)", mids).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MacCjNodeMgr) FetchByPrimaryKey(nodeid uint16) (result MacCjNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacCjNode{}).Where("`nodeid` = ?", nodeid).First(&result).Error

	return
}
