package api

import (
	"blog/controller"
	"blog/service"
	"github.com/gin-gonic/gin"
)

type LinkController struct {
	controller.RestController

	linkService *service.LinkService
}

func (c *LinkController) OnInitController() {
	c.linkService = service.GetService[*service.LinkService](c.linkService)

	c.Group.GET("", c.listLink)
	c.Group.POST("", c.addLink)
	c.Group.PUT(":id", c.updateLink)
	c.Group.DELETE(":id", c.deleteLink)
}

type listLink struct {
	Title       *string `form:"title"`
	Url         *string `form:"url"`
	Status      *int    `form:"status"`
	PageNum     int     `form:"pageNum" binding:"required"`
	PageSize    int     `form:"pageSize" binding:"required"`
	OrderName   *string `form:"orderName"`
	OrderMethod *string `form:"orderMethod"`
}

// @Summary		list link
// @Description	list link
// @Tags		link
// @Accept		json
// @Produce		json
// @Param		title			body	string	false	"title"
// @Param		url				body	string	false	"url"
// @Param		status			body	string	false	"status"
// @Param		pageNum			body	int		true	"pageNum"
// @Param		pageSize		body	int		true	"pageSize"
// @Param		orderName		body	string	false	"orderName"
// @Param		orderMethod		body	string	false	"orderMethod"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/link [GET]
func (c *LinkController) listLink(context *gin.Context) {

}

type addLink struct {
	Title  string `form:"title" binding:"required"`
	Url    string `form:"url" binding:"required"`
	Weight *int   `form:"weight" binding:"required"`
	Status *int   `form:"status" binding:"required"`
}

// @Summary		add link
// @Description	add link
// @Tags		link
// @Accept		json
// @Produce		json
// @Param		title		body	string	true	"title"
// @Param		url			body	string	true	"url"
// @Param		weight		body	int		true	"weight"
// @Param		status		body	int		true	"description"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/link [POST]
func (c *LinkController) addLink(context *gin.Context) {

}

type pathLinkId struct {
	Id int `uri:"id" binding:"required"`
}

type updateLink struct {
	Title  string `form:"title" binding:"required"`
	Url    string `form:"url" binding:"required"`
	Weight *int   `form:"weight" binding:"required"`
	Status *int   `form:"status" binding:"required"`
}

// @Summary		update link
// @Description	update link
// @Tags		link
// @Accept		json
// @Produce		json
// @Param		id			path	int		true	"id"
// @Param		title		body	string	true	"title"
// @Param		url			body	string	true	"url"
// @Param		weight		body	int		true	"weight"
// @Param		status		body	int		true	"description"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/link/{id} [PUT]
func (c *LinkController) updateLink(context *gin.Context) {

}

// @Summary		delete link
// @Description	delete link
// @Tags		link
// @Accept		json
// @Produce		json
// @Param		id			path	int		true	"id"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/link/{id} [DELETE]
func (c *LinkController) deleteLink(context *gin.Context) {

}
