package po

import "time"

type File struct {
	Id          int        `gorm:"column:id;comment:ID;primaryKey;autoIncrement"`
	DisplayName string     `gorm:"column:display_name;comment:文件名"`
	Length      int64      `gorm:"column:length;comment:文件长度"`
	Type        string     `gorm:"column:type;comment:文件类型"`
	Path        string     `gorm:"column:path;comment:存储路径"`
	AuthorId    int        `gorm:"column:author_id;comment:上传者ID"`
	CreateTime  time.Time  `gorm:"column:create_time;comment:上传时间"`
	UpdateTime  *time.Time `gorm:"column:update_time;comment:更新时间"`

	Author User `gorm:"foreignKey:Id;references:AuthorId"`
}
