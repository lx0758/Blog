package next

import (
	nast "blog/markdown/next/ast"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
	"strings"
)

type Catalogue struct {
	Title  string
	Anchor string
	Childs []*Catalogue

	parent   *Catalogue
	rawLevel int
}

type catalogueStore struct {
	root *Catalogue
	last *Catalogue
}

/**
 * 处理某些文章跳级使用标题标签, 导致无法正常解析标题列表的错误
 *
 * #     ........ 1
 * ###   ........ 2
 * ##### ........ 3
 * ##    ........ 2
 * ####  ........ 3
 * ###   ........ 3
 * #     ........ 1
 **/
func (store *catalogueStore) addCatalogue(title string, anchor string, level int) {
	if level == store.last.rawLevel {
		store.last = store.addCatalogueChild(store.last.parent, title, anchor, level)
	} else if level > store.last.rawLevel {
		store.last = store.addCatalogueChild(store.last, title, anchor, level)
	} else {
		current := store.last
		for level <= current.rawLevel {
			if current.parent == nil {
				break
			}
			current = current.parent
		}
		store.last = store.addCatalogueChild(current, title, anchor, level)
	}
}

func (store *catalogueStore) addCatalogueChild(parent *Catalogue, title string, anchor string, level int) *Catalogue {
	catalogue := &Catalogue{
		Title:    title,
		Anchor:   anchor,
		Childs:   make([]*Catalogue, 0),
		parent:   parent,
		rawLevel: level,
	}
	parent.Childs = append(parent.Childs, catalogue)
	return catalogue
}

func (store *catalogueStore) getCatalogues() []*Catalogue {
	return store.root.Childs
}

func newCatalogueStore() *catalogueStore {
	root := &Catalogue{
		Title:    "",
		Anchor:   "",
		Childs:   make([]*Catalogue, 0),
		parent:   nil,
		rawLevel: 0,
	}
	return &catalogueStore{
		root: root,
		last: root,
	}
}

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

var keyCatalogues = parser.NewContextKey()

type cataloguesASTTransformer struct {
}

var defaultCataloguesASTTransformer = &cataloguesASTTransformer{}

func NewCataloguesASTTransformer() parser.ASTTransformer {
	return defaultCataloguesASTTransformer
}

func (t *cataloguesASTTransformer) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	catalogueStore := newCatalogueStore()
	_ = ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}
		switch n := n.(type) {
		case *ast.Heading:
			title := string(n.Text(reader.Source()))
			anchor := ""
			if value, ok := n.AttributeString("id"); ok {
				if id, ok := value.([]byte); ok {
					anchor = string(id)
				}
			}
			level := n.Level
			catalogueStore.addCatalogue(title, anchor, level)
			return ast.WalkSkipChildren, nil
		default:
			return ast.WalkContinue, nil
		}
	})
	pc.Set(keyCatalogues, catalogueStore.getCatalogues())
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

type catalogues struct {
}

func NewCatalogues() goldmark.Extender {
	return &catalogues{}
}

func (e *catalogues) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithBlockParsers(
			util.Prioritized(NewTocParser(), 999),
		),
	)
	m.Parser().AddOptions(
		parser.WithASTTransformers(
			util.Prioritized(NewCataloguesASTTransformer(), 999),
		),
	)
	m.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(NewTocRenderer(), 999),
		),
	)
}

func GetCatalogues(context parser.Context) []*Catalogue {
	value := context.Get(keyCatalogues)
	if catalogues, ok := value.([]*Catalogue); ok {
		return catalogues
	}
	return nil
}
