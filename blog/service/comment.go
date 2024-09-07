package service

import (
	"blog/database"
	"gorm.io/gorm"
)

type CommentService struct {
	Service
	db *gorm.DB
}

func (s *CommentService) OnInitService() {
	s.db = database.GetDB()
}
