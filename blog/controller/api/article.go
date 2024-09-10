package api

import (
	"blog/bean/api_vo"
	"blog/bean/po"
	"blog/controller"
	"blog/service"
	"blog/util"
	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	controller.RestController

	articleService *service.ArticleService
}

func (c *ArticleController) OnInitController() {
	c.articleService = service.GetService[*service.ArticleService](c.articleService)

	c.Group.GET("", c.listArticle)
	c.Group.GET(":id", c.queryArticle)
	c.Group.POST("", c.addArticle)
	c.Group.PUT(":id", c.updateArticle)
	c.Group.PUT(":id/status", c.updateArticleStatus)
	c.Group.DELETE(":id", c.deleteArticle)
}

type listArticle struct {
	Title       *string `form:"title"`
	CategoryId  *int    `form:"categoryId"`
	Url         *string `form:"url"`
	Comment     *bool   `form:"enableComment"`
	Status      *int    `form:"status"`
	PageNum     int     `form:"pageNum" binding:"required"`
	PageSize    int     `form:"pageSize" binding:"required"`
	OrderName   *string `form:"orderName"`
	OrderMethod *string `form:"orderMethod"`
}

func (c *ArticleController) listArticle(context *gin.Context) {
	var listArticle listArticle
	if err := context.ShouldBind(&listArticle); err != nil {
		c.RestValidationError(context, err)
	}

	pagination := c.articleService.PaginationArticleByAdmin(
		listArticle.Title,
		listArticle.CategoryId,
		listArticle.Url,
		listArticle.Comment,
		listArticle.Status,
		listArticle.PageNum,
		listArticle.PageSize,
		listArticle.OrderName,
		listArticle.OrderMethod,
	)
	articleItemVOs := make([]api_vo.ArticleItemVO, 0)
	for _, article := range pagination.List {
		articleVO := api_vo.ArticleItemVO{}
		articleVO.From(article)
		articleItemVOs = append(articleItemVOs, articleVO)
	}
	paginationVO := api_vo.PaginationVO[api_vo.ArticleItemVO]{
		PageNum:  pagination.PageNum,
		PageSize: pagination.PageSize,
		Size:     pagination.Size,
		Total:    pagination.Total,
		List:     articleItemVOs,
	}
	c.RestSucceed(context, paginationVO)
}

func (c *ArticleController) queryArticle(context *gin.Context) {
	id, err := c.GetPathInt(context, "id", 0)
	if err != nil {
		c.RestError(context, "文章ID无效")
	}

	article := c.articleService.QueryArticleByAdmin(id)
	if article == nil {
		c.RestError(context, "文章不存在")
	}

	articleVO := api_vo.ArticleVO{}
	articleVO.From(*article)
	c.RestSucceed(context, articleVO)
}

type addUpdateArticle struct {
	Title         string    `form:"title" binding:"required"`
	CategoryId    int       `form:"categoryId" binding:"required"`
	Content       string    `form:"content" binding:"required"`
	Url           *string   `form:"url"`
	Weight        *int      `form:"weight"`
	EnableComment bool      `form:"enableComment" binding:"required"`
	Status        int       `form:"status" binding:"required"`
	Tags          *[]string `form:"tags"`
}

func (c *ArticleController) addArticle(context *gin.Context) {
	var updateArticle addUpdateArticle
	if err := context.ShouldBind(&updateArticle); err != nil {
		c.RestValidationError(context, err)
	}
	c.articleService.AddArticleByAdmin(
		updateArticle.Title,
		updateArticle.CategoryId,
		updateArticle.Content,
		updateArticle.Url,
		updateArticle.Weight,
		util.If(updateArticle.EnableComment, po.ARTICLE_COMMENT_ENABLE, po.ARTICLE_COMMENT_DISABLE),
		updateArticle.Status,
		updateArticle.Tags,
	)
	c.RestSucceed(context, nil)

}

func (c *ArticleController) updateArticle(context *gin.Context) {
	id, err := c.GetPathInt(context, "id", 0)
	if err != nil {
		c.RestError(context, "文章ID无效")
	}
	var updateArticle addUpdateArticle
	if err := context.ShouldBind(&updateArticle); err != nil {
		c.RestValidationError(context, err)
	}
	c.articleService.UpdateArticleByAdmin(
		id,
		updateArticle.Title,
		updateArticle.CategoryId,
		updateArticle.Content,
		updateArticle.Url,
		updateArticle.Weight,
		util.If(updateArticle.EnableComment, po.ARTICLE_COMMENT_ENABLE, po.ARTICLE_COMMENT_DISABLE),
		util.If(updateArticle.Status == po.ARTICLE_STATUS_PUBLISHED, po.ARTICLE_STATUS_PUBLISHED, po.ARTICLE_STATUS_UNPUBLISHED),
		updateArticle.Tags,
	)
	c.RestSucceed(context, nil)
}

func (c *ArticleController) updateArticleStatus(context *gin.Context) {
	id, err := c.GetPathInt(context, "id", 0)
	if err != nil {
		c.RestError(context, "文章ID无效")
	}
	var updateArticle addUpdateArticle
	if err := context.ShouldBind(&updateArticle); err != nil {
		c.RestValidationError(context, err)
	}
	c.articleService.UpdateArticleStatusByAdmin(
		id,
		util.If(updateArticle.Status == po.ARTICLE_STATUS_PUBLISHED, po.ARTICLE_STATUS_PUBLISHED, po.ARTICLE_STATUS_UNPUBLISHED),
	)
	c.RestSucceed(context, nil)
}

func (c *ArticleController) deleteArticle(context *gin.Context) {
	id, err := c.GetPathInt(context, "id", 0)
	if err != nil {
		c.RestError(context, "文章ID无效")
	}
	c.articleService.DeleteArticleByAdmin(id)
	c.RestSucceed(context, nil)
}
