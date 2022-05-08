package middleware

import (
	"context"
	"goproject/main/ginweb/global"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func Tracer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newCtx context.Context
		var span opentracing.Span
		spanCtx, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(
				ctx.Request.Header,
			),
		)
		if err != nil {
			span, newCtx = opentracing.StartSpanFromContextWithTracer(
				ctx.Request.Context(),
				global.Trace,
				ctx.Request.URL.Path,
			)
		} else {
			span, newCtx = opentracing.StartSpanFromContextWithTracer(
				ctx.Request.Context(),
				global.Trace,
				ctx.Request.URL.Path,
				opentracing.ChildOf(spanCtx),
				opentracing.Tag{
					Key:   string(ext.Component),
					Value: "HTTP",
				},
			)
		}
		defer span.Finish()
		ctx.Request = ctx.Request.WithContext(newCtx)
		ctx.Next()
	}
}
