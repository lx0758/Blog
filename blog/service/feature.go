package service

import (
	"blog/database"
	"gorm.io/gorm"
)

type FeatureService struct {
	Service
	db *gorm.DB
}

func (s *FeatureService) OnInitService() {
	s.db = database.GetDB()
}
