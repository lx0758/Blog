package api_vo

import "time"

type Url struct {
	Id          int
	Key         string
	Url         string
	Description string
	AuthorId    int
	Views       int
	TotalViews  int
	AuthorName  string
	CreateTime  time.Time
	UpdateTime  time.Time
	Status      int
}
