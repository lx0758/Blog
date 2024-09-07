package html

import (
	"blog/bean/html_vo"
	"blog/controller"
	"blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PageController struct {
	controller.Controller
	*service.ThemeService

	articleService *service.ArticleService
}

func (c *PageController) OnInitController() {
	c.ThemeService = service.GetService[*service.ThemeService](c.ThemeService)
	c.articleService = service.GetService[*service.ArticleService](c.articleService)

	c.Group.GET("", c.getPage)
	c.Group.GET(":PageNum", c.getPage)
}

func (c *PageController) getPage(context *gin.Context) {
	pageNum := c.GetIntParam(context, "PageNum", 1)
	articles, pagination := c.articleService.PaginationArticle(pageNum)
	articleVOs := make([]html_vo.ArticleVO, 0)
	for _, article := range articles {
		articleVO := html_vo.ArticleVO{}
		articleVO.From(article)
		articleVOs = append(articleVOs, articleVO)
	}
	c.Render(context, http.StatusOK, "page.html", html_vo.PaginationVO[html_vo.ArticleVO]{
		PageNum:  pagination.PageNum,
		PageSize: pagination.PageSize,
		Size:     pagination.Size,
		Total:    pagination.Total,
		List:     articleVOs,
	})
}
