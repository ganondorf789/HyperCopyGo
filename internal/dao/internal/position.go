// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PositionDao is the data access object for the table position.
type PositionDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  PositionColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// PositionColumns defines and stores column names for the table position.
type PositionColumns struct {
	Id               string //
	User             string // 用户钱包地址
	Symbol           string // 交易对符号
	PositionSize     string // 持仓数量（负数为空头）
	EntryPrice       string // 开仓均价
	MarkPrice        string // 标记价格
	LiqPrice         string // 强平价格
	Leverage         string // 杠杆倍数
	MarginBalance    string // 保证金余额
	PositionValueUsd string // 持仓价值(USD)
	UnrealizedPnl    string // 未实现盈亏
	FundingFee       string // 资金费用
	MarginMode       string // 保证金模式 cross/isolated
	Labels           string // 标签，逗号分隔
	CreatedAt        string //
	UpdatedAt        string //
}

// positionColumns holds the columns for the table position.
var positionColumns = PositionColumns{
	Id:               "id",
	User:             "user",
	Symbol:           "symbol",
	PositionSize:     "position_size",
	EntryPrice:       "entry_price",
	MarkPrice:        "mark_price",
	LiqPrice:         "liq_price",
	Leverage:         "leverage",
	MarginBalance:    "margin_balance",
	PositionValueUsd: "position_value_usd",
	UnrealizedPnl:    "unrealized_pnl",
	FundingFee:       "funding_fee",
	MarginMode:       "margin_mode",
	Labels:           "labels",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
}

// NewPositionDao creates and returns a new DAO object for table data access.
func NewPositionDao(handlers ...gdb.ModelHandler) *PositionDao {
	return &PositionDao{
		group:    "default",
		table:    "position",
		columns:  positionColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PositionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PositionDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PositionDao) Columns() PositionColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PositionDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PositionDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PositionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
