package dto

type (
	RoleInfoCreateRequest struct {
		Name        string `json:"name" binding:"required" validate:"min=1"`
		Description string `json:"description"  validate:"min=1"`
	}
)
