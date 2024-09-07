package markdown

type Catalogue struct {
	Title  string
	Anchor string
	Childs []Catalogue
}

func RenderDescription(markdown string) string {
	return markdown
}

func RenderContent(markdown string) (html string, catalogues []Catalogue) {
	html = markdown
	catalogues = []Catalogue{}
	return
}
