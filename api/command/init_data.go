package command

import (
	"github.com/limoxi/ghost"
)

type InitData struct {
	ghost.ApiTemplate
}

func (this *InitData) Resource() string {
	return "command.init_data"
}

func (this *InitData) Post() ghost.Response {
	return ghost.NewJsonResponse(nil)
}

func init() {
	ghost.RegisterApi(&InitData{})
}
