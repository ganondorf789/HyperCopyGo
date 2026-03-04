package cron_jobs

import (
	"context"
	"encoding/json"

	"demo/internal/dao"
	"demo/internal/model/do"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
)

const hotTradersAPI = "https://hyperbot.network/api/leaderboard/smart/hot?lang=zh&pnlList=false"

type hotTradersResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Address   string   `json:"address"`
		UserPhoto *string  `json:"userPhoto"`
		Remark    string   `json:"remark"`
		Labels    []string `json:"labels"`
	} `json:"data"`
}

func init() {
	registerJob("sync_hot_traders", "0 */30 * * * *", SyncHotTraders)
}

// SyncHotTraders 从排行榜 API 拉取热门交易员列表，
// 根据 address 更新 remark(→twitter_name)、userPhoto(→profile_picture)、labels，
// 并将其标记为热门地址。
func SyncHotTraders(ctx context.Context) {
	resp, err := g.Client().Get(ctx, hotTradersAPI)
	if err != nil {
		g.Log().Errorf(ctx, "SyncHotTraders: 请求排行榜 API 失败: %v", err)
		return
	}
	defer resp.Close()

	var result hotTradersResp
	if err = json.Unmarshal(resp.ReadAll(), &result); err != nil {
		g.Log().Errorf(ctx, "SyncHotTraders: 解析响应失败: %v", err)
		return
	}
	if result.Code != 0 {
		g.Log().Errorf(ctx, "SyncHotTraders: API 返回错误: code=%d, msg=%s", result.Code, result.Msg)
		return
	}

	if len(result.Data) == 0 {
		g.Log().Infof(ctx, "SyncHotTraders: 无热门交易员数据")
		return
	}

	// 先将所有 trader 的热门标记重置，再重新标记
	_, err = dao.Traders.Ctx(ctx).
		Where(do.Traders{IsHotAddress: true}).
		Data(do.Traders{IsHotAddress: false}).
		Update()
	if err != nil {
		g.Log().Errorf(ctx, "SyncHotTraders: 重置热门标记失败: %v", err)
		return
	}

	for _, item := range result.Data {
		if item.Address == "" {
			continue
		}

		updateData := do.Traders{
			TwitterName:  item.Remark,
			IsHotAddress: true,
			Labels:       item.Labels,
		}
		if item.UserPhoto != nil {
			updateData.ProfilePicture = *item.UserPhoto
		}

		_, err = dao.Traders.Ctx(ctx).
			Where(do.Traders{Address: item.Address}).
			Data(updateData).
			Update()
		if err != nil {
			g.Log().Errorf(ctx, "SyncHotTraders: 更新 trader(%s) 失败: %v", item.Address, err)
			continue
		}
	}

	g.Log().Infof(ctx, "SyncHotTraders: 成功同步 %d 个热门交易员", len(result.Data))
}

type cronJob struct {
	Name     string
	CronExpr string
	Fn       func(ctx context.Context)
}

var jobs []cronJob

func registerJob(name, cronExpr string, fn func(ctx context.Context)) {
	jobs = append(jobs, cronJob{Name: name, CronExpr: cronExpr, Fn: fn})
}

// StartAll 启动所有已注册的定时任务
func StartAll(ctx context.Context) {
	for _, job := range jobs {
		j := job
		if _, err := gcron.Add(ctx, j.CronExpr, func(ctx context.Context) {
			j.Fn(ctx)
		}, j.Name); err != nil {
			g.Log().Errorf(ctx, "注册定时任务 [%s] 失败: %v", j.Name, err)
		} else {
			g.Log().Infof(ctx, "定时任务 [%s] 已注册, cron=%s", j.Name, j.CronExpr)
		}
	}
}
