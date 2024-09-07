package service

import (
	"blog/bean/po"
	"blog/database"
	"gorm.io/gorm"
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

func (s *ArticleService) PaginationArticleByAdmin(pageNum int) po.Pagination[po.Article] {
	db := s.db.Model(&po.Article{}).
		Preload("Category").Preload("Urls").
		Order("id DESC")
	return database.Pagination[po.Article](db, pageNum, ARTICLE_PAGE_SIZE)
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
