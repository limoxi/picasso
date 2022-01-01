package file

import (
	"github.com/limoxi/ghost"
	bm_account "picasso/business/model/account"
	bm_file "picasso/business/model/file"
)

type GroupCode struct {
	ghost.ApiTemplate

	PutParams *struct {
		GroupId int `json:"group_id"`
	}
}

// Resource 邀请码
func (this *GroupCode) Resource() string {
	return "file.group_code"
}

func (this *GroupCode) Put() ghost.Response {
	ctx := this.GetCtx()
	user := bm_account.GetUserFromCtx(ctx)
	group := bm_file.NewGroupRepository(ctx).GetForUser(user, this.PutParams.GroupId)
	if group == nil {
		panic(ghost.NewBusinessError("群组不存在"))
	}
	return ghost.NewJsonResponse(ghost.Map{
		"code": group.GenCode(),
	})
}

func init() {
	ghost.RegisterApi(&GroupCode{})
}
