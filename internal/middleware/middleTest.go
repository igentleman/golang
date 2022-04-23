package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Test struct {
	Name string
}

func TestMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("这是第二个中间件")
		c.Set("name", "六月新")
		//c.Abort()
		c.Next()
		log.Println("中间件结束")
	}
}
