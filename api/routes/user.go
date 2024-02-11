package routes

import (
	"github.com/Jerasin/controllers/userController"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUserRoutes(router *gin.RouterGroup, db *gorm.DB) {
	controller := userController.DBController{Database: db}

	router.GET("users", controller.GetUsers) // GET
	// router.GET("collections/:id", ctrls.GetCollectionById)   // GET BY ID
	router.POST("user", controller.CreateUser) // POST
	// router.PATCH("collections", ctrls.UpdateCollection)      // PATCH
	// router.DELETE("collections/:id", ctrls.DeleteCollection) // DELETE
}
