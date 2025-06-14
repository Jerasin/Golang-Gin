package response

import "github.com/Jerasin/app/model"

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	// Password string `json:"password"`
	FullName   string `json:"fullName"`
	Avatar     string `json:"avatar"`
	RoleInfoID uint   `json:"userId"`
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
	Id       int    `json:"id"`
	Username string `json:"username"`
	// Password string `json:"password"`
	FullName   string       `json:"fullName"`
	Avatar     string       `json:"avatar"`
	RoleInfoID uint         `json:"userId"`
	RoleInfo   UserRoleInfo `json:"userRole"`
}
