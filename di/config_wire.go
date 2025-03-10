//go:build wireinject
// +build wireinject

package di

import (
	"github.com/bps-kota-bontang/fiber-starter/config"
	"github.com/google/wire"
)

// Wire Set for Database
var ConfigSet = wire.NewSet(config.LoadConfigDatabase)
