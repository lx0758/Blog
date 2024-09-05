package po

import "time"

type User struct {
	Id            int          `gorm:"column:id;comment:ID;index:index_user_id;primaryKey;autoIncrement"`
	Username      string       `gorm:"column:username;comment:用户名;index:index_user_username"`
	Password      string       `gorm:"column:password;comment:用户密码"`
	Avatar        *string      `gorm:"column:avatar;comment:头像"`
	Nickname      *string      `gorm:"column:nickname;comment:昵称"`
	Description   *string      `gorm:"column:description;comment:个性签名"`
	Email         *string      `gorm:"column:email;comment:电子邮箱"`
	Accounts      UserAccounts `gorm:"column:accounts;comment:社交账号;serializer:json"`
	LastLoginTime *time.Time   `gorm:"column:last_login_time;comment:最后登录时间"`
	Status        int          `gorm:"column:status;comment:是否启用 0_未启用 1_已启用"`
	CreateTime    time.Time    `gorm:"column:create_time;comment:创建时间"`
	UpdateTime    *time.Time   `gorm:"column:update_time;comment:更新时间"`
}

type UserAccounts struct {
	Github        string `json:"github"`
	Weibo         string `json:"weibo"`
	Google        string `json:"google"`
	Twitter       string `json:"twitter"`
	Facebook      string `json:"facebook"`
	StackOverflow string `json:"stackOverflow"`
	Youtube       string `json:"youtube"`
	Instagram     string `json:"instagram"`
	Skype         string `json:"skype"`
}
