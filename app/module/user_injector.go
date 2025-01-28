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

var userServiceSet = wire.NewSet(service.UserServiceInit,
	wire.Bind(new(service.UserServiceInterface), new(*service.UserServiceModel)),
)

var userCtrlSet = wire.NewSet(controller.UserControllerInit,
	wire.Bind(new(controller.UserControllerInterface), new(*controller.UserController)),
)

var userRepoSet = wire.NewSet(repository.UserRepositoryInit,
	wire.Bind(new(repository.UserRepositoryInterface), new(*repository.UserRepository)),
)

type UserModule struct {
	UserRepo repository.UserRepositoryInterface
	UserSvc  service.UserServiceInterface
	UserCtrl controller.UserControllerInterface
}

func NewUserModule(userRepo repository.UserRepositoryInterface,
	userService service.UserServiceInterface,
	userCtrl controller.UserControllerInterface,
) *UserModule {
	return &UserModule{
		UserRepo: userRepo,
		UserSvc:  userService,
		UserCtrl: userCtrl,
	}
}

func UserModuleInit() *UserModule {
	wire.Build(NewUserModule, userServiceSet, userRepoSet, db, userCtrlSet, baseRepoSet)
	return nil
}
