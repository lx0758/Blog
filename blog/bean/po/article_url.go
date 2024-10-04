package po

import "time"

type ArticleUrl struct {
	Url        string    `gorm:"column:url;comment:URL;not null;unique;index:article_url;primaryKey"`
	ArticleId  int       `gorm:"column:article_id;comment:文章ID;not null;index:article_url"`
	Status     int       `gorm:"column:status;comment:URL状态 0_历史的 1_当前的;not null"`
	CreateTime time.Time `gorm:"column:create_time;comment:创建时间;not null"`
}

const (
	ARTICLE_URL_STATUS_OLD  = 0
	ARTICLE_URL_STATUS_LAST = 1
)
