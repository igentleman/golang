package dao

import "goproject/main/ginweb/internal/model"

func (d *Dao) CreateArticle(title, desc, cover_image_url, content, create_by string, state uint8) error {
	m := model.Model{
		CreatedBy: create_by,
	}
	a := model.Article{
		Model:         &m,
		Title:         title,
		Desc:          desc,
		Content:       content,
		CoverImageUrl: cover_image_url,
		State:         state,
	}
	return a.ArticleAdd(d.Engine)
}

func (d *Dao) DelArticle(id uint) error {
	a := model.Article{
		Model: &model.Model{ID: id},
	}
	return a.ArticleDel(d.Engine)
}

func (d *Dao) UpdateArticle(id uint, title, desc, cover_image_url, content, modified_by string, state uint8) error {
	a := model.Article{
		Model:         &model.Model{ID: id, ModifiedBy: modified_by},
		Title:         title,
		Desc:          desc,
		Content:       content,
		CoverImageUrl: cover_image_url,
		State:         state,
	}
	return a.ArticleUpdate(d.Engine)
}

func (d *Dao) GetArticle(id uint, title string, state uint8, page, page_size int) ([]*model.Article, error) {
	a := model.Article{
		Model: &model.Model{ID: id},
		Title: title,
		State: state,
	}
	return a.ArticleGet(d.Engine, page, page_size)
}

func (d *Dao) CountArticle(id uint, title string, state uint8) (int, error) {
	a := model.Article{
		Model: &model.Model{ID: id},
		Title: title,
		State: state,
	}
	return a.ArticleCount(d.Engine)
}
