// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProxyPoolsDao is the data access object for the table proxy_pools.
type ProxyPoolsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ProxyPoolsColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ProxyPoolsColumns defines and stores column names for the table proxy_pools.
type ProxyPoolsColumns struct {
	Id        string // 主键ID
	Host      string // 代理主机地址
	Port      string // 代理端口
	Username  string // 认证用户名
	Password  string // 认证密码
	Status    string // 状态（1=启用 0=禁用）
	Remark    string // 备注
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// proxyPoolsColumns holds the columns for the table proxy_pools.
var proxyPoolsColumns = ProxyPoolsColumns{
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

// NewProxyPoolsDao creates and returns a new DAO object for table data access.
func NewProxyPoolsDao(handlers ...gdb.ModelHandler) *ProxyPoolsDao {
	return &ProxyPoolsDao{
		group:    "default",
		table:    "proxy_pools",
		columns:  proxyPoolsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ProxyPoolsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ProxyPoolsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ProxyPoolsDao) Columns() ProxyPoolsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ProxyPoolsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ProxyPoolsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ProxyPoolsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
