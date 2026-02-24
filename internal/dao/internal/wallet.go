// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WalletDao is the data access object for the table wallet.
type WalletDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  WalletColumns      // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// WalletColumns defines and stores column names for the table wallet.
type WalletColumns struct {
	Id               string //
	UserId           string // 所属用户ID
	Address          string // 钱包地址
	ApiWalletAddress string // API Wallet Address
	ApiSecretKey     string // API Secret Key
	Remark           string // 备注
	Status           string // 状态 1:正常 0:禁用
	CreatedAt        string //
	UpdatedAt        string //
}

// walletColumns holds the columns for the table wallet.
var walletColumns = WalletColumns{
	Id:               "id",
	UserId:           "user_id",
	Address:          "address",
	ApiWalletAddress: "api_wallet_address",
	ApiSecretKey:     "api_secret_key",
	Remark:           "remark",
	Status:           "status",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
}

// NewWalletDao creates and returns a new DAO object for table data access.
func NewWalletDao(handlers ...gdb.ModelHandler) *WalletDao {
	return &WalletDao{
		group:    "default",
		table:    "wallet",
		columns:  walletColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *WalletDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *WalletDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *WalletDao) Columns() WalletColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *WalletDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *WalletDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *WalletDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
