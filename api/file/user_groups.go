package file

import (
	"github.com/limoxi/ghost"
	m_account "picasso/business/model/account"
	bm_file "picasso/business/model/file"
)

type UserGroups struct {
	ghost.ApiTemplate

	GetParams *struct {
		Filters     ghost.Map         `form:"filters"`
		WithOptions ghost.FillOptions `form:"with_option"`
		CurPage     int               `form:"cur_page"`
		PageSize    int               `form:"page_size"`
	}
}

// Resource 当前用户所有的文件群组
func (this *UserGroups) Resource() string {
	return "file.user_groups"
}

func (this *UserGroups) Get() ghost.Response {
	params := this.GetParams
	ctx := this.GetCtx()
	user := m_account.GetUserFromCtx(ctx)
	groupRepo := bm_file.NewGroupRepository(ctx)
	groupRepo.SetPage(params.CurPage, params.PageSize)
	groups := groupRepo.GetGroupsForUser(user, ghost.Map{})
	return ghost.NewJsonResponse(ghost.Map{
		"groups":    bm_file.NewGroupEncodeService(ctx).EncodeMany(groups),
		"page_info": groupRepo.GetPaginator().ToMap(),
	})
}

func init() {
	ghost.RegisterApi(&UserGroups{})
}
