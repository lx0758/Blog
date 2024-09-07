package html

import (
	"blog/controller"
	"blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CategoryController struct {
	controller.Controller
	*service.ThemeService
}

func (c *CategoryController) OnInitController() {
	c.ThemeService = service.GetService[*service.ThemeService](c.ThemeService)

	c.Group.GET("", c.getCategories)
	c.Group.GET(":categoryName", c.getCategory)
	c.Group.GET(":categoryName/:PageNum", c.getCategory)
}

func (c *CategoryController) getCategories(context *gin.Context) {
	_ = context.Param("categoryName")
	c.Render(context, http.StatusOK, "categories.html", nil)
}

func (c *CategoryController) getCategory(context *gin.Context) {
	_ = context.Param("categoryName")
	_ = context.Param("PageNum")
	c.Render(context, http.StatusOK, "category.html", nil)
}
