package model

type PermissionInfo struct {
	BaseModel
	Name        string `gorm:"unique;not null" json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

func (PermissionInfo) TableName() string {
	return "permissionInfos"
}
