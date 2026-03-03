// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TraderFundingsDao is the data access object for the table trader_fundings.
type TraderFundingsDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  TraderFundingsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// TraderFundingsColumns defines and stores column names for the table trader_fundings.
type TraderFundingsColumns struct {
	Id          string // 主键ID
	Address     string // 钱包地址
	Time        string // 资金费时间（毫秒时间戳）
	Hash        string // 交易哈希
	Coin        string // 币种
	Usdc        string // USDC金额（正=收入，负=支出）
	Szi         string // 持仓大小
	FundingRate string // 资金费率
	CreatedAt   string // 创建时间
}

// traderFundingsColumns holds the columns for the table trader_fundings.
var traderFundingsColumns = TraderFundingsColumns{
	Id:          "id",
	Address:     "address",
	Time:        "time",
	Hash:        "hash",
	Coin:        "coin",
	Usdc:        "usdc",
	Szi:         "szi",
	FundingRate: "funding_rate",
	CreatedAt:   "created_at",
}

// NewTraderFundingsDao creates and returns a new DAO object for table data access.
func NewTraderFundingsDao(handlers ...gdb.ModelHandler) *TraderFundingsDao {
	return &TraderFundingsDao{
		group:    "default",
		table:    "trader_fundings",
		columns:  traderFundingsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TraderFundingsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TraderFundingsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TraderFundingsDao) Columns() TraderFundingsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TraderFundingsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TraderFundingsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TraderFundingsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
