// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MyTrackWalletDao is the data access object for the table my_track_wallet.
type MyTrackWalletDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  MyTrackWalletColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// MyTrackWalletColumns defines and stores column names for the table my_track_wallet.
type MyTrackWalletColumns struct {
	Id           string // 主键ID
	UserId       string // 所属用户ID
	Wallet       string // 跟踪的钱包地址
	Remark       string // 备注
	EnableNotify string // 是否开启通知 0:关 1:开
	NotifyAction string // 通知动作 1:开仓 2:平仓 3:加仓 4:减仓
	Lang         string // 语言
	Status       string // 状态 1:正常 0:禁用
	CreatedAt    string // 创建时间
	UpdatedAt    string // 更新时间
}

// myTrackWalletColumns holds the columns for the table my_track_wallet.
var myTrackWalletColumns = MyTrackWalletColumns{
	Id:           "id",
	UserId:       "user_id",
	Wallet:       "wallet",
	Remark:       "remark",
	EnableNotify: "enable_notify",
	NotifyAction: "notify_action",
	Lang:         "lang",
	Status:       "status",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewMyTrackWalletDao creates and returns a new DAO object for table data access.
func NewMyTrackWalletDao(handlers ...gdb.ModelHandler) *MyTrackWalletDao {
	return &MyTrackWalletDao{
		group:    "default",
		table:    "my_track_wallet",
		columns:  myTrackWalletColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *MyTrackWalletDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *MyTrackWalletDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *MyTrackWalletDao) Columns() MyTrackWalletColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *MyTrackWalletDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *MyTrackWalletDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *MyTrackWalletDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
