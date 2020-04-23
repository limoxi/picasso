package command

import (
	"github.com/limoxi/ghost"
)

type InitData struct {
	ghost.ApiTemplate

}

func (this *InitData) GetResource() string{
	return "command.init_data"
}

func (this *InitData) Post() ghost.Response{
	ghost.SyncDB()
	return ghost.NewJsonResponse(nil)
}

func init(){
	ghost.RegisterApi(&InitData{})
}