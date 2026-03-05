// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FetchFailuresDao is the data access object for the table fetch_failures.
type FetchFailuresDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  FetchFailuresColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// FetchFailuresColumns defines and stores column names for the table fetch_failures.
type FetchFailuresColumns struct {
	Id          string // 主键ID
	Type        string // 数据类型（fills/orders/funding）
	Reason      string // 失败原因（exceeds_limit/rate_limited）
	Address     string // 钱包地址
	StartMs     string // 失败时间窗口开始（毫秒时间戳）
	EndMs       string // 失败时间窗口结束（毫秒时间戳）
	RecordCount string // 该窗口返回的记录数
	CreatedAt   string // 创建时间
}

// fetchFailuresColumns holds the columns for the table fetch_failures.
var fetchFailuresColumns = FetchFailuresColumns{
	Id:          "id",
	Type:        "type",
	Reason:      "reason",
	Address:     "address",
	StartMs:     "start_ms",
	EndMs:       "end_ms",
	RecordCount: "record_count",
	CreatedAt:   "created_at",
}

// NewFetchFailuresDao creates and returns a new DAO object for table data access.
func NewFetchFailuresDao(handlers ...gdb.ModelHandler) *FetchFailuresDao {
	return &FetchFailuresDao{
		group:    "default",
		table:    "fetch_failures",
		columns:  fetchFailuresColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FetchFailuresDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FetchFailuresDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FetchFailuresDao) Columns() FetchFailuresColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FetchFailuresDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FetchFailuresDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *FetchFailuresDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
