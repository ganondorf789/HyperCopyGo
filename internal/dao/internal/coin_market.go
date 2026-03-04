// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CoinMarketDao is the data access object for the table coin_market.
type CoinMarketDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  CoinMarketColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// CoinMarketColumns defines and stores column names for the table coin_market.
type CoinMarketColumns struct {
	Id               string // 主键ID
	Coin             string // 币种名称
	Price            string // 当前价格
	Change24H        string // 24h价格变动
	ChangePercent24H string // 24h价格变动百分比
	Open24H          string // 24h开盘价
	Close24H         string // 24h收盘价
	High24H          string // 24h最高价
	Low24H           string // 24h最低价
	Volume24H        string // 24h成交量
	QuoteVolume24H   string // 24h计价成交额
	Funding          string // 资金费率
	OpenInterest     string // 未平仓合约量
	CreatedAt        string // 创建时间
	UpdatedAt        string // 更新时间
}

// coinMarketColumns holds the columns for the table coin_market.
var coinMarketColumns = CoinMarketColumns{
	Id:               "id",
	Coin:             "coin",
	Price:            "price",
	Change24H:        "change24h",
	ChangePercent24H: "change_percent24h",
	Open24H:          "open24h",
	Close24H:         "close24h",
	High24H:          "high24h",
	Low24H:           "low24h",
	Volume24H:        "volume24h",
	QuoteVolume24H:   "quote_volume24h",
	Funding:          "funding",
	OpenInterest:     "open_interest",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
}

// NewCoinMarketDao creates and returns a new DAO object for table data access.
func NewCoinMarketDao(handlers ...gdb.ModelHandler) *CoinMarketDao {
	return &CoinMarketDao{
		group:    "default",
		table:    "coin_market",
		columns:  coinMarketColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CoinMarketDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CoinMarketDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CoinMarketDao) Columns() CoinMarketColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CoinMarketDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CoinMarketDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *CoinMarketDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
