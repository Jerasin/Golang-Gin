package routes

import (
	"github.com/Jerasin/controllers/userController"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUserRoutes(router *gin.RouterGroup, db *gorm.DB) {
	controller := userController.DBController{Database: db}

	router.GET("users", controller.GetUsers)         // GET
	router.GET("user/:id", controller.GetUserById)   // GET BY ID
	router.POST("user", controller.CreateUser)       // POST
	router.PUT("user/:id", controller.UpdateUser)    // PUT
	router.DELETE("user/:id", controller.DeleteUser) // DELETE
}
