package space

import (
	"github.com/limoxi/ghost"
	bm_account "picasso/business/model/account"
	bm_space "picasso/business/model/space"
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

func (this *Spaces) Resource() string{
	return "space.spaces"
}

func (this *Spaces) Get() ghost.Response{
	var params spacesGetParams
	this.Bind(&params)
	ctx := this.GetCtx()
	user := bm_account.GetUserFromCtx(ctx)
	spaceRepository := bm_space.NewSpaceRepository(ctx)
	paginator := ghost.NewPaginator(params.CurPage, params.PageSize)
	spaceRepository.SetPaginator(paginator)
	spaces := spaceRepository.GetSpacesForUser(user, ghost.Map{})
	return ghost.NewJsonResponse(ghost.Map{
		"spaces": bm_space.NewSpaceEncodeService(ctx).EncodeMany(spaces),
		"page_info": paginator.ToResultMap(),
	})
}

func init(){
	ghost.RegisterApi(&Spaces{})
}
