package route

import "github.com/gin-gonic/gin"

func userRouter(app *gin.Engine)  {
	router := app.Group("/user")
	{
		router.GET("/list",)
		router.POST("/list/add",)
	}
}
