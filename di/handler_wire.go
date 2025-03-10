//go:build wireinject
// +build wireinject

package di

import (
	"github.com/bps-kota-bontang/fiber-starter/handlers"

	"github.com/google/wire"
)

// Wire Set
var HandlerSet = wire.NewSet(
	handlers.NewUserHandler,
	ServiceSet,
	ValidatorSet,
)
