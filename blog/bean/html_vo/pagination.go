package html_vo

type PaginationVO[T interface{}] struct {
	PageNum  int
	PageSize int
	Size     int
	Total    int
	List     []T
}
