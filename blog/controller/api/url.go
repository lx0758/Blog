package api

import (
	"blog/controller"
	"blog/service"
	"github.com/gin-gonic/gin"
)

type UrlController struct {
	controller.RestController

	urlService *service.UrlService
}

func (c *UrlController) OnInitController() {
	c.urlService = service.GetService[*service.UrlService](c.urlService)

	c.Group.GET("", c.listUrl)
	c.Group.POST("", c.addUrl)
	c.Group.PUT(":id", c.updateUrl)
	c.Group.DELETE(":id", c.deleteUrl)
}

type listUrl struct {
	Key         *string `form:"key"`
	Url         *string `form:"url"`
	Description *string `form:"description"`
	PageNum     int     `form:"pageNum" binding:"required"`
	PageSize    int     `form:"pageSize" binding:"required"`
	OrderName   *string `form:"orderName"`
	OrderMethod *string `form:"orderMethod"`
}

// @Summary		list url
// @Description	list url
// @Tags		url
// @Accept		json
// @Produce		json
// @Param		key				body	string	false	"key"
// @Param		url				body	string	false	"url"
// @Param		description		body	string	false	"description"
// @Param		pageNum			body	int		true	"pageNum"
// @Param		pageSize		body	int		true	"pageSize"
// @Param		orderName		body	string	false	"orderName"
// @Param		orderMethod		body	string	false	"orderMethod"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/url [GET]
func (c *UrlController) listUrl(context *gin.Context) {

}

type addUrl struct {
	Key         string `form:"key" binding:"required"`
	Url         string `form:"url" binding:"required"`
	Description string `form:"description" binding:"required"`
	Status      int    `form:"status" binding:"required"`
}

// @Summary		add url
// @Description	add url
// @Tags		url
// @Accept		json
// @Produce		json
// @Param		key				body	string	true	"key"
// @Param		url				body	string	true	"url"
// @Param		description		body	string	true	"description"
// @Param		status			body	int		true	"description"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/url [POST]
func (c *UrlController) addUrl(context *gin.Context) {

}

type pathUrlId struct {
	Id int `uri:"id" binding:"required"`
}

type updateUrl struct {
	Key         string `form:"key" binding:"required"`
	Url         string `form:"url" binding:"required"`
	Description string `form:"description" binding:"required"`
	Status      int    `form:"status" binding:"required"`
}

// @Summary		update url
// @Description	update url
// @Tags		url
// @Accept		json
// @Produce		json
// @Param		id			path	int		true	"id"
// @Param		key			body	string	true	"key"
// @Param		url			body	string	true	"url"
// @Param		description	body	string	true	"description"
// @Param		status		body	int		true	"status"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/url/{id} [PUT]

func (c *UrlController) updateUrl(context *gin.Context) {

}

// @Summary		delete url
// @Description	delete url
// @Tags		url
// @Accept		json
// @Produce		json
// @Param		id			path	int		true	"id"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/url/{id} [DELETE]
func (c *UrlController) deleteUrl(context *gin.Context) {

}
