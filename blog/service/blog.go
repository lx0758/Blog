package service

import (
	"blog/bean/html_vo"
	"blog/bean/po"
	"blog/util"
)

type BlogService struct {
	Service
	configService   *ConfigService
	articleService  *ArticleService
	categoryService *CategoryService
	tagService      *TagService
	userService     *UserService
	linkService     *LinkService

	*html_vo.BlogVO
}

func (s *BlogService) OnInitService() {
	s.configService = GetService[*ConfigService](s.configService)
	s.articleService = GetService[*ArticleService](s.articleService)
	s.categoryService = GetService[*CategoryService](s.categoryService)
	s.tagService = GetService[*TagService](s.tagService)
	s.userService = GetService[*UserService](s.userService)
	s.linkService = GetService[*LinkService](s.linkService)

	s.refreshCacheInfo()
}

func (s *BlogService) GetCacheBlog() *html_vo.BlogVO {
	return s.BlogVO
}

func (s *BlogService) refreshCacheInfo() {
	configMap := s.configService.QueryConfigMap()
	articleCount := s.articleService.Count()
	categoryCount := s.categoryService.Count()
	tagCount := s.tagService.Count()

	user := s.userService.QueryOwner()
	if user == nil {
		user = &po.User{
			Accounts: &po.UserAccounts{},
		}
	}

	links := s.linkService.ListLink()
	linkVOs := make([]html_vo.LinkVO, 0)
	for _, link := range links {
		linkVOs = append(linkVOs, html_vo.LinkVO{
			Title: link.Title,
			Url:   link.Url,
		})
	}
	s.BlogVO = &html_vo.BlogVO{
		SiteDomain:         configMap[po.CONFIG_KEY_SITE_DOMAIN],
		SiteTitle:          configMap[po.CONFIG_KEY_SITE_TITLE],
		SiteDescription:    configMap[po.CONFIG_KEY_SITE_DESCRIPTION],
		SiteKeywords:       configMap[po.CONFIG_KEY_SITE_KEYWORDS],
		SiteBeianIcp:       configMap[po.CONFIG_KEY_SITE_BEIAN_ICP],
		SiteBeianPs:        configMap[po.CONFIG_KEY_SITE_BEIAN_PS],
		SiteBaidu:          configMap[po.CONFIG_KEY_SITE_BAIDU],
		SiteCreateYear:     configMap[po.CONFIG_KEY_SITE_CREATE_YEAR],
		SiteArticleCount:   articleCount,
		SiteCategoryCount:  categoryCount,
		SiteTagCount:       tagCount,
		OwnerAvatar:        util.IfNotNil(user.Avatar, ""),
		OwnerNickname:      user.Nickname,
		OwnerDescription:   util.IfNotNil(user.Description, ""),
		OwnerEmail:         util.IfNotNil(user.Email, ""),
		OwnerGithub:        util.IfNotNil(user.Accounts.Github, ""),
		OwnerWeibo:         util.IfNotNil(user.Accounts.Weibo, ""),
		OwnerGoogle:        util.IfNotNil(user.Accounts.Google, ""),
		OwnerTwitter:       util.IfNotNil(user.Accounts.Twitter, ""),
		OwnerFacebook:      util.IfNotNil(user.Accounts.Facebook, ""),
		OwnerStackOverflow: util.IfNotNil(user.Accounts.StackOverflow, ""),
		OwnerYoutube:       util.IfNotNil(user.Accounts.Youtube, ""),
		OwnerInstagram:     util.IfNotNil(user.Accounts.Instagram, ""),
		OwnerSkype:         util.IfNotNil(user.Accounts.Skype, ""),
		Links:              linkVOs,
	}
}
