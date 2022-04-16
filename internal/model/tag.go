package model

type Tag struct {
	Model  *Model
	Name   string `json:"name"`
	Status uint8  `json:"status"`
}

func (this *Tag) TableName() string { //关联表名
	return "blog_tag"
}
