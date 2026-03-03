// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TraderOrdersDao is the data access object for the table trader_orders.
type TraderOrdersDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  TraderOrdersColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// TraderOrdersColumns defines and stores column names for the table trader_orders.
type TraderOrdersColumns struct {
	Id               string // 主键ID
	Address          string // 钱包地址
	Coin             string // 币种
	Side             string // 买卖方向（A=卖/B=买）
	LimitPx          string // 限价
	Sz               string // 委托量
	Oid              string // 订单ID
	Timestamp        string // 委托时间（毫秒时间戳）
	TriggerCondition string // 触发条件
	IsTrigger        string // 是否触发订单
	TriggerPx        string // 触发价
	Children         string // 子订单(JSON)
	IsPositionTpsl   string // 是否为仓位止盈止损
	ReduceOnly       string // 是否只减仓
	OrderType        string // 订单类型
	OrigSz           string // 原始委托量
	Tif              string // 有效期类型
	Cloid            string // 客户端订单ID
	Status           string // 订单状态
	CreatedAt        string // 创建时间
}

// traderOrdersColumns holds the columns for the table trader_orders.
var traderOrdersColumns = TraderOrdersColumns{
	Id:               "id",
	Address:          "address",
	Coin:             "coin",
	Side:             "side",
	LimitPx:          "limit_px",
	Sz:               "sz",
	Oid:              "oid",
	Timestamp:        "timestamp",
	TriggerCondition: "trigger_condition",
	IsTrigger:        "is_trigger",
	TriggerPx:        "trigger_px",
	Children:         "children",
	IsPositionTpsl:   "is_position_tpsl",
	ReduceOnly:       "reduce_only",
	OrderType:        "order_type",
	OrigSz:           "orig_sz",
	Tif:              "tif",
	Cloid:            "cloid",
	Status:           "status",
	CreatedAt:        "created_at",
}

// NewTraderOrdersDao creates and returns a new DAO object for table data access.
func NewTraderOrdersDao(handlers ...gdb.ModelHandler) *TraderOrdersDao {
	return &TraderOrdersDao{
		group:    "default",
		table:    "trader_orders",
		columns:  traderOrdersColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TraderOrdersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TraderOrdersDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TraderOrdersDao) Columns() TraderOrdersColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TraderOrdersDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TraderOrdersDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TraderOrdersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
