package space

import (
	"github.com/limoxi/ghost"
	dm_account "picasso/domain/model/account"
	dm_space "picasso/domain/model/space"
)

type Spaces struct {
	ghost.ApiTemplate
}

type spacesGetParams struct {
	Filters ghost.Map `form:"filters"`
	WithOptions ghost.FillOptions `form:"with_option"`
	CurPage int `form:"cur_page"`
	PageSize int `form:"page_size"`
}

func (this *Spaces) GetResource() string{
	return "space.spaces"
}

func (this *Spaces) Get() ghost.Response{
	var params spacesGetParams
	this.Bind(&params)
	ctx := this.GetCtx()
	user := dm_account.GetUserFromCtx(ctx)
	paginator := ghost.NewPaginator(params.CurPage, params.PageSize)
	spaces := dm_space.NewSpaceRepository(ctx).GetSpacesForUser(user, ghost.Map{}, paginator)
	return ghost.NewJsonResponse(ghost.Map{
		"spaces": dm_space.NewSpaceEncodeService(ctx).EncodeMany(spaces),
		"page_info": paginator.ToResultMap(),
	})
}

func init(){
	ghost.RegisterApi(&Spaces{})
}
