package api

import (
	"blog/controller"
	"blog/service"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
	controller.RestController
	*service.IndexService
}

func (c *IndexController) OnInitController() {
	c.IndexService = service.GetService[*service.IndexService](c.IndexService)

	c.Group.Use(c.authorize)
	c.Group.GET("", c.getIndex)
	c.Group.POST("", c.postIndex)
}

func (c *IndexController) authorize(context *gin.Context) {
	context.Next()
}

func (c *IndexController) getIndex(context *gin.Context) {
	c.Succeed(context, nil)
}

func (c *IndexController) postIndex(context *gin.Context) {
	c.Failed(context, 1, "Not support!", nil)
}
