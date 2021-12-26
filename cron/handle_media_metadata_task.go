package cron

import (
	"github.com/limoxi/ghost"
	"github.com/limoxi/ghost/utils/cron"
	bs_file "picasso/business/app/file"
	db_file "picasso/db/file"
)

// handleMediaMetadataTask 获取图片&视频的元数据
// 包括缩略图、创建时间、创建地点等
type handleMediaMetadataTask struct {
	cron.CronTask
	cron.Pipe
}

func (this *handleMediaMetadataTask) RunConsumer(data interface{}, taskCtx *cron.TaskContext) {
	dbModel := data.(*db_file.File)
	mediaProcessor := bs_file.NewMediaMetadataProcessor(taskCtx.GetCtx())
	mediaProcessor.Process(dbModel)
}

func (this *handleMediaMetadataTask) Run(taskCtx *cron.TaskContext) {
	db := taskCtx.GetDb()
	var dbModels []*db_file.File
	result := db.Model(&db_file.File{}).Where(ghost.Map{
		"type":   db_file.FILE_TYPE_MEDIA,
		"status": db_file.FILE_STATUS_SAVED,
	}).Limit(10).Find(&dbModels)
	if err := result.Error; err != nil {
		panic(err)
	}
	for _, dbModel := range dbModels {
		err := this.AddData(dbModel)
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
	task := NewHandleMediaMetadataTask()
	cron.RegisterPipeTask(task, "*/5 * * * * *")
}
