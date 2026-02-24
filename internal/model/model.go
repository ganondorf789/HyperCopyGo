package model

// 用户注册输入
type UserRegisterInput struct {
	Username string
	Password string
	Nickname string
}

// 用户登录输入
type UserLoginInput struct {
	Username string
	Password string
}

// 用户信息输出
type UserInfoOutput struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Status   int    `json:"status"`
}

// 登录输出
type TokenOutput struct {
	Token  string `json:"token"`
	Expire int64  `json:"expire"`
}

// 管理员登录输入
type AdminLoginInput struct {
	Username string
	Password string
}

// 管理员信息输出
type AdminInfoOutput struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Realname string `json:"realname"`
	Role     string `json:"role"`
	Status   int    `json:"status"`
}

// 管理员用户列表查询
type AdminUserListInput struct {
	Page     int
	PageSize int
	Status   int
}

type AdminUserListOutput struct {
	List  []UserInfoOutput `json:"list"`
	Total int              `json:"total"`
	Page  int              `json:"page"`
}

// 管理员修改用户状态
type AdminUserStatusInput struct {
	Id     int64
	Status int
}
