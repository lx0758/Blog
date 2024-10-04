package api

import (
	"blog/bean/api_vo"
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
	var listLink listLink
	if err := context.ShouldBind(&listLink); err != nil {
		c.RestValidationError(context, err)
	}

	pagination := c.linkService.PaginationByAdmin(
		listLink.Title,
		listLink.Url,
		listLink.Status,
		listLink.PageNum,
		listLink.PageSize,
		listLink.OrderName,
		listLink.OrderMethod,
	)
	linkVOs := make([]api_vo.LinkVO, 0)
	for _, link := range pagination.List {
		linkVO := api_vo.LinkVO{}
		linkVO.From(link)
		linkVOs = append(linkVOs, linkVO)
	}
	paginationVO := api_vo.PaginationVO[api_vo.LinkVO]{
		PageNum:  pagination.PageNum,
		PageSize: pagination.PageSize,
		Size:     pagination.Size,
		Total:    pagination.Total,
		List:     linkVOs,
	}
	c.RestSucceed(context, paginationVO)
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
	var addLink addLink
	if err := context.ShouldBind(&addLink); err != nil {
		c.RestValidationError(context, err)
	}
	result := c.linkService.AddByAdmin(
		addLink.Title,
		addLink.Url,
		*addLink.Weight,
		*addLink.Status,
	)
	if result {
		c.RestSucceed(context, nil)
	} else {
		c.RestError(context, "添加失败")
	}
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
	var pathLinkId pathLinkId
	if err := context.ShouldBindUri(&pathLinkId); err != nil {
		c.RestValidationError(context, err)
	}
	var updateLink updateLink
	if err := context.ShouldBind(&updateLink); err != nil {
		c.RestValidationError(context, err)
	}
	result := c.linkService.UpdateByAdmin(
		pathLinkId.Id,
		updateLink.Title,
		updateLink.Url,
		*updateLink.Weight,
		*updateLink.Status,
	)
	if result {
		c.RestSucceed(context, nil)
	} else {
		c.RestError(context, "更新失败")
	}
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
	var pathUrlId pathUrlId
	if err := context.ShouldBindUri(&pathUrlId); err != nil {
		c.RestValidationError(context, err)
	}
	result := c.linkService.DeleteByAdmin(pathUrlId.Id)
	if result {
		c.RestSucceed(context, nil)
	} else {
		c.RestError(context, "删除失败")
	}
}
