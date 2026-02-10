package api

import (
	"blog/controller"
	"blog/service"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type CaptchaController struct {
	controller.RestController

	captchaService *service.CaptchaService
}

func (c *CaptchaController) OnInitController() {
	c.captchaService = service.GetService[*service.CaptchaService](c.captchaService)

	c.Group.GET("", c.queryCaptcha)
}

// @Summary		admin captcha
// @Description	admin captcha
// @Tags		captcha
// @Accept		*/*
// @Produce		image/png
// @Router		/api/captcha [get]
func (c *CaptchaController) queryCaptcha(context *gin.Context) {
	session := sessions.Default(context)
	captcha := c.captchaService.Generate(session, service.CAPTCHA_TYPE_LOGIN, 1*time.Minute)
	context.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate")
	context.Data(http.StatusOK, "image/png", captcha.Bytes())
}
