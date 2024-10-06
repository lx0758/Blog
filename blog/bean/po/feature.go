package po

import "time"

type Feature struct {
	Key         string     `gorm:"column:key;comment:键;not null;unique;index;primaryKey"`
	Value       *string    `gorm:"column:value;comment:值"`
	Description string     `gorm:"column:description;comment:描述;not null"`
	CreateTime  time.Time  `gorm:"column:create_time;comment:创建时间;not null"`
	UpdateTime  *time.Time `gorm:"column:update_time;comment:更新时间"`
}

const (
	FEATURE_KEY_SMTP = "SMTP"
)

type SMTPFeature struct {
	Enable    bool   `json:"enable"`
	Hostname  string `json:"hostname"`
	Port      int    `json:"port"`
	Ssl       bool   `json:"ssl"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FromName  string `json:"fromName"`
	FromEmail string `json:"fromEmail"`
}
