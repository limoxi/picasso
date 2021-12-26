package cron

import (
	"github.com/limoxi/ghost"
	"github.com/limoxi/ghost/utils/cron"
)

type demoTask struct {
	cron.CronTask
}

func (this *demoTask) Run(taskCtx *cron.TaskContext) {
	ghost.Info("[demo_task] run...")
}

func NewDemoTask() *demoTask {
	task := new(demoTask)
	task.SetName("demo_task")
	return task
}

func init() {
	//task := NewDemoTask()
	//cron.RegisterTask(task, "*/5 * * * * *")
}
