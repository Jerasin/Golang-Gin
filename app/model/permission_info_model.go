package model

type PermissionInfo struct {
	BaseModel
	Name        string `gorm:"unique;not null"`
	Description string
}

func (PermissionInfo) TableName() string {
	return "permissionInfos"
}
