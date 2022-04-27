package model

import (
	"goproject/main/ginweb/pkg/app"

	"github.com/jinzhu/gorm"
)

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t *Tag) TableName() string { //关联表名
	return "blog_tag"
}

func (t Tag) TagList(db *gorm.DB, page, size int) ([]*Tag, error) {
	var tag []*Tag
	var err error
	if page > 0 && size > 0 {
		db = db.Offset(app.GetPageOffset(page, size)).Limit(size)
		// db = db.Offset(size).Limit(size)
	}
	if t.Name != "" {
		db = db.Model(t).Where("name = ?", t.Name)
	}
	db = db.Model(t).Where("state = ?", t.State)
	err = db.Model(t).Where("is_del = 0").Find(&tag).Error
	if err != nil {
		return nil, err
	}
	return tag, nil
}

func (t *Tag) TagAdd(db *gorm.DB) error {
	return db.Model(t).Create(t).Error
}

func (t *Tag) TagDel(db *gorm.DB) error {
	return db.Model(t).Where("id = ? and is_del = ?", t.Model.ID, 0).Delete(t).Error
}

func (t *Tag) TagUpdate(db *gorm.DB) error {
	return db.Model(t).Where("id = ?", t.Model.ID).Update(t).Error
}

func (t *Tag) Count(db *gorm.DB) (int, error) {
	var num int
	if t.Name != "" {
		db = db.Model(t).Where("name = ?", t.Name)
	}
	db = db.Model(t).Where("state = ?", t.State)
	err := db.Model(t).Where("is_del = ?", 0).Count(&num).Error
	if err != nil {
		return 0, err
	}
	return num, nil
}
