package service

import (
	"blog/bean/po"
	"blog/database"
	"blog/util"
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
		Where("status = ?", po.ARTICLE_STATUS_PUBLISHED).
		Count(&count)
	return int(count)
}

func (s *ArticleService) QueryArticle(id int) *po.Article {
	var article po.Article
	s.db.Model(&po.Article{}).
		Preload("Category").Preload("Author").Preload("Urls").Preload("Tags").
		Where("status = ? AND id = ?", po.ARTICLE_STATUS_PUBLISHED, id).
		Take(&article)
	if article.Id == 0 {
		return nil
	}
	s.incrementViews(&article)
	return &article
}

func (s *ArticleService) QueryArticleByAdmin(id int) *po.Article {
	var article po.Article
	s.db.Model(&po.Article{}).
		Preload("Category").Preload("Author").Preload("Urls").Preload("Tags").
		Where("id = ?", id).
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
		Where("status = ? AND id < ?", po.ARTICLE_STATUS_PUBLISHED, id).
		Order("id DESC").
		First(&article)
	if article.Id == 0 {
		return nil
	}
	return &article
}

func (s *ArticleService) QueryNextArticle(id int) *po.Article {
	var article po.Article
	s.db.Model(&po.Article{}).
		Preload("Urls").
		Where("status = ? AND id > ?", po.ARTICLE_STATUS_PUBLISHED, id).
		Order("id DESC").
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
		Where("status = ?", po.ARTICLE_STATUS_PUBLISHED).
		Order("id DESC").
		Find(&articles)
	return articles
}

func (s *ArticleService) PaginationArticleByAdmin(
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
		Joins("left join t_article_url on t_article.id = t_article_url.article_id").
		Distinct().
		Preload("Category").Preload("Urls")
	if title != nil {
		db = db.Where("upper(t_article.title) LIKE upper(?)", "%"+*title+"%")
	}
	if categoryId != nil {
		db = db.Where("t_article.category_id = ?", *categoryId)
	}
	if url != nil {
		db = db.Where("upper(t_article.url) LIKE upper(?)", "%"+*url+"%")
	}
	if comment != nil {
		if *comment {
			db = db.Where("t_article.comment_status = ?", po.ARTICLE_COMMENT_ENABLE)
		} else {
			db = db.Where("t_article.comment_status = ?", po.ARTICLE_COMMENT_DISABLE)
		}
	}
	if status != nil {
		db = db.Where("t_article.status = ?", *status)
	}
	if orderName != nil && orderMethod != nil {
		db = db.Order(clause.OrderBy{Columns: []clause.OrderByColumn{
			{Column: clause.Column{
				Name: util.ToSnakeCase(*orderName)},
				Desc: util.If(*orderMethod == "descending", true, false),
			},
			{Column: clause.Column{Name: "id"}, Desc: true},
		}})
	} else {
		db = db.Order("id DESC")
	}
	return database.Pagination[po.Article](db, pageNum, pageSize)
}

func (s *ArticleService) PaginationArticleByPage(pageNum int) po.Pagination[po.Article] {
	db := s.db.Model(&po.Article{}).
		Preload("Category").Preload("Author").Preload("Urls").
		Where("status = ?", po.ARTICLE_STATUS_PUBLISHED).
		Order("id DESC")
	return database.Pagination[po.Article](db, pageNum, ARTICLE_PAGE_SIZE)
}

func (s *ArticleService) PaginationArticleByArchive(pageNum int) po.Pagination[po.Article] {
	db := s.db.Model(&po.Article{}).
		Preload("Urls").
		Where("status = ?", po.ARTICLE_STATUS_PUBLISHED).
		Order("id DESC")
	return database.Pagination[po.Article](db, pageNum, ARTICLE_PAGE_SIZE)
}

func (s *ArticleService) PaginationArticleByCategory(categoryId int, pageNum int) po.Pagination[po.Article] {
	db := s.db.Model(&po.Article{}).
		Preload("Urls").
		InnerJoins("Category", s.db.Where(&po.Category{Id: categoryId})).
		Where(&po.Article{Status: po.ARTICLE_STATUS_PUBLISHED}).
		Order("id DESC")
	return database.Pagination[po.Article](db, pageNum, ARTICLE_PAGE_SIZE)
}

func (s *ArticleService) PaginationArticleByTag(tagId int, pageNum int) po.Pagination[po.Article] {
	db := s.db.Model(&po.Article{}).
		Joins("LEFT JOIN t_article_tag ON t_article.id = t_article_tag.article_id").
		Where("t_article.status = ? AND t_article_tag.tag_id = ?", po.ARTICLE_STATUS_PUBLISHED, tagId).
		Order("id DESC")
	return database.Pagination[po.Article](db, pageNum, ARTICLE_PAGE_SIZE)
}

func (s *ArticleService) AddArticleByAdmin(
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

func (s *ArticleService) UpdateArticleByAdmin(
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
		Where("id = ?", id).
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

func (s *ArticleService) UpdateArticleStatusByAdmin(id int, status int) int {
	db := s.db.Model(&po.Article{}).
		Where("id = ?", id).
		Update("status", status)
	return int(db.RowsAffected)
}

func (s *ArticleService) DeleteArticleByAdmin(id int) int {
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
