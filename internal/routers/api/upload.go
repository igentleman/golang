package api

import (
	"goproject/main/ginweb/global"
	"goproject/main/ginweb/internal/service"
	"goproject/main/ginweb/pkg/app"
	"goproject/main/ginweb/pkg/convert"
	"goproject/main/ginweb/pkg/errcode"
	"goproject/main/ginweb/pkg/upload"

	"github.com/gin-gonic/gin"
)

type Upload struct{}

func NewUpload() *Upload {
	return &Upload{}
}

func (u *Upload) UploadFile(r *gin.Context) {
	response := app.NewResponse(r)
	file, fileHeader, err := r.Request.FormFile("file")
	if err != nil {
		response.ToErrorResponse(errcode.ErrorUploadFileFail)
		return
	}
	fileType := convert.StrTo(r.PostForm("type")).MustInt()
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New(r.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf("上传文件失败，err=%v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail)
		return
	}
	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
