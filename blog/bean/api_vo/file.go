package api_vo

import "time"

type FileVO struct {
	Id          int
	DisplayName string
	Length      int64
	Type        string
	Path        string
	Url         string
	AuthorId    int
	AuthorName  string
	CreateTime  time.Time
	UpdateTime  time.Time
}
