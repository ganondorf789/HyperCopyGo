// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Wallet is the golang structure of table wallet for DAO operations like Where/Data.
type Wallet struct {
	g.Meta           `orm:"table:wallet, do:true"`
	Id               any         // 主键ID
	UserId           any         // 所属用户ID
	Address          any         // 钱包地址
	ApiWalletAddress any         // API Wallet Address
	ApiSecretKey     any         // API Secret Key
	Remark           any         // 备注
	Status           any         // 状态 1:正常 0:禁用
	CreatedAt        *gtime.Time // 创建时间
	UpdatedAt        *gtime.Time // 更新时间
}
