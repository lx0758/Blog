package po

import "time"

type Link struct {
	Id         int        `gorm:"column:id;comment:ID;not null;index;primaryKey;autoIncrement"`
	Title      string     `gorm:"column:title;comment:标题;not null"`
	Url        string     `gorm:"column:url;comment:链接;not null"`
	Weight     int        `gorm:"column:weight;comment:权重;not null;default:0"`
	Status     int        `gorm:"column:status;comment:状态 0_未启用 1_已启用;not null"`
	CreateTime time.Time  `gorm:"column:create_time;comment:创建时间;not null"`
	UpdateTime *time.Time `gorm:"column:update_time;comment:更新时间"`
}

const (
	LINK_STATUS_UNPUBLISHED = 0
	LINK_STATUS_PUBLISHED   = 1
)
