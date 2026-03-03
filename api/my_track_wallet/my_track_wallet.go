// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package my_track_wallet

import (
	"context"

	"demo/api/my_track_wallet/v1"
)

type IMyTrackWalletV1 interface {
	MyTrackWalletCreate(ctx context.Context, req *v1.MyTrackWalletCreateReq) (res *v1.MyTrackWalletCreateRes, err error)
	MyTrackWalletUpdate(ctx context.Context, req *v1.MyTrackWalletUpdateReq) (res *v1.MyTrackWalletUpdateRes, err error)
	MyTrackWalletDelete(ctx context.Context, req *v1.MyTrackWalletDeleteReq) (res *v1.MyTrackWalletDeleteRes, err error)
	MyTrackWalletDetail(ctx context.Context, req *v1.MyTrackWalletDetailReq) (res *v1.MyTrackWalletDetailRes, err error)
	MyTrackWalletList(ctx context.Context, req *v1.MyTrackWalletListReq) (res *v1.MyTrackWalletListRes, err error)
	MyTrackWalletExport(ctx context.Context, req *v1.MyTrackWalletExportReq) (res *v1.MyTrackWalletExportRes, err error)
	MyTrackWalletImport(ctx context.Context, req *v1.MyTrackWalletImportReq) (res *v1.MyTrackWalletImportRes, err error)
}
