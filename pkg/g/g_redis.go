package g

import (
	"fmt"

	"github.com/go-redis/redis/v8"

	myRedis "github.com/github-muzilong/go-toolkit/redis"

	"github.com/github-muzilong/go-gin-example/config"
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
