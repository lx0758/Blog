package api_vo

import "time"

type TagVO struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	CreateTime   time.Time `json:"createTime"`
	UpdateTime   time.Time `json:"updateTime"`
	ArticleCount int       `json:"articleCount"`
}
