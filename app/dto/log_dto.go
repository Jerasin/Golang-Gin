package dto

type (
	LoginDtoRequest struct {
		Username string `json:"username" binding:"required" example:"admin" validate:"min=1"`
		Password string `json:"password" binding:"required" example:"1234" validate:"min=1"`
	}

	LoginDtoResponse struct {
		BaseDtoResponse `json:"base_response"`
		Data            LoginDtoDataResponse `json:"data"`
	}

	LoginDtoDataResponse struct {
		RefreshToken string `json:"refresh_token" binding:"required" example:"admin" validate:"min=1"`
		Token        string `json:"token" binding:"required" example:"1234" validate:"min=1"`
	}
)
