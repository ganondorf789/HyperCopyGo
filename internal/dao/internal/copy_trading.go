// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CopyTradingDao is the data access object for the table copy_trading.
type CopyTradingDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  CopyTradingColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// CopyTradingColumns defines and stores column names for the table copy_trading.
type CopyTradingColumns struct {
	Id                             string //
	UserId                         string // 所属用户ID
	TargetWallet                   string // 目标钱包地址
	TargetWalletPlatform           string // 目标钱包平台
	Remark                         string // 备注
	Leverage                       string // 杠杆倍数
	MarginMode                     string // 保证金模式 1:逐仓 2:全仓
	FollowModel                    string // 跟单模式 1:固定金额 2:固定比例
	FollowModelValue               string // 跟单模式值
	MinValue                       string // 最小下单金额
	MaxValue                       string // 最大下单金额
	MaxMarginUsage                 string // 最大保证金使用率
	TpValue                        string // 止盈比例
	SlValue                        string // 止损比例
	OptReverseFollowOrder          string // 反向跟单 0:关 1:开
	OptFollowupDecrease            string // 跟随减仓 0:关 1:开
	OptFollowupIncrease            string // 跟随加仓 0:关 1:开
	OptForcedLiquidationProtection string // 强平保护 0:关 1:开
	OptPositionIncreaseOpening     string // 加仓开仓 0:关 1:开
	OptSlippageProtection          string // 滑点保护 0:关 1:开
	SymbolListType                 string // 交易对列表类型 WHITE:白名单 BLACK:黑名单
	SymbolList                     string // 交易对列表,逗号分隔
	MainWallet                     string // 主钱包地址
	MainWalletPlatform             string // 主钱包平台
	Status                         string // 状态 0:停用 1:启用
	CreatedAt                      string //
	UpdatedAt                      string //
}

// copyTradingColumns holds the columns for the table copy_trading.
var copyTradingColumns = CopyTradingColumns{
	Id:                             "id",
	UserId:                         "user_id",
	TargetWallet:                   "target_wallet",
	TargetWalletPlatform:           "target_wallet_platform",
	Remark:                         "remark",
	Leverage:                       "leverage",
	MarginMode:                     "margin_mode",
	FollowModel:                    "follow_model",
	FollowModelValue:               "follow_model_value",
	MinValue:                       "min_value",
	MaxValue:                       "max_value",
	MaxMarginUsage:                 "max_margin_usage",
	TpValue:                        "tp_value",
	SlValue:                        "sl_value",
	OptReverseFollowOrder:          "opt_reverse_follow_order",
	OptFollowupDecrease:            "opt_followup_decrease",
	OptFollowupIncrease:            "opt_followup_increase",
	OptForcedLiquidationProtection: "opt_forced_liquidation_protection",
	OptPositionIncreaseOpening:     "opt_position_increase_opening",
	OptSlippageProtection:          "opt_slippage_protection",
	SymbolListType:                 "symbol_list_type",
	SymbolList:                     "symbol_list",
	MainWallet:                     "main_wallet",
	MainWalletPlatform:             "main_wallet_platform",
	Status:                         "status",
	CreatedAt:                      "created_at",
	UpdatedAt:                      "updated_at",
}

// NewCopyTradingDao creates and returns a new DAO object for table data access.
func NewCopyTradingDao(handlers ...gdb.ModelHandler) *CopyTradingDao {
	return &CopyTradingDao{
		group:    "default",
		table:    "copy_trading",
		columns:  copyTradingColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CopyTradingDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CopyTradingDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CopyTradingDao) Columns() CopyTradingColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CopyTradingDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CopyTradingDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *CopyTradingDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
