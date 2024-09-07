package api

import (
	"blog/config"
	"blog/controller"
	"blog/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type IndexController struct {
	controller.RestController
	*service.BlogService
}

func (c *IndexController) OnInitController() {
	c.BlogService = service.GetService[*service.BlogService](c.BlogService)

	storeKey := config.Config().Session.StoreKey
	if storeKey == "" {
		storeKey = strconv.FormatInt(time.Now().UnixMilli(), 10)
	}
	store := cookie.NewStore([]byte(storeKey))
	c.Group.Use(sessions.Sessions("session", store))
	c.Group.Use(cors.Default())
	c.Group.Use(c.Authorize)

	c.AddController("session", &SessionController{})
}

func (c *IndexController) Authorize(context *gin.Context) {
	context.Next()
}
