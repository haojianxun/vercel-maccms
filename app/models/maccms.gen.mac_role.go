package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _MacRoleMgr struct {
	*_BaseMgr
}

// MacRoleMgr open func
func MacRoleMgr(db *gorm.DB) *_MacRoleMgr {
	if db == nil {
		panic(fmt.Errorf("MacRoleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MacRoleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("mac_role"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_MacRoleMgr) Debug() *_MacRoleMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MacRoleMgr) GetTableName() string {
	return "mac_role"
}

// Reset 重置gorm会话
func (obj *_MacRoleMgr) Reset() *_MacRoleMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MacRoleMgr) Get() (result MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MacRoleMgr) Gets() (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MacRoleMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MacRole{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithRoleID role_id获取 角色ID
func (obj *_MacRoleMgr) WithRoleID(roleID int) Option {
	return optionFunc(func(o *options) { o.query["role_id"] = roleID })
}

// WithRoleRid role_rid获取 关联角色ID
func (obj *_MacRoleMgr) WithRoleRid(roleRid int) Option {
	return optionFunc(func(o *options) { o.query["role_rid"] = roleRid })
}

// WithRoleName role_name获取 角色名称
func (obj *_MacRoleMgr) WithRoleName(roleName string) Option {
	return optionFunc(func(o *options) { o.query["role_name"] = roleName })
}

// WithRoleEn role_en获取 角色英文名
func (obj *_MacRoleMgr) WithRoleEn(roleEn string) Option {
	return optionFunc(func(o *options) { o.query["role_en"] = roleEn })
}

// WithRoleStatus role_status获取 角色状态
func (obj *_MacRoleMgr) WithRoleStatus(roleStatus uint8) Option {
	return optionFunc(func(o *options) { o.query["role_status"] = roleStatus })
}

// WithRoleLock role_lock获取 角色锁定
func (obj *_MacRoleMgr) WithRoleLock(roleLock uint8) Option {
	return optionFunc(func(o *options) { o.query["role_lock"] = roleLock })
}

// WithRoleLetter role_letter获取 角色首字母
func (obj *_MacRoleMgr) WithRoleLetter(roleLetter string) Option {
	return optionFunc(func(o *options) { o.query["role_letter"] = roleLetter })
}

// WithRoleColor role_color获取 角色颜色
func (obj *_MacRoleMgr) WithRoleColor(roleColor string) Option {
	return optionFunc(func(o *options) { o.query["role_color"] = roleColor })
}

// WithRoleActor role_actor获取 关联演员
func (obj *_MacRoleMgr) WithRoleActor(roleActor string) Option {
	return optionFunc(func(o *options) { o.query["role_actor"] = roleActor })
}

// WithRoleRemarks role_remarks获取 角色备注
func (obj *_MacRoleMgr) WithRoleRemarks(roleRemarks string) Option {
	return optionFunc(func(o *options) { o.query["role_remarks"] = roleRemarks })
}

// WithRolePic role_pic获取 角色图片
func (obj *_MacRoleMgr) WithRolePic(rolePic string) Option {
	return optionFunc(func(o *options) { o.query["role_pic"] = rolePic })
}

// WithRoleSort role_sort获取 角色排序
func (obj *_MacRoleMgr) WithRoleSort(roleSort uint16) Option {
	return optionFunc(func(o *options) { o.query["role_sort"] = roleSort })
}

// WithRoleLevel role_level获取 角色等级
func (obj *_MacRoleMgr) WithRoleLevel(roleLevel uint8) Option {
	return optionFunc(func(o *options) { o.query["role_level"] = roleLevel })
}

// WithRoleTime role_time获取 角色时间
func (obj *_MacRoleMgr) WithRoleTime(roleTime int) Option {
	return optionFunc(func(o *options) { o.query["role_time"] = roleTime })
}

// WithRoleTimeAdd role_time_add获取 角色添加时间
func (obj *_MacRoleMgr) WithRoleTimeAdd(roleTimeAdd int) Option {
	return optionFunc(func(o *options) { o.query["role_time_add"] = roleTimeAdd })
}

// WithRoleTimeHits role_time_hits获取 角色点击时间
func (obj *_MacRoleMgr) WithRoleTimeHits(roleTimeHits int) Option {
	return optionFunc(func(o *options) { o.query["role_time_hits"] = roleTimeHits })
}

// WithRoleTimeMake role_time_make获取 角色制作时间
func (obj *_MacRoleMgr) WithRoleTimeMake(roleTimeMake int) Option {
	return optionFunc(func(o *options) { o.query["role_time_make"] = roleTimeMake })
}

// WithRoleHits role_hits获取 角色点击量
func (obj *_MacRoleMgr) WithRoleHits(roleHits int16) Option {
	return optionFunc(func(o *options) { o.query["role_hits"] = roleHits })
}

// WithRoleHitsDay role_hits_day获取 角色日点击量
func (obj *_MacRoleMgr) WithRoleHitsDay(roleHitsDay int16) Option {
	return optionFunc(func(o *options) { o.query["role_hits_day"] = roleHitsDay })
}

// WithRoleHitsWeek role_hits_week获取 角色周点击量
func (obj *_MacRoleMgr) WithRoleHitsWeek(roleHitsWeek int16) Option {
	return optionFunc(func(o *options) { o.query["role_hits_week"] = roleHitsWeek })
}

// WithRoleHitsMonth role_hits_month获取 角色月点击量
func (obj *_MacRoleMgr) WithRoleHitsMonth(roleHitsMonth int16) Option {
	return optionFunc(func(o *options) { o.query["role_hits_month"] = roleHitsMonth })
}

// WithRoleScore role_score获取 角色评分
func (obj *_MacRoleMgr) WithRoleScore(roleScore float64) Option {
	return optionFunc(func(o *options) { o.query["role_score"] = roleScore })
}

// WithRoleScoreAll role_score_all获取 角色总评分
func (obj *_MacRoleMgr) WithRoleScoreAll(roleScoreAll int16) Option {
	return optionFunc(func(o *options) { o.query["role_score_all"] = roleScoreAll })
}

// WithRoleScoreNum role_score_num获取 角色评分人数
func (obj *_MacRoleMgr) WithRoleScoreNum(roleScoreNum int16) Option {
	return optionFunc(func(o *options) { o.query["role_score_num"] = roleScoreNum })
}

// WithRoleUp role_up获取 角色点赞数
func (obj *_MacRoleMgr) WithRoleUp(roleUp int16) Option {
	return optionFunc(func(o *options) { o.query["role_up"] = roleUp })
}

// WithRoleDown role_down获取 角色踩数
func (obj *_MacRoleMgr) WithRoleDown(roleDown int16) Option {
	return optionFunc(func(o *options) { o.query["role_down"] = roleDown })
}

// WithRoleTpl role_tpl获取 角色模板
func (obj *_MacRoleMgr) WithRoleTpl(roleTpl string) Option {
	return optionFunc(func(o *options) { o.query["role_tpl"] = roleTpl })
}

// WithRoleJumpurl role_jumpurl获取 角色跳转URL
func (obj *_MacRoleMgr) WithRoleJumpurl(roleJumpurl string) Option {
	return optionFunc(func(o *options) { o.query["role_jumpurl"] = roleJumpurl })
}

// WithRoleContent role_content获取 角色内容
func (obj *_MacRoleMgr) WithRoleContent(roleContent string) Option {
	return optionFunc(func(o *options) { o.query["role_content"] = roleContent })
}

// GetByOption 功能选项模式获取
func (obj *_MacRoleMgr) GetByOption(opts ...Option) (result MacRole, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MacRoleMgr) GetByOptions(opts ...Option) (results []*MacRole, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MacRoleMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MacRole, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where(options.query)
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

// GetFromRoleID 通过role_id获取内容 角色ID
func (obj *_MacRoleMgr) GetFromRoleID(roleID int) (result MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_id` = ?", roleID).First(&result).Error

	return
}

// GetBatchFromRoleID 批量查找 角色ID
func (obj *_MacRoleMgr) GetBatchFromRoleID(roleIDs []int) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_id` IN (?)", roleIDs).Find(&results).Error

	return
}

// GetFromRoleRid 通过role_rid获取内容 关联角色ID
func (obj *_MacRoleMgr) GetFromRoleRid(roleRid int) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_rid` = ?", roleRid).Find(&results).Error

	return
}

// GetBatchFromRoleRid 批量查找 关联角色ID
func (obj *_MacRoleMgr) GetBatchFromRoleRid(roleRids []int) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_rid` IN (?)", roleRids).Find(&results).Error

	return
}

// GetFromRoleName 通过role_name获取内容 角色名称
func (obj *_MacRoleMgr) GetFromRoleName(roleName string) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_name` = ?", roleName).Find(&results).Error

	return
}

// GetBatchFromRoleName 批量查找 角色名称
func (obj *_MacRoleMgr) GetBatchFromRoleName(roleNames []string) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_name` IN (?)", roleNames).Find(&results).Error

	return
}

// GetFromRoleEn 通过role_en获取内容 角色英文名
func (obj *_MacRoleMgr) GetFromRoleEn(roleEn string) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_en` = ?", roleEn).Find(&results).Error

	return
}

// GetBatchFromRoleEn 批量查找 角色英文名
func (obj *_MacRoleMgr) GetBatchFromRoleEn(roleEns []string) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_en` IN (?)", roleEns).Find(&results).Error

	return
}

// GetFromRoleStatus 通过role_status获取内容 角色状态
func (obj *_MacRoleMgr) GetFromRoleStatus(roleStatus uint8) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_status` = ?", roleStatus).Find(&results).Error

	return
}

// GetBatchFromRoleStatus 批量查找 角色状态
func (obj *_MacRoleMgr) GetBatchFromRoleStatus(roleStatuss []uint8) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_status` IN (?)", roleStatuss).Find(&results).Error

	return
}

// GetFromRoleLock 通过role_lock获取内容 角色锁定
func (obj *_MacRoleMgr) GetFromRoleLock(roleLock uint8) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_lock` = ?", roleLock).Find(&results).Error

	return
}

// GetBatchFromRoleLock 批量查找 角色锁定
func (obj *_MacRoleMgr) GetBatchFromRoleLock(roleLocks []uint8) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_lock` IN (?)", roleLocks).Find(&results).Error

	return
}

