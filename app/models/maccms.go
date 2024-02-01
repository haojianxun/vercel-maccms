package models

// MacActor [...]
type MacActor struct {
	ActorID        int     `gorm:"primaryKey;column:actor_id" json:"-"`           // 演员ID
	TypeID         uint16  `gorm:"column:type_id" json:"typeId"`                  // 类型ID
	TypeID1        uint16  `gorm:"column:type_id_1" json:"typeId1"`               // 另一类型ID
	ActorName      string  `gorm:"column:actor_name" json:"actorName"`            // 演员姓名
	ActorEn        string  `gorm:"column:actor_en" json:"actorEn"`                // 演员英文名
	ActorAlias     string  `gorm:"column:actor_alias" json:"actorAlias"`          // 演员别名
	ActorStatus    uint8   `gorm:"column:actor_status" json:"actorStatus"`        // 演员状态
	ActorLock      uint8   `gorm:"column:actor_lock" json:"actorLock"`            // 演员锁定状态
	ActorLetter    string  `gorm:"column:actor_letter" json:"actorLetter"`        // 演员姓氏首字母
	ActorSex       string  `gorm:"column:actor_sex" json:"actorSex"`              // 演员性别
	ActorColor     string  `gorm:"column:actor_color" json:"actorColor"`          // 演员颜色
	ActorPic       string  `gorm:"column:actor_pic" json:"actorPic"`              // 演员照片链接
	ActorBlurb     string  `gorm:"column:actor_blurb" json:"actorBlurb"`          // 演员简介
	ActorRemarks   string  `gorm:"column:actor_remarks" json:"actorRemarks"`      // 演员备注
	ActorArea      string  `gorm:"column:actor_area" json:"actorArea"`            // 演员地区
	ActorHeight    string  `gorm:"column:actor_height" json:"actorHeight"`        // 演员身高
	ActorWeight    string  `gorm:"column:actor_weight" json:"actorWeight"`        // 演员体重
	ActorBirthday  string  `gorm:"column:actor_birthday" json:"actorBirthday"`    // 演员生日
	ActorBirtharea string  `gorm:"column:actor_birtharea" json:"actorBirtharea"`  // 演员出生地区
	ActorBlood     string  `gorm:"column:actor_blood" json:"actorBlood"`          // 演员血型
	ActorStarsign  string  `gorm:"column:actor_starsign" json:"actorStarsign"`    // 演员星座
	ActorSchool    string  `gorm:"column:actor_school" json:"actorSchool"`        // 演员毕业学校
	ActorWorks     string  `gorm:"column:actor_works" json:"actorWorks"`          // 演员作品
	ActorTag       string  `gorm:"column:actor_tag" json:"actorTag"`              // 演员标签
	ActorClass     string  `gorm:"column:actor_class" json:"actorClass"`          // 演员类别
	ActorLevel     uint8   `gorm:"column:actor_level" json:"actorLevel"`          // 演员级别
	ActorTime      int     `gorm:"column:actor_time" json:"actorTime"`            // 演员时间
	ActorTimeAdd   int     `gorm:"column:actor_time_add" json:"actorTimeAdd"`     // 演员添加时间
	ActorTimeHits  int     `gorm:"column:actor_time_hits" json:"actorTimeHits"`   // 演员点击时间
	ActorTimeMake  int     `gorm:"column:actor_time_make" json:"actorTimeMake"`   // 演员制作时间
	ActorHits      int16   `gorm:"column:actor_hits" json:"actorHits"`            // 演员点击次数
	ActorHitsDay   int16   `gorm:"column:actor_hits_day" json:"actorHitsDay"`     // 演员每日点击次数
	ActorHitsWeek  int16   `gorm:"column:actor_hits_week" json:"actorHitsWeek"`   // 演员每周点击次数
	ActorHitsMonth int16   `gorm:"column:actor_hits_month" json:"actorHitsMonth"` // 演员每月点击次数
	ActorScore     float64 `gorm:"column:actor_score" json:"actorScore"`          // 演员评分
	ActorScoreAll  int16   `gorm:"column:actor_score_all" json:"actorScoreAll"`   // 演员总评分
	ActorScoreNum  int16   `gorm:"column:actor_score_num" json:"actorScoreNum"`   // 演员评分次数
	ActorUp        int16   `gorm:"column:actor_up" json:"actorUp"`                // 演员点赞次数
	ActorDown      int16   `gorm:"column:actor_down" json:"actorDown"`            // 演员踩次数
	ActorTpl       string  `gorm:"column:actor_tpl" json:"actorTpl"`              // 演员模板
	ActorJumpurl   string  `gorm:"column:actor_jumpurl" json:"actorJumpurl"`      // 演员跳转链接
	ActorContent   string  `gorm:"column:actor_content" json:"actorContent"`      // 演员内容
}

// TableName get sql table name.获取数据库表名
func (m *MacActor) TableName() string {
	return "mac_actor"
}

// MacActorColumns get sql column name.获取数据库列名
var MacActorColumns = struct {
	ActorID        string
	TypeID         string
	TypeID1        string
	ActorName      string
	ActorEn        string
	ActorAlias     string
	ActorStatus    string
	ActorLock      string
	ActorLetter    string
	ActorSex       string
	ActorColor     string
	ActorPic       string
	ActorBlurb     string
	ActorRemarks   string
	ActorArea      string
	ActorHeight    string
	ActorWeight    string
	ActorBirthday  string
	ActorBirtharea string
	ActorBlood     string
	ActorStarsign  string
	ActorSchool    string
	ActorWorks     string
	ActorTag       string
	ActorClass     string
	ActorLevel     string
	ActorTime      string
	ActorTimeAdd   string
	ActorTimeHits  string
	ActorTimeMake  string
	ActorHits      string
	ActorHitsDay   string
	ActorHitsWeek  string
	ActorHitsMonth string
	ActorScore     string
	ActorScoreAll  string
	ActorScoreNum  string
	ActorUp        string
	ActorDown      string
	ActorTpl       string
	ActorJumpurl   string
	ActorContent   string
}{
	ActorID:        "actor_id",
	TypeID:         "type_id",
	TypeID1:        "type_id_1",
	ActorName:      "actor_name",
	ActorEn:        "actor_en",
	ActorAlias:     "actor_alias",
	ActorStatus:    "actor_status",
	ActorLock:      "actor_lock",
	ActorLetter:    "actor_letter",
	ActorSex:       "actor_sex",
	ActorColor:     "actor_color",
	ActorPic:       "actor_pic",
	ActorBlurb:     "actor_blurb",
	ActorRemarks:   "actor_remarks",
	ActorArea:      "actor_area",
	ActorHeight:    "actor_height",
	ActorWeight:    "actor_weight",
	ActorBirthday:  "actor_birthday",
	ActorBirtharea: "actor_birtharea",
	ActorBlood:     "actor_blood",
	ActorStarsign:  "actor_starsign",
	ActorSchool:    "actor_school",
	ActorWorks:     "actor_works",
	ActorTag:       "actor_tag",
	ActorClass:     "actor_class",
	ActorLevel:     "actor_level",
	ActorTime:      "actor_time",
	ActorTimeAdd:   "actor_time_add",
	ActorTimeHits:  "actor_time_hits",
	ActorTimeMake:  "actor_time_make",
	ActorHits:      "actor_hits",
	ActorHitsDay:   "actor_hits_day",
	ActorHitsWeek:  "actor_hits_week",
	ActorHitsMonth: "actor_hits_month",
	ActorScore:     "actor_score",
	ActorScoreAll:  "actor_score_all",
	ActorScoreNum:  "actor_score_num",
	ActorUp:        "actor_up",
	ActorDown:      "actor_down",
	ActorTpl:       "actor_tpl",
	ActorJumpurl:   "actor_jumpurl",
	ActorContent:   "actor_content",
}

// MacAdmin [...]
type MacAdmin struct {
	AdminID            uint16 `gorm:"primaryKey;column:admin_id" json:"-"`                    // 管理员ID
	AdminName          string `gorm:"column:admin_name" json:"adminName"`                     // 管理员姓名
	AdminPwd           string `gorm:"column:admin_pwd" json:"adminPwd"`                       // 管理员密码
	AdminRandom        string `gorm:"column:admin_random" json:"adminRandom"`                 // 管理员随机值
	AdminStatus        uint8  `gorm:"column:admin_status" json:"adminStatus"`                 // 管理员状态
	AdminAuth          string `gorm:"column:admin_auth" json:"adminAuth"`                     // 管理员权限
	AdminLoginTime     int    `gorm:"column:admin_login_time" json:"adminLoginTime"`          // 管理员登录时间
	AdminLoginIP       int    `gorm:"column:admin_login_ip" json:"adminLoginIp"`              // 管理员登录IP
	AdminLoginNum      int    `gorm:"column:admin_login_num" json:"adminLoginNum"`            // 管理员登录次数
	AdminLastLoginTime int    `gorm:"column:admin_last_login_time" json:"adminLastLoginTime"` // 管理员上次登录时间
	AdminLastLoginIP   int    `gorm:"column:admin_last_login_ip" json:"adminLastLoginIp"`     // 管理员上次登录IP
}

// TableName get sql table name.获取数据库表名
func (m *MacAdmin) TableName() string {
	return "mac_admin"
}

// MacAdminColumns get sql column name.获取数据库列名
var MacAdminColumns = struct {
	AdminID            string
	AdminName          string
	AdminPwd           string
	AdminRandom        string
	AdminStatus        string
	AdminAuth          string
	AdminLoginTime     string
	AdminLoginIP       string
	AdminLoginNum      string
	AdminLastLoginTime string
	AdminLastLoginIP   string
}{
	AdminID:            "admin_id",
	AdminName:          "admin_name",
	AdminPwd:           "admin_pwd",
	AdminRandom:        "admin_random",
	AdminStatus:        "admin_status",
	AdminAuth:          "admin_auth",
	AdminLoginTime:     "admin_login_time",
	AdminLoginIP:       "admin_login_ip",
	AdminLoginNum:      "admin_login_num",
	AdminLastLoginTime: "admin_last_login_time",
	AdminLastLoginIP:   "admin_last_login_ip",
}

// MacAnnex [...]
type MacAnnex struct {
	AnnexID   int    `gorm:"primaryKey;column:annex_id" json:"-"` // 附件ID
	AnnexTime int    `gorm:"column:annex_time" json:"annexTime"`  // 附件上传时间
	AnnexFile string `gorm:"column:annex_file" json:"annexFile"`  // 附件文件名
	AnnexSize int    `gorm:"column:annex_size" json:"annexSize"`  // 附件大小
	AnnexType string `gorm:"column:annex_type" json:"annexType"`  // 附件类型
}

// TableName get sql table name.获取数据库表名
func (m *MacAnnex) TableName() string {
	return "mac_annex"
}

// MacAnnexColumns get sql column name.获取数据库列名
var MacAnnexColumns = struct {
	AnnexID   string
	AnnexTime string
	AnnexFile string
	AnnexSize string
	AnnexType string
}{
	AnnexID:   "annex_id",
	AnnexTime: "annex_time",
	AnnexFile: "annex_file",
	AnnexSize: "annex_size",
	AnnexType: "annex_type",
}

