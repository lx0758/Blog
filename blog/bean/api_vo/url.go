package api_vo

import "time"

type UrlVO struct {
	Id          int       `json:"id"`
	Key         string    `json:"key"`
	Url         string    `json:"url"`
	Description string    `json:"description"`
	AuthorId    int       `json:"authorId"`
	Views       int       `json:"views"`
	TotalViews  int       `json:"totalViews"`
	AuthorName  string    `json:"authorName"`
	CreateTime  time.Time `json:"createTime"`
	UpdateTime  time.Time `json:"updateTime"`
	Status      int       `json:"status"`
}
