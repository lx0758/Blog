package html

import (
	"blog/bean/html_vo"
	"blog/controller"
	"blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
)

type TagController struct {
	controller.Controller
	*service.ThemeService

	tagService     *service.TagService
	articleService *service.ArticleService
}

func (c *TagController) OnInitController() {
	c.ThemeService = service.GetService[*service.ThemeService](c.ThemeService)
	c.tagService = service.GetService[*service.TagService](c.tagService)
	c.articleService = service.GetService[*service.ArticleService](c.articleService)

	c.Group.GET("/", c.getTags)
	c.Group.GET("/:tagName", c.getTag)
	c.Group.GET("/:tagName/:pageNum", c.getTag)
}

func (c *TagController) getTags(context *gin.Context) {
	tags := c.tagService.ListTag()
	tagVOs := make([]html_vo.TagVO, 0)
	for _, tag := range tags {
		tagVO := html_vo.TagVO{}
		tagVO.From(tag)
		tagVOs = append(tagVOs, tagVO)
	}
	sort.Slice(tagVOs, func(i, j int) bool {
		return tagVOs[i].Name < tagVOs[j].Name
	})
	c.Render(context, http.StatusOK, "tags.gohtml", tagVOs)
}

func (c *TagController) getTag(context *gin.Context) {
	tagName := c.GetPathString(context, "tagName", "")
	tag := c.tagService.QueryTag(tagName)
	if tag == nil {
		c.ErrorNotFound(context)
	}

	pageNum, err := c.GetPathInt(context, "pageNum", 1)
	if err != nil {
		c.Error(context, http.StatusBadRequest, "Wrong parameters:%s", err)
	}
	pagination := c.articleService.PaginationArticleByTag(tag.Id, pageNum)
	articleItemVOs := make([]html_vo.ArticleItemVO, 0)
	for _, article := range pagination.List {
		archiveItemVO := html_vo.ArticleItemVO{}
		archiveItemVO.FromItem(article)
		articleItemVOs = append(articleItemVOs, archiveItemVO)
	}
	c.Render(context, http.StatusOK, "tag.gohtml", html_vo.PaginationVO[html_vo.ArticleItemVO]{
		Name:     tag.Name,
		Path:     "tag/" + tagName,
		PageNum:  pagination.PageNum,
		PageSize: pagination.PageSize,
		Size:     pagination.Size,
		Total:    pagination.Total,
		List:     articleItemVOs,
	})
}
