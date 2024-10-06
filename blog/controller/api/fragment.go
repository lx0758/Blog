package api

import (
	"blog/bean/api_vo"
	"blog/controller"
	"blog/service"
	"github.com/gin-gonic/gin"
)

type FragmentController struct {
	controller.RestController

	fragmentService *service.FragmentService
	sessionService  *service.SessionService
}

func (c *FragmentController) OnInitController() {
	c.fragmentService = service.GetService[*service.FragmentService](c.fragmentService)
	c.sessionService = service.GetService[*service.SessionService](c.sessionService)

	c.Group.GET("", c.listFragment)
	c.Group.GET(":id", c.queryFragment)
	c.Group.POST("", c.addFragment)
	c.Group.PUT(":id", c.updateFragment)
	c.Group.DELETE(":id", c.deleteFragment)
}

type listFragment struct {
	Key         *string `form:"key"`
	Content     *string `form:"content"`
	Status      *int    `form:"status"`
	PageNum     int     `form:"pageNum" binding:"required"`
	PageSize    int     `form:"pageSize" binding:"required"`
	OrderName   *string `form:"orderName"`
	OrderMethod *string `form:"orderMethod"`
}

// @Summary		list fragment
// @Description	list fragment
// @Tags		fragment
// @Accept		json
// @Produce		json
// @Param		key				body	string	false	"key"
// @Param		content			body	string	false	"content"
// @Param		status			body	int		false	"status"
// @Param		pageNum			body	int		true	"pageNum"
// @Param		pageSize		body	int		true	"pageSize"
// @Param		orderName		body	string	false	"orderName"
// @Param		orderMethod		body	string	false	"orderMethod"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/fragment [GET]
func (c *FragmentController) listFragment(context *gin.Context) {
	var listFragment listFragment
	if err := context.ShouldBind(&listFragment); err != nil {
		c.RestValidationError(context, err)
	}
	pagination := c.fragmentService.PaginationByAdmin(
		listFragment.Key,
		listFragment.Content,
		listFragment.Status,
		listFragment.PageNum,
		listFragment.PageSize,
		listFragment.OrderName,
		listFragment.OrderMethod,
	)
	fragmentVOs := make([]api_vo.FragmentVO, 0)
	for _, fragment := range pagination.List {
		fragmentVO := api_vo.FragmentVO{}
		fragmentVO.From(fragment)
		fragmentVOs = append(fragmentVOs, fragmentVO)
	}
	paginationVO := api_vo.PaginationVO[api_vo.FragmentVO]{
		PageNum:  pagination.PageNum,
		PageSize: pagination.PageSize,
		Size:     pagination.Size,
		Total:    pagination.Total,
		List:     fragmentVOs,
	}
	c.RestSucceed(context, paginationVO)
}

type pathFragmentId struct {
	Id int `uri:"id" binding:"required"`
}

// @Summary		query fragment
// @Description	query fragment
// @Tags		fragment
// @Accept		json
// @Produce		json
// @Param		id			path	int		true	"id"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/fragment/{id} [GET]
func (c *FragmentController) queryFragment(context *gin.Context) {
	var pathFragmentId pathFragmentId
	if err := context.ShouldBindUri(&pathFragmentId); err != nil {
		c.RestValidationError(context, err)
	}
	fragment := c.fragmentService.QueryFragmentById(pathFragmentId.Id)
	fragmentVO := api_vo.FragmentVO{}
	fragmentVO.From(*fragment)
	c.RestSucceed(context, fragmentVO)
}

type addFragment struct {
	Key     string `form:"key" binding:"required"`
	Content string `form:"content" binding:"required"`
	Status  *int   `form:"status" binding:"required"`
}

// @Summary		add fragment
// @Description	add fragment
// @Tags		fragment
// @Accept		json
// @Produce		json
// @Param		key				body	string	true	"key"
// @Param		content			body	string	true	"content"
// @Param		status			body	int		true	"status"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/fragment [POST]
func (c *FragmentController) addFragment(context *gin.Context) {
	var addFragment addFragment
	if err := context.ShouldBind(&addFragment); err != nil {
		c.RestValidationError(context, err)
	}
	userId := context.GetInt(KEY_USER_ID)
	fragment := c.fragmentService.AddByAdmin(
		userId,
		addFragment.Key,
		addFragment.Content,
		*addFragment.Status,
	)
	if fragment != nil {
		fragmentVO := api_vo.FragmentVO{}
		fragmentVO.From(*fragment)
		c.RestSucceed(context, fragmentVO)
	} else {
		c.RestError(context, "添加失败")
	}
}

type updateFragment struct {
	Key     string `form:"key" binding:"required"`
	Content string `form:"content" binding:"required"`
	Status  *int   `form:"status" binding:"required"`
}

// @Summary		update fragment
// @Description	update fragment
// @Tags		fragment
// @Accept		json
// @Produce		json
// @Param		id				path	int		true	"id"
// @Param		key				body	string	true	"key"
// @Param		content			body	string	true	"content"
// @Param		status			body	int		true	"status"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/fragment/{key} [PUT]
func (c *FragmentController) updateFragment(context *gin.Context) {
	var pathFragmentId pathFragmentId
	if err := context.ShouldBindUri(&pathFragmentId); err != nil {
		c.RestValidationError(context, err)
	}
	var updateFragment updateFragment
	if err := context.ShouldBind(&updateFragment); err != nil {
		c.RestValidationError(context, err)
	}
	result := c.fragmentService.UpdateByAdmin(pathFragmentId.Id, updateFragment.Key, updateFragment.Content, *updateFragment.Status)
	if result {
		c.RestSucceed(context, nil)
	} else {
		c.RestError(context, "更新失败")
	}
}

// @Summary		delete fragment
// @Description	delete fragment
// @Tags		fragment
// @Accept		json
// @Produce		json
// @Param		id				path	string		true	"id"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/fragment/{key} [DELETE]
func (c *FragmentController) deleteFragment(context *gin.Context) {
	var pathFragmentId pathFragmentId
	if err := context.ShouldBindUri(&pathFragmentId); err != nil {
		c.RestValidationError(context, err)
	}
	result := c.fragmentService.DeleteByAdmin(pathFragmentId.Id)
	if result {
		c.RestSucceed(context, nil)
	} else {
		c.RestError(context, "删除失败")
	}
}
