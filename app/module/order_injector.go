// go:build wireinject
// go:build wireinject
//go:build wireinject
// +build wireinject

package module

import (
	"github.com/Jerasin/app/controller"
	"github.com/Jerasin/app/service"
	"github.com/google/wire"
)

var OrderSvcSet = wire.NewSet(service.OrderServiceInit,
	wire.Bind(new(service.OrderServiceInterface), new(*service.OrderServiceModel)),
)

var OrderCtrlSet = wire.NewSet(controller.OrderControllerInit,
	wire.Bind(new(controller.OrderControllerInterface), new(*controller.OrderController)),
)

// var OrderRepoSet = wire.NewSet(repository.OrderRepositoryInit,
// 	wire.Bind(new(repository.OrderRepositoryInterface), new(*repository.OrderRepository)),
// )

// var OrderDetailRepoSet = wire.NewSet(repository.OrderDetailRepositoryInit,
// 	wire.Bind(new(repository.OrderDetailRepositoryInterface), new(*repository.OrderDetailRepository)),
// )

type OrderModule struct {
	// OrderRepo repository.OrderRepositoryInterface
	OrderSvc  service.OrderServiceInterface
	OrderCtrl controller.OrderControllerInterface
}

func NewOrderModule(
	// OrderRepo repository.OrderRepositoryInterface,
	OrderService service.OrderServiceInterface,
	OrderCtrl controller.OrderControllerInterface,
) *OrderModule {
	return &OrderModule{
		// OrderRepo: OrderRepo,
		OrderSvc:  OrderService,
		OrderCtrl: OrderCtrl,
	}
}

func OrderModuleInit() *OrderModule {
	wire.Build(NewOrderModule, OrderSvcSet, OrderCtrlSet, baseRepoSet, db)
	return nil
}
