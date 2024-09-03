package main

import (
	"blog/config"
	"blog/controller"
	"blog/controller/api"
	"blog/controller/html"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.With(config.Init)

	controller.AddController(engine, "", &html.IndexController{})
	controller.AddController(engine, "api", &api.IndexController{})

	port := config.Config().Server.Port
	if port == 0 {
		port = 8080
	}
	_ = engine.Run(fmt.Sprintf(":%d", port))
}
