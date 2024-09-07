package html_vo

import (
	"blog/bean/po"
	"blog/markdown"
	"blog/util"
	"strings"
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

func (a *ArticleItemVO) FromPage(article po.Article) {
	html, _, _ := markdown.Render(article.Content, true)
	hasPreview := strings.Contains(html, "<!-- more -->")
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

func (a *ArticleItemVO) FromSearch(article po.Article) {
	_, text, _ := markdown.Render(article.Content, false)
	a.Url = article.GetSafeUrl()
	a.Title = article.Title
	a.Content = text
}
