package api_vo

import "time"

type FragmentItem struct {
	Id         int
	Key        string
	AuthorId   int
	AuthorName string
	CreateTime time.Time
	UpdateTime time.Time
	Status     int
}
