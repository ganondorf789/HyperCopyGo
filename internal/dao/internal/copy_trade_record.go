// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CopyTradeRecordDao is the data access object for the table copy_trade_record.
type CopyTradeRecordDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  CopyTradeRecordColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// CopyTradeRecordColumns defines and stores column names for the table copy_trade_record.
type CopyTradeRecordColumns struct {
	Id            string // 主键ID
	UserId        string // 所属用户ID
	Address       string // 钱包地址
	Coin          string // 币种
	Direction     string // 方向（Open Long/Open Short/Close Long/Close Short）
	Size          string // 成交规模（张数）
	Price         string // 成交价格
	ClosedPnl     string // 已实现盈亏（USD）
	ExecuteStatus string // 执行状态 0:待执行 1:成功 2:失败 3:跳过
	OrderStatus   string // 订单状态 open/filled/canceled/triggered
	ErrorMsg      string // 执行失败原因
	TradeTime     string // 触发交易时间（源头成交时间）
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
}

// copyTradeRecordColumns holds the columns for the table copy_trade_record.
var copyTradeRecordColumns = CopyTradeRecordColumns{
	Id:            "id",
	UserId:        "user_id",
	Address:       "address",
	Coin:          "coin",
	Direction:     "direction",
	Size:          "size",
	Price:         "price",
	ClosedPnl:     "closed_pnl",
	ExecuteStatus: "execute_status",
	OrderStatus:   "order_status",
	ErrorMsg:      "error_msg",
	TradeTime:     "trade_time",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewCopyTradeRecordDao creates and returns a new DAO object for table data access.
func NewCopyTradeRecordDao(handlers ...gdb.ModelHandler) *CopyTradeRecordDao {
	return &CopyTradeRecordDao{
		group:    "default",
		table:    "copy_trade_record",
		columns:  copyTradeRecordColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CopyTradeRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CopyTradeRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CopyTradeRecordDao) Columns() CopyTradeRecordColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CopyTradeRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CopyTradeRecordDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *CopyTradeRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
