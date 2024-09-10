package api

import (
	"blog/controller"
	"blog/service"
	"github.com/gin-gonic/gin"
)

type DashboardController struct {
	controller.RestController

	articleService  *service.ArticleService
	categoryService *service.CategoryService
	commentService  *service.CommentService
	fileService     *service.FileService
	tagService      *service.TagService
}

func (c *DashboardController) OnInitController() {
	c.articleService = service.GetService[*service.ArticleService](c.articleService)
	c.categoryService = service.GetService[*service.CategoryService](c.categoryService)
	c.commentService = service.GetService[*service.CommentService](c.commentService)
	c.fileService = service.GetService[*service.FileService](c.fileService)
	c.tagService = service.GetService[*service.TagService](c.tagService)

	c.Group.GET("", c.queryDashboard)
}

func (c *DashboardController) queryDashboard(context *gin.Context) {

}
