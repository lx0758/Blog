package html

import (
	"blog/controller"
	"blog/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {
	controller.Controller
	*service.IndexService
}

func (c *IndexController) OnInitController() {
	c.IndexService = service.GetService[*service.IndexService](c.IndexService)

	c.Group.GET("", c.getIndex)
}

func (c *IndexController) getIndex(context *gin.Context) {
	count := c.IndexService.Increment()
	c.Render(context, http.StatusOK, "index.html", gin.H{
		"Title":   "Hello Gin!",
		"Content": fmt.Sprintf("Request count:%d\n", count),
	})
}
