package controller

import (
	"blog/bean/api_vo"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddController(engine *gin.Engine, relativePath string, controller ControllerInterface) *gin.RouterGroup {
	childGroup := engine.Group(relativePath)
	controller.initController(engine, childGroup)
	controller.OnInitController()
	return childGroup
}

type ControllerInterface interface {
	initController(engine *gin.Engine, group *gin.RouterGroup)
	OnInitController()
}

type Controller struct {
	Engine *gin.Engine
	Group  *gin.RouterGroup
}

func (c *Controller) initController(engine *gin.Engine, group *gin.RouterGroup) {
	c.Engine = engine
	c.Group = group
}

func (c *Controller) AddController(relativePath string, controller ControllerInterface) *gin.RouterGroup {
	childGroup := c.Group.Group(relativePath)
	controller.initController(c.Engine, childGroup)
	controller.OnInitController()
	return childGroup
}

func (c *Controller) Forward(context *gin.Context, url string) {
	context.Request.URL.Path = url
	c.Engine.HandleContext(context)
}

func (c *Controller) GetStringParam(context *gin.Context, key string, def string) string {
	value, exist := context.Params.Get(key)
	if !exist {
		return def
	}
	return value
}

func (c *Controller) GetIntParam(context *gin.Context, key string, def int) int {
	param, exist := context.Params.Get(key)
	if !exist {
		return def
	}
	value, err := strconv.Atoi(param)
	if err != nil {
		return def
	}
	return value
}

type RestController struct {
	Controller
}

func (c *RestController) Succeed(context *gin.Context, data any) {
	context.JSON(http.StatusOK, api_vo.RespVO[any]{Data: data})
}

func (c *RestController) Failed(context *gin.Context, code int, message string, data any) {
	context.JSON(http.StatusOK, api_vo.RespVO[any]{Status: code, Message: message, Data: data})
}
