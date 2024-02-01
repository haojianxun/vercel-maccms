package requests

type MacAdmin struct {
	AdminID            uint16 `json:"admin_id"`                                         // 管理员ID
	AdminName          string `json:"admin_name" form:"admin_name" validate:"required"` // 管理员姓名
	AdminPwd           string `json:"admin_pwd" form:"admin_pwd" validate:"required"`   // 管理员密码
	AdminRandom        string `json:"admin_random"`                                     // 管理员随机值
	AdminStatus        uint8  `json:"admin_status"`                                     // 管理员状态
	AdminAuth          string `json:"admin_auth"`                                       // 管理员权限
	AdminLoginTime     uint   `json:"admin_login_time"`                                 // 管理员登录时间
	AdminLoginIp       uint   `json:"admin_login_ip"`                                   // 管理员登录IP
	AdminLoginNum      uint   `json:"admin_login_num"`                                  // 管理员登录次数
	AdminLastLoginTime uint   `json:"admin_last_login_time"`                            // 管理员上次登录时间
	AdminLastLoginIp   uint   `json:"admin_last_login_ip"`                              // 管理员上次登录IP
}
