package po

import "time"

type Link struct {
	Id         int        `gorm:"column:id;comment:ID;index:index_link_id;primaryKey;autoIncrement"`
	Title      string     `gorm:"column:title;comment:标题"`
	Url        string     `gorm:"column:url;comment:链接"`
	Weight     *int       `gorm:"column:weight;comment:权重"`
	Status     int        `gorm:"column:status;comment:状态 0_未启用 1_已启用"`
	CreateTime time.Time  `gorm:"column:create_time;comment:创建时间"`
	UpdateTime *time.Time `gorm:"column:update_time;comment:更新时间"`
}
