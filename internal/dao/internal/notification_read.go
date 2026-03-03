// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NotificationReadDao is the data access object for the table notification_read.
type NotificationReadDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  NotificationReadColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// NotificationReadColumns defines and stores column names for the table notification_read.
type NotificationReadColumns struct {
	Id             string // 主键ID
	UserId         string // 用户ID
	NotificationId string // 通知ID
	ReadAt         string // 已读时间
}

// notificationReadColumns holds the columns for the table notification_read.
var notificationReadColumns = NotificationReadColumns{
	Id:             "id",
	UserId:         "user_id",
	NotificationId: "notification_id",
	ReadAt:         "read_at",
}

// NewNotificationReadDao creates and returns a new DAO object for table data access.
func NewNotificationReadDao(handlers ...gdb.ModelHandler) *NotificationReadDao {
	return &NotificationReadDao{
		group:    "default",
		table:    "notification_read",
		columns:  notificationReadColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *NotificationReadDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *NotificationReadDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *NotificationReadDao) Columns() NotificationReadColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *NotificationReadDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *NotificationReadDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *NotificationReadDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
