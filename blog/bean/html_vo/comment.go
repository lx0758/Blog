package html_vo

type CommentVO struct {
	Id       int
	Avatar   string
	Nickname string
	Browser  string
	System   string
	Time     string
	Content  string
	Children []CommentVO
}

type CommentPagination struct {
	Total    int
	HasMore  bool
	Comments []CommentVO
}
