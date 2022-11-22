package route

import (
	"github.com/gin-gonic/gin"

	"github.com/muzi-long/go-gin-example/app/http/frontend/controller"
)

func frontendRoute(app *gin.Engine) {
	// 统一前缀
	r := app.Group("/api/")
	{
		// 登录
		r.POST("login", controller.User.Login())

		// 注销，需要认证

		r.GET("logout", controller.User.Login())

	}

}
