// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TraderFillsDao is the data access object for the table trader_fills.
type TraderFillsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  TraderFillsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// TraderFillsColumns defines and stores column names for the table trader_fills.
type TraderFillsColumns struct {
	Id            string // 主键ID
	Address       string // 钱包地址
	Coin          string // 币种
	Px            string // 成交价
	Sz            string // 成交量
	Side          string // 买卖方向（A=卖/B=买）
	Time          string // 成交时间（毫秒时间戳）
	StartPosition string // 成交前仓位大小
	Dir           string // 操作方向（Open Long/Open Short/Close Long/Close Short）
	ClosedPnl     string // 平仓盈亏
	Hash          string // 交易哈希
	Oid           string // 订单ID
	Crossed       string // 是否为全仓模式
	Fee           string // 手续费
	Tid           string // 成交ID
	Cloid         string // 客户端订单ID
	FeeToken      string // 手续费计价币种
	CreatedAt     string // 创建时间
}

// traderFillsColumns holds the columns for the table trader_fills.
var traderFillsColumns = TraderFillsColumns{
	Id:            "id",
	Address:       "address",
	Coin:          "coin",
	Px:            "px",
	Sz:            "sz",
	Side:          "side",
	Time:          "time",
	StartPosition: "start_position",
	Dir:           "dir",
	ClosedPnl:     "closed_pnl",
	Hash:          "hash",
	Oid:           "oid",
	Crossed:       "crossed",
	Fee:           "fee",
	Tid:           "tid",
	Cloid:         "cloid",
	FeeToken:      "fee_token",
	CreatedAt:     "created_at",
}

// NewTraderFillsDao creates and returns a new DAO object for table data access.
func NewTraderFillsDao(handlers ...gdb.ModelHandler) *TraderFillsDao {
	return &TraderFillsDao{
		group:    "default",
		table:    "trader_fills",
		columns:  traderFillsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TraderFillsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TraderFillsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TraderFillsDao) Columns() TraderFillsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TraderFillsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TraderFillsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TraderFillsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
