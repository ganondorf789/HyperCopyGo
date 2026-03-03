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
	Id         string // 主键ID
	Address    string // 钱包地址
	Coin       string // 币种
	MarginMode string // 保证金模式（isolated/cross）
	Direction  string // 方向（long/short）
	Size       string // 最大持仓量
	EntryPrice string // 加权平均入场价
	ClosePrice string // 加权平均平仓价
	StartTime  string // 开仓时间（毫秒时间戳）
	EndTime    string // 平仓时间（毫秒时间戳）
	TotalFee   string // 总手续费
	Pnl        string // 已实现盈亏（closedPnl 之和）
	FillCount  string // 成交笔数
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
}

// completedTradesColumns holds the columns for the table completed_trades.
var completedTradesColumns = CompletedTradesColumns{
	Id:         "id",
	Address:    "address",
	Coin:       "coin",
	MarginMode: "margin_mode",
	Direction:  "direction",
	Size:       "size",
	EntryPrice: "entry_price",
	ClosePrice: "close_price",
	StartTime:  "start_time",
	EndTime:    "end_time",
	TotalFee:   "total_fee",
	Pnl:        "pnl",
	FillCount:  "fill_count",
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
