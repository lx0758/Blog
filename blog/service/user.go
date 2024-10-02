package service

import (
	"blog/bean/po"
	"blog/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type UserService struct {
	Service
	db *gorm.DB
}

func (s *UserService) OnInitService() {
	s.db = database.GetDB()
}

func (s *UserService) ListByAdmin(username *string, nickname *string, status *int) []po.User {
	var users []po.User
	db := s.db.Model(&po.User{})
	if username != nil {
		db = db.Where("? = ?", clause.Column{Name: "username"}, username)
	}
	if nickname != nil {
		db = db.Where("? = ?", clause.Column{Name: "nickname"}, nickname)
	}
	if status != nil {
		db = db.Where("? = ?", clause.Column{Name: "status"}, status)
	}
	db.Find(&users)
	return users
}

func (s *UserService) QueryOwner() *po.User {
	return s.QueryUserById(1)
}

func (s *UserService) QueryUserById(id int) *po.User {
	var user po.User
	s.db.Model(&po.User{}).
		Where("? = ?", clause.Column{Name: "id"}, id).
		Take(&user)
	if user.Id == 0 {
		return nil
	}
	return &user
}

func (s *UserService) QueryUserByUsername(username string) *po.User {
	var user po.User
	s.db.Model(&po.User{}).
		Where("? = ?", clause.Column{Name: "username"}, username).
		Take(&user)
	if user.Id == 0 {
		return nil
	}
	return &user
}

func (s *UserService) UpdateUserLoginTime(user *po.User) {
	s.db.Model(&po.User{}).
		Where("? = ?", clause.Column{Name: "id"}, user.Id).
		Update("last_login_time", time.Now())
}

func (s *UserService) UpdateUserProfile(
	id int,
	avatar *string,
	nickname string,
	description *string,
	email *string,
	github *string,
	weibo *string,
	google *string,
	twitter *string,
	facebook *string,
	overflow *string,
	youtube *string,
	instagram *string,
	skype *string,
) bool {
	updateTime := time.Now()
	db := s.db.Model(&po.User{}).
		Where("? = ?", clause.Column{Name: "id"}, id).
		Updates(po.User{
			Avatar:      avatar,
			Nickname:    nickname,
			Description: description,
			Email:       email,
			Accounts: &po.UserAccounts{
				Github:        github,
				Weibo:         weibo,
				Google:        google,
				Twitter:       twitter,
				Facebook:      facebook,
				StackOverflow: overflow,
				Youtube:       youtube,
				Instagram:     instagram,
				Skype:         skype,
			},
			UpdateTime: &updateTime,
		})
	return db.RowsAffected > 0
}

func (s *UserService) UpdateUserPassword(id int, oldPassword string, newPassword string) bool {
	user := s.QueryUserById(id)
	if user == nil {
		return false
	}
	if !s.VerifyPassword(user, oldPassword) {
		return false
	}
	newHash, err := hashPassword(newPassword)
	if err != nil {
		return false
	}
	db := s.db.Model(&po.User{}).
		Where("? = ?", clause.Column{Name: "id"}, user.Id).
		Update("password", newHash).
		Update("update_time", time.Now())
	return db.RowsAffected > 0
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
