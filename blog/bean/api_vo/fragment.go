package api_vo

import (
	"blog/bean/po"
	"time"
)

type FragmentVO struct {
	Id         int        `json:"id"`
	Key        string     `json:"key"`
	Content    string     `json:"content"`
	AuthorId   int        `json:"authorId"`
	Status     int        `json:"status"`
	CreateTime time.Time  `json:"createTime"`
	UpdateTime *time.Time `json:"updateTime"`
}

func (v *FragmentVO) From(fragment po.Fragment) {
	v.Id = fragment.Id
	v.Key = fragment.Key
	v.Content = fragment.Content
	v.AuthorId = fragment.AuthorId
	v.Status = fragment.Status
	v.CreateTime = fragment.CreateTime
	v.UpdateTime = fragment.UpdateTime
}
