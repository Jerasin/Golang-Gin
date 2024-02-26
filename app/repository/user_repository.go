package repository

import (
	"github.com/Jerasin/app/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAllUser() ([]model.User, error)
	FindOneUser(condition model.User) (model.User, error)
	FindUserById(id int) (model.User, error)
	Save(user *model.User) (model.User, error)
	DeleteUserById(id int) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (u UserRepositoryImpl) FindAllUser() ([]model.User, error) {
	var users []model.User

	var err = u.db.Find(&users).Error
	if err != nil {
		log.Error("Got an error finding all couples. Error: ", err)
		return nil, err
	}

	return users, nil
}

func (u UserRepositoryImpl) FindOneUser(condition model.User) (model.User, error) {
	// var user model.User

	var err = u.db.First(&condition).Error
	if err != nil {
		log.Error("Got an error finding One couples. Error: ", err)
		return condition, err
	}

	return condition, nil
}

func (u UserRepositoryImpl) FindUserById(id int) (model.User, error) {
	var user model.User
	err := u.db.Preload("Role").First(&user, id).Error
	if err != nil {
		log.Error("Got and error when find user by id. Error: ", err)
		return model.User{}, err
	}
	return user, nil
}

func (u UserRepositoryImpl) Save(user *model.User) (model.User, error) {
	var err = u.db.Save(user).Error
	if err != nil {
		log.Error("Got an error when save user. Error: ", err)
		return model.User{}, err
	}
	return *user, nil
}

func (u UserRepositoryImpl) DeleteUserById(id int) error {
	err := u.db.Delete(&model.User{}, id).Error
	if err != nil {
		log.Error("Got an error when delete user. Error: ", err)
		return err
	}
	return nil
}

func UserRepositoryInit(db *gorm.DB) *UserRepositoryImpl {
	db.AutoMigrate(&model.User{})
	return &UserRepositoryImpl{
		db: db,
	}
}
