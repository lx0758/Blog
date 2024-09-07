package html_vo

import "time"

type ArticleItemVO struct {
	Url          string
	Top          bool
	Title        string
	Content      string
	CreateTime   time.Time
	UpdateTime   time.Time
	CategoryName string
}
