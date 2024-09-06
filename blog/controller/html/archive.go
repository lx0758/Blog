package html

import (
	"blog/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArchiveController struct {
	controller.Controller
}

func (c *ArchiveController) OnInitController() {
	c.Group.GET("", c.getArchive)
	c.Group.GET(":pageNumber", c.getArchive)
}

func (c *ArchiveController) getArchive(context *gin.Context) {
	_ = context.Param("pageNumber")
	c.Render(context, http.StatusOK, "archive.html", nil)
}
