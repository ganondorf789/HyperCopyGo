package wallet_grpc

import (
	"context"
	"fmt"

	hyperliquid "github.com/sonirico/go-hyperliquid"

	v1 "demo/api/wallet_grpc/v1"
	"demo/internal/dao"
	"demo/internal/model/entity"
	proxyPool "demo/internal/proxy_pool"
	"demo/internal/service"
	"demo/utility"
)

func init() {
	service.RegisterWalletGrpc(&sWalletGrpc{})
}

type sWalletGrpc struct{}

func (s *sWalletGrpc) GetWalletList(ctx context.Context, appId, appSecret string) (list []*v1.WalletItem, err error) {
	userId, err := utility.ValidateAppKey(ctx, appId, appSecret)
	if err != nil {
		return nil, err
	}

	var items []entity.Wallet
	err = dao.Wallet.Ctx(ctx).
		Where("user_id = ?", userId).
		OrderDesc(dao.Wallet.Columns().Id).
		Scan(&items)
	if err != nil {
		return nil, err
	}

	list = make([]*v1.WalletItem, 0, len(items))
	for _, e := range items {
		list = append(list, entityToProto(ctx, e))
	}
	return list, nil
}

func entityToProto(ctx context.Context, e entity.Wallet) *v1.WalletItem {
	item := &v1.WalletItem{
		Id:               e.Id,
		Address:          e.Address,
		ApiWalletAddress: e.ApiWalletAddress,
		Remark:           e.Remark,
		Balance:          "0.0",
		TotalMarginUsed:  "0.0",
		Withdrawable:     "0.0",
		Upnl:             "0.0000",
		DepositWallet:    e.ApiWalletAddress,
	}
	if e.CreatedAt != nil {
		item.CreatedAt = e.CreatedAt.String()
	}
	if e.UpdatedAt != nil {
		item.UpdatedAt = e.UpdatedAt.String()
	}

	info := hyperliquid.NewInfo(ctx, hyperliquid.MainnetAPIURL, true, nil, nil, nil,
		hyperliquid.InfoOptClientOptions(hyperliquid.ClientOptHTTPClient(proxyPool.HTTPClient())),
	)

	state, err := info.UserState(ctx, e.Address)
	if err == nil {
		item.Balance = state.MarginSummary.AccountValue
		item.TotalMarginUsed = state.MarginSummary.TotalMarginUsed
		item.Withdrawable = state.Withdrawable

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
