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

// @Summary		query smtp features
// @Description	query smtp features
// @Tags		features
// @Accept		json
// @Produce		json
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/features/smtp [GET]
func (c *FeatureController) querySmtp(context *gin.Context) {

}

type querySmtp struct {
	Enable    bool    `form:"enable" binding:"required"`
	Hostname  *string `form:"hostname"`
	Port      *int    `form:"port"`
	Ssl       *bool   `form:"ssl"`
	Username  *string `form:"username"`
	Password  *string `form:"password"`
	FromName  *string `form:"fromName"`
	FromEmail *string `form:"fromEmail"`
}

// @Summary		update smtp features
// @Description	update smtp features
// @Tags		features
// @Accept		json
// @Produce		json
// @Param		enable		body	bool	true	"enable"
// @Param		hostname	body	string	false	"hostname"
// @Param		port		body	int		false	"port"
// @Param		ssl			body	bool	false	"ssl"
// @Param		username	body	string	false	"username"
// @Param		password	body	string	false	"password"
// @Param		fromName	body	string	false	"fromName"
// @Param		fromEmail	body	string	false	"fromEmail"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/features/smtp [PUT]
func (c *FeatureController) updateSmtp(context *gin.Context) {

}

type testSmtp struct {
	Email string `form:"email" binding:"required"`
}

// @Summary		test smtp features
// @Description	test smtp features
// @Tags		features
// @Accept		json
// @Produce		json
// @Param		email	body	string	true	"email"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/features/smtp/test [POST]
func (c *FeatureController) testSmtp(context *gin.Context) {

}
