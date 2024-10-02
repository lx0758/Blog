package service

import (
	"blog/bean/po"
	"blog/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type CategoryService struct {
	Service
	db *gorm.DB

	articleService *ArticleService
}

func (s *CategoryService) OnInitService() {
	s.db = database.GetDB()

	s.articleService = GetService[*ArticleService](s.articleService)
}

func (s *CategoryService) Count() int {
	var count int64
	s.db.Model(&po.Category{}).
		Count(&count)
	return int(count)
}

func (s *CategoryService) ListByHtml() []po.Category {
	var categories []po.Category
	s.db.Model(&po.Category{}).
		Preload("Articles", s.db.Where("? = ?", clause.Column{Name: "status"}, po.ARTICLE_STATUS_PUBLISHED)).
		Order(clause.OrderByColumn{
			Column: clause.Column{Name: "id"},
			Desc:   false,
		}).
		Find(&categories)
	return categories
}

func (s *CategoryService) PaginationByAdmin(
	pageNum int,
	pageSize int,
	_ *string,
	_ *string,
) po.Pagination[po.Category] {
	db := s.db.Model(&po.Category{}).
		Preload("Articles", s.db.Where("? = ?", clause.Column{Name: "status"}, po.ARTICLE_STATUS_PUBLISHED)).
		Order(clause.OrderByColumn{
			Column: clause.Column{Name: "id"},
			Desc:   false,
		})
	return database.Pagination[po.Category](db, pageNum, pageSize)
}

func (s *CategoryService) QueryCategory(name string) *po.Category {
	var category po.Category
	s.db.Model(&po.Category{}).
		Where("? = ?", clause.Column{Name: "name"}, name).
		Take(&category)
	if category.Id == 0 {
		return nil
	}
	return &category
}

func (s *CategoryService) QueryDefaultCategory() *po.Category {
	return s.QueryCategory("默认分类")
}

func (s *CategoryService) AddByAdmin(name string) bool {
	db := s.db.Model(&po.Category{}).
		Create(&po.Category{
			Name:       name,
			CreateTime: time.Now(),
		})
	return db.RowsAffected > 0
}

func (s *CategoryService) UpdateByAdmin(id int, name string) bool {
	db := s.db.Model(&po.Category{}).
		Where("? = ?", clause.Column{Name: "id"}, id).
		Update("name", name)
	return db.RowsAffected > 0
}

func (s *CategoryService) DeleteByAdmin(id int) bool {
	category := s.QueryDefaultCategory()
	if category == nil {
		return false
	}

	s.articleService.UpdateCategoryByAdminDeleteCategory(id, category.Id)

	db := s.db.Delete(&po.Category{Id: id})
	return db.RowsAffected > 0
}
