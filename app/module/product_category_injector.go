// go:build wireinject
// go:build wireinject
//go:build wireinject
// +build wireinject

package module

import (
	"github.com/Jerasin/app/controller"
	"github.com/Jerasin/app/repository"
	"github.com/Jerasin/app/service"
	"github.com/google/wire"
)

var productCategorySvcSet = wire.NewSet(service.ProductCategoryServiceInit,
	wire.Bind(new(service.ProductCategoryServiceInterface), new(*service.ProductCategoryServiceModel)),
)

var productCategoryCtrlSet = wire.NewSet(controller.ProductCategoryControllerInit,
	wire.Bind(new(controller.ProductCategoryControllerInterface), new(*controller.ProductCategoryController)),
)

var productCategoryRepoSet = wire.NewSet(repository.ProductCategoryRepositoryInit,
	wire.Bind(new(repository.ProductCategoryRepositoryInterface), new(*repository.ProductCategoryRepository)),
)

type ProductCategoryModule struct {
	ProductCategoryCtrl controller.ProductCategoryControllerInterface
	ProductCategorySvc  service.ProductCategoryServiceInterface
	ProductCategoryRepo repository.ProductCategoryRepositoryInterface
}

func NewProductCategoryModule(
	productCategoryCtrl controller.ProductCategoryControllerInterface,
	productCategorySvc service.ProductCategoryServiceInterface,
	productCategoryRepo repository.ProductCategoryRepositoryInterface,
) *ProductCategoryModule {
	return &ProductCategoryModule{
		ProductCategoryCtrl: productCategoryCtrl,
		ProductCategorySvc:  productCategorySvc,
		ProductCategoryRepo: productCategoryRepo,
	}
}

func ProductCategoryModuleInit() *ProductCategoryModule {
	wire.Build(NewProductCategoryModule, productCategoryCtrlSet, productCategorySvcSet, productCategoryRepoSet, db, baseRepoSet)
	return nil
}
