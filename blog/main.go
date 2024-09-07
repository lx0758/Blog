package main

import (
	"blog/config"
	"blog/controller"
	"blog/controller/api"
	"blog/controller/html"
	"blog/logger"
	"fmt"
	"github.com/gin-gonic/gin"
)

var htmlIndexController = &html.IndexController{}

func main() {
	conf := config.Config().Server
	if conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.Default()
	engine.Use(htmlIndexController.OnRecovery)

	controller.AddController(engine, "", htmlIndexController)
	controller.AddController(engine, "api", &api.IndexController{})

	port := conf.Port
	err := engine.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		logger.Panicf("Run server error:%s\n", err)
	}
}
