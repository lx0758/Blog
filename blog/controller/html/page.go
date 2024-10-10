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

	articleService  *service.ArticleService
	fragmentService *service.FragmentService
}

func (c *PageController) OnInitController() {
	c.ThemeService = service.GetService[*service.ThemeService](c.ThemeService)
	c.articleService = service.GetService[*service.ArticleService](c.articleService)
	c.fragmentService = service.GetService[*service.FragmentService](c.fragmentService)

	c.Group.GET("/", c.getPage)
	c.Group.GET("/:pageNum", c.getPage)
}

func (c *PageController) getPage(context *gin.Context) {
	pageNum, err := c.GetPathInt(context, "pageNum", 1)
	if err != nil {
		c.Error(context, http.StatusBadRequest, "Wrong parameters:%s", err)
	}
	pagination := c.articleService.PaginationArticleByPage(pageNum)
	articleItemVOs := make([]html_vo.ArticleItemVO, 0)
	for _, article := range pagination.List {
		articleVO := html_vo.ArticleItemVO{}
		articleVO.FromPage(article, c.fragmentService)
		articleItemVOs = append(articleItemVOs, articleVO)
	}
	c.Render(context, "page.gohtml", html_vo.PaginationVO[html_vo.ArticleItemVO]{
		Path:     "page",
		PageNum:  pagination.PageNum,
		PageSize: pagination.PageSize,
		Size:     pagination.Size,
		Total:    pagination.Total,
		List:     articleItemVOs,
	})
}
