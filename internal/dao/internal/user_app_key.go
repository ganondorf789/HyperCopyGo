// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserAppKeyDao is the data access object for the table user_app_key.
type UserAppKeyDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  UserAppKeyColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// UserAppKeyColumns defines and stores column names for the table user_app_key.
type UserAppKeyColumns struct {
	Id        string // 主键ID
	UserId    string // 所属用户ID
	AppId     string // AppID
	AppSecret string // AppSecret
	Remark    string // 备注
	ExpireAt  string // 过期时间,NULL表示永不过期
	Status    string // 状态 1:启用 0:禁用
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// userAppKeyColumns holds the columns for the table user_app_key.
var userAppKeyColumns = UserAppKeyColumns{
	Id:        "id",
	UserId:    "user_id",
	AppId:     "app_id",
	AppSecret: "app_secret",
	Remark:    "remark",
	ExpireAt:  "expire_at",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewUserAppKeyDao creates and returns a new DAO object for table data access.
func NewUserAppKeyDao(handlers ...gdb.ModelHandler) *UserAppKeyDao {
	return &UserAppKeyDao{
		group:    "default",
		table:    "user_app_key",
		columns:  userAppKeyColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserAppKeyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserAppKeyDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserAppKeyDao) Columns() UserAppKeyColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserAppKeyDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserAppKeyDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *UserAppKeyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