// MacArt [...]
type MacArt struct {
	ArtID            int     `gorm:"primaryKey;column:art_id" json:"-"`                 // 文章ID
	TypeID           uint16  `gorm:"column:type_id" json:"typeId"`                      // 类型ID
	TypeID1          uint16  `gorm:"column:type_id_1" json:"typeId1"`                   // 另一类型ID
	GroupID          uint16  `gorm:"column:group_id" json:"groupId"`                    // 分组ID
	ArtName          string  `gorm:"column:art_name" json:"artName"`                    // 文章名称
	ArtSub           string  `gorm:"column:art_sub" json:"artSub"`                      // 文章副标题
	ArtEn            string  `gorm:"column:art_en" json:"artEn"`                        // 文章英文名
	ArtStatus        uint8   `gorm:"column:art_status" json:"artStatus"`                // 文章状态
	ArtLetter        string  `gorm:"column:art_letter" json:"artLetter"`                // 文章姓氏首字母
	ArtColor         string  `gorm:"column:art_color" json:"artColor"`                  // 文章颜色
	ArtFrom          string  `gorm:"column:art_from" json:"artFrom"`                    // 文章来源
	ArtAuthor        string  `gorm:"column:art_author" json:"artAuthor"`                // 文章作者
	ArtTag           string  `gorm:"column:art_tag" json:"artTag"`                      // 文章标签
	ArtClass         string  `gorm:"column:art_class" json:"artClass"`                  // 文章类别
	ArtPic           string  `gorm:"column:art_pic" json:"artPic"`                      // 文章图片链接
	ArtPicThumb      string  `gorm:"column:art_pic_thumb" json:"artPicThumb"`           // 文章缩略图链接
	ArtPicSlide      string  `gorm:"column:art_pic_slide" json:"artPicSlide"`           // 文章幻灯片图片链接
	ArtPicScreenshot string  `gorm:"column:art_pic_screenshot" json:"artPicScreenshot"` // 文章截图图片链接
	ArtBlurb         string  `gorm:"column:art_blurb" json:"artBlurb"`                  // 文章简介
	ArtRemarks       string  `gorm:"column:art_remarks" json:"artRemarks"`              // 文章备注
	ArtJumpurl       string  `gorm:"column:art_jumpurl" json:"artJumpurl"`              // 文章跳转链接
	ArtTpl           string  `gorm:"column:art_tpl" json:"artTpl"`                      // 文章模板
	ArtLevel         uint8   `gorm:"column:art_level" json:"artLevel"`                  // 文章级别
	ArtLock          uint8   `gorm:"column:art_lock" json:"artLock"`                    // 文章锁定状态
	ArtPoints        uint16  `gorm:"column:art_points" json:"artPoints"`                // 文章积分
	ArtPointsDetail  uint16  `gorm:"column:art_points_detail" json:"artPointsDetail"`   // 文章详细积分
	ArtUp            int16   `gorm:"column:art_up" json:"artUp"`                        // 文章点赞次数
	ArtDown          int16   `gorm:"column:art_down" json:"artDown"`                    // 文章踩次数
	ArtHits          int16   `gorm:"column:art_hits" json:"artHits"`                    // 文章点击次数
	ArtHitsDay       int16   `gorm:"column:art_hits_day" json:"artHitsDay"`             // 文章每日点击次数
	ArtHitsWeek      int16   `gorm:"column:art_hits_week" json:"artHitsWeek"`           // 文章每周点击次数
	ArtHitsMonth     int16   `gorm:"column:art_hits_month" json:"artHitsMonth"`         // 文章每月点击次数
	ArtTime          int     `gorm:"column:art_time" json:"artTime"`                    // 文章时间
	ArtTimeAdd       int     `gorm:"column:art_time_add" json:"artTimeAdd"`             // 文章添加时间
	ArtTimeHits      int     `gorm:"column:art_time_hits" json:"artTimeHits"`           // 文章点击时间
	ArtTimeMake      int     `gorm:"column:art_time_make" json:"artTimeMake"`           // 文章制作时间
	ArtScore         float64 `gorm:"column:art_score" json:"artScore"`                  // 文章评分
	ArtScoreAll      int16   `gorm:"column:art_score_all" json:"artScoreAll"`           // 文章总评分
	ArtScoreNum      int16   `gorm:"column:art_score_num" json:"artScoreNum"`           // 文章评分次数
	ArtRelArt        string  `gorm:"column:art_rel_art" json:"artRelArt"`               // 相关文章链接
	ArtRelVod        string  `gorm:"column:art_rel_vod" json:"artRelVod"`               // 相关视频链接
	ArtPwd           string  `gorm:"column:art_pwd" json:"artPwd"`                      // 文章密码
	ArtPwdURL        string  `gorm:"column:art_pwd_url" json:"artPwdUrl"`               // 文章密码链接
	ArtTitle         string  `gorm:"column:art_title" json:"artTitle"`                  // 文章标题
	ArtNote          string  `gorm:"column:art_note" json:"artNote"`                    // 文章注释
	ArtContent       string  `gorm:"column:art_content" json:"artContent"`              // 文章内容
}

// TableName get sql table name.获取数据库表名
func (m *MacArt) TableName() string {
	return "mac_art"
}

// MacArtColumns get sql column name.获取数据库列名
var MacArtColumns = struct {
	ArtID            string
	TypeID           string
	TypeID1          string
	GroupID          string
	ArtName          string
	ArtSub           string
	ArtEn            string
	ArtStatus        string
	ArtLetter        string
	ArtColor         string
	ArtFrom          string
	ArtAuthor        string
	ArtTag           string
	ArtClass         string
	ArtPic           string
	ArtPicThumb      string
	ArtPicSlide      string
	ArtPicScreenshot string
	ArtBlurb         string
	ArtRemarks       string
	ArtJumpurl       string
	ArtTpl           string
	ArtLevel         string
	ArtLock          string
	ArtPoints        string
	ArtPointsDetail  string
	ArtUp            string
	ArtDown          string
	ArtHits          string
	ArtHitsDay       string
	ArtHitsWeek      string
	ArtHitsMonth     string
	ArtTime          string
	ArtTimeAdd       string
	ArtTimeHits      string
	ArtTimeMake      string
	ArtScore         string
	ArtScoreAll      string
	ArtScoreNum      string
	ArtRelArt        string
	ArtRelVod        string
	ArtPwd           string
	ArtPwdURL        string
	ArtTitle         string
	ArtNote          string
	ArtContent       string
}{
	ArtID:            "art_id",
	TypeID:           "type_id",
	TypeID1:          "type_id_1",
	GroupID:          "group_id",
	ArtName:          "art_name",
	ArtSub:           "art_sub",
	ArtEn:            "art_en",
	ArtStatus:        "art_status",
	ArtLetter:        "art_letter",
	ArtColor:         "art_color",
	ArtFrom:          "art_from",
	ArtAuthor:        "art_author",
	ArtTag:           "art_tag",
	ArtClass:         "art_class",
	ArtPic:           "art_pic",
	ArtPicThumb:      "art_pic_thumb",
	ArtPicSlide:      "art_pic_slide",
	ArtPicScreenshot: "art_pic_screenshot",
	ArtBlurb:         "art_blurb",
	ArtRemarks:       "art_remarks",
	ArtJumpurl:       "art_jumpurl",
	ArtTpl:           "art_tpl",
	ArtLevel:         "art_level",
	ArtLock:          "art_lock",
	ArtPoints:        "art_points",
	ArtPointsDetail:  "art_points_detail",
	ArtUp:            "art_up",
	ArtDown:          "art_down",
	ArtHits:          "art_hits",
	ArtHitsDay:       "art_hits_day",
	ArtHitsWeek:      "art_hits_week",
	ArtHitsMonth:     "art_hits_month",
	ArtTime:          "art_time",
	ArtTimeAdd:       "art_time_add",
	ArtTimeHits:      "art_time_hits",
	ArtTimeMake:      "art_time_make",
	ArtScore:         "art_score",
	ArtScoreAll:      "art_score_all",
	ArtScoreNum:      "art_score_num",
	ArtRelArt:        "art_rel_art",
	ArtRelVod:        "art_rel_vod",
	ArtPwd:           "art_pwd",
	ArtPwdURL:        "art_pwd_url",
	ArtTitle:         "art_title",
	ArtNote:          "art_note",
	ArtContent:       "art_content",
}

// MacCard [...]
type MacCard struct {
	CardID         int    `gorm:"primaryKey;column:card_id" json:"-"`            // 卡片ID
	CardNo         string `gorm:"column:card_no" json:"cardNo"`                  // 卡号
	CardPwd        string `gorm:"column:card_pwd" json:"cardPwd"`                // 卡密码
	CardMoney      uint16 `gorm:"column:card_money" json:"cardMoney"`            // 卡上余额
	CardPoints     uint16 `gorm:"column:card_points" json:"cardPoints"`          // 卡积分
	CardUseStatus  uint8  `gorm:"column:card_use_status" json:"cardUseStatus"`   // 卡使用状态
	CardSaleStatus uint8  `gorm:"column:card_sale_status" json:"cardSaleStatus"` // 卡销售状态
	UserID         int    `gorm:"column:user_id" json:"userId"`                  // 用户ID
	CardAddTime    int    `gorm:"column:card_add_time" json:"cardAddTime"`       // 卡添加时间
	CardUseTime    int    `gorm:"column:card_use_time" json:"cardUseTime"`       // 卡使用时间
}

// TableName get sql table name.获取数据库表名
func (m *MacCard) TableName() string {
	return "mac_card"
}

// MacCardColumns get sql column name.获取数据库列名
var MacCardColumns = struct {
	CardID         string
	CardNo         string
	CardPwd        string
	CardMoney      string
	CardPoints     string
	CardUseStatus  string
	CardSaleStatus string
	UserID         string
	CardAddTime    string
	CardUseTime    string
}{
	CardID:         "card_id",
	CardNo:         "card_no",
	CardPwd:        "card_pwd",
	CardMoney:      "card_money",
	CardPoints:     "card_points",
	CardUseStatus:  "card_use_status",
	CardSaleStatus: "card_sale_status",
	UserID:         "user_id",
	CardAddTime:    "card_add_time",
	CardUseTime:    "card_use_time",
}

// MacCash [...]
type MacCash struct {
	CashID        int     `gorm:"primaryKey;column:cash_id" json:"-"`          // 提现记录ID
	UserID        int     `gorm:"column:user_id" json:"userId"`                // 用户ID
	CashStatus    uint8   `gorm:"column:cash_status" json:"cashStatus"`        // 提现状态
	CashPoints    uint16  `gorm:"column:cash_points" json:"cashPoints"`        // 提现积分
	CashMoney     float64 `gorm:"column:cash_money" json:"cashMoney"`          // 提现金额
	CashBankName  string  `gorm:"column:cash_bank_name" json:"cashBankName"`   // 提现银行名称
	CashBankNo    string  `gorm:"column:cash_bank_no" json:"cashBankNo"`       // 提现银行账号
	CashPayeeName string  `gorm:"column:cash_payee_name" json:"cashPayeeName"` // 提现收款人姓名
	CashTime      int     `gorm:"column:cash_time" json:"cashTime"`            // 提现申请时间
	CashTimeAudit int     `gorm:"column:cash_time_audit" json:"cashTimeAudit"` // 提现审核时间
}

// TableName get sql table name.获取数据库表名
func (m *MacCash) TableName() string {
	return "mac_cash"
}

// MacCashColumns get sql column name.获取数据库列名
var MacCashColumns = struct {
	CashID        string
	UserID        string
	CashStatus    string
	CashPoints    string
	CashMoney     string
	CashBankName  string
	CashBankNo    string
	CashPayeeName string
	CashTime      string
	CashTimeAudit string
}{
	CashID:        "cash_id",
	UserID:        "user_id",
	CashStatus:    "cash_status",
	CashPoints:    "cash_points",
	CashMoney:     "cash_money",
	CashBankName:  "cash_bank_name",
	CashBankNo:    "cash_bank_no",
	CashPayeeName: "cash_payee_name",
	CashTime:      "cash_time",
	CashTimeAudit: "cash_time_audit",
}

// MacCjContent [...]
type MacCjContent struct {
	ID     int    `gorm:"primaryKey;column:id" json:"-"` // 内容ID
	Nodeid int    `gorm:"column:nodeid" json:"nodeid"`   // 节点ID
	Status uint8  `gorm:"column:status" json:"status"`   // 状态
	URL    string `gorm:"column:url" json:"url"`         // URL
	Title  string `gorm:"column:title" json:"title"`     // 标题
	Data   string `gorm:"column:data" json:"data"`       // 数据
}

// TableName get sql table name.获取数据库表名
func (m *MacCjContent) TableName() string {
	return "mac_cj_content"
}

// MacCjContentColumns get sql column name.获取数据库列名
var MacCjContentColumns = struct {
	ID     string
	Nodeid string
	Status string
	URL    string
	Title  string
	Data   string
}{
	ID:     "id",
	Nodeid: "nodeid",
	Status: "status",
	URL:    "url",
	Title:  "title",
	Data:   "data",
}

