package html

import (
	"blog/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TagController struct {
	controller.Controller
}

func (c *TagController) OnInitController() {
	c.Group.GET("", c.getTags)
	c.Group.GET(":tagName", c.getTag)
	c.Group.GET(":tagName/:pageNumber", c.getTag)
}

func (c *TagController) getTags(context *gin.Context) {
	c.Render(context, http.StatusOK, "tags.html", nil)
}

func (c *TagController) getTag(context *gin.Context) {
	_ = context.Param("tagName")
	_ = context.Param("pageNumber")
	c.Render(context, http.StatusOK, "tag.html", nil)
}
