package api_vo

type Pagination[T interface{}] struct {
	PageNum  int
	PageSize int
	Pages    int
	Total    int
	List     []T
}
