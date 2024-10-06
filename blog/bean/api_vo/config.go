package api_vo

import (
	"blog/bean/po"
	"time"
)

type ConfigVO struct {
	Key         string     `json:"key"`
	Value       *string    `json:"value"`
	Description string     `json:"description"`
	CreateTime  time.Time  `json:"createTime"`
	UpdateTime  *time.Time `json:"updateTime"`
}

func (v *ConfigVO) From(config po.Config) {
	v.Key = config.Key
	v.Value = config.Value
	v.Description = config.Description
	v.CreateTime = config.CreateTime
	v.UpdateTime = config.UpdateTime
}
