package html

import (
	"blog/controller"
	"blog/res"
	"blog/service"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"strings"
	"time"
)

type IndexController struct {
	controller.Controller
	*service.IndexService
}

func (c *IndexController) OnInitController() {
	c.IndexService = service.GetService[*service.IndexService](c.IndexService)

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
	builder := strings.Builder{}
	tmpl, _ := template.ParseFS(res.TemplatesFS, "sitemap.xml")
	_ = tmpl.Execute(&builder, map[string]any{
		"Host": host,
		"Date": time.DateOnly,
		"Urls": []map[string]any{
			{"Host": host, "Url": "1", "Date": "2024-09-01", "Freq": "daily", "Weight": 1.0},
			{"Host": host, "Url": "2", "Date": "2024-09-01", "Freq": "weekly", "Weight": 1.0},
			{"Host": host, "Url": "3", "Date": "2024-09-01", "Freq": "monthly", "Weight": 1.0},
			{"Host": host, "Url": "4", "Date": "2024-09-01", "Freq": "yearly", "Weight": 1.0},
			{"Host": host, "Url": "5", "Date": "2024-09-01", "Freq": "never", "Weight": 1.0},
		},
	})
	context.Data(http.StatusOK, "application/xml; charset=utf-8", []byte(builder.String()))
}
