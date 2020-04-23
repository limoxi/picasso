package command

import (
	"github.com/limoxi/ghost"
)

type Syncdb struct {
	ghost.ApiTemplate

}

func (this *Syncdb) GetResource() string{
	return "command.syncdb"
}

func (this *Syncdb) Post() ghost.Response{
	ghost.SyncDB()
	return ghost.NewJsonResponse(nil)
}

func init(){
	ghost.RegisterApi(&Syncdb{})
}