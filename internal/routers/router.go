package routers

import (
	"goproject/main/ginweb/global"
	"goproject/main/ginweb/internal/middleware"
	"goproject/main/ginweb/internal/routers/api"
	v1 "goproject/main/ginweb/internal/routers/api/v1"
	"goproject/main/ginweb/pkg/limiter"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-programming-tour-book/blog-service/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
	Key:          "/auth",
	FillInterval: time.Second * 10,
	Capacity:     1,
	Quantum:      1,
})

func NewRoutes() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}
	// r.Use(middleware.Tracer()) //链路追踪
	r.Use(middleware.RateLimiter(methodLimiters))
	// r.Use(middleware.RateLimiter(ipLimiters))
	r.Use(middleware.Translations())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	u := api.NewUpload()
	r.POST("/upload/file", u.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	r.POST("/auth", api.GetAuth)
	// r.StaticFS("/static1", http.Dir("storage/12"))
	v := r.Group("/api/v1")
	// v := r.Group("/api/v1", middleware.Translations(), middleware.JWT())
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
