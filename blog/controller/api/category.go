package api

import (
	"blog/bean/api_vo"
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

	c.Group.GET("", c.listCategory)
	c.Group.POST("", c.addCategory)
	c.Group.PUT(":id", c.updateCategory)
	c.Group.DELETE(":id", c.deleteCategory)
}

type listCategory struct {
	Name        *string `form:"name"`
	PageNum     int     `form:"pageNum" binding:"required"`
	PageSize    int     `form:"pageSize" binding:"required"`
	OrderName   *string `form:"orderName"`
	OrderMethod *string `form:"orderMethod"`
}

// @Summary		list category
// @Description	list category
// @Tags		category
// @Accept		json
// @Produce		json
// @Param		name			body	string		false	"name"
// @Param		pageNum			body	int		true	"pageNum"
// @Param		pageSize		body	int		true	"pageSize"
// @Param		orderName		body	string	false	"orderName"
// @Param		orderMethod		body	string	false	"orderMethod"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/category [GET]
func (c *CategoryController) listCategory(context *gin.Context) {
	var listCategory listCategory
	if err := context.ShouldBind(&listCategory); err != nil {
		c.RestValidationError(context, err)
	}

	pagination := c.categoryService.PaginationByAdmin(
		listCategory.Name,
		listCategory.PageNum,
		listCategory.PageSize,
		listCategory.OrderName,
		listCategory.OrderMethod,
	)
	categoryVOs := make([]api_vo.CategoryVO, 0)
	for _, category := range pagination.List {
		articleVO := api_vo.CategoryVO{}
		articleVO.From(category)
		categoryVOs = append(categoryVOs, articleVO)
	}
	paginationVO := api_vo.PaginationVO[api_vo.CategoryVO]{
		PageNum:  pagination.PageNum,
		PageSize: pagination.PageSize,
		Size:     pagination.Size,
		Total:    pagination.Total,
		List:     categoryVOs,
	}
	c.RestSucceed(context, paginationVO)
}

type addUpdateCategory struct {
	Name string `form:"name" binding:"required"`
}

// @Summary		add category
// @Description	add category
// @Tags		category
// @Accept		json
// @Produce		json
// @Param		name		body		string	true	"name"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/category [POST]
func (c *CategoryController) addCategory(context *gin.Context) {
	var addCategory addUpdateCategory
	if err := context.ShouldBind(&addCategory); err != nil {
		c.RestValidationError(context, err)
	}
	result := c.categoryService.AddByAdmin(addCategory.Name)
	if result {
		c.RestSucceed(context, nil)
	} else {
		c.RestError(context, "添加失败")
	}
}

type pathCategoryId struct {
	Id int `uri:"id" binding:"required"`
}

// @Summary		update category
// @Description	update category
// @Tags		category
// @Accept		json
// @Produce		json
// @Param		id			path	int		true	"id"
// @Param		name		body	string	true	"name"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/category [PUT]
func (c *CategoryController) updateCategory(context *gin.Context) {
	var pathCategoryId pathCategoryId
	if err := context.ShouldBindUri(&pathCategoryId); err != nil {
		c.RestValidationError(context, err)
	}
	var updateCategory addUpdateCategory
	if err := context.ShouldBind(&updateCategory); err != nil {
		c.RestValidationError(context, err)
	}
	result := c.categoryService.UpdateByAdmin(pathCategoryId.Id, updateCategory.Name)
	if result {
		c.RestSucceed(context, nil)
	} else {
		c.RestError(context, "更新失败")
	}
}

// @Summary		delete category
// @Description	delete category
// @Tags		category
// @Accept		json
// @Produce		json
// @Param		id				path	int		true	"id"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/category/{id} [DELETE]
func (c *CategoryController) deleteCategory(context *gin.Context) {
	var pathCategoryId pathCategoryId
	if err := context.ShouldBindUri(&pathCategoryId); err != nil {
		c.RestValidationError(context, err)
	}
	result := c.categoryService.DeleteByAdmin(pathCategoryId.Id)
	if result {
		c.RestSucceed(context, nil)
	} else {
		c.RestError(context, "删除失败")
	}
}
