// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProxyPoolDao is the data access object for the table proxy_pool.
type ProxyPoolDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ProxyPoolColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ProxyPoolColumns defines and stores column names for the table proxy_pool.
type ProxyPoolColumns struct {
	Id        string //
	Host      string // 代理主机地址
	Port      string // 代理端口
	Username  string // 认证用户名
	Password  string // 认证密码
	Status    string // 状态: 1=启用 0=禁用
	Remark    string // 备注
	CreatedAt string //
	UpdatedAt string //
}

// proxyPoolColumns holds the columns for the table proxy_pool.
var proxyPoolColumns = ProxyPoolColumns{
	Id:        "id",
	Host:      "host",
	Port:      "port",
	Username:  "username",
	Password:  "password",
	Status:    "status",
	Remark:    "remark",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewProxyPoolDao creates and returns a new DAO object for table data access.
func NewProxyPoolDao(handlers ...gdb.ModelHandler) *ProxyPoolDao {
	return &ProxyPoolDao{
		group:    "default",
		table:    "proxy_pool",
		columns:  proxyPoolColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ProxyPoolDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ProxyPoolDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ProxyPoolDao) Columns() ProxyPoolColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ProxyPoolDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ProxyPoolDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ProxyPoolDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
