package bullet_service

import (
	"context"
	"encoding/json"
	"github.com/kim118000/game2023/pkg/constant"
	"github.com/kim118000/game2023/pkg/redis"
)

type Bullet struct {
	Ts       int64
	Content  string
	Username string
}

func (b *Bullet) AddBullet() error {
	json, err := json.Marshal(b)
	if err != nil {
		return err
	}
	redis.RDB.LPush(context.Background(), constant.BULLET_REDIS_KEY, json)
	redis.RDB.LTrim(context.Background(), constant.BULLET_REDIS_KEY, 0, 99)
	return nil
}

func GetTopBullet() []*Bullet {
	val := redis.RDB.LRange(context.Background(), constant.BULLET_REDIS_KEY, 0, 99).Val()

	bullets := make([]*Bullet, 0, len(val))
	for _, v := range val {
		var b Bullet
		json.Unmarshal([]byte(v), &b)
		bullets = append(bullets, &b)
	}
	return bullets
}
