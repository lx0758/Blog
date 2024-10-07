package service

import (
	"blog/bean/po"
	"blog/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type CommentService struct {
	Service
	db *gorm.DB

	emailService    *EmailService
	locationService *LocationService
}

func (s *CommentService) OnInitService() {
	s.db = database.GetDB()

	s.emailService = GetService[*EmailService](s.emailService)
	s.locationService = GetService[*LocationService](s.locationService)
}

func (s *CommentService) Count() int {
	var count int64
	s.db.Model(&po.Comment{}).
		Where("? = ?", clause.Column{Name: "status"}, po.COMMENT_STATUS_PUBLISHED).
		Count(&count)
	return int(count)
}

func (s *CommentService) PaginationByAdmin(
	id *int,
	author *string,
	email *string,
	ip *string,
	content *string,
	status *int,
	pageNum int,
	pageSize int,
	orderName *string,
	orderMethod *string,
) po.Pagination[po.Comment] {
	db := s.db.Model(&po.Comment{}).
		Preload("Article")
	if id != nil && *id != 0 {
		db = db.Where("? = ?", clause.Column{Name: "id"}, *id)
	}
	if author != nil && *author != "" {
		db = db.Where("? LIKE ?", clause.Column{Name: "author"}, "%"+*author+"%")
	}
	if email != nil && *email != "" {
		db = db.Where("? LIKE ?", clause.Column{Name: "email"}, "%"+*email+"%")
	}
	if ip != nil && *ip != "" {
		db = db.Where("? LIKE ?", clause.Column{Name: "ip"}, "%"+*ip+"%")
	}
	if content != nil && *content != "" {
		db = db.Where("? LIKE ?", clause.Column{Name: "content"}, "%"+*content+"%")
	}
	if status != nil {
		db = db.Where("? = ?", clause.Column{Name: "status"}, *status)
	}
	db = db.Order(database.TableOrderBy(orderName, orderMethod, "id", true))
	return database.Pagination[po.Comment](db, pageNum, pageSize)
}

func (s *CommentService) PaginationByHtml(
	articleId int,
	pageNum int,
) po.Pagination[po.Comment] {
	db := s.db.Model(&po.Comment{}).
		Preload("Childs", s.childsPreload).
		Where("? is NULL", clause.Column{Name: "parent_id"}).
		Where("? = ?", clause.Column{Name: "article_id"}, articleId).
		Where("? = ?", clause.Column{Name: "status"}, po.COMMENT_STATUS_PUBLISHED).
		Order(clause.OrderByColumn{
			Column: clause.Column{Name: "id"},
			Desc:   true,
		})
	return database.Pagination[po.Comment](db, pageNum, 10)
}

/**
 * 评论嵌套子评论的预加载, 通过递归实现
 * https://github.com/go-gorm/gorm/issues/4027
 */
func (s *CommentService) childsPreload(db *gorm.DB) *gorm.DB {
	return db.Model(&po.Comment{}).
		Preload("Childs", s.childsPreload).
		Where("? = ?", clause.Column{Name: "status"}, po.COMMENT_STATUS_PUBLISHED).
		Order(clause.OrderByColumn{
			Column: clause.Column{Name: "id"},
			Desc:   true,
		})
}

func (s *CommentService) ListNewComment(count int) []po.Comment {
	var comments []po.Comment
	s.db.Model(&po.Comment{}).
		Where("? = ?", clause.Column{Name: "status"}, po.COMMENT_STATUS_PUBLISHED).
		Order(clause.OrderByColumn{
			Column: clause.Column{Name: "id"},
			Desc:   true,
		}).
		Limit(count).
		Find(&comments)
	return comments
}

func (s *CommentService) QueryByAdmin(id int) *po.Comment {
	comment := po.Comment{}
	s.db.Model(&po.Comment{}).
		Where("? = ?", clause.Column{Name: "id"}, id).
		Take(&comment)
	if comment.Id == 0 {
		return nil
	}
	return &comment
}

func (s *CommentService) AddByAdmin(
	articleId int,
	parentId int,
	author string,
	authorId int,
	email string,
	ip string,
	ua string,
	content string,
	notify bool,
) bool {
	location := s.locationService.GetLocationFromIp(ip)
	db := s.db.Model(&po.Comment{}).
		Create(&po.Comment{
			ArticleId:  articleId,
			ParentId:   &parentId,
			Author:     author,
			AuthorId:   &authorId,
			Email:      email,
			Ip:         &ip,
			Location:   &location,
			Ua:         &ua,
			Content:    content,
			Status:     po.COMMENT_STATUS_PUBLISHED,
			CreateTime: time.Now(),
		})
	if db.RowsAffected > 0 && notify {
		// 评论回复完成, 发送评论通知邮件
		s.emailService.SendCommentReplayEmail(author, email, content, articleId, parentId)
	}
	if db.RowsAffected > 0 {
		refreshBlogCacheInfo()
	}
	return db.RowsAffected > 0
}

func (s *CommentService) AddByHtml(
	articleId int,
	parentId *int,
	author string,
	email string,
	ip string,
	ua string,
	content string,
) bool {
	location := s.locationService.GetLocationFromIp(ip)
	if parentId != nil && *parentId == 0 {
		parentId = nil
	}
	db := s.db.Model(&po.Comment{}).
		Create(&po.Comment{
			ArticleId:  articleId,
			ParentId:   parentId,
			Author:     author,
			Email:      email,
			Ip:         &ip,
			Location:   &location,
			Ua:         &ua,
			Content:    content,
			Status:     po.COMMENT_STATUS_UNPUBLISHED,
			CreateTime: time.Now(),
		})
	if db.RowsAffected > 0 {
		// 文章被评论, 向管理员发送评论通知邮件
		s.emailService.SendCommentedEmail(articleId, parentId, author, email, ip, location, ua, content)
	}
	if db.RowsAffected > 0 {
		refreshBlogCacheInfo()
	}
	return db.RowsAffected > 0
}

func (s *CommentService) UpdateStatusByAdmin(
	id int,
	status int,
	notify bool,
) bool {
	updateTime := time.Now()
	comment := po.Comment{}
	s.db.Model(&po.Comment{}).
		Where("? = ?", clause.Column{Name: "id"}, id).
		Find(&comment)
	db := s.db.Model(&po.Comment{}).
		Where("? = ?", clause.Column{Name: "id"}, id).
		Select("status", "update_time").
		Updates(&po.Comment{
			Status:     status,
			UpdateTime: &updateTime,
		})
	if db.RowsAffected > 0 && notify && status == po.COMMENT_STATUS_PUBLISHED && comment.ParentId != nil {
		// 文章评论审核完成, 发送评论通知邮件
		s.emailService.SendCommentReplayEmail(comment.Author, comment.Email, comment.Content, comment.ArticleId, *comment.ParentId)
	}
	if db.RowsAffected > 0 {
		refreshBlogCacheInfo()
	}
	return db.RowsAffected > 0
}

func (s *CommentService) DeleteByAdmin(id int) bool {
	db := s.db.Model(&po.Comment{}).
		Where("? = ?", clause.Column{Name: "id"}, id).
		Delete(&po.Comment{})
	if db.RowsAffected > 0 {
		refreshBlogCacheInfo()
	}
	return db.RowsAffected > 0
}
