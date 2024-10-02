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

	c.Group.GET("", c.listTag)
	c.Group.POST("", c.addTag)
	c.Group.PUT(":id", c.updateTag)
	c.Group.DELETE(":id", c.deleteTag)
}

// @Summary		list tag
// @Description	list tag
// @Tags		tag
// @Accept		json
// @Produce		json
// @Param		name			body	string	false	"name"
// @Param		pageNum			body	int		true	"pageNum"
// @Param		pageSize		body	int		true	"pageSize"
// @Param		orderName		body	string	false	"orderName"
// @Param		orderMethod		body	string	false	"orderMethod"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/tag [GET]
func (c *TagController) listTag(context *gin.Context) {

}

type addTag struct {
	Name string `form:"name" binding:"required"`
}

// @Summary		add tag
// @Description	add tag
// @Tags		tag
// @Accept		json
// @Produce		json
// @Param		name		body	string	true	"name"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/tag [POST]
func (c *TagController) addTag(context *gin.Context) {

}

type pathTagId struct {
	Id int `uri:"id" binding:"required"`
}

type updateTag struct {
	Name string `form:"name" binding:"required"`
}

// @Summary		update tag
// @Description	update tag
// @Tags		tag
// @Accept		json
// @Produce		json
// @Param		id			path	int		true	"id"
// @Param		name		body	string	true	"name"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/tag/{id} [PUT]
func (c *TagController) updateTag(context *gin.Context) {

}

// @Summary		delete tag
// @Description	delete tag
// @Tags		tag
// @Accept		json
// @Produce		json
// @Param		id			path	int		true	"id"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/tag/{id} [DELETE]
func (c *TagController) deleteTag(context *gin.Context) {

}
