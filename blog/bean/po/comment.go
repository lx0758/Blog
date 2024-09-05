package po

import "time"

type Comment struct {
	Id         int        `gorm:"column:id;comment:ID;index:index_comment_id;primaryKey;autoIncrement"`
	ArticleId  int        `gorm:"column:article_id;comment:文章ID"`
	ParentId   *int       `gorm:"column:parent_id;comment:父评论ID"`
	Author     string     `gorm:"column:author;comment:作者"`
	AuthorId   *string    `gorm:"column:author_id;comment:作者ID"`
	Email      string     `gorm:"column:email;comment:电子邮箱"`
	Url        *string    `gorm:"column:url;comment:链接地址"`
	Content    string     `gorm:"column:content;comment:内容"`
	Ip         *string    `gorm:"column:ip;comment:IP地址"`
	Location   *string    `gorm:"column:location;comment:位置"`
	Ua         *string    `gorm:"column:ua;comment:UA标识"`
	Status     int        `gorm:"column:status;comment:状态 0_审核中 1_已审核"`
	CreateTime time.Time  `gorm:"column:create_time;comment:创建时间"`
	UpdateTime *time.Time `gorm:"column:update_time;comment:更新时间"`

	Article Article   `gorm:"foreignKey:ArticleId;references:Id"`
	Childs  []Comment `gorm:"foreignKey:ParentId;references:Id"`
}
