package service

import (
	"blog/database"
	"gorm.io/gorm"
)

type FileService struct {
	Service
	db *gorm.DB
}

func (s *FileService) OnInitService() {
	s.db = database.GetDB()
}
