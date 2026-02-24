// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BaseCopyTrading 跟单配置公共字段，供 API 层复用以减少重复定义
type BaseCopyTrading struct {
	TargetWallet                   string  `json:"targetWallet"                   orm:"target_wallet"`                     // 目标钱包地址
	TargetWalletPlatform           string  `json:"targetWalletPlatform"           orm:"target_wallet_platform"`            // 目标钱包平台
	Remark                         string  `json:"remark"                         orm:"remark"`                            // 备注
	Leverage                       int     `json:"leverage"                       orm:"leverage"`                          // 杠杆倍数
	MarginMode                     int     `json:"marginMode"                     orm:"margin_mode"`                       // 保证金模式 1:逐仓 2:全仓
	FollowModel                    int     `json:"followModel"                    orm:"follow_model"`                      // 跟单模式 1:固定金额 2:固定比例
	FollowModelValue               float64 `json:"followModelValue"               orm:"follow_model_value"`                // 跟单模式值
	MinValue                       float64 `json:"minValue"                       orm:"min_value"`                         // 最小下单金额
	MaxValue                       float64 `json:"maxValue"                       orm:"max_value"`                         // 最大下单金额
	MaxMarginUsage                 float64 `json:"maxMarginUsage"                 orm:"max_margin_usage"`                  // 最大保证金使用率
	TpValue                        float64 `json:"tpValue"                        orm:"tp_value"`                          // 止盈比例
	SlValue                        float64 `json:"slValue"                        orm:"sl_value"`                          // 止损比例
	OptReverseFollowOrder          int     `json:"optReverseFollowOrder"          orm:"opt_reverse_follow_order"`          // 反向跟单 0:关 1:开
	OptFollowupDecrease            int     `json:"optFollowupDecrease"            orm:"opt_followup_decrease"`             // 跟随减仓 0:关 1:开
	OptFollowupIncrease            int     `json:"optFollowupIncrease"            orm:"opt_followup_increase"`             // 跟随加仓 0:关 1:开
	OptForcedLiquidationProtection int     `json:"optForcedLiquidationProtection" orm:"opt_forced_liquidation_protection"` // 强平保护 0:关 1:开
	OptPositionIncreaseOpening     int     `json:"optPositionIncreaseOpening"     orm:"opt_position_increase_opening"`     // 加仓开仓 0:关 1:开
	OptSlippageProtection          int     `json:"optSlippageProtection"          orm:"opt_slippage_protection"`           // 滑点保护 0:关 1:开
	SymbolListType                 string  `json:"symbolListType"                 orm:"symbol_list_type"`                  // 交易对列表类型 WHITE:白名单 BLACK:黑名单
	SymbolList                     string  `json:"symbolList"                     orm:"symbol_list"`                       // 交易对列表,逗号分隔
	MainWallet                     string  `json:"mainWallet"                     orm:"main_wallet"`                       // 主钱包地址
	MainWalletPlatform             string  `json:"mainWalletPlatform"             orm:"main_wallet_platform"`              // 主钱包平台
}

// CopyTrading is the golang structure for table copy_trading.
type CopyTrading struct {
	Id        int64       `json:"id"        orm:"id"`
	UserId    int64       `json:"userId"    orm:"user_id"`    // 所属用户ID
	BaseCopyTrading
	Status    int         `json:"status"    orm:"status"`     // 状态 0:停用 1:启用
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at"`
}
