package space

import "github.com/limoxi/ghost"

type Space struct {
	ghost.BaseModel

	Name string `gorm:"size:128"`
	UserId int
}
func (Space) TableName() string{
	return "space_space"
}


var SPACE_USER_STATUS_NORMAL = 0
var SPACE_USER_STATUS_INVITED = 1
var SPACE_USER_STATUS2TEXT = map[int]string{
	SPACE_USER_STATUS_NORMAL: "normal",
	SPACE_USER_STATUS_INVITED: "invited",
}
type SpaceHasUser struct {
	ghost.BaseModel

	SpaceId int
	UserId int
	IsManager bool
	Status int
}
func (SpaceHasUser) TableName() string{
	return "space_has_user"
}

func init(){
	ghost.RegisterDBModel(&Space{})
	ghost.RegisterDBModel(&SpaceHasUser{})
}