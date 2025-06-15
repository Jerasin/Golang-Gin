package response

import "github.com/Jerasin/app/model"

type User struct {
	model.BaseModel
	Username string `json:"username"`
	// Password string `json:"password"`
	Fullname   string `gorm:"unique;not null" json:"fullname,omitempty"`
	Avatar     string `json:"avatar"`
	RoleInfoID uint   `json:"roleId" gorm:"column:role_info_id"`
	Email      string `json:"email"`
	IsActive   bool   `gorm:"column:is_active" json:"isActive"`
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
