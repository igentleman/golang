package v1

import "github.com/gin-gonic/gin"

type Article struct {
}

func NewArticle() *Article {
	return &Article{}
}

func (this *Article) ArticleCreat(r *gin.Context) {

}

func (this *Article) ArticleDelete(r *gin.Context) {

}

func (this *Article) ArticleUpdate(r *gin.Context) {

}

func (this *Article) ArticleGet(r *gin.Context) {

}
