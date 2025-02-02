package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Jerasin/app/config"
	"github.com/Jerasin/app/router"
	"github.com/Jerasin/app/util"
	docs "github.com/Jerasin/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@host		localhost:3000
//	@BasePath	/api/v1

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	fmt.Println("Main Start")

	config.EnvConfig()
	port := config.GetEnv("PORT", "3000")
	baseSwaggerPath := config.GetEnv("BASE_SWAGGER_PATH", "/api")
	baseModule := router.NewBaseModule()
	app := router.RouterInit(baseModule)
	docs.SwaggerInfo.BasePath = baseSwaggerPath
	swaggerUiPath := fmt.Sprintf("%s/swagger/*any", baseSwaggerPath)
	app.GET(swaggerUiPath, ginSwagger.WrapHandler(swaggerfiles.Handler))
	appInfo := fmt.Sprintf("0.0.0.0:%s", port)

	db := util.InitDbClient()
	initDataClient := util.InitDataClientInit(db)

	permissionInfos := initDataClient.InitPermissionInfo()
	initDataClient.InitRoleInfo(permissionInfos)
	initDataClient.InitUser()
	initDataClient.InitProductCategory()
	initDataClient.InitProduct()

	fmt.Println("swaggerUiPath", swaggerUiPath)
	fmt.Println("appInfo", appInfo)

	for _, item := range app.Routes() {
		println("method:", item.Method, "path:", item.Path)
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Current Directory:", dir)

	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		// ถ้า entry เป็น directory, จะพิมพ์ชื่อ folder
		if file.IsDir() {
			fmt.Println(file.Name()) // ชื่อ folder
		}
	}

	app.Run(appInfo) // listen and serve on 0.0.0.0:8080
}
