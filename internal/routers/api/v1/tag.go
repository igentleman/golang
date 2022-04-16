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

// @Summary 创建标签
// @Produce json
// @Param name query string true "标签名"
// @Param Status query int true "标签状态"
// @Param created_by query string true "创建人"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tag [get]
func (this *Tag) TagCreat(r *gin.Context) {

}

// @Summary 删除标签
// @Produce json
// @Param id query int true "标签id"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tag{id} [delete]
func (this *Tag) TagDelete(r *gin.Context) {

}

// @Summary 更新标签
// @Produce json
// @Param id query int true "标签id"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tag{id} [delete]
func (this *Tag) TagUpdate(r *gin.Context) {

}

func (this *Tag) TagGet(r *gin.Context) {
	app.NewResponse(r).ToResponse("内部服务器错误")
	return
}
