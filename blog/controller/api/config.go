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

	c.Group.GET("", c.listConfigs)
	c.Group.POST("", c.addConfig)
	c.Group.PUT(":key", c.updateConfig)
	c.Group.DELETE(":key", c.deleteConfig)
}

func (c *ConfigController) listConfigs(context *gin.Context) {

}

func (c *ConfigController) addConfig(context *gin.Context) {

}

func (c *ConfigController) updateConfig(context *gin.Context) {

}

func (c *ConfigController) deleteConfig(context *gin.Context) {

}
