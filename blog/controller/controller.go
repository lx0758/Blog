package controller

import (
	"blog/bean/api_vo"
	"errors"
	"fmt"
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

func (c *Controller) GetPathString(context *gin.Context, key string, def string) string {
	value, exist := context.Params.Get(key)
	if !exist {
		return def
	}
	return value
}

func (c *Controller) GetPathInt(context *gin.Context, key string, def int) (int, error) {
	param, exist := context.Params.Get(key)
	if !exist {
		return def, nil
	}
	value, err := strconv.Atoi(param)
	if err != nil {
		return def, err
	}
	return value, nil
}

func (c *Controller) GetPathBool(context *gin.Context, key string, def bool) (bool, error) {
	param, exist := context.Params.Get(key)
	if !exist {
		return def, nil
	}
	value, err := strconv.ParseBool(param)
	if err != nil {
		return def, err
	}
	return value, nil
}

func (c *Controller) Forward(context *gin.Context, url string) {
	context.Request.URL.Path = url
	c.Engine.HandleContext(context)
	context.Abort()
}

func (c *Controller) ErrorNotFound(context *gin.Context) {
	c.Error(context, http.StatusNotFound, "Not found "+context.Request.RequestURI)
}

func (c *Controller) Error(context *gin.Context, status int, format string, args ...any) {
	err := errors.New(fmt.Sprintf(format, args))
	context.AbortWithError(status, err)
	panic(err)
}

const (
	API_STATUS_SUCCEED = 0
	API_STATUS_ERROR   = 1
)

type RestController struct {
	Controller
}

func (c *RestController) RestSucceed(context *gin.Context, data any) {
	context.JSON(http.StatusOK, api_vo.RespVO[any]{
		Status:  API_STATUS_SUCCEED,
		Message: "succeed",
		Data:    data,
	})
}

func (c *RestController) RestFailed(context *gin.Context, code int, message string, data any) {
	context.JSON(http.StatusOK, api_vo.RespVO[any]{
		Status:  code,
		Message: message,
		Data:    data,
	})
}
