package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _MacTypeMgr struct {
	*_BaseMgr
}

// MacTypeMgr open func
func MacTypeMgr(db *gorm.DB) *_MacTypeMgr {
	if db == nil {
		panic(fmt.Errorf("MacTypeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MacTypeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("mac_type"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_MacTypeMgr) Debug() *_MacTypeMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MacTypeMgr) GetTableName() string {
	return "mac_type"
}

// Reset 重置gorm会话
func (obj *_MacTypeMgr) Reset() *_MacTypeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MacTypeMgr) Get() (result MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MacTypeMgr) Gets() (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MacTypeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MacType{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTypeID type_id获取 分类ID
func (obj *_MacTypeMgr) WithTypeID(typeID uint16) Option {
	return optionFunc(func(o *options) { o.query["type_id"] = typeID })
}

// WithTypeName type_name获取 分类名称
func (obj *_MacTypeMgr) WithTypeName(typeName string) Option {
	return optionFunc(func(o *options) { o.query["type_name"] = typeName })
}

// WithTypeEn type_en获取 分类英文名
func (obj *_MacTypeMgr) WithTypeEn(typeEn string) Option {
	return optionFunc(func(o *options) { o.query["type_en"] = typeEn })
}

// WithTypeSort type_sort获取 分类排序
func (obj *_MacTypeMgr) WithTypeSort(typeSort uint16) Option {
	return optionFunc(func(o *options) { o.query["type_sort"] = typeSort })
}

// WithTypeMid type_mid获取 分类所属模型
func (obj *_MacTypeMgr) WithTypeMid(typeMid uint16) Option {
	return optionFunc(func(o *options) { o.query["type_mid"] = typeMid })
}

// WithTypePid type_pid获取 分类父ID
func (obj *_MacTypeMgr) WithTypePid(typePid uint16) Option {
	return optionFunc(func(o *options) { o.query["type_pid"] = typePid })
}

// WithTypeStatus type_status获取 分类状态
func (obj *_MacTypeMgr) WithTypeStatus(typeStatus uint8) Option {
	return optionFunc(func(o *options) { o.query["type_status"] = typeStatus })
}

// WithTypeTpl type_tpl获取 分类模板
func (obj *_MacTypeMgr) WithTypeTpl(typeTpl string) Option {
	return optionFunc(func(o *options) { o.query["type_tpl"] = typeTpl })
}

// WithTypeTplList type_tpl_list获取 分类列表模板
func (obj *_MacTypeMgr) WithTypeTplList(typeTplList string) Option {
	return optionFunc(func(o *options) { o.query["type_tpl_list"] = typeTplList })
}

// WithTypeTplDetail type_tpl_detail获取 分类详情模板
func (obj *_MacTypeMgr) WithTypeTplDetail(typeTplDetail string) Option {
	return optionFunc(func(o *options) { o.query["type_tpl_detail"] = typeTplDetail })
}

// WithTypeTplPlay type_tpl_play获取 分类播放页模板
func (obj *_MacTypeMgr) WithTypeTplPlay(typeTplPlay string) Option {
	return optionFunc(func(o *options) { o.query["type_tpl_play"] = typeTplPlay })
}

// WithTypeTplDown type_tpl_down获取 分类下载页模板
func (obj *_MacTypeMgr) WithTypeTplDown(typeTplDown string) Option {
	return optionFunc(func(o *options) { o.query["type_tpl_down"] = typeTplDown })
}

// WithTypeKey type_key获取 分类关键字
func (obj *_MacTypeMgr) WithTypeKey(typeKey string) Option {
	return optionFunc(func(o *options) { o.query["type_key"] = typeKey })
}

// WithTypeDes type_des获取 分类描述
func (obj *_MacTypeMgr) WithTypeDes(typeDes string) Option {
	return optionFunc(func(o *options) { o.query["type_des"] = typeDes })
}

// WithTypeTitle type_title获取 分类标题
func (obj *_MacTypeMgr) WithTypeTitle(typeTitle string) Option {
	return optionFunc(func(o *options) { o.query["type_title"] = typeTitle })
}

// WithTypeUnion type_union获取 分类关联信息
func (obj *_MacTypeMgr) WithTypeUnion(typeUnion string) Option {
	return optionFunc(func(o *options) { o.query["type_union"] = typeUnion })
}

// WithTypeExtend type_extend获取 分类扩展信息
func (obj *_MacTypeMgr) WithTypeExtend(typeExtend string) Option {
	return optionFunc(func(o *options) { o.query["type_extend"] = typeExtend })
}

// WithTypeLogo type_logo获取 分类Logo
func (obj *_MacTypeMgr) WithTypeLogo(typeLogo string) Option {
	return optionFunc(func(o *options) { o.query["type_logo"] = typeLogo })
}

// WithTypePic type_pic获取 分类图片
func (obj *_MacTypeMgr) WithTypePic(typePic string) Option {
	return optionFunc(func(o *options) { o.query["type_pic"] = typePic })
}

// WithTypeJumpurl type_jumpurl获取 分类跳转URL
func (obj *_MacTypeMgr) WithTypeJumpurl(typeJumpurl string) Option {
	return optionFunc(func(o *options) { o.query["type_jumpurl"] = typeJumpurl })
}

// GetByOption 功能选项模式获取
func (obj *_MacTypeMgr) GetByOption(opts ...Option) (result MacType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MacTypeMgr) GetByOptions(opts ...Option) (results []*MacType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MacTypeMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MacType, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MacType{}).Where(options.query)
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

// GetFromTypeID 通过type_id获取内容 分类ID
func (obj *_MacTypeMgr) GetFromTypeID(typeID uint16) (result MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_id` = ?", typeID).First(&result).Error

	return
}

// GetBatchFromTypeID 批量查找 分类ID
func (obj *_MacTypeMgr) GetBatchFromTypeID(typeIDs []uint16) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_id` IN (?)", typeIDs).Find(&results).Error

	return
}

// GetFromTypeName 通过type_name获取内容 分类名称
func (obj *_MacTypeMgr) GetFromTypeName(typeName string) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_name` = ?", typeName).Find(&results).Error

	return
}

// GetBatchFromTypeName 批量查找 分类名称
func (obj *_MacTypeMgr) GetBatchFromTypeName(typeNames []string) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_name` IN (?)", typeNames).Find(&results).Error

	return
}

// GetFromTypeEn 通过type_en获取内容 分类英文名
func (obj *_MacTypeMgr) GetFromTypeEn(typeEn string) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_en` = ?", typeEn).Find(&results).Error

	return
}

// GetBatchFromTypeEn 批量查找 分类英文名
func (obj *_MacTypeMgr) GetBatchFromTypeEn(typeEns []string) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_en` IN (?)", typeEns).Find(&results).Error

	return
}

// GetFromTypeSort 通过type_sort获取内容 分类排序
func (obj *_MacTypeMgr) GetFromTypeSort(typeSort uint16) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_sort` = ?", typeSort).Find(&results).Error

	return
}

// GetBatchFromTypeSort 批量查找 分类排序
func (obj *_MacTypeMgr) GetBatchFromTypeSort(typeSorts []uint16) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_sort` IN (?)", typeSorts).Find(&results).Error

	return
}

// GetFromTypeMid 通过type_mid获取内容 分类所属模型
func (obj *_MacTypeMgr) GetFromTypeMid(typeMid uint16) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_mid` = ?", typeMid).Find(&results).Error

	return
}

// GetBatchFromTypeMid 批量查找 分类所属模型
func (obj *_MacTypeMgr) GetBatchFromTypeMid(typeMids []uint16) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_mid` IN (?)", typeMids).Find(&results).Error

	return
}

