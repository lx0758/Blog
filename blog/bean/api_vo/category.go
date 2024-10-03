package api_vo

import (
	"blog/bean/po"
	"time"
)

type CategoryVO struct {
	Id           int        `json:"id"`
	Name         string     `json:"name"`
	CreateTime   time.Time  `json:"createTime"`
	UpdateTime   *time.Time `json:"updateTime"`
	ArticleCount int        `json:"articleCount"`
}

func (c *CategoryVO) From(category po.Category) {
	c.Id = category.Id
	c.Name = category.Name
	c.CreateTime = category.CreateTime
	c.UpdateTime = category.UpdateTime
	c.ArticleCount = category.ArticleCount
}
