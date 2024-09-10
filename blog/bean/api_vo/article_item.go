package api_vo

import (
	"blog/bean/po"
	"sort"
	"time"
)

type ArticleItemVO struct {
	Id            int        `json:"id"`
	Title         string     `json:"title"`
	CategoryName  string     `json:"categoryName"`
	Url           string     `json:"url"`
	Urls          []string   `json:"urls"`
	Weight        int        `json:"weight"`
	Views         int        `json:"views"`
	EnableComment bool       `json:"enableComment"`
	Status        int        `json:"status"`
	CreateTime    time.Time  `json:"createTime"`
	UpdateTime    *time.Time `json:"updateTime"`
}

func (a *ArticleItemVO) From(article po.Article) {
	oldUrls := make([]string, 0)
	sort.Slice(article.Urls, func(i, j int) bool {
		return article.Urls[i].CreateTime.Unix() > article.Urls[j].CreateTime.Unix()
	})
	for _, articleUrl := range article.Urls {
		if articleUrl.Status == po.ARTICLE_URL_STATUS_OLD {
			oldUrls = append(oldUrls, articleUrl.Url)
		}
	}
	a.Id = article.Id
	a.Title = article.Title
	a.CategoryName = article.Category.Name
	a.Url = article.GetSafeUrl()
	a.Urls = oldUrls
	a.Weight = article.Weight
	a.Views = article.Views
	a.EnableComment = article.CommentStatus == po.ARTICLE_COMMENT_ENABLE
	a.Status = article.Status
	a.CreateTime = article.CreateTime
	a.UpdateTime = article.UpdateTime
}
