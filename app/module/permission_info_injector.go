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

var PermissionInfoSvcSet = wire.NewSet(service.PermissionInfoServiceInit,
	wire.Bind(new(service.PermissionInfoServiceInterface), new(*service.PermissionInfoServiceModel)),
)

var PermissionInfoCtrlSet = wire.NewSet(controller.PermissionInfoControllerInit,
	wire.Bind(new(controller.PermissionInfoControllerInterface), new(*controller.PermissionInfoController)),
)

type PermissionInfoModule struct {
	PermissionInfoSvc  service.PermissionInfoServiceInterface
	PermissionInfoCtrl controller.PermissionInfoControllerInterface
}

func NewPermissionInfoModule(
	PermissionInfoService service.PermissionInfoServiceInterface,
	PermissionInfoCtrl controller.PermissionInfoControllerInterface,
) *PermissionInfoModule {
	return &PermissionInfoModule{
		PermissionInfoSvc:  PermissionInfoService,
		PermissionInfoCtrl: PermissionInfoCtrl,
	}
}

func PermissionInfoModuleInit() *PermissionInfoModule {
	wire.Build(NewPermissionInfoModule, PermissionInfoSvcSet, PermissionInfoCtrlSet, baseRepoSet, db)
	return nil
}
