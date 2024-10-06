package api

import (
	"blog/bean/api_vo"
	"blog/controller"
	"blog/service"
	"github.com/gin-gonic/gin"
)

type ConfigController struct {
	controller.RestController

	configService *service.ConfigService
}

func (c *ConfigController) OnInitController() {
	c.configService = service.GetService[*service.ConfigService](c.configService)

	c.Group.GET("", c.listConfig)
	c.Group.POST("", c.addConfig)
	c.Group.PUT(":key", c.updateConfig)
	c.Group.DELETE(":key", c.deleteConfig)
}

type listConfig struct {
	Key         *string `form:"key"`
	Value       *string `form:"value"`
	Description *string `form:"description"`
	PageNum     int     `form:"pageNum" binding:"required"`
	PageSize    int     `form:"pageSize" binding:"required"`
	OrderName   *string `form:"orderName"`
	OrderMethod *string `form:"orderMethod"`
}

// @Summary		list config
// @Description	list config
// @Tags		config
// @Accept		json
// @Produce		json
// @Param		key				body	string	false	"key"
// @Param		value			body	string	false	"value"
// @Param		description		body	string	false	"description"
// @Param		pageNum			body	int		true	"pageNum"
// @Param		pageSize		body	int		true	"pageSize"
// @Param		orderName		body	string	false	"orderName"
// @Param		orderMethod		body	string	false	"orderMethod"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/config [GET]
func (c *ConfigController) listConfig(context *gin.Context) {
	var listConfig listConfig
	if err := context.ShouldBind(&listConfig); err != nil {
		c.RestValidationError(context, err)
	}
	pagination := c.configService.PaginationByAdmin(
		listConfig.Key,
		listConfig.Value,
		listConfig.Description,
		listConfig.PageNum,
		listConfig.PageSize,
		listConfig.OrderName,
		listConfig.OrderMethod,
	)
	configVOs := make([]api_vo.ConfigVO, 0)
	for _, config := range pagination.List {
		configVO := api_vo.ConfigVO{}
		configVO.From(config)
		configVOs = append(configVOs, configVO)
	}
	paginationVO := api_vo.PaginationVO[api_vo.ConfigVO]{
		PageNum:  pagination.PageNum,
		PageSize: pagination.PageSize,
		Size:     pagination.Size,
		Total:    pagination.Total,
		List:     configVOs,
	}
	c.RestSucceed(context, paginationVO)
}

type addConfig struct {
	Key         string  `form:"key" binding:"required"`
	Value       *string `form:"value"`
	Description string  `form:"description" binding:"required"`
}

// @Summary		add config
// @Description	add config
// @Tags		config
// @Accept		json
// @Produce		json
// @Param		key				body	string	true	"key"
// @Param		value			body	string	false	"value"
// @Param		description		body	string	true	"description"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/config [POST]
func (c *ConfigController) addConfig(context *gin.Context) {
	var addConfig addConfig
	if err := context.ShouldBind(&addConfig); err != nil {
		c.RestValidationError(context, err)
	}
	result := c.configService.AddByAdmin(addConfig.Key, addConfig.Value, addConfig.Description)
	if result {
		c.RestSucceed(context, nil)
	} else {
		c.RestError(context, "添加失败")
	}
}

type pathConfigKey struct {
	Key string `uri:"key" binding:"required"`
}

type updateConfig struct {
	Value       *string `form:"value"`
	Description string  `form:"description" binding:"required"`
}

// @Summary		add config
// @Description	add config
// @Tags		config
// @Accept		json
// @Produce		json
// @Param		key				path	string	true	"key"
// @Param		value			body	string	false	"value"
// @Param		description		body	string	true	"description"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/config/{key} [PUT]
func (c *ConfigController) updateConfig(context *gin.Context) {
	var pathConfigKey pathConfigKey
	if err := context.ShouldBindUri(&pathConfigKey); err != nil {
		c.RestValidationError(context, err)
	}
	var updateConfig updateConfig
	if err := context.ShouldBind(&updateConfig); err != nil {
		c.RestValidationError(context, err)
	}
	result := c.configService.UpdateByAdmin(pathConfigKey.Key, updateConfig.Description, updateConfig.Value)
	if result {
		c.RestSucceed(context, nil)
	} else {
		c.RestError(context, "更新失败")
	}
}

// @Summary		delete config
// @Description	delete config
// @Tags		config
// @Accept		json
// @Produce		json
// @Param		key				path	string	true	"key"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/config/{key} [DELETE]
func (c *ConfigController) deleteConfig(context *gin.Context) {
	var pathConfigKey pathConfigKey
	if err := context.ShouldBindUri(&pathConfigKey); err != nil {
		c.RestValidationError(context, err)
	}
	result := c.configService.DeleteByAdmin(pathConfigKey.Key)
	if result {
		c.RestSucceed(context, nil)
	} else {
		c.RestError(context, "删除失败")
	}
}