// GetFromRoleLetter 通过role_letter获取内容 角色首字母
func (obj *_MacRoleMgr) GetFromRoleLetter(roleLetter string) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_letter` = ?", roleLetter).Find(&results).Error

	return
}

// GetBatchFromRoleLetter 批量查找 角色首字母
func (obj *_MacRoleMgr) GetBatchFromRoleLetter(roleLetters []string) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_letter` IN (?)", roleLetters).Find(&results).Error

	return
}

// GetFromRoleColor 通过role_color获取内容 角色颜色
func (obj *_MacRoleMgr) GetFromRoleColor(roleColor string) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_color` = ?", roleColor).Find(&results).Error

	return
}

// GetBatchFromRoleColor 批量查找 角色颜色
func (obj *_MacRoleMgr) GetBatchFromRoleColor(roleColors []string) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_color` IN (?)", roleColors).Find(&results).Error

	return
}

// GetFromRoleActor 通过role_actor获取内容 关联演员
func (obj *_MacRoleMgr) GetFromRoleActor(roleActor string) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_actor` = ?", roleActor).Find(&results).Error

	return
}

// GetBatchFromRoleActor 批量查找 关联演员
func (obj *_MacRoleMgr) GetBatchFromRoleActor(roleActors []string) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_actor` IN (?)", roleActors).Find(&results).Error

	return
}

