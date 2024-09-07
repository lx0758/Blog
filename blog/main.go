package main

import (
	"blog/config"
	"blog/controller"
	"blog/controller/api"
	"blog/controller/html"
	"blog/database"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()

	engine := gin.Default()

	config.Init(engine)

	controller.AddController(engine, "", &html.IndexController{})
	controller.AddController(engine, "api", &api.IndexController{})

	port := config.Config().Server.Port
	_ = engine.Run(fmt.Sprintf(":%d", port))
}
