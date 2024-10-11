package ast

import (
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
)

type Toc struct {
	gast.BaseBlock
	context *parser.Context
}

func (n *Toc) Context() *parser.Context {
	return n.context
}

func (n *Toc) Dump(source []byte, level int) {
	m := map[string]string{}
	gast.DumpHelper(n, source, level, m, nil)
}

var KindToc = gast.NewNodeKind("Toc")

func (n *Toc) Kind() gast.NodeKind {
	return KindToc
}

func NewToc(context *parser.Context) *Toc {
	return &Toc{
		context: context,
	}
}
