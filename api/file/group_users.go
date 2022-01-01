package file

import (
	"github.com/limoxi/ghost"
	bm_account "picasso/business/model/account"
	bm_file "picasso/business/model/file"
)

type GroupUsers struct {
	ghost.ApiTemplate

	GetParams *struct {
		GroupId int `form:"group_id"`
	}
}

// Resource 群组中所有的用户
func (this *GroupUsers) Resource() string {
	return "file.group_users"
}

func (this *GroupUsers) Get() ghost.Response {
	ctx := this.GetCtx()
	user := bm_account.GetUserFromCtx(ctx)
	group := bm_file.NewGroupRepository(ctx).GetById(this.GetParams.GroupId)
	if group == nil {
		panic(ghost.NewBusinessError("空间不存在"))
	}
	if !group.HasUser(user) {
		panic(ghost.NewBusinessError("当前用户没有权限"))
	}

	users := group.GetUsers()
	return ghost.NewJsonResponse(ghost.Map{
		"users": bm_file.NewGroupEncodeService(ctx).EncodeManyUsers(users),
	})
}

func init() {
	ghost.RegisterApi(&GroupUsers{})
}
