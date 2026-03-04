package model

// ==================== 配置 ====================

type Config struct {
	Jwt      JwtConfig
	SendGrid SendGridConfig
}

type SendGridConfig struct {
	ApiKey    string `json:"apiKey"`
	FromEmail string `json:"fromEmail"`
	FromName  string `json:"fromName"`
}

type JwtConfig struct {
	Secret string `json:"secret"`
	Expire int64  `json:"expire"`
}

// ==================== Trader ====================

type PopularTraderItem struct {
	Address           string   `json:"address"`
	UserPhoto         string   `json:"userPhoto"`
	WinRate           float64  `json:"winRate"`
	RealizedPnl       float64  `json:"realizedPnl"`
	AccountTotalValue float64  `json:"accountTotalValue"`
	CurrentPosition   float64  `json:"currentPosition"`
	Labels            []string `json:"labels"`
	Remark            string   `json:"remark"`
}

// ==================== 管理员 ====================
