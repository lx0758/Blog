package html_vo

type PaginationVO[T interface{}] struct {
	Name     string
	Path     string
	PageNum  int
	PageSize int
	Size     int
	Total    int
	List     []T
}
