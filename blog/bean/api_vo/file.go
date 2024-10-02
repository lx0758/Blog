package api_vo

import "time"

type FileVO struct {
	Id          int       `json:"id"`
	DisplayName string    `json:"displayName"`
	Length      int64     `json:"length"`
	Type        string    `json:"type"`
	Path        string    `json:"path"`
	Url         string    `json:"url"`
	AuthorId    int       `json:"authorId"`
	AuthorName  string    `json:"authorName"`
	CreateTime  time.Time `json:"createTime"`
	UpdateTime  time.Time `json:"updateTime"`
}
