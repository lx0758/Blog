package po

import "time"

type Url struct {
	Id          int        `gorm:"column:id;comment:ID;not null;index;primaryKey;autoIncrement"`
	Key         string     `gorm:"column:key;comment:键;not null;unique;index"`
	Url         string     `gorm:"column:url;comment:链接地址;not null"`
	Description string     `gorm:"column:description;comment:描述信息;not null"`
	AuthorId    int        `gorm:"column:author_id;comment:作者ID;not null"`
	Views       int        `gorm:"column:views;comment:访问次数;not null;default:0"`
	TotalViews  int        `gorm:"column:total_views;comment:总访问次数;not null;default:0"`
	Status      int        `gorm:"column:status;comment:状态 0_禁用 1_启用;not null"`
	CreateTime  time.Time  `gorm:"column:create_time;comment:创建时间;not null"`
	UpdateTime  *time.Time `gorm:"column:update_time;comment:更新时间"`

	Author User `gorm:"foreignKey:AuthorId;references:Id"`
}

const (
	URL_STATUS_UNPUBLISHED = 0
	URL_STATUS_PUBLISHED   = 1
)
