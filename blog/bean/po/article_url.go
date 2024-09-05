package po

import "time"

type ArticleUrl struct {
	Url        string    `gorm:"column:url;comment:URL;index:index_article_url"`
	ArticleId  int       `gorm:"column:article_id;comment:文章ID;index:index_article_url"`
	Status     int       `gorm:"column:status;comment:URL状态 0_历史的 1_当前的"`
	CreateTime time.Time `gorm:"column:create_time;comment:创建时间"`
}

const STATUS_LAST = 1
const STATUS_OLD = 1
