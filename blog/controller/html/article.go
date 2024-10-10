package html

import (
	"blog/bean/html_vo"
	"blog/bean/po"
	"blog/controller"
	"blog/service"
	"blog/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type ArticleController struct {
	controller.Controller
	*service.ThemeService

	articleService    *service.ArticleService
	articleUrlService *service.ArticleUrlService
	fragmentService   *service.FragmentService
	captchaService    *service.CaptchaService
	commentService    *service.CommentService
}

func (c *ArticleController) OnInitController() {
	c.ThemeService = service.GetService[*service.ThemeService](c.ThemeService)
	c.articleService = service.GetService[*service.ArticleService](c.articleService)
	c.articleUrlService = service.GetService[*service.ArticleUrlService](c.articleUrlService)
	c.fragmentService = service.GetService[*service.FragmentService](c.fragmentService)
	c.captchaService = service.GetService[*service.CaptchaService](c.captchaService)
	c.commentService = service.GetService[*service.CommentService](c.commentService)

	c.Group.GET("/:article", c.getArticle)
	c.Group.GET("/:article/comment", c.getArticleComment)
	c.Group.POST("/:article/comment", c.addArticleComment)
	c.Engine.GET("/captcha", c.getArticleCommentCaptcha)
}

func (c *ArticleController) getArticle(context *gin.Context) {
	url := c.GetPathString(context, "article", "")
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
	articleVO.From(*article, prevArticle, nextArticle, c.fragmentService)
	c.RenderArticle(context, articleVO)
}

type commentArticleId struct {
	Id int `uri:"article" binding:"required"`
}

type getArticleComment struct {
	SubjectId *int `form:"subjectId"`
	PageNum   int  `form:"page" binding:"required"`
}

// @Summary		query comment
// @Description	query comment
// @Tags		html
// @Accept		json
// @Produce		json
// @Param		url			path	int		false	"url"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/{url}/comment [GET]
func (c *ArticleController) getArticleComment(context *gin.Context) {
	var commentArticleId commentArticleId
	if err := context.ShouldBindUri(&commentArticleId); err != nil {
		c.ErrorNotFound(context)
	}
	var getArticleComment getArticleComment
	if err := context.ShouldBindQuery(&getArticleComment); err != nil {
		context.JSON(http.StatusOK, map[string]any{
			"status":  http.StatusBadRequest,
			"message": "参数错误",
		})
		return
	}
	pagination := c.commentService.PaginationByHtml(
		commentArticleId.Id,
		getArticleComment.PageNum,
	)
	commentVOs := make([]html_vo.CommentVO, 0)
	for _, comment := range pagination.List {
		commentVO := html_vo.CommentVO{}
		commentVO.From(comment)
		commentVOs = append(commentVOs, commentVO)
	}
	paginationVO := html_vo.CommentPaginationVO{
		Total:    pagination.Total,
		HasMore:  pagination.PageNum < pagination.Size,
		Comments: commentVOs,
	}
	context.JSON(http.StatusOK, map[string]any{
		"status":  http.StatusOK,
		"message": "获取成功",
		"data":    paginationVO,
	})
}

type addArticleComment struct {
	SubjectId *int    `form:"subjectId"`
	ParentId  *int    `form:"parentId"`
	Nickname  string  `form:"nickname" binding:"required"`
	Email     string  `form:"email" binding:"required"`
	Link      *string `form:"link"`
	Content   string  `form:"content" binding:"required"`
	Captcha   string  `form:"captcha" binding:"required"`
}

// @Summary		add comment
// @Description	add comment
// @Tags		html
// @Accept		json
// @Produce		json
// @Param		url			path	int		false	"url"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/{url}/comment [POST]
func (c *ArticleController) addArticleComment(context *gin.Context) {
	var commentArticleId commentArticleId
	if err := context.ShouldBindUri(&commentArticleId); err != nil {
		c.ErrorNotFound(context)
	}
	var addArticleComment addArticleComment
	if err := context.ShouldBind(&addArticleComment); err != nil {
		context.JSON(http.StatusOK, map[string]any{
			"status":  http.StatusBadRequest,
			"message": "参数错误",
		})
		return
	}

	session := sessions.Default(context)
	captcha := c.captchaService.Verify(
		session,
		service.CAPTCHA_TYPE_COMMENT,
		addArticleComment.Captcha,
	)
	if !captcha {
		context.JSON(http.StatusOK, map[string]any{
			"status":  http.StatusBadRequest,
			"message": "验证码错误",
		})
		return
	}

	result := c.commentService.AddByHtml(
		commentArticleId.Id,
		addArticleComment.ParentId,
		addArticleComment.Nickname,
		addArticleComment.Email,
		util.GetRequestIp(context),
		util.GetRequestUa(context),
		addArticleComment.Content,
	)
	if result {
		context.JSON(http.StatusOK, map[string]any{
			"status":  http.StatusOK,
			"message": "评论成功",
		})
	} else {
		context.JSON(http.StatusOK, map[string]any{
			"status":  http.StatusBadRequest,
			"message": "评论失败",
		})
	}
}

// @Summary		comment captcha
// @Description	comment captcha
// @Tags		html
// @Accept		*/*
// @Produce		image/png
// @Router		/captcha [get]
func (c *ArticleController) getArticleCommentCaptcha(context *gin.Context) {
	session := sessions.Default(context)
	captcha := c.captchaService.Generate(session, service.CAPTCHA_TYPE_COMMENT, 30*time.Minute)
	context.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate")
	context.Data(http.StatusOK, "image/png", captcha.Bytes())
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
