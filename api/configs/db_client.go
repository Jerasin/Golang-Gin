package configs

import (
	"fmt"

	"github.com/Jerasin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}

func InitDbClient() *gorm.DB {
	DB_HOST := GetEnv("DB_HOST")
	DB_NAME := GetEnv("DB_NAME")
	DB_USER := GetEnv("DB_USER")
	DB_PASSWORD := GetEnv("DB_PASSWORD")

	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_HOST, DB_NAME)
	n, err := fmt.Printf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local\n", DB_USER, DB_PASSWORD, DB_HOST, DB_NAME)
	fmt.Println("n", n)
	fmt.Println("err", err)
	fmt.Println("mysqlInfo", mysqlInfo)

	db, err := gorm.Open(mysql.Open(mysqlInfo), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Init Db")

	// Migrate the schema
	autoMigrate(db)

	return db
}
