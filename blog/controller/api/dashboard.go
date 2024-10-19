package api

import (
	"blog/bean/api_vo"
	"blog/controller"
	"blog/service"
	"github.com/gin-gonic/gin"
)

type DashboardController struct {
	controller.RestController

	blogService     *service.BlogService
	articleService  *service.ArticleService
	categoryService *service.CategoryService
	commentService  *service.CommentService
	fileService     *service.FileService
	tagService      *service.TagService
}

func (c *DashboardController) OnInitController() {
	c.blogService = service.GetService[*service.BlogService](c.blogService)
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
	articleCount, categoryCount, tagCount, fileCount, commentCount := c.blogService.GetCacheCount()
	browseCount := c.articleService.CountViews()

	newArticleItemVOs := make([]api_vo.ArticleItemVO, 0)
	for _, article := range c.articleService.ListNewArticle(10) {
		articleItemVO := api_vo.ArticleItemVO{}
		articleItemVO.From(article)
		newArticleItemVOs = append(newArticleItemVOs, articleItemVO)
	}
	newCommentVOs := make([]api_vo.CommentVO, 0)
	for _, comment := range c.commentService.ListNewComment(10) {
		commentVO := api_vo.CommentVO{}
		commentVO.From(comment)
		newCommentVOs = append(newCommentVOs, commentVO)
	}
	newFileVOs := make([]api_vo.FileVO, 0)
	for _, file := range c.fileService.ListNewFile(10) {
		fileVO := api_vo.FileVO{}
		fileVO.From(file, c.fileService.GetFileUrl(file))
		newFileVOs = append(newFileVOs, fileVO)
	}

	dashboardVO := api_vo.DashboardVO{
		ArticleCount:  articleCount,
		CategoryCount: categoryCount,
		TagCount:      tagCount,
		UploadCount:   fileCount,
		CommentCount:  commentCount,
		BrowseCount:   browseCount,
		NewArticles:   newArticleItemVOs,
		NewComments:   newCommentVOs,
		NewUploads:    newFileVOs,
	}
	c.RestSucceed(context, dashboardVO)
}
