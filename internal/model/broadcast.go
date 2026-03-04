package model

type NewPositionBroadcast struct {
	Id             int64  `json:"id"`
	Address        string `json:"address"`
	Coin           string `json:"coin"`
	Szi            string `json:"szi"`
	LeverageType   string `json:"leverageType"`
	Leverage       int    `json:"leverage"`
	EntryPx        string `json:"entryPx"`
	PositionValue  string `json:"positionValue"`
	UnrealizedPnl  string `json:"unrealizedPnl"`
	ReturnOnEquity string `json:"returnOnEquity"`
	LiquidationPx  string `json:"liquidationPx"`
	MarginUsed     string `json:"marginUsed"`
	MaxLeverage    int    `json:"maxLeverage"`
}

type MarketAlertBroadcast struct {
	Id       int64  `json:"id"`
	Category string `json:"category"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Level    int    `json:"level"`
}
