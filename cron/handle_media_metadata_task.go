package cron

import (
	"fmt"
	"github.com/limoxi/ghost"
	"github.com/limoxi/ghost/utils/cron"
	bs_media "picasso/business/app/media"
	db_media "picasso/db/media"
)

// handleMediaMetadataTask 获取图片&视频的元数据
// 包括缩略图、创建时间、创建地点等
type handleMediaMetadataTask struct {
	cron.CronTask
	cron.Pipe
}

func (this *handleMediaMetadataTask) RunConsumer(data interface{}, taskCtx *cron.TaskContext){
	dbModel := data.(*db_media.Media)
	mediaProcessor := bs_media.NewMediaMetadataProcessor(taskCtx.GetCtx())
	switch dbModel.Type {
	case db_media.MEDIA_TYPE_IMAGE:
		mediaProcessor.ProcessImage(nil)
	case db_media.MEDIA_TYPE_VIDEO:
		mediaProcessor.ProcessVideo()
	default:
		ghost.Error(fmt.Sprintf("unknown media type: %d", dbModel.Type))
	}
}

func (this *handleMediaMetadataTask) Run(taskCtx *cron.TaskContext) {
	db := taskCtx.GetDb()
	var dbModels []*db_media.Media
	result := db.Model(&db_media.Media{}).Where(ghost.Map{
		"status": db_media.MEDIA_STATUS_SAVED,
	}).Limit(10).Find(&dbModels)
	if err := result.Error; err != nil{
		panic(err)
	}
	for _, dbModel := range dbModels{
		err := this.AddData(dbModel)
		if err != nil{
			break
		}
	}
}

func NewHandleMediaMetadataTask() *handleMediaMetadataTask{
	task := new(handleMediaMetadataTask)
	task.SetName("handle_media_metadata_task")
	task.Init(100)
	return task
}

func init() {
	task := NewHandleMediaMetadataTask()
	cron.RegisterPipeTask(task, "*/5 * * * * *")
}