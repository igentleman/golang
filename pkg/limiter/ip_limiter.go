package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

type IpLimiter struct {
	*Limiter
}

func NewIPLimiter() LimiteInterface {
	return &IpLimiter{
		Limiter: &Limiter{limiterBuckets: make(map[string]*ratelimit.Bucket)},
	}
}

func (m *IpLimiter) Key(r *gin.Context) string {
	return r.ClientIP()
}

func (m *IpLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := m.limiterBuckets[key]
	if ok {
		return bucket, true
	}
	return nil, false
}

func (m *IpLimiter) AddBuckets(rules ...LimiterBucketRule) LimiteInterface {
	for _, v := range rules {
		_, ok := m.limiterBuckets[v.Key]
		if !ok {
			m.limiterBuckets[v.Key] = ratelimit.NewBucketWithQuantum(v.FillInterval, v.Capacity, v.Quantum)
		}
	}
	return m
}

func GetIp(r *gin.Context) string {
	return r.ClientIP()
}
