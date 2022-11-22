package middleware

import (
	"github.com/gin-gonic/gin"
)

func Permission(perms ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 查询用户

		// 用户是否有该权限

		ctx.Next()
	}
}