// MacCjHistory [...]
type MacCjHistory struct {
	Md5 string `gorm:"primaryKey;column:md5" json:"-"` // MD5值
}

// TableName get sql table name.获取数据库表名
func (m *MacCjHistory) TableName() string {
	return "mac_cj_history"
}

// MacCjHistoryColumns get sql column name.获取数据库列名
var MacCjHistoryColumns = struct {
	Md5 string
}{
	Md5: "md5",
}

// MacCjNode [...]
type MacCjNode struct {
	Nodeid           uint16 `gorm:"primaryKey;column:nodeid" json:"-"`                 // 节点ID
	Name             string `gorm:"column:name" json:"name"`                           // 节点名称
	Lastdate         int    `gorm:"column:lastdate" json:"lastdate"`                   // 最后更新日期
	Sourcecharset    string `gorm:"column:sourcecharset" json:"sourcecharset"`         // 源字符集
	Sourcetype       uint8  `gorm:"column:sourcetype" json:"sourcetype"`               // 源类型
	URLpage          string `gorm:"column:urlpage" json:"urlpage"`                     // URL页面
	PagesizeStart    uint8  `gorm:"column:pagesize_start" json:"pagesizeStart"`        // 页面大小起始值
	PagesizeEnd      int16  `gorm:"column:pagesize_end" json:"pagesizeEnd"`            // 页面大小结束值
	PageBase         string `gorm:"column:page_base" json:"pageBase"`                  // 页面基础URL
	ParNum           uint8  `gorm:"column:par_num" json:"parNum"`                      // 参数数量
	URLContain       string `gorm:"column:url_contain" json:"urlContain"`              // URL包含规则
	URLExcept        string `gorm:"column:url_except" json:"urlExcept"`                // URL排除规则
	URLStart         string `gorm:"column:url_start" json:"urlStart"`                  // URL起始规则
	URLEnd           string `gorm:"column:url_end" json:"urlEnd"`                      // URL结束规则
	TitleRule        string `gorm:"column:title_rule" json:"titleRule"`                // 标题规则
	TitleHTMLRule    string `gorm:"column:title_html_rule" json:"titleHtmlRule"`       // 标题HTML规则
	TypeRule         string `gorm:"column:type_rule" json:"typeRule"`                  // 类型规则
	TypeHTMLRule     string `gorm:"column:type_html_rule" json:"typeHtmlRule"`         // 类型HTML规则
	ContentRule      string `gorm:"column:content_rule" json:"contentRule"`            // 内容规则
	ContentHTMLRule  string `gorm:"column:content_html_rule" json:"contentHtmlRule"`   // 内容HTML规则
	ContentPageStart string `gorm:"column:content_page_start" json:"contentPageStart"` // 内容页起始规则
	ContentPageEnd   string `gorm:"column:content_page_end" json:"contentPageEnd"`     // 内容页结束规则
	ContentPageRule  uint8  `gorm:"column:content_page_rule" json:"contentPageRule"`   // 内容页规则
	ContentPage      uint8  `gorm:"column:content_page" json:"contentPage"`            // 内容页
	ContentNextpage  string `gorm:"column:content_nextpage" json:"contentNextpage"`    // 下一页规则
	DownAttachment   uint8  `gorm:"column:down_attachment" json:"downAttachment"`      // 下载附件
	Watermark        uint8  `gorm:"column:watermark" json:"watermark"`                 // 水印
	CollOrder        uint8  `gorm:"column:coll_order" json:"collOrder"`                // 采集顺序
	CustomizeConfig  string `gorm:"column:customize_config" json:"customizeConfig"`    // 自定义配置
	ProgramConfig    string `gorm:"column:program_config" json:"programConfig"`        // 程序配置
	Mid              uint8  `gorm:"column:mid" json:"mid"`                             // 模型ID
}

// TableName get sql table name.获取数据库表名
func (m *MacCjNode) TableName() string {
	return "mac_cj_node"
}

// MacCjNodeColumns get sql column name.获取数据库列名
var MacCjNodeColumns = struct {
	Nodeid           string
	Name             string
	Lastdate         string
	Sourcecharset    string
	Sourcetype       string
	URLpage          string
	PagesizeStart    string
	PagesizeEnd      string
	PageBase         string
	ParNum           string
	URLContain       string
	URLExcept        string
	URLStart         string
	URLEnd           string
	TitleRule        string
	TitleHTMLRule    string
	TypeRule         string
	TypeHTMLRule     string
	ContentRule      string
	ContentHTMLRule  string
	ContentPageStart string
	ContentPageEnd   string
	ContentPageRule  string
	ContentPage      string
	ContentNextpage  string
	DownAttachment   string
	Watermark        string
	CollOrder        string
	CustomizeConfig  string
	ProgramConfig    string
	Mid              string
}{
	Nodeid:           "nodeid",
	Name:             "name",
	Lastdate:         "lastdate",
	Sourcecharset:    "sourcecharset",
	Sourcetype:       "sourcetype",
	URLpage:          "urlpage",
	PagesizeStart:    "pagesize_start",
	PagesizeEnd:      "pagesize_end",
	PageBase:         "page_base",
	ParNum:           "par_num",
	URLContain:       "url_contain",
	URLExcept:        "url_except",
	URLStart:         "url_start",
	URLEnd:           "url_end",
	TitleRule:        "title_rule",
	TitleHTMLRule:    "title_html_rule",
	TypeRule:         "type_rule",
	TypeHTMLRule:     "type_html_rule",
	ContentRule:      "content_rule",
	ContentHTMLRule:  "content_html_rule",
	ContentPageStart: "content_page_start",
	ContentPageEnd:   "content_page_end",
	ContentPageRule:  "content_page_rule",
	ContentPage:      "content_page",
	ContentNextpage:  "content_nextpage",
	DownAttachment:   "down_attachment",
	Watermark:        "watermark",
	CollOrder:        "coll_order",
	CustomizeConfig:  "customize_config",
	ProgramConfig:    "program_config",
	Mid:              "mid",
}

// MacCollect [...]
type MacCollect struct {
	CollectID         int    `gorm:"primaryKey;column:collect_id" json:"-"`                // 采集ID
	CollectName       string `gorm:"column:collect_name" json:"collectName"`               // 采集名称
	CollectURL        string `gorm:"column:collect_url" json:"collectUrl"`                 // 采集URL
	CollectType       uint8  `gorm:"column:collect_type" json:"collectType"`               // 采集类型
	CollectMid        uint8  `gorm:"column:collect_mid" json:"collectMid"`                 // 模型ID
	CollectAppid      string `gorm:"column:collect_appid" json:"collectAppid"`             // 采集AppID
	CollectAppkey     string `gorm:"column:collect_appkey" json:"collectAppkey"`           // 采集AppKey
	CollectParam      string `gorm:"column:collect_param" json:"collectParam"`             // 采集参数
	CollectFilter     uint8  `gorm:"column:collect_filter" json:"collectFilter"`           // 采集过滤
	CollectFilterFrom string `gorm:"column:collect_filter_from" json:"collectFilterFrom"`  // 采集过滤来源
	CollectFilterYear string `gorm:"column:collect_filter_year" json:"collectFilterYear"`  // 采集时，过滤年份
	CollectOpt        uint8  `gorm:"column:collect_opt" json:"collectOpt"`                 // 采集选项
	CollectSyncPicOpt uint8  `gorm:"column:collect_sync_pic_opt" json:"collectSyncPicOpt"` // 同步图片选项，0-跟随全局，1-开启，2-关闭
}

// TableName get sql table name.获取数据库表名
func (m *MacCollect) TableName() string {
	return "mac_collect"
}

// MacCollectColumns get sql column name.获取数据库列名
var MacCollectColumns = struct {
	CollectID         string
	CollectName       string
	CollectURL        string
	CollectType       string
	CollectMid        string
	CollectAppid      string
	CollectAppkey     string
	CollectParam      string
	CollectFilter     string
	CollectFilterFrom string
	CollectFilterYear string
	CollectOpt        string
	CollectSyncPicOpt string
}{
	CollectID:         "collect_id",
	CollectName:       "collect_name",
	CollectURL:        "collect_url",
	CollectType:       "collect_type",
	CollectMid:        "collect_mid",
	CollectAppid:      "collect_appid",
	CollectAppkey:     "collect_appkey",
	CollectParam:      "collect_param",
	CollectFilter:     "collect_filter",
	CollectFilterFrom: "collect_filter_from",
	CollectFilterYear: "collect_filter_year",
	CollectOpt:        "collect_opt",
	CollectSyncPicOpt: "collect_sync_pic_opt",
}

// MacComment [...]
type MacComment struct {
	CommentID      int    `gorm:"primaryKey;column:comment_id" json:"-"`        // 评论ID
	CommentMid     uint8  `gorm:"column:comment_mid" json:"commentMid"`         // 模型ID
	CommentRid     int    `gorm:"column:comment_rid" json:"commentRid"`         // 回复评论ID
	CommentPid     int    `gorm:"column:comment_pid" json:"commentPid"`         // 父评论ID
	UserID         int    `gorm:"column:user_id" json:"userId"`                 // 用户ID
	CommentStatus  uint8  `gorm:"column:comment_status" json:"commentStatus"`   // 评论状态
	CommentName    string `gorm:"column:comment_name" json:"commentName"`       // 评论者姓名
	CommentIP      int    `gorm:"column:comment_ip" json:"commentIp"`           // 评论者IP
	CommentTime    int    `gorm:"column:comment_time" json:"commentTime"`       // 评论时间
	CommentContent string `gorm:"column:comment_content" json:"commentContent"` // 评论内容
	CommentUp      int16  `gorm:"column:comment_up" json:"commentUp"`           // 评论点赞数
	CommentDown    int16  `gorm:"column:comment_down" json:"commentDown"`       // 评论踩数
	CommentReply   int16  `gorm:"column:comment_reply" json:"commentReply"`     // 评论回复数
	CommentReport  int16  `gorm:"column:comment_report" json:"commentReport"`   // 评论举报数
}

// TableName get sql table name.获取数据库表名
func (m *MacComment) TableName() string {
	return "mac_comment"
}

// MacCommentColumns get sql column name.获取数据库列名
var MacCommentColumns = struct {
	CommentID      string
	CommentMid     string
	CommentRid     string
	CommentPid     string
	UserID         string
	CommentStatus  string
	CommentName    string
	CommentIP      string
	CommentTime    string
	CommentContent string
	CommentUp      string
	CommentDown    string
	CommentReply   string
	CommentReport  string
}{
	CommentID:      "comment_id",
	CommentMid:     "comment_mid",
	CommentRid:     "comment_rid",
	CommentPid:     "comment_pid",
	UserID:         "user_id",
	CommentStatus:  "comment_status",
	CommentName:    "comment_name",
	CommentIP:      "comment_ip",
	CommentTime:    "comment_time",
	CommentContent: "comment_content",
	CommentUp:      "comment_up",
	CommentDown:    "comment_down",
	CommentReply:   "comment_reply",
	CommentReport:  "comment_report",
}

// MacGbook [...]
type MacGbook struct {
	GbookID        int    `gorm:"primaryKey;column:gbook_id" json:"-"`           // 留言ID
	GbookRid       int    `gorm:"column:gbook_rid" json:"gbookRid"`              // 回复留言ID
	UserID         int    `gorm:"column:user_id" json:"userId"`                  // 用户ID
	GbookStatus    uint8  `gorm:"column:gbook_status" json:"gbookStatus"`        // 留言状态
	GbookName      string `gorm:"column:gbook_name" json:"gbookName"`            // 留言者姓名
	GbookIP        int    `gorm:"column:gbook_ip" json:"gbookIp"`                // 留言者IP
	GbookTime      int    `gorm:"column:gbook_time" json:"gbookTime"`            // 留言时间
	GbookReplyTime int    `gorm:"column:gbook_reply_time" json:"gbookReplyTime"` // 留言回复时间
	GbookContent   string `gorm:"column:gbook_content" json:"gbookContent"`      // 留言内容
	GbookReply     string `gorm:"column:gbook_reply" json:"gbookReply"`          // 留言回复
}

// TableName get sql table name.获取数据库表名
func (m *MacGbook) TableName() string {
	return "mac_gbook"
}

