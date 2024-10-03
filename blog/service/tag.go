package service

import (
	"blog/bean/po"
	"blog/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type TagService struct {
	Service
	db *gorm.DB
}

func (s *TagService) OnInitService() {
	s.db = database.GetDB()
}

func (s *TagService) Count() int {
	var count int64
	db := s.db.Model(&po.Tag{})
	db.Count(&count)
	return int(count)
}

func (s *TagService) PaginationByAdmin(
	name *string,
	pageNum int,
	pageSize int,
	orderName *string,
	orderMethod *string,
) po.Pagination[po.Tag] {
	articleDb := s.db.Model(&po.Article{}).
		Select("COUNT(?)", clause.Column{Name: "id"}).
		Joins("LEFT JOIN ? ON ? = ?", clause.Table{Name: "t_article_tag"}, clause.Column{Name: "t_article.id"}, clause.Column{Name: "t_article_tag.article_id"}).
		Where("? = ?", clause.Column{Name: "t_article_tag.tag_id"}, clause.Column{Name: "t_tag.id"})
	db := s.db.Model(&po.Tag{}).
		Select("*, (?) AS ?", articleDb, clause.Column{Name: "article_count"})
	if name != nil && *name != "" {
		db = db.Where("upper(?) LIKE upper(?)", clause.Column{Table: "t_tag", Name: "name"}, "%"+*name+"%")
	}
	db = db.Order(database.TableOrderBy(orderName, orderMethod, "id", true))
	return database.Pagination[po.Tag](db, pageNum, pageSize)
}

func (s *TagService) ListTag() []po.Tag {
	var tags []po.Tag
	articleDb := s.db.Model(&po.Article{}).
		Select("COUNT(?)", clause.Column{Name: "id"}).
		Joins("LEFT JOIN ? ON ? = ?", clause.Table{Name: "t_article_tag"}, clause.Column{Name: "t_article.id"}, clause.Column{Name: "t_article_tag.article_id"}).
		Where("? = ?", clause.Column{Name: "t_article.status"}, po.ARTICLE_STATUS_PUBLISHED).
		Where("? = ?", clause.Column{Name: "t_article_tag.tag_id"}, clause.Column{Name: "t_tag.id"})
	s.db.Model(&po.Tag{}).
		Select("*, (?) AS ?", articleDb, clause.Column{Name: "article_count"}).
		Order(clause.OrderByColumn{
			Column: clause.Column{Name: "id"},
			Desc:   false,
		}).
		Find(&tags)
	return tags
}

func (s *TagService) QueryTag(name string) *po.Tag {
	var tag po.Tag
	s.db.
		Where("? = ?", clause.Column{Name: "name"}, name).
		Take(&tag)
	if tag.Id == 0 {
		return nil
	}
	return &tag
}

func (s *TagService) AddByAdmin(name string) bool {
	db := s.db.Model(&po.Tag{}).
		Create(&po.Tag{
			Name:       name,
			CreateTime: time.Now(),
		})
	return db.RowsAffected > 0
}

func (s *TagService) UpdateByAdmin(id int, name string) bool {
	updateTime := time.Now()
	db := s.db.Model(&po.Tag{}).
		Where("? = ?", clause.Column{Name: "id"}, id).
		Select("name", "update_time").
		Updates(&po.Tag{
			Name:       name,
			UpdateTime: &updateTime,
		})
	return db.RowsAffected > 0
}

func (s *TagService) DeleteByAdmin(id int) bool {
	db := s.db.Model(&po.Tag{}).Delete(&po.Tag{Id: id})
	if db.RowsAffected > 0 {
		s.db.Table("t_article_tag").Delete("? = ?", clause.Column{Name: "tag_id"}, id)
	}
	return db.RowsAffected > 0
}
