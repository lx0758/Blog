package html_vo

import "time"

type ArticleItem struct {
	Url          string
	Top          bool
	Title        string
	Content      string
	CreateTime   time.Time
	UpdateTime   time.Time
	CategoryName string
}
