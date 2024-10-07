package service

import (
	"blog/bean/po"
	"blog/database"
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

const ARTICLE_PAGE_SIZE = 10

type ArticleService struct {
	Service
	db *gorm.DB
}

func (s *ArticleService) OnInitService() {
	s.db = database.GetDB()
}

func (s *ArticleService) Count() int {
	var count int64
	s.db.Model(&po.Article{}).
		Where("? = ?", clause.Column{Name: "status"}, po.ARTICLE_STATUS_PUBLISHED).
		Count(&count)
	return int(count)
}

func (s *ArticleService) CountViews() int {
	var count int64
	s.db.Model(&po.Article{}).
		Select("SUM(?)", clause.Column{Name: "views"}).
		Where("? = ?", clause.Column{Name: "status"}, po.ARTICLE_STATUS_PUBLISHED).
		Take(&count)
	return int(count)
}

func (s *ArticleService) PaginationByAdmin(
	title *string,
	categoryId *int,
	url *string,
	comment *bool,
	status *int,
	pageNum int,
	pageSize int,
	orderName *string,
	orderMethod *string,
) po.Pagination[po.Article] {
	// 为了 Count 能够正确应用 DISTINCT 函数, 应该是 GORM 的一个 BUG
	// SELECT COUNT(DISTINCT("t_article"."id")) FROM "t_article" LEFT JOIN "t_article_url" ON "t_article"."id" = "t_article_url"."article_id"
	// SELECT DISTINCT t_article.* FROM "t_article" LEFT JOIN "t_article_url" ON "t_article"."id" = "t_article_url"."article_id" ORDER BY "id" DESC LIMIT 20
	db := s.db.Model(&po.Article{}).
		Select("t_article.id").
		Distinct("?", clause.Column{Name: "t_article.*", Raw: true}).
		Joins("LEFT JOIN ? ON ? = ?", clause.Table{Name: "t_article_url"}, clause.Column{Name: "t_article.id"}, clause.Column{Name: "t_article_url.article_id"}).
		Preload("Category").Preload("Urls")
	if title != nil {
		db = db.Where("upper(?) LIKE upper(?)", clause.Column{Name: "t_article.title"}, "%"+*title+"%")
	}
	if categoryId != nil {
		db = db.Where("? = ?", clause.Column{Name: "t_article.category_id"}, *categoryId)
	}
	if url != nil && *url != "" {
		db = db.Where("upper(?) LIKE upper(?)", clause.Column{Name: "t_article_url.url"}, "%"+*url+"%")
	}
	if comment != nil {
		var commentStatus int
		if *comment {
			commentStatus = po.ARTICLE_COMMENT_ENABLE
		} else {
			commentStatus = po.ARTICLE_COMMENT_DISABLE
		}
		db = db.Where("? = ?", clause.Column{Name: "t_article.comment_status"}, commentStatus)
	}
	if status != nil {
		db = db.Where("? = ?", clause.Column{Name: "t_article.status"}, *status)
	}
	db = db.Order(database.TableOrderBy(orderName, orderMethod, "id", true))
	return database.Pagination[po.Article](db, pageNum, pageSize)
}

func (s *ArticleService) PaginationArticleByPage(pageNum int) po.Pagination[po.Article] {
	db := s.db.Model(&po.Article{}).
		Preload("Category").Preload("Author").Preload("Urls").
		Where("? = ?", clause.Column{Name: "status"}, po.ARTICLE_STATUS_PUBLISHED).
		Order(clause.OrderBy{
			Columns: []clause.OrderByColumn{
				{Column: clause.Column{Name: "weight"}, Desc: true},
				{Column: clause.Column{Name: "id"}, Desc: true},
			},
		})
	return database.Pagination[po.Article](db, pageNum, ARTICLE_PAGE_SIZE)
}

func (s *ArticleService) PaginationArticleByArchive(pageNum int) po.Pagination[po.Article] {
	db := s.db.Model(&po.Article{}).
		Preload("Urls").
		Where("? = ?", clause.Column{Name: "status"}, po.ARTICLE_STATUS_PUBLISHED).
		Order(clause.OrderBy{
			Columns: []clause.OrderByColumn{
				{Column: clause.Column{Name: "create_time"}, Desc: true},
				{Column: clause.Column{Name: "id"}, Desc: true},
			},
		})
	return database.Pagination[po.Article](db, pageNum, ARTICLE_PAGE_SIZE)
}

func (s *ArticleService) PaginationArticleByCategory(categoryId int, pageNum int) po.Pagination[po.Article] {
	db := s.db.Model(&po.Article{}).
		Preload("Urls").
		InnerJoins("Category", s.db.Where(&po.Category{Id: categoryId})).
		Where(&po.Article{Status: po.ARTICLE_STATUS_PUBLISHED}).
		Order(clause.OrderBy{
			Columns: []clause.OrderByColumn{
				{Column: clause.Column{Name: "weight"}, Desc: true},
				{Column: clause.Column{Name: "id"}, Desc: true},
			},
		})
	return database.Pagination[po.Article](db, pageNum, ARTICLE_PAGE_SIZE)
}

func (s *ArticleService) PaginationArticleByTag(tagId int, pageNum int) po.Pagination[po.Article] {
	db := s.db.Model(&po.Article{}).
		Joins("LEFT JOIN ? ON ? = ?", clause.Table{Name: "t_article_tag"}, clause.Column{Name: "t_article.id"}, clause.Column{Name: "t_article_tag.article_id"}).
		Where("? = ? AND ? = ?", clause.Column{Name: "t_article.status"}, po.ARTICLE_STATUS_PUBLISHED, clause.Column{Name: "t_article_tag.tag_id"}, tagId).
		Order(clause.OrderBy{
			Columns: []clause.OrderByColumn{
				{Column: clause.Column{Name: "weight"}, Desc: true},
				{Column: clause.Column{Name: "id"}, Desc: true},
			},
		})
	return database.Pagination[po.Article](db, pageNum, ARTICLE_PAGE_SIZE)
}

func (s *ArticleService) ListNewArticle(count int) []po.Article {
	var articles []po.Article
	s.db.Model(&po.Article{}).
		Preload("Urls").
		Where("? = ?", clause.Column{Name: "status"}, po.ARTICLE_STATUS_PUBLISHED).
		Order(clause.OrderByColumn{
			Column: clause.Column{Name: "id"},
			Desc:   true,
		}).
		Limit(count).
		Find(&articles)
	return articles
}

func (s *ArticleService) ListArticle() []po.Article {
	var articles []po.Article
	s.db.Model(&po.Article{}).
		Preload("Urls").
		Where("? = ?", clause.Column{Name: "status"}, po.ARTICLE_STATUS_PUBLISHED).
		Order(clause.OrderByColumn{
			Column: clause.Column{Name: "id"},
			Desc:   true,
		}).
		Find(&articles)
	return articles
}

func (s *ArticleService) QueryArticle(id int) *po.Article {
	var article po.Article
	s.db.Model(&po.Article{}).
		Preload("Category").Preload("Author").Preload("Urls").Preload("Tags").
		Where("? = ? AND ? = ?", clause.Column{Name: "status"}, po.ARTICLE_STATUS_PUBLISHED, clause.Column{Name: "id"}, id).
		Take(&article)
	if article.Id == 0 {
		return nil
	}
	s.incrementViews(&article)
	return &article
}

func (s *ArticleService) QueryByAdmin(id int) *po.Article {
	var article po.Article
	s.db.Model(&po.Article{}).
		Preload("Category").Preload("Author").Preload("Urls").Preload("Tags").
		Where("? = ?", clause.Column{Name: "id"}, id).
		Take(&article)
	if article.Id == 0 {
		return nil
	}
	return &article
}

func (s *ArticleService) QueryPrevArticle(id int) *po.Article {
	var article po.Article
	s.db.Model(&po.Article{}).
		Preload("Urls").
		Where("? = ? AND ? < ?", clause.Column{Name: "status"}, po.ARTICLE_STATUS_PUBLISHED, clause.Column{Name: "id"}, id).
		Order(clause.OrderBy{
			Columns: []clause.OrderByColumn{
				{Column: clause.Column{Name: "weight"}, Desc: true},
				{Column: clause.Column{Name: "id"}, Desc: true},
			},
		}).
		Limit(1).
		Take(&article)
	if article.Id == 0 {
		return nil
	}
	return &article
}

func (s *ArticleService) QueryNextArticle(id int) *po.Article {
	var article po.Article
	s.db.Model(&po.Article{}).
		Preload("Urls").
		Where("? = ? AND ? > ?", clause.Column{Name: "status"}, po.ARTICLE_STATUS_PUBLISHED, clause.Column{Name: "id"}, id).
		Order(clause.OrderBy{
			Columns: []clause.OrderByColumn{
				{Column: clause.Column{Name: "weight"}, Desc: false},
				{Column: clause.Column{Name: "id"}, Desc: false},
			},
		}).
		Limit(1).
		Take(&article)
	if article.Id == 0 {
		return nil
	}
	return &article
}

func (s *ArticleService) AddByAdmin(
	userId int,
	title string,
	categoryId int,
	content string,
	url *string,
	weight int,
	commentStatus int,
	status int,
	tags *[]string,
) (*po.Article, error) {
	var db *gorm.DB
	tdb := s.db.Begin()

	article := po.Article{
		Title:         title,
		Content:       content,
		CategoryId:    categoryId,
		AuthorId:      userId,
		Weight:        weight,
		CommentStatus: commentStatus,
		Status:        status,
		CreateTime:    time.Now(),
	}
	tdb.Model(&article).
		Create(&article)
	if article.Id == 0 {
		tdb.Rollback()
		return nil, errors.New("添加文章失败")
	}

	if url != nil && *url != "" {
		db = tdb.Model(&po.ArticleUrl{}).
			Create(&po.ArticleUrl{
				Url:        *url,
				ArticleId:  article.Id,
				Status:     po.ARTICLE_URL_STATUS_LAST,
				CreateTime: time.Now(),
			})
		if db.RowsAffected == 0 {
			tdb.Rollback()
			return nil, errors.New("添加文章链接失败")
		}
	}

	if tags != nil {
		for _, tag := range *tags {
			tagPO := po.Tag{}
			tdb.Model(&po.Tag{}).
				Where("? = ?", clause.Column{Name: "name"}, tag).
				Take(&tag)
			if tagPO.Id == 0 {
				tagPO = po.Tag{
					Name:       tag,
					CreateTime: time.Now(),
				}
				db = tdb.Model(&tagPO).
					Create(&tagPO)
				if tagPO.Id == 0 {
					tdb.Rollback()
					return nil, errors.New("添加标签失败")
				}
			}
			db = tdb.Table("t_article_tag").
				Create(map[string]interface{}{
					"article_id": article.Id,
					"tag_id":     tagPO.Id,
				})
			if db.RowsAffected == 0 {
				tdb.Rollback()
				return nil, errors.New("关联文章标签失败")
			}
		}
	}

	tdb.Commit()
	refreshBlogCacheInfo()
	return &article, nil
}

func (s *ArticleService) UpdateByAdmin(
	id int,
	title string,
	categoryId int,
	content string,
	url *string,
	weight int,
	commentStatus int,
	status int,
	tags *[]string,
) error {
	var db *gorm.DB
	tdb := s.db.Begin()

	updateTime := time.Now()
	article := po.Article{
		Title:         title,
		Content:       content,
		CategoryId:    categoryId,
		Weight:        weight,
		CommentStatus: commentStatus,
		Status:        status,
		UpdateTime:    &updateTime,
	}
	db = tdb.Model(&article).
		Where("? = ?", clause.Column{Name: "id"}, id).
		Select("title", "content", "category_id", "weight", "comment_status", "status", "update_time").
		Updates(&article)
	if db.RowsAffected == 0 {
		tdb.Rollback()
		return errors.New("更新文章失败")
	}

	if url != nil && *url != "" {
		articleUrl := po.ArticleUrl{}
		db = tdb.Model(&articleUrl).
			Where("? = ? AND ? = ?", clause.Column{Name: "url"}, url, clause.Column{Name: "article_id"}, id).
			First(&articleUrl)
		if articleUrl.Url == "" {
			articleUrl = po.ArticleUrl{
				Url:        *url,
				ArticleId:  id,
				Status:     po.ARTICLE_URL_STATUS_LAST,
				CreateTime: time.Now(),
			}
			db = tdb.Model(&articleUrl).
				Create(&articleUrl)
			if db.RowsAffected == 0 {
				tdb.Rollback()
				return errors.New("添加文章链接失败")
			}
		}
		if articleUrl.Status != po.ARTICLE_URL_STATUS_LAST {
			tdb.Model(&articleUrl).
				Where("? = ?", clause.Column{Name: "url"}, url).
				Update("status", po.ARTICLE_URL_STATUS_LAST)
			if db.RowsAffected == 0 {
				tdb.Rollback()
				return errors.New("刷新文章链接失败")
			}
		}
		tdb.Model(&po.ArticleUrl{}).
			Where("? = ? AND ? <> ?", clause.Column{Name: "article_id"}, id, clause.Column{Name: "url"}, *url).
			Update("status", po.ARTICLE_URL_STATUS_OLD)
	} else {
		tdb.Model(&po.ArticleUrl{}).
			Where("? = ?", clause.Column{Name: "article_id"}, id).
			Update("status", po.ARTICLE_URL_STATUS_OLD)
	}

	tdb.Model(&article).
		Preload("Tags").
		Where("? = ?", clause.Column{Name: "id"}, id).
		Take(&article)
	articleTags := article.Tags
	articleNewTags := make([]string, 0)
	if tags != nil {
		for _, tag := range *tags {
			articleNewTags = append(articleNewTags, tag)
		}
	}
	needLinkTags := make([]string, 0)
	for _, tag := range articleNewTags {
		found := false
		for _, tagPO := range articleTags {
			if tag == tagPO.Name {
				found = true
				break
			}
		}
		if !found {
			needLinkTags = append(needLinkTags, tag)
		}
	}
	for _, tag := range needLinkTags {
		tagPO := po.Tag{}
		tdb.Model(&po.Tag{}).
			Where("? = ?", clause.Column{Name: "name"}, tag).
			Take(&tagPO)
		if tagPO.Id == 0 {
			tagPO = po.Tag{
				Name:       tag,
				CreateTime: time.Now(),
			}
			tdb.Model(&tagPO).
				Create(&tagPO)
			if tagPO.Id == 0 {
				tdb.Rollback()
				return errors.New("添加标签失败")
			}
		}
		db = tdb.Table("t_article_tag").
			Create(map[string]interface{}{
				"article_id": article.Id,
				"tag_id":     tagPO.Id,
			})
		if db.RowsAffected == 0 {
			tdb.Rollback()
			return errors.New("关联文章标签失败")
		}
	}
	needUnlinkTags := make([]po.Tag, 0)
	for _, tagPO := range articleTags {
		found := false
		for _, newTag := range articleNewTags {
			if tagPO.Name == newTag {
				found = true
				break
			}
		}
		if !found {
			needUnlinkTags = append(needUnlinkTags, tagPO)
		}
	}
	for _, tag := range needUnlinkTags {
		db = tdb.Table("t_article_tag").
			Where("? = ? AND ? = ?", clause.Column{Name: "article_id"}, id, clause.Column{Name: "tag_id"}, tag.Id).
			Delete(nil)
		if db.RowsAffected == 0 {
			tdb.Rollback()
			return errors.New("解除关联文章标签失败")
		}
	}

	tdb.Commit()
	refreshBlogCacheInfo()
	return nil
}

func (s *ArticleService) UpdateStatusByAdmin(id int, status int) bool {
	db := s.db.Model(&po.Article{}).
		Where("? = ?", clause.Column{Name: "id"}, id).
		Update("status", status)
	if db.RowsAffected == 0 {
		refreshBlogCacheInfo()
	}
	return db.RowsAffected > 0
}

func (s *ArticleService) DeleteByAdmin(id int) bool {
	db := s.db.Delete(&po.Article{Id: id})
	if db.RowsAffected == 0 {
		refreshBlogCacheInfo()
	}
	return db.RowsAffected > 0
}

func (s *ArticleService) incrementViews(article *po.Article) {
	if article == nil || article.Id == 0 {
		return
	}
	s.db.Model(&po.Article{Id: article.Id}).
		Update("views", gorm.Expr("views + 1"))
}
