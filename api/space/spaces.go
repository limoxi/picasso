package space

import (
	"github.com/limoxi/ghost"
	m_account "picasso/business/model/account"
	m_space "picasso/business/model/space"
)

type Spaces struct {
	ghost.ApiTemplate

	GetParams *struct{
		Filters ghost.Map `form:"filters"`
		WithOptions ghost.FillOptions `form:"with_option"`
		CurPage int `form:"cur_page"`
		PageSize int `form:"page_size"`
	}
}

func (this *Spaces) Resource() string{
	return "space.spaces"
}

func (this *Spaces) Get() ghost.Response{
	params := this.GetParams
	ctx := this.GetCtx()
	user := m_account.GetUserFromCtx(ctx)
	spaceRepository := m_space.NewSpaceRepository(ctx)
	paginator := ghost.NewPaginator(params.CurPage, params.PageSize)
	spaceRepository.SetPaginator(paginator)
	spaces := spaceRepository.GetSpacesForUser(user, ghost.Map{})
	return ghost.NewJsonResponse(ghost.Map{
		"spaces": m_space.NewSpaceEncodeService(ctx).EncodeMany(spaces),
		"page_info": paginator.ToMap(),
	})
}

func init(){
	ghost.RegisterApi(&Spaces{})
}
