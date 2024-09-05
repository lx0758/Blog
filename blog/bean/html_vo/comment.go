package html_vo

type Comment struct {
	Id       int
	Avatar   string
	Nickname string
	Browser  string
	System   string
	Time     string
	Content  string
	Children []Comment
}
