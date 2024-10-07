package service

import (
	"blog/bean/po"
	"blog/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type LinkService struct {
	Service
	db *gorm.DB
}

func (s *LinkService) OnInitService() {
	s.db = database.GetDB()
}

func (s *LinkService) PaginationByAdmin(
	title *string,
	url *string,
	status *int,
	pageNum int,
	pageSize int,
	orderName *string,
	orderMethod *string,
) po.Pagination[po.Link] {
	db := s.db.Model(&po.Link{})
	if title != nil && *title != "" {
		db = db.Where("upper(?) LIKE upper(?)", clause.Column{Name: "title"}, "%"+*title+"%")
	}
	if url != nil && *url != "" {
		db = db.Where("upper(?) LIKE upper(?)", clause.Column{Name: "url"}, "%"+*url+"%")
	}
	if status != nil {
		db = db.Where("? = ?", clause.Column{Name: "status"}, *status)
	}
	db = db.Order(database.TableOrderBy(orderName, orderMethod, "id", true))
	return database.Pagination[po.Link](db, pageNum, pageSize)
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

func (s *LinkService) AddByAdmin(title string, url string, weight int, status int) bool {
	db := s.db.Model(&po.Link{}).
		Create(&po.Link{
			Title:      title,
			Url:        url,
			Weight:     weight,
			Status:     status,
			CreateTime: time.Now(),
		})
	if db.RowsAffected > 0 {
		refreshBlogCacheInfo()
	}
	return db.RowsAffected > 0
}

func (s *LinkService) UpdateByAdmin(id int, title string, url string, weight int, status int) bool {
	updateTime := time.Now()
	db := s.db.Model(&po.Link{}).
		Where("? = ?", clause.Column{Name: "id"}, id).
		Select("title", "url", "weight", "status", "update_time").
		Updates(&po.Link{
			Title:      title,
			Url:        url,
			Weight:     weight,
			Status:     status,
			UpdateTime: &updateTime,
		})
	if db.RowsAffected > 0 {
		refreshBlogCacheInfo()
	}
	return db.RowsAffected > 0
}

func (s *LinkService) DeleteByAdmin(id int) bool {
	db := s.db.Model(&po.Link{}).Delete(&po.Link{Id: id})
	if db.RowsAffected > 0 {
		refreshBlogCacheInfo()
	}
	return db.RowsAffected > 0
}
