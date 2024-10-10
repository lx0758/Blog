package next

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	east "github.com/yuin/goldmark/extension/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

func init() {
	// Extend
	html.HeadingAttributeFilter = html.HeadingAttributeFilter.Extend(
		[]byte("href"),
	)
	// BUG(?)
	extension.StrikethroughAttributeFilter = extension.StrikethroughAttributeFilter.Extend(
		[]byte("del"),
	)
}

type customASTTransformer struct {
}

var defaultCustomASTTransformer = &customASTTransformer{}

func NewCustomASTTransformer() parser.ASTTransformer {
	return defaultCustomASTTransformer
}

func (t *customASTTransformer) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	_ = ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}
		switch n := n.(type) {
		case *ast.Heading:
			if value, ok := n.AttributeString("id"); ok {
				if id, ok := value.([]byte); ok {
					n.SetAttributeString("href", "#"+string(id))
				}
			}
		case *ast.Link:
		case *ast.AutoLink:
			n.SetAttributeString("target", "_blank")
		}
		return ast.WalkContinue, nil
	})
}

type customStyleRenderer struct {
	html.Config
}

func NewCustomStyleRenderer(opts ...html.Option) renderer.NodeRenderer {
	r := &customStyleRenderer{
		Config: html.NewConfig(),
	}
	for _, opt := range opts {
		opt.SetHTMLOption(&r.Config)
	}
	return r
}

func (r *customStyleRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindFencedCodeBlock, r.renderFencedCodeBlock)
	reg.Register(east.KindTaskCheckBox, r.renderTaskCheckBox)
}

func (r *customStyleRenderer) renderFencedCodeBlock(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.FencedCodeBlock)
	if entering {
		_, _ = w.WriteString(`<pre class="line-numbers"><code`)
		language := n.Language(source)
		if language != nil {
			_, _ = w.WriteString(` class="language-`)
			r.Writer.Write(w, language)
			_, _ = w.WriteString(`"`)
		}
		_ = w.WriteByte('>')
		l := n.Lines().Len()
		for i := 0; i < l; i++ {
			line := n.Lines().At(i)
			r.Writer.RawWrite(w, line.Value(source))
		}
	} else {
		_, _ = w.WriteString("</code></pre>\n")
	}
	return ast.WalkContinue, nil
}

func (r *customStyleRenderer) renderTaskCheckBox(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}
	n := node.(*east.TaskCheckBox)

	if n.IsChecked {
		_, _ = w.WriteString(`<input checked="" disabled="" type="checkbox"`)
	} else {
		_, _ = w.WriteString(`<input disabled="" type="checkbox"`)
	}
	if r.XHTML {
		_, _ = w.WriteString(" /> ")
	} else {
		_, _ = w.WriteString("> ")
	}
	return ast.WalkContinue, nil
}

type customStyle struct {
}

func NewCustomStyle() goldmark.Extender {
	return &customStyle{}
}

func (e *customStyle) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithASTTransformers(
			util.Prioritized(NewCustomASTTransformer(), 999),
		),
	)
	m.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(NewCustomStyleRenderer(), 999),
		),
	)
}
