// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LeaderboardDao is the data access object for the table leaderboard.
type LeaderboardDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  LeaderboardColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// LeaderboardColumns defines and stores column names for the table leaderboard.
type LeaderboardColumns struct {
	Id           string // 主键ID
	EthAddress   string // 钱包地址
	AccountValue string // 账户价值
	Pnl          string // 盈亏
	Roi          string // 投资回报率
	Vlm          string // 交易量
	CreatedAt    string // 创建时间
	UpdatedAt    string // 更新时间
	Window       string // 统计窗口 day/week/month/allTime
}

// leaderboardColumns holds the columns for the table leaderboard.
var leaderboardColumns = LeaderboardColumns{
	Id:           "id",
	EthAddress:   "eth_address",
	AccountValue: "account_value",
	Pnl:          "pnl",
	Roi:          "roi",
	Vlm:          "vlm",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	Window:       "window",
}

// NewLeaderboardDao creates and returns a new DAO object for table data access.
func NewLeaderboardDao(handlers ...gdb.ModelHandler) *LeaderboardDao {
	return &LeaderboardDao{
		group:    "default",
		table:    "leaderboard",
		columns:  leaderboardColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *LeaderboardDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *LeaderboardDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *LeaderboardDao) Columns() LeaderboardColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *LeaderboardDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *LeaderboardDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *LeaderboardDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
