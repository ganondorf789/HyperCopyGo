package v1

import (
	"demo/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// 创建定时任务（管理员）
type CronTaskCreateReq struct {
	g.Meta `path:"/cron-task" tags:"CronTask" method:"post" summary:"创建定时任务" login_required:"true" admin_required:"true"`
	model.BaseCronTask
	Status int `json:"status" d:"1"`
}
type CronTaskCreateRes struct {
	g.Meta `mime:"application/json"`
	Id     int64 `json:"id"`
}

// 更新定时任务（管理员）
type CronTaskUpdateReq struct {
	g.Meta `path:"/cron-task/{id}" tags:"CronTask" method:"put" summary:"更新定时任务" login_required:"true" admin_required:"true"`
	Id     int64 `json:"id" in:"path" v:"required"`
	model.BaseCronTask
	Status int `json:"status"`
}
type CronTaskUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// 删除定时任务（管理员）
type CronTaskDeleteReq struct {
	g.Meta `path:"/cron-task/{id}" tags:"CronTask" method:"delete" summary:"删除定时任务" login_required:"true" admin_required:"true"`
	Id     int64 `json:"id" in:"path" v:"required"`
}
type CronTaskDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// 定时任务详情（管理员）
type CronTaskDetailReq struct {
	g.Meta `path:"/cron-task/{id}" tags:"CronTask" method:"get" summary:"定时任务详情" login_required:"true" admin_required:"true"`
	Id     int64 `json:"id" in:"path" v:"required"`
}
type CronTaskDetailRes struct {
	g.Meta `mime:"application/json"`
	model.CronTaskItem
}

// 定时任务列表（管理员）
type CronTaskListReq struct {
	g.Meta   `path:"/cron-task" tags:"CronTask" method:"get" summary:"定时任务列表" login_required:"true" admin_required:"true"`
	TaskType string `json:"taskType"`
	Status   int    `json:"status" d:"-1"`
	Page     int    `json:"page" d:"1"`
	PageSize int    `json:"pageSize" d:"20" v:"max:100#每页最多100条"`
}
type CronTaskListRes struct {
	g.Meta `mime:"application/json"`
	List   []model.CronTaskItem `json:"list"`
	Total  int                  `json:"total"`
	Page   int                  `json:"page"`
}