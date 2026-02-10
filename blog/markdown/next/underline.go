package next

import (
	nast "blog/markdown/next/ast"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type underlineDelimiterProcessor struct {
}

var defaultUnderlineDelimiterProcessor = &underlineDelimiterProcessor{}

func NewUnderlineDelimiterProcessor() parser.DelimiterProcessor {
	return defaultUnderlineDelimiterProcessor
}

func (p *underlineDelimiterProcessor) IsDelimiter(b byte) bool {
	return b == '+'
}

func (p *underlineDelimiterProcessor) CanOpenCloser(opener, closer *parser.Delimiter) bool {
	return opener.Char == closer.Char
}

func (p *underlineDelimiterProcessor) OnMatch(consumes int) ast.Node {
	return nast.NewUnderline()
}

type underlineParser struct{}

var defaultUnderlineParser = &underlineParser{}

func NewUnderlineParser() parser.InlineParser {
	return defaultUnderlineParser
}

func (p *underlineParser) Trigger() []byte {
	return []byte{'+'}
}

func (p *underlineParser) Parse(parent ast.Node, block text.Reader, pc parser.Context) ast.Node {
	before := block.PrecendingCharacter()
	line, segment := block.PeekLine()
	node := parser.ScanDelimiter(line, before, 1, NewUnderlineDelimiterProcessor())
	if node == nil || node.OriginalLength > 2 || before == '+' {
		return nil
	}

	node.Segment = segment.WithStop(segment.Start + node.OriginalLength)
	block.Advance(node.OriginalLength)
	pc.PushDelimiter(node)
	return node
}

type underlineRenderer struct{}

var defaultUnderlineRenderer = &underlineRenderer{}

func NewUnderlineRenderer() renderer.NodeRenderer {
	return defaultUnderlineRenderer
}

func (r *underlineRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(nast.KindUnderline, r.renderUnderline)
}

var UnderlineAttributeFilter = html.GlobalAttributeFilter.Extend(
	[]byte("u"),
)

func (r *underlineRenderer) renderUnderline(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		if n.Attributes() != nil {
			_, _ = w.WriteString("<u")
			html.RenderAttributes(w, n, UnderlineAttributeFilter)
			_ = w.WriteByte('>')
		} else {
			_, _ = w.WriteString("<u>")
		}
	} else {
		_, _ = w.WriteString("</u>")
	}
	return ast.WalkContinue, nil
}

type underline struct {
}

func NewUnderline() goldmark.Extender {
	return &underline{}
}

func (l *underline) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithInlineParsers(
			util.Prioritized(NewUnderlineParser(), 999),
		),
	)
	m.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(NewUnderlineRenderer(), 999),
		),
	)
}
