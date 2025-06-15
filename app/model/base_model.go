package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint            `gorm:"column:id;primary_key" json:"id"`
	CreatedAt time.Time       `gorm:"column:created_at;not null" json:"createdAt,omitempty"`
	UpdatedAt time.Time       `gorm:"column:updated_at;not null" json:"updatedAt,omitempty"`
	DeletedAt *gorm.DeletedAt `gorm:"column:deleted_at;index" json:"deletedAt,omitempty"`
}
