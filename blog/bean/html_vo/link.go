package html_vo

import "blog/bean/po"

type LinkVO struct {
	Title string
	Url   string
}

func (c *LinkVO) From(link po.Link) {
	c.Title = link.Title
	c.Url = link.Url
}
