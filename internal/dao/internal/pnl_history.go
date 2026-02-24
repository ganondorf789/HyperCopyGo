// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PnlHistoryDao is the data access object for the table pnl_history.
type PnlHistoryDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  PnlHistoryColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// PnlHistoryColumns defines and stores column names for the table pnl_history.
type PnlHistoryColumns struct {
	Id        string //
	User      string // 交易员钱包地址
	Timeframe string // 时间框架: 1D, 7D, 30D, All
	PnlList   string // PnL数据点列表
	CreatedAt string //
	UpdatedAt string //
}

// pnlHistoryColumns holds the columns for the table pnl_history.
var pnlHistoryColumns = PnlHistoryColumns{
	Id:        "id",
	User:      "user",
	Timeframe: "timeframe",
	PnlList:   "pnl_list",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewPnlHistoryDao creates and returns a new DAO object for table data access.
func NewPnlHistoryDao(handlers ...gdb.ModelHandler) *PnlHistoryDao {
	return &PnlHistoryDao{
		group:    "default",
		table:    "pnl_history",
		columns:  pnlHistoryColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PnlHistoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PnlHistoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PnlHistoryDao) Columns() PnlHistoryColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PnlHistoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PnlHistoryDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PnlHistoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
