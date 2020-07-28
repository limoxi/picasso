package cron

import (
	"github.com/limoxi/ghost"
	"github.com/limoxi/ghost/utils/cron"
)

// mergeSlicedMediaTask 合并切片上传的文件
type mergeSlicedMediaTask struct {
	cron.CronTask
}


func (this *mergeSlicedMediaTask) Run(taskCtx *cron.TaskContext) error {
	ghost.Info("[merge_sliced_media_task] run...")
	return nil
}

func NewMergeSlicedMediaTask() *mergeSlicedMediaTask{
	task := new(mergeSlicedMediaTask)
	task.CronTask = cron.NewCronTask("merge_sliced_media_task")
	return task
}

func init() {
	task := NewMergeSlicedMediaTask()
	cron.RegisterTask(task, "0 */5 * * * *", true)
}