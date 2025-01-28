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

var RoleInfoSvcSet = wire.NewSet(service.RoleInfoServiceInit,
	wire.Bind(new(service.RoleInfoServiceInterface), new(*service.RoleInfoServiceModel)),
)

var RoleInfoCtrlSet = wire.NewSet(controller.RoleInfoControllerInit,
	wire.Bind(new(controller.RoleInfoControllerInterface), new(*controller.RoleInfoController)),
)

type RoleInfoModule struct {
	RoleInfoSvc  service.RoleInfoServiceInterface
	RoleInfoCtrl controller.RoleInfoControllerInterface
}

func NewRoleInfoModule(
	RoleInfoService service.RoleInfoServiceInterface,
	RoleInfoCtrl controller.RoleInfoControllerInterface,
) *RoleInfoModule {
	return &RoleInfoModule{
		RoleInfoSvc:  RoleInfoService,
		RoleInfoCtrl: RoleInfoCtrl,
	}
}

func RoleInfoModuleInit() *RoleInfoModule {
	wire.Build(NewRoleInfoModule, RoleInfoSvcSet, RoleInfoCtrlSet, baseRepoSet, db)
	return nil
}
