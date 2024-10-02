package api_vo

import "time"

type LinkVO struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Url        string    `json:"url"`
	Weight     int       `json:"weight"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	Status     int       `json:"status"`
}
