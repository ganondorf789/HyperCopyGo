// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TraderPnlHistoriesDao is the data access object for the table trader_pnl_histories.
type TraderPnlHistoriesDao struct {
	table    string                    // table is the underlying table name of the DAO.
	group    string                    // group is the database configuration group name of the current DAO.
	columns  TraderPnlHistoriesColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler        // handlers for customized model modification.
}

// TraderPnlHistoriesColumns defines and stores column names for the table trader_pnl_histories.
type TraderPnlHistoriesColumns struct {
	Id        string // 主键ID
	Address   string // 钱包地址
	Window    string // 统计窗口（day/week/month/allTime）
	History   string // 盈亏历史数据（[[timestamp, value], ...]）
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// traderPnlHistoriesColumns holds the columns for the table trader_pnl_histories.
var traderPnlHistoriesColumns = TraderPnlHistoriesColumns{
	Id:        "id",
	Address:   "address",
	Window:    "window",
	History:   "history",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewTraderPnlHistoriesDao creates and returns a new DAO object for table data access.
func NewTraderPnlHistoriesDao(handlers ...gdb.ModelHandler) *TraderPnlHistoriesDao {
	return &TraderPnlHistoriesDao{
		group:    "default",
		table:    "trader_pnl_histories",
		columns:  traderPnlHistoriesColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TraderPnlHistoriesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TraderPnlHistoriesDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TraderPnlHistoriesDao) Columns() TraderPnlHistoriesColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TraderPnlHistoriesDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TraderPnlHistoriesDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TraderPnlHistoriesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
