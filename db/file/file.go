package file

import (
	"github.com/limoxi/ghost"
	"time"
)

const FILE_TYPE_DIR = 1  // 目录
const FILE_TYPE_FILE = 2 // 文件

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

const CATEGORY_FILE = "FILE"
const CATEGORY_GALLERY = "GALLERY"
const CATEGORY_VIDEO = "VIDEO"
const CATEGORY_SHARE = "SHARE"
const CATEGORY_RECYCLE = "RECYCLE"

var CATEGORY2CN = map[string]string{
	CATEGORY_FILE:    "文件",
	CATEGORY_GALLERY: "相册",
	CATEGORY_VIDEO:   "影音",
	CATEGORY_SHARE:   "分享",
	CATEGORY_RECYCLE: "回收站",
}

// File 文件信息
type File struct {
	ghost.BaseDBModel

	UserId int
	Type   int
	Path   string `gorm:"size:512"`
	Hash   string `gorm:"size:128"`
	Status int

	// 元信息
	Name             string    `gorm:"size:128"`
	Size             int64     // 大小，单位B
	Metadata         string    `gorm:"type:text"` // 元信息
	Thumbnail        string    `gorm:"size:256"`
	LastModifiedTime time.Time `gorm:"autoCreateTime;null"` // 原始文件最后一次修改时间
	CreatedTime      time.Time `gorm:"autoCreateTime;null"` // 原始文件创建时间
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
	Path            string `gorm:"size:256"`
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
