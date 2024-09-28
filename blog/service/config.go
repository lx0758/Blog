package service

import (
	"blog/bean/po"
	"blog/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ConfigService struct {
	Service
	db *gorm.DB
}

func (s *ConfigService) OnInitService() {
	s.db = database.GetDB()
}

func (s *ConfigService) QueryConfigMap() map[string]string {
	var configs []po.Config
	s.db.Model(&po.Config{}).
		Find(&configs)
	configMap := make(map[string]string)
	for _, config := range configs {
		configMap[config.Key] = config.Value
	}
	return configMap
}

func (s *ConfigService) QueryConfig(key string) *po.Config {
	var config po.Config
	s.db.Model(&po.Config{}).
		Where("? = ?", clause.Column{Name: "key"}, key).
		Take(&config)
	if config.Key == "" {
		return nil
	}
	return &config
}

func (s *ConfigService) ListConfig() []po.Config {
	var configs []po.Config
	s.db.Model(&po.Config{}).
		Order(clause.OrderByColumn{
			Column: clause.Column{Name: "key"},
			Desc:   false,
		}).
		Find(&configs)
	return configs
}
