package g

import (
	"github.com/github-muzilong/go-toolkit/logger"
	"go.uber.org/zap"

	"github.com/github-muzilong/go-gin-example/config"
)

var gLogger *zap.Logger

func InitLogger(c *config.Config) {
	cfg := c.Logger
	lCfg := &logger.Config{
		Path:       cfg.Path,
		FileName:   cfg.File,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
		Level:      cfg.Level,
	}
	z, err := logger.New(lCfg)
	if err != nil {
		panic(err)
	}
	gLogger = z
}
