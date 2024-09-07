package service

import (
	"blog/bean/po"
	"blog/database"
	"gorm.io/gorm"
)

type CategoryService struct {
	Service
	db *gorm.DB
}

func (s *CategoryService) OnInitService() {
	s.db = database.GetDB()
}

func (s *CategoryService) Count() int {
	var count int64
	s.db.Model(&po.Category{}).
		Count(&count)
	return int(count)
}

func (s *CategoryService) ListCategory() []po.Category {
	var categories []po.Category
	s.db.Model(&po.Category{}).
		Preload("Articles", s.db.Where("status = ?", po.ARTICLE_STATUS_PUBLISHED)).
		Order("id DESC").
		Find(&categories)
	return categories
}

func (s *CategoryService) QueryCategory(name string) *po.Category {
	var category po.Category
	s.db.Model(&po.Category{}).
		Where("name = ?", name).
		Take(&category)
	if category.Id == 0 {
		return nil
	}
	return &category
}
