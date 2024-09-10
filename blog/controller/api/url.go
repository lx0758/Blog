package api

import (
	"blog/controller"
	"blog/service"
	"github.com/gin-gonic/gin"
)

type UrlController struct {
	controller.RestController

	urlService *service.UrlService
}

func (c *UrlController) OnInitController() {
	c.urlService = service.GetService[*service.UrlService](c.urlService)

	c.Group.GET("", c.listUrls)
	c.Group.POST("", c.addUrl)
	c.Group.PUT(":id", c.updateUrl)
	c.Group.DELETE(":id", c.deleteUrl)
}

func (c *UrlController) listUrls(context *gin.Context) {

}

func (c *UrlController) addUrl(context *gin.Context) {

}

func (c *UrlController) updateUrl(context *gin.Context) {

}

func (c *UrlController) deleteUrl(context *gin.Context) {

}
