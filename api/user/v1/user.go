package v1

import "github.com/gogf/gf/v2/frame/g"

// 用户注册
type UserRegisterReq struct {
	g.Meta   `path:"/user/register" tags:"User" method:"post" summary:"用户注册"`
	Username string `json:"username" v:"required|length:3,32#请输入用户名|用户名长度3-32位"`
	Password string `json:"password" v:"required|length:6,32#请输入密码|密码长度6-32位"`
	Nickname string `json:"nickname"`
}
type UserRegisterRes struct {
	g.Meta `mime:"application/json"`
}

// 用户登录
type UserLoginReq struct {
	g.Meta   `path:"/user/login" tags:"User" method:"post" summary:"用户登录"`
	Username string `json:"username" v:"required#请输入用户名"`
	Password string `json:"password" v:"required#请输入密码"`
}
type UserLoginRes struct {
	g.Meta `mime:"application/json"`
	Token  string `json:"token"`
	Expire int64  `json:"expire"`
}

// 获取当前用户信息（需登录）
type UserProfileReq struct {
	g.Meta `path:"/user/profile" tags:"User" method:"get" summary:"获取当前用户信息"`
}
type UserProfileRes struct {
	g.Meta   `mime:"application/json"`
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
