package api_vo

import (
	"blog/bean/po"
	"time"
)

type UserVO struct {
	Id            int            `json:"id"`
	Username      string         `json:"username"`
	Avatar        *string        `json:"avatar"`
	Nickname      string         `json:"nickname"`
	Description   *string        `json:"description"`
	Email         *string        `json:"email"`
	Accounts      UserAccountsVO `json:"accounts"`
	LastLoginTime *time.Time     `json:"lastLoginTime"`
	Status        int            `json:"status"`
	CreateTime    time.Time      `json:"createTime"`
	UpdateTime    *time.Time     `json:"updateTime"`
}

func (v *UserVO) From(user po.User) {
	v.Id = user.Id
	v.Username = user.Username
	v.Avatar = user.Avatar
	v.Nickname = user.Nickname
	v.Description = user.Description
	v.Email = user.Email
	v.Accounts.From(user.Accounts)
	v.LastLoginTime = user.LastLoginTime
	v.Status = user.Status
	v.CreateTime = user.CreateTime
	v.UpdateTime = user.UpdateTime
}

type UserAccountsVO struct {
	Github        string `json:"github"`
	Weibo         string `json:"weibo"`
	Google        string `json:"google"`
	Twitter       string `json:"twitter"`
	Facebook      string `json:"facebook"`
	StackOverflow string `json:"stackOverflow"`
	Youtube       string `json:"youtube"`
	Instagram     string `json:"instagram"`
	Skype         string `json:"skype"`
}

func (v *UserAccountsVO) From(accounts *po.UserAccounts) {
	if accounts == nil {
		return
	}
	v.Github = *accounts.Github
	v.Weibo = *accounts.Weibo
	v.Google = *accounts.Google
	v.Twitter = *accounts.Twitter
	v.Facebook = *accounts.Facebook
	v.StackOverflow = *accounts.StackOverflow
	v.Youtube = *accounts.Youtube
	v.Instagram = *accounts.Instagram
	v.Skype = *accounts.Skype
}
