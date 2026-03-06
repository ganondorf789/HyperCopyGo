package cron_jobs

import (
	"context"
	"encoding/json"

	"demo/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
)

const hotCoinAPI = "https://hyperbot.network/api/leaderboard/coin/most-popular"

type hotCoinResp struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data []hotCoinRow `json:"data"`
}

type hotCoinRow struct {
	Coin string `json:"coin"`
}

func init() {
	Register("sync_hot_coin", SyncHotCoin)
}

// SyncHotCoin 从 hyperbot 拉取热门币种列表，先清空旧数据再批量写入。
func SyncHotCoin(ctx context.Context, _ string) {
	resp, err := g.Client().Get(ctx, hotCoinAPI)
	if err != nil {
		g.Log().Errorf(ctx, "SyncHotCoin: 请求热门币种 API 失败: %v", err)
		return
	}
	defer resp.Close()

	var result hotCoinResp
	if err = json.Unmarshal(resp.ReadAll(), &result); err != nil {
		g.Log().Errorf(ctx, "SyncHotCoin: 解析响应失败: %v", err)
		return
	}

	if result.Code != 0 {
		g.Log().Errorf(ctx, "SyncHotCoin: API 返回错误码 %d: %s", result.Code, result.Msg)
		return
	}

	if len(result.Data) == 0 {
		g.Log().Infof(ctx, "SyncHotCoin: 无热门币种数据，跳过本次同步")
		return
	}

	if _, err = dao.HotCoin.Ctx(ctx).WhereGTE("id", 1).Delete(); err != nil {
		g.Log().Errorf(ctx, "SyncHotCoin: 清空旧数据失败: %v", err)
		return
	}

	rows := make([]g.Map, 0, len(result.Data))
	for _, item := range result.Data {
		if item.Coin == "" {
			continue
		}
		rows = append(rows, g.Map{"coin": item.Coin})
	}

	if _, err = dao.HotCoin.Ctx(ctx).Data(rows).Insert(); err != nil {
		g.Log().Errorf(ctx, "SyncHotCoin: 批量写入热门币种失败: %v", err)
		return
	}

	g.Log().Infof(ctx, "SyncHotCoin: 成功同步 %d 个热门币种", len(rows))
}
