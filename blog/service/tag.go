package service

import (
	"blog/bean/po"
	"blog/database"
	"gorm.io/gorm"
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

func (s *TagService) ListTag() []po.Tag {
	var tags []po.Tag
	s.db.Model(&po.Tag{}).
		Preload("Articles", s.db.Where("status = ?", po.ARTICLE_STATUS_PUBLISHED)).
		Order("id ASC").Find(&tags)
	return tags
}

func (s *TagService) QueryTag(name string) *po.Tag {
	var tag po.Tag
	s.db.
		Where("name = ?", name).
		Take(&tag)
	if tag.Id == 0 {
		return nil
	}
	return &tag
}
