package api

import (
	"blog/controller"
	"blog/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	controller.RestController

	userService *service.UserService
}

func (c *UserController) OnInitController() {
	c.userService = service.GetService[*service.UserService](c.userService)

	c.Group.GET("", c.listUsers)
	c.Group.GET("profile", c.queryProfile)
	c.Group.PUT("profile", c.updateProfile)
	c.Group.PUT("password", c.updatePassword)
}

func (c *UserController) listUsers(context *gin.Context) {

}

func (c *UserController) queryProfile(context *gin.Context) {

}

func (c *UserController) updateProfile(context *gin.Context) {

}

func (c *UserController) updatePassword(context *gin.Context) {

}
