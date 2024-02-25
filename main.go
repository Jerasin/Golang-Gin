package main

import (
	"fmt"

	"github.com/Jerasin/app/config"
	"github.com/gin-gonic/gin"
)

func initMain() {
	config.EnvConfig()
	config.InitDbClient()
	// config.InitLog()
}

func main() {
	initMain()
	PORT := config.GetEnv("PORT", "3000")
	// init := config.Init()
	// app := router.Init(init)

	app := gin.Default()

	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "HELLO GOLANG RESTFUL API.",
		})
	})

	appInfo := fmt.Sprintf("0.0.0.0:%s", PORT)
	fmt.Println("appInfo", appInfo)
	app.Run(appInfo) // listen and serve on 0.0.0.0:8080
}
