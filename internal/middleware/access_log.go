package middleware

import (
	"bytes"
	"fmt"
	"goproject/main/ginweb/global"
	"goproject/main/ginweb/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (a *AccessLogWriter) Write(b []byte) (int, error) {
	n, e := a.body.Write(b)
	if e != nil {
		fmt.Println("写入失败！")
		return n, e
	}
	return a.ResponseWriter.Write(b)
}

func AccessLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bodyWrite := &AccessLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: ctx.Writer,
		}
		ctx.Writer = bodyWrite
		beginTime := time.Now().Unix()

		ctx.Next()

		endTime := time.Now().Unix()
		var methodS string
		if ctx.Request.Method == "GET" {
			methodS = ctx.Request.RequestURI
		} else {
			methodS = ctx.Request.PostForm.Encode()
		}
		fields := logger.Fields{
			"request":  methodS,
			"response": bodyWrite.body.String(),
		}
		global.Logger.WithFields(fields).Infof("access log: method: %s, status_code: %d, begin_time: %d, end_time: %d",
			ctx.Request.Method,
			bodyWrite.Status(),
			beginTime,
			endTime,
		)
	}
}
