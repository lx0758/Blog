package service

import (
	"blog/bean/html_vo"
	"github.com/gin-gonic/gin"
	"strings"
)

type ThemeService struct {
	Service
	blogService *BlogService
}

func (s *ThemeService) OnInitService() {
	s.blogService = GetService[*BlogService](s.blogService)
}

func (s *ThemeService) Render(context *gin.Context, code int, page string, data interface{}) {
	var pageTitle string
	switch page {
	case "archive.html":
		pageTitle = "归档"
		break
	case "article.html":
		pageTitle = "文章"
		break
	case "category.html":
		pageTitle = "分类详情"
		break
	case "categories.html":
		pageTitle = "分类"
		break
	case "tag.html":
		pageTitle = "标签详情"
		break
	case "tags.html":
		pageTitle = "标签"
		break
	case "page.html":
	default:
		pageTitle = ""
		break
	}
	pageUrl := context.Request.RequestURI
	context.HTML(code, page, map[string]any{
		"Blog":            s.blogService.GetCacheBlog(),
		"Page":            page,
		"PageTitle":       pageTitle,
		"PageDescription": nil,
		"PageKeywords":    nil,
		"PageUrl":         pageUrl,
		"PageData":        data,
	})
}

func (s *ThemeService) RenderArticle(context *gin.Context, code int, page string, article *html_vo.ArticleVO) {
	pageDescription := article.Description
	pageKeywords := strings.Join(article.Tags, ",")
	pageUrl := article.GetSafeUrl()
	pageData := article
	context.HTML(code, page, map[string]any{
		"Blog":            s.blogService.GetCacheBlog(),
		"Page":            page,
		"PageTitle":       article.Title,
		"PageDescription": pageDescription,
		"PageKeywords":    pageKeywords,
		"PageUrl":         pageUrl,
		"PageData":        pageData,
	})
}
