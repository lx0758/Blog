package api_vo

import "time"

type FragmentVO struct {
	Id         int
	Key        string
	Content    string
	AuthorId   int
	CreateTime time.Time
	UpdateTime time.Time
	Status     int
}
