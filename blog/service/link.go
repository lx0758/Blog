package service

import (
	"blog/bean/po"
	"blog/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
		Where("? = ?", clause.Column{Name: "status"}, po.LINK_STATUS_PUBLISHED).
		Order(clause.OrderByColumn{
			Column: clause.Column{Name: "id"},
			Desc:   false,
		}).
		Find(&urls)
	return urls
}
