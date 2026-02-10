package next

import (
	nast "blog/markdown/next/ast"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type tocParser struct {
}

var defaultTocParser = &tocParser{}

func NewTocParser() parser.BlockParser {
	return defaultTocParser
}

func (p *tocParser) Trigger() []byte {
	return []byte{'['}
}

func (p *tocParser) Open(parent ast.Node, reader text.Reader, pc parser.Context) (ast.Node, parser.State) {
	bytes, _ := reader.PeekLine()
	if bytes == nil || len(bytes) < 5 {
		return nil, parser.None
	}

	line := string(bytes)
	line = strings.Trim(line, " \n\r\t")
	line = strings.ToLower(line)
	line = strings.ReplaceAll(line, " ", "")
	if line != "[toc]" {
		return nil, parser.None
	}

	return nast.NewToc(&pc), parser.Close
}

func (p *tocParser) Continue(node ast.Node, reader text.Reader, pc parser.Context) parser.State {
	return parser.Close
}

func (p *tocParser) Close(node ast.Node, reader text.Reader, pc parser.Context) {
	// nothing to do
}

func (p *tocParser) CanInterruptParagraph() bool {
	return false
}

func (p *tocParser) CanAcceptIndentedLine() bool {
	return true
}

type tocRender struct {
}

var defaultTocRender = &tocRender{}

func NewTocRenderer() renderer.NodeRenderer {
	return defaultTocRender
}

func (r *tocRender) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(nast.KindToc, r.renderToc)
}

func (r *tocRender) renderToc(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		node := node.(*nast.Toc)
		catalogues := GetCatalogues(*node.Context())
		if catalogues == nil || len(catalogues) == 0 {
			return ast.WalkContinue, nil
		}
		r.renderTocCatalogues(w, catalogues)
	}
	return ast.WalkContinue, nil
}

func (r *tocRender) renderTocCatalogues(w util.BufWriter, catalogues []*Catalogue) {
	_, _ = w.WriteString(`<ol>`)
	for _, catalogue := range catalogues {
		_, _ = w.WriteString(`<li><strong>`)
		_, _ = w.WriteString(`<a href="#` + catalogue.Anchor + `">` + catalogue.Title + `</a>`)
		_, _ = w.WriteString(`</strong></li>`)
		if len(catalogue.Childs) == 0 {
			continue
		}
		r.renderTocCatalogues(w, catalogue.Childs)
	}
	_, _ = w.WriteString(`</ol>`)
}

type toc struct {
}

func NewToc() goldmark.Extender {
	return &toc{}
}

func (e *toc) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithBlockParsers(
			util.Prioritized(NewTocParser(), 999),
		),
	)
	m.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(NewTocRenderer(), 999),
		),
	)
}
