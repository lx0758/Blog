package html

import (
	"blog/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CategoryController struct {
	controller.Controller
}

func (c *CategoryController) OnInitController() {
	c.Group.GET("", c.getCategories)
	c.Group.GET(":categoryName", c.getCategory)
	c.Group.GET(":categoryName/:pageNumber", c.getCategory)
}

func (c *CategoryController) getCategories(context *gin.Context) {
	_ = context.Param("categoryName")
	c.Render(context, http.StatusOK, "categories.html", nil)
}

func (c *CategoryController) getCategory(context *gin.Context) {
	_ = context.Param("categoryName")
	_ = context.Param("pageNumber")
	c.Render(context, http.StatusOK, "category.html", nil)
}
