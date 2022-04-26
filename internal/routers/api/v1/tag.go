package v1

import (
	"goproject/main/ginweb/global"
	"goproject/main/ginweb/internal/service"
	"goproject/main/ginweb/pkg/app"
	"goproject/main/ginweb/pkg/errcode"

	"github.com/gin-gonic/gin"
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
func (t *Tag) TagCreat(r *gin.Context) {
	creatStruct := service.TagCreatQuery{}
	creatStruct.CreatedBy = "刘粤新"
	response := app.NewResponse(r)
	b, err := app.BindAndValid(r, &creatStruct)
	if !b {
		global.Logger.Errorf(err.Error())
		response.ToErrorResponse(errcode.ErrorCreateTagFail.WithDetails(err.Errors()...))
		return
	}
	s := service.New(r.Request.Context())
	e := s.TagCreate(&creatStruct)
	if e != nil {
		global.Logger.Errorf("创建标签失败，失败原因：%v", e)
		response.ToErrorResponse(errcode.ErrorCreateTagFail.WithDetails(e.Error()))
		return
	}
	response.ToResponse(errcode.Success)
}

// @Summary 删除标签
// @Produce json
// @Param id query int true "标签id"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tag{id} [delete]
func (t *Tag) TagDelete(r *gin.Context) {
	delStruct := service.TagDelQuery{}
	response := app.NewResponse(r)
	b, e := app.BindAndValid(r, &delStruct)
	if !b {
		global.Logger.Errorf("绑定标签失败，失败原因：%v", e)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}
	s := service.New(r.Request.Context())
	err := s.TagDel(&delStruct)
	if e != nil {
		global.Logger.Errorf("删除标签失败，失败原因：%v", err)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}
	response.ToResponse(errcode.Success)
}

// @Summary 更新标签
// @Produce json
// @Param id query int true "标签id"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tag{id} [delete]
func (t *Tag) TagUpdate(r *gin.Context) {
	upStruct := service.TagUpdateQuery{}
	upStruct.ModifiedOn = "刘先生"
	response := app.NewResponse(r)
	b, err := app.BindAndValid(r, &upStruct)
	if !b {
		global.Logger.Errorf("更新标签绑定出错，详情：", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail.WithDetails(err.Errors()...))
		return
	}
	s := service.New(r.Request.Context())
	e := s.TagUpdate(&upStruct)
	if e != nil {
		global.Logger.Errorf("更新标签失败！err=", e)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}
	response.ToResponse(errcode.Success)
}

// @Summary 标签列表
// @Produce json
// @Param page query int true "页码"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tag{id} [delete]
func (t *Tag) TagGet(r *gin.Context) {
	param := service.TagGetQuery{}
	response := app.NewResponse(r)
	b, err := app.BindAndValid(r, &param)
	if !b {
		global.Logger.Errorf("参数绑定出错，详情：", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail.WithDetails(err.Errors()...))
		return
	}
	s := service.New(r.Request.Context())
	pager := app.Pager{
		Page:     app.GetPage(r),
		PageSize: app.GetPageSize(r),
	}
	num, eNum := s.TagCount(&service.TagCount{Name: param.Name, State: param.State})
	if eNum != nil {
		global.Logger.Errorf("获取tag总数出错，错误原因：%v", eNum)
		response.ToErrorResponse(errcode.ErrorCountTagFail.WithDetails(eNum.Error()))
		return
	}
	tagList, e := s.TagGet(&param, &pager)
	if e != nil {
		global.Logger.Errorf("获取tag list失败，错误原因：%v", e)
		response.ToErrorResponse(errcode.ErrorGetTagListFail.WithDetails(e.Error()))
		return
	}
	response.ToResponseList(tagList, num)
}
