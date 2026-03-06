// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HotCoinDao is the data access object for the table hot_coin.
type HotCoinDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  HotCoinColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// HotCoinColumns defines and stores column names for the table hot_coin.
type HotCoinColumns struct {
	Id        string // 主键ID
	Coin      string // 币种名称
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// hotCoinColumns holds the columns for the table hot_coin.
var hotCoinColumns = HotCoinColumns{
	Id:        "id",
	Coin:      "coin",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewHotCoinDao creates and returns a new DAO object for table data access.
func NewHotCoinDao(handlers ...gdb.ModelHandler) *HotCoinDao {
	return &HotCoinDao{
		group:    "default",
		table:    "hot_coin",
		columns:  hotCoinColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *HotCoinDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *HotCoinDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *HotCoinDao) Columns() HotCoinColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *HotCoinDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *HotCoinDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *HotCoinDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
