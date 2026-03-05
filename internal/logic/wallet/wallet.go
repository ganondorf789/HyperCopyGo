package wallet

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	hyperliquid "github.com/sonirico/go-hyperliquid"

	v1 "demo/api/wallet/v1"
	proxyPool "demo/internal/proxy_pool"
	"demo/internal/dao"
	"demo/internal/model"
	"demo/internal/model/entity"
	"demo/internal/service"
)

func init() {
	service.RegisterWallet(&sWallet{})
}

type sWallet struct{}

func (s *sWallet) Create(ctx context.Context, userId int64, in v1.WalletCreateReq) (res *v1.WalletCreateRes, err error) {
	var existing entity.Wallet
	err = dao.Wallet.Ctx(ctx).
		Where("user_id = ? AND address = ?", userId, in.Address).
		Scan(&existing)
	if err != nil {
		return nil, err
	}
	if existing.Id != 0 {
		return nil, fmt.Errorf("钱包地址已存在")
	}

	id, err := dao.Wallet.Ctx(ctx).Data(g.Map{
		"user_id":            userId,
		"address":            in.Address,
		"api_wallet_address": in.ApiWalletAddress,
		"api_secret_key":     in.ApiSecretKey,
		"remark":             in.Remark,
	}).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return &v1.WalletCreateRes{Id: id}, nil
}

func (s *sWallet) Update(ctx context.Context, userId int64, in v1.WalletUpdateReq) error {
	count, err := dao.Wallet.Ctx(ctx).
		Where("id = ? AND user_id = ?", in.Id, userId).
		Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("钱包不存在")
	}
	// 只允许编辑 Remark
	_, err = dao.Wallet.Ctx(ctx).
		Where("id = ? AND user_id = ?", in.Id, userId).
		Data(g.Map{"remark": in.Remark}).
		Update()
	return err
}

func (s *sWallet) Delete(ctx context.Context, userId int64, id int64) error {
	_, err := dao.Wallet.Ctx(ctx).
		Where("id = ? AND user_id = ?", id, userId).
		Delete()
	return err
}

func (s *sWallet) Detail(ctx context.Context, userId int64, id int64) (res *v1.WalletDetailRes, err error) {
	var item entity.Wallet
	err = dao.Wallet.Ctx(ctx).
		Where("id = ? AND user_id = ?", id, userId).
		Scan(&item)
	if err != nil {
		return nil, err
	}
	if item.Id == 0 {
		return nil, fmt.Errorf("钱包不存在")
	}
	walletItem := s.entityToItem(ctx, item)
	return &v1.WalletDetailRes{WalletItem: walletItem}, nil
}

func (s *sWallet) List(ctx context.Context, userId int64, in v1.WalletListReq) (res *v1.WalletListRes, err error) {
	m := dao.Wallet.Ctx(ctx).Where("user_id = ?", userId)

	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	var items []entity.Wallet
	err = m.Page(in.Page, in.PageSize).
		OrderDesc(dao.Wallet.Columns().Id).
		Scan(&items)
	if err != nil {
		return nil, err
	}

	list := make([]model.WalletItem, 0, len(items))
	for _, item := range items {
		list = append(list, s.entityToItem(ctx, item))
	}

	return &v1.WalletListRes{
		List:  list,
		Total: total,
		Page:  in.Page,
	}, nil
}

// entityToItem 将 DB entity 转为 API 返回项，并从 Hyperliquid 拉取链上数据
func (s *sWallet) entityToItem(ctx context.Context, e entity.Wallet) model.WalletItem {
	createdAt := ""
	if e.CreatedAt != nil {
		createdAt = e.CreatedAt.String()
	}
	updatedAt := ""
	if e.UpdatedAt != nil {
		updatedAt = e.UpdatedAt.String()
	}
	item := model.WalletItem{
		Id:               e.Id,
		Address:          e.Address,
		ApiWalletAddress: e.ApiWalletAddress,
		Remark:           e.Remark,
		Balance:          "0.0",
		TotalMarginUsed:  "0.0",
		Withdrawable:     "0.0",
		Upnl:             "0.0000",
		DepositWallet:    e.ApiWalletAddress,
		ArbWithdrawAble:  nil,
		BscWithdrawAble:  nil,
		CreatedAt:        createdAt,
		UpdatedAt:        updatedAt,
	}

	info := hyperliquid.NewInfo(ctx, hyperliquid.MainnetAPIURL, true, nil, nil, nil,
		hyperliquid.InfoOptClientOptions(hyperliquid.ClientOptHTTPClient(proxyPool.HTTPClient())),
	)

	// 1. 获取永续合约账户状态（余额、保证金、可提现、持仓）
	state, err := info.UserState(ctx, e.Address)
	if err == nil {
		item.Balance = state.MarginSummary.AccountValue
		item.TotalMarginUsed = state.MarginSummary.TotalMarginUsed
		item.Withdrawable = state.Withdrawable

		// 汇总所有持仓的未实现盈亏
		if len(state.AssetPositions) > 0 {
			totalUpnl := 0.0
			for _, ap := range state.AssetPositions {
				var pnl float64
				fmt.Sscanf(ap.Position.UnrealizedPnl, "%f", &pnl)
				totalUpnl += pnl
			}
			item.Upnl = fmt.Sprintf("%.4f", totalUpnl)
		}
	}

	// 2. 获取现货账户状态，将 USDC 余额累加到 balance
	spotState, err := info.SpotUserState(ctx, e.Address)
	if err == nil {
		for _, b := range spotState.Balances {
			if b.Coin == "USDC" {
				var spotTotal float64
				fmt.Sscanf(b.Total, "%f", &spotTotal)
				if spotTotal > 0 {
					var currentBalance float64
					fmt.Sscanf(item.Balance, "%f", &currentBalance)
					item.Balance = fmt.Sprintf("%.4f", currentBalance+spotTotal)
				}
				break
			}
		}
	}

	return item
}
