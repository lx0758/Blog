package api

import (
	"blog/bean/api_vo"
	"blog/controller"
	"blog/service"
	"github.com/gin-gonic/gin"
	"mime/multipart"
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
	Type        *string `form:"type"`
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
// @Param		type			body	string	false	"type"
// @Param		pageNum			body	int		true	"pageNum"
// @Param		pageSize		body	int		true	"pageSize"
// @Param		orderName		body	string	false	"orderName"
// @Param		orderMethod		body	string	false	"orderMethod"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/file [GET]
func (c *FileController) listFile(context *gin.Context) {
	var listFile listFile
	if err := context.ShouldBind(&listFile); err != nil {
		c.RestValidationError(context, err)
	}
	pagination := c.fileService.PaginationByAdmin(
		listFile.Name,
		listFile.Type,
		listFile.PageNum,
		listFile.PageSize,
		listFile.OrderName,
		listFile.OrderMethod,
	)
	fileVOs := make([]api_vo.FileVO, 0)
	for _, file := range pagination.List {
		fileVO := api_vo.FileVO{}
		fileVO.From(file, c.fileService.GetFileUrl(file))
		fileVOs = append(fileVOs, fileVO)
	}
	paginationVO := api_vo.PaginationVO[api_vo.FileVO]{
		PageNum:  pagination.PageNum,
		PageSize: pagination.PageSize,
		Size:     pagination.Size,
		Total:    pagination.Total,
		List:     fileVOs,
	}
	c.RestSucceed(context, paginationVO)
}

type addFile struct {
	File multipart.FileHeader `form:"file" binding:"required"`
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
	var addFile addFile
	if err := context.ShouldBind(&addFile); err != nil {
		c.RestValidationError(context, err)
	}
	userId := context.GetInt("userId")
	file := c.fileService.AddByAdmin(userId, addFile.File)
	if file != nil {
		fileVO := api_vo.FileVO{}
		fileVO.From(*file, c.fileService.GetFileUrl(*file))
		c.RestSucceed(context, fileVO)
	} else {
		c.RestError(context, "上传失败")
	}
}

type pathFileId struct {
	Id int `uri:"id" binding:"required"`
}

type updateFile struct {
	File multipart.FileHeader `form:"file" binding:"required"`
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
	var pathFileId pathFileId
	if err := context.ShouldBindUri(&pathFileId); err != nil {
		c.RestValidationError(context, err)
	}
	var updateFile updateFile
	if err := context.ShouldBind(&updateFile); err != nil {
		c.RestValidationError(context, err)
	}
	err := c.fileService.UpdateByAdmin(pathFileId.Id, updateFile.File)
	if err == nil {
		c.RestSucceed(context, nil)
	} else {
		c.RestError(context, err.Error())
	}
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
	var pathFileId pathFileId
	if err := context.ShouldBindUri(&pathFileId); err != nil {
		c.RestValidationError(context, err)
	}
	result := c.fileService.DeleteByAdmin(pathFileId.Id)
	if result {
		c.RestSucceed(context, nil)
	} else {
		c.RestError(context, "删除失败")
	}
}
