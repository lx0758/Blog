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

	articleService *service.ArticleService
}

func (c *IndexController) OnInitController() {
	c.articleService = service.GetService[*service.ArticleService](c.articleService)

	c.Group.GET("", c.getIndex)
	c.Group.GET("admin", c.getAdmin)
	c.Group.GET("search.json", c.getSearchJson)
	c.Group.GET("robots.txt", c.getRobotsTxt)
	c.Group.GET("sitemap.xml", c.getSitemapXml)

	c.AddController("archive", &ArchiveController{})
	c.AddController("article", &ArticleController{})
	c.AddController("category", &CategoryController{})
	c.AddController("page", &PageController{})
	c.AddController("tag", &TagController{})
	c.AddController("u", &UrlController{})
}

func (c *IndexController) getIndex(context *gin.Context) {
	c.Forward(context, "/page")
}

func (c *IndexController) getAdmin(context *gin.Context) {
	context.Redirect(http.StatusTemporaryRedirect, "/admin/index.html")
}

func (c *IndexController) getSearchJson(context *gin.Context) {
	context.JSON(http.StatusOK, map[string]any{})
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
		articleVO := &html_vo.ArticleVO{}
		articleVO.From(article)
		freq := c.getArticleUpdateFreq(articleVO)
		urls = append(urls, map[string]any{
			"Host":   host,
			"Url":    "/article/" + articleVO.GetSafeUrl(),
			"Date":   articleVO.GetSafeDate().Format(time.DateOnly),
			"Freq":   freq,
			"Weight": util.If(article.Weight > 0, 0.9, 0.6),
		})
	}
	builder := &strings.Builder{}
	tmpl, _ := template.ParseFS(res.TemplatesFS, "sitemap.xml")
	_ = tmpl.Execute(builder, urls)
	context.Data(http.StatusOK, "application/xml; charset=utf-8", []byte(builder.String()))
}

func (c *IndexController) getArticleUpdateFreq(article *html_vo.ArticleVO) string {
	lastUpdateTime := article.GetSafeDate()
	differenceDay := (time.Now().Unix() - lastUpdateTime.Unix()) / (24 * 60 * 60)
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
