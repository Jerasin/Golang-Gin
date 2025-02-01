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

var productSvcSet = wire.NewSet(service.ProductServiceInit,
	wire.Bind(new(service.ProductServiceInterface), new(*service.ProductServiceModel)),
)

var productCtrlSet = wire.NewSet(controller.ProductControllerInit,
	wire.Bind(new(controller.ProductControllerInterface), new(*controller.ProductController)),
)

var productRepoSet = wire.NewSet(repository.ProductRepositoryInit,
	wire.Bind(new(repository.ProductRepositoryInterface), new(*repository.ProductRepository)),
)

type ProductModule struct {
	ProductRepo repository.ProductRepositoryInterface
	ProductSvc  service.ProductServiceInterface
	ProductCtrl controller.ProductControllerInterface
}

func NewProductModule(productRepo repository.ProductRepositoryInterface,
	productService service.ProductServiceInterface,
	productCtrl controller.ProductControllerInterface,
) *ProductModule {
	return &ProductModule{
		ProductRepo: productRepo,
		ProductSvc:  productService,
		ProductCtrl: productCtrl,
	}
}

func ProductModuleInit() *ProductModule {
	wire.Build(NewProductModule, productSvcSet, productRepoSet, productCtrlSet, baseRepoSet, db)
	return nil
}
