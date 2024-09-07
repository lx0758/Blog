package po

type Pagination[PO interface{}] struct {
	PageNum  int
	PageSize int
	Size     int
	Total    int
	List     []PO
}
