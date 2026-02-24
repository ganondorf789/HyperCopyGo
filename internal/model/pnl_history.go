package model

import "demo/internal/model/entity"

// PnlPoint PnL 数据点
type PnlPoint struct {
	Ts int64  `json:"ts"` // 时间戳(ms)
	V  string `json:"v"`  // PnL 值
}

// PnlHistoryItem 扩展 entity，将 PnlList 覆盖为切片类型，ORM 自动做 JSON 转换
type PnlHistoryItem struct {
	entity.PnlHistory
	PnlList []PnlPoint `json:"pnlList" orm:"pnl_list"`
}
