package service

import (
	"blog/bean/po"
	"blog/database"
	"encoding/json"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type FeatureService struct {
	Service
	db *gorm.DB
}

func (s *FeatureService) OnInitService() {
	s.db = database.GetDB()
}

func (s *FeatureService) QuerySmtp() *po.SMTPFeature {
	feature := po.Feature{}
	s.db.Model(&po.Feature{}).
		Where("? = ?", clause.Column{Name: "key"}, po.FEATURE_KEY_SMTP).
		Take(&feature)
	if feature.Value == nil {
		return nil
	}
	var smtpFeature po.SMTPFeature
	err := json.Unmarshal([]byte(*feature.Value), &smtpFeature)
	if err != nil {
		return nil
	}
	return &smtpFeature
}

func (s *FeatureService) UpdateSmtp(
	enable bool,
	hostname *string,
	port *int,
	ssl *bool,
	username *string,
	password *string,
	fromName *string,
	fromEmail *string,
) bool {
	smtpFeature := po.SMTPFeature{
		Enable:    enable,
		Hostname:  *hostname,
		Port:      *port,
		Ssl:       *ssl,
		Username:  *username,
		Password:  *password,
		FromName:  *fromName,
		FromEmail: *fromEmail,
	}
	bytes, err := json.Marshal(smtpFeature)
	if err != nil {
		return false
	}
	value := string(bytes)

	updateTime := time.Now()
	db := s.db.Model(&po.Feature{}).
		Where("? = ?", clause.Column{Name: "key"}, po.FEATURE_KEY_SMTP).
		Select("value", "update_time").
		Updates(&po.Feature{
			Value:      &value,
			UpdateTime: &updateTime,
		})
	return db.RowsAffected > 0
}
