// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MembershipDao is the data access object for the table membership.
type MembershipDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  MembershipColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// MembershipColumns defines and stores column names for the table membership.
type MembershipColumns struct {
	Id        string // 主键ID
	UserId    string // 所属用户ID
	Level     string // 会员等级 0:免费 1:基础 2:高级 3:专业
	StartAt   string // 会员开始时间
	ExpireAt  string // 会员到期时间
	Status    string // 状态 1:正常 0:禁用
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// membershipColumns holds the columns for the table membership.
var membershipColumns = MembershipColumns{
	Id:        "id",
	UserId:    "user_id",
	Level:     "level",
	StartAt:   "start_at",
	ExpireAt:  "expire_at",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewMembershipDao creates and returns a new DAO object for table data access.
func NewMembershipDao(handlers ...gdb.ModelHandler) *MembershipDao {
	return &MembershipDao{
		group:    "default",
		table:    "membership",
		columns:  membershipColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *MembershipDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *MembershipDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *MembershipDao) Columns() MembershipColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *MembershipDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *MembershipDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *MembershipDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
