// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TraderPositionsDao is the data access object for the table trader_positions.
type TraderPositionsDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  TraderPositionsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// TraderPositionsColumns defines and stores column names for the table trader_positions.
type TraderPositionsColumns struct {
	Id                    string // 主键ID
	Address               string // 钱包地址
	Coin                  string // 币种
	Szi                   string // 仓位大小（正值为多头，负值为空头）
	LeverageType          string // 杠杆类型（cross/isolated）
	Leverage              string // 杠杆倍数
	EntryPx               string // 入场价
	PositionValue         string // 持仓价值
	UnrealizedPnl         string // 未实现盈亏
	ReturnOnEquity        string // 权益回报率
	LiquidationPx         string // 清算价
	MarginUsed            string // 已用保证金
	MaxLeverage           string // 最大允许杠杆
	CumFundingAllTime     string // 累计资金费（全部时间）
	CumFundingSinceOpen   string // 累计资金费（开仓以来）
	CumFundingSinceChange string // 累计资金费（最近变更以来）
	CreatedAt             string // 创建时间
	UpdatedAt             string // 更新时间
}

// traderPositionsColumns holds the columns for the table trader_positions.
var traderPositionsColumns = TraderPositionsColumns{
	Id:                    "id",
	Address:               "address",
	Coin:                  "coin",
	Szi:                   "szi",
	LeverageType:          "leverage_type",
	Leverage:              "leverage",
	EntryPx:               "entry_px",
	PositionValue:         "position_value",
	UnrealizedPnl:         "unrealized_pnl",
	ReturnOnEquity:        "return_on_equity",
	LiquidationPx:         "liquidation_px",
	MarginUsed:            "margin_used",
	MaxLeverage:           "max_leverage",
	CumFundingAllTime:     "cum_funding_all_time",
	CumFundingSinceOpen:   "cum_funding_since_open",
	CumFundingSinceChange: "cum_funding_since_change",
	CreatedAt:             "created_at",
	UpdatedAt:             "updated_at",
}

// NewTraderPositionsDao creates and returns a new DAO object for table data access.
func NewTraderPositionsDao(handlers ...gdb.ModelHandler) *TraderPositionsDao {
	return &TraderPositionsDao{
		group:    "default",
		table:    "trader_positions",
		columns:  traderPositionsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TraderPositionsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TraderPositionsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TraderPositionsDao) Columns() TraderPositionsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TraderPositionsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TraderPositionsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TraderPositionsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
