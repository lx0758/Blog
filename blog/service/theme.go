package service

import (
	"blog/bean/html_vo"
	"blog/bean/po"
	"blog/util"
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
		Blog:            s.getBlogVO(),
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
		Blog:            s.getBlogVO(),
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

func (s *ThemeService) getBlogVO() html_vo.BlogVO {
	configMap := s.blogService.GetCacheConfigMap()
	owner := s.blogService.GetCacheOwner()
	articleCount, categoryCount, tagCount, _, _ := s.blogService.GetCacheCount()
	links := s.blogService.GetCacheLinks()

	linkVOs := make([]html_vo.LinkVO, 0)
	for _, link := range links {
		linkVOs = append(linkVOs, html_vo.LinkVO{
			Title: link.Title,
			Url:   link.Url,
		})
	}
	return html_vo.BlogVO{
		SiteDomain:         util.IfNotNil(configMap[po.CONFIG_KEY_SITE_DOMAIN], ""),
		SiteTitle:          util.IfNotNil(configMap[po.CONFIG_KEY_SITE_TITLE], ""),
		SiteDescription:    util.IfNotNil(configMap[po.CONFIG_KEY_SITE_DESCRIPTION], ""),
		SiteKeywords:       util.IfNotNil(configMap[po.CONFIG_KEY_SITE_KEYWORDS], ""),
		SiteBeianIcp:       util.IfNotNil(configMap[po.CONFIG_KEY_SITE_BEIAN_ICP], ""),
		SiteBeianPs:        util.IfNotNil(configMap[po.CONFIG_KEY_SITE_BEIAN_PS], ""),
		SiteBaidu:          util.IfNotNil(configMap[po.CONFIG_KEY_SITE_BAIDU], ""),
		SiteCreateYear:     util.IfNotNil(configMap[po.CONFIG_KEY_SITE_CREATE_YEAR], ""),
		SiteArticleCount:   articleCount,
		SiteCategoryCount:  categoryCount,
		SiteTagCount:       tagCount,
		OwnerAvatar:        util.IfNotNil(owner.Avatar, ""),
		OwnerNickname:      owner.Nickname,
		OwnerDescription:   util.IfNotNil(owner.Description, ""),
		OwnerEmail:         util.IfNotNil(owner.Email, ""),
		OwnerGithub:        util.IfNotNil(owner.Accounts.Github, ""),
		OwnerWeibo:         util.IfNotNil(owner.Accounts.Weibo, ""),
		OwnerGoogle:        util.IfNotNil(owner.Accounts.Google, ""),
		OwnerTwitter:       util.IfNotNil(owner.Accounts.Twitter, ""),
		OwnerFacebook:      util.IfNotNil(owner.Accounts.Facebook, ""),
		OwnerStackOverflow: util.IfNotNil(owner.Accounts.StackOverflow, ""),
		OwnerYoutube:       util.IfNotNil(owner.Accounts.Youtube, ""),
		OwnerInstagram:     util.IfNotNil(owner.Accounts.Instagram, ""),
		OwnerSkype:         util.IfNotNil(owner.Accounts.Skype, ""),
		Links:              linkVOs,
	}
}
