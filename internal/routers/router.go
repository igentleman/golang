package routers

import (
	"goproject/main/ginweb/global"
	"goproject/main/ginweb/internal/middleware"
	"goproject/main/ginweb/internal/routers/api"
	"goproject/main/ginweb/internal/routers/api/admin"
	v1 "goproject/main/ginweb/internal/routers/api/v1"
	"goproject/main/ginweb/pkg/limiter"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-programming-tour-book/blog-service/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
	Key:          "/auth",
	FillInterval: time.Second * 10,
	Capacity:     100,
	Quantum:      100,
})

func NewRoutes() *gin.Engine {
	r := gin.New()
	// r.Use(middleware.Cors())
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}
	mwCORS := cors.New(cors.Config{
		//准许跨域请求网站,多个使用,分开,限制使用*
		AllowOrigins: []string{"*"},
		//准许使用的请求方式
		AllowMethods: []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		//准许使用的请求表头
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
		//显示的请求表头
		ExposeHeaders: []string{"Content-Type"},
		//凭证共享,确定共享
		AllowCredentials: true,
		//容许跨域的原点网站,可以直接return true就万事大吉了
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		//超时时间设定
		MaxAge: 24 * time.Hour,
	})
	r.Use(mwCORS)

	// r.Use(middleware.Tracer()) //链路追踪
	// r.Use(middleware.RateLimiter(methodLimiters))
	// r.Use(middleware.RateLimiter(ipLimiters))
	r.Use(middleware.Translations())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	u := api.NewUpload()
	r.POST("/upload/file", u.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	// r.StaticFS("/static1", http.Dir("storage/12"))
	v := r.Group("/api/v1")
	v.POST("/login", api.GetAuth)
	// v := r.Group("/api/v1", middleware.Translations(), middleware.JWT())
	{
		//login

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

	a := r.Group("/admin/api/v1", middleware.JWT(), middleware.Authorize())
	{
		//policy
		p := admin.NewPolicy()
		a.POST("/policy", p.PolicyAdd)
	}
	return r
}
