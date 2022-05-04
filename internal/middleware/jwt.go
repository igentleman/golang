package middleware

import (
	"goproject/main/ginweb/pkg/app"
	"goproject/main/ginweb/pkg/errcode"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			token = ""
			ecode = errcode.Success
		)
		response := app.NewResponse(ctx)
		qToken := ctx.DefaultQuery("token", "")
		if qToken == "" {
			hToken := ctx.GetHeader("token")
			if hToken == "" {
				ecode = errcode.InvalidParams
			} else {
				token = hToken
			}
		} else {
			token = qToken
		}
		if token == "" {
			ctx.Abort()
			response.ToErrorResponse(ecode)
			return
		}
		_, err := app.ParseToken(token)
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				ecode = errcode.UnauthorizedTokenTimeout
			default:
				ecode = errcode.UnauthorizedTokenError
			}
		}
		if ecode != errcode.Success {
			response.ToErrorResponse(ecode)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
