package api_vo

import (
	"blog/bean/po"
	"time"
)

type FileVO struct {
	Id          int        `json:"id"`
	DisplayName string     `json:"displayName"`
	Length      int64      `json:"length"`
	Type        string     `json:"type"`
	Path        string     `json:"path"`
	Url         string     `json:"url"`
	AuthorId    int        `json:"authorId"`
	AuthorName  string     `json:"authorName"`
	CreateTime  time.Time  `json:"createTime"`
	UpdateTime  *time.Time `json:"updateTime"`
}

func (v *FileVO) From(link po.File, url string) {
	v.Id = link.Id
	v.DisplayName = link.DisplayName
	v.Length = link.Length
	v.Type = link.Type
	v.Path = link.Path
	v.Url = url
	v.AuthorId = link.AuthorId
	v.AuthorName = link.Author.Nickname
	v.CreateTime = link.CreateTime
	v.UpdateTime = link.UpdateTime
}
