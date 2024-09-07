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

func (s *ArticleService) QueryArticle(id int64) *po.Article {
	var article po.Article
	db := s.db.Preload("Category").Preload("Author").Preload("Urls").Preload("Tags")
	db.First(&article, "id = ?", id)
	return &article
}

func (s *ArticleService) ListArticle() []po.Article {
	var articles []po.Article
	db := s.db.Preload("Urls")
	db = db.Where("status = ?", []any{1})
	db.Order("id DESC").Find(&articles)
	return articles
}

func (s *ArticleService) PaginationArticle(pageNum int) ([]po.Article, po.Pagination) {
	var articles []po.Article
	db := s.db.Preload("Category").Preload("Author").Preload("Urls")
	db = db.Where("status = ?", []any{1})
	db = db.Order("id DESC").Offset((pageNum - 1) * ARTICLE_PAGE_SIZE).Limit(ARTICLE_PAGE_SIZE).Find(&articles)
	pagination := database.Pagination(db, pageNum, ARTICLE_PAGE_SIZE)
	return articles, pagination
}