// MacGbookColumns get sql column name.获取数据库列名
var MacGbookColumns = struct {
	GbookID        string
	GbookRid       string
	UserID         string
	GbookStatus    string
	GbookName      string
	GbookIP        string
	GbookTime      string
	GbookReplyTime string
	GbookContent   string
	GbookReply     string
}{
	GbookID:        "gbook_id",
	GbookRid:       "gbook_rid",
	UserID:         "user_id",
	GbookStatus:    "gbook_status",
	GbookName:      "gbook_name",
	GbookIP:        "gbook_ip",
	GbookTime:      "gbook_time",
	GbookReplyTime: "gbook_reply_time",
	GbookContent:   "gbook_content",
	GbookReply:     "gbook_reply",
}

// MacGroup [...]
type MacGroup struct {
	GroupID          int16  `gorm:"primaryKey;column:group_id" json:"-"`               // 用户组ID
	GroupName        string `gorm:"column:group_name" json:"groupName"`                // 用户组名称
	GroupStatus      uint8  `gorm:"column:group_status" json:"groupStatus"`            // 用户组状态
	GroupType        string `gorm:"column:group_type" json:"groupType"`                // 用户组类型
	GroupPopedom     string `gorm:"column:group_popedom" json:"groupPopedom"`          // 用户组权限
	GroupPointsDay   uint16 `gorm:"column:group_points_day" json:"groupPointsDay"`     // 用户组每日积分
	GroupPointsWeek  int16  `gorm:"column:group_points_week" json:"groupPointsWeek"`   // 用户组每周积分
	GroupPointsMonth uint16 `gorm:"column:group_points_month" json:"groupPointsMonth"` // 用户组每月积分
	GroupPointsYear  uint16 `gorm:"column:group_points_year" json:"groupPointsYear"`   // 用户组每年积分
	GroupPointsFree  uint8  `gorm:"column:group_points_free" json:"groupPointsFree"`   // 用户组免费积分
}

// TableName get sql table name.获取数据库表名
func (m *MacGroup) TableName() string {
	return "mac_group"
}

// MacGroupColumns get sql column name.获取数据库列名
var MacGroupColumns = struct {
	GroupID          string
	GroupName        string
	GroupStatus      string
	GroupType        string
	GroupPopedom     string
	GroupPointsDay   string
	GroupPointsWeek  string
	GroupPointsMonth string
	GroupPointsYear  string
	GroupPointsFree  string
}{
	GroupID:          "group_id",
	GroupName:        "group_name",
	GroupStatus:      "group_status",
	GroupType:        "group_type",
	GroupPopedom:     "group_popedom",
	GroupPointsDay:   "group_points_day",
	GroupPointsWeek:  "group_points_week",
	GroupPointsMonth: "group_points_month",
	GroupPointsYear:  "group_points_year",
	GroupPointsFree:  "group_points_free",
}

// MacLink [...]
type MacLink struct {
	LinkID      uint16 `gorm:"primaryKey;column:link_id" json:"-"`      // 友情链接ID
	LinkType    uint8  `gorm:"column:link_type" json:"linkType"`        // 链接类型
	LinkName    string `gorm:"column:link_name" json:"linkName"`        // 链接名称
	LinkSort    int16  `gorm:"column:link_sort" json:"linkSort"`        // 链接排序
	LinkAddTime int    `gorm:"column:link_add_time" json:"linkAddTime"` // 链接添加时间
	LinkTime    int    `gorm:"column:link_time" json:"linkTime"`        // 链接更新时间
	LinkURL     string `gorm:"column:link_url" json:"linkUrl"`          // 链接URL
	LinkLogo    string `gorm:"column:link_logo" json:"linkLogo"`        // 链接Logo
}

// TableName get sql table name.获取数据库表名
func (m *MacLink) TableName() string {
	return "mac_link"
}

// MacLinkColumns get sql column name.获取数据库列名
var MacLinkColumns = struct {
	LinkID      string
	LinkType    string
	LinkName    string
	LinkSort    string
	LinkAddTime string
	LinkTime    string
	LinkURL     string
	LinkLogo    string
}{
	LinkID:      "link_id",
	LinkType:    "link_type",
	LinkName:    "link_name",
	LinkSort:    "link_sort",
	LinkAddTime: "link_add_time",
	LinkTime:    "link_time",
	LinkURL:     "link_url",
	LinkLogo:    "link_logo",
}

// MacMsg [...]
type MacMsg struct {
	MsgID      int    `gorm:"primaryKey;column:msg_id" json:"-"`    // 消息ID
	UserID     int    `gorm:"column:user_id" json:"userId"`         // 用户ID
	MsgType    uint8  `gorm:"column:msg_type" json:"msgType"`       // 消息类型
	MsgStatus  uint8  `gorm:"column:msg_status" json:"msgStatus"`   // 消息状态
	MsgTo      string `gorm:"column:msg_to" json:"msgTo"`           // 消息接收者
	MsgCode    string `gorm:"column:msg_code" json:"msgCode"`       // 消息代码
	MsgContent string `gorm:"column:msg_content" json:"msgContent"` // 消息内容
	MsgTime    int    `gorm:"column:msg_time" json:"msgTime"`       // 消息时间
}

// TableName get sql table name.获取数据库表名
func (m *MacMsg) TableName() string {
	return "mac_msg"
}

// MacMsgColumns get sql column name.获取数据库列名
var MacMsgColumns = struct {
	MsgID      string
	UserID     string
	MsgType    string
	MsgStatus  string
	MsgTo      string
	MsgCode    string
	MsgContent string
	MsgTime    string
}{
	MsgID:      "msg_id",
	UserID:     "user_id",
	MsgType:    "msg_type",
	MsgStatus:  "msg_status",
	MsgTo:      "msg_to",
	MsgCode:    "msg_code",
	MsgContent: "msg_content",
	MsgTime:    "msg_time",
}

// MacOrder [...]
type MacOrder struct {
	OrderID      int     `gorm:"primaryKey;column:order_id" json:"-"`       // 订单ID
	UserID       int     `gorm:"column:user_id" json:"userId"`              // 用户ID
	OrderStatus  uint8   `gorm:"column:order_status" json:"orderStatus"`    // 订单状态
	OrderCode    string  `gorm:"column:order_code" json:"orderCode"`        // 订单代码
	OrderPrice   float64 `gorm:"column:order_price" json:"orderPrice"`      // 订单价格
	OrderTime    int     `gorm:"column:order_time" json:"orderTime"`        // 订单时间
	OrderPoints  int16   `gorm:"column:order_points" json:"orderPoints"`    // 订单积分
	OrderPayType string  `gorm:"column:order_pay_type" json:"orderPayType"` // 支付方式
	OrderPayTime int     `gorm:"column:order_pay_time" json:"orderPayTime"` // 支付时间
	OrderRemarks string  `gorm:"column:order_remarks" json:"orderRemarks"`  // 订单备注
}

// TableName get sql table name.获取数据库表名
func (m *MacOrder) TableName() string {
	return "mac_order"
}

// MacOrderColumns get sql column name.获取数据库列名
var MacOrderColumns = struct {
	OrderID      string
	UserID       string
	OrderStatus  string
	OrderCode    string
	OrderPrice   string
	OrderTime    string
	OrderPoints  string
	OrderPayType string
	OrderPayTime string
	OrderRemarks string
}{
	OrderID:      "order_id",
	UserID:       "user_id",
	OrderStatus:  "order_status",
	OrderCode:    "order_code",
	OrderPrice:   "order_price",
	OrderTime:    "order_time",
	OrderPoints:  "order_points",
	OrderPayType: "order_pay_type",
	OrderPayTime: "order_pay_time",
	OrderRemarks: "order_remarks",
}

// MacPlog [...]
type MacPlog struct {
	PlogID      int    `gorm:"primaryKey;column:plog_id" json:"-"`     // 积分日志ID
	UserID      int    `gorm:"column:user_id" json:"userId"`           // 用户ID
	UserID1     int    `gorm:"column:user_id_1" json:"userId1"`        // 关联用户ID
	PlogType    uint8  `gorm:"column:plog_type" json:"plogType"`       // 日志类型
	PlogPoints  uint16 `gorm:"column:plog_points" json:"plogPoints"`   // 积分数量
	PlogTime    int    `gorm:"column:plog_time" json:"plogTime"`       // 日志时间
	PlogRemarks string `gorm:"column:plog_remarks" json:"plogRemarks"` // 日志备注
}

// TableName get sql table name.获取数据库表名
func (m *MacPlog) TableName() string {
	return "mac_plog"
}

// MacPlogColumns get sql column name.获取数据库列名
var MacPlogColumns = struct {
	PlogID      string
	UserID      string
	UserID1     string
	PlogType    string
	PlogPoints  string
	PlogTime    string
	PlogRemarks string
}{
	PlogID:      "plog_id",
	UserID:      "user_id",
	UserID1:     "user_id_1",
	PlogType:    "plog_type",
	PlogPoints:  "plog_points",
	PlogTime:    "plog_time",
	PlogRemarks: "plog_remarks",
}

// MacRole [...]
type MacRole struct {
	RoleID        int     `gorm:"primaryKey;column:role_id" json:"-"`          // 角色ID
	RoleRid       int     `gorm:"column:role_rid" json:"roleRid"`              // 关联角色ID
	RoleName      string  `gorm:"column:role_name" json:"roleName"`            // 角色名称
	RoleEn        string  `gorm:"column:role_en" json:"roleEn"`                // 角色英文名
	RoleStatus    uint8   `gorm:"column:role_status" json:"roleStatus"`        // 角色状态
	RoleLock      uint8   `gorm:"column:role_lock" json:"roleLock"`            // 角色锁定
	RoleLetter    string  `gorm:"column:role_letter" json:"roleLetter"`        // 角色首字母
	RoleColor     string  `gorm:"column:role_color" json:"roleColor"`          // 角色颜色
	RoleActor     string  `gorm:"column:role_actor" json:"roleActor"`          // 关联演员
	RoleRemarks   string  `gorm:"column:role_remarks" json:"roleRemarks"`      // 角色备注
	RolePic       string  `gorm:"column:role_pic" json:"rolePic"`              // 角色图片
	RoleSort      uint16  `gorm:"column:role_sort" json:"roleSort"`            // 角色排序
	RoleLevel     uint8   `gorm:"column:role_level" json:"roleLevel"`          // 角色等级
	RoleTime      int     `gorm:"column:role_time" json:"roleTime"`            // 角色时间
	RoleTimeAdd   int     `gorm:"column:role_time_add" json:"roleTimeAdd"`     // 角色添加时间
	RoleTimeHits  int     `gorm:"column:role_time_hits" json:"roleTimeHits"`   // 角色点击时间
	RoleTimeMake  int     `gorm:"column:role_time_make" json:"roleTimeMake"`   // 角色制作时间
	RoleHits      int16   `gorm:"column:role_hits" json:"roleHits"`            // 角色点击量
	RoleHitsDay   int16   `gorm:"column:role_hits_day" json:"roleHitsDay"`     // 角色日点击量
	RoleHitsWeek  int16   `gorm:"column:role_hits_week" json:"roleHitsWeek"`   // 角色周点击量
	RoleHitsMonth int16   `gorm:"column:role_hits_month" json:"roleHitsMonth"` // 角色月点击量
	RoleScore     float64 `gorm:"column:role_score" json:"roleScore"`          // 角色评分
	RoleScoreAll  int16   `gorm:"column:role_score_all" json:"roleScoreAll"`   // 角色总评分
	RoleScoreNum  int16   `gorm:"column:role_score_num" json:"roleScoreNum"`   // 角色评分人数
	RoleUp        int16   `gorm:"column:role_up" json:"roleUp"`                // 角色点赞数
	RoleDown      int16   `gorm:"column:role_down" json:"roleDown"`            // 角色踩数
	RoleTpl       string  `gorm:"column:role_tpl" json:"roleTpl"`              // 角色模板
	RoleJumpurl   string  `gorm:"column:role_jumpurl" json:"roleJumpurl"`      // 角色跳转URL
	RoleContent   string  `gorm:"column:role_content" json:"roleContent"`      // 角色内容
}

