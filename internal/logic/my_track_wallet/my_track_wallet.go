package my_track_wallet

import (
	"context"
	"fmt"
	"math"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	hyperliquid "github.com/sonirico/go-hyperliquid"

	v1 "demo/api/my_track_wallet/v1"
	proxyPool "demo/internal/proxy_pool"
	"demo/internal/dao"
	"demo/internal/model"
	"demo/internal/model/entity"
	"demo/internal/service"
	"demo/utility"
)

func init() {
	service.RegisterMyTrackWallet(&sMyTrackWallet{})
}

type sMyTrackWallet struct{}

func (s *sMyTrackWallet) Create(ctx context.Context, userId int64, in v1.MyTrackWalletCreateReq) (res *v1.MyTrackWalletCreateRes, err error) {
	ids := make([]int64, 0, len(in.Records))
	for _, r := range in.Records {
		id, err := dao.MyTrackWallet.Ctx(ctx).Data(g.Map{
			"user_id":       userId,
			"wallet":        r.Wallet,
			"remark":        r.Remark,
			"enable_notify": r.EnableNotify,
			"notify_action": r.NotifyAction,
		}).InsertAndGetId()
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return &v1.MyTrackWalletCreateRes{Ids: ids}, nil
}

func (s *sMyTrackWallet) Update(ctx context.Context, userId int64, in v1.MyTrackWalletUpdateReq) error {
	count, err := dao.MyTrackWallet.Ctx(ctx).
		Where("id = ? AND user_id = ?", in.Id, userId).
		Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("跟踪钱包不存在")
	}

	data := g.Map{
		"remark":        in.Remark,
		"notify_action": in.NotifyAction,
	}
	if in.EnableNotify != nil {
		data["enable_notify"] = *in.EnableNotify
	}
	if in.Lang != "" {
		data["lang"] = in.Lang
	}

	_, err = dao.MyTrackWallet.Ctx(ctx).
		Where("id = ? AND user_id = ?", in.Id, userId).
		Data(data).
		Update()
	return err
}

func (s *sMyTrackWallet) Delete(ctx context.Context, userId int64, id int64) error {
	_, err := dao.MyTrackWallet.Ctx(ctx).
		Where("id = ? AND user_id = ?", id, userId).
		Delete()
	return err
}

func (s *sMyTrackWallet) Detail(ctx context.Context, userId int64, id int64) (res *v1.MyTrackWalletDetailRes, err error) {
	var item entity.MyTrackWallet
	err = dao.MyTrackWallet.Ctx(ctx).
		Where("id = ? AND user_id = ?", id, userId).
		Scan(&item)
	if err != nil {
		return nil, err
	}
	if item.Id == 0 {
		return nil, fmt.Errorf("跟踪钱包不存在")
	}
	walletItem := s.entityToItem(ctx, item)
	return &v1.MyTrackWalletDetailRes{TrackWalletItem: walletItem}, nil
}

func (s *sMyTrackWallet) List(ctx context.Context, userId int64, in v1.MyTrackWalletListReq) (res *v1.MyTrackWalletListRes, err error) {
	m := dao.MyTrackWallet.Ctx(ctx).Where("user_id = ?", userId)

	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	var items []entity.MyTrackWallet
	err = m.Page(in.Page, in.PageSize).
		OrderDesc(dao.MyTrackWallet.Columns().Id).
		Scan(&items)
	if err != nil {
		return nil, err
	}

	list := make([]model.TrackWalletItem, 0, len(items))
	for _, item := range items {
		list = append(list, s.entityToItem(ctx, item))
	}

	return &v1.MyTrackWalletListRes{
		List:  list,
		Total: total,
		Page:  in.Page,
	}, nil
}

