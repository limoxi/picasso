package space

import (
	"github.com/limoxi/ghost"
	bm_account "picasso/business/model/account"
	bm_space "picasso/business/model/space"
)

type Member struct {
	ghost.ApiTemplate
}

type spaceMemberPutParams struct {
	SpaceId int `json:"space_id"`
	Code string `json:"code"`
}

func (this *Member) Resource() string{
	return "space.member"
}

func (this *Member) Put() ghost.Response{
	var params spaceMemberPutParams
	this.Bind(&params)
	ctx := this.GetCtx()
	user := bm_account.GetUserFromCtx(ctx)
	space := bm_space.NewSpaceRepository(ctx).GetById(params.SpaceId)
	if space == nil{
		panic(ghost.NewBusinessError("空间不存在"))
	}
	space.AddMember(user, params.Code)
	return ghost.NewJsonResponse(ghost.Map{})
}

func init(){
	ghost.RegisterApi(&Member{})
}