// TableName get sql table name.获取数据库表名
func (m *MacRole) TableName() string {
	return "mac_role"
}

// MacRoleColumns get sql column name.获取数据库列名
var MacRoleColumns = struct {
	RoleID        string
	RoleRid       string
	RoleName      string
	RoleEn        string
	RoleStatus    string
	RoleLock      string
	RoleLetter    string
	RoleColor     string
	RoleActor     string
	RoleRemarks   string
	RolePic       string
	RoleSort      string
	RoleLevel     string
	RoleTime      string
	RoleTimeAdd   string
	RoleTimeHits  string
	RoleTimeMake  string
	RoleHits      string
	RoleHitsDay   string
	RoleHitsWeek  string
	RoleHitsMonth string
	RoleScore     string
	RoleScoreAll  string
	RoleScoreNum  string
	RoleUp        string
	RoleDown      string
	RoleTpl       string
	RoleJumpurl   string
	RoleContent   string
}{
	RoleID:        "role_id",
	RoleRid:       "role_rid",
	RoleName:      "role_name",
	RoleEn:        "role_en",
	RoleStatus:    "role_status",
	RoleLock:      "role_lock",
	RoleLetter:    "role_letter",
	RoleColor:     "role_color",
	RoleActor:     "role_actor",
	RoleRemarks:   "role_remarks",
	RolePic:       "role_pic",
	RoleSort:      "role_sort",
	RoleLevel:     "role_level",
	RoleTime:      "role_time",
	RoleTimeAdd:   "role_time_add",
	RoleTimeHits:  "role_time_hits",
	RoleTimeMake:  "role_time_make",
	RoleHits:      "role_hits",
	RoleHitsDay:   "role_hits_day",
	RoleHitsWeek:  "role_hits_week",
	RoleHitsMonth: "role_hits_month",
	RoleScore:     "role_score",
	RoleScoreAll:  "role_score_all",
	RoleScoreNum:  "role_score_num",
	RoleUp:        "role_up",
	RoleDown:      "role_down",
	RoleTpl:       "role_tpl",
	RoleJumpurl:   "role_jumpurl",
	RoleContent:   "role_content",
}

// MacTmpvod [...]
type MacTmpvod struct {
	ID1   int    `gorm:"column:id1" json:"id1"`     // 临时视频ID
	Name1 string `gorm:"column:name1" json:"name1"` // 视频名称
}

// TableName get sql table name.获取数据库表名
func (m *MacTmpvod) TableName() string {
	return "mac_tmpvod"
}

// MacTmpvodColumns get sql column name.获取数据库列名
var MacTmpvodColumns = struct {
	ID1   string
	Name1 string
}{
	ID1:   "id1",
	Name1: "name1",
}

// MacTopic [...]
type MacTopic struct {
	TopicID        uint16  `gorm:"primaryKey;column:topic_id" json:"-"`           // 专题ID
	TopicName      string  `gorm:"column:topic_name" json:"topicName"`            // 专题名称
	TopicEn        string  `gorm:"column:topic_en" json:"topicEn"`                // 专题英文名
	TopicSub       string  `gorm:"column:topic_sub" json:"topicSub"`              // 专题副标题
	TopicStatus    uint8   `gorm:"column:topic_status" json:"topicStatus"`        // 专题状态
	TopicSort      uint16  `gorm:"column:topic_sort" json:"topicSort"`            // 专题排序
	TopicLetter    string  `gorm:"column:topic_letter" json:"topicLetter"`        // 专题首字母
	TopicColor     string  `gorm:"column:topic_color" json:"topicColor"`          // 专题颜色
	TopicTpl       string  `gorm:"column:topic_tpl" json:"topicTpl"`              // 专题模板
	TopicType      string  `gorm:"column:topic_type" json:"topicType"`            // 专题类型
	TopicPic       string  `gorm:"column:topic_pic" json:"topicPic"`              // 专题图片
	TopicPicThumb  string  `gorm:"column:topic_pic_thumb" json:"topicPicThumb"`   // 专题缩略图
	TopicPicSlide  string  `gorm:"column:topic_pic_slide" json:"topicPicSlide"`   // 专题幻灯片
	TopicKey       string  `gorm:"column:topic_key" json:"topicKey"`              // 专题关键字
	TopicDes       string  `gorm:"column:topic_des" json:"topicDes"`              // 专题描述
	TopicTitle     string  `gorm:"column:topic_title" json:"topicTitle"`          // 专题标题
	TopicBlurb     string  `gorm:"column:topic_blurb" json:"topicBlurb"`          // 专题简介
	TopicRemarks   string  `gorm:"column:topic_remarks" json:"topicRemarks"`      // 专题备注
	TopicLevel     uint8   `gorm:"column:topic_level" json:"topicLevel"`          // 专题级别
	TopicUp        int16   `gorm:"column:topic_up" json:"topicUp"`                // 专题点赞数
	TopicDown      int16   `gorm:"column:topic_down" json:"topicDown"`            // 专题踩数
	TopicScore     float64 `gorm:"column:topic_score" json:"topicScore"`          // 专题评分
	TopicScoreAll  int16   `gorm:"column:topic_score_all" json:"topicScoreAll"`   // 专题总评分
	TopicScoreNum  int16   `gorm:"column:topic_score_num" json:"topicScoreNum"`   // 专题评分次数
	TopicHits      int16   `gorm:"column:topic_hits" json:"topicHits"`            // 专题点击数
	TopicHitsDay   int16   `gorm:"column:topic_hits_day" json:"topicHitsDay"`     // 专题日点击数
	TopicHitsWeek  int16   `gorm:"column:topic_hits_week" json:"topicHitsWeek"`   // 专题周点击数
	TopicHitsMonth int16   `gorm:"column:topic_hits_month" json:"topicHitsMonth"` // 专题月点击数
	TopicTime      int     `gorm:"column:topic_time" json:"topicTime"`            // 专题时间
	TopicTimeAdd   int     `gorm:"column:topic_time_add" json:"topicTimeAdd"`     // 专题添加时间
	TopicTimeHits  int     `gorm:"column:topic_time_hits" json:"topicTimeHits"`   // 专题点击时间
	TopicTimeMake  int     `gorm:"column:topic_time_make" json:"topicTimeMake"`   // 专题制作时间
	TopicTag       string  `gorm:"column:topic_tag" json:"topicTag"`              // 专题标签
	TopicRelVod    string  `gorm:"column:topic_rel_vod" json:"topicRelVod"`       // 专题关联视频
	TopicRelArt    string  `gorm:"column:topic_rel_art" json:"topicRelArt"`       // 专题关联文章
	TopicContent   string  `gorm:"column:topic_content" json:"topicContent"`      // 专题内容
	TopicExtend    string  `gorm:"column:topic_extend" json:"topicExtend"`        // 专题扩展信息
}

// TableName get sql table name.获取数据库表名
func (m *MacTopic) TableName() string {
	return "mac_topic"
}

// MacTopicColumns get sql column name.获取数据库列名
var MacTopicColumns = struct {
	TopicID        string
	TopicName      string
	TopicEn        string
	TopicSub       string
	TopicStatus    string
	TopicSort      string
	TopicLetter    string
	TopicColor     string
	TopicTpl       string
	TopicType      string
	TopicPic       string
	TopicPicThumb  string
	TopicPicSlide  string
	TopicKey       string
	TopicDes       string
	TopicTitle     string
	TopicBlurb     string
	TopicRemarks   string
	TopicLevel     string
	TopicUp        string
	TopicDown      string
	TopicScore     string
	TopicScoreAll  string
	TopicScoreNum  string
	TopicHits      string
	TopicHitsDay   string
	TopicHitsWeek  string
	TopicHitsMonth string
	TopicTime      string
	TopicTimeAdd   string
	TopicTimeHits  string
	TopicTimeMake  string
	TopicTag       string
	TopicRelVod    string
	TopicRelArt    string
	TopicContent   string
	TopicExtend    string
}{
	TopicID:        "topic_id",
	TopicName:      "topic_name",
	TopicEn:        "topic_en",
	TopicSub:       "topic_sub",
	TopicStatus:    "topic_status",
	TopicSort:      "topic_sort",
	TopicLetter:    "topic_letter",
	TopicColor:     "topic_color",
	TopicTpl:       "topic_tpl",
	TopicType:      "topic_type",
	TopicPic:       "topic_pic",
	TopicPicThumb:  "topic_pic_thumb",
	TopicPicSlide:  "topic_pic_slide",
	TopicKey:       "topic_key",
	TopicDes:       "topic_des",
	TopicTitle:     "topic_title",
	TopicBlurb:     "topic_blurb",
	TopicRemarks:   "topic_remarks",
	TopicLevel:     "topic_level",
	TopicUp:        "topic_up",
	TopicDown:      "topic_down",
	TopicScore:     "topic_score",
	TopicScoreAll:  "topic_score_all",
	TopicScoreNum:  "topic_score_num",
	TopicHits:      "topic_hits",
	TopicHitsDay:   "topic_hits_day",
	TopicHitsWeek:  "topic_hits_week",
	TopicHitsMonth: "topic_hits_month",
	TopicTime:      "topic_time",
	TopicTimeAdd:   "topic_time_add",
	TopicTimeHits:  "topic_time_hits",
	TopicTimeMake:  "topic_time_make",
	TopicTag:       "topic_tag",
	TopicRelVod:    "topic_rel_vod",
	TopicRelArt:    "topic_rel_art",
	TopicContent:   "topic_content",
	TopicExtend:    "topic_extend",
}

// MacType [...]
type MacType struct {
	TypeID        uint16 `gorm:"primaryKey;column:type_id" json:"-"`          // 分类ID
	TypeName      string `gorm:"column:type_name" json:"typeName"`            // 分类名称
	TypeEn        string `gorm:"column:type_en" json:"typeEn"`                // 分类英文名
	TypeSort      uint16 `gorm:"column:type_sort" json:"typeSort"`            // 分类排序
	TypeMid       uint16 `gorm:"column:type_mid" json:"typeMid"`              // 分类所属模型
	TypePid       uint16 `gorm:"column:type_pid" json:"typePid"`              // 分类父ID
	TypeStatus    uint8  `gorm:"column:type_status" json:"typeStatus"`        // 分类状态
	TypeTpl       string `gorm:"column:type_tpl" json:"typeTpl"`              // 分类模板
	TypeTplList   string `gorm:"column:type_tpl_list" json:"typeTplList"`     // 分类列表模板
	TypeTplDetail string `gorm:"column:type_tpl_detail" json:"typeTplDetail"` // 分类详情模板
	TypeTplPlay   string `gorm:"column:type_tpl_play" json:"typeTplPlay"`     // 分类播放页模板
	TypeTplDown   string `gorm:"column:type_tpl_down" json:"typeTplDown"`     // 分类下载页模板
	TypeKey       string `gorm:"column:type_key" json:"typeKey"`              // 分类关键字
	TypeDes       string `gorm:"column:type_des" json:"typeDes"`              // 分类描述
	TypeTitle     string `gorm:"column:type_title" json:"typeTitle"`          // 分类标题
	TypeUnion     string `gorm:"column:type_union" json:"typeUnion"`          // 分类关联信息
	TypeExtend    string `gorm:"column:type_extend" json:"typeExtend"`        // 分类扩展信息
	TypeLogo      string `gorm:"column:type_logo" json:"typeLogo"`            // 分类Logo
	TypePic       string `gorm:"column:type_pic" json:"typePic"`              // 分类图片
	TypeJumpurl   string `gorm:"column:type_jumpurl" json:"typeJumpurl"`      // 分类跳转URL
}

// TableName get sql table name.获取数据库表名
func (m *MacType) TableName() string {
	return "mac_type"
}

