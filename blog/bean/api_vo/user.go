package api_vo

import "time"

type UserVO struct {
	Id            int
	Username      string
	Avatar        string
	Nickname      string
	Description   string
	Email         string
	Accounts      UserAccountsVO
	LastLoginTime time.Time
	Status        int
	CreateTime    time.Time
	UpdateTime    time.Time
}

type UserAccountsVO struct {
	Github        string
	Weibo         string
	Google        string
	Twitter       string
	Facebook      string
	StackOverflow string
	Youtube       string
	Instagram     string
	Skype         string
}