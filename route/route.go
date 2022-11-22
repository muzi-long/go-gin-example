package route

import (
	"fmt"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/muzi-long/go-gin-example/pkg/g"
)

// Run 加载运行所有路由文件
func Run() {
	if g.Config().App.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	app := gin.New()
	app.Use(ginzap.Ginzap(g.Logger(), time.RFC3339, true))
	app.Use(ginzap.RecoveryWithZap(g.Logger(), true))

	g.Logger().Info("start http server", zap.Int("port", g.Config().App.Port))

	// 后台路由
	backendRoute(app)

	// 前台路由
	frontendRoute(app)

	// 启动http服务
	err := app.Run(fmt.Sprintf(":%d", g.Config().App.Port))
	if err != nil {
		panic(err)
	}
}
