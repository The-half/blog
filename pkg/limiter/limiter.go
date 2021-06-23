package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

type LimiterIface interface {
	//获取对应的限流器的键值对名称
	Key(c *gin.Context) string
	//获取令牌桶
	GetBucket(key string) (*ratelimit.Bucket, bool)
	//新增多个令牌桶
	AddBuckets(rules ...LimiterBucketRule) LimiterIface
}

//限流器
type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

//定义限流桶的规则
type LimiterBucketRule struct {
	//自定义键值对名称
	Key string
	//间隔多久时间放n个令牌
	FillInterval time.Duration
	//令牌桶的容量
	Capacity int64
	//每次到达间隔时间后缩放的具体令牌数量
	Quantum int64
}


