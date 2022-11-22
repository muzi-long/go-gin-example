package g

import (
	"github.com/muzi-long/go-gin-example/pkg/config"

	c "github.com/muzi-long/go-gin-example/config"
)

var gConfig = new(c.Config)

func InitConfig(file string) {
	err := config.New(file, gConfig)
	if err != nil {
		panic(err)
	}
}
