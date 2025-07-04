package request

type UserRequest struct {
	Username   string `json:"username" binding:"required" example:"admin"`
	Password   string `json:"password" binding:"required" example:"1234"`
	Fullname   string `json:"fullname" binding:"required" example:"admin test"`
	Email      string `json:"email" binding:"required" example:"admin@gmail.com"`
	Avatar     string `json:"avatar" example:"admin"`
	RoleInfoID uint   `json:"roleInfoId" binding:"required" example:"1"`
}

type UpdateUser struct {
	Username string `json:"username" example:"admin"`
	Fullname string `json:"fullname" example:"admin test"`
	Avatar   string `json:"avatar" example:"admin"`
	IsActive bool   `json:"isActive"`
}
