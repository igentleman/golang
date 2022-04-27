package service

import (
	"goproject/main/ginweb/internal/model"
	"goproject/main/ginweb/pkg/app"
)

type TagCreatQuery struct {
	Name      string `form:"name" binding:"required,min=1,max=100"`
	Status    uint8  `form:"status,default=1" binding:"oneof=0 1"`
	CreatedBy string `form:"created_by" binding:"required"`
}

type TagDelQuery struct {
	Id uint `form:"id" binding:"required,gte=1"`
}

type TagUpdateQuery struct {
	Name       string `form:"name" binding:"required,min=1,max=100"`
	ModifiedOn string `form:"modified_on" binding:"required"`
	State      uint8  `form:"state" binding:"required,oneof=0 1"`
	Id         uint   `form:"id" binding:"required,gte=1"`
}

type TagGetQuery struct {
	// Name  string `form:"name" binding:"min=1,max=100"`
	// State uint8  `form:"state" binding:"required,oneof=0 1"`
	Name  string `form:"name"`
	State uint8  `form:"state"`
}

type TagCount struct {
	Name  string `form:"name" binding:"min=1,max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

func (s *Service) TagCount(param *TagCount) (int, error) {
	return s.Dao.TagCount(param.Name, param.State)
}

func (s *Service) TagCreate(param *TagCreatQuery) error {
	return s.Dao.CreateTag(param.Name, param.Status, param.CreatedBy)
}

func (s *Service) TagDel(param *TagDelQuery) error {
	return s.Dao.TagDel(param.Id)
}

func (s *Service) TagUpdate(param *TagUpdateQuery) error {
	return s.Dao.TagUpdate(param.Id, param.Name, param.State, param.ModifiedOn)
}

func (s *Service) TagGet(param *TagGetQuery, pager *app.Pager) ([]*model.Tag, error) {
	return s.Dao.TagList(param.Name, param.State, pager.Page, pager.PageSize)
}
