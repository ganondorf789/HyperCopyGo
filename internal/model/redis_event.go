package model

type NewPositionEvent struct {
	Address               string `json:"address"`
	Coin                  string `json:"coin"`
	Szi                   string `json:"szi"`
	LeverageType          string `json:"leverageType"`
	Leverage              int    `json:"leverage"`
	EntryPx               string `json:"entryPx"`
	PositionValue         string `json:"positionValue"`
	UnrealizedPnl         string `json:"unrealizedPnl"`
	ReturnOnEquity        string `json:"returnOnEquity"`
	LiquidationPx         string `json:"liquidationPx"`
	MarginUsed            string `json:"marginUsed"`
	MaxLeverage           int    `json:"maxLeverage"`
	CumFundingAllTime     string `json:"cumFundingAllTime"`
	CumFundingSinceOpen   string `json:"cumFundingSinceOpen"`
	CumFundingSinceChange string `json:"cumFundingSinceChange"`
}

type MarketAlert struct {
	Count     int64 `json:"count"`
	Minutes   int   `json:"minutes"`
	Threshold int   `json:"threshold"`
}

