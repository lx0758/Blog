package api_vo

import (
	"blog/bean/po"
	"time"
)

type UrlVO struct {
	Id          int        `json:"id"`
	Key         string     `json:"key"`
	Url         string     `json:"url"`
	Description string     `json:"description"`
	AuthorId    int        `json:"authorId"`
	Views       int        `json:"views"`
	TotalViews  int        `json:"totalViews"`
	AuthorName  string     `json:"authorName"`
	Status      int        `json:"status"`
	CreateTime  time.Time  `json:"createTime"`
	UpdateTime  *time.Time `json:"updateTime"`
}

func (v *UrlVO) From(url po.Url) {
	v.Id = url.Id
	v.Key = url.Key
	v.Url = url.Url
	v.Description = url.Description
	v.AuthorId = url.AuthorId
	v.Views = url.Views
	v.TotalViews = url.TotalViews
	v.AuthorName = url.Author.Nickname
	v.Status = url.Status
	v.CreateTime = url.CreateTime
	v.UpdateTime = url.UpdateTime
}
