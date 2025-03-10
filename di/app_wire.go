//go:build wireinject
// +build wireinject

package di

import (
	"github.com/bps-kota-bontang/fiber-starter/app"

	"github.com/google/wire"
)

// Wire Set for App dependencies
var AppSet = wire.NewSet(app.NewFiberApp, HandlerSet)
