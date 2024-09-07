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

func (s *ThemeService) RenderError(context *gin.Context) {
	status := context.Writer.Status()
	if status == 404 {
		s.Render(context, http.StatusNotFound, "404.gohtml", map[string]any{
			"Url": context.Request.RequestURI,
		})
	} else {
		errs := context.Errors
		errStr := "Server internal error"
		if errs != nil {
			errStr = errs.Last().Error()
		}
		if status == http.StatusOK {
			status = http.StatusInternalServerError
		}
		s.Render(context, status, "error.gohtml", map[string]any{
			"Status": status,
			"Error":  errStr,
		})
	}
}
