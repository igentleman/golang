package v1

import (
	"github.com/gin-gonic/gin"
	"goproject/main/ginweb/pkg/app"
)

type Tag struct {
}

func NewTag() *Tag {
	return &Tag{}
}

func (this *Tag) TagCreat(r *gin.Context) {

}

func (this *Tag) TagDelete(r *gin.Context) {

}

func (this *Tag) TagUpdate(r *gin.Context) {

}

func (this *Tag) TagGet(r *gin.Context) {
	app.NewResponse(r).ToResponse("内部服务器错误")
	return
}
