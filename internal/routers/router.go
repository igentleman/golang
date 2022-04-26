package routers

import (
	"goproject/main/ginweb/internal/middleware"
	v1 "goproject/main/ginweb/internal/routers/api/v1"

	"github.com/gin-gonic/gin"
	_ "github.com/go-programming-tour-book/blog-service/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v := r.Group("/api/v1", middleware.Translations())
	{
		//tag
		t := v1.NewTag()
		v.GET("/tag", t.TagGet)       //获取文章标签
		v.POST("/tag", t.TagCreat)    //添加文章标签
		v.PUT("/tag", t.TagUpdate)    //修改文章标签
		v.DELETE("/tag", t.TagDelete) //删除文章标签

		//article
		a := v1.NewArticle()
		v.GET("/article", a.ArticleGet)       //获取文章
		v.POST("/article", a.ArticleCreat)    //添加文章
		v.PUT("/article", a.ArticleUpdate)    //修改文章
		v.DELETE("/article", a.ArticleDelete) //删除文章
	}
	return r
}
