// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TraderStatisticsDao is the data access object for the table trader_statistics.
type TraderStatisticsDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  TraderStatisticsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// TraderStatisticsColumns defines and stores column names for the table trader_statistics.
type TraderStatisticsColumns struct {
	Id                 string // 主键ID
	Address            string // 钱包地址
	Window             string // 统计窗口（day/week/month/allTime）
	Sharpe             string // 夏普比率
	Drawdown           string // 最大回撤
	PositionCount      string // 持仓数
	TotalValue         string // 账户总价值
	PerpValue          string // 永续合约总价值
	PositionValue      string // 持仓价值
	LongPositionValue  string // 多仓仓位价值
	ShortPositionValue string // 空仓仓位价值
	MarginUsage        string // 保证金使用率
	UsedMargin         string // 已用保证金
	ProfitCount        string // 盈利次数
	WinRate            string // 胜率
	TotalPnl           string // 总盈亏
	LongCount          string // 多仓数
	LongRealizedPnl    string // 多仓已实现盈亏
	LongWinRate        string // 多仓胜率
	ShortCount         string // 空仓数
	ShortRealizedPnl   string // 空仓已实现盈亏
	ShortWinRate       string // 空仓胜率
	UnrealizedPnl      string // 未实现盈亏
	AvgLeverage        string // 平均杠杆
	Coins              string // 交易过的币种
	CreatedAt          string // 创建时间
	UpdatedAt          string // 更新时间
	TotalRealizedPnl   string // 已实现总盈亏（正为盈利，负为亏损）
}

// traderStatisticsColumns holds the columns for the table trader_statistics.
var traderStatisticsColumns = TraderStatisticsColumns{
	Id:                 "id",
	Address:            "address",
	Window:             "window",
	Sharpe:             "sharpe",
	Drawdown:           "drawdown",
	PositionCount:      "position_count",
	TotalValue:         "total_value",
	PerpValue:          "perp_value",
	PositionValue:      "position_value",
	LongPositionValue:  "long_position_value",
	ShortPositionValue: "short_position_value",
	MarginUsage:        "margin_usage",
	UsedMargin:         "used_margin",
	ProfitCount:        "profit_count",
	WinRate:            "win_rate",
	TotalPnl:           "total_pnl",
	LongCount:          "long_count",
	LongRealizedPnl:    "long_realized_pnl",
	LongWinRate:        "long_win_rate",
	ShortCount:         "short_count",
	ShortRealizedPnl:   "short_realized_pnl",
	ShortWinRate:       "short_win_rate",
	UnrealizedPnl:      "unrealized_pnl",
	AvgLeverage:        "avg_leverage",
	Coins:              "coins",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
	TotalRealizedPnl:   "total_realized_pnl",
}

// NewTraderStatisticsDao creates and returns a new DAO object for table data access.
func NewTraderStatisticsDao(handlers ...gdb.ModelHandler) *TraderStatisticsDao {
	return &TraderStatisticsDao{
		group:    "default",
		table:    "trader_statistics",
		columns:  traderStatisticsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TraderStatisticsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TraderStatisticsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TraderStatisticsDao) Columns() TraderStatisticsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TraderStatisticsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TraderStatisticsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TraderStatisticsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
