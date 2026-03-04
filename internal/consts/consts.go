package consts

const (
	// JWT context key
	CtxUserIdKey   = "userId"
	CtxUserTypeKey = "userType"

	// 用户类型
	UserTypeUser  = "user"
	UserTypeAdmin = "admin"

	// 用户状态
	UserStatusEnabled  = 1
	UserStatusDisabled = 0

	// 管理员角色
	RoleAdmin      = "admin"
	RoleSuperAdmin = "super_admin"
)
