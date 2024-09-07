package html

import (
	"blog/bean/html_vo"
	"blog/bean/po"
	"blog/controller"
	"blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ArticleController struct {
	controller.Controller
	*service.ThemeService

	articleService    *service.ArticleService
	articleUrlService *service.ArticleUrlService
	commentService    *service.CommentService
}

func (c *ArticleController) OnInitController() {
	c.ThemeService = service.GetService[*service.ThemeService](c.ThemeService)
	c.articleService = service.GetService[*service.ArticleService](c.articleService)
	c.articleUrlService = service.GetService[*service.ArticleUrlService](c.articleUrlService)
	c.commentService = service.GetService[*service.CommentService](c.commentService)

	c.Group.GET("/:url", c.getArticle)
	c.Group.GET("/:url/comment", c.getArticleComment)
	c.Group.POST("/:url/comment", c.postArticleComment)
}

func (c *ArticleController) getArticle(context *gin.Context) {
	url := c.GetPathString(context, "url", "")
	if url == "" {
		c.ErrorNotFound(context)
	}

	articleUrl, article := c.getKeyArticle(url)
	if articleUrl != nil {
		context.Redirect(http.StatusTemporaryRedirect, articleUrl.Url)
		return
	}
	if article == nil {
		c.ErrorNotFound(context)
	}

	prevArticle := c.articleService.QueryPrevArticle(article.Id)
	nextArticle := c.articleService.QueryNextArticle(article.Id)

	articleVO := &html_vo.ArticleVO{}
	articleVO.From(*article, prevArticle, nextArticle)
	c.Render(context, http.StatusOK, "article.gohtml", articleVO)
}

func (c *ArticleController) getArticleComment(context *gin.Context) {
	id, err := c.GetPathInt(context, "url", 0)
	if err != nil || id == 0 {
		c.ErrorNotFound(context)
	}
}

func (c *ArticleController) postArticleComment(context *gin.Context) {
	id, err := c.GetPathInt(context, "url", 0)
	if err != nil || id == 0 {
		c.ErrorNotFound(context)
	}
}

func (c *ArticleController) getKeyArticle(url string) (*po.ArticleUrl, *po.Article) {
	articleUrl := c.articleUrlService.QueryArticleUrlByUrl(url)
	if articleUrl != nil {
		if articleUrl.Status == po.ARTICLE_URL_STATUS_LAST {
			// 使用最新的 URL 访问, 直接查库
			return nil, c.articleService.QueryArticle(articleUrl.ArticleId)
		} else {
			articleUrl := c.articleUrlService.QueryArticleUrlById(articleUrl.ArticleId)
			// 使用旧的 Url 请求, 重定向到新的
			return articleUrl, nil
		}
	}
	articleId, err := strconv.Atoi(url)
	if err != nil {
		// 传入参数不是有效 URL 也不是 ID
		return nil, nil
	}

	articleUrl = c.articleUrlService.QueryArticleUrlById(articleId)
	if articleUrl != nil {
		// 使用ID请求, 但是存在 URL, 重定向到URL
		return articleUrl, nil
	}

	// 传入参数可能是个 ID, 直接查库
	return nil, c.articleService.QueryArticle(articleId)
}