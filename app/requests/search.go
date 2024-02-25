package requests

type Search struct {
	Name     string `json:"name" form:"name"`
	PageNum  int64  `json:"pageNum" form:"pageNum"`   // 第几页
	PageSize int64  `json:"pageSize" form:"pageSize"` // 每页多少条
}
