package service

import (
	"blog/bean/po"
	"blog/database"
	"gorm.io/gorm"
)

type LinkService struct {
	Service
	db *gorm.DB
}

func (s *LinkService) OnInitService() {
	s.db = database.GetDB()
}

func (s *LinkService) ListLink() []po.Link {
	var urls []po.Link
	s.db.Model(&po.Link{}).
		Where("status = ?", po.LINK_STATUS_PUBLISHED).
		Order("id ASC").
		Find(&urls)
	return urls
}
