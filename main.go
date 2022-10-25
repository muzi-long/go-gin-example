package main

import (
	"flag"

	"github.com/github-muzilong/go-gin-example/bootstrap"
)

var configFile string

func main() {
	flag.StringVar(&configFile, "config", "config/config.toml", "the config file")
	flag.Parse()
	bootstrap.Run(configFile)
}
