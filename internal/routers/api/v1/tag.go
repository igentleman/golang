package v1

import (
	"github.com/gin-gonic/gin"
	"goproject/main/ginweb/internal/dao"
	"log"
)

type TagCreatQuery struct {
	Name   string `from:"name" binding:"min=1,max=100"`
	Status uint8  `from:"status,default=1" binding:"oneof=0 1"`
}

type TagDelQuery struct {
	Id string `from:"id" binding:"required,gte=1"`
}

type TagUpdateQuery struct {
	Name   string `from:"name" binding:"required,min=1,max=100"`
	Status string `from:"status" binding:"required,oneof=0 1"`
	Id     string `from:"id" binding:"required,gte=1"`
}

type TagGetQuery struct {
	Id string `from:"id" binding:"required,gte=1"`
}

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

// @Summary 标签列表
// @Produce json
// @Param page query int true "页码"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tag{id} [delete]
func (this *Tag) TagGet(r *gin.Context) {
	//list, err := dao.New().TagList("", 100, 2, 10)
	err := dao.New().TagDel(8)
	if err != nil {
		log.Println("错误提示：", err)
		return
	}
	//for _, v := range list {
	//	fmt.Println(v)
	//}

}
