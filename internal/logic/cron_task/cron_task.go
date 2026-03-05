package cron_task

import (
	"context"
	"fmt"
	"sort"
	"time"

	v1 "demo/api/cron_task/v1"
	cronJobs "demo/internal/cron_jobs"
	"demo/internal/dao"
	"demo/internal/model"
	"demo/internal/model/entity"
	"demo/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gtime"
)

func init() {
	service.RegisterCronTask(&sCronTask{})
}

type sCronTask struct{}

func (s *sCronTask) Create(ctx context.Context, in v1.CronTaskCreateReq) (res *v1.CronTaskCreateRes, err error) {
	id, err := dao.CronTask.Ctx(ctx).Data(entity.CronTask{
		Name:     in.Name,
		CronExpr: in.CronExpr,
		TaskType: in.TaskType,
		Params:   in.Params,
		Remark:   in.Remark,
		Status:   in.Status,
	}).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return &v1.CronTaskCreateRes{Id: id}, nil
}

func (s *sCronTask) Update(ctx context.Context, in v1.CronTaskUpdateReq) error {
	count, err := dao.CronTask.Ctx(ctx).Where("id = ?", in.Id).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("定时任务不存在")
	}

	_, err = dao.CronTask.Ctx(ctx).
		Where("id = ?", in.Id).
		Data(entity.CronTask{
			Name:     in.Name,
			CronExpr: in.CronExpr,
			TaskType: in.TaskType,
			Params:   in.Params,
			Remark:   in.Remark,
			Status:   in.Status,
		}).
		Update()
	return err
}

func (s *sCronTask) Delete(ctx context.Context, id int64) error {
	_, err := dao.CronTask.Ctx(ctx).Where("id = ?", id).Delete()
	return err
}

func (s *sCronTask) Detail(ctx context.Context, id int64) (res *v1.CronTaskDetailRes, err error) {
	var item entity.CronTask
	err = dao.CronTask.Ctx(ctx).Where("id = ?", id).Scan(&item)
	if err != nil {
		return nil, err
	}
	if item.Id == 0 {
		return nil, fmt.Errorf("定时任务不存在")
	}
	return &v1.CronTaskDetailRes{
		CronTaskItem: entityToItem(item),
	}, nil
}

func (s *sCronTask) List(ctx context.Context, in v1.CronTaskListReq) (res *v1.CronTaskListRes, err error) {
	m := dao.CronTask.Ctx(ctx)
	if in.TaskType != "" {
		m = m.Where("task_type = ?", in.TaskType)
	}
	if in.Status >= 0 {
		m = m.Where("status = ?", in.Status)
	}

	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	var items []entity.CronTask
	err = m.Page(in.Page, in.PageSize).
		OrderDesc(dao.CronTask.Columns().Id).
		Scan(&items)
	if err != nil {
		return nil, err
	}

	list := make([]model.CronTaskItem, 0, len(items))
	for _, e := range items {
		list = append(list, entityToItem(e))
	}

	return &v1.CronTaskListRes{
		List:  list,
		Total: total,
		Page:  in.Page,
	}, nil
}

// Execute 手动执行指定定时任务
func (s *sCronTask) Execute(ctx context.Context, id int64) (res *v1.CronTaskExecuteRes, err error) {
	var task entity.CronTask
	err = dao.CronTask.Ctx(ctx).Where("id = ?", id).Scan(&task)
	if err != nil {
		return nil, err
	}
	if task.Id == 0 {
		return nil, fmt.Errorf("定时任务不存在")
	}

	handler, ok := cronJobs.Handlers[task.TaskType]
	if !ok {
		return nil, fmt.Errorf("未注册的任务类型: %s", task.TaskType)
	}

	costMs, lastErr := s.runTask(ctx, task, handler)

	return &v1.CronTaskExecuteRes{
		LastRunCost: costMs,
		LastError:   lastErr,
	}, nil
}

// TaskTypes 返回所有已注册的可用任务类型
func (s *sCronTask) TaskTypes(ctx context.Context) (res *v1.CronTaskTypesRes, err error) {
	types := make([]string, 0, len(cronJobs.Handlers))
	for t := range cronJobs.Handlers {
		types = append(types, t)
	}
	sort.Strings(types)
	return &v1.CronTaskTypesRes{Types: types}, nil
}

// StartAll 从数据库加载所有启用的定时任务，根据 cron_expr 注册到 gcron 调度器
func (s *sCronTask) StartAll(ctx context.Context) {
	var items []entity.CronTask
	err := dao.CronTask.Ctx(ctx).Where("status = ?", 1).Scan(&items)
	if err != nil {
		g.Log().Errorf(ctx, "StartAll: 加载定时任务失败: %v", err)
		return
	}

	for _, item := range items {
		handler, ok := cronJobs.Handlers[item.TaskType]
		if !ok {
			g.Log().Warningf(ctx, "StartAll: 任务[%s]类型[%s]未注册，跳过", item.Name, item.TaskType)
			continue
		}

		task := item
		fn := handler
		if _, err := gcron.Add(ctx, task.CronExpr, func(ctx context.Context) {
			s.runTask(ctx, task, fn)
		}, task.Name); err != nil {
			g.Log().Errorf(ctx, "StartAll: 注册任务[%s]失败: %v", task.Name, err)
		} else {
			g.Log().Infof(ctx, "StartAll: 任务[%s]已注册, cron=%s, type=%s", task.Name, task.CronExpr, task.TaskType)
		}
	}
}

// runTask 执行任务并更新运行统计（last_run_at、last_run_cost、last_error、run_count）
func (s *sCronTask) runTask(ctx context.Context, task entity.CronTask, handler cronJobs.TaskHandler) (costMs int64, lastErr string) {
	start := time.Now()

	func() {
		defer func() {
			if r := recover(); r != nil {
				lastErr = fmt.Sprintf("panic: %v", r)
			}
		}()
		handler(ctx, task.Params)
	}()

	costMs = time.Since(start).Milliseconds()

	_, err := dao.CronTask.Ctx(ctx).
		Where("id = ?", task.Id).
		Data(g.Map{
			dao.CronTask.Columns().LastRunAt:   gtime.Now(),
			dao.CronTask.Columns().LastRunCost: costMs,
			dao.CronTask.Columns().LastError:   lastErr,
			dao.CronTask.Columns().RunCount:    gdb.Raw("run_count + 1"),
		}).
		Update()
	if err != nil {
		g.Log().Errorf(ctx, "runTask: 更新任务[%s]运行状态失败: %v", task.Name, err)
	}

	return
}

func entityToItem(e entity.CronTask) model.CronTaskItem {
	return model.CronTaskItem{
		Id: e.Id,
		BaseCronTask: model.BaseCronTask{
			Name:     e.Name,
			CronExpr: e.CronExpr,
			TaskType: e.TaskType,
			Params:   e.Params,
			Remark:   e.Remark,
		},
		LastRunAt:   e.LastRunAt,
		LastRunCost: e.LastRunCost,
		LastError:   e.LastError,
		RunCount:    e.RunCount,
		Status:      e.Status,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}
