package html

import (
	"blog/controller"
	"blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArchiveController struct {
	controller.Controller
	*service.ThemeService
}

func (c *ArchiveController) OnInitController() {
	c.ThemeService = service.GetService[*service.ThemeService](c.ThemeService)

	c.Group.GET("", c.getArchive)
	c.Group.GET(":PageNum", c.getArchive)
}

func (c *ArchiveController) getArchive(context *gin.Context) {
	_ = context.Param("PageNum")
	c.Render(context, http.StatusOK, "archive.html", nil)
}
