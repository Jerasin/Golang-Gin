package main

import (
	"fmt"

	"github.com/Jerasin/configs"
	"github.com/Jerasin/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Register struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Avatar   string `json:"avatar" binding:"required"`
}

func main() {

	PORT := configs.GetEnv("PORT")
	app := gin.Default()
	app.Use(cors.Default())
	api := app.Group("/api")

	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "HELLO GOLANG RESTFUL API.",
		})
	})

	db := configs.InitDbClient()
	routes.SetUserRoutes(api, db)

	// var json Register

	// app.POST("/register", func(c *gin.Context) {
	// 	if err := c.ShouldBindJSON(&json); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// 	c.JSON(200, gin.H{
	// 		"message": json,
	// 		"status":  200,
	// 	})
	// })

	appInfo := fmt.Sprintf("0.0.0.0:%s", PORT)
	fmt.Println("appInfo", appInfo)
	app.Run(appInfo) // listen and serve on 0.0.0.0:8080
}
