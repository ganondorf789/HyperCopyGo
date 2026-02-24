package v1

import (
	"demo/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// 批量创建跟踪钱包（需登录）
type MyTrackWalletCreateReq struct {
	g.Meta  `path:"/my-track-wallet" tags:"MyTrackWallet" method:"post" summary:"批量创建跟踪钱包" login_required:"true"`
	Records []MyTrackWalletRecord `json:"records" v:"required#请输入跟踪钱包记录"`
}
type MyTrackWalletRecord struct {
	Wallet       string `json:"wallet"       v:"required#请输入钱包地址"`
	Remark       string `json:"remark"`
	EnableNotify int    `json:"enableNotify" d:"0"`
	NotifyAction string `json:"notifyAction"`
}
type MyTrackWalletCreateRes struct {
	g.Meta `mime:"application/json"`
	Ids    []int64 `json:"ids"`
}

// 更新跟踪钱包（需登录）
type MyTrackWalletUpdateReq struct {
	g.Meta       `path:"/my-track-wallet/{id}" tags:"MyTrackWallet" method:"put" summary:"更新跟踪钱包" login_required:"true"`
	Id           int64  `json:"id"           in:"path" v:"required"`
	Remark       string `json:"remark"`
	EnableNotify *int   `json:"enableNotify"`
	NotifyAction string `json:"notifyAction"`
	Lang         string `json:"lang"`
}
type MyTrackWalletUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// 删除跟踪钱包（需登录）
type MyTrackWalletDeleteReq struct {
	g.Meta `path:"/my-track-wallet/{id}" tags:"MyTrackWallet" method:"delete" summary:"删除跟踪钱包" login_required:"true"`
	Id     int64 `json:"id" in:"path" v:"required"`
}
type MyTrackWalletDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// 获取跟踪钱包详情（需登录）
type MyTrackWalletDetailReq struct {
	g.Meta `path:"/my-track-wallet/{id}" tags:"MyTrackWallet" method:"get" summary:"跟踪钱包详情" login_required:"true"`
	Id     int64 `json:"id" in:"path" v:"required"`
}
type MyTrackWalletDetailRes struct {
	g.Meta `mime:"application/json"`
	model.TrackWalletItem
}

// 获取跟踪钱包列表（需登录）
type MyTrackWalletListReq struct {
	g.Meta   `path:"/my-track-wallet" tags:"MyTrackWallet" method:"get" summary:"跟踪钱包列表" login_required:"true"`
	Page     int `json:"page" d:"1"`
	PageSize int `json:"pageSize" d:"20" v:"max:100#每页最多100条"`
}
type MyTrackWalletListRes struct {
	g.Meta `mime:"application/json"`
	List   []model.TrackWalletItem `json:"list"`
	Total  int                     `json:"total"`
	Page   int                     `json:"page"`
}
