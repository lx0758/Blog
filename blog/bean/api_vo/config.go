package api_vo

import "time"

type ConfigVO struct {
	Key         string
	Value       string
	Description string
	CreateTime  time.Time
	UpdateTime  time.Time
}
