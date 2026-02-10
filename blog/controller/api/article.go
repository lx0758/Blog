package api

import (
	"blog/bean/api_vo"
	"blog/bean/po"
	"blog/controller"
	"blog/service"
	"blog/util"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	controller.RestController

	articleService *service.ArticleService
}

func (c *ArticleController) OnInitController() {
	c.articleService = service.GetService[*service.ArticleService](c.articleService)

	c.Group.GET("", c.listArticle)
	c.Group.GET(":id", c.queryArticle)
	c.Group.POST("", c.addArticle)
	c.Group.PUT(":id", c.updateArticle)
	c.Group.PUT(":id/status", c.updateArticleStatus)
	c.Group.DELETE(":id", c.deleteArticle)
}

type listArticle struct {
	Title       *string `form:"title"`
	CategoryId  *int    `form:"categoryId"`
	Url         *string `form:"url"`
	Comment     *bool   `form:"enableComment"`
	Status      *int    `form:"status"`
	PageNum     int     `form:"pageNum" binding:"required"`
	PageSize    int     `form:"pageSize" binding:"required"`
	OrderName   *string `form:"orderName"`
	OrderMethod *string `form:"orderMethod"`
}

// @Summary		list article
// @Description	list article
// @Tags		article
// @Accept		json
// @Produce		json
// @Param		title			body	string	false	"title"
// @Param		categoryId		body	string	false	"categoryId"
// @Param		url				body	string	false	"url"
// @Param		enableComment	body	string	false	"enableComment"
// @Param		status			body	string	false	"status"
// @Param		pageNum			body	int		true	"pageNum"
// @Param		pageSize		body	int		true	"pageSize"
// @Param		orderName		body	string	false	"orderName"
// @Param		orderMethod		body	string	false	"orderMethod"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/article [GET]
func (c *ArticleController) listArticle(context *gin.Context) {
	var listArticle listArticle
	if err := context.ShouldBind(&listArticle); err != nil {
		c.RestValidationError(context, err)
	}

	pagination := c.articleService.PaginationByAdmin(
		listArticle.Title,
		listArticle.CategoryId,
		listArticle.Url,
		listArticle.Comment,
		listArticle.Status,
		listArticle.PageNum,
		listArticle.PageSize,
		listArticle.OrderName,
		listArticle.OrderMethod,
	)
	articleItemVOs := make([]api_vo.ArticleItemVO, 0)
	for _, article := range pagination.List {
		articleVO := api_vo.ArticleItemVO{}
		articleVO.From(article)
		articleItemVOs = append(articleItemVOs, articleVO)
	}
	paginationVO := api_vo.PaginationVO[api_vo.ArticleItemVO]{
		PageNum:  pagination.PageNum,
		PageSize: pagination.PageSize,
		Size:     pagination.Size,
		Total:    pagination.Total,
		List:     articleItemVOs,
	}
	c.RestSucceed(context, paginationVO)
}

type pathArticleId struct {
	Id int `uri:"id" binding:"required"`
}

// @Summary		query article
// @Description	query article
// @Tags		article
// @Accept		json
// @Produce		json
// @Param		id			path	int	true	"id"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/article/{id} [GET]
func (c *ArticleController) queryArticle(context *gin.Context) {
	var pathArticleId pathArticleId
	if err := context.ShouldBindUri(&pathArticleId); err != nil {
		c.RestValidationError(context, err)
	}

	article := c.articleService.QueryByAdmin(pathArticleId.Id)
	if article == nil {
		c.RestError(context, "文章不存在")
	}

	articleVO := api_vo.ArticleVO{}
	articleVO.From(*article)
	c.RestSucceed(context, articleVO)
}

type addUpdateArticle struct {
	Title         string    `form:"title" binding:"required"`
	CategoryId    int       `form:"categoryId" binding:"required"`
	Content       string    `form:"content" binding:"required"`
	Url           *string   `form:"url"`
	Weight        *int      `form:"weight"`
	EnableComment *bool     `form:"enableComment" binding:"required"`
	Status        *int      `form:"status" binding:"required"`
	Tags          *[]string `form:"tags"`
}

