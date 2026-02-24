package v1

import "github.com/gogf/gf/v2/frame/g"

// 管理员登录
type AdminLoginReq struct {
	g.Meta   `path:"/admin/login" tags:"Admin" method:"post" summary:"管理员登录"`
	Username string `json:"username" v:"required#请输入用户名"`
	Password string `json:"password" v:"required#请输入密码"`
}
type AdminLoginRes struct {
	g.Meta `mime:"application/json"`
	Token  string `json:"token"`
	Expire int64  `json:"expire"`
}

// 管理员获取自身信息
type AdminProfileReq struct {
	g.Meta `path:"/admin/profile" tags:"Admin" method:"get" summary:"管理员信息"`
}
type AdminProfileRes struct {
	g.Meta   `mime:"application/json"`
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Realname string `json:"realname"`
	Role     string `json:"role"`
}

// 管理员获取用户列表
type AdminUserListReq struct {
	g.Meta   `path:"/admin/users" tags:"Admin" method:"get" summary:"用户列表"`
	Page     int `json:"page" d:"1"`
	PageSize int `json:"pageSize" d:"20" v:"max:100#每页最多100条"`
	Status   int `json:"status" d:"-1"`
}
type AdminUserListRes struct {
	g.Meta `mime:"application/json"`
	List   []AdminUserItem `json:"list"`
	Total  int             `json:"total"`
	Page   int             `json:"page"`
}
type AdminUserItem struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Status   int    `json:"status"`
}

// 管理员修改用户状态（启用/禁用）
type AdminUserStatusReq struct {
	g.Meta `path:"/admin/users/{id}/status" tags:"Admin" method:"put" summary:"修改用户状态"`
	Id     int64 `json:"id" in:"path" v:"required"`
	Status int   `json:"status" v:"required|in:0,1#请输入状态|状态只能是0或1"`
}
type AdminUserStatusRes struct {
	g.Meta `mime:"application/json"`
}

// 管理员删除用户
type AdminUserDeleteReq struct {
	g.Meta `path:"/admin/users/{id}" tags:"Admin" method:"delete" summary:"删除用户"`
	Id     int64 `json:"id" in:"path" v:"required"`
}
type AdminUserDeleteRes struct {
	g.Meta `mime:"application/json"`
}
