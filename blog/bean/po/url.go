package po

import "time"

type Url struct {
	Id          int        `gorm:"column:id;comment:ID;index:index_url_id;primaryKey;autoIncrement"`
	Key         string     `gorm:"column:key;comment:键;index:index_url_key"`
	Url         string     `gorm:"column:url;comment:链接地址"`
	Description string     `gorm:"column:description;comment:描述信息"`
	AuthorId    int        `gorm:"column:author_id;comment:作者ID"`
	Views       int        `gorm:"column:views;comment:访问次数"`
	TotalViews  int        `gorm:"column:total_views;comment:总访问次数"`
	Status      int        `gorm:"column:status;comment:状态 0_禁用 1_启用"`
	CreateTime  time.Time  `gorm:"column:create_time;comment:创建时间"`
	UpdateTime  *time.Time `gorm:"column:update_time;comment:更新时间"`

	Author User `gorm:"foreignKey:Id;references:AuthorId"`
}