// GetFromRoleRemarks 通过role_remarks获取内容 角色备注
func (obj *_MacRoleMgr) GetFromRoleRemarks(roleRemarks string) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_remarks` = ?", roleRemarks).Find(&results).Error

	return
}

// GetBatchFromRoleRemarks 批量查找 角色备注
func (obj *_MacRoleMgr) GetBatchFromRoleRemarks(roleRemarkss []string) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_remarks` IN (?)", roleRemarkss).Find(&results).Error

	return
}

// GetFromRolePic 通过role_pic获取内容 角色图片
func (obj *_MacRoleMgr) GetFromRolePic(rolePic string) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_pic` = ?", rolePic).Find(&results).Error

	return
}

// GetBatchFromRolePic 批量查找 角色图片
func (obj *_MacRoleMgr) GetBatchFromRolePic(rolePics []string) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_pic` IN (?)", rolePics).Find(&results).Error

	return
}

// GetFromRoleSort 通过role_sort获取内容 角色排序
func (obj *_MacRoleMgr) GetFromRoleSort(roleSort uint16) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_sort` = ?", roleSort).Find(&results).Error

	return
}

// GetBatchFromRoleSort 批量查找 角色排序
func (obj *_MacRoleMgr) GetBatchFromRoleSort(roleSorts []uint16) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_sort` IN (?)", roleSorts).Find(&results).Error

	return
}

