package config

import (
	"fmt"

	"github.com/Jerasin/app/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
}

func InitDbClient() *gorm.DB {
	DB_HOST := GetEnv("DB_HOST", "localhost:3306")
	DB_NAME := GetEnv("DB_NAME", "api")
	DB_USER := GetEnv("DB_USER", "api")
	DB_PASSWORD := GetEnv("DB_PASSWORD", "123456")

	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_HOST, DB_NAME)
	_, err := fmt.Printf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local\n", DB_USER, DB_PASSWORD, DB_HOST, DB_NAME)
	// fmt.Println("n", n)
	// fmt.Println("err", err)

	if err != nil {
		panic("failed to mapping string")
	}

	fmt.Println("mysqlInfo", mysqlInfo)
	db, err := gorm.Open(mysql.Open(mysqlInfo), &gorm.Config{TranslateError: true})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Init Db")

	// Migrate the schema
	autoMigrate(db)

	return db
}
