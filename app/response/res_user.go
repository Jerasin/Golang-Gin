package response

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	// Password string `json:"password"`
	FullName string `json:"fullName"`
	Avatar   string `json:"avatar"`
}

type UserPagination struct {
	PaginationResponse
	Data []User `json:"data"`
}
