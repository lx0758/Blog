package html

import (
	"blog/bean/html_vo"
	"blog/controller"
	"blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CategoryController struct {
	controller.Controller
	*service.ThemeService

	articleService  *service.ArticleService
	categoryService *service.CategoryService
}

func (c *CategoryController) OnInitController() {
	c.ThemeService = service.GetService[*service.ThemeService](c.ThemeService)
	c.articleService = service.GetService[*service.ArticleService](c.articleService)
	c.categoryService = service.GetService[*service.CategoryService](c.categoryService)

	c.Group.GET("/", c.getCategories)
	c.Group.GET("/:categoryName", c.getCategory)
	c.Group.GET("/:categoryName/:pageNum", c.getCategory)
}

func (c *CategoryController) getCategories(context *gin.Context) {
	categories := c.categoryService.ListCategory()
	categoryVOs := make([]html_vo.CategoryVO, 0)
	for _, article := range categories {
		articleVO := html_vo.CategoryVO{}
		articleVO.From(article)
		categoryVOs = append(categoryVOs, articleVO)
	}
	c.Render(context, http.StatusOK, "categories.gohtml", categoryVOs)
}

func (c *CategoryController) getCategory(context *gin.Context) {
	categoryName := c.GetPathString(context, "categoryName", "")
	category := c.categoryService.QueryCategory(categoryName)
	if category == nil {
		c.ErrorNotFound(context)
	}

	pageNum, err := c.GetPathInt(context, "pageNum", 1)
	if err != nil {
		c.Error(context, http.StatusBadRequest, "Wrong parameters:%s", err)
	}
	pagination := c.articleService.PaginationArticleByCategory(category.Id, pageNum)
	articleItemVOs := make([]html_vo.ArticleItemVO, 0)
	for _, article := range pagination.List {
		archiveItemVO := html_vo.ArticleItemVO{}
		archiveItemVO.FromItem(article)
		articleItemVOs = append(articleItemVOs, archiveItemVO)
	}
	c.Render(context, http.StatusOK, "category.gohtml", html_vo.PaginationVO[html_vo.ArticleItemVO]{
		Name:     category.Name,
		Path:     "category/" + categoryName,
		PageNum:  pagination.PageNum,
		PageSize: pagination.PageSize,
		Size:     pagination.Size,
		Total:    pagination.Total,
		List:     articleItemVOs,
	})
}
