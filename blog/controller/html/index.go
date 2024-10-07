package html

import (
	"blog/bean/html_vo"
	"blog/controller"
	"blog/res"
	"blog/service"
	"blog/util"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"strings"
	"time"
)

type IndexController struct {
	controller.Controller
	*service.ThemeService
	articleService *service.ArticleService

	pageController *PageController
}

func (c *IndexController) OnInitController() {
	c.ThemeService = service.GetService[*service.ThemeService](c.ThemeService)
	c.articleService = service.GetService[*service.ArticleService](c.articleService)

	c.Group.GET("/", c.getIndex)
	c.Group.GET("/admin", c.getAdmin)
	c.Group.GET("/search.json", c.getSearchJson)
	c.Group.GET("/robots.txt", c.getRobotsTxt)
	c.Group.GET("/sitemap.xml", c.getSitemapXml)

	c.pageController = &PageController{}
	c.AddController("/archive", &ArchiveController{})
	c.AddController("/article", &ArticleController{})
	c.AddController("/category", &CategoryController{})
	c.AddController("/page", c.pageController)
	c.AddController("/tag", &TagController{})
	c.AddController("/u", &UrlController{})
}

func (c *IndexController) getIndex(context *gin.Context) {
	c.pageController.getPage(context)
}

func (c *IndexController) getAdmin(context *gin.Context) {
	context.Redirect(http.StatusTemporaryRedirect, "/admin/index.html")
}

func (c *IndexController) getSearchJson(context *gin.Context) {
	articles := c.articleService.ListArticle()
	searchArray := make([]map[string]any, 0)
	for _, article := range articles {
		articleItemVO := html_vo.ArticleItemVO{}
		articleItemVO.FromSearch(article)
		searchArray = append(searchArray, map[string]any{
			"title":   articleItemVO.Title,
			"content": articleItemVO.Content,
			"url":     articleItemVO.Url,
		})
	}
	context.JSON(http.StatusOK, searchArray)
}

func (c *IndexController) getRobotsTxt(context *gin.Context) {
	builder := strings.Builder{}
	tmpl, _ := template.ParseFS(res.TemplatesFS, "robots.txt")
	_ = tmpl.Execute(&builder, map[string]string{
		"Host": context.Request.Host,
	})
	context.String(http.StatusOK, builder.String())
}

func (c *IndexController) getSitemapXml(context *gin.Context) {
	host := context.Request.Host
	articles := c.articleService.ListArticle()
	urls := make([]map[string]any, 0)
	urls = append(urls, map[string]any{
		"Host":   host,
		"Url":    "/",
		"Date":   time.Now().Format(time.DateOnly),
		"Freq":   "daily",
		"Weight": 0.6,
	})
	for _, article := range articles {
		articleItemVO := &html_vo.ArticleItemVO{}
		articleItemVO.FromItem(article)
		freq := c.getArticleUpdateFreq(articleItemVO)
		urls = append(urls, map[string]any{
			"Host":   host,
			"Url":    "/article/" + articleItemVO.Url,
			"Date":   articleItemVO.UpdateTime.Format(time.DateOnly),
			"Freq":   freq,
			"Weight": util.If(article.Weight > 0, 0.9, 0.6),
		})
	}
	builder := &strings.Builder{}
	tmpl, _ := template.ParseFS(res.TemplatesFS, "sitemap.xml")
	_ = tmpl.Execute(builder, urls)
	context.Data(http.StatusOK, "application/xml; charset=utf-8", []byte(builder.String()))
}

func (c *IndexController) getArticleUpdateFreq(article *html_vo.ArticleItemVO) string {
	lastUpdateTime := article.UpdateTime
	differenceDay := time.Now().Sub(lastUpdateTime).Seconds() / (24 * 60 * 60)
	if differenceDay <= 0 {
		return "daily"
	}
	if differenceDay <= 7 {
		return "weekly"
	}
	if differenceDay <= 30 {
		return "monthly"
	}
	if differenceDay <= 365 {
		return "yearly"
	}
	return "never"
}
