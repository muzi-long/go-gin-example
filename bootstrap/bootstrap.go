package bootstrap

import (
	"github.com/muzi-long/go-gin-example/pkg/g"
	"github.com/muzi-long/go-gin-example/route"
)

// Run app运行需要加载的启动项
func Run(configFile string) {
	// 初始化配置文件
	g.InitConfig(configFile)
	// 初始化日志
	g.InitLogger(g.Config())
	// 初始化mysql连接
	g.InitMysql(g.Config())
	// 初始化redis连接
	g.InitRedis(g.Config())
	// 初始化gin
	route.Run()
}
