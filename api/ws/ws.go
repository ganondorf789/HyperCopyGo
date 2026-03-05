// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package ws

import (
	"context"

	"demo/api/ws/v1"
)

type IWsV1 interface {
	WsConnect(ctx context.Context, req *v1.WsConnectReq) (res *v1.WsConnectRes, err error)
}
