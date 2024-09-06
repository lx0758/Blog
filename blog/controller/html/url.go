package html

import (
	"blog/controller"
	"github.com/gin-gonic/gin"
)

type UrlController struct {
	controller.Controller
}

func (c *UrlController) OnInitController() {
	c.Group.GET(":key", c.getUrl)
}

func (c *UrlController) getUrl(context *gin.Context) {
	_ = context.Param("key")
}
