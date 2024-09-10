package service

import (
	"blog/bean/po"
	"blog/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type UserService struct {
	Service
	db *gorm.DB
}

func (s *UserService) OnInitService() {
	s.db = database.GetDB()
}

func (s *UserService) QueryOwner() *po.User {
	return s.QueryUserById(1)
}

func (s *UserService) QueryUserById(id int) *po.User {
	var user po.User
	s.db.Model(&po.User{}).
		Where("id = ?", id).
		Take(&user)
	if user.Id == 0 {
		return nil
	}
	return &user
}

func (s *UserService) QueryUserByUsername(username string) *po.User {
	var user po.User
	s.db.Model(&po.User{}).
		Where("username = ?", username).
		Take(&user)
	if user.Id == 0 {
		return nil
	}
	return &user
}

func (s *UserService) UpdateUserLoginTime(user *po.User) {
	s.db.Model(&po.User{}).
		Where("id = ?", user.Id).
		Update("last_login_time", time.Now())
}

func (s *UserService) VerifyPassword(user *po.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