// GetFromTypePid 通过type_pid获取内容 分类父ID
func (obj *_MacTypeMgr) GetFromTypePid(typePid uint16) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_pid` = ?", typePid).Find(&results).Error

	return
}

// GetBatchFromTypePid 批量查找 分类父ID
func (obj *_MacTypeMgr) GetBatchFromTypePid(typePids []uint16) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_pid` IN (?)", typePids).Find(&results).Error

	return
}

// GetFromTypeStatus 通过type_status获取内容 分类状态
func (obj *_MacTypeMgr) GetFromTypeStatus(typeStatus uint8) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_status` = ?", typeStatus).Find(&results).Error

	return
}

// GetBatchFromTypeStatus 批量查找 分类状态
func (obj *_MacTypeMgr) GetBatchFromTypeStatus(typeStatuss []uint8) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_status` IN (?)", typeStatuss).Find(&results).Error

	return
}

// GetFromTypeTpl 通过type_tpl获取内容 分类模板
func (obj *_MacTypeMgr) GetFromTypeTpl(typeTpl string) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_tpl` = ?", typeTpl).Find(&results).Error

	return
}

// GetBatchFromTypeTpl 批量查找 分类模板
func (obj *_MacTypeMgr) GetBatchFromTypeTpl(typeTpls []string) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_tpl` IN (?)", typeTpls).Find(&results).Error

	return
}

