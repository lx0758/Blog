package api_vo

import (
	"blog/bean/po"
	"time"
)

type LinkVO struct {
	Id         int        `json:"id"`
	Title      string     `json:"title"`
	Url        string     `json:"url"`
	Weight     int        `json:"weight"`
	Status     int        `json:"status"`
	CreateTime time.Time  `json:"createTime"`
	UpdateTime *time.Time `json:"updateTime"`
}

func (v *LinkVO) From(link po.Link) {
	v.Id = link.Id
	v.Title = link.Title
	v.Url = link.Url
	v.Weight = link.Weight
	v.Status = link.Status
	v.CreateTime = link.CreateTime
	v.UpdateTime = link.UpdateTime
}
