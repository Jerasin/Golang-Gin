package response

type RoleInfo struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type RoleInfoPagination struct {
	PaginationResponse
	Data []RoleInfo `json:"data"`
}
