// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CompletedTradesDao is the data access object for the table completed_trades.
type CompletedTradesDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  CompletedTradesColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// CompletedTradesColumns defines and stores column names for the table completed_trades.
type CompletedTradesColumns struct {
	Id         string //
	User       string // 交易员钱包地址
	Coin       string // 币种
	Side       string // 方向 long/short
	EntryPx    string // 开仓均价
	ClosePx    string // 平仓均价
	Sz         string // 交易数量
	ClosedPnl  string // 已实现盈亏
	TotalFee   string // 总手续费
	OpenTime   string // 开仓时间戳(ms)
	CloseTime  string // 平仓时间戳(ms)
	DurationMs string // 持仓时长(ms)
	CreatedAt  string //
	UpdatedAt  string //
}

// completedTradesColumns holds the columns for the table completed_trades.
var completedTradesColumns = CompletedTradesColumns{
	Id:         "id",
	User:       "user",
	Coin:       "coin",
	Side:       "side",
	EntryPx:    "entry_px",
	ClosePx:    "close_px",
	Sz:         "sz",
	ClosedPnl:  "closed_pnl",
	TotalFee:   "total_fee",
	OpenTime:   "open_time",
	CloseTime:  "close_time",
	DurationMs: "duration_ms",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewCompletedTradesDao creates and returns a new DAO object for table data access.
func NewCompletedTradesDao(handlers ...gdb.ModelHandler) *CompletedTradesDao {
	return &CompletedTradesDao{
		group:    "default",
		table:    "completed_trades",
		columns:  completedTradesColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CompletedTradesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CompletedTradesDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CompletedTradesDao) Columns() CompletedTradesColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CompletedTradesDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CompletedTradesDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *CompletedTradesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
