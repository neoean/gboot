package weApp

import (
	"time"
	"{{.Package}}/common/config"
	"{{.Package}}/common/logs"
)

type WeCache struct{}

func (cc *WeCache) Set(key string, val interface{}, timeout time.Duration) {
	redis := config.Redis
	redis.Set(key, val, timeout)
}

func (cc *WeCache) Get(key string) (interface{}, bool) {
	redis := config.Redis
	result, err := redis.Get(key).Result()
	if err != nil {
		logs.Error("wechat cache get error: %v", err)
	}
	return result, err == nil
}
