package service

import (
	"blog/bean/po"
	"blog/database"
	"blog/logger"
	"blog/res"
	"blog/util"
	"errors"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"html/template"
	"strings"
)

type EmailService struct {
	Service
	db *gorm.DB

	featureService *FeatureService
	userService    *UserService
	blogService    *BlogService
	articleService *ArticleService
	commentService *CommentService
}

type recipient struct {
	Name  *string
	Email string
}

func (s *EmailService) OnInitService() {
	s.db = database.GetDB()

	s.featureService = GetService[*FeatureService](s.featureService)
	s.userService = GetService[*UserService](s.userService)
	s.blogService = GetService[*BlogService](s.blogService)
	s.articleService = GetService[*ArticleService](s.articleService)
	s.commentService = GetService[*CommentService](s.commentService)
}

func (s *EmailService) SendTestEmail(email string) error {
	to := []recipient{
		{Email: email},
	}
	body := s.renderEmailTemplate("comment_test", nil)
	err := s.sendEmail(to, nil, "博客邮件发送测试", body)
	if err != nil {
		return errors.New("邮件发送失败")
	}
	return nil
}

func (s *EmailService) SendCommentedEmail(
	articleId int,
	parentId *int,
	author string,
	email string,
	ip string,
	location string,
	ua string,
	content string,
) {
	owner := s.userService.QueryOwner()
	if owner == nil || owner.Email == nil {
		return
	}

	article := s.articleService.QueryByAdmin(articleId)
	if article == nil {
		return
	}

	var parentComment *po.Comment
	if parentId != nil {
		parentComment = s.commentService.QueryByAdmin(*parentId)
	}

	to := []recipient{
		{Name: &owner.Nickname, Email: *owner.Email},
	}
	data := map[string]any{
		"ArticleOwner":         util.If(owner.Nickname != "", owner.Nickname, "博客管理员"),
		"ArticleTitle":         article.Title,
		"ArticleUrl":           "https://" + s.blogService.GetCacheBlog().SiteDomain + "/article/" + article.GetSafeUrl(),
		"ParentCommentAuthor":  nil,
		"ParentCommentEmail":   nil,
		"ParentCommentContent": nil,
		"CommentAuthor":        author,
		"CommentEmail":         email,
		"CommentContent":       content,
		"CommentIp":            ip,
		"CommentLocation":      location,
		"CommentUa":            ua,
	}
	if parentComment != nil {
		data["ParentCommentAuthor"] = parentComment.Author
		data["ParentCommentEmail"] = parentComment.Email
		data["ParentCommentContent"] = parentComment.Content
	}
	body := s.renderEmailTemplate("comment_notify", data)
	_ = s.sendEmail(to, nil, "文章评论通知", body)
}

func (s *EmailService) SendCommentReplayEmail(
	author string,
	email string,
	content string,
	articleId int,
	parentId int,
) {
	article := s.articleService.QueryByAdmin(articleId)
	if article == nil {
		return
	}

	var parentComment *po.Comment
	recipients := make([]recipient, 0)
	for {
		comment := po.Comment{}
		s.db.Model(&po.Comment{}).
			Where("? = ?", clause.Column{Name: "id"}, parentId).
			First(&comment)
		recipients = append(recipients, recipient{Name: &comment.Author, Email: comment.Email})
		if comment.Id == parentId {
			parentComment = &comment
		}
		if comment.ParentId == nil || *comment.ParentId == 0 {
			break
		}
		parentId = *comment.ParentId
	}
	// 抄收自己
	recipients = append(recipients, recipient{Name: &author, Email: email})

	to := recipients[0:1]
	cc := recipients[1:]
	data := map[string]any{
		"ReplayAuthor":         author,
		"ReplayContent":        content,
		"ArticleTitle":         article.Title,
		"ArticleUrl":           "https://" + s.blogService.GetCacheBlog().SiteDomain + "/article/" + article.GetSafeUrl(),
		"ParentCommentAuthor":  parentComment.Author,
		"ParentCommentEmail":   parentComment.Email,
		"ParentCommentContent": parentComment.Content,
	}
	body := s.renderEmailTemplate("comment_reply", data)
	_ = s.sendEmail(to, &cc, "评论回复通知", body)
}

func (s *EmailService) sendEmail(to []recipient, cc *[]recipient, subject string, content string) error {
	smtpFeature := s.featureService.QuerySmtp()
	if smtpFeature == nil {
		return errors.New("SMTP 未配置")
	}
	if !smtpFeature.Enable {
		return errors.New("SMTP 未开启")
	}

	message := gomail.NewMessage()
	message.SetHeader("From", message.FormatAddress(smtpFeature.FromEmail, smtpFeature.FromName))
	message.SetHeader("To", s.recipientToStrings(message, to)...)
	if cc != nil {
		message.SetHeader("Cc", s.recipientToStrings(message, *cc)...)
	}
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", content)

	dialer := gomail.NewDialer(smtpFeature.Hostname, smtpFeature.Port, smtpFeature.Username, smtpFeature.Password)
	dialer.SSL = smtpFeature.Ssl
	err := dialer.DialAndSend(message)
	if err != nil {
		logger.Printf("send email to %s(cc:%s) error:%s\n", message.GetHeader("To"), message.GetHeader("Cc"), err)
	} else {
		logger.Printf("send email to %s(cc:%s) succeed.", message.GetHeader("To"), message.GetHeader("Cc"))
	}
	return err
}

func (s *EmailService) renderEmailTemplate(name string, data map[string]any) string {
	content := strings.Builder{}
	tmpl, _ := template.ParseFS(res.TemplatesFS, "email/"+name+".gohtml")
	_ = tmpl.Execute(&content, data)
	return content.String()
}

func (s *EmailService) recipientToStrings(message *gomail.Message, recipients []recipient) []string {
	rs := make([]string, 0)
	for _, r := range recipients {
		if r.Name != nil {
			rs = append(rs, message.FormatAddress(r.Email, *r.Name))
		} else {
			rs = append(rs, r.Email)
		}
	}
	return rs
}
