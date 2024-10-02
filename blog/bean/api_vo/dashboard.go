package api_vo

type DashboardVO struct {
	ArticleCount  int         `json:"articleCount"`
	CategoryCount int         `json:"categoryCount"`
	TagCount      int         `json:"tagCount"`
	UploadCount   int         `json:"uploadCount"`
	CommentCount  int         `json:"commentCount"`
	BrowseCount   int         `json:"browseCount"`
	NewArticles   []ArticleVO `json:"newArticles"`
	NewComments   []CommentVO `json:"newComments"`
	NewUploads    []FileVO    `json:"newUploads"`
}
