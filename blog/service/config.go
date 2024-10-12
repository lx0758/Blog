package service

import (
	"blog/bean/po"
	"blog/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type ConfigService struct {
	Service
	db *gorm.DB
}

func (s *ConfigService) OnInitService() {
	s.db = database.GetDB()
}

func (s *ConfigService) PaginationByAdmin(
	key *string,
	value *string,
	description *string,
	pageNum int,
	pageSize int,
	orderName *string,
	orderMethod *string,
) po.Pagination[po.Config] {
	db := s.db.Model(&po.Config{})
	if key != nil && *key != "" {
		db = db.Where("upper(?) LIKE upper(?)", clause.Column{Name: "key"}, "%"+*key+"%")
	}
	if value != nil && *value != "" {
		db = db.Where("upper(?) LIKE upper(?)", clause.Column{Name: "value"}, "%"+*value+"%")
	}
	if description != nil && *description != "" {
		db = db.Where("upper(?) LIKE upper(?)", clause.Column{Name: "description"}, "%"+*description+"%")
	}
	db = db.Order(database.TableOrderBy(orderName, orderMethod, "key", false))
	return database.Pagination[po.Config](db, pageNum, pageSize)
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

func (s *ConfigService) QueryConfigMap() map[string]*string {
	var configs []po.Config
	s.db.Model(&po.Config{}).
		Find(&configs)
	configMap := make(map[string]*string)
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

func (s *ConfigService) AddByAdmin(key string, value *string, description string) bool {
	db := s.db.Model(&po.Config{}).
		Create(&po.Config{
			Key:         key,
			Value:       value,
			Description: description,
			CreateTime:  time.Now(),
		})
	if db.RowsAffected > 0 {
		refreshBlogCacheInfo()
	}
	return db.RowsAffected > 0
}

func (s *ConfigService) UpdateByAdmin(key string, description string, value *string) bool {
	updateTime := time.Now()
	db := s.db.Model(&po.Config{}).
		Where("? = ?", clause.Column{Name: "key"}, key).
		Select("description", "value", "update_time").
		Updates(&po.Config{
			Description: description,
			Value:       value,
			UpdateTime:  &updateTime,
		})
	if db.RowsAffected > 0 {
		refreshBlogCacheInfo()
	}
	return db.RowsAffected > 0
}

func (s *ConfigService) DeleteByAdmin(key string) bool {
	db := s.db.Model(&po.Config{}).
		Where("? = ?", clause.Column{Name: "key"}, key).
		Delete(nil)
	if db.RowsAffected > 0 {
		refreshBlogCacheInfo()
	}
	return db.RowsAffected > 0
}
