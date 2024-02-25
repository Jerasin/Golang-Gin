// go:build wireinject
//go:build wireinject
// +build wireinject

package config

import (
	"github.com/Jerasin/app/controller"
	"github.com/Jerasin/app/repository"
	"github.com/Jerasin/app/service"

	"github.com/google/wire"
)

var db = wire.NewSet(InitDbClient)

var userServiceSet = wire.NewSet(service.UserServiceInit,
	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
)

var userCtrlSet = wire.NewSet(controller.UserControllerInit,
	wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)),
)

var userRepoSet = wire.NewSet(repository.UserRepositoryInit,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
)

func Init() *Initialization {
	wire.Build(NewInitialization, db, userCtrlSet, userServiceSet, userRepoSet)
	return nil
}
