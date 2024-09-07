package api_vo

import "time"

type LinkVO struct {
	Id         int
	Title      string
	Url        string
	Weight     int
	CreateTime time.Time
	UpdateTime time.Time
	Status     int
}
