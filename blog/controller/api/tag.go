package api

import (
	"blog/bean/api_vo"
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

type listTag struct {
	Name        *string `form:"name"`
	PageNum     int     `form:"pageNum" binding:"required"`
	PageSize    int     `form:"pageSize" binding:"required"`
	OrderName   *string `form:"orderName"`
	OrderMethod *string `form:"orderMethod"`
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
	var listTag listTag
	if err := context.ShouldBind(&listTag); err != nil {
		c.RestValidationError(context, err)
	}

	pagination := c.tagService.PaginationByAdmin(
		listTag.Name,
		listTag.PageNum,
		listTag.PageSize,
		listTag.OrderName,
		listTag.OrderMethod,
	)
	tagVOs := make([]api_vo.TagVO, 0)
	for _, tag := range pagination.List {
		tagVO := api_vo.TagVO{}
		tagVO.From(tag)
		tagVOs = append(tagVOs, tagVO)
	}
	paginationVO := api_vo.PaginationVO[api_vo.TagVO]{
		PageNum:  pagination.PageNum,
		PageSize: pagination.PageSize,
		Size:     pagination.Size,
		Total:    pagination.Total,
		List:     tagVOs,
	}
	c.RestSucceed(context, paginationVO)
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
	var addTag addTag
	if err := context.ShouldBind(&addTag); err != nil {
		c.RestValidationError(context, err)
	}
	result := c.tagService.AddByAdmin(addTag.Name)
	if result {
		c.RestSucceed(context, nil)
	} else {
		c.RestError(context, "添加失败")
	}
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
	var pathTagId pathTagId
	if err := context.ShouldBindUri(&pathTagId); err != nil {
		c.RestValidationError(context, err)
	}
	var updateTag updateTag
	if err := context.ShouldBind(&updateTag); err != nil {
		c.RestValidationError(context, err)
	}
	result := c.tagService.UpdateByAdmin(pathTagId.Id, updateTag.Name)
	if result {
		c.RestSucceed(context, nil)
	} else {
		c.RestError(context, "更新失败")
	}
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
	var pathUrlId pathUrlId
	if err := context.ShouldBindUri(&pathUrlId); err != nil {
		c.RestValidationError(context, err)
	}
	result := c.tagService.DeleteByAdmin(pathUrlId.Id)
	if result {
		c.RestSucceed(context, nil)
	} else {
		c.RestError(context, "删除失败")
	}
}
