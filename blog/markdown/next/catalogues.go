package next

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
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

type catalogues struct {
}

func NewCatalogues() goldmark.Extender {
	return &catalogues{}
}

func (e *catalogues) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithASTTransformers(
			util.Prioritized(NewCataloguesASTTransformer(), 999),
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
