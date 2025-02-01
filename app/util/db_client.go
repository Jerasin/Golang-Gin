package util

import (
	"fmt"

	"github.com/Jerasin/app/config"
	"github.com/Jerasin/app/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func autoMigrate(db *gorm.DB) {
	fmt.Println("autoMigrate")
	db.AutoMigrate(&model.ProductCategory{}, &model.Product{}, &model.PermissionInfo{}, &model.RoleInfo{}, &model.User{}, &model.Order{}, &model.OrderDetail{}, &model.Wallet{})
}

func InitDbClient() *gorm.DB {
	DB_HOST := config.GetEnv("DB_HOST", "localhost:3306")
	DB_NAME := config.GetEnv("DB_NAME", "api")
	DB_USER := config.GetEnv("DB_USER", "api")
	DB_PASSWORD := config.GetEnv("DB_PASSWORD", "123456")
	DB_LOG_ENABLE := config.GetEnv("DB_LOG_ENABLE", "false")

	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_HOST, DB_NAME)
	_, err := fmt.Printf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local\n", DB_USER, DB_PASSWORD, DB_HOST, DB_NAME)
	// fmt.Println("n", n)
	// fmt.Println("err", err)

	if err != nil {
		panic("failed to mapping string")
	}

	dbLogLevel := logger.Info

	if DB_LOG_ENABLE == "true" {
		dbLogLevel = logger.Info
	}

	fmt.Println("mysqlInfo", mysqlInfo)
	db, err := gorm.Open(mysql.Open(mysqlInfo), &gorm.Config{TranslateError: true, Logger: logger.Default.LogMode(dbLogLevel), SkipDefaultTransaction: true})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Init Db")

	// Migrate the schema
	autoMigrate(db)

	return db
}
