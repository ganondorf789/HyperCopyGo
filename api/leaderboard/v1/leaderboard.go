package v1

import (
	"demo/internal/model"
	"demo/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// 排行榜（ROI Top10）
type LeaderboardTopReq struct {
	g.Meta `path:"/leaderboard/top" tags:"Leaderboard" method:"get" summary:"排行榜（ROI Top10）"`
}
type LeaderboardTopRes struct {
	g.Meta `mime:"application/json"`
	List   []entity.Leaderboard `json:"list"`
}

// 跌幅榜（ROI Bottom10）
type LeaderboardDeclineReq struct {
	g.Meta `path:"/leaderboard/decline" tags:"Leaderboard" method:"get" summary:"跌幅榜（ROI Bottom10）"`
}
type LeaderboardDeclineRes struct {
	g.Meta `mime:"application/json"`
	List   []entity.Leaderboard `json:"list"`
}

// 交易量列表（Vlm Top10）
type LeaderboardVolumeReq struct {
	g.Meta `path:"/leaderboard/volume" tags:"Leaderboard" method:"get" summary:"交易量列表（Vlm Top10）"`
}
type LeaderboardVolumeRes struct {
	g.Meta `mime:"application/json"`
	List   []entity.Leaderboard `json:"list"`
}

// 热门币种列表
type HotCoinReq struct {
	g.Meta `path:"/leaderboard/hot-coin" tags:"Leaderboard" method:"get" summary:"热门币种列表"`
}
type HotCoinRes struct {
	g.Meta `mime:"application/json"`
	List   []entity.CoinMarket `json:"list"`
}

// 盈利榜（按 PnL 降序，支持分页和时间窗口）
type LeaderboardProfitReq struct {
	g.Meta   `path:"/leaderboard/profit" tags:"Leaderboard" method:"get" summary:"盈利榜"`
	Window   string `json:"window"   v:"required|in:day,week,month,allTime#请选择时间窗口|时间窗口只能是day/week/month/allTime"`
	Page     int    `json:"page"     d:"1"`
	PageSize int    `json:"pageSize" d:"20" v:"max:100#每页最多100条"`
}
type LeaderboardProfitRes struct {
	g.Meta `mime:"application/json"`
	List   []model.LeaderboardProfitItem `json:"list"`
	Total  int                           `json:"total"`
	Page   int                           `json:"page"`
}
