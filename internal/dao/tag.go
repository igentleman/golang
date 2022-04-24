package dao

import "goproject/main/ginweb/internal/model"

func (t *Dao) TagCount(name string, state uint8) (int, error) {
	tag := model.Tag{
		Name:  name,
		State: state,
	}
	tagNum, err := tag.Count(t.Engine)
	if err != nil {
		return 0, err
	}
	return tagNum, nil
}

func (t *Dao) TagList(name string, state uint8, page, size int) ([]*model.Tag, error) {
	tag := model.Tag{
		Name:  name,
		State: state,
	}
	return tag.TagList(t.Engine, page, size)
}

func (t *Dao) CreateTag(name string, state uint8, createBy string) error {
	tag := model.Tag{
		Model: &model.Model{CreatedBy: createBy},
		Name:  name,
		State: state,
	}
	return tag.TagAdd(t.Engine)
}

func (t *Dao) TagDel(id uint) error {
	tag := model.Tag{
		Model: &model.Model{ID: id},
	}
	return tag.TagDel(t.Engine)
}

func (t *Dao) TagUpdate(name string, state uint8, modifiedOn string) error {
	tag := model.Tag{
		Model: &model.Model{ModifiedBy: modifiedOn},
		Name:  name,
		State: state,
	}
	return tag.TagUpdate(t.Engine)
}
