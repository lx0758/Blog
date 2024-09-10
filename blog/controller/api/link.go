package api

import (
	"blog/controller"
	"blog/service"
	"github.com/gin-gonic/gin"
)

type LinkController struct {
	controller.RestController

	linkService *service.LinkService
}

func (c *LinkController) OnInitController() {
	c.linkService = service.GetService[*service.LinkService](c.linkService)

	c.Group.GET("", c.listLinks)
	c.Group.POST("", c.addLink)
	c.Group.PUT(":id", c.updateLink)
	c.Group.DELETE(":id", c.deleteLink)
}

func (c *LinkController) listLinks(context *gin.Context) {

}

func (c *LinkController) addLink(context *gin.Context) {

}

func (c *LinkController) updateLink(context *gin.Context) {

}

func (c *LinkController) deleteLink(context *gin.Context) {

}
