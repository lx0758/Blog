package html_vo

import (
	"blog/bean/po"
	"blog/markdown"
	"blog/util"
	"strconv"
	"time"
)

type ArticleVO struct {
	Id             int
	Url            string
	Title          string
	Description    string
	Content        string
	CreateTime     time.Time
	UpdateTime     time.Time
	CategoryName   string
	AuthorNickname string
	Views          int
	EnableComment  bool
	Tags           []string
	Catalogues     []CatalogueVO
}

func (a *ArticleVO) From(article po.Article) {
	url := ""
	for _, u := range article.Urls {
		if u.Status == po.ARTICLE_URL_STATUS_LAST {
			url = u.Url
			break
		}
	}
	description := markdown.RenderDescription(article.Content)
	content, catalogues := markdown.RenderContent(article.Content)
	tags := make([]string, 0)
	for _, tag := range article.Tags {
		tags = append(tags, tag.Name)
	}
	categoryName := util.If(article.Category.Name != "", article.Category.Name, "")
	authorNickname := util.If(article.Author.Nickname != "", article.Author.Nickname, "")
	a.Id = article.Id
	a.Url = url
	a.Title = article.Title
	a.Description = description
	a.Content = content
	a.CreateTime = article.CreateTime
	a.UpdateTime = article.UpdateTime
	a.CategoryName = categoryName
	a.AuthorNickname = authorNickname
	a.Views = article.Views
	a.EnableComment = article.CommentStatus == po.ARTICLE_COMMENT_ENABLE
	a.Tags = tags
	a.Catalogues = convertCatalog(catalogues)
}

func (a *ArticleVO) GetSafeUrl() string {
	return util.If(a.Url != "", a.Url, strconv.Itoa(a.Id))
}

func (a *ArticleVO) GetSafeDate() time.Time {
	return util.If(!a.UpdateTime.IsZero(), a.UpdateTime, a.CreateTime)
}

func convertCatalog(catalogues []markdown.Catalogue) []CatalogueVO {
	vos := make([]CatalogueVO, 0)
	for _, catalog := range catalogues {
		vos = append(vos, CatalogueVO{
			Title:  catalog.Title,
			Anchor: catalog.Anchor,
			Childs: convertCatalog(catalog.Childs),
		})
	}
	return vos
}
