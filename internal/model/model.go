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

// ==================== Discover ====================

type DiscoverSort struct {
	Field string `json:"field" v:"required#请指定排序字段"`
	Dir   string `json:"dir" d:"DESC" v:"in:ASC,DESC#排序方向仅支持ASC/DESC"`
}

type DiscoverFilter struct {
	Field string `json:"field" v:"required#请指定筛选字段"`
	Op    string `json:"op" v:"required|in:<,=,>,>=,<=,!=,<>,exist#请指定操作符|操作符仅支持< = > >= <= != <> exist"`
	Val   string `json:"val" v:"required#请指定筛选值"`
}

type DiscoverTraderItem struct {
	Address                string     `json:"address"`
	AvgLeverage            float64    `json:"avgLeverage"`
	DdDrawdown             float64    `json:"ddDrawdown"`
	LongPnl                float64    `json:"longPnl"`
	LongWinRate            float64    `json:"longWinRate"`
	PnlList                []PnlPoint `json:"pnlList"`
	Sharpe                 float64    `json:"sharpe"`
	ShortPnl               float64    `json:"shortPnl"`
	ShortWinRate           float64    `json:"shortWinRate"`
	SnapEffLeverage        float64    `json:"snapEffLeverage"`
	SnapLongPositionCount  int64      `json:"snapLongPositionCount"`
	SnapLongPositionValue  float64    `json:"snapLongPositionValue"`
	SnapMarginUsageRate    float64    `json:"snapMarginUsageRate"`
	SnapPerpValue          float64    `json:"snapPerpValue"`
	SnapPositionCount      int64      `json:"snapPositionCount"`
	SnapPositionValue      float64    `json:"snapPositionValue"`
	SnapShortPositionCount int64      `json:"snapShortPositionCount"`
	SnapShortPositionValue float64    `json:"snapShortPositionValue"`
	SnapSpotValue          float64    `json:"snapSpotValue"`
	SnapTotalMarginUsed    float64    `json:"snapTotalMarginUsed"`
	SnapTotalValue         float64    `json:"snapTotalValue"`
	SnapUnrealizedPnl      float64    `json:"snapUnrealizedPnl"`
	Tags                   []string   `json:"tags"`
	TotalPnl               float64    `json:"totalPnl"`
	WinRate                float64    `json:"winRate"`
}

// ==================== 管理员 ====================
