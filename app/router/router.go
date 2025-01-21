package router

import (
	"github.com/Jerasin/app/middleware"
	"github.com/Jerasin/app/module"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type BaseModuleInit struct {
	UserModule            *module.UserModule
	AuthModule            *module.AuthModule
	ProductCategoryModule *module.ProductCategoryModule
	ProductModule         *module.ProductModule
	OrderModule           *module.OrderModule
	PermissionInfoModule  *module.PermissionInfoModule
	RoleInfoModule        *module.RoleInfoModule
	WalletModule          *module.WalletModule
}

func NewBaseModule() BaseModuleInit {
	userInit := module.UserModuleInit()
	authInit := module.AuthModuleInit()
	productCategoryInit := module.ProductCategoryModuleInit()
	productInit := module.ProductModuleInit()
	orderInit := module.OrderModuleInit()
	permissionInfoInit := module.PermissionInfoModuleInit()
	roleInfoInit := module.RoleInfoModuleInit()
	walletInit := module.WalletModuleInit()

	return BaseModuleInit{
		UserModule:            userInit,
		AuthModule:            authInit,
		ProductCategoryModule: productCategoryInit,
		ProductModule:         productInit,
		OrderModule:           orderInit,
		PermissionInfoModule:  permissionInfoInit,
		RoleInfoModule:        roleInfoInit,
		WalletModule:          walletInit,
	}
}

func RouterInit(init BaseModuleInit) *gin.Engine {
	router := gin.New()

	router.SetTrustedProxies(nil)
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	api := router.Group("/api")

	user := api.Group("/users")
	user.Use(middleware.AuthorizeJwt())
	user.GET("", init.UserModule.UserCtrl.GetAllUsers)
	user.POST("", init.UserModule.UserCtrl.CreateUser)
	user.GET("/:userID", init.UserModule.UserCtrl.GetUserById)
	user.PUT("/:userID", init.UserModule.UserCtrl.UpdateUserData)
	user.DELETE("/:userID", init.UserModule.UserCtrl.DeleteUser)

	auth := api.Group("/auth")

	auth.POST("/register", init.AuthModule.AuthCtrl.Register)
	auth.POST("/login", init.AuthModule.AuthCtrl.Login)
	auth.POST("/refresh/token", init.AuthModule.AuthCtrl.RefreshToken)

	product := api.Group("/products")
	product.POST("", init.ProductModule.ProductCtrl.CreateProduct)
	product.GET("", init.ProductModule.ProductCtrl.GetAllProducts)
	product.GET("/:productID", init.ProductModule.ProductCtrl.GetProductById)
	product.PUT("/:productID", init.ProductModule.ProductCtrl.UpdateProductData)
	product.DELETE("/:productID", init.ProductModule.ProductCtrl.DeleteProduct)

	productCategory := product.Group("/categories")
	productCategory.Use(middleware.AuthorizeJwt())
	productCategory.POST("", init.ProductCategoryModule.ProductCategoryCtrl.CreateProductCategory)
	productCategory.GET("", init.ProductCategoryModule.ProductCategoryCtrl.GetListProductCategory)
	productCategory.GET("/:productCategoryID", init.ProductCategoryModule.ProductCategoryCtrl.GetProductCategoryById)
	productCategory.PUT("/:productCategoryID", init.ProductCategoryModule.ProductCategoryCtrl.UpdateProductCategoryData)

	order := api.Group("/orders")
	order.Use(middleware.AuthorizeJwt())
	order.POST("", init.OrderModule.OrderCtrl.CreateOrder)
	order.GET("", init.OrderModule.OrderCtrl.GetAllProducts)

	permissionInfo := api.Group("/permission_infos")
	permissionInfo.Use(middleware.AuthorizeJwt())
	permissionInfo.POST("", init.PermissionInfoModule.PermissionInfoCtrl.CreatePermissionInfo)
	permissionInfo.GET("", init.PermissionInfoModule.PermissionInfoCtrl.GetListPermissionInfo)
	permissionInfo.GET("/:permissionInfoID", init.PermissionInfoModule.PermissionInfoCtrl.GetPermissionInfoById)
	permissionInfo.PUT("/:permissionInfoID", init.PermissionInfoModule.PermissionInfoCtrl.UpdatePermissionInfoData)
	permissionInfo.DELETE("/:permissionInfoID", init.PermissionInfoModule.PermissionInfoCtrl.DeletePermissionInfo)

	role_info := api.Group("/role_infos")
	permissionInfo.Use(middleware.AuthorizeJwt())
	role_info.POST("", init.RoleInfoModule.RoleInfoCtrl.CreateRoleInfo)
	role_info.GET("", init.RoleInfoModule.RoleInfoCtrl.GetListRoleInfo)
	role_info.GET("/:roleInfoID", init.RoleInfoModule.RoleInfoCtrl.GetRoleInfoById)
	role_info.PUT("/:roleInfoID", init.RoleInfoModule.RoleInfoCtrl.UpdateRoleInfoData)
	role_info.DELETE("/:roleInfoID", init.RoleInfoModule.RoleInfoCtrl.DeleteRoleInfo)

	wallet := api.Group("/wallets")
	permissionInfo.Use(middleware.AuthorizeJwt())
	wallet.POST("", init.WalletModule.WalletCtrl.CreateWallet)
	wallet.GET("", init.WalletModule.WalletCtrl.GetListWallet)
	wallet.GET("/:walletID", init.WalletModule.WalletCtrl.GetWalletById)
	wallet.PUT("/:walletID", init.WalletModule.WalletCtrl.UpdateWalletData)
	wallet.DELETE("/:walletID", init.WalletModule.WalletCtrl.DeleteWallet)
	return router
}
