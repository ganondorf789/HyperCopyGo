// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FillsDao is the data access object for the table fills.
type FillsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  FillsColumns       // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// FillsColumns defines and stores column names for the table fills.
type FillsColumns struct {
	Id            string //
	User          string // 交易员钱包地址
	Coin          string // 币种
	Dir           string // 方向描述
	Side          string // 买卖方向 B/A
	Px            string // 成交价格
	Sz            string // 成交数量
	ClosedPnl     string // 已实现盈亏
	Fee           string // 手续费
	FeeToken      string // 手续费币种
	BuilderFee    string // builder手续费
	Hash          string // 交易哈希
	Oid           string // 订单ID
	Tid           string // 成交ID
	Crossed       string // 是否crossed
	StartPosition string // 成交前持仓
	FillTime      string // 成交时间戳(ms)
	CreatedAt     string //
	UpdatedAt     string //
}

// fillsColumns holds the columns for the table fills.
var fillsColumns = FillsColumns{
	Id:            "id",
	User:          "user",
	Coin:          "coin",
	Dir:           "dir",
	Side:          "side",
	Px:            "px",
	Sz:            "sz",
	ClosedPnl:     "closed_pnl",
	Fee:           "fee",
	FeeToken:      "fee_token",
	BuilderFee:    "builder_fee",
	Hash:          "hash",
	Oid:           "oid",
	Tid:           "tid",
	Crossed:       "crossed",
	StartPosition: "start_position",
	FillTime:      "fill_time",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewFillsDao creates and returns a new DAO object for table data access.
func NewFillsDao(handlers ...gdb.ModelHandler) *FillsDao {
	return &FillsDao{
		group:    "default",
		table:    "fills",
		columns:  fillsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FillsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FillsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FillsDao) Columns() FillsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FillsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FillsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *FillsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
