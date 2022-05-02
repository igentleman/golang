package api

import (
	"goproject/main/ginweb/internal/service"
	"goproject/main/ginweb/pkg/app"
	"goproject/main/ginweb/pkg/errcode"

	"github.com/gin-gonic/gin"
)

func GetAuth(r *gin.Context) {
	response := app.NewResponse(r)
	s := service.AuthGetQuery{}
	b, e := app.BindAndValid(r, &s)
	if !b {
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist.WithDetails(e.Error()))
		return
	}
	service := service.New(r.Request.Context())
	err := service.CheckAuth(&s)
	if err != nil {
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}
	token, err := app.GenerateToken(s.AppKey, s.AppSecret)
	if err != nil {
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	response.ToResponse(errcode.Success.WithDetails(token))
}
