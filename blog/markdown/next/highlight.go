package next

import (
	nast "blog/markdown/next/ast"
	"github.com/yuin/goldmark/renderer/html"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type highlightDelimiterProcessor struct {
}

var defaultHighlightDelimiterProcessor = &highlightDelimiterProcessor{}

func NewHighlightDelimiterProcessor() parser.DelimiterProcessor {
	return defaultHighlightDelimiterProcessor
}

func (p *highlightDelimiterProcessor) IsDelimiter(b byte) bool {
	return b == '='
}

func (p *highlightDelimiterProcessor) CanOpenCloser(opener, closer *parser.Delimiter) bool {
	return opener.Char == closer.Char
}

func (p *highlightDelimiterProcessor) OnMatch(consumes int) ast.Node {
	return nast.NewHighlight()
}

type highlightParser struct{}

var defaultHighlightParser = &highlightParser{}

func NewHighlightParser() parser.InlineParser {
	return defaultHighlightParser
}

func (p *highlightParser) Trigger() []byte {
	return []byte{'='}
}

func (p *highlightParser) Parse(parent ast.Node, block text.Reader, pc parser.Context) ast.Node {
	before := block.PrecendingCharacter()
	line, segment := block.PeekLine()
	node := parser.ScanDelimiter(line, before, 1, NewHighlightDelimiterProcessor())
	if node == nil || node.OriginalLength > 2 || before == '=' {
		return nil
	}

	node.Segment = segment.WithStop(segment.Start + node.OriginalLength)
	block.Advance(node.OriginalLength)
	pc.PushDelimiter(node)
	return node
}

type highlightRenderer struct{}

var defaultHighlightRenderer = &highlightRenderer{}

func NewHighlightRenderer() renderer.NodeRenderer {
	return defaultHighlightRenderer
}

func (r *highlightRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(nast.KindHighlight, r.renderHighlight)
}

var HighlightAttributeFilter = html.GlobalAttributeFilter.Extend(
	[]byte("mark"),
)

func (r *highlightRenderer) renderHighlight(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		if n.Attributes() != nil {
			_, _ = w.WriteString("<mark")
			html.RenderAttributes(w, n, HighlightAttributeFilter)
			_ = w.WriteByte('>')
		} else {
			_, _ = w.WriteString("<mark>")
		}
	} else {
		_, _ = w.WriteString("</mark>")
	}
	return ast.WalkContinue, nil
}

type highlight struct {
}

var Highlight = &highlight{}

func NewHighlight() goldmark.Extender {
	return &highlight{}
}

func (l *highlight) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithInlineParsers(
			util.Prioritized(NewHighlightParser(), 999),
		),
	)
	m.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(NewHighlightRenderer(), 999),
		),
	)
}
