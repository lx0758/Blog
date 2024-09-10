package service

import (
	"blog/bean/html_vo"
	"github.com/gin-gonic/gin"
	"net/http"
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
	case "archive.gohtml":
		pageTitle = "归档"
	case "article.gohtml":
		pageTitle = "文章"
	case "category.gohtml":
		pageTitle = "分类详情"
	case "categories.gohtml":
		pageTitle = "分类"
	case "tag.gohtml":
		pageTitle = "标签详情"
	case "tags.gohtml":
		pageTitle = "标签"
	case "page.gohtml":
	default:
		pageTitle = ""
	}
	pageUrl := context.Request.RequestURI
	context.HTML(code, page, html_vo.ThemeVO{
		Blog:            *s.blogService.GetCacheBlog(),
		Page:            page,
		PageTitle:       pageTitle,
		PageDescription: "",
		PageKeywords:    "",
		PageUrl:         pageUrl,
		PageData:        data,
	})
}

func (s *ThemeService) RenderArticle(context *gin.Context, code int, page string, article *html_vo.ArticleVO) {
	pageDescription := article.Description
	pageKeywords := strings.Join(article.Tags, ",")
	pageUrl := article.Url
	pageData := article
	context.HTML(code, page, html_vo.ThemeVO{
		Blog:            *s.blogService.GetCacheBlog(),
		Page:            page,
		PageTitle:       article.Title,
		PageDescription: pageDescription,
		PageKeywords:    pageKeywords,
		PageUrl:         pageUrl,
		PageData:        pageData,
	})
}

func (s *ThemeService) RenderError(context *gin.Context, err error) {
	status := context.Writer.Status()
	if status == http.StatusOK {
		status = http.StatusInternalServerError
	}
	page := "error.gohtml"
	if status == http.StatusNotFound {
		page = "404.gohtml"
	}
	errStr := err.Error()
	errs := context.Errors
	if errs != nil {
		errStr = errs.Last().Error()
	}
	s.Render(context, status, page, map[string]any{
		"Status": status,
		"Url":    context.Request.RequestURI,
		"Error":  errStr,
	})
}
