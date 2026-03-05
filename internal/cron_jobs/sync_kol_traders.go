package cron_jobs

import (
	"context"
	"encoding/json"
	"fmt"

	"demo/internal/dao"
	"demo/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

const (
	kolTradersAPI  = "https://hyperbot.network/api/leaderboard/kol"
	kolPageSize    = 50
	kolDefaultLang = "zh"
	kolPeriod      = 7
)

type kolTradersResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Total int `json:"total"`
		List  []struct {
			TwitterName    string   `json:"twitterName"`
			Username       string   `json:"username"`
			Address        string   `json:"address"`
			ProfilePicture string   `json:"profilePicture"`
			Labels         []string `json:"labels"`
		} `json:"list"`
		Pages int `json:"pages"`
	} `json:"data"`
}

func init() {
	Register("sync_kol_traders", SyncKolTraders)
}

// SyncKolTraders 从 KOL 排行榜 API 拉取全部推特 KOL 列表（自动翻页），
// 根据 address 更新 twitterName、username、profilePicture、labels，
// 并将其标记为推特 KOL。
func SyncKolTraders(ctx context.Context, _ string) {
	_, err := dao.Traders.Ctx(ctx).
		Where(entity.Traders{IsTwitterKol: true}).
		Data(g.Map{"is_twitter_kol": false}).
		Update()
	if err != nil {
		g.Log().Errorf(ctx, "SyncKolTraders: 重置 KOL 标记失败: %v", err)
		return
	}

	totalSynced := 0
	for page := 1; ; page++ {
		url := fmt.Sprintf("%s?pageNum=%d&pageSize=%d&period=%d&lang=%s",
			kolTradersAPI, page, kolPageSize, kolPeriod, kolDefaultLang)

		resp, err := g.Client().Get(ctx, url)
		if err != nil {
			g.Log().Errorf(ctx, "SyncKolTraders: 请求 KOL API 第%d页失败: %v", page, err)
			return
		}

		var result kolTradersResp
		if err = json.Unmarshal(resp.ReadAll(), &result); err != nil {
			resp.Close()
			g.Log().Errorf(ctx, "SyncKolTraders: 解析第%d页响应失败: %v", page, err)
			return
		}
		resp.Close()

		if result.Code != 0 {
			g.Log().Errorf(ctx, "SyncKolTraders: API 返回错误: code=%d, msg=%s", result.Code, result.Msg)
			return
		}

		if len(result.Data.List) == 0 {
			break
		}

		for _, item := range result.Data.List {
			if item.Address == "" {
				continue
			}

			_, err = dao.Traders.Ctx(ctx).
				Where(entity.Traders{Address: item.Address}).
				Data(entity.Traders{
					TwitterName:    item.TwitterName,
					Username:       item.Username,
					ProfilePicture: item.ProfilePicture,
					Labels:         item.Labels,
					IsTwitterKol:   true,
				}).
				OmitEmpty().
				Update()
			if err != nil {
				g.Log().Errorf(ctx, "SyncKolTraders: 更新 trader(%s) 失败: %v", item.Address, err)
				continue
			}
		}

		totalSynced += len(result.Data.List)

		if page >= result.Data.Pages {
			break
		}
	}

	g.Log().Infof(ctx, "SyncKolTraders: 成功同步 %d 个 KOL 交易员", totalSynced)
}
