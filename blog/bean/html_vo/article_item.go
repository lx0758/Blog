package html_vo

import (
	"blog/bean/po"
	"blog/markdown"
	"blog/util"
	"time"
)

type ArticleItemVO struct {
	Url          string
	Top          bool
	Title        string
	Content      string
	Preview      bool
	CreateTime   time.Time
	UpdateTime   time.Time
	CategoryName string
}

func (a *ArticleItemVO) FromPage(article po.Article, fs markdown.FragmentSource) {
	html, hasPreview := markdown.RenderByPreview(article.Content, fs)
	a.Url = article.GetSafeUrl()
	a.Top = article.Weight > 0
	a.Title = article.Title
	a.Content = html
	a.Preview = hasPreview
	a.CreateTime = article.CreateTime
	a.UpdateTime = util.IfNotNil(article.UpdateTime, article.CreateTime)
	a.CategoryName = article.Category.Name
}

func (a *ArticleItemVO) FromItem(article po.Article) {
	a.Url = article.GetSafeUrl()
	a.Title = article.Title
	a.CreateTime = article.CreateTime
	a.UpdateTime = util.IfNotNil(article.UpdateTime, article.CreateTime)
}

func (a *ArticleItemVO) FromSearch(article po.Article, fs markdown.FragmentSource) {
	text := markdown.RenderBySearch(article.Content, fs)
	a.Url = article.GetSafeUrl()
	a.Title = article.Title
	a.Content = text
}
