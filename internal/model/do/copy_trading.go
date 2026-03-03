// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CopyTrading is the golang structure of table copy_trading for DAO operations like Where/Data.
type CopyTrading struct {
	g.Meta                         `orm:"table:copy_trading, do:true"`
	Id                             any         // 主键ID
	UserId                         any         // 所属用户ID
	TargetWallet                   any         // 目标钱包地址
	TargetWalletPlatform           any         // 目标钱包平台
	Remark                         any         // 备注
	FollowType                     any         // 跟单类型 1:自动跟单 2:条件跟单 3:实时跟单
	FollowOnce                     any         // 是否只跟一次 0:否 1:是
	PositionConditions             any         // 持仓筛选条件(JSON数组)
	TraderConditions               any         // 交易员筛选条件(JSON数组)
	TagAccountValue                any         // 账户总价值 small/medium/whale
	TagProfitScale                 any         // 盈利规模 small/medium/large
	TagDirection                   any         // 方向偏好 short/neutral/long
	TagTradingRhythm               any         // 交易节奏 longterm/swing/short/scalping
	TagProfitStatus                any         // 盈利状态 steady/volatile/balanced
	TagTradingStyles               []string    // 交易风格(多选) hf_stable/hf_aggressive/lf_stable/lf_aggressive/steady_profit/high_risk_high_return/asymmetric/low_drawdown/volatility
	TraderMetricPeriod             any         // 交易员指标周期 1d/7d/30d/90d/all
	FollowMarginMode               any         // 跟单保证金模式 1:逐仓 2:全仓
	FollowSymbol                   any         // 跟单币种
	Leverage                       any         // 杠杆倍数
	MarginMode                     any         // 保证金模式 1:逐仓 2:全仓
	FollowModel                    any         // 跟单模式 1:固定金额 2:固定比例
	FollowModelValue               any         // 跟单模式值
	MinValue                       any         // 最小下单金额
	MaxValue                       any         // 最大下单金额
	MaxMarginUsage                 any         // 最大保证金使用率
	TpValue                        any         // 止盈比例
	SlValue                        any         // 止损比例
	OptReverseFollowOrder          any         // 反向跟单 0:关 1:开
	OptFollowupDecrease            any         // 跟随减仓 0:关 1:开
	OptFollowupIncrease            any         // 跟随加仓 0:关 1:开
	OptForcedLiquidationProtection any         // 强平保护 0:关 1:开
	OptPositionIncreaseOpening     any         // 加仓开仓 0:关 1:开
	OptSlippageProtection          any         // 滑点保护 0:关 1:开
	SymbolListType                 any         // 交易对列表类型 WHITE:白名单 BLACK:黑名单
	SymbolList                     any         // 交易对列表,逗号分隔
	MainWallet                     any         // 主钱包地址
	MainWalletPlatform             any         // 主钱包平台
	Status                         any         // 状态 0:停用 1:启用
	CreatedAt                      *gtime.Time // 创建时间
	UpdatedAt                      *gtime.Time // 更新时间
}
