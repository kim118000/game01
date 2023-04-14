package redis

import (
	"github.com/kim118000/game2023/pkg/setting"
	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func Setup() {
	RDB = redis.NewClient(&redis.Options{
		Addr:         setting.RedisSetting.Host,
		Password:     setting.RedisSetting.Password, // 没有密码，默认值
		MinIdleConns: setting.RedisSetting.MaxIdle,
		PoolSize:     setting.RedisSetting.MaxActive,
		DB:           0, // 默认DB 0
	})
}
