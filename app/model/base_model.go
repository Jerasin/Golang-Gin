package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint            `gorm:"primary_key" json:"id"`
	CreatedAt time.Time       `gorm:"not null" json:"createdAt,omitempty"`
	UpdatedAt time.Time       `gorm:"not null" json:"updatedAt,omitempty"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
}
