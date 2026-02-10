package markdown

import (
	"regexp"
	"strconv"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
)

type customIDs struct {
	idMap map[string]int
}

func (ids *customIDs) Generate(value []byte, kind ast.NodeKind) []byte {
	name := string(value)
	name = regexp.MustCompile("[^a-zA-Z0-9\u4E00-\u9FA5]+").ReplaceAllString(name, "-")
	ids.idMap[name]++
	index := ids.idMap[name]
	if index == 1 {
		return []byte(name)
	}
	return []byte(name + "-" + strconv.Itoa(index))
}

func (ids *customIDs) Put(value []byte) {
	// nothing to do
}

func newCustomIDs() parser.IDs {
	return &customIDs{
		idMap: map[string]int{},
	}
}

func NewCustomIDs() parser.ContextOption {
	return parser.WithIDs(newCustomIDs())
}
