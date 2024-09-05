package po

import "time"

type Fragment struct {
	Id         int        `gorm:"column:id;comment:ID;index:index_fragment_id;primaryKey;autoIncrement"`
	Key        string     `gorm:"column:key;comment:KEY;index:index_fragment_key"`
	Content    string     `gorm:"column:content;comment:内容"`
	AuthorId   int        `gorm:"column:author_id;comment:作者ID"`
	Status     int        `gorm:"column:status;comment:状态 0_禁用 1_启用"`
	CreateTime time.Time  `gorm:"column:create_time;comment:创建时间"`
	UpdateTime *time.Time `gorm:"column:update_time;comment:更新时间"`

	Author User `gorm:"foreignKey:Id;references:AuthorId"`
}
