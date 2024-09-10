package api

import (
	"blog/controller"
	"blog/service"
	"github.com/gin-gonic/gin"
)

type FeatureController struct {
	controller.RestController

	featureService *service.FeatureService
}

func (c *FeatureController) OnInitController() {
	c.featureService = service.GetService[*service.FeatureService](c.featureService)

	c.Group.GET("smtp", c.querySmtp)
	c.Group.PUT("smtp", c.updateSmtp)
	c.Group.POST("smtp/test", c.testSmtp)
}

func (c *FeatureController) querySmtp(context *gin.Context) {

}

func (c *FeatureController) updateSmtp(context *gin.Context) {

}

func (c *FeatureController) testSmtp(context *gin.Context) {

}
