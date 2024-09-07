package api_vo

import (
	"time"
)

type CategoryVO struct {
	Id           int
	Name         string
	CreateTime   time.Time
	UpdateTime   time.Time
	ArticleCount int
}