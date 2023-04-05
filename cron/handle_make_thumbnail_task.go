package cron

import (
	"github.com/limoxi/ghost"
	"github.com/limoxi/ghost/utils/cron"
	bm_file "picasso/business/model/file"
	bs_file "picasso/business/service/file"
	db_file "picasso/db/file"
)

// handleMakeThumbnailTask 生成文件的缩略图
type handleMakeThumbnailTask struct {
	cron.CronTask
	cron.Pipe
}

func (this *handleMakeThumbnailTask) RunConsumer(data interface{}, taskCtx *cron.TaskContext) {
	f := data.(*bm_file.File)
	bs_file.NewFileService(taskCtx.GetCtx()).MakeThumbnail(f)
}

func (this *handleMakeThumbnailTask) Run(taskCtx *cron.TaskContext) {
	files := bm_file.NewFileRepository(taskCtx.GetCtx()).GetByFilters(ghost.Map{
		"type":      db_file.FILE_TYPE_FILE,
		"thumbnail": "",
	})
	for _, file := range files {
		err := this.AddData(file)
		if err != nil {
			break
		}
	}
}

func NewHandleMakeThumbnailTask() *handleMakeThumbnailTask {
	task := new(handleMakeThumbnailTask)
	task.SetName("handle_make_thumbnail_task")
	task.Init(10)
	return task
}

func init() {
	task := NewHandleMakeThumbnailTask()
	cron.RegisterPipeTask(task, "*/30 * * * * *")
}
