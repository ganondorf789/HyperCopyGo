package cron_jobs

import (
	"context"
	"math"
	"sync"

	"demo/internal/dao"
	"demo/internal/model/entity"
	proxyPool "demo/internal/proxy_pool"

	"github.com/gogf/gf/v2/frame/g"
	hyperliquid "github.com/sonirico/go-hyperliquid"
)

const l2DepthWorkers = 10

func init() {
	Register("sync_whale_anchor", SyncWhaleAnchor)
}

type coinBaseData struct {
	coin            string
	markPrice       float64
	dayNtlVlm      float64
	openInterestUSD float64
}

// SyncWhaleAnchor 计算所有币种的巨鲸仓位锚点。
// 公式：whale_threshold = max(0.4%×24hVolume, 1%×OI, 30%×1%盘口深度)
func SyncWhaleAnchor(ctx context.Context, _ string) {
	info := hyperliquid.NewInfo(ctx, hyperliquid.MainnetAPIURL, true, nil, nil, nil,
		hyperliquid.InfoOptClientOptions(hyperliquid.ClientOptHTTPClient(proxyPool.HTTPClient())),
	)

	result, err := info.MetaAndAssetCtxs(ctx, hyperliquid.MetaAndAssetCtxsParams{})
	if err != nil {
		g.Log().Errorf(ctx, "SyncWhaleAnchor: 请求 MetaAndAssetCtxs 失败: %v", err)
		return
	}

	universe := result.Universe
	ctxs := result.Ctxs
	if len(universe) != len(ctxs) {
		g.Log().Errorf(ctx, "SyncWhaleAnchor: Universe(%d) 与 Ctxs(%d) 长度不匹配", len(universe), len(ctxs))
		return
	}

	// 1. 构建基础数据
	var coins []coinBaseData
	for i, asset := range universe {
		if asset.IsDelisted {
			continue
		}
		ac := ctxs[i]
		markPx := parseFloat(ac.MarkPx)
		if markPx <= 0 {
			continue
		}
		oiBase := parseFloat(ac.OpenInterest)
		coins = append(coins, coinBaseData{
			coin:            asset.Name,
			markPrice:       markPx,
			dayNtlVlm:      parseFloat(ac.DayNtlVlm),
			openInterestUSD: oiBase * markPx,
		})
	}

	if len(coins) == 0 {
		g.Log().Infof(ctx, "SyncWhaleAnchor: 无可用币种数据")
		return
	}

	// 2. 并发获取 L2 盘口深度
	type depthResult struct {
		idx      int
		depth1pc float64
	}

	depthCh := make(chan depthResult, len(coins))
	sem := make(chan struct{}, l2DepthWorkers)
	var wg sync.WaitGroup

	for i, c := range coins {
		wg.Add(1)
		go func(idx int, coin string, mid float64) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			depth := calcDepth1Pct(ctx, info, coin, mid)
			depthCh <- depthResult{idx: idx, depth1pc: depth}
		}(i, c.coin, c.markPrice)
	}

	wg.Wait()
	close(depthCh)

	depths := make([]float64, len(coins))
	for r := range depthCh {
		depths[r.idx] = r.depth1pc
	}

	// 3. 计算巨鲸锚点并写入数据库
	synced := 0
	for i, c := range coins {
		volComp := 0.004 * c.dayNtlVlm
		oiComp := 0.01 * c.openInterestUSD
		depthComp := 0.30 * depths[i]
		threshold := math.Max(volComp, math.Max(oiComp, depthComp))

		data := entity.WhaleAnchor{
			Symbol:         c.coin,
			Volume24H:      c.dayNtlVlm,
			OpenInterest:   c.openInterestUSD,
			Depth1Pct:      depths[i],
			ValVolume:      volComp,
			ValOi:          oiComp,
			ValDepth:       depthComp,
			WhaleThreshold: threshold,
		}

		affected, err := dao.WhaleAnchor.Ctx(ctx).
			Where("symbol = ?", c.coin).
			Data(data).
			UpdateAndGetAffected()
		if err != nil {
			g.Log().Errorf(ctx, "SyncWhaleAnchor: 更新 %s 失败: %v", c.coin, err)
			continue
		}
		if affected == 0 {
			if _, err = dao.WhaleAnchor.Ctx(ctx).Data(data).Insert(); err != nil {
				g.Log().Errorf(ctx, "SyncWhaleAnchor: 插入 %s 失败: %v", c.coin, err)
				continue
			}
		}
		synced++
	}

	g.Log().Infof(ctx, "SyncWhaleAnchor: 成功同步 %d 个币种巨鲸锚点", synced)
}

// calcDepth1Pct 计算指定币种 1% 盘口深度（USD）。
// 将 mid 上下 1% 范围内的挂单量 × 价格求和（买 + 卖两侧）。
func calcDepth1Pct(ctx context.Context, info *hyperliquid.Info, coin string, mid float64) float64 {
	book, err := info.L2Snapshot(ctx, coin)
	if err != nil {
		g.Log().Warningf(ctx, "SyncWhaleAnchor: 获取 %s L2 快照失败: %v", coin, err)
		return 0
	}
	if len(book.Levels) < 2 {
		return 0
	}

	lowerBound := mid * 0.99
	upperBound := mid * 1.01

	var depth float64

	// bids: book.Levels[0]
	for _, lvl := range book.Levels[0] {
		if lvl.Px >= lowerBound {
			depth += lvl.Px * lvl.Sz
		}
	}
	// asks: book.Levels[1]
	for _, lvl := range book.Levels[1] {
		if lvl.Px <= upperBound {
			depth += lvl.Px * lvl.Sz
		}
	}

	return depth
}
