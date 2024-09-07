package po

import "time"

type Config struct {
	Key         string     `gorm:"column:key;comment:键;not null;index;primaryKey"`
	Value       string     `gorm:"column:value;comment:值;not null"`
	Description string     `gorm:"column:description;comment:描述;not null"`
	CreateTime  time.Time  `gorm:"column:create_time;comment:创建时间;not null"`
	UpdateTime  *time.Time `gorm:"column:update_time;comment:更新时间"`
}
