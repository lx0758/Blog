package ast

import (
	gast "github.com/yuin/goldmark/ast"
)

type More struct {
	gast.BaseBlock
}

func (n *More) Dump(source []byte, level int) {
	m := map[string]string{}
	gast.DumpHelper(n, source, level, m, nil)
}

var KindMore = gast.NewNodeKind("More")

func (n *More) Kind() gast.NodeKind {
	return KindMore
}

func NewMore() *More {
	return &More{}
}
