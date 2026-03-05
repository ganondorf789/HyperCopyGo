// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FillFetchFailuresDao is the data access object for the table fill_fetch_failures.
type FillFetchFailuresDao struct {
	table    string                   // table is the underlying table name of the DAO.
	group    string                   // group is the database configuration group name of the current DAO.
	columns  FillFetchFailuresColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler       // handlers for customized model modification.
}

// FillFetchFailuresColumns defines and stores column names for the table fill_fetch_failures.
type FillFetchFailuresColumns struct {
	Id        string // 主键ID
	Address   string // 钱包地址
	StartMs   string // 失败时间窗口开始（毫秒时间戳）
	EndMs     string // 失败时间窗口结束（毫秒时间戳）
	FillCount string // 该窗口返回的记录数
	CreatedAt string // 创建时间
}

// fillFetchFailuresColumns holds the columns for the table fill_fetch_failures.
var fillFetchFailuresColumns = FillFetchFailuresColumns{
	Id:        "id",
	Address:   "address",
	StartMs:   "start_ms",
	EndMs:     "end_ms",
	FillCount: "fill_count",
	CreatedAt: "created_at",
}

// NewFillFetchFailuresDao creates and returns a new DAO object for table data access.
func NewFillFetchFailuresDao(handlers ...gdb.ModelHandler) *FillFetchFailuresDao {
	return &FillFetchFailuresDao{
		group:    "default",
		table:    "fill_fetch_failures",
		columns:  fillFetchFailuresColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FillFetchFailuresDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FillFetchFailuresDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FillFetchFailuresDao) Columns() FillFetchFailuresColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FillFetchFailuresDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FillFetchFailuresDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *FillFetchFailuresDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
