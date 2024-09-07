package api

import (
	"blog/controller"
	"blog/service"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
	controller.RestController
	*service.BlogService
}

func (c *IndexController) OnInitController() {
	c.BlogService = service.GetService[*service.BlogService](c.BlogService)

	c.Group.Use(c.authorize)

	c.AddController("session", &SessionController{})
}

func (c *IndexController) authorize(context *gin.Context) {
	context.Next()
}
