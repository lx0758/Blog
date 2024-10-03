package api

import (
	"blog/controller"
	"blog/service"
	"github.com/gin-gonic/gin"
)

type CommentController struct {
	controller.RestController

	commentService *service.CommentService
}

func (c *CommentController) OnInitController() {
	c.commentService = service.GetService[*service.CommentService](c.commentService)

	c.Group.GET("", c.listComment)
	c.Group.POST("", c.addComment)
	c.Group.PUT(":id", c.updateCommentStatus)
	c.Group.DELETE(":id", c.deleteComment)
}

type listComment struct {
	ArticleId   *int    `form:"articleId"`
	Author      *string `form:"author"`
	Email       *string `form:"email"`
	Ip          *string `form:"ip"`
	Content     *string `form:"content"`
	Status      *int    `form:"status"`
	PageNum     int     `form:"pageNum" binding:"required"`
	PageSize    int     `form:"pageSize" binding:"required"`
	OrderName   *string `form:"orderName"`
	OrderMethod *string `form:"orderMethod"`
}

// @Summary		list comment
// @Description	list comment
// @Tags		comment
// @Accept		json
// @Produce		json
// @Param		articleId		body	int		false	"articleId"
// @Param		author			body	string	false	"author"
// @Param		email			body	string	false	"email"
// @Param		ip				body	string	false	"ip"
// @Param		content			body	string	false	"content"
// @Param		status			body	int		false	"status"
// @Param		pageNum			body	int		true	"pageNum"
// @Param		pageSize		body	int		true	"pageSize"
// @Param		orderName		body	string	false	"orderName"
// @Param		orderMethod		body	string	false	"orderMethod"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/comment [GET]
func (c *CommentController) listComment(context *gin.Context) {

}

type addComment struct {
	ArticleId   int    `form:"articleId" binding:"required"`
	ParentId    int    `form:"parentId" binding:"required"`
	Content     string `form:"content" binding:"required"`
	EmailNotify string `form:"emailNotify" binding:"required"`
}

// @Summary		add comment
// @Description	add comment
// @Tags		comment
// @Accept		json
// @Produce		json
// @Param		articleId	body		int		true	"articleId"
// @Param		parentId	body		int		true	"parentId"
// @Param		content		body		string	true	"content"
// @Param		emailNotify	body		bool	true	"emailNotify"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/comment [POST]
func (c *CommentController) addComment(context *gin.Context) {

}

type pathCommentId struct {
	Id int `uri:"id" binding:"required"`
}

type updateCommentStatus struct {
	Status *int `form:"status" binding:"required"`
}

// @Summary		update comment status
// @Description	update comment status
// @Tags		comment
// @Accept		json
// @Produce		json
// @Param		id			path	int		true	"id"
// @Param		status		path	int		true	"status"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/comment/{id} [PUT]
func (c *CommentController) updateCommentStatus(context *gin.Context) {

}

// @Summary		delete comment status
// @Description	delete comment status
// @Tags		comment
// @Accept		json
// @Produce		json
// @Param		id			path	int		true	"id"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/comment/{id} [DELETE]
func (c *CommentController) deleteComment(context *gin.Context) {

}
