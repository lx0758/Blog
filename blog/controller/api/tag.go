package api

import (
	"blog/controller"
	"blog/service"
	"github.com/gin-gonic/gin"
)

type TagController struct {
	controller.RestController

	tagService *service.TagService
}

func (c *TagController) OnInitController() {
	c.tagService = service.GetService[*service.TagService](c.tagService)

	c.Group.GET("", c.listTags)
	c.Group.POST("", c.addTag)
	c.Group.PUT(":id", c.updateTag)
	c.Group.DELETE(":id", c.deleteTag)
}

func (c *TagController) listTags(context *gin.Context) {

}

func (c *TagController) addTag(context *gin.Context) {

}

func (c *TagController) updateTag(context *gin.Context) {

}

func (c *TagController) deleteTag(context *gin.Context) {

}
