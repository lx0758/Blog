package api

import (
	"blog/controller"
	"blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FileController struct {
	controller.RestController

	fileService *service.FileService
}

func (c *FileController) OnInitController() {
	c.fileService = service.GetService[*service.FileService](c.fileService)

	c.Group.GET("", c.listFile)
	c.Group.POST("", c.addFile)
	c.Group.PUT(":id", c.updateFile)
	c.Group.DELETE(":id", c.deleteFile)
}

type listFile struct {
	Name        *string `form:"name"`
	Type        *int    `form:"type"`
	PageNum     int     `form:"pageNum" binding:"required"`
	PageSize    int     `form:"pageSize" binding:"required"`
	OrderName   *string `form:"orderName"`
	OrderMethod *string `form:"orderMethod"`
}

// @Summary		list file
// @Description	list file
// @Tags		file
// @Accept		json
// @Produce		json
// @Param		name			body	string	false	"name"
// @Param		type			body	int		false	"type"
// @Param		pageNum			body	int		true	"pageNum"
// @Param		pageSize		body	int		true	"pageSize"
// @Param		orderName		body	string	false	"orderName"
// @Param		orderMethod		body	string	false	"orderMethod"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/file [GET]
func (c *FileController) listFile(context *gin.Context) {

}

type addFile struct {
	File http.File `form:"file" binding:"required"`
}

// @Summary		add file
// @Description	add file
// @Tags		file
// @Accept		mpfd
// @Produce		json
// @Param		file		formData	file	true	"file"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/file [POST]
func (c *FileController) addFile(context *gin.Context) {

}

type pathFileId struct {
	Id int `uri:"id" binding:"required"`
}

type updateFile struct {
	File http.File `form:"file" binding:"required"`
}

// @Summary		update file
// @Description	update file
// @Tags		file
// @Accept		mpfd
// @Produce		json
// @Param		id			path		int		true	"id"
// @Param		file		formData	file	true	"file"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/file/{id} [PUT]
func (c *FileController) updateFile(context *gin.Context) {

}

// @Summary		delete file
// @Description	delete file
// @Tags		file
// @Accept		json
// @Produce		json
// @Param		id			path	int		true	"id"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/file/{id} [DELETE]
func (c *FileController) deleteFile(context *gin.Context) {

}
