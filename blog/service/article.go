package service

import (
	"blog/bean/po"
	"blog/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

const ARTICLE_PAGE_SIZE = 10

type ArticleService struct {
	Service
	db *gorm.DB
}

func (s *ArticleService) OnInitService() {
	s.db = database.GetDB()
}

func (s *ArticleService) Count() int {
	var count int64
	s.db.Model(&po.Article{}).
		Where("? = ?", clause.Column{Name: "status"}, po.ARTICLE_STATUS_PUBLISHED).
		Count(&count)
	return int(count)
}

func (s *ArticleService) QueryArticle(id int) *po.Article {
	var article po.Article
	s.db.Model(&po.Article{}).
		Preload("Category").Preload("Author").Preload("Urls").Preload("Tags").
		Where("? = ? AND ? = ?", clause.Column{Name: "status"}, po.ARTICLE_STATUS_PUBLISHED, clause.Column{Name: "id"}, id).
		Take(&article)
	if article.Id == 0 {
		return nil
	}
	s.incrementViews(&article)
	return &article
}

func (s *ArticleService) QueryByAdmin(id int) *po.Article {
	var article po.Article
	s.db.Model(&po.Article{}).
		Preload("Category").Preload("Author").Preload("Urls").Preload("Tags").
		Where("? = ?", clause.Column{Name: "id"}, id).
		Take(&article)
	if article.Id == 0 {
		return nil
	}
	return &article
}

func (s *ArticleService) QueryPrevArticle(id int) *po.Article {
	var article po.Article
	s.db.Model(&po.Article{}).
		Preload("Urls").
		Where("? = ? AND ? < ?", clause.Column{Name: "status"}, po.ARTICLE_STATUS_PUBLISHED, clause.Column{Name: "id"}, id).
		Last(&article)
	if article.Id == 0 {
		return nil
	}
	return &article
}

func (s *ArticleService) QueryNextArticle(id int) *po.Article {
	var article po.Article
	s.db.Model(&po.Article{}).
		Preload("Urls").
		Where("? = ? AND ? > ?", clause.Column{Name: "status"}, po.ARTICLE_STATUS_PUBLISHED, clause.Column{Name: "id"}, id).
		First(&article)
	if article.Id == 0 {
		return nil
	}
	return &article
}

func (s *ArticleService) ListArticle() []po.Article {
	var articles []po.Article
	s.db.Model(&po.Article{}).
		Preload("Urls").
		Where("? = ?", clause.Column{Name: "status"}, po.ARTICLE_STATUS_PUBLISHED).
		Order(clause.OrderByColumn{
			Column: clause.Column{Name: "id"},
			Desc:   true,
		}).
		Find(&articles)
	return articles
}

func (s *ArticleService) PaginationByAdmin(
	title *string,
	categoryId *int,
	url *string,
	comment *bool,
	status *int,
	pageNum int,
	pageSize int,
	orderName *string,
	orderMethod *string,
) po.Pagination[po.Article] {
	db := s.db.Model(&po.Article{}).
		Joins("LEFT JOIN ? ON ? = ?", clause.Table{Name: "t_article_url"}, clause.Column{Name: "t_article.id"}, clause.Column{Name: "t_article_url.article_id"}).
		Joins("LEFT JOIN ? ON ? = ?", clause.Table{Name: "t_article_tag"}, clause.Column{Name: "t_article.id"}, clause.Column{Name: "t_article_tag.article_id"}).
		Distinct().
		Preload("Category").Preload("Urls")
	if title != nil {
		db = db.Where("upper(?) LIKE upper(?)", clause.Column{Name: "t_article.title"}, "%"+*title+"%")
	}
	if categoryId != nil {
		db = db.Where("? = ?", clause.Column{Name: "t_article.category_id"}, *categoryId)
	}
	if url != nil && *url != "" {
		db = db.Where("upper(?) LIKE upper(?)", clause.Column{Name: "t_article_url.url"}, "%"+*url+"%")
	}
	if comment != nil {
		var commentStatus int
		if *comment {
			commentStatus = po.ARTICLE_COMMENT_ENABLE
		} else {
			commentStatus = po.ARTICLE_COMMENT_DISABLE
		}
		db = db.Where("? = ?", clause.Column{Name: "t_article.comment_status"}, commentStatus)
	}
	if status != nil {
		db = db.Where("? = ?", clause.Column{Name: "t_article.status"}, *status)
	}
	db = db.Order(database.TableOrderBy(orderName, orderMethod, "id", true))
	return database.Pagination[po.Article](db, pageNum, pageSize)
}

func (s *ArticleService) PaginationArticleByPage(pageNum int) po.Pagination[po.Article] {
	db := s.db.Model(&po.Article{}).
		Preload("Category").Preload("Author").Preload("Urls").
		Where("? = ?", clause.Column{Name: "status"}, po.ARTICLE_STATUS_PUBLISHED).
		Order(clause.OrderByColumn{
			Column: clause.Column{Name: "id"},
			Desc:   true,
		})
	return database.Pagination[po.Article](db, pageNum, ARTICLE_PAGE_SIZE)
}

func (s *ArticleService) PaginationArticleByArchive(pageNum int) po.Pagination[po.Article] {
	db := s.db.Model(&po.Article{}).
		Preload("Urls").
		Where("? = ?", clause.Column{Name: "status"}, po.ARTICLE_STATUS_PUBLISHED).
		Order(clause.OrderByColumn{
			Column: clause.Column{Name: "id"},
			Desc:   true,
		})
	return database.Pagination[po.Article](db, pageNum, ARTICLE_PAGE_SIZE)
}

func (s *ArticleService) PaginationArticleByCategory(categoryId int, pageNum int) po.Pagination[po.Article] {
	db := s.db.Model(&po.Article{}).
		Preload("Urls").
		InnerJoins("Category", s.db.Where(&po.Category{Id: categoryId})).
		Where(&po.Article{Status: po.ARTICLE_STATUS_PUBLISHED}).
		Order(clause.OrderByColumn{
			Column: clause.Column{Name: "id"},
			Desc:   true,
		})
	return database.Pagination[po.Article](db, pageNum, ARTICLE_PAGE_SIZE)
}

func (s *ArticleService) PaginationArticleByTag(tagId int, pageNum int) po.Pagination[po.Article] {
	db := s.db.Model(&po.Article{}).
		Joins("LEFT JOIN ? ON ? = ?", clause.Table{Name: "t_article_tag"}, clause.Column{Name: "t_article.id"}, clause.Column{Name: "t_article_tag.article_id"}).
		Where("? = ? AND ? = ?", clause.Column{Name: "t_article.status"}, po.ARTICLE_STATUS_PUBLISHED, clause.Column{Name: "t_article_tag.tag_id"}, tagId).
		Order(clause.OrderByColumn{
			Column: clause.Column{Name: "id"},
			Desc:   true,
		})
	return database.Pagination[po.Article](db, pageNum, ARTICLE_PAGE_SIZE)
}

func (s *ArticleService) AddByAdmin(
	title string,
	categoryId int,
	content string,
	url *string,
	weight *int,
	commentStatus int,
	status int,
	tags *[]string,
) *po.Article {
	article := po.Article{}
	db := s.db.Begin()
	db = db.Model(&po.Article{}).
		Create(map[string]interface{}{
			"title":          title,
			"content":        content,
			"category_id":    categoryId,
			"author_id":      0,
			"weight":         weight,
			"comment_status": commentStatus,
			"status":         status,
			"update_time":    time.Now(),
		})
	db.Commit()
	return &article
}

func (s *ArticleService) UpdateByAdmin(
	id int,
	title string,
	categoryId int,
	content string,
	url *string,
	weight *int,
	commentStatus int,
	status int,
	tags *[]string,
) int {
	db := s.db.Begin()
	db = db.Model(&po.Article{}).
		Where("? = ?", clause.Column{Name: "id"}, id).
		Updates(map[string]interface{}{
			"title":          title,
			"content":        content,
			"category_id":    categoryId,
			"author_id":      0,
			"weight":         weight,
			"comment_status": commentStatus,
			"status":         status,
			"update_time":    time.Now(),
		})
	db.Commit()
	return int(db.RowsAffected)
}

func (s *ArticleService) UpdateStatusByAdmin(id int, status int) int {
	db := s.db.Model(&po.Article{}).
		Where("? = ?", clause.Column{Name: "id"}, id).
		Update("status", status)
	return int(db.RowsAffected)
}

func (s *ArticleService) UpdateCategoryByAdminDeleteCategory(deleteCategoryId int, defaultCategory int) {
	s.db.Model(&po.Article{}).
		Where("? = ?", clause.Column{Name: "category_id"}, deleteCategoryId).
		Update("category_id", defaultCategory)
}

func (s *ArticleService) DeleteByAdmin(id int) int {
	db := s.db.Delete(&po.Article{Id: id})
	return int(db.RowsAffected)
}

func (s *ArticleService) incrementViews(article *po.Article) {
	if article == nil || article.Id == 0 {
		return
	}
	article.Views++
	s.db.Model(&po.Article{Id: article.Id}).
		UpdateColumns(&po.Article{
			Views: article.Views,
		})
}
