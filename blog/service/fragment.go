package service

import (
	"blog/bean/po"
	"blog/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type FragmentService struct {
	Service
	db *gorm.DB
}

func (s *FragmentService) OnInitService() {
	s.db = database.GetDB()
}

func (s *FragmentService) GetFragment(key string) string {
	var fragment po.Fragment
	s.db.Model(&po.Fragment{}).
		Where("? = ?", clause.Column{Name: "key"}, key).
		First(&fragment)
	return fragment.Content
}

func (s *FragmentService) PaginationByAdmin(
	key *string,
	context *string,
	status *int,
	pageNum int,
	pageSize int,
	orderName *string,
	orderMethod *string,
) po.Pagination[po.Fragment] {
	db := s.db.Model(&po.Fragment{}).
		Preload("Author")
	if key != nil && *key != "" {
		db = db.Where("upper(?) LIKE upper(?)", clause.Column{Name: "key"}, "%"+*key+"%")
	}
	if context != nil && *context != "" {
		db = db.Where("upper(?) LIKE upper(?)", clause.Column{Name: "context"}, "%"+*context+"%")
	}
	if status != nil {
		db = db.Where("? = ?", clause.Column{Name: "status"}, *status)
	}
	db = db.Order(database.TableOrderBy(orderName, orderMethod, "id", true))
	return database.Pagination[po.Fragment](db, pageNum, pageSize)
}

func (s *FragmentService) QueryFragmentById(id int) *po.Fragment {
	var fragment po.Fragment
	s.db.Model(&po.Fragment{}).
		Where("? = ?", clause.Column{Name: "id"}, id).
		First(&fragment)
	if fragment.Id == 0 {
		return nil
	}
	return &fragment
}

func (s *FragmentService) QueryFragmentByKey(key string) *po.Fragment {
	var fragment po.Fragment
	s.db.Model(&po.Fragment{}).
		Where("? = ?", clause.Column{Name: "key"}, key).
		First(&fragment)
	if fragment.Id == 0 {
		return nil
	}
	return &fragment
}

func (s *FragmentService) AddByAdmin(userId int, key string, content string, status int) *po.Fragment {
	fragment := po.Fragment{
		Key:        key,
		Content:    content,
		AuthorId:   userId,
		Status:     status,
		CreateTime: time.Now(),
	}
	s.db.Model(&po.Fragment{}).
		Create(&fragment)
	if fragment.Id == 0 {
		return nil
	}
	return &fragment
}

func (s *FragmentService) UpdateByAdmin(id int, key string, content string, status int) bool {
	updateTime := time.Now()
	db := s.db.Model(&po.Fragment{}).
		Where("id = ?", id).
		Select("content", "key", "status").
		Updates(&po.Fragment{
			Key:        key,
			Content:    content,
			Status:     status,
			UpdateTime: &updateTime,
		})
	return db.RowsAffected > 0
}

func (s *FragmentService) DeleteByAdmin(id int) bool {
	db := s.db.Model(&po.Fragment{}).
		Where("? = ?", clause.Column{Name: "id"}, id).
		Delete(nil)
	return db.RowsAffected > 0
}
