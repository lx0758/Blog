package po

import "time"

type File struct {
	Id          int        `gorm:"column:id;comment:ID;not null;index;primaryKey;autoIncrement"`
	DisplayName string     `gorm:"column:display_name;comment:文件名;not null"`
	Length      int64      `gorm:"column:length;comment:文件长度;not null"`
	Type        string     `gorm:"column:type;comment:文件类型;not null"`
	Path        string     `gorm:"column:path;comment:存储路径;not null"`
	AuthorId    int        `gorm:"column:author_id;comment:上传者ID;not null"`
	CreateTime  time.Time  `gorm:"column:create_time;comment:上传时间;not null"`
	UpdateTime  *time.Time `gorm:"column:update_time;comment:更新时间"`

	Author User `gorm:"foreignKey:AuthorId;references:Id"`
}
