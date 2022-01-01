package file

import (
	"github.com/limoxi/ghost"
	bm_account "picasso/business/model/account"
	"picasso/business/model/file"
)

type GroupUser struct {
	ghost.ApiTemplate

	PutParams *struct {
		GroupId int    `json:"group_id"`
		Code    string `json:"code"`
	}
}

// Resource 文件群组用户
func (this *GroupUser) Resource() string {
	return "file.group_user"
}

func (this *GroupUser) Put() ghost.Response {
	ctx := this.GetCtx()
	user := bm_account.GetUserFromCtx(ctx)
	group := file.NewGroupRepository(ctx).GetById(this.PutParams.GroupId)
	if group == nil {
		panic(ghost.NewBusinessError("空间不存在"))
	}
	group.AddUser(user, this.PutParams.Code)
	return ghost.NewJsonResponse(ghost.Map{})
}

func init() {
	ghost.RegisterApi(&GroupUser{})
}
