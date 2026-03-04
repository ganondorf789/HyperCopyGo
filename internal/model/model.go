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

type KolTraderItem struct {
	TwitterName       string   `json:"twitterName"`
	Username          string   `json:"username"`
	Address           string   `json:"address"`
	AccountTotalValue float64  `json:"accountTotalValue"`
	WinRate           float64  `json:"winRate"`
	PositionCount     float64  `json:"positionCount"`
	TotalPnl          float64  `json:"totalPnl"`
	ProfilePicture    string   `json:"profilePicture"`
	Labels            []string `json:"labels"`
}

// ==================== 管理员 ====================
