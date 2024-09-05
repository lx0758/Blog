package po

import "time"

type Tag struct {
	Id         int        `gorm:"column:id;commentID;index:index_tag_id;primaryKey;autoIncrement"`
	Name       string     `gorm:"column:name;comment名称"`
	CreateTime time.Time  `gorm:"column:create_time;comment创建时间"`
	UpdateTime *time.Time `gorm:"column:update_time;comment更新时间"`

	Articles []Article `gorm:"many2many:t_article_tag;foreignKey:Id;joinForeignKey:TagId;References:ArticleId;joinReferences:Id"`
}
