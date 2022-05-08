package cron

import (
	"github.com/limoxi/ghost"
	"github.com/limoxi/ghost/utils/cron"
	bm_file "picasso/business/model/file"
	bs_file "picasso/business/service/file"
	db_file "picasso/db/file"
)

// handleMediaMetadataTask 获取图片&视频的元数据
// 包括缩略图、创建时间、创建地点等
type handleMediaMetadataTask struct {
	cron.CronTask
	cron.Pipe
}

func (this *handleMediaMetadataTask) RunConsumer(data interface{}, taskCtx *cron.TaskContext) {
	mediaFile := data.(*bm_file.File)
	mediaProcessor := bs_file.NewMediaMetadataProcessor(taskCtx.GetCtx())
	mediaProcessor.Process(mediaFile)
}

func (this *handleMediaMetadataTask) Run(taskCtx *cron.TaskContext) {
	files := bm_file.NewFileRepository(taskCtx.GetCtx()).GetByFilters(ghost.Map{
		"status": db_file.FILE_STATUS_SAVED,
	})
	for _, file := range files {
		err := this.AddData(file)
		if err != nil {
			break
		}
	}
}

func NewHandleMediaMetadataTask() *handleMediaMetadataTask {
	task := new(handleMediaMetadataTask)
	task.SetName("handle_media_metadata_task")
	task.Init(100)
	return task
}

func init() {
	//task := NewHandleMediaMetadataTask()
	//cron.RegisterPipeTask(task, "*/5 * * * * *")
}
