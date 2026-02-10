package next

import (
	east "github.com/yuin/goldmark/extension/ast"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
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

// Transform
// 1. 给标题加锚点
// 2. 给链接加新标签页打开
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

type customRenderer struct {
	html.Config
}

func NewCustomRenderer(opts ...html.Option) renderer.NodeRenderer {
	r := &customRenderer{
		Config: html.NewConfig(),
	}
	for _, opt := range opts {
		opt.SetHTMLOption(&r.Config)
	}
	return r
}

func (r *customRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindListItem, r.renderListItem)
	reg.Register(ast.KindFencedCodeBlock, r.renderFencedCodeBlock)
}

// 优化 ListItem, 渲染 TaskCheckBox 的父容器为 div
func (r *customRenderer) renderListItem(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	itemTag := "li"
	if tbn, ok := n.FirstChild().(*ast.TextBlock); ok {
		if _, ok := tbn.FirstChild().(*east.TaskCheckBox); ok {
			itemTag = "div"
		}
	}

	if entering {
		if n.Attributes() != nil {
			_, _ = w.WriteString("<" + itemTag)
			html.RenderAttributes(w, n, html.ListItemAttributeFilter)
			_ = w.WriteByte('>')
		} else {
			_, _ = w.WriteString("<" + itemTag + ">")
		}
		fc := n.FirstChild()
		if fc != nil {
			if _, ok := fc.(*ast.TextBlock); !ok {
				_ = w.WriteByte('\n')
			}
		}
	} else {
		_, _ = w.WriteString("</" + itemTag + ">\n")
	}
	return ast.WalkContinue, nil
}

// 适配前台代码高亮框架
func (r *customRenderer) renderFencedCodeBlock(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
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

type custom struct {
}

func NewCustom() goldmark.Extender {
	return &custom{}
}

func (e *custom) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithASTTransformers(
			util.Prioritized(NewCustomASTTransformer(), 999),
		),
	)
	m.Renderer().AddOptions(
		// see:
		//	html.Config.HardWraps
		//	html.SetOption
		renderer.WithOption("HardWraps", true),
		renderer.WithOption("EastAsianLineBreaks", html.EastAsianLineBreaksCSS3Draft),
		renderer.WithNodeRenderers(
			util.Prioritized(NewCustomRenderer(), 999),
		),
	)
}
