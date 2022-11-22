package main

import (
	"flag"

	"github.com/muzi-long/go-gin-example/bootstrap"
)

var configFile = flag.String("config", "config/config.toml", "the config file")

func main() {
	flag.Parse()
	bootstrap.Run(*configFile)
}
