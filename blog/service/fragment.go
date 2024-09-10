package service

import (
	"blog/database"
	"gorm.io/gorm"
)

type FragmentService struct {
	Service
	db *gorm.DB
}

func (s *FragmentService) OnInitService() {
	s.db = database.GetDB()
}
