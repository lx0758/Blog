package api

import (
	"blog/bean/api_vo"
	"blog/controller"
	"blog/service"

	"github.com/gin-gonic/gin"
)

type FeatureController struct {
	controller.RestController

	featureService *service.FeatureService
	emailService   *service.EmailService
}

func (c *FeatureController) OnInitController() {
	c.featureService = service.GetService[*service.FeatureService](c.featureService)
	c.emailService = service.GetService[*service.EmailService](c.emailService)

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

	smtpVO := api_vo.SMTPVO{}
	smtpFeature := c.featureService.QuerySmtp()
	if smtpFeature != nil {
		smtpVO.Enable = smtpFeature.Enable
		smtpVO.Hostname = smtpFeature.Hostname
		smtpVO.Port = smtpFeature.Port
		smtpVO.Ssl = smtpFeature.Ssl
		smtpVO.Username = smtpFeature.Username
		smtpVO.Password = smtpFeature.Password
		smtpVO.FromName = smtpFeature.FromName
		smtpVO.FromEmail = smtpFeature.FromEmail
	}
	c.RestSucceed(context, smtpVO)
}

type updateSmtp struct {
	Enable    *bool   `form:"enable" binding:"required"`
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
	var updateSmtp updateSmtp
	if err := context.ShouldBind(&updateSmtp); err != nil {
		c.RestValidationError(context, err)
	}

	result := c.featureService.UpdateSmtp(
		*updateSmtp.Enable,
		updateSmtp.Hostname,
		updateSmtp.Port,
		updateSmtp.Ssl,
		updateSmtp.Username,
		updateSmtp.Password,
		updateSmtp.FromName,
		updateSmtp.FromEmail,
	)
	if result {
		c.RestSucceed(context, nil)
	} else {
		c.RestError(context, "更新失败")
	}
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
	var testSmtp testSmtp
	if err := context.ShouldBind(&testSmtp); err != nil {
		c.RestValidationError(context, err)
	}

	err := c.emailService.SendTestEmail(testSmtp.Email)
	if err == nil {
		c.RestSucceed(context, nil)
	} else {
		c.RestError(context, err.Error())
	}
}
