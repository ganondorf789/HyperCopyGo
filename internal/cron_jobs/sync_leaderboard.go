package cron_jobs

import (
	"context"
	"encoding/json"

	"demo/internal/dao"
	"demo/internal/model/do"

	"github.com/gogf/gf/v2/frame/g"
)

const leaderboardAPI = "https://stats-data.hyperliquid.xyz/Mainnet/leaderboard"

type leaderboardResp struct {
	LeaderboardRows []struct {
		EthAddress         string          `json:"ethAddress"`
		AccountValue       json.Number     `json:"accountValue"`
		WindowPerformances []windowPerfRow `json:"windowPerformances"`
	} `json:"leaderboardRows"`
}

// windowPerfRow 是 ["week", {"pnl":"...", "roi":"...", "vlm":"..."}] 形式的元组
type windowPerfRow struct {
	Window string
	Perf   struct {
		Pnl json.Number `json:"pnl"`
		Roi json.Number `json:"roi"`
		Vlm json.Number `json:"vlm"`
	}
}

func (w *windowPerfRow) UnmarshalJSON(data []byte) error {
	var raw [2]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	if err := json.Unmarshal(raw[0], &w.Window); err != nil {
		return err
	}
	return json.Unmarshal(raw[1], &w.Perf)
}

func init() {
	Register("sync_leaderboard", SyncLeaderboard)
}

// SyncLeaderboard 从 Hyperliquid 排行榜拉取全部交易员数据，
// 提取 week 周期的 pnl/roi/vlm 保存到 leaderboard 表。
func SyncLeaderboard(ctx context.Context, _ string) {
	resp, err := g.Client().Get(ctx, leaderboardAPI)
	if err != nil {
		g.Log().Errorf(ctx, "SyncLeaderboard: 请求排行榜 API 失败: %v", err)
		return
	}
	defer resp.Close()

	var result leaderboardResp
	if err = json.Unmarshal(resp.ReadAll(), &result); err != nil {
		g.Log().Errorf(ctx, "SyncLeaderboard: 解析响应失败: %v", err)
		return
	}

	if len(result.LeaderboardRows) == 0 {
		g.Log().Infof(ctx, "SyncLeaderboard: 无排行榜数据")
		return
	}

	synced := 0
	for _, row := range result.LeaderboardRows {
		if row.EthAddress == "" {
			continue
		}

		accountValue, _ := row.AccountValue.Float64()

		var pnl, roi, vlm float64
		for _, wp := range row.WindowPerformances {
			if wp.Window == "week" {
				pnl, _ = wp.Perf.Pnl.Float64()
				roi, _ = wp.Perf.Roi.Float64()
				vlm, _ = wp.Perf.Vlm.Float64()
				break
			}
		}

		data := do.Leaderboard{
			EthAddress:   row.EthAddress,
			AccountValue: accountValue,
			Pnl:          pnl,
			Roi:          roi,
			Vlm:          vlm,
		}

		affected, err := dao.Leaderboard.Ctx(ctx).
			Where(do.Leaderboard{EthAddress: row.EthAddress}).
			Data(data).
			UpdateAndGetAffected()
		if err != nil {
			g.Log().Errorf(ctx, "SyncLeaderboard: 更新 %s 失败: %v", row.EthAddress, err)
			continue
		}
		if affected == 0 {
			if _, err = dao.Leaderboard.Ctx(ctx).Data(data).Insert(); err != nil {
				g.Log().Errorf(ctx, "SyncLeaderboard: 插入 %s 失败: %v", row.EthAddress, err)
				continue
			}
		}
		synced++
	}

	g.Log().Infof(ctx, "SyncLeaderboard: 成功同步 %d 个排行榜交易员", synced)
}
