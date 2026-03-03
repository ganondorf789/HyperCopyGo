// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TraderPerformancesDao is the data access object for the table trader_performances.
type TraderPerformancesDao struct {
	table    string                    // table is the underlying table name of the DAO.
	group    string                    // group is the database configuration group name of the current DAO.
	columns  TraderPerformancesColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler        // handlers for customized model modification.
}

// TraderPerformancesColumns defines and stores column names for the table trader_performances.
type TraderPerformancesColumns struct {
	Id        string // 主键ID
	Address   string // 钱包地址
	Window    string // 统计窗口（day/week/month/allTime）
	Pnl       string // 盈亏
	Roi       string // 收益率
	Vlm       string // 交易量
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// traderPerformancesColumns holds the columns for the table trader_performances.
var traderPerformancesColumns = TraderPerformancesColumns{
	Id:        "id",
	Address:   "address",
	Window:    "window",
	Pnl:       "pnl",
	Roi:       "roi",
	Vlm:       "vlm",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewTraderPerformancesDao creates and returns a new DAO object for table data access.
func NewTraderPerformancesDao(handlers ...gdb.ModelHandler) *TraderPerformancesDao {
	return &TraderPerformancesDao{
		group:    "default",
		table:    "trader_performances",
		columns:  traderPerformancesColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TraderPerformancesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TraderPerformancesDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TraderPerformancesDao) Columns() TraderPerformancesColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TraderPerformancesDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TraderPerformancesDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TraderPerformancesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
