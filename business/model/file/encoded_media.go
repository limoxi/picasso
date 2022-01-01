package file

type EncodedFile struct {
	GroupId       int    `json:"group_id"`
	Type          int    `json:"type"`
	SimpleHash    string `json:"simple_hash"`
	FullHash      string `json:"full_hash"`
	ThumbnailPath string `json:"thumbnail_path"`
	StoragePath   string `json:"storage_path"`
	Status        int    `json:"status"`
	Metadata      string `json:"metadata"`
	ShootTime     string `json:"shoot_time"`
	ShootLocation string `json:"shoot_location"`
	Size          int64  `json:"size"`
	Duration      int    `json:"duration"`
}
