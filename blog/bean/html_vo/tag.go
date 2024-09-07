package html_vo

import "blog/bean/po"

type TagVO struct {
	Name         string
	ArticleCount int
}

func (c *TagVO) From(tag po.Tag) {
	c.Name = tag.Name
	c.ArticleCount = len(tag.Articles)
}
