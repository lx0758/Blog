package api_vo

import "time"

type CommentVO struct {
	Id           int
	ArticleId    int
	ArticleTitle string
	ArticleUrl   string
	Nickname     string
	Email        string
	Url          string
	Ip           string
	Location     string
	Browser      string
	System       string
	CreateTime   time.Time
	UpdateTime   time.Time
	Content      string
	Status       int
}
