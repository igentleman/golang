package middleware

import (
	"goproject/main/ginweb/pkg/app"
	"goproject/main/ginweb/pkg/errcode"
	"goproject/main/ginweb/pkg/limiter"

	"github.com/gin-gonic/gin"
)

func RateLimiter(l limiter.LimiteInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := l.Key(ctx)
		bucket, b := l.GetBucket(key)
		if b {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(ctx)
				response.ToErrorResponse(errcode.TooManyRequests)
				ctx.Abort()
				return
			}
		}
		ctx.Next()
	}
}
