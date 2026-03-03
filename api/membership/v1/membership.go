package v1

import (
	"demo/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 开通会员（管理员）
type MembershipCreateReq struct {
	g.Meta   `path:"/membership" tags:"Membership" method:"post" summary:"开通会员" login_required:"true" admin_required:"true"`
	UserId   int64       `json:"userId" v:"required#请输入用户ID"`
	Level    int         `json:"level" v:"required|in:0,1,2,3#请选择会员等级#会员等级不合法"`
	StartAt  *gtime.Time `json:"startAt" v:"required#请输入开始时间"`
	ExpireAt *gtime.Time `json:"expireAt" v:"required#请输入到期时间"`
	Status   int         `json:"status" d:"1"`
}
type MembershipCreateRes struct {
	g.Meta `mime:"application/json"`
	Id     int64 `json:"id"`
}

// 编辑会员（管理员）
type MembershipUpdateReq struct {
	g.Meta   `path:"/membership/{id}" tags:"Membership" method:"put" summary:"编辑会员" login_required:"true" admin_required:"true"`
	Id       int64       `json:"id" in:"path" v:"required"`
	Level    int         `json:"level"`
	StartAt  *gtime.Time `json:"startAt"`
	ExpireAt *gtime.Time `json:"expireAt"`
	Status   int         `json:"status"`
}
type MembershipUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// 删除会员（管理员）
type MembershipDeleteReq struct {
	g.Meta `path:"/membership/{id}" tags:"Membership" method:"delete" summary:"删除会员" login_required:"true" admin_required:"true"`
	Id     int64 `json:"id" in:"path" v:"required"`
}
type MembershipDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// 会员详情（管理员）
type MembershipDetailReq struct {
	g.Meta `path:"/membership/{id}" tags:"Membership" method:"get" summary:"会员详情" login_required:"true" admin_required:"true"`
	Id     int64 `json:"id" in:"path" v:"required"`
}
type MembershipDetailRes struct {
	g.Meta `mime:"application/json"`
	model.MembershipItem
}

// 会员列表（管理员）
type MembershipListReq struct {
	g.Meta   `path:"/membership" tags:"Membership" method:"get" summary:"会员列表" login_required:"true" admin_required:"true"`
	UserId   int64 `json:"userId"`
	Level    int   `json:"level" d:"-1"`
	Status   int   `json:"status" d:"-1"`
	Page     int   `json:"page" d:"1"`
	PageSize int   `json:"pageSize" d:"20" v:"max:100#每页最多100条"`
}
type MembershipListRes struct {
	g.Meta `mime:"application/json"`
	List   []model.MembershipItem `json:"list"`
	Total  int                    `json:"total"`
	Page   int                    `json:"page"`
}
