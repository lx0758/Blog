package api

import (
	"blog/controller"
	"github.com/gin-gonic/gin"
)

type SessionController struct {
	controller.RestController
}

func (c *SessionController) OnInitController() {
	c.Group.POST("", c.postSession)
	c.Group.DELETE("", c.deleteSession)
}

func (c *SessionController) postSession(context *gin.Context) {
	c.Succeed(context, nil)
}

func (c *SessionController) deleteSession(context *gin.Context) {
	c.Failed(context, 1, "Not support!", nil)
}
