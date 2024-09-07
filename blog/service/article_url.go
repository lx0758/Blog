package service

import (
	"blog/bean/po"
	"blog/database"
	"gorm.io/gorm"
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
		Where("url = ?", url).
		Take(&articleUrl)
	if articleUrl.Url == "" {
		return nil
	}
	return &articleUrl
}

func (s *ArticleUrlService) QueryArticleUrlById(id int) *po.ArticleUrl {
	var articleUrl po.ArticleUrl
	s.db.Model(&po.ArticleUrl{}).
		Where("status = ? AND article_id = ?", po.ARTICLE_URL_STATUS_LAST, id).
		Take(&articleUrl)
	if articleUrl.Url == "" {
		return nil
	}
	return &articleUrl
}
