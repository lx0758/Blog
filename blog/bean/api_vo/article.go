package api_vo

import (
	"blog/bean/po"
	"time"
)

type ArticleVO struct {
	Id            int        `json:"id"`
	Title         string     `json:"title"`
	CategoryId    int        `json:"categoryId"`
	Url           string     `json:"url"`
	Tags          []string   `json:"tags"`
	Content       string     `json:"content"`
	Weight        int        `json:"weight"`
	EnableComment bool       `json:"enableComment"`
	Status        int        `json:"status"`
	CreateTime    time.Time  `json:"createTime"`
	UpdateTime    *time.Time `json:"updateTime"`
}

func (a *ArticleVO) From(article po.Article) {
	a.Id = article.Id
	a.Title = article.Title
	a.CategoryId = article.CategoryId
	a.Url = article.GetSafeUrl()
	a.Tags = article.GetSafeTags()
	a.Content = article.Content
	a.Weight = article.Weight
	a.EnableComment = article.CommentStatus == po.ARTICLE_COMMENT_ENABLE
	a.Status = article.Status
	a.CreateTime = article.CreateTime
	a.UpdateTime = article.UpdateTime
}