// @Summary		add article
// @Description	add article
// @Tags		article
// @Accept		json
// @Produce		json
// @Param		id				path	int		true	"id"
// @Param		title			body	string	true	"title"
// @Param		categoryId		body	int		true	"categoryId"
// @Param		content			body	string	true	"content"
// @Param		url				body	string	false	"url"
// @Param		weight			body	int		false	"weight"
// @Param		enableComment	body	bool	true	"enableComment"
// @Param		status			body	int		true	"status"
// @Param		tags			body	string	false	"tags"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/article [POST]
func (c *ArticleController) addArticle(context *gin.Context) {
	var addArticle addUpdateArticle
	if err := context.ShouldBind(&addArticle); err != nil {
		c.RestValidationError(context, err)
	}
	userId := context.GetInt(KEY_USER_ID)
	article, err := c.articleService.AddByAdmin(
		userId,
		addArticle.Title,
		addArticle.CategoryId,
		addArticle.Content,
		addArticle.Url,
		util.IfNotNil(addArticle.Weight, 0),
		util.If(*addArticle.EnableComment, po.ARTICLE_COMMENT_ENABLE, po.ARTICLE_COMMENT_DISABLE),
		*addArticle.Status,
		addArticle.Tags,
	)
	if article != nil {
		article = c.articleService.QueryByAdmin(article.Id)
		var articleVO api_vo.ArticleVO
		articleVO.From(*article)
		c.RestSucceed(context, articleVO)
	} else {
		c.RestError(context, err.Error())
	}
}

// @Summary		update article
// @Description	update article
// @Tags		article
// @Accept		json
// @Produce		json
// @Param		id				path	int		true	"id"
// @Param		title			body	string	true	"title"
// @Param		categoryId		body	int		true	"categoryId"
// @Param		content			body	string	true	"content"
// @Param		url				body	string	false	"url"
// @Param		weight			body	int		false	"weight"
// @Param		enableComment	body	bool	true	"enableComment"
// @Param		status			body	int		true	"status"
// @Param		tags			body	string	false	"tags"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/article [PUT]
func (c *ArticleController) updateArticle(context *gin.Context) {
	var pathArticleId pathArticleId
	if err := context.ShouldBindUri(&pathArticleId); err != nil {
		c.RestValidationError(context, err)
	}
	var updateArticle addUpdateArticle
	if err := context.ShouldBind(&updateArticle); err != nil {
		c.RestValidationError(context, err)
	}
	err := c.articleService.UpdateByAdmin(
		pathArticleId.Id,
		updateArticle.Title,
		updateArticle.CategoryId,
		updateArticle.Content,
		updateArticle.Url,
		util.IfNotNil(updateArticle.Weight, 0),
		util.If(*updateArticle.EnableComment, po.ARTICLE_COMMENT_ENABLE, po.ARTICLE_COMMENT_DISABLE),
		util.If(*updateArticle.Status == po.ARTICLE_STATUS_PUBLISHED, po.ARTICLE_STATUS_PUBLISHED, po.ARTICLE_STATUS_UNPUBLISHED),
		updateArticle.Tags,
	)
	if err == nil {
		c.RestSucceed(context, nil)
	} else {
		c.RestError(context, err.Error())
	}
}

type updateArticleStatus struct {
	Status *int `form:"status" binding:"required"`
}

// @Summary		update article status
// @Description	update article status
// @Tags		article
// @Accept		json
// @Produce		json
// @Param		id				path	int		true	"id"
// @Param		status			body	string	true	"status"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/article/{id}/status [PUT]
func (c *ArticleController) updateArticleStatus(context *gin.Context) {
	var pathArticleId pathArticleId
	if err := context.ShouldBindUri(&pathArticleId); err != nil {
		c.RestValidationError(context, err)
	}
	var updateArticleStatus updateArticleStatus
	if err := context.ShouldBind(&updateArticleStatus); err != nil {
		c.RestValidationError(context, err)
	}
	result := c.articleService.UpdateStatusByAdmin(
		pathArticleId.Id,
		util.If(*updateArticleStatus.Status == po.ARTICLE_STATUS_PUBLISHED, po.ARTICLE_STATUS_PUBLISHED, po.ARTICLE_STATUS_UNPUBLISHED),
	)
	if result {
		c.RestSucceed(context, nil)
	} else {
		c.RestError(context, "更新失败")
	}
}

// @Summary		delete article
// @Description	delete article
// @Tags		article
// @Accept		json
// @Produce		json
// @Param		id				path	int		true	"id"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/article/{id} [DELETE]
func (c *ArticleController) deleteArticle(context *gin.Context) {
	var pathArticleId pathArticleId
	if err := context.ShouldBindUri(&pathArticleId); err != nil {
		c.RestValidationError(context, err)
	}
	result := c.articleService.DeleteByAdmin(pathArticleId.Id)
	if result {
		c.RestSucceed(context, nil)
	} else {
		c.RestError(context, "更新失败")
	}
}
