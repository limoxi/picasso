package space

type EncodedSpace struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type EncodedSpaceMember struct {
	UserId int `json:"user_id"`
	NickName string `json:"nick_name"`
	IsManager bool `json:"is_manager"`
}