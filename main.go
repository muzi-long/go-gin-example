package main

import (
	"flag"

	"go.uber.org/zap"

	"github.com/github-muzilong/go-gin-example/bootstrap"
	"github.com/github-muzilong/go-gin-example/pkg/g"
)

var configFile string

func main() {
	flag.StringVar(&configFile, "config", "config/config.toml", "the config file")
	flag.Parse()
	bootstrap.Run(configFile)
	g.Logger().Info("start http server", zap.Int("port", g.Config().App.Port))
}