// GetFromRoleLevel 通过role_level获取内容 角色等级
func (obj *_MacRoleMgr) GetFromRoleLevel(roleLevel uint8) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_level` = ?", roleLevel).Find(&results).Error

	return
}

// GetBatchFromRoleLevel 批量查找 角色等级
func (obj *_MacRoleMgr) GetBatchFromRoleLevel(roleLevels []uint8) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_level` IN (?)", roleLevels).Find(&results).Error

	return
}

// GetFromRoleTime 通过role_time获取内容 角色时间
func (obj *_MacRoleMgr) GetFromRoleTime(roleTime int) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_time` = ?", roleTime).Find(&results).Error

	return
}

// GetBatchFromRoleTime 批量查找 角色时间
func (obj *_MacRoleMgr) GetBatchFromRoleTime(roleTimes []int) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_time` IN (?)", roleTimes).Find(&results).Error

	return
}

// GetFromRoleTimeAdd 通过role_time_add获取内容 角色添加时间
func (obj *_MacRoleMgr) GetFromRoleTimeAdd(roleTimeAdd int) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_time_add` = ?", roleTimeAdd).Find(&results).Error

	return
}

// GetBatchFromRoleTimeAdd 批量查找 角色添加时间
func (obj *_MacRoleMgr) GetBatchFromRoleTimeAdd(roleTimeAdds []int) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_time_add` IN (?)", roleTimeAdds).Find(&results).Error

	return
}

// GetFromRoleTimeHits 通过role_time_hits获取内容 角色点击时间
func (obj *_MacRoleMgr) GetFromRoleTimeHits(roleTimeHits int) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_time_hits` = ?", roleTimeHits).Find(&results).Error

	return
}

// GetBatchFromRoleTimeHits 批量查找 角色点击时间
func (obj *_MacRoleMgr) GetBatchFromRoleTimeHits(roleTimeHitss []int) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_time_hits` IN (?)", roleTimeHitss).Find(&results).Error

	return
}

// GetFromRoleTimeMake 通过role_time_make获取内容 角色制作时间
func (obj *_MacRoleMgr) GetFromRoleTimeMake(roleTimeMake int) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_time_make` = ?", roleTimeMake).Find(&results).Error

	return
}

// GetBatchFromRoleTimeMake 批量查找 角色制作时间
func (obj *_MacRoleMgr) GetBatchFromRoleTimeMake(roleTimeMakes []int) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_time_make` IN (?)", roleTimeMakes).Find(&results).Error

	return
}

// GetFromRoleHits 通过role_hits获取内容 角色点击量
func (obj *_MacRoleMgr) GetFromRoleHits(roleHits int16) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_hits` = ?", roleHits).Find(&results).Error

	return
}

// GetBatchFromRoleHits 批量查找 角色点击量
func (obj *_MacRoleMgr) GetBatchFromRoleHits(roleHitss []int16) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_hits` IN (?)", roleHitss).Find(&results).Error

	return
}

