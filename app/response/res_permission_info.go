package response

type PermissionInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type PermissionInfoPagination struct {
	PaginationResponse
	Data []PermissionInfo `json:"data"`
}
