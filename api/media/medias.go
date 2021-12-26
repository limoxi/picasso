package media

import (
	"github.com/limoxi/ghost"
	m_account "picasso/business/model/account"
	m_media "picasso/business/model/media"
)

type Medias struct {
	ghost.ApiTemplate

	GetParams *struct {
		Filters     ghost.Map         `form:"filters"`
		WithOptions ghost.FillOptions `form:"with_option"`
		CurPage     int               `form:"cur_page"`
		PageSize    int               `form:"page_size"`
	}
}

func (this *Medias) Resource() string {
	return "file.medias"
}

func (this *Medias) Get() ghost.Response {
	params := this.GetParams
	ctx := this.GetCtx()
	user := m_account.GetUserFromCtx(ctx)

	mediaRepository := m_media.NewMediaRepository(ctx)
	paginator := ghost.NewPaginator(params.CurPage, params.PageSize)
	mediaRepository.SetPaginator(paginator)

	medias := mediaRepository.GetPagedMediasForUser(user.Id, params.Filters)
	return ghost.NewJsonResponse(ghost.Map{
		"spaces":    m_media.NewMediaEncodeService(ctx).EncodeMany(medias),
		"page_info": paginator.ToMap(),
	})
}

func init() {
	ghost.RegisterApi(&Spaces{})
}
