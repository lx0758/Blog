package service

import (
	"blog/bean/po"
	"blog/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type UrlService struct {
	Service
	db *gorm.DB
}

func (s *UrlService) OnInitService() {
	s.db = database.GetDB()
}

func (s *UrlService) PaginationByAdmin(
	key *string,
	url *string,
	description *string,
	status *int,
	pageNum int,
	pageSize int,
	orderName *string,
	orderMethod *string,
) po.Pagination[po.Url] {
	db := s.db.Model(&po.Url{}).
		Preload("Author")
	if key != nil && *key != "" {
		db = db.Where("upper(?) LIKE upper(?)", clause.Column{Name: "key"}, "%"+*key+"%")
	}
	if url != nil && *url != "" {
		db = db.Where("upper(?) LIKE upper(?)", clause.Column{Name: "url"}, "%"+*key+"%")
	}
	if description != nil && *description != "" {
		db = db.Where("upper(?) LIKE upper(?)", clause.Column{Name: "description"}, "%"+*description+"%")
	}
	if status != nil {
		db = db.Where("? = ?", clause.Column{Name: "status"}, *status)
	}
	db = db.Order(database.TableOrderBy(orderName, orderMethod, "id", true))
	return database.Pagination[po.Url](db, pageNum, pageSize)
}

func (s *UrlService) QueryUrl(key string) *po.Url {
	var url po.Url
	s.db.Model(&po.Url{}).
		Where("? = ? AND ? = ?", clause.Column{Name: "status"}, po.URL_STATUS_PUBLISHED, clause.Column{Name: "key"}, key).
		Take(&url)
	if url.Id == 0 {
		return nil
	}
	s.incrementViews(&url)
	return &url
}

func (s *UrlService) AddByAdmin(userId int, key string, url string, description string, status int) bool {
	db := s.db.Model(&po.Url{}).
		Create(&po.Url{
			Key:         key,
			Url:         url,
			Description: description,
			AuthorId:    userId,
			Status:      status,
			CreateTime:  time.Now(),
		})
	return db.RowsAffected > 0
}

func (s *UrlService) UpdateByAdmin(id int, key string, url string, description string, status int) bool {
	updateTime := time.Now()
	db := s.db.Model(&po.Url{}).
		Where("? = ?", clause.Column{Name: "id"}, id).
		Select("key", "url", "description", "status", "views", "update_time").
		Updates(&po.Url{
			Key:         key,
			Url:         url,
			Description: description,
			Status:      status,
			Views:       0,
			UpdateTime:  &updateTime,
		})
	return db.RowsAffected > 0
}

func (s *UrlService) DeleteByAdmin(id int) bool {
	db := s.db.Model(&po.Url{}).
		Where("? = ?", clause.Column{Name: "id"}, id).
		Delete(nil)
	return db.RowsAffected > 0
}

func (s *UrlService) incrementViews(url *po.Url) {
	if url == nil || url.Id == 0 {
		return
	}
	s.db.Model(&po.Url{Id: url.Id}).
		Update("views", gorm.Expr("views + 1")).
		Update("total_views", gorm.Expr("total_views + 1"))
}
