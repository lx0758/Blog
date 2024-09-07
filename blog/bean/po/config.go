package po

import "time"

type Config struct {
	Key         string     `gorm:"column:key;comment:键;not null;index;primaryKey"`
	Value       string     `gorm:"column:value;comment:值;not null"`
	Description string     `gorm:"column:description;comment:描述;not null"`
	CreateTime  time.Time  `gorm:"column:create_time;comment:创建时间;not null"`
	UpdateTime  *time.Time `gorm:"column:update_time;comment:更新时间"`
}

const (
	CONFIG_KEY_SITE_BAIDU       = "SITE_BAIDU"
	CONFIG_KEY_SITE_BEIAN_ICP   = "SITE_BEIAN_ICP"
	CONFIG_KEY_SITE_BEIAN_PS    = "SITE_BEIAN_PS"
	CONFIG_KEY_SITE_CREATE_YEAR = "SITE_CREATE_YEAR"
	CONFIG_KEY_SITE_DESCRIPTION = "SITE_DESCRIPTION"
	CONFIG_KEY_SITE_DOMAIN      = "SITE_DOMAIN"
	CONFIG_KEY_SITE_KEYWORDS    = "SITE_KEYWORDS"
	CONFIG_KEY_SITE_TITLE       = "SITE_TITLE"
)