// GetFromTypeTplList 通过type_tpl_list获取内容 分类列表模板
func (obj *_MacTypeMgr) GetFromTypeTplList(typeTplList string) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_tpl_list` = ?", typeTplList).Find(&results).Error

	return
}

// GetBatchFromTypeTplList 批量查找 分类列表模板
func (obj *_MacTypeMgr) GetBatchFromTypeTplList(typeTplLists []string) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_tpl_list` IN (?)", typeTplLists).Find(&results).Error

	return
}

// GetFromTypeTplDetail 通过type_tpl_detail获取内容 分类详情模板
func (obj *_MacTypeMgr) GetFromTypeTplDetail(typeTplDetail string) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_tpl_detail` = ?", typeTplDetail).Find(&results).Error

	return
}

// GetBatchFromTypeTplDetail 批量查找 分类详情模板
func (obj *_MacTypeMgr) GetBatchFromTypeTplDetail(typeTplDetails []string) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_tpl_detail` IN (?)", typeTplDetails).Find(&results).Error

	return
}

// GetFromTypeTplPlay 通过type_tpl_play获取内容 分类播放页模板
func (obj *_MacTypeMgr) GetFromTypeTplPlay(typeTplPlay string) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_tpl_play` = ?", typeTplPlay).Find(&results).Error

	return
}

// GetBatchFromTypeTplPlay 批量查找 分类播放页模板
func (obj *_MacTypeMgr) GetBatchFromTypeTplPlay(typeTplPlays []string) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_tpl_play` IN (?)", typeTplPlays).Find(&results).Error

	return
}

// GetFromTypeTplDown 通过type_tpl_down获取内容 分类下载页模板
func (obj *_MacTypeMgr) GetFromTypeTplDown(typeTplDown string) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_tpl_down` = ?", typeTplDown).Find(&results).Error

	return
}

// GetBatchFromTypeTplDown 批量查找 分类下载页模板
func (obj *_MacTypeMgr) GetBatchFromTypeTplDown(typeTplDowns []string) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_tpl_down` IN (?)", typeTplDowns).Find(&results).Error

	return
}

// GetFromTypeKey 通过type_key获取内容 分类关键字
func (obj *_MacTypeMgr) GetFromTypeKey(typeKey string) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_key` = ?", typeKey).Find(&results).Error

	return
}

// GetBatchFromTypeKey 批量查找 分类关键字
func (obj *_MacTypeMgr) GetBatchFromTypeKey(typeKeys []string) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_key` IN (?)", typeKeys).Find(&results).Error

	return
}

// GetFromTypeDes 通过type_des获取内容 分类描述
func (obj *_MacTypeMgr) GetFromTypeDes(typeDes string) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_des` = ?", typeDes).Find(&results).Error

	return
}

// GetBatchFromTypeDes 批量查找 分类描述
func (obj *_MacTypeMgr) GetBatchFromTypeDes(typeDess []string) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_des` IN (?)", typeDess).Find(&results).Error

	return
}

