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

// @Summary		query dashboard
// @Description	query dashboard
// @Tags		dashboard
// @Accept		json
// @Produce		json
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/dashboard [GET]
func (c *DashboardController) queryDashboard(context *gin.Context) {

}
