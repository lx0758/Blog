package service

import (
	"blog/bean/po"
	"blog/database"
	"gorm.io/gorm"
)

type UrlService struct {
	Service
	db *gorm.DB
}

func (s *UrlService) OnInitService() {
	s.db = database.GetDB()
}

func (s *UrlService) QueryUrl(key string) *po.Url {
	var url po.Url
	s.db.Model(&po.Url{}).
		Where("status = ? AND key = ?", po.URL_STATUS_PUBLISHED, key).
		Take(&url)
	if url.Id == 0 {
		return nil
	}
	s.incrementViews(&url)
	return &url
}

func (s *UrlService) incrementViews(url *po.Url) {
	if url == nil || url.Id == 0 {
		return
	}
	url.Views++
	url.TotalViews++
	s.db.Model(&po.Url{Id: url.Id}).
		UpdateColumns(&po.Url{
			Views:      url.Views,
			TotalViews: url.TotalViews,
		})
}