// GetFromTypeTitle 通过type_title获取内容 分类标题
func (obj *_MacTypeMgr) GetFromTypeTitle(typeTitle string) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_title` = ?", typeTitle).Find(&results).Error

	return
}

// GetBatchFromTypeTitle 批量查找 分类标题
func (obj *_MacTypeMgr) GetBatchFromTypeTitle(typeTitles []string) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_title` IN (?)", typeTitles).Find(&results).Error

	return
}

// GetFromTypeUnion 通过type_union获取内容 分类关联信息
func (obj *_MacTypeMgr) GetFromTypeUnion(typeUnion string) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_union` = ?", typeUnion).Find(&results).Error

	return
}

// GetBatchFromTypeUnion 批量查找 分类关联信息
func (obj *_MacTypeMgr) GetBatchFromTypeUnion(typeUnions []string) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_union` IN (?)", typeUnions).Find(&results).Error

	return
}

// GetFromTypeExtend 通过type_extend获取内容 分类扩展信息
func (obj *_MacTypeMgr) GetFromTypeExtend(typeExtend string) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_extend` = ?", typeExtend).Find(&results).Error

	return
}

// GetBatchFromTypeExtend 批量查找 分类扩展信息
func (obj *_MacTypeMgr) GetBatchFromTypeExtend(typeExtends []string) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_extend` IN (?)", typeExtends).Find(&results).Error

	return
}

// GetFromTypeLogo 通过type_logo获取内容 分类Logo
func (obj *_MacTypeMgr) GetFromTypeLogo(typeLogo string) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_logo` = ?", typeLogo).Find(&results).Error

	return
}

// GetBatchFromTypeLogo 批量查找 分类Logo
func (obj *_MacTypeMgr) GetBatchFromTypeLogo(typeLogos []string) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_logo` IN (?)", typeLogos).Find(&results).Error

	return
}

// GetFromTypePic 通过type_pic获取内容 分类图片
func (obj *_MacTypeMgr) GetFromTypePic(typePic string) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_pic` = ?", typePic).Find(&results).Error

	return
}

// GetBatchFromTypePic 批量查找 分类图片
func (obj *_MacTypeMgr) GetBatchFromTypePic(typePics []string) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_pic` IN (?)", typePics).Find(&results).Error

	return
}

// GetFromTypeJumpurl 通过type_jumpurl获取内容 分类跳转URL
func (obj *_MacTypeMgr) GetFromTypeJumpurl(typeJumpurl string) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_jumpurl` = ?", typeJumpurl).Find(&results).Error

	return
}

// GetBatchFromTypeJumpurl 批量查找 分类跳转URL
func (obj *_MacTypeMgr) GetBatchFromTypeJumpurl(typeJumpurls []string) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_jumpurl` IN (?)", typeJumpurls).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MacTypeMgr) FetchByPrimaryKey(typeID uint16) (result MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_id` = ?", typeID).First(&result).Error

	return
}

// FetchIndexByTypeName  获取多个内容
func (obj *_MacTypeMgr) FetchIndexByTypeName(typeName string) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_name` = ?", typeName).Find(&results).Error

	return
}

// FetchIndexByTypeEn  获取多个内容
func (obj *_MacTypeMgr) FetchIndexByTypeEn(typeEn string) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_en` = ?", typeEn).Find(&results).Error

	return
}

// FetchIndexByTypeSort  获取多个内容
func (obj *_MacTypeMgr) FetchIndexByTypeSort(typeSort uint16) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_sort` = ?", typeSort).Find(&results).Error

	return
}

// FetchIndexByTypeMid  获取多个内容
func (obj *_MacTypeMgr) FetchIndexByTypeMid(typeMid uint16) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_mid` = ?", typeMid).Find(&results).Error

	return
}

// FetchIndexByTypePid  获取多个内容
func (obj *_MacTypeMgr) FetchIndexByTypePid(typePid uint16) (results []*MacType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacType{}).Where("`type_pid` = ?", typePid).Find(&results).Error

	return
}
