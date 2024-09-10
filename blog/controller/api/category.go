package api

import (
	"blog/controller"
	"blog/service"
	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	controller.RestController

	categoryService *service.CategoryService
}

func (c *CategoryController) OnInitController() {
	c.categoryService = service.GetService[*service.CategoryService](c.categoryService)

	c.Group.GET("", c.listCategories)
	c.Group.POST("", c.addCategory)
	c.Group.PUT(":id", c.updateCategory)
	c.Group.DELETE(":id", c.deleteCategory)
}

func (c *CategoryController) listCategories(context *gin.Context) {

}

func (c *CategoryController) addCategory(context *gin.Context) {

}

func (c *CategoryController) updateCategory(context *gin.Context) {

}

func (c *CategoryController) deleteCategory(context *gin.Context) {

}
