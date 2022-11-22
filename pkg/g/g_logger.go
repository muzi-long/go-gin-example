package g

import (
	"time"

	"go.uber.org/zap"

	"github.com/muzi-long/go-gin-example/config"
	"github.com/muzi-long/go-gin-example/pkg/logger"
)

var gLogger *zap.Logger

func InitLogger(c *config.Config) {
	cfg := c.Logger
	fileName := cfg.File
	if cfg.File == "2006-01-02.log" {
		fileName = time.Now().Format("2006-01-02") + ".log"
	}
	lCfg := &logger.Config{
		Path:       cfg.Path,
		FileName:   fileName,
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
