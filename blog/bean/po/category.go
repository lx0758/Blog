package po

import "time"

type Category struct {
	Id         int        `gorm:"column:id;comment:ID;not null;index;primaryKey;autoIncrement"`
	Name       string     `gorm:"column:name;comment:名称;not null"`
	CreateTime time.Time  `gorm:"column:create_time;comment:创建时间;not null"`
	UpdateTime *time.Time `gorm:"column:update_time;comment:更新时间"`

	Articles []Article `gorm:"foreignKey:CategoryId;references:Id"`
}
