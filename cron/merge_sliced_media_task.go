package cron

import (
	"bufio"
	"github.com/limoxi/ghost"
	"github.com/limoxi/ghost/utils/cron"
	"io/ioutil"
	"os"
	"path"
	bs_media "picasso/business/service/media"
	db_media "picasso/db/media"
	"sort"
	"strconv"
	"strings"
)

// mergeSlicedMediaTask 合并切片上传的文件
type mergeSlicedMediaTask struct {
	cron.CronTask
}

func (this *mergeSlicedMediaTask) allSlicesIsHere(slices []string) bool{
	curIndexes := make([]int, 0)
	totalSliceCount := 0
	for _, slice := range slices{
		if !strings.HasSuffix(slice, ".slice"){
			continue
		}
		slice = strings.Split(slice, ".slice")[0]
		sps := strings.Split(slice, "_")
		l := len(sps)
		curIndex, _ := strconv.Atoi(sps[l-1])
		curIndexes = append(curIndexes, curIndex)
		totalSliceCount, _ = strconv.Atoi(sps[l-2])
	}
	if len(curIndexes) != totalSliceCount{
		return false
	}
	sort.Ints(curIndexes)
	for i, index := range curIndexes{
		if i != index{
			return false
		}
	}
	return true
}

func (this *mergeSlicedMediaTask) Run(taskCtx *cron.TaskContext) {
	db := taskCtx.GetDb()

	var dbModel db_media.Media
	result := db.Model(&db_media.Media{}).Where(ghost.Map{
		"status": db_media.MEDIA_STATUS_SLICE_SAVED,
	}).First(&dbModel)
	if err := result.Error; err != nil{
		if result.RecordNotFound(){
			return
		}
		ghost.Error(err)
		panic(err)
	}

	ghost.Info("[merge_sliced_media_task] start handle media: " + dbModel.Hash)

	sps := strings.Split(dbModel.StoragePath, "___")
	tmpDirPath := sps[0]
	filename := sps[1]

	fs, err := ioutil.ReadDir(tmpDirPath)
	if err != nil{
		ghost.Error(err)
		panic(err)
	}
	slices := make([]string, 0, len(fs))
	for _, f := range fs{
		name := f.Name()
		if strings.HasSuffix(name, ".slice"){
			slices = append(slices, f.Name())
		}
	}
	if !this.allSlicesIsHere(slices){
		ghost.Warn("wait left slices...")
		return
	}
	sort.Strings(slices)
	targetFilePath := path.Join(bs_media.MEDIA_TYPE2STORAGE_PATH[dbModel.Type], filename)
	targetFile, err := os.OpenFile(targetFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil{
		ghost.Error(err)
		panic(err)
	}
	defer targetFile.Close()
	writer := bufio.NewWriterSize(targetFile, 10 << 21)
	for _, slice := range slices{
		ghost.Info("merge slice: ", slice)
		bytes, err := ioutil.ReadFile(path.Join(tmpDirPath, slice))
		if err != nil{
			ghost.Error(err)
			panic(err)
		}
		_, err = writer.Write(bytes)
		if err != nil{
			ghost.Error(err)
			panic(err)
		}
	}

	err = writer.Flush()
	if err != nil{
		ghost.Error(err)
		panic(err)
	}

	result = db.Model(&db_media.Media{}).Where(ghost.Map{
		"type": dbModel.Type,
		"hash": dbModel.Hash,
	}).Update(ghost.Map{
		"storage_path": targetFilePath,
		"status": db_media.MEDIA_STATUS_SAVED,
	})
	if err = result.Error; err != nil{
		os.Remove(targetFilePath)
		panic(err)
	}

	err = os.RemoveAll(tmpDirPath)
	if err != nil{
		panic(err)
	}
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