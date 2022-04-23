package dao

import (
	"github.com/jinzhu/gorm"
	"goproject/main/ginweb/global"
	"goproject/main/ginweb/internal/model"
)

type dao struct {
	Engine *gorm.DB
}

func New() *dao {
	return &dao{
		Engine: global.DbEngine,
	}
}

func (this *dao) TagCount(name string, state uint8) (int, error) {
	tag := model.Tag{
		Name:  name,
		State: state,
	}
	tagNum, err := tag.Count(this.Engine)
	if err != nil {
		return 0, err
	}
	return tagNum, nil
}

func (this *dao) TagList(name string, state uint8, page, size int) ([]*model.Tag, error) {
	tag := model.Tag{
		Name:  name,
		State: state,
	}
	return tag.TagList(this.Engine, page, size)
}

func (this *dao) CreateTag(name string, state uint8, createBy string) error {
	tag := model.Tag{
		Model: &model.Model{CreatedBy: createBy},
		Name:  name,
		State: state,
	}
	return tag.TagAdd(this.Engine)
}

func (this *dao) TagDel(id uint) error {
	tag := model.Tag{
		Model: &model.Model{ID: id},
	}
	return tag.TagDel(this.Engine)
}

func (this *dao) TagUpdate(name string, state uint8, modifiedOn string) error {
	tag := model.Tag{
		Model: &model.Model{ModifiedBy: modifiedOn},
		Name:  name,
		State: state,
	}
	return tag.TagUpdate(this.Engine)
}
