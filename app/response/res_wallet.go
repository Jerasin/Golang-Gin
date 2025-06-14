package response

type WalletDetail struct {
	Name  string `json:"name"`
	Token string `json:"token"`
	Uuid  string `json:"uuid"`
	// UserID uint   `json:"user_id"`
	User User `json:"user"`
}

type Wallet struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Token  string `json:"token"`
	Uuid   string `json:"uuid"`
	UserID uint   `json:"user_id"`
	Value  int    `json:"value"`
}

type WalletPagination struct {
	PaginationResponse
	Data []Wallet `json:"data"`
}
