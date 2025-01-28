package response

type ProductCategory struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProductCategoryPagination struct {
	PaginationResponse
	Data []ProductCategory `json:"data"`
}
