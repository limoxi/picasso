package media

import (
	"github.com/limoxi/ghost"
	"time"
)

const MEDIA_STATUS_SLICE_SAVED = -1
const MEDIA_STATUS_SAVED = 0
const MEDIA_STATUS_PENDING = 1
const MEDIA_STATUS_COMPLETE = 2
var MEDIA_STATUS2TEXT = map[int]string{
	MEDIA_STATUS_SLICE_SAVED: "slice_saved",
	MEDIA_STATUS_SAVED: "saved",
	MEDIA_STATUS_PENDING: "pending",
	MEDIA_STATUS_COMPLETE: "complete",
}

const SLICED_MEDIA_STATUS_SAVED = 0  // 已保存
const SLICED_MEDIA_STATUS_MERGED = 1  // 已合并

const MEDIA_TYPE_IMAGE = 1
const MEDIA_TYPE_VIDEO = 2
var MEDIA_TYPE2TEXT = map[int]string{
	MEDIA_TYPE_IMAGE: "image",
	MEDIA_TYPE_VIDEO: "video",
}
var MEDIA_TEXT2TYPE = map[string]int{
	"image": MEDIA_TYPE_IMAGE,
	"video": MEDIA_TYPE_VIDEO,
}

// Media 媒体资源信息
type Media struct {
	ghost.BaseModel
	SpaceId int
	Type int
	Hash string `gorm:"size:128"`
	ThumbnailPath string `gorm:"size:256;default('')"`
	StoragePath string `gorm:"size:256"`
	Status int
	Metadata string `gorm:"type:text"` // 元信息
	ShootTime time.Time // 拍摄时间
	ShootLocation string `gorm:"size:256;default('')"` // 拍摄地点
	Size int64 // 大小，单位B
	Duration int // 时长，单位秒
}
func (Media) TableName() string{
	return "media_media"
}

// SlicedMedia 分片的媒体资源
type SlicedMedia struct{
	ghost.BaseModel
	MediaType int
	MediaHash string `gorm:"size:128"`
	SliceHash string `gorm:"size:128"`
	SliceIndex int
	TotalSliceCount int
	StoragePath string `gorm:"size:256"`
	Status int
	Size int64 // 大小，单位B
}
func (SlicedMedia) TableName() string{
	return "media_sliced_media"
}

func init(){
	ghost.RegisterDBModel(&Media{})
	ghost.RegisterDBModel(&SlicedMedia{})
}