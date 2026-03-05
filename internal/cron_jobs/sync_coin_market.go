package cron_jobs

import (
	"context"
	"strconv"

	"demo/internal/dao"
	"demo/internal/model/entity"
	proxyPool "demo/internal/proxy_pool"

	"github.com/gogf/gf/v2/frame/g"
	hyperliquid "github.com/sonirico/go-hyperliquid"
)

func init() {
	Register("sync_coin_market", SyncCoinMarket)
}

// SyncCoinMarket 通过 Hyperliquid API 拉取全部永续合约资产的行情数据，
// 更新 coin_market 表中的价格、涨跌幅、成交量、资金费率、未平仓量等字段。
func SyncCoinMarket(ctx context.Context, _ string) {
	info := hyperliquid.NewInfo(ctx, hyperliquid.MainnetAPIURL, true, nil, nil, nil,
		hyperliquid.InfoOptClientOptions(hyperliquid.ClientOptHTTPClient(proxyPool.HTTPClient())),
	)

	result, err := info.MetaAndAssetCtxs(ctx, hyperliquid.MetaAndAssetCtxsParams{})
	if err != nil {
		g.Log().Errorf(ctx, "SyncCoinMarket: 请求 MetaAndAssetCtxs 失败: %v", err)
		return
	}

	universe := result.Universe
	ctxs := result.Ctxs

	if len(universe) != len(ctxs) {
		g.Log().Errorf(ctx, "SyncCoinMarket: Universe(%d) 与 Ctxs(%d) 长度不匹配", len(universe), len(ctxs))
		return
	}

	synced := 0
	for i, asset := range universe {
		if asset.IsDelisted {
			continue
		}

		ac := ctxs[i]
		coin := asset.Name

		markPx := parseFloat(ac.MarkPx)
		prevDayPx := parseFloat(ac.PrevDayPx)
		funding := parseFloat(ac.Funding)
		openInterest := parseFloat(ac.OpenInterest)
		dayNtlVlm := parseFloat(ac.DayNtlVlm)
		dayBaseVlm := parseFloat(ac.DayBaseVlm)

		change24h := markPx - prevDayPx
		changePercent24h := 0.0
		if prevDayPx > 0 {
			changePercent24h = change24h / prevDayPx * 100
		}

		data := entity.CoinMarket{
			Coin:             coin,
			Price:            markPx,
			Change24H:        change24h,
			ChangePercent24H: changePercent24h,
			Open24H:          prevDayPx,
			Close24H:         markPx,
			Volume24H:        dayBaseVlm,
			QuoteVolume24H:   dayNtlVlm,
			Funding:          funding,
			OpenInterest:     openInterest,
		}

		affected, err := dao.CoinMarket.Ctx(ctx).
			Where("coin = ?", coin).
			Data(data).
			UpdateAndGetAffected()
		if err != nil {
			g.Log().Errorf(ctx, "SyncCoinMarket: 更新 %s 失败: %v", coin, err)
			continue
		}

		if affected == 0 {
			if _, err = dao.CoinMarket.Ctx(ctx).Data(data).Insert(); err != nil {
				g.Log().Errorf(ctx, "SyncCoinMarket: 插入 %s 失败: %v", coin, err)
				continue
			}
		}

		synced++
	}

	g.Log().Infof(ctx, "SyncCoinMarket: 成功同步 %d 个币种行情", synced)
}

func parseFloat(s string) float64 {
	if s == "" {
		return 0
	}
	v, _ := strconv.ParseFloat(s, 64)
	return v
}
