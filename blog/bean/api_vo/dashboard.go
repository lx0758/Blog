package api_vo

type DashboardVO struct {
	ArticleCount  int
	CategoryCount int
	TagCount      int
	UploadCount   int
	CommentCount  int
	BrowseCount   int
	NewArticles   []ArticleVO
	NewComments   []CommentVO
	NewUploads    []FileVO
}
