package admin

import (
	"goproject/main/ginweb/global"
	"goproject/main/ginweb/pkg/app"
	"goproject/main/ginweb/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Policy struct {
	Sub string //用户的角色
	Act string //请求方法
	Obj string //请求的URI
}

func NewPolicy() *Policy {
	return &Policy{}
}

func (p *Policy) PolicyAdd(c *gin.Context) {
	response := app.NewResponse(c)
	// user, isOk := c.Get("user")
	// if !isOk {
	// 	response.ToErrorResponse(errcode.ErrorAddPolicyFail)
	// 	return
	// }
	// u, _ := user.(Policy)
	if ok := global.Casbin.AddPolicy("admin11", "/admin/api/v1/policy", "POST"); !ok {
		response.ToErrorResponse(errcode.ErrorAddPolicyFail)
		return
	}
	response.ToResponse(errcode.Success)
}
