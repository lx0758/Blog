package service

import (
	"blog/bean/po"
	"blog/logger"
)

const (
	EVENT_REFRESH_CACHE_INGO = 1
)

var blogChannel = make(chan int, 1)

func refreshBlogCacheInfo() {
	blogChannel <- EVENT_REFRESH_CACHE_INGO
}

type BlogService struct {
	Service
	configService   *ConfigService
	articleService  *ArticleService
	categoryService *CategoryService
	tagService      *TagService
	fileService     *FileService
	commentService  *CommentService
	userService     *UserService
	linkService     *LinkService

	configMap     map[string]*string
	owner         po.User
	links         []po.Link
	articleCount  int
	categoryCount int
	tagCount      int
	fileCount     int
	commentCount  int
}

func (s *BlogService) OnInitService() {
	s.configService = GetService[*ConfigService](s.configService)
	s.articleService = GetService[*ArticleService](s.articleService)
	s.categoryService = GetService[*CategoryService](s.categoryService)
	s.tagService = GetService[*TagService](s.tagService)
	s.fileService = GetService[*FileService](s.fileService)
	s.commentService = GetService[*CommentService](s.commentService)
	s.userService = GetService[*UserService](s.userService)
	s.linkService = GetService[*LinkService](s.linkService)

	refreshBlogCacheInfo()
	go func() {
		for event := range blogChannel {
			logger.Printf("received a blog event event: %d", event)
			switch event {
			case EVENT_REFRESH_CACHE_INGO:
				s.refreshCacheInfo()
			}
		}
	}()
}

func (s *BlogService) GetCacheConfigMap() map[string]*string {
	return s.configMap
}

func (s *BlogService) GetCacheCount() (articleCount, categoryCount, tagCount, fileCount, commentCount int) {
	return s.articleCount, s.categoryCount, s.tagCount, s.fileCount, s.commentCount
}

func (s *BlogService) GetCacheOwner() po.User {
	return s.owner
}

func (s *BlogService) GetCacheLinks() []po.Link {
	return s.links
}

func (s *BlogService) refreshCacheInfo() {
	s.configMap = s.configService.QueryConfigMap()

	user := s.userService.QueryOwner()
	if user == nil {
		user = &po.User{
			Accounts: &po.UserAccounts{},
		}
	}
	s.owner = *user

	s.links = s.linkService.ListLink()

	s.articleCount = s.articleService.Count()
	s.categoryCount = s.categoryService.Count()
	s.tagCount = s.tagService.Count()
	s.fileCount = s.fileService.Count()
	s.commentCount = s.commentService.Count()
}
