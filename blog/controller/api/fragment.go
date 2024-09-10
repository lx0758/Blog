package api

import (
	"blog/controller"
	"blog/service"
	"github.com/gin-gonic/gin"
)

type FragmentController struct {
	controller.RestController

	fragmentService *service.FragmentService
}

func (c *FragmentController) OnInitController() {
	c.fragmentService = service.GetService[*service.FragmentService](c.fragmentService)

	c.Group.GET("", c.listFragments)
	c.Group.GET(":id", c.queryFragment)
	c.Group.POST("", c.addFragment)
	c.Group.PUT(":id", c.updateFragment)
	c.Group.DELETE(":id", c.deleteFragment)
}

func (c *FragmentController) listFragments(context *gin.Context) {

}

func (c *FragmentController) queryFragment(context *gin.Context) {

}

func (c *FragmentController) addFragment(context *gin.Context) {

}

func (c *FragmentController) updateFragment(context *gin.Context) {

}

func (c *FragmentController) deleteFragment(context *gin.Context) {

}