// entityToItem 将 DB entity 转为列表项，通过 Hyperliquid API 获取账户余额和持仓
func (s *sMyTrackWallet) entityToItem(ctx context.Context, e entity.MyTrackWallet) model.TrackWalletItem {
	var remark *string
	if e.Remark != "" {
		remark = &e.Remark
	}

	item := model.TrackWalletItem{
		Id:                 e.Id,
		Wallet:             e.Wallet,
		Remark:             remark,
		EnableNotify:       e.EnableNotify,
		NotifyAction:       e.NotifyAction,
		Lang:               e.Lang,
		Balance:            "0.0",
		Pnl:                "0.0000",
		MarginUsedRatio:    "0.0000%",
		TotalPositionValue: "0.0000",
		Positions:          make([]model.TrackWalletPosition, 0),
	}

	info := hyperliquid.NewInfo(ctx, hyperliquid.MainnetAPIURL, true, nil, nil, nil,
		hyperliquid.InfoOptClientOptions(hyperliquid.ClientOptHTTPClient(proxyPool.HTTPClient())),
	)

	state, err := info.UserState(ctx, e.Wallet)
	if err != nil {
		return item
	}

	// 账户余额
	item.Balance = state.MarginSummary.AccountValue

	// 保证金使用率 = totalMarginUsed / accountValue * 100%
	var accountValue, totalMarginUsed float64
	fmt.Sscanf(state.MarginSummary.AccountValue, "%f", &accountValue)
	fmt.Sscanf(state.MarginSummary.TotalMarginUsed, "%f", &totalMarginUsed)
	if accountValue > 0 {
		ratio := totalMarginUsed / accountValue * 100
		item.MarginUsedRatio = fmt.Sprintf("%.4f%%", ratio)
	}

	// 遍历持仓，计算 pnl 合计和总持仓价值
	totalPnl := 0.0
	totalPosValue := 0.0

	for _, ap := range state.AssetPositions {
		pos := ap.Position

		// 方向：szi > 0 为 long，< 0 为 short
		var szi float64
		fmt.Sscanf(pos.Szi, "%f", &szi)
		direction := "long"
		if szi < 0 {
			direction = "short"
		}

		// 保证金类型
		marginType := "cross"
		if pos.Leverage.Type == "isolated" {
			marginType = "isolated"
		}

		// 持仓价值 = |szi| * markPx（从 AssetCtx 获取不到时用 positionValue）
		positionValue := pos.PositionValue

		// 未实现盈亏
		var unrealizedPnl float64
		fmt.Sscanf(pos.UnrealizedPnl, "%f", &unrealizedPnl)
		totalPnl += unrealizedPnl

		// 总持仓价值
		var posVal float64
		fmt.Sscanf(positionValue, "%f", &posVal)
		totalPosValue += math.Abs(posVal)

		// 未实现盈亏比例 = unrealizedPnl / marginUsed * 100%
		var marginUsed float64
		fmt.Sscanf(pos.MarginUsed, "%f", &marginUsed)
		pnlRatio := "0.0000%"
		if marginUsed > 0 {
			pnlRatio = fmt.Sprintf("%.4f%%", unrealizedPnl/marginUsed*100)
		}

		// 开仓均价
		entryPx := ""
		if pos.EntryPx != nil {
			entryPx = *pos.EntryPx
		}

		// 强平价格
		liquidationPx := ""
		if pos.LiquidationPx != nil {
			liquidationPx = *pos.LiquidationPx
		}

		item.Positions = append(item.Positions, model.TrackWalletPosition{
			Coin:               pos.Coin,
			Leverage:           pos.Leverage.Value,
			Direction:          direction,
			Type:               marginType,
			Szi:                pos.Szi,
			PositionValue:      positionValue,
			EntryPx:            entryPx,
			MarkPx:             "", // markPx 需要从 MetaAndAssetCtxs 获取，此处留空
			UnrealizedPnl:      pos.UnrealizedPnl,
			UnrealizedPnlRatio: pnlRatio,
			LiquidationPx:      liquidationPx,
			MarginUsed:         pos.MarginUsed,
		})
	}

	item.Pnl = fmt.Sprintf("%.4f", totalPnl)
	item.TotalPositionValue = fmt.Sprintf("%.4f", totalPosValue)

	return item
}

// ==================== 导出/导入分享码 ====================

// Export 导出当前用户所有跟踪钱包为分享码
func (s *sMyTrackWallet) Export(ctx context.Context, userId int64) (res *v1.MyTrackWalletExportRes, err error) {
	var items []entity.MyTrackWallet
	err = dao.MyTrackWallet.Ctx(ctx).
		Where("user_id = ?", userId).
		Scan(&items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, fmt.Errorf("没有可导出的跟踪钱包")
	}

	wallets := make([]string, 0, len(items))
	for _, item := range items {
		wallets = append(wallets, item.Wallet)
	}

	code, err := utility.EncodeWallets(wallets)
	if err != nil {
		return nil, err
	}
	return &v1.MyTrackWalletExportRes{Code: code}, nil
}

// Import 通过分享码导入跟踪钱包（跳过已存在的地址）
func (s *sMyTrackWallet) Import(ctx context.Context, userId int64, in v1.MyTrackWalletImportReq) (res *v1.MyTrackWalletImportRes, err error) {
	wallets, err := utility.DecodeWallets(in.Code)
	if err != nil {
		return nil, err
	}

	// 查询当前用户已有的钱包地址，避免重复
	var existing []entity.MyTrackWallet
	err = dao.MyTrackWallet.Ctx(ctx).
		Where("user_id = ?", userId).
		Scan(&existing)
	if err != nil {
		return nil, err
	}
	existMap := make(map[string]bool, len(existing))
	for _, e := range existing {
		existMap[strings.ToLower(e.Wallet)] = true
	}

	ids := make([]int64, 0)
	for _, w := range wallets {
		if existMap[strings.ToLower(w)] {
			continue // 跳过已存在的
		}
		id, err := dao.MyTrackWallet.Ctx(ctx).Data(entity.MyTrackWallet{
			UserId: userId,
			Wallet: w,
		}).InsertAndGetId()
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	return &v1.MyTrackWalletImportRes{Ids: ids}, nil
}
