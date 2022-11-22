package controller

import (
	"github.com/gin-gonic/gin"
)

var User = &userController{}

type userController struct {
}

func (*userController) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
