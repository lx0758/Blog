package html_vo

type CommentPagination struct {
	Total    int
	HasMore  bool
	Comments []Comment
}
