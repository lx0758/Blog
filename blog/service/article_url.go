package service

import (
	"blog/bean/po"
	"blog/database"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ArticleUrlService struct {
	Service
	db *gorm.DB
}

func (s *ArticleUrlService) OnInitService() {
	s.db = database.GetDB()
}

func (s *ArticleUrlService) QueryArticleUrlByUrl(url string) *po.ArticleUrl {
	var articleUrl po.ArticleUrl
	s.db.Model(&po.ArticleUrl{}).
		Where("? = ?", clause.Column{Name: "url"}, url).
		Take(&articleUrl)
	if articleUrl.Url == "" {
		return nil
	}
	return &articleUrl
}

func (s *ArticleUrlService) QueryArticleUrlById(id int) *po.ArticleUrl {
	var articleUrl po.ArticleUrl
	s.db.Model(&po.ArticleUrl{}).
		Where("? = ? AND ? = ?", clause.Column{Name: "status"}, po.ARTICLE_URL_STATUS_LAST, clause.Column{Name: "article_id"}, id).
		Take(&articleUrl)
	if articleUrl.Url == "" {
		return nil
	}
	return &articleUrl
}
