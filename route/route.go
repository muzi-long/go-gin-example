package route

import (
	"fmt"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"

	"github.com/github-muzilong/go-gin-example/pkg/g"
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
	err := app.Run(fmt.Sprintf(":%d", g.Config().App.Port))
	if err != nil {
		panic(err)
	}
	// 认证路由
	authRoute(app)
	// 用户路由
	userRouter(app)

}
