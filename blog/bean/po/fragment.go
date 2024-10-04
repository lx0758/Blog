package po

import "time"

type Fragment struct {
	Id         int        `gorm:"column:id;comment:ID;not null;index;primaryKey;autoIncrement"`
	Key        string     `gorm:"column:key;comment:KEY;not null;unique;index"`
	Content    string     `gorm:"column:content;comment:内容;not null"`
	AuthorId   int        `gorm:"column:author_id;comment:作者ID;not null"`
	Status     int        `gorm:"column:status;comment:状态 0_禁用 1_启用;not null"`
	CreateTime time.Time  `gorm:"column:create_time;comment:创建时间;not null"`
	UpdateTime *time.Time `gorm:"column:update_time;comment:更新时间"`

	Author User `gorm:"foreignKey:AuthorId;references:Id"`
}
