package controller

import (
	"blog/bean/api_vo"
	"github.com/gin-gonic/gin"
	"net/http"
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

func (c *Controller) Render(context *gin.Context, code int, page string, data any) {
	context.HTML(code, page, data)
}

type RestController struct {
	Controller
}

func (c *RestController) Succeed(context *gin.Context, data any) {
	context.JSON(http.StatusOK, api_vo.Resp[any]{Data: data})
}

func (c *RestController) Failed(context *gin.Context, code int, message string, data any) {
	context.JSON(http.StatusOK, api_vo.Resp[any]{Status: code, Message: message, Data: data})
}
