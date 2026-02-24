// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Wallet is the golang structure for table wallet.
type Wallet struct {
	Id               int64       `json:"id"               orm:"id"                 description:""`                   //
	UserId           int64       `json:"userId"           orm:"user_id"            description:"所属用户ID"`             // 所属用户ID
	Address          string      `json:"address"          orm:"address"            description:"钱包地址"`               // 钱包地址
	ApiWalletAddress string      `json:"apiWalletAddress" orm:"api_wallet_address" description:"API Wallet Address"` // API Wallet Address
	ApiSecretKey     string      `json:"apiSecretKey"     orm:"api_secret_key"     description:"API Secret Key"`     // API Secret Key
	Remark           string      `json:"remark"           orm:"remark"             description:"备注"`                 // 备注
	Status           int         `json:"status"           orm:"status"             description:"状态 1:正常 0:禁用"`       // 状态 1:正常 0:禁用
	CreatedAt        *gtime.Time `json:"createdAt"        orm:"created_at"         description:""`                   //
	UpdatedAt        *gtime.Time `json:"updatedAt"        orm:"updated_at"         description:""`                   //
}
