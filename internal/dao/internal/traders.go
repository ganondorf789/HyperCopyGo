// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TradersDao is the data access object for the table traders.
type TradersDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  TradersColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// TradersColumns defines and stores column names for the table traders.
type TradersColumns struct {
	Id                     string // 主键ID
	TwitterName            string // 推特显示名
	Username               string // 推特用户名
	Address                string // 钱包地址
	ProfilePicture         string // 头像链接
	Labels                 string // 标签列表
	SnapEffLeverage        string // 快照-有效杠杆
	SnapLongPositionCount  string // 快照-多头持仓数
	SnapLongPositionValue  string // 快照-多头持仓价值
	SnapMarginUsageRate    string // 快照-保证金使用率
	SnapPerpValue          string // 快照-永续合约价值
	SnapPositionCount      string // 快照-总持仓数
	SnapPositionValue      string // 快照-总持仓价值
	SnapShortPositionCount string // 快照-空头持仓数
	SnapShortPositionValue string // 快照-空头持仓价值
	SnapSpotValue          string // 快照-现货价值
	SnapTotalMarginUsed    string // 快照-已用保证金
	SnapTotalValue         string // 快照-总价值
	SnapUnrealizedPnl      string // 快照-未实现盈亏
	ShortPnl               string // 空头盈亏
	ShortWinRate           string // 空头胜率
	LongPnl                string // 多头盈亏
	LongWinRate            string // 多头胜率
	CreatedAt              string // 创建时间
	UpdatedAt              string // 更新时间
}

// tradersColumns holds the columns for the table traders.
var tradersColumns = TradersColumns{
	Id:                     "id",
	TwitterName:            "twitter_name",
	Username:               "username",
	Address:                "address",
	ProfilePicture:         "profile_picture",
	Labels:                 "labels",
	SnapEffLeverage:        "snap_eff_leverage",
	SnapLongPositionCount:  "snap_long_position_count",
	SnapLongPositionValue:  "snap_long_position_value",
	SnapMarginUsageRate:    "snap_margin_usage_rate",
	SnapPerpValue:          "snap_perp_value",
	SnapPositionCount:      "snap_position_count",
	SnapPositionValue:      "snap_position_value",
	SnapShortPositionCount: "snap_short_position_count",
	SnapShortPositionValue: "snap_short_position_value",
	SnapSpotValue:          "snap_spot_value",
	SnapTotalMarginUsed:    "snap_total_margin_used",
	SnapTotalValue:         "snap_total_value",
	SnapUnrealizedPnl:      "snap_unrealized_pnl",
	ShortPnl:               "short_pnl",
	ShortWinRate:           "short_win_rate",
	LongPnl:                "long_pnl",
	LongWinRate:            "long_win_rate",
	CreatedAt:              "created_at",
	UpdatedAt:              "updated_at",
}

// NewTradersDao creates and returns a new DAO object for table data access.
func NewTradersDao(handlers ...gdb.ModelHandler) *TradersDao {
	return &TradersDao{
		group:    "default",
		table:    "traders",
		columns:  tradersColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TradersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TradersDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TradersDao) Columns() TradersColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TradersDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TradersDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TradersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
