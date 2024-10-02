package api

import (
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

}

type addConfig struct {
	Key         string  `uri:"key" binding:"required"`
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

}
