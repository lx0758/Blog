package service

import (
	"blog/bean/po"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	SESSION_KEY_USER_ID = "userId"
)

type SessionService struct {
	Service

	userService *UserService
}

func (s *SessionService) OnInitService() {
	s.userService = GetService[*UserService](s.userService)
}

func (s *SessionService) Login(session sessions.Session, username string, password string) error {
	user := s.userService.QueryUserByUsername(username)
	if user == nil || !s.userService.VerifyPassword(user, password) {
		return errors.New("账号或密码错误")
	}
	if user.Status != po.USER_STATUS_ENABLED {
		return errors.New("账号已禁用")
	}
	s.userService.UpdateUserLoginTime(user)
	session.Set(SESSION_KEY_USER_ID, user.Id)
	_ = session.Save()
	return nil
}

func (s *SessionService) Logout(session sessions.Session) {
	session.Delete(SESSION_KEY_USER_ID)
	_ = session.Save()
}

func (s *SessionService) GetLoggedUser(context *gin.Context) *po.User {
	session := sessions.Default(context)
	userId := session.Get(SESSION_KEY_USER_ID)
	userIdValue, ok := userId.(int)
	if !ok {
		return nil
	}
	user := s.userService.QueryUserById(userIdValue)
	if user == nil {
		return nil
	}
	if user.Status != po.USER_STATUS_ENABLED {
		return nil
	}
	return user
}
