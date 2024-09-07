package html

import (
	"blog/controller"
	"blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UrlController struct {
	controller.Controller

	urlService *service.UrlService
}

func (c *UrlController) OnInitController() {
	c.urlService = service.GetService[*service.UrlService](c.urlService)

	c.Group.GET("/:key", c.getUrl)
}

func (c *UrlController) getUrl(context *gin.Context) {
	key := c.GetPathString(context, "key", "")
	url := c.urlService.QueryUrl(key)
	if url == nil {
		c.ErrorNotFound(context)
	}
	context.Redirect(http.StatusTemporaryRedirect, url.Url)
}