// MacTypeColumns get sql column name.获取数据库列名
var MacTypeColumns = struct {
	TypeID        string
	TypeName      string
	TypeEn        string
	TypeSort      string
	TypeMid       string
	TypePid       string
	TypeStatus    string
	TypeTpl       string
	TypeTplList   string
	TypeTplDetail string
	TypeTplPlay   string
	TypeTplDown   string
	TypeKey       string
	TypeDes       string
	TypeTitle     string
	TypeUnion     string
	TypeExtend    string
	TypeLogo      string
	TypePic       string
	TypeJumpurl   string
}{
	TypeID:        "type_id",
	TypeName:      "type_name",
	TypeEn:        "type_en",
	TypeSort:      "type_sort",
	TypeMid:       "type_mid",
	TypePid:       "type_pid",
	TypeStatus:    "type_status",
	TypeTpl:       "type_tpl",
	TypeTplList:   "type_tpl_list",
	TypeTplDetail: "type_tpl_detail",
	TypeTplPlay:   "type_tpl_play",
	TypeTplDown:   "type_tpl_down",
	TypeKey:       "type_key",
	TypeDes:       "type_des",
	TypeTitle:     "type_title",
	TypeUnion:     "type_union",
	TypeExtend:    "type_extend",
	TypeLogo:      "type_logo",
	TypePic:       "type_pic",
	TypeJumpurl:   "type_jumpurl",
}

// MacUlog [...]
type MacUlog struct {
	UlogID     int    `gorm:"primaryKey;column:ulog_id" json:"-"`   // 播放记录ID
	UserID     int    `gorm:"column:user_id" json:"userId"`         // 用户ID
	UlogMid    uint8  `gorm:"column:ulog_mid" json:"ulogMid"`       // 模型ID
	UlogType   uint8  `gorm:"column:ulog_type" json:"ulogType"`     // 记录类型
	UlogRid    int    `gorm:"column:ulog_rid" json:"ulogRid"`       // 资源ID
	UlogSid    uint8  `gorm:"column:ulog_sid" json:"ulogSid"`       // 影片集数ID
	UlogNid    uint16 `gorm:"column:ulog_nid" json:"ulogNid"`       // 影片ID
	UlogPoints uint16 `gorm:"column:ulog_points" json:"ulogPoints"` // 观看进度
	UlogTime   int    `gorm:"column:ulog_time" json:"ulogTime"`     // 记录时间
}

// TableName get sql table name.获取数据库表名
func (m *MacUlog) TableName() string {
	return "mac_ulog"
}

// MacUlogColumns get sql column name.获取数据库列名
var MacUlogColumns = struct {
	UlogID     string
	UserID     string
	UlogMid    string
	UlogType   string
	UlogRid    string
	UlogSid    string
	UlogNid    string
	UlogPoints string
	UlogTime   string
}{
	UlogID:     "ulog_id",
	UserID:     "user_id",
	UlogMid:    "ulog_mid",
	UlogType:   "ulog_type",
	UlogRid:    "ulog_rid",
	UlogSid:    "ulog_sid",
	UlogNid:    "ulog_nid",
	UlogPoints: "ulog_points",
	UlogTime:   "ulog_time",
}

// MacUser [...]
type MacUser struct {
	UserID            int    `gorm:"primaryKey;column:user_id" json:"-"`                   // 用户ID
	GroupID           uint16 `gorm:"column:group_id" json:"groupId"`                       // 用户组ID
	UserName          string `gorm:"column:user_name" json:"userName"`                     // 用户名
	UserPwd           string `gorm:"column:user_pwd" json:"userPwd"`                       // 用户密码
	UserNickName      string `gorm:"column:user_nick_name" json:"userNickName"`            // 用户昵称
	UserQq            string `gorm:"column:user_qq" json:"userQq"`                         // 用户QQ
	UserEmail         string `gorm:"column:user_email" json:"userEmail"`                   // 用户邮箱
	UserPhone         string `gorm:"column:user_phone" json:"userPhone"`                   // 用户手机号
	UserStatus        uint8  `gorm:"column:user_status" json:"userStatus"`                 // 用户状态
	UserPortrait      string `gorm:"column:user_portrait" json:"userPortrait"`             // 用户头像
	UserPortraitThumb string `gorm:"column:user_portrait_thumb" json:"userPortraitThumb"`  // 用户头像缩略图
	UserOpenidQq      string `gorm:"column:user_openid_qq" json:"userOpenidQq"`            // 用户QQ登录OpenID
	UserOpenidWeixin  string `gorm:"column:user_openid_weixin" json:"userOpenidWeixin"`    // 用户微信登录OpenID
	UserQuestion      string `gorm:"column:user_question" json:"userQuestion"`             // 找回密码问题
	UserAnswer        string `gorm:"column:user_answer" json:"userAnswer"`                 // 找回密码答案
	UserPoints        int    `gorm:"column:user_points" json:"userPoints"`                 // 用户积分
	UserPointsFroze   int    `gorm:"column:user_points_froze" json:"userPointsFroze"`      // 冻结积分
	UserRegTime       int    `gorm:"column:user_reg_time" json:"userRegTime"`              // 注册时间
	UserRegIP         int    `gorm:"column:user_reg_ip" json:"userRegIp"`                  // 注册IP
	UserLoginTime     int    `gorm:"column:user_login_time" json:"userLoginTime"`          // 登录时间
	UserLoginIP       int    `gorm:"column:user_login_ip" json:"userLoginIp"`              // 登录IP
	UserLastLoginTime int    `gorm:"column:user_last_login_time" json:"userLastLoginTime"` // 上次登录时间
	UserLastLoginIP   int    `gorm:"column:user_last_login_ip" json:"userLastLoginIp"`     // 上次登录IP
	UserLoginNum      uint16 `gorm:"column:user_login_num" json:"userLoginNum"`            // 登录次数
	UserExtend        uint16 `gorm:"column:user_extend" json:"userExtend"`                 // 用户扩展字段
	UserRandom        string `gorm:"column:user_random" json:"userRandom"`                 // 用户随机码
	UserEndTime       int    `gorm:"column:user_end_time" json:"userEndTime"`              // VIP到期时间
	UserPid           int    `gorm:"column:user_pid" json:"userPid"`                       // 上级推广员ID
	UserPid2          int    `gorm:"column:user_pid_2" json:"userPid2"`                    // 上级推广员ID_2
	UserPid3          int    `gorm:"column:user_pid_3" json:"userPid3"`                    // 上级推广员ID_3
}

// TableName get sql table name.获取数据库表名
func (m *MacUser) TableName() string {
	return "mac_user"
}

// MacUserColumns get sql column name.获取数据库列名
var MacUserColumns = struct {
	UserID            string
	GroupID           string
	UserName          string
	UserPwd           string
	UserNickName      string
	UserQq            string
	UserEmail         string
	UserPhone         string
	UserStatus        string
	UserPortrait      string
	UserPortraitThumb string
	UserOpenidQq      string
	UserOpenidWeixin  string
	UserQuestion      string
	UserAnswer        string
	UserPoints        string
	UserPointsFroze   string
	UserRegTime       string
	UserRegIP         string
	UserLoginTime     string
	UserLoginIP       string
	UserLastLoginTime string
	UserLastLoginIP   string
	UserLoginNum      string
	UserExtend        string
	UserRandom        string
	UserEndTime       string
	UserPid           string
	UserPid2          string
	UserPid3          string
}{
	UserID:            "user_id",
	GroupID:           "group_id",
	UserName:          "user_name",
	UserPwd:           "user_pwd",
	UserNickName:      "user_nick_name",
	UserQq:            "user_qq",
	UserEmail:         "user_email",
	UserPhone:         "user_phone",
	UserStatus:        "user_status",
	UserPortrait:      "user_portrait",
	UserPortraitThumb: "user_portrait_thumb",
	UserOpenidQq:      "user_openid_qq",
	UserOpenidWeixin:  "user_openid_weixin",
	UserQuestion:      "user_question",
	UserAnswer:        "user_answer",
	UserPoints:        "user_points",
	UserPointsFroze:   "user_points_froze",
	UserRegTime:       "user_reg_time",
	UserRegIP:         "user_reg_ip",
	UserLoginTime:     "user_login_time",
	UserLoginIP:       "user_login_ip",
	UserLastLoginTime: "user_last_login_time",
	UserLastLoginIP:   "user_last_login_ip",
	UserLoginNum:      "user_login_num",
	UserExtend:        "user_extend",
	UserRandom:        "user_random",
	UserEndTime:       "user_end_time",
	UserPid:           "user_pid",
	UserPid2:          "user_pid_2",
	UserPid3:          "user_pid_3",
}

// MacVisit [...]
type MacVisit struct {
	VisitID   int    `gorm:"primaryKey;column:visit_id" json:"-"` // 访问记录ID
	UserID    int    `gorm:"column:user_id" json:"userId"`        // 用户ID
	VisitIP   int    `gorm:"column:visit_ip" json:"visitIp"`      // 访问IP
	VisitLy   string `gorm:"column:visit_ly" json:"visitLy"`      // 访问来源
	VisitTime int    `gorm:"column:visit_time" json:"visitTime"`  // 访问时间
}

// TableName get sql table name.获取数据库表名
func (m *MacVisit) TableName() string {
	return "mac_visit"
}

// MacVisitColumns get sql column name.获取数据库列名
var MacVisitColumns = struct {
	VisitID   string
	UserID    string
	VisitIP   string
	VisitLy   string
	VisitTime string
}{
	VisitID:   "visit_id",
	UserID:    "user_id",
	VisitIP:   "visit_ip",
	VisitLy:   "visit_ly",
	VisitTime: "visit_time",
}

