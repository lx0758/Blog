package next

import (
	nast "blog/markdown/next/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

var keyHasPreview = parser.NewContextKey()

type moreParser struct {
	preview bool
}

func NewMoreParser(preview bool) parser.BlockParser {
	return &moreParser{
		preview: preview,
	}
}

func (p *moreParser) Trigger() []byte {
	return []byte{'<'}
}

func (p *moreParser) Open(parent ast.Node, reader text.Reader, pc parser.Context) (ast.Node, parser.State) {
	bytes, _ := reader.PeekLine()
	if bytes == nil || len(bytes) < 10 {
		return nil, parser.None
	}

	line := string(bytes)
	line = strings.Trim(line, " \n\r\t")
	line = strings.ReplaceAll(line, " ", "")
	if line != "<!--more-->" {
		return nil, parser.None
	}

	if p.preview {
		pc.Set(keyHasPreview, true)
	}
	return nast.NewMore(), parser.Close
}

func (p *moreParser) Continue(node ast.Node, reader text.Reader, pc parser.Context) parser.State {
	if p.preview {
		// 如果是预览模式，则继续解析
		return parser.Continue
	} else {
		return parser.Close
	}
}

func (p *moreParser) Close(node ast.Node, reader text.Reader, pc parser.Context) {
	// nothing to do
}

func (p *moreParser) CanInterruptParagraph() bool {
	return false
}

func (p *moreParser) CanAcceptIndentedLine() bool {
	return true
}

type morePreview struct {
}

var MorePreview = &morePreview{}

func NewMorePreview() goldmark.Extender {
	return &morePreview{}
}

func (e *morePreview) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithBlockParsers(
			util.Prioritized(NewMoreParser(true), 399),
		),
	)
}

func HasPreview(context parser.Context) bool {
	value := context.Get(keyHasPreview)
	if result, ok := value.(bool); ok {
		return result
	}
	return false
}

type moreAnchorRenderer struct {
}

func NewMoreAnchorRenderer() renderer.NodeRenderer {
	return &moreAnchorRenderer{}
}

func (m *moreAnchorRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(nast.KindMore, m.renderMoreAnchor)
}

func (m *moreAnchorRenderer) renderMoreAnchor(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		_, _ = w.WriteString("<a")
		n.SetAttributeString("id", "more")
		n.SetAttributeString("href", "#more")
		if n.Attributes() != nil {
			html.RenderAttributes(w, n, nil)
		}
		_ = w.WriteByte('>')
	} else {
		_, _ = w.WriteString("</a>\n")
	}
	return ast.WalkContinue, nil
}

type moreAnchor struct {
}

var MoreAnchor = &moreAnchor{}

func NewMoreAnchor() goldmark.Extender {
	return &moreAnchor{}
}

func (e *moreAnchor) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithBlockParsers(
			util.Prioritized(NewMoreParser(false), 399),
		),
	)
	m.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(NewMoreAnchorRenderer(), 999),
		),
	)
}
