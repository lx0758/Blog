package html

import (
	"blog/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PageController struct {
	controller.Controller
}

func (c *PageController) OnInitController() {
	c.Group.GET("", c.getPage)
	c.Group.GET(":pageNumber", c.getPage)
}

func (c *PageController) getPage(context *gin.Context) {
	_ = context.Param("pageNumber")
	c.Render(context, http.StatusOK, "page.html", nil)
}
