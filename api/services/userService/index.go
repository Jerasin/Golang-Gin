package userService

import "gorm.io/gorm"

// create database controller
type userServices struct {
	Database *gorm.DB
}
