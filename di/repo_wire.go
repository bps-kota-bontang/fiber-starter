//go:build wireinject
// +build wireinject

package di

import (
	"github.com/bps-kota-bontang/fiber-starter/repositories"

	"github.com/google/wire"
)

// Wire Set
var RepositorySet = wire.NewSet(
	repositories.NewUserRepository,
	DBSet,
)
