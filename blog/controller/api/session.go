package api

import (
	"blog/controller"
	"blog/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type SessionController struct {
	controller.RestController

	sessionService *service.SessionService
	captchaService *service.CaptchaService
}

func (c *SessionController) OnInitController() {
	c.sessionService = service.GetService[*service.SessionService](c.sessionService)
	c.captchaService = service.GetService[*service.CaptchaService](c.captchaService)

	c.Group.POST("", c.postSession)
	c.Group.DELETE("", c.deleteSession)
}

type postSession struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Captcha  string `form:"captcha" binding:"required"`
}

// @Summary		login
// @Description	login admin
// @Tags		session
// @Accept		json
// @Produce		json
// @Param		username	body	string	true	"username"
// @Param		password	body	string	true	"password"
// @Param		captcha		body	string	true	"captcha"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/session [post]
func (c *SessionController) postSession(context *gin.Context) {
	var postSession postSession
	if err := context.ShouldBind(&postSession); err != nil {
		c.RestValidationError(context, err)
	}

	session := sessions.Default(context)
	if !c.captchaService.Verify(session, service.CAPTCHA_TYPE_LOGIN, postSession.Captcha) {
		c.RestError(context, "验证码错误")
	}
	if postSession.Username == "" || postSession.Password == "" {
		c.RestError(context, "用户名或密码为空")
	}

	err := c.sessionService.Login(session, postSession.Username, postSession.Password)
	if err != nil {
		c.RestError(context, err.Error())
	}
	c.RestSucceed(context, nil)
}

// @Summary		logout
// @Description	logout admin
// @Tags		session
// @Accept		json
// @Produce		json
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/session [delete]
func (c *SessionController) deleteSession(context *gin.Context) {
	c.sessionService.Logout(sessions.Default(context))
	c.RestSucceed(context, nil)
}
