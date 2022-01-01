package file

import (
	"github.com/limoxi/ghost"
	"time"
)

const FILE_STATUS_SLICE_SAVED = -1
const FILE_STATUS_SAVED = 0
const FILE_STATUS_PENDING = 1
const FILE_STATUS_COMPLETE = 2

var FILE_STATUS2TEXT = map[int]string{
	FILE_STATUS_SLICE_SAVED: "slice_saved",
	FILE_STATUS_SAVED:       "saved",
	FILE_STATUS_PENDING:     "pending",
	FILE_STATUS_COMPLETE:    "complete",
}

const SLICED_FILE_STATUS_SAVED = 0  // 已保存
const SLICED_FILE_STATUS_MERGED = 1 // 已合并

const FILE_TYPE_FILE = 1
const FILE_TYPE_MEDIA = 2

var FILE_TEXT2TYPE = map[string]int{
	"file":  FILE_TYPE_FILE,
	"media": FILE_TYPE_MEDIA,
}

const MEDIA_TYPE_IMAGE = 1
const MEDIA_TYPE_VIDEO = 2
const MEDIA_TYPE_OTHERS = 3

var MEDIA_TYPE2TEXT = map[int]string{
	MEDIA_TYPE_IMAGE:  "image",
	MEDIA_TYPE_VIDEO:  "video",
	MEDIA_TYPE_OTHERS: "others",
}
var MEDIA_TEXT2TYPE = map[string]int{
	"image":  MEDIA_TYPE_IMAGE,
	"video":  MEDIA_TYPE_VIDEO,
	"others": MEDIA_TYPE_OTHERS,
}

// File 文件信息
type File struct {
	ghost.BaseDBModel

	UserId      int
	GroupId     int
	Type        int
	Hash        string `gorm:"size:128"`
	Name        string `gorm:"size:128"`
	StoragePath string `gorm:"size:256"`
	Size        int64  // 大小，单位B
	Metadata    string `gorm:"type:text"` // 元信息
	Status      int

	// 媒体信息
	ThumbnailPath string    `gorm:"size:256"`
	CreatedTime   time.Time // 原始文件创建时间
}

func (File) TableName() string {
	return "file_file"
}

// SlicedFile 分片的文件信息
type SlicedFile struct {
	ghost.BaseDBModel

	FileHash        string `gorm:"size:128"`
	SliceHash       string `gorm:"size:128"`
	SliceIndex      int
	TotalSliceCount int
	StoragePath     string `gorm:"size:256"`
	Status          int
	Size            int64 // 大小，单位B
}

func (SlicedFile) TableName() string {
	return "file_sliced_file"
}

func init() {
	ghost.RegisterDBModel(&File{})
	ghost.RegisterDBModel(&SlicedFile{})
}
