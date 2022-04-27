package v1

import (
	"goproject/main/ginweb/global"
	"goproject/main/ginweb/internal/service"
	"goproject/main/ginweb/pkg/app"
	"goproject/main/ginweb/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Article struct {
}

func NewArticle() *Article {
	return &Article{}
}

func (t *Article) ArticleCreat(r *gin.Context) {
	a := service.ArticleCreatQuery{}
	a.CreateBy = "刘粤新"
	response := app.NewResponse(r)
	b, e := app.BindAndValid(r, &a)
	if !b {
		global.Logger.Errorf("绑定创建文章结构体失败，错误原因：%v", e)
		response.ToErrorResponse(errcode.ErrorCreateArticleFail)
		return
	}
	s := service.New(r.Request.Context())
	err := s.ArticleAdd(&a)
	if err != nil {
		global.Logger.Errorf("创建文章失败，失败原因：%v", err)
		response.ToErrorResponse(errcode.ErrorCreateArticleFail.WithDetails(e.Errors()...))
		return
	}
	response.ToResponse(errcode.Success)
}

func (t *Article) ArticleDelete(r *gin.Context) {
	s := service.ArticleDelQuery{}
	response := app.NewResponse(r)
	b, e := app.BindAndValid(r, &s)
	if !b {
		global.Logger.Errorf("绑定删除文章结构体失败，错误原因：%v", e)
		response.ToErrorResponse(errcode.ErrorDeleteArticleFail)
		return
	}
	d := service.New(r.Request.Context())
	err := d.ArticleDel(&s)
	if err != nil {
		global.Logger.Errorf("删除文章结失败，错误原因：%v", err)
		response.ToErrorResponse(errcode.ErrorDeleteArticleFail)
		return
	}
	response.ToResponse(errcode.Success)
}

func (t *Article) ArticleUpdate(r *gin.Context) {
	s := service.ArticleUpdateQuery{}
	response := app.NewResponse(r)
	b, e := app.BindAndValid(r, &s)
	if !b {
		global.Logger.Errorf("绑定失败，err:%v", e)
		response.ToErrorResponse(errcode.ErrorUpdateArticleFail)
		return
	}
	d := service.New(r.Request.Context())
	err := d.ArticleUpdate(&s)
	if err != nil {
		global.Logger.Errorf("文章更新失败，err=%v", err)
		response.ToErrorResponse(errcode.ErrorUpdateArticleFail)
		return
	}
	response.ToResponse(errcode.Success)
}

func (t *Article) ArticleGet(r *gin.Context) {
	s := service.ArticleGetQuery{}
	response := app.NewResponse(r)
	b, e := app.BindAndValid(r, &s)
	if !b {
		global.Logger.Errorf("绑定失败，err:%v", e)
		response.ToErrorResponse(errcode.ErrorGetArticleListFail)
		return
	}
	pager := app.Pager{
		Page:     app.GetPage(r),
		PageSize: app.GetPageSize(r),
	}
	d := service.New(r.Request.Context())
	d.ArticleGet(&s, &pager)
}
