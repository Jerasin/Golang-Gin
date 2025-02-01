// go:build wireinject
// go:build wireinject
//go:build wireinject
// +build wireinject

package module

import (
	"github.com/Jerasin/app/repository"
	"github.com/Jerasin/app/util"
	"github.com/google/wire"
)

var db = wire.NewSet(util.InitDbClient)

var baseRepoSet = wire.NewSet(repository.BaseRepositoryInit,
	wire.Bind(new(repository.BaseRepositoryInterface), new(*repository.BaseRepository)),
)
