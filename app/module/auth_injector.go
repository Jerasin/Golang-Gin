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

var authSvcSet = wire.NewSet(service.AuthServiceInit,
	wire.Bind(new(service.AuthServiceInterface), new(*service.AuthServiceModel)),
)

var authCtrlSet = wire.NewSet(controller.AuthControllerInit,
	wire.Bind(new(controller.AuthControllerInterface), new(*controller.AuthController)),
)

type AuthModule struct {
	AuthCtrl controller.AuthControllerInterface
	AuthSvc  service.AuthServiceInterface
	UserRepo repository.UserRepositoryInterface
	UserSvc  service.UserServiceInterface
}

func NewAuthModule(
	authCtrl controller.AuthControllerInterface,
	authSvc service.AuthServiceInterface,
	userRepo repository.UserRepositoryInterface,
	userSvc service.UserServiceInterface,
) *AuthModule {
	return &AuthModule{
		AuthSvc:  authSvc,
		AuthCtrl: authCtrl,
		UserRepo: userRepo,
		UserSvc:  userSvc,
	}
}

func AuthModuleInit() *AuthModule {
	wire.Build(NewAuthModule, authCtrlSet, authSvcSet, userRepoSet, userServiceSet, db, baseRepoSet)
	return nil
}
