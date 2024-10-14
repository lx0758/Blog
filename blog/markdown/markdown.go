package markdown

import (
	"blog/markdown/next"
	"bytes"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

func RenderByArticle(source string, fs FragmentSource) (html string, description string, catalogues []*next.Catalogue) {
	sourceBytes := getFinalSource(source, fs)
	markdown := getMarkdown(
		next.NewCatalogues(),
		next.NewMoreAnchor(),
	)
	context := parser.NewContext(NewCustomIDs())
	var buffer bytes.Buffer
	if err := markdown.Convert(sourceBytes, &buffer, parser.WithContext(context)); err != nil {
		panic(err)
	}
	html = buffer.String()
	description = trimHtml(html)
	catalogues = next.GetCatalogues(context)
	return
}

func RenderByPreview(source string, fs FragmentSource) (html string, hasPreview bool) {
	sourceBytes := getFinalSource(source, fs)
	markdown := getMarkdown(
		next.NewMorePreview(),
	)
	var buffer bytes.Buffer
	context := parser.NewContext(NewCustomIDs())
	if err := markdown.Convert(sourceBytes, &buffer, parser.WithContext(context)); err != nil {
		panic(err)
	}
	html = buffer.String()
	hasPreview = next.HasPreview(context)
	return
}

func RenderBySearch(source string, fs FragmentSource) (text string) {
	sourceBytes := getFinalSource(source, fs)
	markdown := getMarkdown()
	var buffer bytes.Buffer
	if err := markdown.Convert(sourceBytes, &buffer); err != nil {
		panic(err)
	}
	html := buffer.String()
	text = trimHtml(html)
	return
}

func getMarkdown(ext ...goldmark.Extender) goldmark.Markdown {
	extends := []goldmark.Extender{
		//extension.GFM,            // GitHub Flavored Markdown(Linkify, Table, Strikethrough, TaskList)
		extension.NewLinkify(),     // 自动链接
		extension.NewTable(),       // 表格
		extension.Strikethrough,    // 删除线
		extension.TaskList,         // 任务列表
		extension.DefinitionList,   // 定义列表
		extension.NewCJK(),         // 换行优化
		extension.NewFootnote(),    // 脚注
		extension.NewTypographer(), // 排版优化

		next.NewCustom(),
		next.NewUnderline(),
		next.NewHighlight(),
		next.NewToc(),
	}
	extends = append(extends, ext...)
	return goldmark.New(
		goldmark.WithParserOptions(
			parser.WithEscapedSpace(),
			parser.WithAutoHeadingID(),
			parser.WithHeadingAttribute(),
		),
		goldmark.WithRendererOptions(),
		goldmark.WithExtensions(
			extends...,
		),
	)
}
