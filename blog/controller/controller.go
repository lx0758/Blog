package controller

import (
	"blog/bean/api_vo"
	"blog/bean/html_vo"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddController(engine *gin.Engine, relativePath string, controller ControllerInterface) *gin.RouterGroup {
	childGroup := engine.Group(relativePath)
	controller.initController(engine, childGroup)
	controller.OnInitController()
	return childGroup
}

type ControllerInterface interface {
	initController(engine *gin.Engine, group *gin.RouterGroup)
	OnInitController()
}

type Controller struct {
	Engine *gin.Engine
	Group  *gin.RouterGroup
}

func (c *Controller) initController(engine *gin.Engine, group *gin.RouterGroup) {
	c.Engine = engine
	c.Group = group
}

func (c *Controller) AddController(relativePath string, controller ControllerInterface) *gin.RouterGroup {
	childGroup := c.Group.Group(relativePath)
	controller.initController(c.Engine, childGroup)
	controller.OnInitController()
	return childGroup
}

func (c *Controller) Forward(context *gin.Context, url string) {
	context.Request.URL.Path = url
	c.Engine.HandleContext(context)
}

func (c *Controller) Render(context *gin.Context, code int, page string, data any) {
	pageTitle := ""
	if pageTitle == "" {
		switch page {
		case "archive":
			pageTitle = "归档"
			break
		case "article":
			pageTitle = "文章"
			break
		case "category":
			pageTitle = "分类详情"
			break
		case "categorys":
			pageTitle = "分类"
			break
		case "page":
			pageTitle = ""
			break
		case "tag":
			pageTitle = "标签详情"
			break
		case "tags":
			pageTitle = "标签"
			break
		}
	}
	blogVO := html_vo.Blog{
		SiteDomain:         "6xyun.cn",
		SiteTitle:          "6x's Website",
		SiteDescription:    "我的网站",
		SiteKeywords:       "6x,Liux",
		SiteBeianIcp:       "浙B2-20080101-4",
		SiteBeianPs:        "浙公网安备33010602009975号",
		SiteBaidu:          "123",
		SiteCreateYear:     "2014",
		SiteArticleCount:   1001,
		SiteCategoryCount:  20,
		SiteTagCount:       63,
		OwnerAvatar:        "",
		OwnerNickname:      "6x",
		OwnerDescription:   "一个程序员",
		OwnerEmail:         "lx0758@qq.com",
		OwnerGithub:        "lx0758",
		OwnerWeibo:         "lx0758",
		OwnerGoogle:        "lx0758",
		OwnerTwitter:       "lx0758",
		OwnerFacebook:      "lx0758",
		OwnerStackOverflow: "lx0758",
		OwnerYoutube:       "lx0758",
		OwnerInstagram:     "lx0758",
		OwnerSkype:         "lx0758",
		Links: []html_vo.Link{{
			Title: "百度",
			Url:   "https://www.baidu.com",
		}, {
			Title: "Google",
			Url:   "https://www.google.com",
		}},
	}
	context.HTML(code, page, map[string]any{
		"Blog":            blogVO,
		"Page":            page,
		"PageTitle":       pageTitle,
		"PageDescription": "测试描述",
		"PageKeywords":    "测试关键字",
		"PageUrl":         "/page/1",
		"PageData":        data,
	})
}

type RestController struct {
	Controller
}

func (c *RestController) Succeed(context *gin.Context, data any) {
	context.JSON(http.StatusOK, api_vo.Resp[any]{Data: data})
}

func (c *RestController) Failed(context *gin.Context, code int, message string, data any) {
	context.JSON(http.StatusOK, api_vo.Resp[any]{Status: code, Message: message, Data: data})
}
