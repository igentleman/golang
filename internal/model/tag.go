package model

import (
	"github.com/jinzhu/gorm"
)

type Tag struct {
	Model *Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (this *Tag) TableName() string { //关联表名
	return "blog_tag"
}

func (this Tag) TagList(db *gorm.DB, page, size int) ([]*Tag, error) {
	var t []*Tag
	var err error
	db.Limit(size).Select("name").Offset(page).Find(&t)
	if page > 0 && size > 0 {
		db = db.Offset(page).Limit(size)
	}
	if this.Name != "" {
		db = db.Model(this).Where("name = ?", this.Name)
	}
	db = db.Model(this).Where("state = ?", this.State)
	err = db.Model(this).Where("is_del = 0").Find(&t).Error
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (this *Tag) TagAdd(db *gorm.DB) error {
	return db.Model(this).Create(this).Error
}

func (this *Tag) TagDel(db *gorm.DB) error {
	return db.Debug().Model(this).Where("id = ? and is_del = ?", this.Model.ID, 0).Delete(this).Error
	//r := db.Debug().Model(this).Where("id = ? and is_del = ?", this.Model.ID, 0).Delete(this)
	//fmt.Println(r)
	//return nil
}

func (this *Tag) TagUpdate(db *gorm.DB) error {
	return db.Model(this).Where("id = ?", this.Model.ID).Update(this).Error
}

func (this *Tag) Count(db *gorm.DB) (int, error) {
	var num int
	if this.Name != "" {
		db = db.Model(this).Where("name = ?", this.Name)
	}
	db = db.Model(this).Where("state = ?", this.State)
	err := db.Model(this).Where("is_del = ?", 0).Count(&num).Error
	if err != nil {
		return 0, err
	}
	return num, nil
}