// GetFromRoleHitsDay 通过role_hits_day获取内容 角色日点击量
func (obj *_MacRoleMgr) GetFromRoleHitsDay(roleHitsDay int16) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_hits_day` = ?", roleHitsDay).Find(&results).Error

	return
}

// GetBatchFromRoleHitsDay 批量查找 角色日点击量
func (obj *_MacRoleMgr) GetBatchFromRoleHitsDay(roleHitsDays []int16) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_hits_day` IN (?)", roleHitsDays).Find(&results).Error

	return
}

// GetFromRoleHitsWeek 通过role_hits_week获取内容 角色周点击量
func (obj *_MacRoleMgr) GetFromRoleHitsWeek(roleHitsWeek int16) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_hits_week` = ?", roleHitsWeek).Find(&results).Error

	return
}

// GetBatchFromRoleHitsWeek 批量查找 角色周点击量
func (obj *_MacRoleMgr) GetBatchFromRoleHitsWeek(roleHitsWeeks []int16) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_hits_week` IN (?)", roleHitsWeeks).Find(&results).Error

	return
}

// GetFromRoleHitsMonth 通过role_hits_month获取内容 角色月点击量
func (obj *_MacRoleMgr) GetFromRoleHitsMonth(roleHitsMonth int16) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_hits_month` = ?", roleHitsMonth).Find(&results).Error

	return
}

// GetBatchFromRoleHitsMonth 批量查找 角色月点击量
func (obj *_MacRoleMgr) GetBatchFromRoleHitsMonth(roleHitsMonths []int16) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_hits_month` IN (?)", roleHitsMonths).Find(&results).Error

	return
}

// GetFromRoleScore 通过role_score获取内容 角色评分
func (obj *_MacRoleMgr) GetFromRoleScore(roleScore float64) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_score` = ?", roleScore).Find(&results).Error

	return
}

// GetBatchFromRoleScore 批量查找 角色评分
func (obj *_MacRoleMgr) GetBatchFromRoleScore(roleScores []float64) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_score` IN (?)", roleScores).Find(&results).Error

	return
}

// GetFromRoleScoreAll 通过role_score_all获取内容 角色总评分
func (obj *_MacRoleMgr) GetFromRoleScoreAll(roleScoreAll int16) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_score_all` = ?", roleScoreAll).Find(&results).Error

	return
}

// GetBatchFromRoleScoreAll 批量查找 角色总评分
func (obj *_MacRoleMgr) GetBatchFromRoleScoreAll(roleScoreAlls []int16) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_score_all` IN (?)", roleScoreAlls).Find(&results).Error

	return
}

// GetFromRoleScoreNum 通过role_score_num获取内容 角色评分人数
func (obj *_MacRoleMgr) GetFromRoleScoreNum(roleScoreNum int16) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_score_num` = ?", roleScoreNum).Find(&results).Error

	return
}

// GetBatchFromRoleScoreNum 批量查找 角色评分人数
func (obj *_MacRoleMgr) GetBatchFromRoleScoreNum(roleScoreNums []int16) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_score_num` IN (?)", roleScoreNums).Find(&results).Error

	return
}

// GetFromRoleUp 通过role_up获取内容 角色点赞数
func (obj *_MacRoleMgr) GetFromRoleUp(roleUp int16) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_up` = ?", roleUp).Find(&results).Error

	return
}

// GetBatchFromRoleUp 批量查找 角色点赞数
func (obj *_MacRoleMgr) GetBatchFromRoleUp(roleUps []int16) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_up` IN (?)", roleUps).Find(&results).Error

	return
}

// GetFromRoleDown 通过role_down获取内容 角色踩数
func (obj *_MacRoleMgr) GetFromRoleDown(roleDown int16) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_down` = ?", roleDown).Find(&results).Error

	return
}

// GetBatchFromRoleDown 批量查找 角色踩数
func (obj *_MacRoleMgr) GetBatchFromRoleDown(roleDowns []int16) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_down` IN (?)", roleDowns).Find(&results).Error

	return
}

// GetFromRoleTpl 通过role_tpl获取内容 角色模板
func (obj *_MacRoleMgr) GetFromRoleTpl(roleTpl string) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_tpl` = ?", roleTpl).Find(&results).Error

	return
}

// GetBatchFromRoleTpl 批量查找 角色模板
func (obj *_MacRoleMgr) GetBatchFromRoleTpl(roleTpls []string) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_tpl` IN (?)", roleTpls).Find(&results).Error

	return
}

