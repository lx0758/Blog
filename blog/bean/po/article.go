package po

import "time"

type Article struct {
	Id            int        `gorm:"column:id;comment:ID;primaryKey;autoIncrement;index:index_article_id"`
	Title         string     `gorm:"column:title;comment:标题"`
	Content       string     `gorm:"column:content;comment:内容"`
	CategoryId    int        `gorm:"column:category_id;comment:分类ID"`
	AuthorId      int        `gorm:"column:author_id;comment:作者ID"`
	Weight        *int       `gorm:"column:weight;comment:展示权重"`
	Views         int        `gorm:"column:views;comment:阅读量"`
	CommentStatus int        `gorm:"column:comment_status;comment:允许评论 0_禁止评论 1_允许评论"`
	Status        int        `gorm:"column:status;comment:状态 0_草稿 1_发布"`
	CreateTime    time.Time  `gorm:"column:create_time;comment:创建时间"`
	UpdateTime    *time.Time `gorm:"column:update_time;comment:更新时间"`

	Category Category     `gorm:"foreignKey:Id;references:CategoryId"`
	Author   User         `gorm:"foreignKey:Id;references:AuthorId"`
	Urls     []ArticleUrl `gorm:"foreignKey:ArticleId;references:Id"`
	Tags     []Tag        `gorm:"many2many:t_article_tag;foreignKey:Id;joinForeignKey:ArticleId;References:TagId;joinReferences:Id"`
}
