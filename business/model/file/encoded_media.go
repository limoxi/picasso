package file

type EncodedFile struct {
	GroupId         int    `json:"group_id"`
	Type            int    `json:"type"`
	Hash            string `json:"hash"`
	Name            string `json:"name"`
	StorageBasePath string `json:"storage_base_path"`
	StorageDirPath  string `json:"storage_dir_path"`
	StoragePath     string `json:"storage_path"`
	ThumbnailPath   string `json:"thumbnail_path"`
	Size            int64  `json:"size"`
	Status          int    `json:"status"`
	Metadata        string `json:"metadata"`
	CreatedTime     string `json:"created_time"`
}
