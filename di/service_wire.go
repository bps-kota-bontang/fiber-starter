//go:build wireinject
// +build wireinject

package di

import (
	"github.com/bps-kota-bontang/fiber-starter/services"

	"github.com/google/wire"
)

// Wire Set
var ServiceSet = wire.NewSet(
	services.NewUserService,
	RepositorySet,
)
