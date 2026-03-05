package model

// PnlPoint PnL 数据点
type PnlPoint struct {
	Ts int64  `json:"ts"` // 时间戳(ms)
	V  string `json:"v"`  // PnL 值
}
