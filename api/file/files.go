package file

import (
	"github.com/limoxi/ghost"
	bm_account "picasso/business/model/account"
	bm_file "picasso/business/model/file"
)

type Files struct {
	ghost.ApiTemplate

	GetParams *struct {
		Filters     ghost.Map         `form:"filters"`
		WithOptions ghost.FillOptions `form:"with_option"`
		CurPage     int               `form:"cur_page"`
		PageSize    int               `form:"page_size"`
		OrderAttrs  []string          `form:"order_attrs"`
	}
}

func (this *Files) Resource() string {
	return "file.files"
}

func (this *Files) Get() ghost.Response {
	params := this.GetParams
	ctx := this.GetCtx()
	user := bm_account.GetUserFromCtx(ctx)

	fileRepo := bm_file.NewFileRepository(ctx)
	fileRepo.SetPage(params.CurPage, params.PageSize)

	files := fileRepo.GetOrderedFilesForUser(user.Id, params.Filters, params.OrderAttrs)
	return ghost.NewJsonResponse(ghost.Map{
		"files":     bm_file.NewFileEncodeService(ctx).EncodeMany(files),
		"page_info": fileRepo.GetPaginator().ToMap(),
	})
}

func init() {
	ghost.RegisterApi(&Files{})
}
