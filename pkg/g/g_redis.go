package g

import (
	"fmt"

	"github.com/go-redis/redis/v8"

	"github.com/muzi-long/go-gin-example/config"
	myRedis "github.com/muzi-long/go-gin-example/pkg/redis"
)

var gRedis *redis.Client

func InitRedis(c *config.Config) {
	cfg := c.Redis
	cRedis := &myRedis.Config{
		Address:  []string{fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)},
		Password: cfg.Password,
		Db:       0,
	}
	client, err := myRedis.New(cRedis)
	if err != nil {
		panic(err)
	}
	gRedis = client
}
