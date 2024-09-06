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

	c.AddController("session", &SessionController{})
}

func (c *IndexController) authorize(context *gin.Context) {
	context.Next()
}
