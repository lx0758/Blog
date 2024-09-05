package html_vo

import "time"

type Article struct {
	Id             int
	Url            string
	Title          string
	Description    string
	Content        string
	CreateTime     time.Time
	UpdateTime     time.Time
	CategoryName   string
	AuthorNickname string
	Views          int
	EnableComment  bool
	Tags           []string
	Catalogues     []Catalogue
}
