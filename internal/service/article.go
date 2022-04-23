package service

type ArticleCreatQuery struct {
	Title         string `form:"title" binding:"required,min=1,max=100"`
	Desc          string `form:"desc" binding:"required,min=1,max=100"`
	CoverImageUrl string `form:"coverUrl" binding:"required,min=1,max=255"`
	Content       string `form:"content" binding:"required"`
	State         uint8  `form:"state,default=1" binding:"oneof=1 0"`
}

type ArticleDelQuery struct {
	Id int `form:"id" binding:"required,gte=1"`
}

type ArticleUpdateQuery struct {
	Id            int    `form:"id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"required,min=1,max=100"`
	Desc          string `form:"desc" binding:"required,min=1,max=100"`
	CoverImageUrl string `form:"coverUrl" binding:"required,min=1,max=255"`
	Content       string `form:"content" binding:"required"`
	State         uint8  `form:"state,default=1" binding:"oneof=1 0"`
}

type ArticleGetQuery struct {
	Id int `form:"id" binding:"required,gte=1"`
}
