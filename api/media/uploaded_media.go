package media

import (
	"github.com/limoxi/ghost"
	"mime/multipart"
	dm_media "picasso/domain/model/media"
	//ds_space "picasso/domain/service/space"
)

type UploadedMedia struct {
	ghost.ApiTemplate
}

func (this *UploadedMedia) Resource() string{
	return "media.uploaded_media"
}

type uploadMediasPutParams struct {
	SpaceId int `form:"space_id"`
	Hash string `form:"hash"`
	MediaType string `form:"media_type"`
	Files []*multipart.FileHeader `form:"file"`
}

func (this *UploadedMedia) Put() ghost.Response{
	ctx := this.GetCtx()
	var params uploadMediasPutParams
	this.Bind(&params)

	var uploader dm_media.Uploader
	switch params.MediaType {
	case "image":
		uploader = dm_media.NewImageUploadService(ctx)
	case "video":
		uploader = dm_media.NewImageUploadService(ctx)
	default:
		panic(ghost.NewBusinessError("不支持的媒体类型: "+params.MediaType))
	}
	//user := dm_account.GetUserFromCtx(ctx)
	//space, err := ds_space.NewSpaceValidator(ctx).Check(user, params.SpaceId)
	//if err != nil{
	//	panic(err)
	//}

	uploader.Upload(1, params.Files[0], params.Hash)
	return ghost.NewJsonResponse(nil)
}

func init(){
	ghost.RegisterApi(&UploadedMedia{})
}
