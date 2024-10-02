package api_vo

import "time"

type FragmentVO struct {
	Id         int       `json:"id"`
	Key        string    `json:"key"`
	Content    string    `json:"content"`
	AuthorId   int       `json:"authorId"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	Status     int       `json:"status"`
}
