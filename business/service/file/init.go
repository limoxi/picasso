package file

import (
	"github.com/limoxi/ghost"
	"os"
	"path"
	db_file "picasso/db/file"
	"strings"
)

var STORAGE_ROOT_PATH string
var FILE_STORAGE_PATH string
var VIDEO_STORAGE_PATH string
var GALLERY_STORAGE_PATH string
var SHARE_STORAGE_PATH string
var THUMBNAIL_STORAGE_PATH string
var SLICE_TMP_STORAGE_PATH string

func prepareDirs() {
	err := os.Mkdir(GALLERY_STORAGE_PATH, os.ModeDir)
	ghost.Warn(err)

	err = os.Mkdir(SHARE_STORAGE_PATH, os.ModeDir)
	ghost.Warn(err)

	err = os.Mkdir(VIDEO_STORAGE_PATH, os.ModeDir)
	ghost.Warn(err)

	err = os.Mkdir(THUMBNAIL_STORAGE_PATH, os.ModeDir)
	ghost.Warn(err)

	err = os.Mkdir(FILE_STORAGE_PATH, os.ModeDir)
	ghost.Warn(err)

	err = os.Mkdir(SLICE_TMP_STORAGE_PATH, os.ModeDir)
	ghost.Warn(err)
}

func init() {
	if ghost.OS == "windows" {
		STORAGE_ROOT_PATH = "E:\\picasso"
	} else {
		STORAGE_ROOT_PATH = "/picasso"
	}
	FILE_STORAGE_PATH = path.Join(STORAGE_ROOT_PATH, strings.ToLower(db_file.CATEGORY_FILE))
	GALLERY_STORAGE_PATH = path.Join(STORAGE_ROOT_PATH, strings.ToLower(db_file.CATEGORY_GALLERY))
	VIDEO_STORAGE_PATH = path.Join(STORAGE_ROOT_PATH, strings.ToLower(db_file.CATEGORY_VIDEO))
	THUMBNAIL_STORAGE_PATH = path.Join(STORAGE_ROOT_PATH, "thumbnail")
	SHARE_STORAGE_PATH = path.Join(STORAGE_ROOT_PATH, "share")
	SLICE_TMP_STORAGE_PATH = path.Join(STORAGE_ROOT_PATH, "slices_tmp")

	prepareDirs()
}
