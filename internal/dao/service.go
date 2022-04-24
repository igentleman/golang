package dao

type serviceCreateRequest struct {
	Title string `form:"title" binding:"required,max=100"`
}
