package api_vo

import (
	"blog/bean/po"
	"time"
)

type TagVO struct {
	Id           int        `json:"id"`
	Name         string     `json:"name"`
	CreateTime   time.Time  `json:"createTime"`
	UpdateTime   *time.Time `json:"updateTime"`
	ArticleCount int        `json:"articleCount"`
}

func (v *TagVO) From(tag po.Tag) {
	v.Id = tag.Id
	v.Name = tag.Name
	v.CreateTime = tag.CreateTime
	v.UpdateTime = tag.UpdateTime
	v.ArticleCount = tag.ArticleCount
}
