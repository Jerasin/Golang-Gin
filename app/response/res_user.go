package response

import "github.com/Jerasin/app/model"

type User struct {
	model.BaseModel
	Username string `json:"username"`
	// Password string `json:"password"`
	FullName   string `json:"fullname"`
	Avatar     string `json:"avatar"`
	RoleInfoID uint   `json:"userId"`
	Email      string `json:"email"`
}

type UserPagination struct {
	PaginationResponse
	Data []User `json:"data"`
}

type UserRoleInfo struct {
	Name            string                 `gorm:"unique;not null" json:"name"`
	Description     string                 `json:"description"`
	PermissionInfos []model.PermissionInfo `json:"permissionInfos"`
}

type UserInfo struct {
	model.BaseModel
	Email    string `json:"email"`
	Username string `json:"username"`
	// Password string `json:"password"`
	FullName   string       `json:"fullname"`
	Avatar     string       `json:"avatar"`
	RoleInfoID uint         `json:"userId"`
	RoleInfo   UserRoleInfo `json:"userRole"`
}
