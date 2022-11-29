package route

import (
	"github.com/gin-gonic/gin"

	"github.com/muzi-long/go-gin-example/app/http/backend/controller"
	"github.com/muzi-long/go-gin-example/app/http/backend/middleware"
)

func backendRoute(app *gin.Engine) {
	// 统一前缀 /admin
	r := app.Group("/admin/")
	{
		// 登录
		r.POST("login", controller.User.Login)

		// 注销，需要认证
		auth := r.Use(middleware.Auth())
		auth.GET("logout", controller.User.Login)
	}

}
