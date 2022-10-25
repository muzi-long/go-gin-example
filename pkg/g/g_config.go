package g

import (
	"github.com/github-muzilong/go-toolkit/config"

	c "github.com/github-muzilong/go-gin-example/config"
)

var gConfig = new(c.Config)

func InitConfig(file string) {
	err := config.New(file, gConfig)
	if err != nil {
		panic(err)
	}
}
