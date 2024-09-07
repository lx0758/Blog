package po

import (
	"strconv"
	"time"
)

type Article struct {
	Id            int        `gorm:"column:id;comment:ID;not null;index;primaryKey;autoIncrement"`
	Title         string     `gorm:"column:title;comment:标题;not null"`
	Content       string     `gorm:"column:content;comment:内容;not null"`
	CategoryId    int        `gorm:"column:category_id;comment:分类ID;not null"`
	AuthorId      int        `gorm:"column:author_id;comment:作者ID;not null"`
	Weight        int        `gorm:"column:weight;comment:展示权重;not null;default:0"`
	Views         int        `gorm:"column:views;comment:阅读量;not null;default:0"`
	CommentStatus int        `gorm:"column:comment_status;comment:允许评论 0_禁止评论 1_允许评论;not null"`
	Status        int        `gorm:"column:status;comment:状态 0_草稿 1_发布;not null"`
	CreateTime    time.Time  `gorm:"column:create_time;comment:创建时间;not null"`
	UpdateTime    *time.Time `gorm:"column:update_time;comment:更新时间"`

	Category Category     `gorm:"foreignKey:CategoryId;references:Id"`
	Author   User         `gorm:"foreignKey:AuthorId;references:Id"`
	Urls     []ArticleUrl `gorm:"foreignKey:ArticleId;references:Id"` // 描述一对多的关联时, 外键列和关联列和一对一时是相反的
	Tags     []Tag        `gorm:"many2many:article_tag;foreignKey:Id;joinForeignKey:article_id;References:Id;joinReferences:tag_id"`
}

func (a *Article) GetSafeUrl() string {
	url := ""
	for _, u := range a.Urls {
		if u.Status == ARTICLE_URL_STATUS_LAST {
			url = u.Url
			break
		}
	}
	if url == "" {
		url = strconv.Itoa(a.Id)
	}
	return url
}

func (a *Article) GetSafeTags() []string {
	tags := make([]string, 0)
	for _, tag := range a.Tags {
		tags = append(tags, tag.Name)
	}
	return tags
}

const (
	ARTICLE_STATUS_UNPUBLISHED = 0
	ARTICLE_STATUS_PUBLISHED   = 1
	ARTICLE_COMMENT_DISABLE    = 0
	ARTICLE_COMMENT_ENABLE     = 1
)
