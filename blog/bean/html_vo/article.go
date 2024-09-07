package html_vo

import (
	"blog/bean/po"
	"blog/markdown"
	"blog/util"
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
	Prev           *ArticleItemVO
	Next           *ArticleItemVO
}

type CatalogueVO struct {
	Title  string
	Anchor string
	Childs []CatalogueVO
}

func (a *ArticleVO) From(article po.Article, prevArticle, nextArticle *po.Article) {
	html, text, catalogues := markdown.Render(article.Content, false)
	var prevArticleItemVO, nextArticleItemVO *ArticleItemVO
	if prevArticle != nil {
		prevArticleItemVO = &ArticleItemVO{}
		prevArticleItemVO.FromItem(*prevArticle)
	}
	if nextArticle != nil {
		nextArticleItemVO = &ArticleItemVO{}
		nextArticleItemVO.FromItem(*nextArticle)
	}
	a.Id = article.Id
	a.Url = article.GetSafeUrl()
	a.Title = article.Title
	a.Description = text
	a.Content = html
	a.CreateTime = article.CreateTime
	a.UpdateTime = util.IfNotNil(article.UpdateTime, article.CreateTime)
	a.CategoryName = article.Category.Name
	a.AuthorNickname = article.Author.Nickname
	a.Views = article.Views
	a.EnableComment = article.CommentStatus == po.ARTICLE_COMMENT_ENABLE
	a.Tags = article.GetSafeTags()
	a.Catalogues = convertCatalog(catalogues)
	a.Prev = prevArticleItemVO
	a.Next = nextArticleItemVO
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