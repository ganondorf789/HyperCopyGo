package v1

import (
	"demo/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// 创建钱包（需登录）
type WalletCreateReq struct {
	g.Meta           `path:"/wallet" tags:"Wallet" method:"post" summary:"创建钱包" login_required:"true"`
	Address          string `json:"address"          v:"required#请输入钱包地址"`
	ApiWalletAddress string `json:"apiWalletAddress" v:"required#请输入API Wallet Address"`
	ApiSecretKey     string `json:"apiSecretKey"     v:"required#请输入API Secret Key"`
	Remark           string `json:"remark"`
}
type WalletCreateRes struct {
	g.Meta `mime:"application/json"`
	Id     int64 `json:"id"`
}

// 更新钱包备注（需登录，只能编辑Remark）
type WalletUpdateReq struct {
	g.Meta `path:"/wallet/{id}" tags:"Wallet" method:"put" summary:"更新钱包备注" login_required:"true"`
	Id     int64  `json:"id"     in:"path" v:"required"`
	Remark string `json:"remark"`
}
type WalletUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// 删除钱包（需登录）
type WalletDeleteReq struct {
	g.Meta `path:"/wallet/{id}" tags:"Wallet" method:"delete" summary:"删除钱包" login_required:"true"`
	Id     int64 `json:"id" in:"path" v:"required"`
}
type WalletDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// 获取钱包详情（需登录）
type WalletDetailReq struct {
	g.Meta `path:"/wallet/{id}" tags:"Wallet" method:"get" summary:"钱包详情" login_required:"true"`
	Id     int64 `json:"id" in:"path" v:"required"`
}
type WalletDetailRes struct {
	g.Meta `mime:"application/json"`
	model.WalletItem
}

// 获取钱包列表（需登录）
type WalletListReq struct {
	g.Meta   `path:"/wallet" tags:"Wallet" method:"get" summary:"钱包列表" login_required:"true"`
	Page     int `json:"page" d:"1"`
	PageSize int `json:"pageSize" d:"20" v:"max:100#每页最多100条"`
}
type WalletListRes struct {
	g.Meta `mime:"application/json"`
	List   []model.WalletItem `json:"list"`
	Total  int                `json:"total"`
	Page   int                `json:"page"`
}

