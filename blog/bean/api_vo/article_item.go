package api_vo

import (
	"time"
)

type ArticleItemVO struct {
	Id            int
	Title         string
	CategoryName  string
	Url           string
	Urls          []string
	Weight        int
	Views         int
	EnableComment bool
	Status        int
	CreateTime    time.Time
	UpdateTime    time.Time
}
