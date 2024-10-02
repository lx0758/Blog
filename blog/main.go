package main

import (
	"blog/config"
	"blog/controller"
	"blog/controller/api"
	"blog/controller/html"
	"blog/logger"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var htmlIndexController = &html.IndexController{}
var apiIndexController = &api.IndexController{}

// go install github.com/swaggo/swag/cmd/swag@latest
// https://github.com/swaggo/swag/blob/master/operation.go
// swag init
//
// @title			个人博客
// @version		1.0
// @description	个人博客接口
// @contact.name	6x
// @contact.url	http://github.com/lx0758/
// @contact.email	lx0758@qq.com
// @host			127.0.0.1:8080
// @basePath		/
func main() {
	conf := config.Config().Server
	if conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.Default()
	engine.Use(Recovery)

	controller.AddController(engine, "", htmlIndexController)
	controller.AddController(engine, "api", apiIndexController)

	port := conf.Port
	err := engine.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		logger.Panicf("Run server error:%s\n", err)
	}
}

func Recovery(context *gin.Context) {
	defer func() {
		if context.Writer.Status() == http.StatusNotFound {
			htmlIndexController.RenderError(context, errors.New("not found"))
		} else if err := recover(); err != nil {
			htmlIndexController.RenderError(context, err.(error))
		}
	}()
	context.Next()
}
