package g

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/github-muzilong/go-gin-example/config"
)

func Config() *config.Config {
	return gConfig
}

func DB() *gorm.DB {
	return gDB
}

func Logger() *zap.Logger {
	return gLogger
}

func Redis() *redis.Client {
	return gRedis
}
