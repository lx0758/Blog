package po

import "time"

type Tag struct {
	Id         int        `gorm:"column:id;comment:ID;not null;index;primaryKey;autoIncrement"`
	Name       string     `gorm:"column:name;comment:名称;not null;unique;index;"`
	CreateTime time.Time  `gorm:"column:create_time;comment:创建时间;not null"`
	UpdateTime *time.Time `gorm:"column:update_time;comment:更新时间"`

	ArticleCount int `gorm:"column:article_count;-:migration;->"`
}
