package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/muzi-long/go-gin-example/app/model"
)

var User = &userController{}

type userController struct {
	UserModel model.UserModel
}

func (c *userController) Login(ctx *gin.Context) {

}

func (c *userController) Logout(ctx *gin.Context) {
	ctx.Abort()
}