// GetFromRoleJumpurl 通过role_jumpurl获取内容 角色跳转URL
func (obj *_MacRoleMgr) GetFromRoleJumpurl(roleJumpurl string) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_jumpurl` = ?", roleJumpurl).Find(&results).Error

	return
}

// GetBatchFromRoleJumpurl 批量查找 角色跳转URL
func (obj *_MacRoleMgr) GetBatchFromRoleJumpurl(roleJumpurls []string) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_jumpurl` IN (?)", roleJumpurls).Find(&results).Error

	return
}

// GetFromRoleContent 通过role_content获取内容 角色内容
func (obj *_MacRoleMgr) GetFromRoleContent(roleContent string) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_content` = ?", roleContent).Find(&results).Error

	return
}

// GetBatchFromRoleContent 批量查找 角色内容
func (obj *_MacRoleMgr) GetBatchFromRoleContent(roleContents []string) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_content` IN (?)", roleContents).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MacRoleMgr) FetchByPrimaryKey(roleID int) (result MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_id` = ?", roleID).First(&result).Error

	return
}

// FetchIndexByRoleRid  获取多个内容
func (obj *_MacRoleMgr) FetchIndexByRoleRid(roleRid int) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_rid` = ?", roleRid).Find(&results).Error

	return
}

// FetchIndexByRoleName  获取多个内容
func (obj *_MacRoleMgr) FetchIndexByRoleName(roleName string) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_name` = ?", roleName).Find(&results).Error

	return
}

// FetchIndexByRoleEn  获取多个内容
func (obj *_MacRoleMgr) FetchIndexByRoleEn(roleEn string) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_en` = ?", roleEn).Find(&results).Error

	return
}

// FetchIndexByRoleLetter  获取多个内容
func (obj *_MacRoleMgr) FetchIndexByRoleLetter(roleLetter string) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_letter` = ?", roleLetter).Find(&results).Error

	return
}

// FetchIndexByRoleActor  获取多个内容
func (obj *_MacRoleMgr) FetchIndexByRoleActor(roleActor string) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_actor` = ?", roleActor).Find(&results).Error

	return
}

// FetchIndexByRoleLevel  获取多个内容
func (obj *_MacRoleMgr) FetchIndexByRoleLevel(roleLevel uint8) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_level` = ?", roleLevel).Find(&results).Error

	return
}

// FetchIndexByRoleTime  获取多个内容
func (obj *_MacRoleMgr) FetchIndexByRoleTime(roleTime int) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_time` = ?", roleTime).Find(&results).Error

	return
}

// FetchIndexByRoleTimeAdd  获取多个内容
func (obj *_MacRoleMgr) FetchIndexByRoleTimeAdd(roleTimeAdd int) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_time_add` = ?", roleTimeAdd).Find(&results).Error

	return
}

// FetchIndexByRoleScore  获取多个内容
func (obj *_MacRoleMgr) FetchIndexByRoleScore(roleScore float64) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_score` = ?", roleScore).Find(&results).Error

	return
}

// FetchIndexByRoleScoreAll  获取多个内容
func (obj *_MacRoleMgr) FetchIndexByRoleScoreAll(roleScoreAll int16) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_score_all` = ?", roleScoreAll).Find(&results).Error

	return
}

// FetchIndexByRoleScoreNum  获取多个内容
func (obj *_MacRoleMgr) FetchIndexByRoleScoreNum(roleScoreNum int16) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_score_num` = ?", roleScoreNum).Find(&results).Error

	return
}

// FetchIndexByRoleUp  获取多个内容
func (obj *_MacRoleMgr) FetchIndexByRoleUp(roleUp int16) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_up` = ?", roleUp).Find(&results).Error

	return
}

// FetchIndexByRoleDown  获取多个内容
func (obj *_MacRoleMgr) FetchIndexByRoleDown(roleDown int16) (results []*MacRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MacRole{}).Where("`role_down` = ?", roleDown).Find(&results).Error

	return
}
