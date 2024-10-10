package ast

import (
	gast "github.com/yuin/goldmark/ast"
)

type Underline struct {
	gast.BaseInline
}

func (n *Underline) Dump(source []byte, level int) {
	m := map[string]string{}
	gast.DumpHelper(n, source, level, m, nil)
}

var KindUnderline = gast.NewNodeKind("Underline")

func (n *Underline) Kind() gast.NodeKind {
	return KindUnderline
}

func NewUnderline() *Underline {
	return &Underline{}
}
