package html_vo

type ThemeVO struct {
	Blog            BlogVO
	Page            string
	PageTitle       string
	PageDescription string
	PageKeywords    string
	PageUrl         string
	PageData        any
	Catalogues      []CatalogueVO
}
