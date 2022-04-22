package service

type TagCreatQuery struct {
	Name   string `form:"name" binding:"min=1,max=100"`
	Status uint8  `form:"status,default=1" binding:"oneof=0 1"`
}

type TagDelQuery struct {
	Id string `form:"id" binding:"required,gte=1"`
}

type TagUpdateQuery struct {
	Name   string `form:"name" binding:"required,min=1,max=100"`
	Status string `form:"status" binding:"required,oneof=0 1"`
	Id     int    `form:"id" binding:"required,gte=1"`
}

type TagGetQuery struct {
	Id int `form:"id" binding:"required,gte=1"`
}
