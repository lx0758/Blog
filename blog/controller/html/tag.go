package html

import (
	"blog/controller"
	"blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TagController struct {
	controller.Controller
	*service.ThemeService
}

func (c *TagController) OnInitController() {
	c.ThemeService = service.GetService[*service.ThemeService](c.ThemeService)

	c.Group.GET("", c.getTags)
	c.Group.GET(":tagName", c.getTag)
	c.Group.GET(":tagName/:PageNum", c.getTag)
}

func (c *TagController) getTags(context *gin.Context) {
	c.Render(context, http.StatusOK, "tags.html", nil)
}

func (c *TagController) getTag(context *gin.Context) {
	_ = context.Param("tagName")
	_ = context.Param("PageNum")
	c.Render(context, http.StatusOK, "tag.html", nil)
}