// MacVod [...]
type MacVod struct {
	VodID            int     `gorm:"primaryKey;column:vod_id" json:"vodID"`             // 视频ID
	TypeID           int16   `gorm:"column:type_id" json:"typeId"`                      // 栏目ID
	TypeID1          uint16  `gorm:"column:type_id_1" json:"typeId1"`                   // 子栏目ID
	GroupID          uint16  `gorm:"column:group_id" json:"groupId"`                    // 分组ID
	VodName          string  `gorm:"column:vod_name" json:"vodName"`                    // 视频标题
	VodSub           string  `gorm:"column:vod_sub" json:"vodSub"`                      // 视频子标题
	VodEn            string  `gorm:"column:vod_en" json:"vodEn"`                        // 视频英文或者拼音名称
	VodStatus        uint8   `gorm:"column:vod_status" json:"vodStatus"`                // 视频状态
	VodLetter        string  `gorm:"column:vod_letter" json:"vodLetter"`                // 首字母索引排序
	VodColor         string  `gorm:"column:vod_color" json:"vodColor"`                  // 颜色值
	VodTag           string  `gorm:"column:vod_tag" json:"vodTag"`                      // 视频标签
	VodClass         string  `gorm:"column:vod_class" json:"vodClass"`                  // 视频分类
	VodPic           string  `gorm:"column:vod_pic" json:"vodPic"`                      // 视频封面图
	VodPicThumb      string  `gorm:"column:vod_pic_thumb" json:"vodPicThumb"`           // 视频封面图的缩略图
	VodPicSlide      string  `gorm:"column:vod_pic_slide" json:"vodPicSlide"`           // 视频封面图的幻灯片
	VodPicScreenshot string  `gorm:"column:vod_pic_screenshot" json:"vodPicScreenshot"` // 视频截图
	VodActor         string  `gorm:"column:vod_actor" json:"vodActor"`                  // 视频演员
	VodDirector      string  `gorm:"column:vod_director" json:"vodDirector"`            // 视频导演
	VodWriter        string  `gorm:"column:vod_writer" json:"vodWriter"`                // 视频编剧
	VodBehind        string  `gorm:"column:vod_behind" json:"vodBehind"`                // 幕后制作人员
	VodBlurb         string  `gorm:"column:vod_blurb" json:"vodBlurb"`                  // 视频简介
	VodRemarks       string  `gorm:"column:vod_remarks" json:"vodRemarks"`              // 备注
	VodPubdate       string  `gorm:"column:vod_pubdate" json:"vodPubdate"`              // 视频发布日期
	VodTotal         int16   `gorm:"column:vod_total" json:"vodTotal"`                  // 视频总数
	VodSerial        string  `gorm:"column:vod_serial" json:"vodSerial"`                // 视频连载信息
	VodTv            string  `gorm:"column:vod_tv" json:"vodTv"`                        // 视频播出的电视台
	VodWeekday       string  `gorm:"column:vod_weekday" json:"vodWeekday"`              // 视频播出的星期几
	VodArea          string  `gorm:"column:vod_area" json:"vodArea"`                    // 视频制作地区
	VodLang          string  `gorm:"column:vod_lang" json:"vodLang"`                    // 视频语言
	VodYear          string  `gorm:"column:vod_year" json:"vodYear"`                    // 视频制作年份
	VodVersion       string  `gorm:"column:vod_version" json:"vodVersion"`              // 视频版本
	VodState         string  `gorm:"column:vod_state" json:"vodState"`                  // 视频状态
	VodAuthor        string  `gorm:"column:vod_author" json:"vodAuthor"`                // 视频作者
	VodJumpurl       string  `gorm:"column:vod_jumpurl" json:"vodJumpurl"`              // 跳转链接
	VodTpl           string  `gorm:"column:vod_tpl" json:"vodTpl"`                      // 模板信息
	VodTplPlay       string  `gorm:"column:vod_tpl_play" json:"vodTplPlay"`             // 播放模板信息
	VodTplDown       string  `gorm:"column:vod_tpl_down" json:"vodTplDown"`             // 下载模板信息
	VodIsend         uint8   `gorm:"column:vod_isend" json:"vodIsend"`                  // 视频是否完结
	VodLock          uint8   `gorm:"column:vod_lock" json:"vodLock"`                    // 视频是否锁定
	VodLevel         uint8   `gorm:"column:vod_level" json:"vodLevel"`                  // 视频级别
	VodCopyright     uint8   `gorm:"column:vod_copyright" json:"vodCopyright"`          // 视频版权信息
	VodPoints        uint16  `gorm:"column:vod_points" json:"vodPoints"`                // 视频积分
	VodPointsPlay    uint16  `gorm:"column:vod_points_play" json:"vodPointsPlay"`       // 播放积分
	VodPointsDown    uint16  `gorm:"column:vod_points_down" json:"vodPointsDown"`       // 下载积分
	VodHits          int16   `gorm:"column:vod_hits" json:"vodHits"`                    // 视频点击数
	VodHitsDay       int16   `gorm:"column:vod_hits_day" json:"vodHitsDay"`             // 视频每日点击数
	VodHitsWeek      int16   `gorm:"column:vod_hits_week" json:"vodHitsWeek"`           // 视频每周点击数
	VodHitsMonth     int16   `gorm:"column:vod_hits_month" json:"vodHitsMonth"`         // 视频每月点击数
	VodDuration      string  `gorm:"column:vod_duration" json:"vodDuration"`            // 视频时长
	VodUp            int16   `gorm:"column:vod_up" json:"vodUp"`                        // 视频点赞数
	VodDown          int16   `gorm:"column:vod_down" json:"vodDown"`                    // 视频踩数
	VodScore         float64 `gorm:"column:vod_score" json:"vodScore"`                  // 视频评分
	VodScoreAll      int16   `gorm:"column:vod_score_all" json:"vodScoreAll"`           // 视频总评分数
	VodScoreNum      int16   `gorm:"column:vod_score_num" json:"vodScoreNum"`           // 视频评分人数
	VodTime          int     `gorm:"column:vod_time" json:"vodTime"`                    // 视频总时长
	VodTimeAdd       int     `gorm:"column:vod_time_add" json:"vodTimeAdd"`             // 视频添加时长
	VodTimeHits      int     `gorm:"column:vod_time_hits" json:"vodTimeHits"`           // 视频点击时长
	VodTimeMake      int     `gorm:"column:vod_time_make" json:"vodTimeMake"`           // 视频制作时长
	VodTrysee        uint16  `gorm:"column:vod_trysee" json:"vodTrysee"`                // 试看时长
	VodDoubanID      int     `gorm:"column:vod_douban_id" json:"vodDoubanId"`           // 豆瓣ID
	VodDoubanScore   float64 `gorm:"column:vod_douban_score" json:"vodDoubanScore"`     // 豆瓣评分
	VodReurl         string  `gorm:"column:vod_reurl" json:"vodReurl"`                  // 重定向链接
	VodRelVod        string  `gorm:"column:vod_rel_vod" json:"vodRelVod"`               // 相关视频链接
	VodRelArt        string  `gorm:"column:vod_rel_art" json:"vodRelArt"`               // 相关文章链接
	VodPwd           string  `gorm:"column:vod_pwd" json:"vodPwd"`                      // 视频密码
	VodPwdURL        string  `gorm:"column:vod_pwd_url" json:"vodPwdUrl"`               // 密码链接
	VodPwdPlay       string  `gorm:"column:vod_pwd_play" json:"vodPwdPlay"`             // 播放密码
	VodPwdPlayURL    string  `gorm:"column:vod_pwd_play_url" json:"vodPwdPlayUrl"`      // 播放密码链接
	VodPwdDown       string  `gorm:"column:vod_pwd_down" json:"vodPwdDown"`             // 下载密码
	VodPwdDownURL    string  `gorm:"column:vod_pwd_down_url" json:"vodPwdDownUrl"`      // 下载密码链接
	VodContent       string  `gorm:"column:vod_content" json:"vodContent"`              // 视频内容
	VodPlayFrom      string  `gorm:"column:vod_play_from" json:"vodPlayFrom"`           // 播放来源
	VodPlayServer    string  `gorm:"column:vod_play_server" json:"vodPlayServer"`       // 播放服务器
	VodPlayNote      string  `gorm:"column:vod_play_note" json:"vodPlayNote"`           // 播放注意事项
	VodPlayURL       string  `gorm:"column:vod_play_url" json:"vodPlayUrl"`             // 播放链接
	VodDownFrom      string  `gorm:"column:vod_down_from" json:"vodDownFrom"`           // 下载来源
	VodDownServer    string  `gorm:"column:vod_down_server" json:"vodDownServer"`       // 下载服务器
	VodDownNote      string  `gorm:"column:vod_down_note" json:"vodDownNote"`           // 下载注意事项
	VodDownURL       string  `gorm:"column:vod_down_url" json:"vodDownUrl"`             // 下载链接
	VodPlot          uint8   `gorm:"column:vod_plot" json:"vodPlot"`                    // 剧情类型
	VodPlotName      string  `gorm:"column:vod_plot_name" json:"vodPlotName"`           // 剧情名称
	VodPlotDetail    string  `gorm:"column:vod_plot_detail" json:"vodPlotDetail"`       // 剧情详情
}

// TableName get sql table name.获取数据库表名
func (m *MacVod) TableName() string {
	return "mac_vod"
}

// MacVodColumns get sql column name.获取数据库列名
var MacVodColumns = struct {
	VodID            string
	TypeID           string
	TypeID1          string
	GroupID          string
	VodName          string
	VodSub           string
	VodEn            string
	VodStatus        string
	VodLetter        string
	VodColor         string
	VodTag           string
	VodClass         string
	VodPic           string
	VodPicThumb      string
	VodPicSlide      string
	VodPicScreenshot string
	VodActor         string
	VodDirector      string
	VodWriter        string
	VodBehind        string
	VodBlurb         string
	VodRemarks       string
	VodPubdate       string
	VodTotal         string
	VodSerial        string
	VodTv            string
	VodWeekday       string
	VodArea          string
	VodLang          string
	VodYear          string
	VodVersion       string
	VodState         string
	VodAuthor        string
	VodJumpurl       string
	VodTpl           string
	VodTplPlay       string
	VodTplDown       string
	VodIsend         string
	VodLock          string
	VodLevel         string
	VodCopyright     string
	VodPoints        string
	VodPointsPlay    string
	VodPointsDown    string
	VodHits          string
	VodHitsDay       string
	VodHitsWeek      string
	VodHitsMonth     string
	VodDuration      string
	VodUp            string
	VodDown          string
	VodScore         string
	VodScoreAll      string
	VodScoreNum      string
	VodTime          string
	VodTimeAdd       string
	VodTimeHits      string
	VodTimeMake      string
	VodTrysee        string
	VodDoubanID      string
	VodDoubanScore   string
	VodReurl         string
	VodRelVod        string
	VodRelArt        string
	VodPwd           string
	VodPwdURL        string
	VodPwdPlay       string
	VodPwdPlayURL    string
	VodPwdDown       string
	VodPwdDownURL    string
	VodContent       string
	VodPlayFrom      string
	VodPlayServer    string
	VodPlayNote      string
	VodPlayURL       string
	VodDownFrom      string
	VodDownServer    string
	VodDownNote      string
	VodDownURL       string
	VodPlot          string
	VodPlotName      string
	VodPlotDetail    string
}{
	VodID:            "vod_id",
	TypeID:           "type_id",
	TypeID1:          "type_id_1",
	GroupID:          "group_id",
	VodName:          "vod_name",
	VodSub:           "vod_sub",
	VodEn:            "vod_en",
	VodStatus:        "vod_status",
	VodLetter:        "vod_letter",
	VodColor:         "vod_color",
	VodTag:           "vod_tag",
	VodClass:         "vod_class",
	VodPic:           "vod_pic",
	VodPicThumb:      "vod_pic_thumb",
	VodPicSlide:      "vod_pic_slide",
	VodPicScreenshot: "vod_pic_screenshot",
	VodActor:         "vod_actor",
	VodDirector:      "vod_director",
	VodWriter:        "vod_writer",
	VodBehind:        "vod_behind",
	VodBlurb:         "vod_blurb",
	VodRemarks:       "vod_remarks",
	VodPubdate:       "vod_pubdate",
	VodTotal:         "vod_total",
	VodSerial:        "vod_serial",
	VodTv:            "vod_tv",
	VodWeekday:       "vod_weekday",
	VodArea:          "vod_area",
	VodLang:          "vod_lang",
	VodYear:          "vod_year",
	VodVersion:       "vod_version",
	VodState:         "vod_state",
	VodAuthor:        "vod_author",
	VodJumpurl:       "vod_jumpurl",
	VodTpl:           "vod_tpl",
	VodTplPlay:       "vod_tpl_play",
	VodTplDown:       "vod_tpl_down",
	VodIsend:         "vod_isend",
	VodLock:          "vod_lock",
	VodLevel:         "vod_level",
	VodCopyright:     "vod_copyright",
	VodPoints:        "vod_points",
	VodPointsPlay:    "vod_points_play",
	VodPointsDown:    "vod_points_down",
	VodHits:          "vod_hits",
	VodHitsDay:       "vod_hits_day",
	VodHitsWeek:      "vod_hits_week",
	VodHitsMonth:     "vod_hits_month",
	VodDuration:      "vod_duration",
	VodUp:            "vod_up",
	VodDown:          "vod_down",
	VodScore:         "vod_score",
	VodScoreAll:      "vod_score_all",
	VodScoreNum:      "vod_score_num",
	VodTime:          "vod_time",
	VodTimeAdd:       "vod_time_add",
	VodTimeHits:      "vod_time_hits",
	VodTimeMake:      "vod_time_make",
	VodTrysee:        "vod_trysee",
	VodDoubanID:      "vod_douban_id",
	VodDoubanScore:   "vod_douban_score",
	VodReurl:         "vod_reurl",
	VodRelVod:        "vod_rel_vod",
	VodRelArt:        "vod_rel_art",
	VodPwd:           "vod_pwd",
	VodPwdURL:        "vod_pwd_url",
	VodPwdPlay:       "vod_pwd_play",
	VodPwdPlayURL:    "vod_pwd_play_url",
	VodPwdDown:       "vod_pwd_down",
	VodPwdDownURL:    "vod_pwd_down_url",
	VodContent:       "vod_content",
	VodPlayFrom:      "vod_play_from",
	VodPlayServer:    "vod_play_server",
	VodPlayNote:      "vod_play_note",
	VodPlayURL:       "vod_play_url",
	VodDownFrom:      "vod_down_from",
	VodDownServer:    "vod_down_server",
	VodDownNote:      "vod_down_note",
	VodDownURL:       "vod_down_url",
	VodPlot:          "vod_plot",
	VodPlotName:      "vod_plot_name",
	VodPlotDetail:    "vod_plot_detail",
}

