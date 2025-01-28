//go:build wireinject
// +build wireinject

// go:build wireinject
// go:build wireinject
package module

import (
	"github.com/Jerasin/app/controller"
	"github.com/Jerasin/app/service"
	"github.com/google/wire"
)

var WalletSvcSet = wire.NewSet(service.WalletServiceInit,
	wire.Bind(new(service.WalletServiceInterface), new(*service.WalletServiceModel)),
)

var WalletCtrlSet = wire.NewSet(controller.WalletControllerInit,
	wire.Bind(new(controller.WalletControllerInterface), new(*controller.WalletController)),
)

type WalletModule struct {
	WalletSvc  service.WalletServiceInterface
	WalletCtrl controller.WalletControllerInterface
}

func NewWalletModule(
	WalletService service.WalletServiceInterface,
	WalletCtrl controller.WalletControllerInterface,
) *WalletModule {
	return &WalletModule{
		WalletSvc:  WalletService,
		WalletCtrl: WalletCtrl,
	}
}

func WalletModuleInit() *WalletModule {
	wire.Build(NewWalletModule, WalletSvcSet, WalletCtrlSet, baseRepoSet, db)
	return nil
}
