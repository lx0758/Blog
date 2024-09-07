package service

import (
	"blog/bean/po"
	"blog/database"
	"gorm.io/gorm"
)

type UserService struct {
	Service
	db *gorm.DB
}

func (s *UserService) OnInitService() {
	s.db = database.GetDB()
}

func (s *UserService) QueryOwner() *po.User {
	var user po.User
	s.db.Model(&po.User{}).
		Where("id = ?", 1).
		Take(&user)
	if user.Id == 0 {
		return nil
	}
	return &user
}
