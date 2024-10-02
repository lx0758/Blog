package api_vo

import "time"

type CommentVO struct {
	Id           int       `json:"id"`
	ArticleId    int       `json:"articleId"`
	ArticleTitle string    `json:"articleTitle"`
	ArticleUrl   string    `json:"articleUrl"`
	Nickname     string    `json:"nickname"`
	Email        string    `json:"email"`
	Url          string    `json:"url"`
	Ip           string    `json:"ip"`
	Location     string    `json:"location"`
	Browser      string    `json:"browser"`
	System       string    `json:"system"`
	CreateTime   time.Time `json:"createTime"`
	UpdateTime   time.Time `json:"updateTime"`
	Content      string    `json:"content"`
	Status       int       `json:"status"`
}
