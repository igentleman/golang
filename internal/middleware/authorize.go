package middleware

import (
	"fmt"
	"goproject/main/ginweb/global"
	"goproject/main/ginweb/pkg/app"
	"goproject/main/ginweb/pkg/errcode"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := app.NewResponse(c)
		//获取请求的url
		obj := strings.Split(c.Request.URL.RequestURI(), "?")[0]
		//获取请求方法
		act := c.Request.Method
		var tokenAlive int
		var token string
		qToken := c.DefaultQuery("token", "")
		if qToken == "" {
			hToken := c.GetHeader("token")
			if hToken == "" {
				tokenAlive = 1
			} else {
				token = hToken
			}
		} else {
			token = qToken
		}
		claims, e := app.ParseToken(token)
		if e != nil || tokenAlive == 1 {
			response.ToErrorResponse(errcode.RbacAuthorizeFail)
			c.Abort()
			return
		}
		//获取用户的角色
		sub := claims.Role
		//判断策略中是否存在
		err := global.Casbin.LoadPolicy()
		if err != nil {
			fmt.Println(err)
			return
		}
		if b := global.Casbin.Enforce(sub, obj, act); !b {
			response.ToErrorResponse(errcode.RbacAuthorizeFail)
			c.Abort()
			return
		}
		c.Next()
	}
}
