package api_vo

import (
	"blog/bean/po"
	"time"
)

type CommentVO struct {
	Id           int        `json:"id"`
	ArticleId    int        `json:"articleId"`
	ArticleTitle string     `json:"articleTitle"`
	ArticleUrl   string     `json:"articleUrl"`
	Nickname     string     `json:"nickname"`
	Email        string     `json:"email"`
	Url          *string    `json:"url"`
	Content      string     `json:"content"`
	Ip           *string    `json:"ip"`
	Location     *string    `json:"location"`
	Browser      string     `json:"browser"`
	System       string     `json:"system"`
	Status       int        `json:"status"`
	CreateTime   time.Time  `json:"createTime"`
	UpdateTime   *time.Time `json:"updateTime"`
}

func (v *CommentVO) From(comment po.Comment) {
	v.Id = comment.Id
	v.ArticleId = comment.ArticleId
	v.ArticleTitle = comment.Article.Title
	v.ArticleUrl = comment.Article.GetSafeUrl()
	v.Nickname = comment.Author
	v.Email = comment.Email
	v.Url = comment.Url
	v.Content = comment.Content
	v.Ip = comment.Ip
	v.Location = comment.Location
	v.Browser = comment.GetBrowser()
	v.System = comment.GetSystem()
	v.Status = comment.Status
	v.CreateTime = comment.CreateTime
	v.UpdateTime = comment.UpdateTime
}
