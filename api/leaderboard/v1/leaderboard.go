package v1

import (
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
