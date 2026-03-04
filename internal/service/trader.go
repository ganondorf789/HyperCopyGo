// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "demo/api/trader/v1"
)

type (
	ITrader interface {
		// Popular 获取热门地址列表：IsHotAddress 的 trader LEFT JOIN month 窗口统计
		Popular(ctx context.Context) (res *v1.TraderPopularRes, err error)
		// KolList 获取 X KOL 列表，支持分页和窗口筛选
		KolList(ctx context.Context, in v1.TraderKolListReq) (res *v1.TraderKolListRes, err error)
	}
)

var (
	localTrader ITrader
)

func Trader() ITrader {
	if localTrader == nil {
		panic("implement not found for interface ITrader, forgot register?")
	}
	return localTrader
}

func RegisterTrader(i ITrader) {
	localTrader = i
}
