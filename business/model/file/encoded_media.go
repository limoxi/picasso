package file

type EncodedFile struct {
	Id     int    `json:"id"`
	Type   int    `json:"type"`
	Hash   string `json:"hash"`
	Path   string `json:"path"`
	Status int    `json:"status"`

	Name             string `json:"name"`
	Size             int64  `json:"size"`
	Metadata         string `json:"metadata"`
	Thumbnail        string `json:"thumbnail"`
	LastModifiedTime string `json:"last_modified_time"`
	CreatedTime      string `json:"created_time"`
}
