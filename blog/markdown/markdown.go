package markdown

import (
	"bytes"
	"github.com/yuin/goldmark"
	"regexp"
	"strings"
)

type Catalogue struct {
	Title  string
	Anchor string
	Childs []Catalogue
}

func Render(markdown string, preview bool) (html string, text string, catalogues []Catalogue) {
	catalogues = []Catalogue{}
	var buffer bytes.Buffer
	if err := goldmark.Convert([]byte(markdown), &buffer); err != nil {
		panic(err)
	}
	html = buffer.String()
	text = TrimHtml(html)
	return
}

func TrimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "")
	return strings.TrimSpace(src)
}
