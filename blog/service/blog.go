package service

import "blog/bean/html_vo"

type BlogService struct {
	Service
	configService   *ConfigService
	articleService  *ArticleService
	categoryService *CategoryService
	tagService      *TagService
	linkService     *LinkService

	*html_vo.BlogVO
}

func (s *BlogService) OnInitService() {
	s.configService = GetService[*ConfigService](s.configService)
	s.articleService = GetService[*ArticleService](s.articleService)
	s.categoryService = GetService[*CategoryService](s.categoryService)
	s.tagService = GetService[*TagService](s.tagService)
	s.linkService = GetService[*LinkService](s.linkService)

	s.refreshCacheInfo()
}

func (s *BlogService) GetCacheBlog() *html_vo.BlogVO {
	return s.BlogVO
}

func (s *BlogService) refreshCacheInfo() {
	s.BlogVO = &html_vo.BlogVO{
		SiteDomain:         "6xyun.cn",
		SiteTitle:          "6x's Website",
		SiteDescription:    "我的网站",
		SiteKeywords:       "6x,Liux",
		SiteBeianIcp:       "浙B2-20080101-4",
		SiteBeianPs:        "浙公网安备33010602009975号",
		SiteBaidu:          "123",
		SiteCreateYear:     "2014",
		SiteArticleCount:   1001,
		SiteCategoryCount:  20,
		SiteTagCount:       63,
		OwnerAvatar:        "",
		OwnerNickname:      "6x",
		OwnerDescription:   "一个程序员",
		OwnerEmail:         "lx0758@qq.com",
		OwnerGithub:        "lx0758",
		OwnerWeibo:         "lx0758",
		OwnerGoogle:        "lx0758",
		OwnerTwitter:       "lx0758",
		OwnerFacebook:      "lx0758",
		OwnerStackOverflow: "lx0758",
		OwnerYoutube:       "lx0758",
		OwnerInstagram:     "lx0758",
		OwnerSkype:         "lx0758",
		Links: []html_vo.LinkVO{{
			Title: "百度",
			Url:   "https://www.baidu.com",
		}, {
			Title: "Google",
			Url:   "https://www.google.com",
		}},
	}
}
