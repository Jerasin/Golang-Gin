package userService

import (
	"github.com/Jerasin/models"
	"gorm.io/gorm"
)

type UserServiceInterface interface {
	GetUser(id int) (models.User, error)
	GetUsers() ([]models.User, error)
}

func (db *userServices) GetUsers() ([]models.User, error) {
	var users []models.User
	result := db.Database.Find(&users)

	if result.Error != nil {
		return users, result.Error
	}
	return users, result.Error
}

func (db *userServices) GetUser(id int) (models.User, error) {
	var user models.User
	result := db.Database.First(&user, id)

	if result.Error != nil {
		return user, result.Error
	}

	return user, result.Error
}

// func (db *userServices) CreateUser(body models.User) (models.User, error) {
// 	var user models.User
// 	result := db.Database.First(&user, id)

// 	if result.Error != nil {
// 		return user, result.Error
// 	}

// 	return user, result.Error
// }

func UserService(db *gorm.DB) UserServiceInterface {
	return &userServices{
		Database: db,
	}
}
