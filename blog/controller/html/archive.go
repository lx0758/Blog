package html

import (
	"blog/bean/html_vo"
	"blog/controller"
	"blog/service"

	"github.com/gin-gonic/gin"
)

type ArchiveController struct {
	controller.Controller
	*service.ThemeService

	articleService *service.ArticleService
}

func (c *ArchiveController) OnInitController() {
	c.ThemeService = service.GetService[*service.ThemeService](c.ThemeService)
	c.articleService = service.GetService[*service.ArticleService](c.articleService)

	c.Group.GET("/", c.getArchive)
	c.Group.GET("/:pageNum", c.getArchive)
}

func (c *ArchiveController) getArchive(context *gin.Context) {
	pageNum, err := c.GetPathInt(context, "pageNum", 1)
	if err != nil {
		c.ErrorNotFound(context)
	}
	pagination := c.articleService.PaginationArticleByArchive(pageNum)
	articleItemVOs := make([]html_vo.ArticleItemVO, 0)
	for _, article := range pagination.List {
		articleItemVO := html_vo.ArticleItemVO{}
		articleItemVO.FromItem(article)
		articleItemVOs = append(articleItemVOs, articleItemVO)
	}
	c.Render(context, "archive.gohtml", html_vo.PaginationVO[html_vo.ArticleItemVO]{
		Path:     "archive",
		PageNum:  pagination.PageNum,
		PageSize: pagination.PageSize,
		Size:     pagination.Size,
		Total:    pagination.Total,
		List:     articleItemVOs,
	})
}
