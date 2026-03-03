// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package wallet

import (
	"context"

	"demo/api/wallet/v1"
)

type IWalletV1 interface {
	WalletCreate(ctx context.Context, req *v1.WalletCreateReq) (res *v1.WalletCreateRes, err error)
	WalletUpdate(ctx context.Context, req *v1.WalletUpdateReq) (res *v1.WalletUpdateRes, err error)
	WalletDelete(ctx context.Context, req *v1.WalletDeleteReq) (res *v1.WalletDeleteRes, err error)
	WalletDetail(ctx context.Context, req *v1.WalletDetailReq) (res *v1.WalletDetailRes, err error)
	WalletList(ctx context.Context, req *v1.WalletListReq) (res *v1.WalletListRes, err error)
}
