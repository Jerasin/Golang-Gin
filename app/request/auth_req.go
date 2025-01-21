package request

type LoginRequest struct {
	Username string `json:"username" binding:"required" example:"admin"`
	Password string `json:"password" binding:"required" example:"1234"`
}

type TokenReqBody struct {
	RefreshToken string `json:"refresh_token"`
}
