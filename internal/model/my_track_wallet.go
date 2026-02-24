package model

// TrackWalletPosition 跟踪钱包持仓信息
type TrackWalletPosition struct {
	Coin              string `json:"coin"`              // 币种
	Leverage          int    `json:"leverage"`          // 杠杆倍数
	Direction         string `json:"direction"`         // 方向 long/short
	Type              string `json:"type"`              // 保证金类型 cross/isolated
	Szi               string `json:"szi"`               // 持仓数量
	PositionValue     string `json:"positionValue"`     // 持仓价值
	EntryPx           string `json:"entryPx"`           // 开仓均价
	MarkPx            string `json:"markPx"`            // 标记价格
	UnrealizedPnl     string `json:"unrealizedPnl"`     // 未实现盈亏
	UnrealizedPnlRatio string `json:"unrealizedPnlRatio"` // 未实现盈亏比例
	LiquidationPx     string `json:"liquidationPx"`     // 强平价格
	MarginUsed        string `json:"marginUsed"`        // 已用保证金
}

// TrackWalletItem 跟踪钱包列表项（含 Hyperliquid 链上数据）
type TrackWalletItem struct {
	Id                 int64                 `json:"id"`
	Wallet             string                `json:"wallet"`             // 钱包地址
	Remark             *string               `json:"remark"`             // 备注
	EnableNotify       int                   `json:"enableNotify"`       // 是否开启通知
	NotifyAction       string                `json:"notifyAction"`       // 通知动作
	Lang               string                `json:"lang"`               // 语言
	Balance            string                `json:"balance"`            // 账户余额
	Pnl                string                `json:"pnl"`                // 未实现盈亏合计
	MarginUsedRatio    string                `json:"marginUsedRatio"`    // 保证金使用率
	TotalPositionValue string                `json:"totalPositionValue"` // 总持仓价值
	Positions          []TrackWalletPosition `json:"positions"`          // 持仓列表
}
