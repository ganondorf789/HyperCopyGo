package cron_task

import (
	"context"
	"fmt"

	v1 "demo/api/cron_task/v1"
	"demo/internal/dao"
	"demo/internal/model"
	"demo/internal/model/do"
	"demo/internal/model/entity"
	"demo/internal/service"
)

func init() {
	service.RegisterCronTask(&sCronTask{})
}

type sCronTask struct{}

func (s *sCronTask) Create(ctx context.Context, in v1.CronTaskCreateReq) (res *v1.CronTaskCreateRes, err error) {
	id, err := dao.CronTask.Ctx(ctx).Data(do.CronTask{
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
	count, err := dao.CronTask.Ctx(ctx).Where(do.CronTask{Id: in.Id}).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("定时任务不存在")
	}

	_, err = dao.CronTask.Ctx(ctx).
		Where(do.CronTask{Id: in.Id}).
		Data(do.CronTask{
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
	_, err := dao.CronTask.Ctx(ctx).Where(do.CronTask{Id: id}).Delete()
	return err
}

func (s *sCronTask) Detail(ctx context.Context, id int64) (res *v1.CronTaskDetailRes, err error) {
	var item entity.CronTask
	err = dao.CronTask.Ctx(ctx).Where(do.CronTask{Id: id}).Scan(&item)
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
		m = m.Where(do.CronTask{TaskType: in.TaskType})
	}
	if in.Status >= 0 {
		m = m.Where(do.CronTask{Status: in.Status})
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

	list := make([]v1.CronTaskItem, 0, len(items))
	for _, e := range items {
		list = append(list, entityToItem(e))
	}

	return &v1.CronTaskListRes{
		List:  list,
		Total: total,
		Page:  in.Page,
	}, nil
}

func entityToItem(e entity.CronTask) v1.CronTaskItem {
	return v1.CronTaskItem{
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
