package cron_jobs

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
)

type cronJob struct {
	Name     string
	CronExpr string
	Fn       func(ctx context.Context)
}

var jobs []cronJob

func registerJob(name, cronExpr string, fn func(ctx context.Context)) {
	jobs = append(jobs, cronJob{Name: name, CronExpr: cronExpr, Fn: fn})
}

// StartAll 启动所有已注册的定时任务
func StartAll(ctx context.Context) {
	for _, job := range jobs {
		j := job
		if _, err := gcron.Add(ctx, j.CronExpr, func(ctx context.Context) {
			j.Fn(ctx)
		}, j.Name); err != nil {
			g.Log().Errorf(ctx, "注册定时任务 [%s] 失败: %v", j.Name, err)
		} else {
			g.Log().Infof(ctx, "定时任务 [%s] 已注册, cron=%s", j.Name, j.CronExpr)
		}
	}
}
