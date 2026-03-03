package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// WsConnectReq WebSocket 连接请求，通过 query 参数传递 JWT token
type WsConnectReq struct {
	g.Meta `path:"/ws" method:"get" tags:"WebSocket" summary:"WebSocket连接"`
	Token  string `json:"token" in:"query" v:"required#请传入token"`
}

type WsConnectRes struct {
	g.Meta `mime:"text/html" type:"string" example:" "`
}
