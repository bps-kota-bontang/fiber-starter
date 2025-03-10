//go:build wireinject
// +build wireinject

package di

import (
	"github.com/bps-kota-bontang/fiber-starter/database"
	"github.com/google/wire"
)

// Wire Set for Database
var DBSet = wire.NewSet(database.ConnectDB, ConfigSet)
