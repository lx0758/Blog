package html_vo

import (
	"blog/bean/po"
)

type CategoryVO struct {
	Name         string
	ArticleCount int
}

func (c *CategoryVO) From(category po.Category) {
	c.Name = category.Name
	c.ArticleCount = len(category.Articles)
}
