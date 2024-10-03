package controller

import (
	"blog/bean/api_vo"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
}

func (c *Controller) Error(context *gin.Context, status int, format string, args ...any) {
	var err error
	if len(args) != 0 {
		err = errors.New(fmt.Sprintf(format, args))
	} else {
		err = errors.New(format)
	}
	context.Status(status)
	panic(err)
}

func (c *Controller) ErrorNotFound(context *gin.Context) {
	c.Error(context, http.StatusNotFound, "Not found "+context.Request.RequestURI)
}

const (
	API_STATUS_SUCCEED      = http.StatusOK
	API_STATUS_UNAUTHORIZED = http.StatusUnauthorized
	API_STATUS_ERROR        = http.StatusInternalServerError
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

func (c *RestController) RestFailed(context *gin.Context, code int, message string) {
	context.JSON(http.StatusOK, api_vo.RespVO[any]{
		Status:  code,
		Message: message,
		Data:    nil,
	})
	context.Abort()
}

func (c *RestController) RestError(_ *gin.Context, format string, args ...any) {
	var err error
	if len(args) != 0 {
		err = errors.New(fmt.Sprintf(format, args))
	} else {
		err = errors.New(format)
	}
	panic(err)
}

func (c *RestController) RestValidationError(context *gin.Context, err error) {
	msg := err.Error()
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		validationError := validationErrors[0]
		msg = validationError.StructField() + " " + validationError.Tag()
	}
	panic(errors.New(fmt.Sprintf("参数验证失败: %s", msg)))
}
