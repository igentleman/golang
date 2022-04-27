package service

import (
	"goproject/main/ginweb/internal/model"
	"goproject/main/ginweb/pkg/app"
)

type ArticleCreatQuery struct {
	Title         string `form:"title" binding:"required,min=1,max=100"`
	Desc          string `form:"desc" binding:"required,min=1,max=100"`
	CoverImageUrl string `form:"cover_image_url" binding:"required,min=1,max=255"`
	Content       string `form:"content" binding:"required"`
	State         uint8  `form:"state,default=1" binding:"oneof=1 0"`
	CreateBy      string `form:"create_by" binding:"required"`
}

type ArticleDelQuery struct {
	Id uint `form:"id" binding:"required,gte=1"`
}

type ArticleUpdateQuery struct {
	Id            uint   `form:"id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"required,min=1,max=100"`
	Desc          string `form:"desc" binding:"required,min=1,max=100"`
	CoverImageUrl string `form:"coverUrl" binding:"required,min=1,max=255"`
	Content       string `form:"content" binding:"required"`
	State         uint8  `form:"state,default=1" binding:"oneof=1 0"`
	ModifiedBy    string `form:"modified_by" binding:"required"`
}

type ArticleGetQuery struct {
	Id    uint   `form:"id" binding:"gte=1"`
	Title string `form:"id" binding:"min=1,max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=1 0"`
}

type ArticleCount struct {
	Id    uint   `form:"id" binding:"gte=1"`
	Title string `form:"id" binding:"min=1,max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=1 0"`
}

func (s *Service) ArticleAdd(a *ArticleCreatQuery) error {
	return s.Dao.CreateArticle(a.Title, a.Desc, a.CoverImageUrl, a.Content, a.CreateBy, a.State)
}

func (s *Service) ArticleDel(a *ArticleDelQuery) error {
	return s.Dao.DelArticle(a.Id)
}

func (s *Service) ArticleUpdate(a *ArticleUpdateQuery) error {
	return s.Dao.UpdateArticle(a.Id, a.Title, a.Desc, a.CoverImageUrl, a.Content, a.ModifiedBy, a.State)
}

func (s *Service) ArticleGet(a *ArticleGetQuery, pager *app.Pager) ([]*model.Article, error) {
	return s.Dao.GetArticle(a.Id, a.Title, a.State, pager.Page, pager.PageSize)
}

func (s *Service) ArticleCount(a *ArticleCount) (int, error) {
	return s.Dao.CountArticle(a.Id, a.Title, a.State)
}
