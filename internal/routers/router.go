package routers

import (
	"github.com/gin-gonic/gin"
	"goproject/main/ginweb/internal/routers/api/v1"
)

func NewRoutes() *gin.Engine {
	r := gin.Default()
	v := r.Group("/api/v1")
	{
		//tag
		t := v1.NewTag()
		v.GET("/tag", t.TagGet)          //获取文章标签
		v.POST("/tag", t.TagCreat)       //添加文章标签
		v.PUT("/tag:id", t.TagUpdate)    //修改文章标签
		v.DELETE("/tag:id", t.TagDelete) //删除文章标签

		//article
		a := v1.NewArticle()
		v.GET("/article", a.ArticleGet)          //获取文章
		v.POST("/article", a.ArticleCreat)       //添加文章
		v.PUT("/article:id", a.ArticleUpdate)    //修改文章
		v.DELETE("/article:id", a.ArticleDelete) //删除文章
	}
	return r
}
