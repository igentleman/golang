package limiter

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

type MethodLimiter struct {
	*Limiter
}

func NewMethodLimiter() LimiteInterface {
	return &MethodLimiter{
		Limiter: &Limiter{limiterBuckets: make(map[string]*ratelimit.Bucket)},
	}
}

func (m *MethodLimiter) Key(r *gin.Context) string {
	uri := r.Request.RequestURI
	index := strings.Index(uri, "?")
	if index == -1 {
		return uri
	}
	return uri[:index]
}

func (m *MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := m.limiterBuckets[key]
	if ok {
		return bucket, true
	}
	return nil, false
}

func (m *MethodLimiter) AddBuckets(rules ...LimiterBucketRule) LimiteInterface {
	for _, v := range rules {
		_, ok := m.limiterBuckets[v.Key]
		if !ok {
			m.limiterBuckets[v.Key] = ratelimit.NewBucketWithQuantum(v.FillInterval, v.Capacity, v.Quantum)
		}
	}
	return m
}
