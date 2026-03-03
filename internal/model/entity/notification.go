// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Notification is the golang structure for table notification.
type Notification struct {
	Id        int64       `json:"id"        orm:"id"         description:"主键ID"`                                        // 主键ID
	UserId    int64       `json:"userId"    orm:"user_id"    description:"所属用户ID,0表示公共通知"`                              // 所属用户ID,0表示公共通知
	Category  string      `json:"category"  orm:"category"   description:"通知类型 public/copy_trading/whale/track/market"` // 通知类型 public/copy_trading/whale/track/market
	Title     string      `json:"title"     orm:"title"      description:"通知标题"`                                        // 通知标题
	Content   string      `json:"content"   orm:"content"    description:"通知内容"`                                        // 通知内容
	RefId     int64       `json:"refId"     orm:"ref_id"     description:"关联业务ID,0表示无关联"`                               // 关联业务ID,0表示无关联
	RefType   string      `json:"refType"   orm:"ref_type"   description:"关联业务类型 copy_trading/track_wallet/position"`   // 关联业务类型 copy_trading/track_wallet/position
	Level     int         `json:"level"     orm:"level"      description:"通知级别 0:普通 1:重要 2:紧急"`                         // 通知级别 0:普通 1:重要 2:紧急
	Status    int         `json:"status"    orm:"status"     description:"状态 1:正常 0:已撤回"`                               // 状态 1:正常 0:已撤回
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`                                        // 创建时间
}
