package api_vo

type PaginationVO[T interface{}] struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	Size     int `json:"size"`
	Total    int `json:"total"`
	List     []T `json:"list"`
}
