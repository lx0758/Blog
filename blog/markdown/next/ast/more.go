package ast

import (
	gast "github.com/yuin/goldmark/ast"
	"strconv"
)

type More struct {
	gast.BaseBlock
	preview bool
}

func (n *More) Preview() bool {
	return n.preview
}

func (n *More) Dump(source []byte, level int) {
	m := map[string]string{
		"preview": strconv.FormatBool(n.preview),
	}
	gast.DumpHelper(n, source, level, m, nil)
}

var KindMore = gast.NewNodeKind("More")

func (n *More) Kind() gast.NodeKind {
	return KindMore
}

func NewMore(preview bool) *More {
	return &More{
		preview: preview,
	}
}
