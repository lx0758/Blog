package api

import (
	"blog/bean/api_vo"
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
	var listUrl listUrl
	if err := context.ShouldBind(&listUrl); err != nil {
		c.RestValidationError(context, err)
	}

	pagination := c.urlService.PaginationByAdmin(
		listUrl.Key,
		listUrl.Url,
		listUrl.Description,
		listUrl.PageNum,
		listUrl.PageSize,
		listUrl.OrderName,
		listUrl.OrderMethod,
	)
	urlVOs := make([]api_vo.UrlVO, 0)
	for _, url := range pagination.List {
		urlVO := api_vo.UrlVO{}
		urlVO.From(url)
		urlVOs = append(urlVOs, urlVO)
	}
	paginationVO := api_vo.PaginationVO[api_vo.UrlVO]{
		PageNum:  pagination.PageNum,
		PageSize: pagination.PageSize,
		Size:     pagination.Size,
		Total:    pagination.Total,
		List:     urlVOs,
	}
	c.RestSucceed(context, paginationVO)
}

type addUrl struct {
	Key         string `form:"key" binding:"required"`
	Url         string `form:"url" binding:"required"`
	Description string `form:"description" binding:"required"`
	Status      *int   `form:"status" binding:"required"`
}

// @Summary		add url
// @Description	add url
// @Tags		url
// @Accept		json
// @Produce		json
// @Param		key				body	string	true	"key"
// @Param		url				body	string	true	"url"
// @Param		description		body	string	true	"description"
// @Param		status			body	int		true	"status"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/url [POST]
func (c *UrlController) addUrl(context *gin.Context) {
	var addUrl addUrl
	if err := context.ShouldBind(&addUrl); err != nil {
		c.RestValidationError(context, err)
	}
	userId := context.GetInt(KEY_USER_ID)
	result := c.urlService.AddByAdmin(userId, addUrl.Key, addUrl.Url, addUrl.Description, *addUrl.Status)
	if result {
		c.RestSucceed(context, nil)
	} else {
		c.RestError(context, "添加失败")
	}
}

type pathUrlId struct {
	Id int `uri:"id" binding:"required"`
}

type updateUrl struct {
	Key         string `form:"key" binding:"required"`
	Url         string `form:"url" binding:"required"`
	Description string `form:"description" binding:"required"`
	Status      *int   `form:"status" binding:"required"`
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
	var pathUrlId pathUrlId
	if err := context.ShouldBindUri(&pathUrlId); err != nil {
		c.RestValidationError(context, err)
	}
	var updateUrl updateUrl
	if err := context.ShouldBind(&updateUrl); err != nil {
		c.RestValidationError(context, err)
	}
	result := c.urlService.UpdateByAdmin(pathUrlId.Id, updateUrl.Key, updateUrl.Url, updateUrl.Description, *updateUrl.Status)
	if result {
		c.RestSucceed(context, nil)
	} else {
		c.RestError(context, "更新失败")
	}
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
	var pathUrlId pathUrlId
	if err := context.ShouldBindUri(&pathUrlId); err != nil {
		c.RestValidationError(context, err)
	}
	result := c.urlService.DeleteByAdmin(pathUrlId.Id)
	if result {
		c.RestSucceed(context, nil)
	} else {
		c.RestError(context, "删除失败")
	}
}
