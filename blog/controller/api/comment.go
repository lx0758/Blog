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

	c.Group.GET("", c.listComments)
	c.Group.POST("", c.addComment)
	c.Group.PUT(":id", c.updateComment)
	c.Group.DELETE(":id", c.deleteComment)
}

func (c *CommentController) listComments(context *gin.Context) {

}

func (c *CommentController) addComment(context *gin.Context) {

}

func (c *CommentController) updateComment(context *gin.Context) {

}

func (c *CommentController) deleteComment(context *gin.Context) {

}
