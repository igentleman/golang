package model

import (
	"goproject/main/ginweb/pkg/app"

	"github.com/jinzhu/gorm"
)

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a *Article) TableName() string {
	return "blog_article"
}

func (a *Article) ArticleAdd(db *gorm.DB) error {
	return db.Create(a).Error
}

func (a *Article) ArticleDel(db *gorm.DB) error {
	return db.Model(a).Delete(a).Error
}

func (a *Article) ArticleUpdate(db *gorm.DB) error {
	return db.Model(a).Where("id = ?", a.ID).Update(a).Error
}

func (a *Article) ArticleGet(db *gorm.DB, page, page_size int) ([]*Article, error) {
	var articles []*Article
	if page > 0 && page_size > 0 {
		db = db.Offset(app.GetPageOffset(page, page_size)).Limit(page_size)
	}
	if a.ID != 0 {
		db = db.Where("id = ?", a.ID)
	}
	if a.Title != "" {
		db = db.Where("title like ?", "%"+a.Title+"%")
	}
	if a.State != 0 {
		db = db.Where("state = ?", a.State)
	}
	err := db.Model(a).Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (a *Article) ArticleCount(db *gorm.DB) (int, error) {
	var articlesNum int
	if a.ID != 0 {
		db = db.Where("id = ?", a.ID)
	}
	if a.Title != "" {
		db = db.Where("title like ?", "%"+a.Title+"%")
	}
	if a.State != 0 {
		db = db.Where("state = ?", a.State)
	}
	err := db.Model(a).Count(&articlesNum).Error
	if err != nil {
		return 0, err
	}
	return articlesNum, nil
}
