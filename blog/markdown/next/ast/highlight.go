package ast

import (
	gast "github.com/yuin/goldmark/ast"
)

type Highlight struct {
	gast.BaseInline
}

func (n *Highlight) Dump(source []byte, level int) {
	m := map[string]string{}
	gast.DumpHelper(n, source, level, m, nil)
}

var KindHighlight = gast.NewNodeKind("Highlight")

func (n *Highlight) Kind() gast.NodeKind {
	return KindHighlight
}

func NewHighlight() *Highlight {
	return &Highlight{}
}
