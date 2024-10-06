package html_vo

import "blog/bean/po"

type CommentVO struct {
	Id       int         `json:"id"`
	Avatar   string      `json:"avatar"`
	Nickname string      `json:"nickname"`
	Browser  string      `json:"browser"`
	System   string      `json:"system"`
	Time     string      `json:"time"`
	Content  string      `json:"content"`
	Children []CommentVO `json:"children"`
}

func (v *CommentVO) From(comment po.Comment) {
	v.Id = comment.Id
	v.Avatar = "/blog/images/avatar.gif"
	v.Nickname = comment.Author
	v.Browser = comment.GetBrowser()
	v.System = comment.GetSystem()
	v.Time = comment.CreateTime.Format("2006-01-02 15:04:05")
	v.Content = comment.Content

	children := make([]CommentVO, 0)
	for _, child := range comment.Childs {
		childVO := CommentVO{}
		childVO.From(child)
		children = append(children, childVO)
	}
	v.Children = children
}

type CommentPaginationVO struct {
	Total    int         `json:"total"`
	HasMore  bool        `json:"hasMore"`
	Comments []CommentVO `json:"comments"`
}
