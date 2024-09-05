package api_vo

import "time"

type Article struct {
	Id            int
	Title         string
	CategoryId    int
	Url           string
	Tags          []string
	Content       string
	Weight        int
	EnableComment bool
	Status        int
	CreateTime    time.Time
	UpdateTime    time.Time
}
