package api_vo

import "time"

type FragmentItemVO struct {
	Id         int       `json:"id"`
	Key        string    `json:"key"`
	AuthorId   int       `json:"authorId"`
	AuthorName string    `json:"authorName"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	Status     int       `json:"status"`
}
