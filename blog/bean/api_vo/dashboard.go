package api_vo

type Dashboard struct {
	ArticleCount  int
	CategoryCount int
	TagCount      int
	UploadCount   int
	CommentCount  int
	BrowseCount   int
	NewArticles   []Article
	NewComments   []Comment
	NewUploads    []File
}