// MacVodSearch vod搜索缓存表
type MacVodSearch struct {
	SearchKey         string `gorm:"primaryKey;column:search_key" json:"-"`                // 搜索键（关键词md5）
	SearchWord        string `gorm:"column:search_word" json:"searchWord"`                 // 搜索关键词
	SearchField       string `gorm:"column:search_field" json:"searchField"`               // 搜索字段名（可有多个，用|分隔）
	SearchHitCount    uint64 `gorm:"column:search_hit_count" json:"searchHitCount"`        // 搜索命中次数
	SearchLastHitTime int    `gorm:"column:search_last_hit_time" json:"searchLastHitTime"` // 最近命中时间
	SearchUpdateTime  int    `gorm:"column:search_update_time" json:"searchUpdateTime"`    // 添加时间
	SearchResultCount int    `gorm:"column:search_result_count" json:"searchResultCount"`  // 结果Id数量
	SearchResultIDs   string `gorm:"column:search_result_ids" json:"searchResultIds"`      // 搜索结果Id列表，英文半角逗号分隔
}

// TableName get sql table name.获取数据库表名
func (m *MacVodSearch) TableName() string {
	return "mac_vod_search"
}

// MacVodSearchColumns get sql column name.获取数据库列名
var MacVodSearchColumns = struct {
	SearchKey         string
	SearchWord        string
	SearchField       string
	SearchHitCount    string
	SearchLastHitTime string
	SearchUpdateTime  string
	SearchResultCount string
	SearchResultIDs   string
}{
	SearchKey:         "search_key",
	SearchWord:        "search_word",
	SearchField:       "search_field",
	SearchHitCount:    "search_hit_count",
	SearchLastHitTime: "search_last_hit_time",
	SearchUpdateTime:  "search_update_time",
	SearchResultCount: "search_result_count",
	SearchResultIDs:   "search_result_ids",
}

// MacWebsite [...]
type MacWebsite struct {
	WebsiteID            int     `gorm:"primaryKey;column:website_id" json:"-"`                     // 网站ID
	TypeID               uint16  `gorm:"column:type_id" json:"typeId"`                              // 类型ID
	TypeID1              uint16  `gorm:"column:type_id_1" json:"typeId1"`                           // 类型ID_1
	WebsiteName          string  `gorm:"column:website_name" json:"websiteName"`                    // 网站名称
	WebsiteSub           string  `gorm:"column:website_sub" json:"websiteSub"`                      // 网站子标题
	WebsiteEn            string  `gorm:"column:website_en" json:"websiteEn"`                        // 网站英文名称
	WebsiteStatus        uint8   `gorm:"column:website_status" json:"websiteStatus"`                // 网站状态
	WebsiteLetter        string  `gorm:"column:website_letter" json:"websiteLetter"`                // 网站首字母
	WebsiteColor         string  `gorm:"column:website_color" json:"websiteColor"`                  // 网站颜色
	WebsiteLock          uint8   `gorm:"column:website_lock" json:"websiteLock"`                    // 网站锁定状态
	WebsiteSort          int     `gorm:"column:website_sort" json:"websiteSort"`                    // 网站排序
	WebsiteJumpurl       string  `gorm:"column:website_jumpurl" json:"websiteJumpurl"`              // 网站跳转URL
	WebsitePic           string  `gorm:"column:website_pic" json:"websitePic"`                      // 网站图片
	WebsitePicScreenshot string  `gorm:"column:website_pic_screenshot" json:"websitePicScreenshot"` // 网站图片截图
	WebsiteLogo          string  `gorm:"column:website_logo" json:"websiteLogo"`                    // 网站Logo
	WebsiteArea          string  `gorm:"column:website_area" json:"websiteArea"`                    // 网站地区
	WebsiteLang          string  `gorm:"column:website_lang" json:"websiteLang"`                    // 网站语言
	WebsiteLevel         uint8   `gorm:"column:website_level" json:"websiteLevel"`                  // 网站级别
	WebsiteTime          int     `gorm:"column:website_time" json:"websiteTime"`                    // 网站时间
	WebsiteTimeAdd       int     `gorm:"column:website_time_add" json:"websiteTimeAdd"`             // 网站添加时间
	WebsiteTimeHits      int     `gorm:"column:website_time_hits" json:"websiteTimeHits"`           // 网站点击时间
	WebsiteTimeMake      int     `gorm:"column:website_time_make" json:"websiteTimeMake"`           // 网站制作时间
	WebsiteTimeReferer   int     `gorm:"column:website_time_referer" json:"websiteTimeReferer"`     // 网站引用时间
	WebsiteHits          int16   `gorm:"column:website_hits" json:"websiteHits"`                    // 网站点击次数
	WebsiteHitsDay       int16   `gorm:"column:website_hits_day" json:"websiteHitsDay"`             // 网站日点击次数
	WebsiteHitsWeek      int16   `gorm:"column:website_hits_week" json:"websiteHitsWeek"`           // 网站周点击次数
	WebsiteHitsMonth     int16   `gorm:"column:website_hits_month" json:"websiteHitsMonth"`         // 网站月点击次数
	WebsiteScore         float64 `gorm:"column:website_score" json:"websiteScore"`                  // 网站评分
	WebsiteScoreAll      int16   `gorm:"column:website_score_all" json:"websiteScoreAll"`           // 网站总评分
	WebsiteScoreNum      int16   `gorm:"column:website_score_num" json:"websiteScoreNum"`           // 网站评分次数
	WebsiteUp            int16   `gorm:"column:website_up" json:"websiteUp"`                        // 网站点赞数
	WebsiteDown          int16   `gorm:"column:website_down" json:"websiteDown"`                    // 网站踩数
	WebsiteReferer       int16   `gorm:"column:website_referer" json:"websiteReferer"`              // 网站引用数
	WebsiteRefererDay    int16   `gorm:"column:website_referer_day" json:"websiteRefererDay"`       // 网站日引用数
	WebsiteRefererWeek   int16   `gorm:"column:website_referer_week" json:"websiteRefererWeek"`     // 网站周引用数
	WebsiteRefererMonth  int16   `gorm:"column:website_referer_month" json:"websiteRefererMonth"`   // 网站月引用数
	WebsiteTag           string  `gorm:"column:website_tag" json:"websiteTag"`                      // 网站标签
	WebsiteClass         string  `gorm:"column:website_class" json:"websiteClass"`                  // 网站类别
	WebsiteRemarks       string  `gorm:"column:website_remarks" json:"websiteRemarks"`              // 网站备注
	WebsiteTpl           string  `gorm:"column:website_tpl" json:"websiteTpl"`                      // 网站模板
	WebsiteBlurb         string  `gorm:"column:website_blurb" json:"websiteBlurb"`                  // 网站简介
	WebsiteContent       string  `gorm:"column:website_content" json:"websiteContent"`              // 网站内容
}

// TableName get sql table name.获取数据库表名
func (m *MacWebsite) TableName() string {
	return "mac_website"
}

// MacWebsiteColumns get sql column name.获取数据库列名
var MacWebsiteColumns = struct {
	WebsiteID            string
	TypeID               string
	TypeID1              string
	WebsiteName          string
	WebsiteSub           string
	WebsiteEn            string
	WebsiteStatus        string
	WebsiteLetter        string
	WebsiteColor         string
	WebsiteLock          string
	WebsiteSort          string
	WebsiteJumpurl       string
	WebsitePic           string
	WebsitePicScreenshot string
	WebsiteLogo          string
	WebsiteArea          string
	WebsiteLang          string
	WebsiteLevel         string
	WebsiteTime          string
	WebsiteTimeAdd       string
	WebsiteTimeHits      string
	WebsiteTimeMake      string
	WebsiteTimeReferer   string
	WebsiteHits          string
	WebsiteHitsDay       string
	WebsiteHitsWeek      string
	WebsiteHitsMonth     string
	WebsiteScore         string
	WebsiteScoreAll      string
	WebsiteScoreNum      string
	WebsiteUp            string
	WebsiteDown          string
	WebsiteReferer       string
	WebsiteRefererDay    string
	WebsiteRefererWeek   string
	WebsiteRefererMonth  string
	WebsiteTag           string
	WebsiteClass         string
	WebsiteRemarks       string
	WebsiteTpl           string
	WebsiteBlurb         string
	WebsiteContent       string
}{
	WebsiteID:            "website_id",
	TypeID:               "type_id",
	TypeID1:              "type_id_1",
	WebsiteName:          "website_name",
	WebsiteSub:           "website_sub",
	WebsiteEn:            "website_en",
	WebsiteStatus:        "website_status",
	WebsiteLetter:        "website_letter",
	WebsiteColor:         "website_color",
	WebsiteLock:          "website_lock",
	WebsiteSort:          "website_sort",
	WebsiteJumpurl:       "website_jumpurl",
	WebsitePic:           "website_pic",
	WebsitePicScreenshot: "website_pic_screenshot",
	WebsiteLogo:          "website_logo",
	WebsiteArea:          "website_area",
	WebsiteLang:          "website_lang",
	WebsiteLevel:         "website_level",
	WebsiteTime:          "website_time",
	WebsiteTimeAdd:       "website_time_add",
	WebsiteTimeHits:      "website_time_hits",
	WebsiteTimeMake:      "website_time_make",
	WebsiteTimeReferer:   "website_time_referer",
	WebsiteHits:          "website_hits",
	WebsiteHitsDay:       "website_hits_day",
	WebsiteHitsWeek:      "website_hits_week",
	WebsiteHitsMonth:     "website_hits_month",
	WebsiteScore:         "website_score",
	WebsiteScoreAll:      "website_score_all",
	WebsiteScoreNum:      "website_score_num",
	WebsiteUp:            "website_up",
	WebsiteDown:          "website_down",
	WebsiteReferer:       "website_referer",
	WebsiteRefererDay:    "website_referer_day",
	WebsiteRefererWeek:   "website_referer_week",
	WebsiteRefererMonth:  "website_referer_month",
	WebsiteTag:           "website_tag",
	WebsiteClass:         "website_class",
	WebsiteRemarks:       "website_remarks",
	WebsiteTpl:           "website_tpl",
	WebsiteBlurb:         "website_blurb",
	WebsiteContent:       "website_content",
}

// MacWechat [...]
type MacWechat struct {
	ID          int64  `gorm:"primaryKey;column:id" json:"-"`
	BelongWx    string `gorm:"column:belong_wx" json:"belongWx"`       // 所属账号
	Account     string `gorm:"column:account" json:"account"`          // 微信账号
	Avatar      string `gorm:"column:avatar" json:"avatar"`            // 微信头像
	City        string `gorm:"column:city" json:"city"`                // 微信城市
	Country     string `gorm:"column:country" json:"country"`          // 国家
	LabelidList string `gorm:"column:labelid_list" json:"labelidList"` // 标签列表
	Nickname    string `gorm:"column:nickname" json:"nickname"`        // 昵称
	Province    string `gorm:"column:province" json:"province"`        // 省份
	Remark      string `gorm:"column:remark" json:"remark"`            // 备注
	Sex         int    `gorm:"column:sex" json:"sex"`                  // 性别:1-男，2-女 0-未知
	Wxid        string `gorm:"column:wxid" json:"wxid"`                // 微信ID
	CreateTs    int64  `gorm:"column:create_ts" json:"createTs"`
	UpdateTs    int64  `gorm:"column:update_ts" json:"updateTs"`
	DeleteTs    int64  `gorm:"column:delete_ts" json:"deleteTs"`
}

// TableName get sql table name.获取数据库表名
func (m *MacWechat) TableName() string {
	return "mac_wechat"
}

// MacWechatColumns get sql column name.获取数据库列名
var MacWechatColumns = struct {
	ID          string
	BelongWx    string
	Account     string
	Avatar      string
	City        string
	Country     string
	LabelidList string
	Nickname    string
	Province    string
	Remark      string
	Sex         string
	Wxid        string
	CreateTs    string
	UpdateTs    string
	DeleteTs    string
}{
	ID:          "id",
	BelongWx:    "belong_wx",
	Account:     "account",
	Avatar:      "avatar",
	City:        "city",
	Country:     "country",
	LabelidList: "labelid_list",
	Nickname:    "nickname",
	Province:    "province",
	Remark:      "remark",
	Sex:         "sex",
	Wxid:        "wxid",
	CreateTs:    "create_ts",
	UpdateTs:    "update_ts",
	DeleteTs:    "delete_ts",
}
