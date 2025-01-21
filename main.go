package main

import (
	"fmt"

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
	db := util.InitDbClient()
	swaggerUiPath := fmt.Sprintf("%s/swagger/*any", baseSwaggerPath)
	app.GET(swaggerUiPath, ginSwagger.WrapHandler(swaggerfiles.Handler))
	appInfo := fmt.Sprintf("0.0.0.0:%s", port)

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

	app.Run(appInfo) // listen and serve on 0.0.0.0:8080
}
