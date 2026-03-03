// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CronTaskDao is the data access object for the table cron_task.
type CronTaskDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  CronTaskColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// CronTaskColumns defines and stores column names for the table cron_task.
type CronTaskColumns struct {
	Id          string // 主键ID
	Name        string // 任务名称
	CronExpr    string // Cron表达式
	TaskType    string // 任务类型 sync_leaderboard/sync_positions/sync_fills/copy_trade/track_wallet/market_alert
	Params      string // 任务参数(JSON)
	LastRunAt   string // 上次执行时间
	LastRunCost string // 上次执行耗时(毫秒)
	LastError   string // 上次执行错误信息
	RunCount    string // 累计执行次数
	Status      string // 状态 1:启用 0:停用
	Remark      string // 备注
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
}

// cronTaskColumns holds the columns for the table cron_task.
var cronTaskColumns = CronTaskColumns{
	Id:          "id",
	Name:        "name",
	CronExpr:    "cron_expr",
	TaskType:    "task_type",
	Params:      "params",
	LastRunAt:   "last_run_at",
	LastRunCost: "last_run_cost",
	LastError:   "last_error",
	RunCount:    "run_count",
	Status:      "status",
	Remark:      "remark",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewCronTaskDao creates and returns a new DAO object for table data access.
func NewCronTaskDao(handlers ...gdb.ModelHandler) *CronTaskDao {
	return &CronTaskDao{
		group:    "default",
		table:    "cron_task",
		columns:  cronTaskColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CronTaskDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CronTaskDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CronTaskDao) Columns() CronTaskColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CronTaskDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CronTaskDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *CronTaskDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
