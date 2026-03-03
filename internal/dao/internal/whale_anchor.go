// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WhaleAnchorDao is the data access object for the table whale_anchor.
type WhaleAnchorDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  WhaleAnchorColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// WhaleAnchorColumns defines and stores column names for the table whale_anchor.
type WhaleAnchorColumns struct {
	Id             string // 主键ID
	Symbol         string // 交易对符号
	Volume24H      string // 24h成交量(USD)
	OpenInterest   string // 当前未平仓合约量(USD)
	Depth1Pct      string // 1%盘口深度(USD)
	ValVolume      string // 0.4% x 24h Volume
	ValOi          string // 1% x OI
	ValDepth       string // 30% x 1% Depth
	WhaleThreshold string // 巨鲸仓位阈值 max(val_volume,val_oi,val_depth)
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
}

// whaleAnchorColumns holds the columns for the table whale_anchor.
var whaleAnchorColumns = WhaleAnchorColumns{
	Id:             "id",
	Symbol:         "symbol",
	Volume24H:      "volume24h",
	OpenInterest:   "open_interest",
	Depth1Pct:      "depth1pct",
	ValVolume:      "val_volume",
	ValOi:          "val_oi",
	ValDepth:       "val_depth",
	WhaleThreshold: "whale_threshold",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewWhaleAnchorDao creates and returns a new DAO object for table data access.
func NewWhaleAnchorDao(handlers ...gdb.ModelHandler) *WhaleAnchorDao {
	return &WhaleAnchorDao{
		group:    "default",
		table:    "whale_anchor",
		columns:  whaleAnchorColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *WhaleAnchorDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *WhaleAnchorDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *WhaleAnchorDao) Columns() WhaleAnchorColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *WhaleAnchorDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *WhaleAnchorDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *WhaleAnchorDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
