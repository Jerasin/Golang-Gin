package request

import "time"

type UserRequest struct {
	Username   string `json:"username" binding:"required" example:"admin"`
	Password   string `json:"password" binding:"required" example:"1234"`
	Fullname   string `json:"fullname" binding:"required" example:"admin test"`
	Email      string `json:"email" binding:"required" example:"admin@gmail.com"`
	Avatar     string `json:"avatar" example:"admin"`
	RoleInfoID uint   `json:"roleInfoId" example:"1"`
}

type UpdateUser struct {
	Username string `json:"username" example:"admin"`
	// Password string `json:"password" example:"1234"`
	Fullname  string    `json:"fullname" example:"admin test"`
	Avatar    string    `json:"avatar" example:"admin"`
	UpdatedAt time.Time `json:"updated_at"`
}
