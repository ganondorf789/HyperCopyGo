package v1

import "github.com/gogf/gf/v2/frame/g"

// 发送邮箱验证码（免鉴权）
type SendVerifyCodeReq struct {
	g.Meta `path:"/user/send-code" tags:"User" method:"post" summary:"发送邮箱验证码"`
	Email  string `json:"email" v:"required|email#请输入邮箱|邮箱格式不正确"`
}
type SendVerifyCodeRes struct {
	g.Meta `mime:"application/json"`
}

// 邮箱验证码登录（免鉴权）
type UserLoginReq struct {
	g.Meta `path:"/user/login" tags:"User" method:"post" summary:"邮箱验证码登录"`
	Email  string `json:"email" v:"required|email#请输入邮箱|邮箱格式不正确"`
	Code   string `json:"code" v:"required|length:6,6#请输入验证码|验证码为6位数字"`
}
type UserLoginRes struct {
	g.Meta `mime:"application/json"`
	Token  string `json:"token"`
	Expire int64  `json:"expire"`
}

// 获取当前用户信息（需登录）
type UserProfileReq struct {
	g.Meta `path:"/user/profile" tags:"User" method:"get" summary:"获取当前用户信息" login_required:"true"`
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
