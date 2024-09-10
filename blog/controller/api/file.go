package api

import (
	"blog/controller"
	"blog/service"
	"github.com/gin-gonic/gin"
)

type FileController struct {
	controller.RestController

	fileService *service.FileService
}

func (c *FileController) OnInitController() {
	c.fileService = service.GetService[*service.FileService](c.fileService)

	c.Group.GET("", c.listFiles)
	c.Group.POST("", c.addFile)
	c.Group.PUT(":id", c.updateFile)
	c.Group.DELETE(":id", c.deleteFile)
}

func (c *FileController) listFiles(context *gin.Context) {

}

func (c *FileController) addFile(context *gin.Context) {

}

func (c *FileController) updateFile(context *gin.Context) {

}

func (c *FileController) deleteFile(context *gin.Context) {

}
